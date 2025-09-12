# 高并发系统设计与技术问答

## Q1: 千万QPS秒杀系统设计

### 极致架构设计 (1000万QPS)

#### 1. 分层架构
```
用户(亿级) -> 全球CDN -> 智能DNS -> 多级网关 -> 内存计算集群 -> 异步写入层
```

#### 2. 前端极致优化

**CDN + 边缘计算**
- 全球200+ CDN节点，就近接入
- 边缘计算：CDN节点直接处理秒杀逻辑
- 静态化：整个秒杀页面CDN缓存
- WebAssembly：前端高性能计算

**前端限流 + 预处理**
```javascript
// 客户端令牌桶限流
class TokenBucket {
    constructor(capacity = 10, refillRate = 1) {
        this.capacity = capacity;
        this.tokens = capacity;
        this.refillRate = refillRate;
        this.lastRefill = Date.now();
    }
    
    consume() {
        this.refill();
        if (this.tokens > 0) {
            this.tokens--;
            return true;
        }
        return false;
    }
}
```

#### 3. 网关层设计

**多级网关架构**
- L1网关：全局限流 + 黑名单 (10万QPS/实例)
- L2网关：业务路由 + 用户认证 (5万QPS/实例)  
- L3网关：精细化控制 + 熔断降级 (2万QPS/实例)

**超高并发网关实现**
```go
// 基于Golang + epoll的高性能网关
type Gateway struct {
    rateLimiter *RateLimiter  // 基于令牌桶
    circuitBreaker *CircuitBreaker
    loadBalancer *LoadBalancer
}

func (g *Gateway) Handle(req *Request) {
    // 1. 限流检查 - 纳秒级
    if !g.rateLimiter.Allow() {
        return fastReject() 
    }
    
    // 2. 热点检测 - 微秒级
    if isHotKey(req.GoodsId) {
        return routeToHotCluster(req)
    }
    
    // 3. 正常路由 - 微秒级
    return g.loadBalancer.Forward(req)
}
```

#### 4. 内存计算集群

**Redis Cluster + 一致性哈希**
```
Redis集群配置:
- 100个节点 * 16GB = 1.6TB内存
- 每节点10万QPS，总计1000万QPS
- 3副本，数据安全
```

**内存数据结构优化**
```lua
-- 原子性库存扣减 Lua脚本
local key = KEYS[1]
local delta = ARGV[1] 
local current = redis.call('GET', key)

if current == false then
    return -1  -- key不存在
end

current = tonumber(current)
if current >= tonumber(delta) then
    redis.call('DECRBY', key, delta)
    return current - delta
else 
    return -2  -- 库存不足
end
```

**分片策略**
```
商品分片: goods_id % 100 -> 对应Redis节点
用户分片: user_id % 100 -> 对应业务服务器
时间分片: timestamp % 10 -> 对应消息队列分区
```

#### 5. 核心服务集群

**服务配置**
- 1000台服务器 * 16核 * 4进程 = 64000个工作进程
- 每进程处理200QPS，总计1280万QPS能力
- 基于Golang/Rust实现，单机性能极致优化

**内存池 + 协程池**
```go
type SecKillService struct {
    memPool     *sync.Pool        // 内存池
    workerPool  *ants.Pool        // 协程池
    localCache  *freecache.Cache  // 本地缓存
}

func (s *SecKillService) ProcessSecKill(req *SecKillReq) {
    // 从内存池获取对象，避免GC
    ctx := s.memPool.Get().(*Context)
    defer s.memPool.Put(ctx)
    
    // 提交到协程池异步处理
    s.workerPool.Submit(func() {
        s.handleSecKill(ctx, req)
    })
}
```

#### 6. 极致秒杀流程

**0-拷贝设计**
```
请求 -> 内存解析 -> 内存计算 -> 内存响应 (全程无磁盘IO)
```

