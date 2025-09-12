

# **可扩展 A/B 测试分流服务的设计评估与实施策略**

## **引言：验证现代实验架构的基石**

本报告旨在对您提出的 A/B 测试分流服务架构进行一次深入的评估与分析。首先需要明确的是，您所构思的设计方案不仅是可行的，而且在核心原则上是高度健全和先进的。它所采用的无状态、双重哈希分流机制，与 Netflix、Uber 等行业领导者在构建大规模、高可靠性实验平台时所遵循的最佳实践高度一致 1。因此，本报告的目的并非颠覆或重构您的设计，而是作为一次协作式的深度探讨，旨在加固现有设计、剖析其深层含义，并针对您提出的动态分配挑战，提供一套可直接投入生产的解决方案。

我们的目标是将一个优秀的高层设计，转化为一份详尽、严谨且可执行的工程蓝图。报告将从核心分流机制的原理分析入手，深入探讨实验并发控制的策略，最终聚焦于动态流量分配这一核心难题，并给出一套完整的、包含数据结构和算法的 prescriptive（指令性）解决方案。

---

## **第 1 节：核心分流机制分析：双重哈希无状态设计**

### **1.1. 架构选择的战略意义：无状态（Stateless）架构**

您设计的核心是一种“即时计算”（on-the-fly）的分流模式，这本质上定义了一个无状态（Stateless）系统。这一选择并非仅仅是技术实现上的偏好，而是决定了整个系统运维模式和可扩展性上限的战略性决策。

与此相对的是“伪随机+缓存”（Pseudorandom with Caching, PwC）的方法，这是一种有状态（Stateful）的设计 1。PwC 方法需要依赖服务端缓存或数据库来持久化用户的实验分配结果。这种设计在分布式环境中引入了显著的复杂性，包括数据一致性、单点故障风险以及缓存同步等难题，尤其是在并行运行多个实验时，其局限性更为突出 1。

相比之下，您选择的无状态“哈希+分区”（Hash and Partition, HP）架构具备以下决定性优势：

* **卓越的水平扩展能力**：由于用户分配结果无需存储，分流服务的任何节点都可以独立处理请求。这意味着可以通过简单地增加或减少服务实例来应对流量波动，而无需进行复杂的状态同步 4。  
* **高容错性与可靠性**：任何一个分流服务节点宕机都不会影响用户的实验分组状态，因为下一次请求可以由任何其他健康节点重新计算并得到完全相同的结果。这避免了因状态存储系统故障而导致整个实验平台不可用的风险 4。  
* **简化的部署与运维**：无状态服务可以实现滚动更新、蓝绿部署等现代化部署策略，而无需担心复杂的数据迁移或状态版本兼容问题。这极大地降低了运维负担，并与现代云原生架构理念完美契合 7。

可以说，选择无状态架构，是从根本上规避了与会话管理、数据复制和缓存一致性相关的整整一个类别的问题，为构建一个大规模、高可用的系统奠定了坚实的基础。

### **1.2. 哈希算法的角色与特性：MurmurHash**

在哈希算法的选择上，MurmurHash（特别是 MurmurHash3 版本）是一个非常明智的决策。作为一个非加密哈希函数，它专为性能和良好的分布特性而设计，这使其成为高吞吐量分流服务的理想选择 9。

MurmurHash 的关键特性包括：

* **高性能**：基准测试表明，MurmurHash 的速度远超如 MD5、SHA-1 等加密哈希算法，同时也优于 FNV-1a 等其他流行的非加密哈希算法 9。在需要对每次请求进行实时计算的分流服务中，这种性能优势至关重要。  
* **出色的分布均匀性**：它能将输入值均匀地映射到输出空间，具有极低的哈希碰撞率。这确保了用户被随机且均匀地分配到 1000 个流量桶中，为所有实验提供了公平的流量基础 9。  
* **雪崩效应（Avalanche Effect）**：MurmurHash 内部通过一系列精妙的位运算（乘法、位旋转、异或）和精心挑选的“魔术数字”来实现比特位的充分混合 9。这使得输入值的微小变化（例如，从  
  user123\_layerA 变为 user123\_layerB）会导致输出哈希值的巨大且无关联的变化。正是这一数学特性，为您设计的“层间正交”提供了坚实的理论保障。

