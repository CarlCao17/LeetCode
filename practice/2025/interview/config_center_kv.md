# 配置中心

- 强一致
- KV 存储
- Watch 机制 实现近实时的更新

ETCD: revision, key-space
    - Establishing a Watch (gRPC streaming)
    - Streaming Events
    - Guarantees: totally ordered by revision, reliable, atomic changes
Consul: service mesh solution, includes a key-value store
    - Blocking query
    - Long Polling
    - Client loop
    - 
Zookeeper