**毫秒级处理链路**
```
网关接收(0.1ms) -> 参数校验(0.1ms) -> Redis库存检查(0.5ms) -> 
业务逻辑(0.2ms) -> 异步写入(0.1ms) -> 响应返回(0.1ms) = 总计1.1ms
```

### Redis全内存秒杀方案

#### 1. Redis秒杀架构设计

**核心思想：将整个秒杀流程在Redis内存中完成**
```
请求 -> 网关 -> Redis Lua脚本 -> 返回结果 -> 异步持久化
```

**Redis集群配置**
```
Redis Cluster: 
- 200个节点 * 32GB = 6.4TB总内存
- 每节点支持5万QPS，总计1000万QPS
- 6主6从架构，保证高可用
```

#### 2. Lua脚本实现完整秒杀逻辑

**秒杀核心脚本**
```lua
-- 秒杀主逻辑 Lua脚本
local goodsKey = KEYS[1]           -- 商品库存key
local userSetKey = KEYS[2]         -- 已购买用户集合key  
local orderListKey = KEYS[3]       -- 订单列表key
local userId = ARGV[1]             -- 用户ID
local orderId = ARGV[2]            -- 订单ID
local timestamp = ARGV[3]          -- 时间戳

-- 1. 检查用户是否已购买
if redis.call('SISMEMBER', userSetKey, userId) == 1 then
    return {0, 'already_bought'}
end

-- 2. 检查并扣减库存
local stock = redis.call('GET', goodsKey)
if not stock or tonumber(stock) <= 0 then
    return {0, 'no_stock'}
end

-- 3. 原子性操作：扣库存 + 记录用户 + 生成订单
redis.call('DECR', goodsKey)
redis.call('SADD', userSetKey, userId)
redis.call('LPUSH', orderListKey, cjson.encode({
    order_id = orderId,
    user_id = userId,
    timestamp = timestamp,
    status = 'pending'
}))

-- 4. 设置订单过期时间(15分钟)
local orderKey = 'order:' .. orderId
redis.call('HSET', orderKey, 
    'user_id', userId,
    'goods_id', goodsKey,
    'status', 'pending',
    'created_at', timestamp
)
redis.call('EXPIRE', orderKey, 900)  -- 15分钟过期

return {1, 'success', orderId}
```

**库存初始化脚本**
```lua
-- 库存预热脚本
local goodsKey = KEYS[1]
local totalStock = ARGV[1]
local userSetKey = KEYS[2]
local orderListKey = KEYS[3]

-- 初始化库存
redis.call('SET', goodsKey, totalStock)
-- 清空历史数据
redis.call('DEL', userSetKey, orderListKey)

return 'OK'
```

#### 3. 一致性保证机制

**3.1 Redis层面一致性**

**原子性保证**
```go
// Go客户端调用
func (s *SecKillService) SecKill(goodsId, userId int64) (*SecKillResult, error) {
    luaScript := redis.NewScript(SECKILL_LUA_SCRIPT)
    
    keys := []string{
        fmt.Sprintf("stock:%d", goodsId),      // 库存key
        fmt.Sprintf("buyers:%d", goodsId),     // 购买者集合
        fmt.Sprintf("orders:%d", goodsId),     // 订单列表
    }
    
    args := []interface{}{
        userId,
        generateOrderId(),
        time.Now().Unix(),
    }
    
    // Lua脚本保证原子性
    result, err := luaScript.Run(s.redisClient, keys, args...).Result()
    return parseSecKillResult(result)
}
```

**集群一致性：使用Hash Tag**
```go
// 确保相关数据在同一Redis节点
stockKey := fmt.Sprintf("stock:{%d}", goodsId)      // Hash tag: goodsId
buyersKey := fmt.Sprintf("buyers:{%d}", goodsId)    // 同一节点
ordersKey := fmt.Sprintf("orders:{%d}", goodsId)    // 同一节点
```

**3.2 Redis与数据库一致性**