虽然 MurmurHash 不具备抵御恶意哈希冲突攻击（如 HashDoS）的能力，但这在 A/B 测试的场景下通常不成问题，因为输入值（用户 ID、实验名等）均由系统内部生成和信任，而非来自不可信的外部用户输入 9。

### **1.3. 双层哈希策略的解构**

您的设计巧妙地运用了两次哈希，分别服务于不同的目的，共同构成了整个分流逻辑的核心。

* 流量分层哈希 (hash("${id}\_${layer\_name}") % 1000)：  
  这第一次哈希是实现实验间隔离的关键。通过将用户 ID 与一个稳定的 layer\_name 进行拼接（即加盐），我们为同一个用户在每一个流量层都生成了一个独立的、与其他层不相关的哈希值。随后对 1000 取模，将用户映射到该层专属的 1000 个流量桶之一。这个过程确保了用户在 A 层的分桶结果，与其在 B 层的分桶结果在统计上是完全独立的。这正是实现层间正交（Orthogonality）的底层机制。  
* 实验分组哈希 (hash("${id}:${flight\_name}") % ${version\_number})：  
  这第二次哈希则保证了用户在某个特定实验内部的分组稳定性。通过将用户 ID 与一个在实验生命周期内不变的 flight\_name（实验的唯一标识符）进行拼接，可以确保一个用户一旦被分配到某个实验组（如控制组或某个实验版本），在实验持续期间，无论被评估多少次，其分配结果都保持不变 1。这种分配的  
  **确定性**和**一致性**是 A/B 测试结果有效性的基本前提 1。

---

## **第 2 节：正交性与互斥性：实验并发的层级化管理**

### **2.1. 通过“层”实现隔离**

您的设计明确提出了“同层互斥，层间正交”的原则，这是一种管理实验并发、避免相互干扰的经典且强大的模式。

* **同层互斥（Mutually Exclusive）**：在同一个流量层（layer）内，所有实验共享同一组 1000 个流量桶。一个流量桶在同一时间只能被分配给一个实验。因此，当一个用户根据其分层哈希值落入某个桶时，他最多只会进入该桶所对应的那个实验。这就自然地保证了同一层内的实验是互斥的，用户不会同时参与两个或多个同层实验，从而避免了因体验叠加而造成的交互效应（Interaction Effect），保证了实验结果的纯净性 13。这种机制在概念上等同于 Optimizely 等成熟平台中的“互斥组”（Exclusion Groups） 13。  
* **层间正交（Orthogonal）**：由于不同层的分桶哈希使用了不同的盐（layer\_name），用户在不同层的分桶结果是相互独立的随机事件。这意味着，一个用户被分配到 A 层的某个实验，并不会影响他被分配到 B 层任何实验的概率。这种统计上的独立性使得跨层实验之间不会产生系统性的偏差，可以安全地并行运行，从而极大地提升了整个组织的实验效率和吞吐量 15。

这种分层模型不仅是一个技术实现，更是一种有效的组织扩展机制。它允许不同的业务团队（例如，主页团队使用“主页层”，支付团队使用“支付层”）在各自的领域内独立、自主地开展实验，无需频繁地跨团队协调，也无需担心污染彼此的实验数据。这种架构上的解耦，是支撑大型企业实现高实验迭代速度的关键 16。

### **2.2. 行业实践对比：分层模型 vs. 统一配置模型**

将您所提出的显式分层模型与业界的其他先进设计进行对比，可以更好地理解其优缺点。一个值得关注的对比对象是 Uber 新一代的实验平台，它采用了一种“统一配置”的思路 3。

