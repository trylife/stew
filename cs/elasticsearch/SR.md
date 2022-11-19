# 容量规划

- 【推荐】 shard容量，业务型10-30GB(最佳)，日志型不超过100GB。 [[ref](https://www.elastic.co/guide/en/elasticsearch/reference/current/size-your-shards.html#use-ds-ilm-for-time-series)]

- 【参考】避免过渡分片，分片越多开销越大。

- 【强制】单个shard文档数量不超过2^31。

- 【强制】至少有1个副本(replicas)。

- # 索引设计

- 【推荐】动态映射仅适合数据探索阶段。[[ref](https://www.elastic.co/guide/en/elasticsearch/reference/current/dynamic-mapping.html)]

- 【推荐】使用别名aliases，方便不停机更换索引。[[ref](https://www.elastic.co/guide/en/elasticsearch/reference/current/aliases.html)]

- 【推荐】使用自动生成ID [[ref](https://www.elastic.co/guide/en/elasticsearch/reference/current/tune-for-indexing-speed.html#_use_auto_generated_ids)]

- 【参考】设置合理的 routing key (默认是ID)。

- 【参考】字段类型不支持修改必须谨慎。

- 【参考】ES只允许新增字段(lucene决定的)。

- ES中只储存要检索的字段，再通过ID去RDS中查。

- # 查询建议

- 【推荐】返回大结果集、深度分页使用scroll api。[[ref](https://www.elastic.co/guide/en/elasticsearch/reference/current/general-recommendations.html#large-size)]

- 【推荐】text类型仅用于分词检索，不适合聚合、排序。

- 【强制】不要在对text类型字段做wildc。

- 【强制】text类型fileddata会加大对内存的占用，不建议使用。

- # 数据初始化场景建议

- 【推荐】数据初始化时禁用副本(index.number_of_replicas=0) [[ref](https://www.elastic.co/guide/en/elasticsearch/reference/current/tune-for-indexing-speed.html#_disable_replicas_for_initial_loads)]

- 【推荐】禁用刷新时间refesh_interval = -1。

- 【推荐】使用批量写入bulk api，每批多少个最快需要自己先实验。[[ref](https://www.elastic.co/guide/en/elasticsearch/reference/8.5/docs-bulk.html)]

- 【推荐】使用ES的发布订阅模式快速导入数据。TODO REF

- # 运维建议

- 【推荐】使用SSD。

- 【推荐】配置slowlog。

- 【推荐】刷新频率refresh_interval，业务型1-5s，日志型30s - 120s。[[ref](https://www.elastic.co/guide/en/elasticsearch/reference/current/tune-for-indexing-speed.html#_unset_or_increase_the_refresh_interval)] 

- 【推荐】CPU内存比1:4或1:8。

- 【推荐】集群最大节点数 = 单节点CPU*5。

- 【推荐】建议shard的个数（包括副本）要尽可能等于数据节点数，或者是数据节点数的整数倍。

- 【推荐】保留一半的内存给文件系统缓冲I/O。[[ref](https://www.elastic.co/guide/en/elasticsearch/reference/current/tune-for-indexing-speed.html#_give_memory_to_the_filesystem_cache)]

- # 其他

- TODO 阶段最大内存与强制GC

- TODO 磁盘使用率与索引

# References