**方案一：异步最终一致性**
```go
// Redis操作成功后异步写DB
func (s *SecKillService) asyncPersist(order *Order) {
    // 发送到消息队列
    s.mqProducer.Send(&OrderMessage{
        OrderId:   order.OrderId,
        UserId:    order.UserId,
        GoodsId:   order.GoodsId,
        Timestamp: order.Timestamp,
        Action:    "create_order",
    })
}

// 消费者处理持久化
func (c *OrderConsumer) ProcessOrder(msg *OrderMessage) error {
    // 写入数据库
    err := c.dbService.CreateOrder(msg)
    if err != nil {
        // 重试机制
        return c.retryCreateOrder(msg)
    }
    
    // 确认消息
    return nil
}
```

**方案二：基于Redisson的分布式锁**
```java
// Java示例：Redis + DB事务一致性
@Transactional
public SecKillResult secKill(Long goodsId, Long userId) {
    RLock lock = redisson.getLock("seckill:lock:" + goodsId);
    
    try {
        // 获取分布式锁
        if (!lock.tryLock(100, TimeUnit.MILLISECONDS)) {
            return SecKillResult.busy();
        }
        
        // 1. Redis操作
        SecKillResult redisResult = redisSecKill(goodsId, userId);
        if (!redisResult.isSuccess()) {
            return redisResult;
        }
        
        // 2. 数据库操作  
        orderService.createOrder(redisResult.getOrderId());
        
        return redisResult;
        
    } finally {
        lock.unlock();
    }
}
```

**方案三：双写一致性 + 补偿机制**
```go
type ConsistencyManager struct {
    redis      *redis.Client
    db         *sql.DB
    compensate chan *CompensateTask
}

func (cm *ConsistencyManager) SecKillWithConsistency(goodsId, userId int64) error {
    // 1. Redis秒杀
    redisResult, err := cm.redisSecKill(goodsId, userId)
    if err != nil {
        return err
    }
    
    // 2. 异步写DB
    go func() {
        if err := cm.db.CreateOrder(redisResult.OrderId); err != nil {
            // 写DB失败，发起补偿
            cm.compensate <- &CompensateTask{
                Type: "rollback_redis",
                OrderId: redisResult.OrderId,
                GoodsId: goodsId,
                UserId: userId,
            }
        }
    }()
    
    return nil
}

// 补偿任务处理
func (cm *ConsistencyManager) compensateWorker() {
    for task := range cm.compensate {
        switch task.Type {
        case "rollback_redis":
            cm.rollbackRedisOrder(task)
        case "retry_db":
            cm.retryDBOrder(task)
        }
    }
}
```

#### 4. 数据恢复与校验

**4.1 定时数据校验**
```go
// 每5分钟校验一次数据一致性
func (s *ConsistencyChecker) CheckConsistency() {
    // 获取Redis中的订单
    redisOrders := s.getRedisOrders()
    
    // 获取DB中的订单  
    dbOrders := s.getDBOrders()
    
    // 比较差异
    diff := s.compareOrders(redisOrders, dbOrders)
    
    // 处理不一致数据
    for _, inconsistency := range diff {
        s.handleInconsistency(inconsistency)
    }
}
```

**4.2 数据修复机制**
```lua
-- Redis数据修复脚本
local function repairData(goodsId)
    local stockKey = 'stock:' .. goodsId
    local buyersKey = 'buyers:' .. goodsId
    local ordersKey = 'orders:' .. goodsId
    
    -- 获取实际已售出数量
    local soldCount = redis.call('SCARD', buyersKey)
    local orderCount = redis.call('LLEN', ordersKey)
    
    -- 修复库存不一致
    if soldCount ~= orderCount then
        redis.call('LTRIM', ordersKey, 0, soldCount - 1)
    end
    
    return soldCount
end
```

#### 5. 性能优化与监控