Uber 的新平台有意地抽象掉了“层”或“域”这样的物理概念。取而代之的是，每个实验都通过一个唯一的盐（experiment\_key）进行独立的随机化，使得**正交成为默认行为**。而**互斥关系则是通过配置来动态定义**的：让多个实验去控制同一个配置参数，并通过优先级规则来决定在特定上下文中哪个实验生效。这种模型提供了极大的灵活性，任何一组实验都可以根据需要被设置为互斥，而无需预先规划到同一个“层”中。相比之下，Netflix 的平台虽然也支持并行实验，但更多地依赖工具来帮助实验者发现潜在冲突，并由人工进行协调 2。

* **您的分层模型**：优点在于**简单、直观且可预测**。开发者和产品经理可以非常清晰地理解实验间的关系：“实验 A 和实验 B 在同一个层吗？是的。那么它们就是互斥的。” 这种模型的规则是全局性的，易于管理。其缺点在于灵活性稍差，互斥关系是静态绑定的，一旦划分好层，跨层实现互斥会比较困难。  
* **Uber 的统一配置模型**：优点在于**极致的灵活性**。它可以支持更复杂的实验设计，如依赖实验、层级实验等。其缺点在于配置管理的复杂性更高，需要一个更强大的控制平面来处理实验间的依赖和优先级裁决。

对于一个正在起步或快速发展的实验平台而言，您选择的分层模型是一个绝佳的起点。它在实现核心隔离需求的同时，保持了系统的简洁性。未来，如果业务复杂度增长到一定程度，可以借鉴 Uber 的思路，在现有哈希机制的基础上，演进到一个更加动态和灵活的配置模型。

---

## **第 3 节：动态分配的挑战：管理流量桶与实验的映射关系**

### **3.1. 问题定义：分配稳定性的必要性**

您提出的核心问题在于：当一个流量层内有新实验加入或旧实验结束时，如何动态地维护 1000 个流量桶与实验之间的映射关系。这里的核心约束是**必须保证存量实验的用户分配稳定性** 12。

如果一个正在参与 exp\_A 的用户，因为 exp\_B 的加入或 exp\_C 的结束，突然被重新分配到了 exp\_B，或者被移出了 exp\_A，那么 exp\_A 和 exp\_B 这两个实验的统计有效性都将遭到毁灭性破坏。因为实验组和对照组的用户群体发生了变化，违背了随机分组的基本原则。这是一个必须规避的严重故障模式。Optimizely 的文档明确警告，在正在运行的互斥组中添加新实验可能会导致流量重新分配和用户重算分组 13，这凸显了正确处理动态分配逻辑的至关重要性。

### **3.2. 朴素实现的风险：模运算与流量桶碎片化**

一个常见的错误是使用简单的模运算（modulo）来进行分配。例如，假设层内有 exp\_A 和 exp\_B 两个实验，我们通过 bucket\_id % 2 来决定用户进入哪个实验。此时如果新加入 exp\_C，并将分配逻辑改为 bucket\_id % 3，那么几乎所有用户的分配都会发生改变，这将瞬间摧毁所有正在运行的实验。这清晰地表明，必须采用一种更稳健的、基于范围的分配方法。

此外，还需要考虑**流量桶碎片化**（Bucket Fragmentation）的问题。随着不同大小的实验不断地加入和结束，层内可用的流量桶可能会变得像操作系统中的内存碎片一样，被分割成许多个不连续的小块。如果分配算法不够智能，可能会出现这样一种情况：虽然剩余可用桶的总数足够支撑一个新实验，但由于没有足够大的**连续**可用区块，导致新实验无法启动。

---

## **第 4 节：动态流量桶管理的指令性解决方案**

为了解决上述挑战，我们提出一套基于“范围分配图”的数据结构和相应的原子化管理算法。

