package design

package main

import (
	"fmt"
	"strconv"
	"time"
	"your-module/lru" // 替换为实际模块路径
)

func main() {
	fmt.Println("=== 原生泛型LRU缓存示例 ===\n")
	
	// 基本使用示例
	basicExample()
	fmt.Println()
	
	// 类型安全示例
	typeSafetyExample()
	fmt.Println()
	
	// 淘汰回调示例
	evictionCallbackExample()
	fmt.Println()
	
	// 遍历示例
	iterationExample()
	fmt.Println()
	
	// 统计功能示例
	statsExample()
	fmt.Println()
	
	// 实际应用场景示例
	cacheScenarioExample()
}

// 基本使用示例
func basicExample() {
	fmt.Println("--- 基本使用示例 ---")
	
	// 创建string -> int类型的LRU缓存
	cache, _ := lru.New[string, int](3)
	
	// 添加数据
	cache.Put("one", 1)
	cache.Put("two", 2)
	cache.Put("three", 3)
	
	fmt.Printf("缓存大小: %d/%d\n", cache.Len(), cache.Cap())
	
	// 访问数据
	if val, ok := cache.Get("one"); ok {
		fmt.Printf("获取 'one': %d\n", val)
	}
	
	// 添加新数据，触发淘汰
	cache.Put("four", 4)
	
	fmt.Printf("添加 'four' 后的键列表: %v\n", cache.Keys())
	
	// 检查被淘汰的项
	if !cache.Contains("two") {
		fmt.Println("'two' 已被淘汰（最久未使用）")
	}
}

// 类型安全示例
func typeSafetyExample() {
	fmt.Println("--- 类型安全示例 ---")
	
	// 不同类型的缓存
	intCache, _ := lru.New[int, string](5)
	userCache, _ := lru.New[string, User](3)
	
	// int -> string 缓存
	intCache.Put(1, "one")
	intCache.Put(2, "two")
	
	if val, ok := intCache.Get(1); ok {
		fmt.Printf("intCache[1] = %s\n", val)
	}
	
	// string -> User 缓存
	userCache.Put("alice", User{Name: "Alice", Age: 25})
	userCache.Put("bob", User{Name: "Bob", Age: 30})
	
	if user, ok := userCache.Get("alice"); ok {
		fmt.Printf("userCache['alice'] = %+v\n", user)
	}
	
	// 编译时类型检查
	// intCache.Put("invalid", 123) // 编译错误：key必须是int
	// userCache.Put("charlie", "invalid") // 编译错误：value必须是User
}

type User struct {
	Name string
	Age  int
}

// 淘汰回调示例
func evictionCallbackExample() {
	fmt.Println("--- 淘汰回调示例 ---")
	
	cache, _ := lru.New[string, string](2,
		lru.WithEvictCallback(func(key, value string) {
			fmt.Printf("  淘汰: %s -> %s\n", key, value)
		}))
	
	cache.Put("a", "apple")
	cache.Put("b", "banana")
	fmt.Println("添加 a, b")
	
	cache.Put("c", "cherry") // 淘汰 a
	fmt.Println("添加 c")
	
	cache.Put("d", "date") // 淘汰 b
	fmt.Println("添加 d")
	
	fmt.Printf("最终键列表: %v\n", cache.Keys())
}

// 遍历示例
func iterationExample() {
	fmt.Println("--- 遍历示例 ---")
	
	cache, _ := lru.New[int, string](5)
	
	// 添加数据
	for i := 1; i <= 5; i++ {
		cache.Put(i, fmt.Sprintf("value_%d", i))
	}
	
	// 访问某些项以改变LRU顺序
	cache.Get(2)
	cache.Get(4)
	
	fmt.Println("正向遍历（从最新到最旧）:")
	cache.ForEach(func(key int, value string) bool {
		fmt.Printf("  %d -> %s\n", key, value)
		return true // 继续遍历
	})
	
	fmt.Println("反向遍历（从最旧到最新）:")
	cache.ForEachReverse(func(key int, value string) bool {
		fmt.Printf("  %d -> %s\n", key, value)
		return true
	})
	
	// 获取所有键值对
	entries := cache.Entries()
	fmt.Printf("所有条目: %+v\n", entries)
}

// 统计功能示例
func statsExample() {
	fmt.Println("--- 统计功能示例 ---")
	
	cache, _ := lru.NewWithStats[string, int](3)
	
	// 添加数据
	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)
	
	// 模拟访问模式
	cache.Get("a") // 命中
	cache.Get("b") // 命中
	cache.Get("x") // 未命中
	cache.Get("y") // 未命中
	cache.Get("a") // 命中
	cache.Get("c") // 命中
	
	stats := cache.GetStats()
	fmt.Printf("统计信息:\n")
	fmt.Printf("  容量: %d\n", stats.Capacity)
	fmt.Printf("  大小: %d\n", stats.Size)
	fmt.Printf("  命中率: %.2f%%\n", stats.HitRatio*100)
}