**5.1 Pipeline批量操作**
```go
// 批量处理订单状态更新
func (s *SecKillService) BatchUpdateOrderStatus(orders []Order) error {
    pipe := s.redisClient.Pipeline()
    
    for _, order := range orders {
        orderKey := fmt.Sprintf("order:%s", order.OrderId)
        pipe.HSet(orderKey, "status", order.Status)
        pipe.HSet(orderKey, "updated_at", time.Now().Unix())
    }
    
    _, err := pipe.Exec()
    return err
}
```

**5.2 实时监控指标**
```go
type SecKillMetrics struct {
    TotalRequests    int64  // 总请求数
    SuccessCount     int64  // 成功次数
    FailureCount     int64  // 失败次数
    RedisLatency     time.Duration // Redis延迟
    DBSyncLatency    time.Duration // DB同步延迟
    InconsistentCount int64 // 不一致数据量
}

// 监控告警
func (m *SecKillMetrics) CheckAlerts() {
    successRate := float64(m.SuccessCount) / float64(m.TotalRequests)
    if successRate < 0.999 {
        m.sendAlert("秒杀成功率低于99.9%")
    }
    
    if m.InconsistentCount > 100 {
        m.sendAlert("数据不一致数量超过阈值")
    }
}
```

#### 6. 优势与限制

**优势**
- **极致性能**: Redis内存操作，QPS可达千万级
- **原子性**: Lua脚本保证操作原子性，无并发问题
- **低延迟**: 内存操作，响应时间微秒级
- **扩展性**: Redis集群可水平扩展

**限制与解决方案**
- **内存成本**: 需要大量内存 → 数据分层存储
- **持久化风险**: 内存数据易丢失 → AOF+RDB双重保障
- **一致性复杂**: 需要额外保证 → 补偿机制+定时校验
- **运维复杂**: 集群管理复杂 → 自动化运维工具

**最终一致性保证策略**
1. **实时层**: Redis Lua脚本保证原子性
2. **异步层**: 消息队列异步持久化
3. **补偿层**: 定时任务检查并修复不一致
4. **监控层**: 实时监控告警，快速发现问题
```

### 极致性能优化

#### 1. 硬件级优化

**CPU绑定**
```bash
# 绑定进程到指定CPU核心
taskset -c 0-3 ./seckill-server
```

**网络优化**
```bash
# 万兆网卡 + DPDK用户态网络栈
echo 'net.core.rmem_max = 134217728' >> /etc/sysctl.conf
echo 'net.ipv4.tcp_rmem = 4096 87380 134217728' >> /etc/sysctl.conf
```

**内存优化**
- 大页内存：减少TLB缺页
- NUMA感知：内存就近访问
- 内存预分配：避免运行时分配

#### 2. 算法级优化

**布隆过滤器防刷**
```go
// 10亿用户，误判率0.001%
bloomFilter := bloom.NewWithEstimates(1000000000, 0.00001)

func checkUser(userId int64) bool {
    key := fmt.Sprintf("user:%d", userId)
    if !bloomFilter.TestString(key) {
        return false  // 一定不存在
    }
    // 可能存在，需要进一步验证
    return checkUserInRedis(userId)
}
```

**LRU-K热点检测**
```go
type HotSpotDetector struct {
    lru *lru.Cache  // 热点商品缓存
    counter *Counter // 访问计数器
}

func (h *HotSpotDetector) IsHot(goodsId int64) bool {
    count := h.counter.Get(goodsId)
    return count > HOTSPOT_THRESHOLD
}
```

#### 3. 数据库优化

**分库分表策略**
```sql
-- 按商品ID分表，1000张表
CREATE TABLE seckill_order_{0000-0999} (
    order_id BIGINT PRIMARY KEY,
    user_id BIGINT,
    goods_id BIGINT,
    created_at TIMESTAMP,
    INDEX idx_user_goods (user_id, goods_id)
) ENGINE=InnoDB;
```

**读写分离 + 连接池**
```go
// HikariCP连接池配置
dbPool := &sql.DB{
    MaxOpenConns: 1000,        // 最大连接数
    MaxIdleConns: 100,         // 空闲连接数
    ConnMaxLifetime: time.Hour, // 连接最大生命周期
}
```

### 容错与监控

#### 1. 多级降级策略

**自动降级开关**
```go
type DegradeController struct {
    cpuThreshold    float64 // CPU阈值
    memoryThreshold float64 // 内存阈值  
    qpsThreshold    int64   // QPS阈值
}