### **4.1. 配置数据结构：基于范围的分配图（Range-Based Allocation Map）**

建议将每个流量层的状态，用一个**有序的、覆盖整个流量桶区间的对象列表**来表示。这个列表必须完整覆盖 \`\` 的所有桶，且区间之间无重叠、无缝隙。

**层分配图示例结构 (JSON):**

JSON

// 一个流量层的分配图示例

设计原理：  
这种结构远优于简单的键值对或哈希表，因为它显式地将“连续的桶范围”作为一等公民。这使得查找可用空间、修改分配等操作变得高效（例如，可以在这个小型的有序列表上进行线性扫描或二分查找），并从根本上解决了碎片化问题。这种设计思路与行业实践中将流量分配到连续“槽位”或“桶”的做法是一致的 12。

### **4.2. 原子化的配置更新**

对每个层的分配图的任何修改，都必须被视为一个**原子操作**。当管理员添加、结束或调整实验时，正确的流程应该是：读取完整的分配图 \-\> 在内存中进行修改 \-\> 将修改后的完整分配图一次性写回。

技术建议：  
强烈建议使用支持“比较并交换”（Compare-and-Swap, CAS）或提供事务语义的配置存储系统，例如 etcd、Consul，或者一个支持事务的关系型数据库。这可以确保所有更新都是基于最新的配置版本进行的，从而避免了因并发操作导致的数据不一致或配置损坏。这是保障生产环境稳定性的关键一环。

### **4.3. 生命周期管理算法**

以下是针对实验不同生命周期阶段的具体管理算法，旨在提供清晰、可实现的操作逻辑。

#### **4.3.1. 新增一个实验**

1. **读取**：获取指定流量层的当前分配图。  
2. **查找**：遍历分配图，寻找一个 status 为 AVAILABLE 且大小（end\_bucket \- start\_bucket \+ 1）足以容纳新实验所需桶数的区块。可以使用“首次适应”（First-Fit）或“最佳适应”（Best-Fit）等策略。  
3. **分配与切分**：  
   * 如果找到合适的区块，则“切分”该区块。创建一个代表新实验的新条目，并相应地调整原 AVAILABLE 区块的边界（如果原区块被完全占用，则删除该条目）。  
   * 例如，要在 的可用区块中分配 100 个桶给 \`exp\_C\`，则原条目变为，并插入新条目 {"start\_bucket": 200, "end\_bucket": 299, "experiment\_id": "exp\_C", "status": "ACTIVE"}。  
4. **写入**：将修改后的**整个**分配图原子化地写回配置中心。  
5. **失败处理**：如果没有找到足够大的连续区块，则分配请求失败，并向用户明确提示。这可以促使管理员进行碎片整理或规划。

#### **4.3.2. 结束一个实验**

1. **读取**：获取当前分配图。  
2. **定位**：找到代表已结束实验的条目。  
3. **释放**：将其 experiment\_id 设为 null，status 设为 AVAILABLE。  
4. **碎片整理（Defragmentation）**：这是至关重要的一步。检查该区块的**前一个**和**后一个**条目是否也处于 AVAILABLE 状态。如果是，则将这些连续的 AVAILABLE 区块**合并**成一个更大的 AVAILABLE 区块。  
5. **写入**：将修改后的分配图原子化地写回。

#### **4.3.3. 调整流量（渐进式发布）**

渐进式发布（Gradual Rollout）是现代功能交付的核心实践 19。我们的算法必须优雅地支持它，同时保证用户分配的稳定性。

1. **增加流量**：  
   * 定位到实验区块。检查其**紧邻的下一个**区块是否为 AVAILABLE 且空间足够。  
   * 如果是，则扩大实验区块的 end\_bucket，并同时缩小 AVAILABLE 区块的 start\_bucket。  
   * 如果下一个区块不是 AVAILABLE 或空间不足，则无法从此方向增加流量。  
2. **减少流量**：  
   * 缩小实验区块的 end\_bucket。  
   * 检查其紧邻的下一个区块是否为 AVAILABLE。如果是，则将释放出的桶合并到该 AVAILABLE 区块中（通过减小其 start\_bucket）。如果不是，则为释放出的桶创建一个新的 AVAILABLE 区块。

该逻辑确保了流量调整总是从范围的一端进行，从而保证了所有已在实验中的用户的分组不会改变。

### **4.4. 分配图状态演进表示例**

下表通过一系列连续操作，直观地展示了所提出的数据结构和算法如何维护数据完整性并有效对抗碎片化。

| 操作 | 实验需求 | 流量桶分配图 (JSON 表示) | 备注 |
| :---- | :---- | :---- | :---- |
| 初始状态 | 无 | \`\` | 流量层为空，存在一个完整的可用区块。 |
| 新增 Exp\_A (20%) | 200 个桶 | \`\` | 首次适应策略，切分初始区块。 |
| 新增 Exp\_B (30%) | 300 个桶 | \`\` | 从剩余的可用区块中分配。 |
| 结束 Exp\_A | \- | \`\` | Exp\_A 的区块被释放。此时出现了**碎片**。 |
| 新增 Exp\_C (40%) | 400 个桶 | \`\` | 算法选择最大的可用区块 \`\` 进行分配。 |
| 结束 Exp\_B | \- | \`\` | **关键步骤**：算法检测到 和 相邻且均可用，自动将它们**合并**为一个大的可用区块 \`\`。 |

---

## **第 5 节：高级架构考量与建议**

### **5.1. 支持全局对照组（Universal Holdouts）**

像 Uber 这样成熟的实验平台，通常会支持“全局对照组”的概念，用于衡量所有实验累积的、长期的影响 22。

实现建议：  
使用我们提出的范围分配模型可以非常优雅地实现这一点。只需在所有流量层的分配图中，永久性地保留一小段固定的桶区间（例如 \`\`）作为全局对照组。分流服务在处理请求时，可以优先检查用户是否落入此区间。如果是，则直接将其分配到全局对照组，并跳过该层后续的所有实验分配逻辑。此功能对核心设计改动极小，却能增加一种极其强大的分析能力。

### **5.2. 容错与降级策略**

实现建议：  
必须考虑当分流服务无法从配置中心获取分配图时的异常情况。系统必须能够安全地失败（fail safe）。推荐的降级策略是：

1. **首选策略**：使用一份本地缓存的、最近一次成功获取的“最后可用配置”（last-known-good）。这需要配合合理的缓存时间（TTL）和更新机制，以最大化系统可用性。  
2. **最终策略**：如果没有任何可用配置，则将所有用户分配到默认体验（即不进入任何实验）。

Uber 的系统在其 SDK 设计中也明确考虑了这种降级逻辑 3。

### **5.3. 可观测性与监控**

实现建议：  
分流服务必须被充分地埋点和监控，以洞察其运行状态。需要追踪的关键指标包括：

* **延迟（Latency）**：分流决策的 p50, p90, p99 延迟。  
* **错误率（Error Rate）**：获取配置失败、哈希计算错误等异常的发生率。  
* **分配计数（Allocation Counts）**：为每个实验的每个分组（包括控制组）设置一个计数器。通过监控这些计数器，可以实时验证流量是否按照预期比例分配，这对于及早发现“样本比例不匹配”（Sample Ratio Mismatch, SRM）等严重问题至关重要 23。  
* **配置版本（Configuration Version）**：在日志中记录每次分流决策所依据的配置版本号，这将极大地简化问题排查过程。

---

## **结论：构建一个弹性、可扩展的实验服务蓝图**

总结而言，您所设计的基于双重哈希和流量分层的核心分流架构，是一个极为出色的起点。它在设计理念上与业界最佳实践保持一致，为构建一个高性能、高可靠的实验平台奠定了坚实的基础。

本报告的核心贡献在于，针对您提出的动态分配挑战，提供了一套完整的、可落地的解决方案。这套方案以**基于范围的分配图**为核心数据结构，通过**原子化的配置更新**和**包含碎片整理的生命周期管理算法**，将一个优秀的设计理念，提升为了一个能够应对复杂生产环境挑战的、稳健的工程系统。

将这些建议付诸实践，您将构建的不仅仅是一个分流服务，而是一个能够支撑组织进行可信、高速、数据驱动决策的强大引擎。

#### **引用的著作**

1. Building a Trustworthy A/B Testing Platform — Practical Guide and ..., 访问时间为 九月 1, 2025， [https://medium.com/@chuan-zhang/building-a-trustworthy-a-b-testing-platform-practical-guide-and-an-architecture-demonstration-332446724ba0](https://medium.com/@chuan-zhang/building-a-trustworthy-a-b-testing-platform-practical-guide-and-an-architecture-demonstration-332446724ba0)  
2. It's All A/Bout Testing: The Netflix Experimentation Platform | by ..., 访问时间为 九月 1, 2025， [https://netflixtechblog.com/its-all-a-bout-testing-the-netflix-experimentation-platform-4e1ca458c15](https://netflixtechblog.com/its-all-a-bout-testing-the-netflix-experimentation-platform-4e1ca458c15)  
3. Supercharging A/B Testing at Uber | Uber Blog, 访问时间为 九月 1, 2025， [https://www.uber.com/blog/supercharging-a-b-testing-at-uber/](https://www.uber.com/blog/supercharging-a-b-testing-at-uber/)  
4. Converting stateful application to stateless using AWS services, 访问时间为 九月 1, 2025， [https://aws.amazon.com/blogs/architecture/converting-stateful-application-to-stateless-using-aws-services/](https://aws.amazon.com/blogs/architecture/converting-stateful-application-to-stateless-using-aws-services/)  
5. Stateful vs. Stateless Applications: What's the Difference? \- Pure Storage Blog, 访问时间为 九月 1, 2025， [https://blog.purestorage.com/purely-educational/stateful-vs-stateless-applications-whats-the-difference/](https://blog.purestorage.com/purely-educational/stateful-vs-stateless-applications-whats-the-difference/)  
6. Stateful vs stateless applications \- Red Hat, 访问时间为 九月 1, 2025， [https://www.redhat.com/en/topics/cloud-native-apps/stateful-vs-stateless](https://www.redhat.com/en/topics/cloud-native-apps/stateful-vs-stateless)  
7. A Guide to Stateful and Stateless Applications Best Practices \- XenonStack, 访问时间为 九月 1, 2025， [https://www.xenonstack.com/insights/stateful-and-stateless-applications](https://www.xenonstack.com/insights/stateful-and-stateless-applications)  
8. A Deep Dive into Stateless and Stateful Services: Architecting Scalable, Efficient Applications | by Scaibu, 访问时间为 九月 1, 2025， [https://scaibu.medium.com/a-deep-dive-into-stateless-and-stateful-services-architecting-scalable-efficient-applications-d6eddf795edf](https://scaibu.medium.com/a-deep-dive-into-stateless-and-stateful-services-architecting-scalable-efficient-applications-d6eddf795edf)  
9. MurmurHash: The Scrappy Algorithm That Secretly Powers Half the Internet \- Medium, 访问时间为 九月 1, 2025， [https://medium.com/@thealonemusk/murmurhash-the-scrappy-algorithm-that-secretly-powers-half-the-internet-2d3f79b4509b](https://medium.com/@thealonemusk/murmurhash-the-scrappy-algorithm-that-secretly-powers-half-the-internet-2d3f79b4509b)  
10. Hashing and Validation of MurmurHash in Go Implementation \- MojoAuth, 访问时间为 九月 1, 2025， [https://mojoauth.com/hashing/murmurhash-in-go/](https://mojoauth.com/hashing/murmurhash-in-go/)  
11. MurmurHash \- Wikipedia, 访问时间为 九月 1, 2025， [https://en.wikipedia.org/wiki/MurmurHash](https://en.wikipedia.org/wiki/MurmurHash)  
12. Principles behind Traffic Distribution \- Getting Started \- Hackle, 访问时间为 九月 1, 2025， [https://docs-en.hackle.io/docs/ab-bucketing](https://docs-en.hackle.io/docs/ab-bucketing)  
13. Mutually exclusive experiments – Support Help Center, 访问时间为 九月 1, 2025， [https://support.optimizely.com/hc/en-us/articles/4410289064205-Mutually-exclusive-experiments](https://support.optimizely.com/hc/en-us/articles/4410289064205-Mutually-exclusive-experiments)  
14. Mutually Exclusive Experiments: Preventing the Interaction Effect \- AB Tasty, 访问时间为 九月 1, 2025， [https://www.abtasty.com/blog/mutually-exclusive-experiments/](https://www.abtasty.com/blog/mutually-exclusive-experiments/)  
15. Intro to A/B Testing — Do it right, do it always | by Richard Chen | Medium, 访问时间为 九月 1, 2025， [https://medium.com/@richardchen\_81235/intro-to-a-b-testing-do-it-right-do-it-always-620c81d5f0fc](https://medium.com/@richardchen_81235/intro-to-a-b-testing-do-it-right-do-it-always-620c81d5f0fc)  
16. 1\. Layered Architecture \- Software Architecture Patterns \[Book\] \- O'Reilly Media, 访问时间为 九月 1, 2025， [https://www.oreilly.com/library/view/software-architecture-patterns/9781491971437/ch01.html](https://www.oreilly.com/library/view/software-architecture-patterns/9781491971437/ch01.html)  
17. A/B Testing — What are the Solution Design and Engineering Processes for Your Startup? | by Dinh-Cuong DUONG | Problem Solving Blog | Medium, 访问时间为 九月 1, 2025， [https://medium.com/problem-solving-blog/a-b-testing-what-is-the-solution-for-your-web-and-mobile-apps-part-1-68faf132fceb](https://medium.com/problem-solving-blog/a-b-testing-what-is-the-solution-for-your-web-and-mobile-apps-part-1-68faf132fceb)  
18. A/B Test Bucketing using Hashing \- Engineering at Depop, 访问时间为 九月 1, 2025， [https://engineering.depop.com/a-b-test-bucketing-using-hashing-475c4ce5d07](https://engineering.depop.com/a-b-test-bucketing-using-hashing-475c4ce5d07)  
19. A/B Testing \- PromptLayer, 访问时间为 九月 1, 2025， [https://docs.promptlayer.com/why-promptlayer/ab-releases](https://docs.promptlayer.com/why-promptlayer/ab-releases)  
20. Gradually rolling out A/B tests to your audience \- Leanplum Documentation, 访问时间为 九月 1, 2025， [https://docs.leanplum.com/docs/gradually-rolling-out-ab-tests-to-your-audience](https://docs.leanplum.com/docs/gradually-rolling-out-ab-tests-to-your-audience)  
21. Faster and Safer Releases with AB Tasty Rollouts, 访问时间为 九月 1, 2025， [https://www.abtasty.com/rollouts/](https://www.abtasty.com/rollouts/)  
22. Under the Hood of Uber's Experimentation Platform | Uber Blog, 访问时间为 九月 1, 2025， [https://www.uber.com/blog/xp/](https://www.uber.com/blog/xp/)  
23. How can I create an A/B testing framework? \- Kameleoon, 访问时间为 九月 1, 2025， [https://www.kameleoon.com/blog/ab-testing-frameworks](https://www.kameleoon.com/blog/ab-testing-frameworks)