// 实际应用场景示例：数据库查询缓存
func cacheScenarioExample() {
	fmt.Println("--- 数据库查询缓存示例 ---")
	
	// 查询结果结构
	type QueryResult struct {
		Data      []map[string]interface{}
		Timestamp time.Time
	}
	
	// 创建查询缓存
	queryCache, _ := lru.New[string, *QueryResult](100,
		lru.WithEvictCallback(func(query string, result *QueryResult) {
			fmt.Printf("  查询结果过期: %s (缓存时间: %v)\n",
				query, time.Since(result.Timestamp).Truncate(time.Millisecond))
		}))
	
	// 模拟数据库查询函数
	executeQuery := func(sql string) *QueryResult {
		// 检查缓存
		if result, ok := queryCache.Get(sql); ok {
			fmt.Printf("缓存命中: %s\n", sql)
			return result
		}
		
		// 模拟数据库查询（耗时操作）
		fmt.Printf("执行查询: %s\n", sql)
		time.Sleep(50 * time.Millisecond) // 模拟查询延迟
		
		result := &QueryResult{
			Data: []map[string]interface{}{
				{"id": 1, "name": "test"},
			},
			Timestamp: time.Now(),
		}
		
		// 缓存结果
		queryCache.Put(sql, result)
		return result
	}
	
	// 执行查询
	queries := []string{
		"SELECT * FROM users WHERE active = 1",
		"SELECT COUNT(*) FROM orders",
		"SELECT * FROM products WHERE category = 'electronics'",
	}
	
	// 第一轮查询（都会执行）
	fmt.Println("第一轮查询:")
	for _, query := range queries {
		executeQuery(query)
	}
	
	fmt.Printf("缓存大小: %d\n", queryCache.Len())
	
	// 第二轮查询（从缓存获取）
	fmt.Println("\n第二轮查询:")
	for _, query := range queries {
		executeQuery(query)
	}
	
	// 添加更多查询直到触发淘汰
	fmt.Println("\n添加更多查询:")
	for i := 0; i < 100; i++ {
		query := fmt.Sprintf("SELECT * FROM table_%d", i)
		executeQuery(query)
	}
	
	fmt.Printf("最终缓存大小: %d\n", queryCache.Len())
}

// 性能测试示例
func performanceExample() {
	fmt.Println("--- 性能测试示例 ---")
	
	cache, _ := lru.New[int, string](10000)
	
	// 写入性能测试
	start := time.Now()
	for i := 0; i < 100000; i++ {
		cache.Put(i, strconv.Itoa(i))
	}
	writeTime := time.Since(start)
	
	// 读取性能测试
	start = time.Now()
	for i := 0; i < 100000; i++ {
		cache.Get(i % 10000)
	}
	readTime := time.Since(start)
	
	fmt.Printf("写入 100,000 项耗时: %v\n", writeTime)
	fmt.Printf("读取 100,000 项耗时: %v\n", readTime)
	fmt.Printf("最终缓存大小: %d\n", cache.Len())
}

// 缓存替换策略演示
func replacementPolicyDemo() {
	fmt.Println("--- LRU替换策略演示 ---")
	
	cache, _ := lru.New[string, int](4)
	
	// 添加初始数据
	cache.Put("A", 1)
	cache.Put("B", 2)
	cache.Put("C", 3)
	cache.Put("D", 4)
	fmt.Printf("初始状态: %v\n", cache.Keys())
	
	// 访问A，使其成为最新
	cache.Get("A")
	fmt.Printf("访问A后: %v\n", cache.Keys())
	
	// 访问C，使其成为最新
	cache.Get("C")
	fmt.Printf("访问C后: %v\n", cache.Keys())
	
	// 添加新项E，应该淘汰B（最久未使用）
	cache.Put("E", 5)
	fmt.Printf("添加E后: %v\n", cache.Keys())
	
	// 验证B被淘汰
	if !cache.Contains("B") {
		fmt.Println("B已被淘汰（符合LRU策略）")
	}
}

// 内存使用分析
func memoryAnalysis() {
	fmt.Println("--- 内存使用分析 ---")
	
	const capacity = 1000
	cache, _ := lru.New[int, [100]byte](capacity)
	
	// 填充缓存
	for i := 0; i < capacity; i++ {
		var data [100]byte
		for j := range data {
			data[j] = byte(i % 256)
		}
		cache.Put(i, data)
	}
	
	fmt.Printf("缓存容量: %d\n", cache.Cap())
	fmt.Printf("当前大小: %d\n", cache.Len())
	fmt.Printf("估计内存使用: ~%.2f KB\n", 
		float64(capacity * (100 + 16)) / 1024) // 100字节数据 + 约16字节开销
}