func (dc *DegradeController) ShouldDegrade() bool {
    return getCurrentCPU() > dc.cpuThreshold ||
           getCurrentMemory() > dc.memoryThreshold ||
           getCurrentQPS() > dc.qpsThreshold
}
```

**降级方案**
- Level1：关闭非核心功能
- Level2：返回预设静态数据
- Level3：直接返回"系统繁忙"

#### 2. 实时监控

**关键指标**
```
系统指标: CPU、内存、网络、磁盘IO
业务指标: QPS、响应时间、成功率、库存准确性
```

**报警机制**
- QPS > 1200万：告警
- 平均响应时间 > 5ms：告警
- 成功率 < 99.9%：告警
- 出现超卖：紧急告警

### 性能指标

- **峰值QPS**: 1000万+ (支持瞬时1200万)
- **平均响应时间**: 2ms
- **99.9%响应时间**: 5ms  
- **系统可用性**: 99.99%
- **库存准确性**: 100% (绝对不超卖)
- **并发用户数**: 1000万+

### 成本估算

**硬件成本**
- 应用服务器：1000台 * 2万 = 2000万
- Redis集群：100台 * 3万 = 300万
- 网络设备：万兆交换机 + 负载均衡 = 500万
- **总计：2800万/年**

**运维成本**  
- CDN费用：500万/年
- 带宽费用：1000万/年
- 人力成本：200万/年
- **总计：1700万/年**

**整体投入：4500万/年，支撑千万QPS秒杀**

---

## Q2: LLM应用场景

### 主要应用领域

#### 1. 内容生成
- **文本创作**: 文章写作、营销文案、技术文档
- **代码生成**: 自动编程、代码补全、bug修复
- **创意设计**: 广告创意、产品描述、SEO内容

#### 2. 智能客服与助手
- **对话系统**: 7x24小时客户服务
- **知识问答**: 企业内部知识库查询
- **任务自动化**: 邮件处理、日程安排

#### 3. 数据分析与处理
- **文档理解**: PDF/Word文档解析
- **数据提取**: 结构化信息抽取
- **内容审核**: 违规内容检测

#### 4. 教育培训
- **个性化教学**: 自适应学习路径
- **作业批改**: 自动评分和反馈
- **语言学习**: 口语练习、语法纠错

### Q2.1 Prompt调优方法论

#### 1. 结构化Prompt设计

**基础模板**
```
Role: 你是一个{专业角色}
Context: 在{具体场景}下
Task: 请{具体任务描述}
Format: 输出格式为{指定格式}
Examples: {提供示例}
Constraints: 注意{约束条件}
```

#### 2. 优化策略

**Few-shot Learning**
```
示例1: 输入 -> 期望输出
示例2: 输入 -> 期望输出  
示例3: 输入 -> 期望输出
现在处理: {实际输入}
```

**Chain of Thought (CoT)**
```
请按照以下步骤思考：
1. 分析问题的关键要素
2. 确定解决思路
3. 逐步推理
4. 给出最终答案
```

**思维树 (Tree of Thoughts)**
- 生成多个解决方案分支
- 评估每个分支的可行性
- 选择最优路径继续推理

#### 3. 评估与迭代

**评估维度**
- 准确性：答案是否正确
- 相关性：是否回答了问题
- 完整性：信息是否充分
- 一致性：多次运行结果稳定性

**A/B测试框架**
- 对照组：原始prompt
- 实验组：优化后prompt  
- 评估指标：成功率、用户满意度

### Q2.2 RAG系统优化策略

#### 1. RAG架构设计

```
查询 -> 查询理解 -> 向量检索 -> 重排序 -> 上下文构建 -> LLM生成 -> 答案输出
```

#### 2. 检索准确性优化

**文档处理优化**
- 智能分块：基于语义而非固定长度
- 多层级索引：段落级 + 文档级
- 元数据增强：添加标题、时间、来源信息

**向量检索优化**
- 混合检索：Dense + Sparse检索结合
- 查询扩展：同义词扩展、查询重写
- 负采样：提高向量区分度

**代码示例**
```python
def hybrid_retrieval(query, top_k=10):
    # Dense检索
    dense_results = vector_db.similarity_search(
        embed_query(query), k=top_k*2
    )
    
    # Sparse检索 (BM25)
    sparse_results = bm25_index.search(query, k=top_k*2)
    
    # 结果融合
    return rerank_results(dense_results, sparse_results, top_k)
