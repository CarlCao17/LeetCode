

# **A/B 实验平台技术架构方案**

## **1\. 核心设计原则**

本方案基于以下四个核心原则构建，旨在打造一个高性能、高可靠、可扩展且科学严谨的实验平台：

1. **无状态计算 (Stateless Computation)**：分流决策服务本身不存储任何用户状态。每次决策都基于传入的用户标识和实验配置实时计算得出。这确保了卓越的水平扩展能力和高容错性 1。  
2. **确定性哈希 (Deterministic Hashing)**：对于固定的用户和实验配置，分流结果永远是相同的。这保证了用户体验的一致性和实验数据的有效性 6。  
3. **正交与互斥 (Orthogonality & Exclusivity)**：通过分层设计，天然支持不同业务域的实验并行（层间正交），同时保证同一业务域内的实验互不干扰（同层互斥），最大化流量利用效率 8。  
4. **配置驱动 (Configuration-Driven)**：实验的定义、流量的分配和生命周期管理完全由中心化的配置驱动，业务代码与实验逻辑解耦。这与 Uber 等公司的先进实践一致，提供了极高的灵活性和安全性 13。

---

## **2\. 系统架构概览**

整个平台由五个核心部分组成，构成一个完整的闭环生态系统：

\!([https://i.imgur.com/8iYfLzW.png](https://i.imgur.com/8iYfLzW.png))

1. **实验管理平台 (Experiment Management UI)**：  
   * **角色**：一个 Web 前端应用，是实验科学家、产品经理和工程师进行所有实验操作的唯一入口。  
   * **功能**：提供实验的创建、配置、启动、暂停、结束等全生命周期管理功能。可视化展示流量层的桶分配情况。  
2. **配置与元数据服务 (Config & Metadata Service)**：  
   * **角色**：一个后端的 CRUD 服务，负责处理来自管理平台的请求，并对实验配置进行持久化存储。它是实验“事实状态”的管理者。  
   * **功能**：维护实验、流量层、指标等元数据。核心是实现我们之前讨论的**原子化**的流量桶分配图更新算法。  
3. **配置中心 (Configuration Center)**：  
   * **角色**：一个高可用的分布式键值存储系统，作为实验配置的“真理之源”（Source of Truth）。  
   * **功能**：存储由元数据服务生成的最终分流配置。分流服务通过订阅（watch）此处的变更来动态更新内存中的规则。  
4. **分流服务 (Diversion Service)**：  
   * **角色**：一个高性能、低延迟、无状态的 gRPC/HTTP 服务。它是整个系统的执行引擎。  
   * **功能**：接收来自业务服务的请求（包含用户标识等上下文），根据从配置中心同步的最新规则，实时计算并返回实验分组结果。  
5. **数据管道与分析 (Data Pipeline & Analytics)**：  
   * **角色**：一个异步的数据处理系统，负责收集、处理和分析实验数据。  
   * **功能**：业务方将分流服务返回的实验分组信息，连同业务行为日志（如点击、转化）一同上报到数据管道。数据经过处理后存入数据仓库，用于后续的统计分析和效果计算。

---

## **3\. 核心组件深度设计**

### **3.1. 分流服务 (Diversion Service)**

这是系统的性能核心，设计上必须追求极致的简洁和高效。

* **职责**：  
  * 启动时从配置中心拉取全量实验配置，并在内存中构建高效的查找结构（bucket\_map）。  
  * 监听配置中心的变更，一旦有更新，则原子地替换内存中的查找结构。  
  * 对外提供一个单一的、轻量级的 API 接口，用于分流决策。  
* **API 设计 (gRPC/REST)**：  
  Protocol Buffers  
  // gRPC Definition  
  service DiversionService {  
    rpc Assign (AssignRequest) returns (AssignResponse);  
  }

  message AssignRequest {  
    string user\_id \= 1;         // 核心用户标识 (设备ID/登录ID)  
    map\<string, string\> context \= 2; // 其他上下文信息，如 app\_version, city\_id 等  
  }

  message AssignResponse {  
    map\<string, Assignment\> assignments \= 1; // key: layer\_name, value: 实验分组信息  
  }

  message Assignment {  
    string flight\_name \= 1;  
    string version\_name \= 2;  
  }

  * **设计考量**：一次请求返回该用户在**所有**流量层命中的实验结果。这避免了业务方需要为不同层的实验多次请求分流服务，极大地提升了效率。  
* **容错与降级**：  
  * **本地缓存**：分流服务应在本地文件系统缓存一份“最后可用的”配置。如果启动时无法连接配置中心，则加载本地缓存启动，保证服务可用性。  
  * **默认返回**：在任何异常情况下（如配置解析失败），服务应安全降级，返回不命中任何实验的默认结果。

### **3.2. 配置与元数据服务**

这是系统的管理核心，负责维护逻辑的正确性和数据的一致性。

* **数据模型 (Data Models)**：  
  * **流量层 (Layer)**  
    JSON  
    {  
      "layer\_name": "homepage\_rec\_algo",  
      "total\_buckets": 1000,  
      "allocation\_map":  
    }

  * **实验 (Experiment)**  
    JSON  
    {  
      "flight\_name": "rec\_model\_v3\_exp",  
      "description": "推荐模型V3版本",  
      "owner": "user@example.com",  
      "status": "RUNNING", // DRAFT, RUNNING, PAUSED, ENDED  
      "layer\_name": "homepage\_rec\_algo",  
      "required\_buckets": 100, // 10% 流量  
      "versions": \[  
        {"name": "control", "weight": 50},  
        {"name": "treatment\_v3", "weight": 50}  
      \]  
    }

  * **流量桶分配图 (Allocation Map)**: 这是我们之前设计的核心数据结构，用于管理一个层内的所有桶。  
    JSON

* **核心逻辑**：  
  * 实现上一份报告中详述的、包含**碎片整理**的实验生命周期管理算法（新增、结束、调整流量）。  
  * 每次对 allocation\_map 的修改都必须是**事务性的**。  
  * 当一个实验的配置发生变更并被“发布”时，该服务会：  
    1. 根据实验元数据和 allocation\_map 生成一份给分流服务消费的、精简的最终配置 JSON。  
    2. 将这份 JSON 发布到配置中心（例如，key 为 /ab/config/layers/homepage\_rec\_algo）。

---

## **4\. 核心算法详解**

### **4.1. 双重哈希分流算法**

这是分流决策的数学基础。

Python

\# 伪代码  
import mmh3 \# MurmurHash3 library

TOTAL\_BUCKETS \= 1000

def get\_assignment(user\_id, layer\_config):  
    \# 1\. 流量分桶 (Layer Bucketing)  
    \# 使用 layer\_name 作为盐，实现层间正交  
    layer\_salt \= layer\_config\["layer\_name"\]  
    bucket\_seed\_str \= f"{user\_id}\_{layer\_salt}"  
    bucket\_hash \= mmh3.hash(bucket\_seed\_str, signed=False)  
    bucket\_id \= bucket\_hash % TOTAL\_BUCKETS

    \# 2\. 查找实验 (Experiment Lookup)  
    \# O(1) 复杂度  
    experiment\_config \= layer\_config\["bucket\_map"\]\[bucket\_id\]  
    if experiment\_config is None:  
        return None \# 未命中实验

    \# 3\. 实验分组 (Experiment Grouping)  
    \# 使用 flight\_name 作为盐，保证实验内部分组稳定  
    flight\_name \= experiment\_config\["flight\_name"\]  
    group\_seed\_str \= f"{user\_id}:{flight\_name}"  
    group\_hash \= mmh3.hash(group\_seed\_str, signed=False)  
      
    \# 使用权重进行分组  
    total\_weight \= sum(v\["weight"\] for v in experiment\_config\["versions"\])  
    group\_mod\_result \= group\_hash % total\_weight  
      
    current\_weight \= 0  
    for version in experiment\_config\["versions"\]:  
        current\_weight \+= version\["weight"\]  
        if group\_mod\_result \< current\_weight:  
            return {"flight\_name": flight\_name, "version\_name": version\["name"\]}  
              
    return None \# Should not happen

### **4.2. 渐进式发布 (Gradual Rollout) 算法**

渐进式发布是现代软件交付的关键实践，我们的架构必须原生支持 14。

* **增加流量**：假设 exp\_A 占用了 \`\` 号桶，需要增加到 150 个桶。  
  1. 配置服务读取 allocation\_map。  
  2. 定位到 exp\_A 的区块 {"start": 0, "end": 99,...}。  
  3. 检查其**紧邻的下一个**区块，假设是 {"start": 100, "end": 499, "status": "AVAILABLE"}。  
  4. 检查该可用区块是否有足够空间（500 \- 100 \>= 50）。  
  5. **原子地**更新两个区块：exp\_A 的 end 变为 149，可用区块的 start 变为 150。  
  6. 发布新配置。  
* **减少流量**：逻辑相反，缩小实验区块的 end，并将释放的桶合并到下一个可用区块或创建一个新的可用区块。

这个基于范围的操作确保了所有已在 \`\` 桶内的用户，其分组**绝对不会**发生改变，完美支持了安全的渐进式发布。

---

## **5\. 技术选型建议**

| 组件 | 推荐技术 | 备选方案 | 备注 |
| :---- | :---- | :---- | :---- |
| **管理平台 (UI)** | React / Vue.js | Angular | 成熟的前端框架，生态丰富。 |
| **配置元数据服务** | Go / Java (Spring Boot) | Python (Django/FastAPI) | 需要支持事务性操作，与数据库交互紧密。 |
| **元数据存储** | PostgreSQL / MySQL | \- | 关系型数据库，保证数据强一致性。 |
| **配置中心** | etcd | Consul, Zookeeper | 支持 Watch 机制，高可用，是架构关键。 |
| **分流服务** | Go / Rust | Java, C++ | 性能敏感，选择编译型、高并发语言。 |
| **数据管道** | Kafka | Pulsar | 行业标准的异步消息队列，用于解耦日志上报。 |
| **数据处理** | Spark / Flink | \- | 用于处理海量曝光日志和计算指标。 |
| **数据仓库** | ClickHouse / BigQuery | Snowflake, Redshift | 用于存储最终的实验结果数据，支持快速查询。 |

---

## **6\. 结论**

该技术方案为您提供了一个从零到一构建世界级实验平台的完整指南。它通过**分层、无状态、配置驱动**的设计，解决了实验平台在**科学性、可扩展性和可维护性**三个维度上的核心挑战。

* **对于开发者**：提供了一个简单、清晰、低延迟的接口，将复杂的实验逻辑完全解耦。  
* **对于产品经理/数据科学家**：提供了一个强大、灵活的管理平台，可以安全、高效地并行运行大量实验，加速产品迭代和数据驱动决策。

遵循此蓝图，您将能够构建一个不仅能满足当前需求，更能支撑未来多年业务增长的强大实验基础设施。

#### **引用的著作**

1. Converting stateful application to stateless using AWS services, 访问时间为 九月 1, 2025， [https://aws.amazon.com/blogs/architecture/converting-stateful-application-to-stateless-using-aws-services/](https://aws.amazon.com/blogs/architecture/converting-stateful-application-to-stateless-using-aws-services/)  
2. Stateful vs. Stateless Applications: What's the Difference? \- Pure Storage Blog, 访问时间为 九月 1, 2025， [https://blog.purestorage.com/purely-educational/stateful-vs-stateless-applications-whats-the-difference/](https://blog.purestorage.com/purely-educational/stateful-vs-stateless-applications-whats-the-difference/)  
3. Stateful vs stateless applications \- Red Hat, 访问时间为 九月 1, 2025， [https://www.redhat.com/en/topics/cloud-native-apps/stateful-vs-stateless](https://www.redhat.com/en/topics/cloud-native-apps/stateful-vs-stateless)  
4. A Deep Dive into Stateless and Stateful Services: Architecting Scalable, Efficient Applications | by Scaibu, 访问时间为 九月 1, 2025， [https://scaibu.medium.com/a-deep-dive-into-stateless-and-stateful-services-architecting-scalable-efficient-applications-d6eddf795edf](https://scaibu.medium.com/a-deep-dive-into-stateless-and-stateful-services-architecting-scalable-efficient-applications-d6eddf795edf)  
5. Stateful vs. Stateless: Understanding Key Differences for Apps and IT Systems | Splunk, 访问时间为 九月 1, 2025， [https://www.splunk.com/en\_us/blog/learn/stateful-vs-stateless.html](https://www.splunk.com/en_us/blog/learn/stateful-vs-stateless.html)  
6. MurmurHash: The Scrappy Algorithm That Secretly Powers Half the Internet \- Medium, 访问时间为 九月 1, 2025， [https://medium.com/@thealonemusk/murmurhash-the-scrappy-algorithm-that-secretly-powers-half-the-internet-2d3f79b4509b](https://medium.com/@thealonemusk/murmurhash-the-scrappy-algorithm-that-secretly-powers-half-the-internet-2d3f79b4509b)  
7. Principles behind Traffic Distribution \- Getting Started \- Hackle, 访问时间为 九月 1, 2025， [https://docs-en.hackle.io/docs/ab-bucketing](https://docs-en.hackle.io/docs/ab-bucketing)  
8. Mutually exclusive experiments – Support Help Center, 访问时间为 九月 1, 2025， [https://support.optimizely.com/hc/en-us/articles/4410289064205-Mutually-exclusive-experiments](https://support.optimizely.com/hc/en-us/articles/4410289064205-Mutually-exclusive-experiments)  
9. Mutually Exclusive Experiments: Preventing the Interaction Effect \- AB Tasty, 访问时间为 九月 1, 2025， [https://www.abtasty.com/blog/mutually-exclusive-experiments/](https://www.abtasty.com/blog/mutually-exclusive-experiments/)  
10. Intro to A/B Testing — Do it right, do it always | by Richard Chen | Medium, 访问时间为 九月 1, 2025， [https://medium.com/@richardchen\_81235/intro-to-a-b-testing-do-it-right-do-it-always-620c81d5f0fc](https://medium.com/@richardchen_81235/intro-to-a-b-testing-do-it-right-do-it-always-620c81d5f0fc)  
11. When Should I Use Mutually Exclusive Experiments? \- Eppo, 访问时间为 九月 1, 2025， [https://www.geteppo.com/blog/when-should-i-use-mutually-exclusive-experiments](https://www.geteppo.com/blog/when-should-i-use-mutually-exclusive-experiments)  
12. Should You Run Simultaneous Experiments? A Guide to Avoiding Conflicting Results, 访问时间为 九月 1, 2025， [https://www.convert.com/blog/a-b-testing/should-you-run-simultaneous-ab-tests/](https://www.convert.com/blog/a-b-testing/should-you-run-simultaneous-ab-tests/)  
13. Supercharging A/B Testing at Uber | Uber Blog, 访问时间为 九月 1, 2025， [https://www.uber.com/blog/supercharging-a-b-testing-at-uber/](https://www.uber.com/blog/supercharging-a-b-testing-at-uber/)  
14. A/B Testing \- PromptLayer, 访问时间为 九月 1, 2025， [https://docs.promptlayer.com/why-promptlayer/ab-releases](https://docs.promptlayer.com/why-promptlayer/ab-releases)  
15. Gradually rolling out A/B tests to your audience \- Leanplum Documentation, 访问时间为 九月 1, 2025， [https://docs.leanplum.com/docs/gradually-rolling-out-ab-tests-to-your-audience](https://docs.leanplum.com/docs/gradually-rolling-out-ab-tests-to-your-audience)  
16. Faster and Safer Releases with AB Tasty Rollouts, 访问时间为 九月 1, 2025， [https://www.abtasty.com/rollouts/](https://www.abtasty.com/rollouts/)  
17. Gradual deployments \- Workers \- Cloudflare Docs, 访问时间为 九月 1, 2025， [https://developers.cloudflare.com/workers/configuration/versions-and-deployments/gradual-deployments/](https://developers.cloudflare.com/workers/configuration/versions-and-deployments/gradual-deployments/)