```

#### 3. 重排序策略

**多阶段排序**
1. 粗排：向量相似度
2. 精排：Cross-encoder重排序
3. 多样性：MMR算法避免重复

**上下文质量评估**
```python
def context_quality_score(context, query):
    relevance_score = calc_relevance(context, query)
    freshness_score = calc_freshness(context.timestamp)
    authority_score = calc_authority(context.source)
    
    return 0.6 * relevance_score + 0.3 * freshness_score + 0.1 * authority_score
```

#### 4. 生成优化

**Prompt工程**
```
基于以下上下文信息回答问题：

上下文：
{retrieved_contexts}

问题：{query}

要求：
1. 仅基于上下文信息回答
2. 如果信息不足，明确说明
3. 标注信息来源
```

**幻觉检测与修正**
- 事实校验：将答案与原始文档对比
- 置信度评估：模型输出概率分析
- 引用验证：确保引用内容真实存在

---

## Q3: 事务机制与隔离级别

### 事务ACID特性

#### 1. 原子性 (Atomicity)
- **定义**: 事务中的所有操作要么全部成功，要么全部失败
- **实现**: 通过日志记录和回滚机制
- **示例**: 转账操作中扣款和加款必须同时成功

#### 2. 一致性 (Consistency)  
- **定义**: 事务执行前后，数据库状态保持一致
- **实现**: 通过约束检查和触发器
- **示例**: 转账前后总金额保持不变

#### 3. 隔离性 (Isolation)
- **定义**: 并发事务之间相互隔离，不互相干扰
- **实现**: 通过锁机制和MVCC
- **示例**: 两个用户同时查询余额，看到一致的结果

#### 4. 持久性 (Durability)
- **定义**: 事务提交后，数据永久保存
- **实现**: 通过WAL日志和刷盘机制
- **示例**: 转账成功后即使断电也不会丢失

### 事务隔离级别

#### 1. 读未提交 (Read Uncommitted)
```sql
SET TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;
```
- **特点**: 可以读取未提交的数据
- **问题**: 脏读、不可重复读、幻读
- **应用**: 对一致性要求不高的统计查询

#### 2. 读已提交 (Read Committed)
```sql
SET TRANSACTION ISOLATION LEVEL READ COMMITTED;
```
- **特点**: 只能读取已提交的数据
- **问题**: 不可重复读、幻读
- **应用**: PostgreSQL默认级别，适合大部分应用

#### 3. 可重复读 (Repeatable Read)
```sql
SET TRANSACTION ISOLATION LEVEL REPEATABLE READ;
```
- **特点**: 事务期间多次读取同一数据结果一致
- **问题**: 幻读（部分数据库已解决）
- **应用**: MySQL InnoDB默认级别

#### 4. 串行化 (Serializable)
```sql
SET TRANSACTION ISOLATION LEVEL SERIALIZABLE;
```
- **特点**: 最严格的隔离级别，完全串行化执行
- **问题**: 性能最低，锁竞争严重
- **应用**: 对一致性要求极高的场景

### 并发问题与解决方案

#### 1. 脏读 (Dirty Read)
**问题**: 读取到未提交的数据
```sql
-- 事务A
BEGIN;
UPDATE account SET balance = 500 WHERE id = 1;
-- 未提交

-- 事务B同时执行
SELECT balance FROM account WHERE id = 1; -- 读到500（脏数据）
```
**解决**: 使用读已提交及以上隔离级别

#### 2. 不可重复读 (Non-Repeatable Read)
**问题**: 同一事务中多次读取同一数据结果不同
```sql
-- 事务A
BEGIN;
SELECT balance FROM account WHERE id = 1; -- 第一次读取：1000

-- 事务B执行并提交
UPDATE account SET balance = 500 WHERE id = 1;
COMMIT;

-- 事务A继续
SELECT balance FROM account WHERE id = 1; -- 第二次读取：500
```
**解决**: 使用可重复读隔离级别

#### 3. 幻读 (Phantom Read)
**问题**: 同一事务中多次查询，结果集中数据条数不同
```sql
-- 事务A
BEGIN;
SELECT COUNT(*) FROM account WHERE balance > 1000; -- 第一次：5条

-- 事务B插入新数据
INSERT INTO account VALUES (6, 'user6', 1500);
COMMIT;

-- 事务A继续
SELECT COUNT(*) FROM account WHERE balance > 1000; -- 第二次：6条
```
**解决**: 使用串行化隔离级别或间隙锁

### 事务实现机制

#### 1. 锁机制
**共享锁 (S锁)**
```sql
SELECT * FROM account WHERE id = 1 LOCK IN SHARE MODE;
```

**排他锁 (X锁)**
```sql
SELECT * FROM account WHERE id = 1 FOR UPDATE;
```

**意向锁**: 表级锁，提高锁检查效率

#### 2. MVCC (多版本并发控制)
**实现原理**:
- 每行数据维护多个版本
- 通过时间戳或事务ID标识版本
- 读操作根据隔离级别选择合适版本

**MySQL InnoDB实现**:
```
数据行格式: [DATA] [TRX_ID] [ROLL_PTR] [其他字段]
- TRX_ID: 修改该行的事务ID
- ROLL_PTR: 指向undo log的指针
```

#### 3. WAL (Write-Ahead Logging)
**执行流程**:
1. 修改操作先写入日志
2. 日志刷盘后才修改数据页
3. 数据页可延迟刷盘

**恢复机制**:
- Redo log: 重做已提交事务
- Undo log: 回滚未提交事务

### 分布式事务

#### 1. 两阶段提交 (2PC)
```
阶段1 - 准备阶段:
协调者 -> 所有参与者: prepare()
参与者 -> 协调者: yes/no

阶段2 - 提交阶段:
如果都是yes: 协调者 -> 所有参与者: commit()
如果有no: 协调者 -> 所有参与者: abort()
```

#### 2. TCC模式
```java
// Try: 资源预留
public boolean try(String userId, BigDecimal amount) {
    return accountService.freeze(userId, amount);
}

// Confirm: 确认执行
public boolean confirm(String userId, BigDecimal amount) {
    return accountService.deduct(userId, amount);
}

// Cancel: 取消回滚  
public boolean cancel(String userId, BigDecimal amount) {
    return accountService.unfreeze(userId, amount);
}
```

#### 3. Saga模式
```java
// 正向操作链
transferMoney() -> deductAccount() -> addAccount() -> updateTransferRecord()

// 补偿操作链  
cancelTransfer() -> rollbackRecord() -> rollbackAdd() -> rollbackDeduct()
```

### 实际应用建议

#### 1. 隔离级别选择
- **读密集应用**: 读已提交，配合缓存
- **金融系统**: 可重复读或串行化
- **统计分析**: 读未提交（允许脏读）

#### 2. 性能优化
- 减少事务持有时间
- 避免长事务
- 合理使用索引减少锁范围
- 读写分离减少锁竞争

#### 3. 监控指标
- 事务等待时间
- 死锁发生频率  
- 长事务数量
- 锁等待超时次数