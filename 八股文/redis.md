Redis中的值对象都是由 redisObject 结构来表示。
Redis 会使用一个全局哈希表保存所有键值对，哈希表的每一项是一个 dictEntry 的结构体，用来指向一个键值对

### 数据结构
https://www.cnblogs.com/xuxh120/p/14960156.html

+ 组成sds+redisObject
+ dictEntry 结构中有三个 8 字节的指针，分别指向 key、value 以及下一个 dictEntry，三
+ 实际是分配32个字节，因为Redis 使用的内存分配库 jemalloc
```

typedef struct dictht {
    dictEntry **table; //数组指针，每个元素都是一个指向dictEntry的指针
    unsigned long size; //表示这个dictht已经分配空间的大小，大小总是2^n
    unsigned long sizemask;//sizemask = size - 1; 是用来求hash值的掩码，为2^n-1
    unsigned long used; //目前已有的元素数量
} dictht

typedef struct dictEntry {
    void *key; //key void*表示任意类型指针
    union {//联合体中对于数字类型提供了专门的类型优化
        void *val;
        uint64_t u64;
        int64_t s64;
        double d;
    } v;
    struct dictEntry *next; //next指针，用拉链法解决哈希冲突
} dictEntry;
```

#### 内存占用
数据存储占用=数据大小+dictEntry大小+根据数据类型来看redisObject占用内存

---
### string 底层结构
sds + redisObject 

### string 内部编码方式
embstr一次分配内存，raw两次分配内存空间(回收同理)。 然后是使用了引用计数技术来共享对象。
#### Int
  + 保存long 型的64位有符号整数
  + 直接存redisObject
  + RedisObject 中的指针就直接赋值为整数数据了,不需要额外的指针空间指向数据
#### embstr
  + 字符串值的长度<=39字节，使用 embstr 编码的方式来保存这个字符串值。
#### raw
  + 字符串值的长度>39字节，使用一个简单动态字符串（SDS）来保存这个字符串值， 并将对象的编码设置为 raw

### embstr编码方式

 + 字符串<=32 字节：redisObject 中的元数据、指针和 SDS 是一块连续的内存区域，这样就可以避免内存碎片
### 压缩列表

Redis 还对 Long 类型整数和 SDS 的内存布局作了专门的设计。

### string 扩容机制
  + 扩容*2倍
  + >1MB，则按照1MB来
---

### key优化
#### case1:
```
一是请求的链路优化，二是cpu核专用内存
减少网络请求的次数，这个需要考虑，比如做多级缓存，还要注意类似于 mget 的那种问题
控制kv 大小
对大key做压缩，value头部预留压缩标记位就行了。
热key做二级缓存，避免频繁解压导致cpu过高问题。
```
#### case2: 热key处理
```
redis之前加一个代理模块，在db访问前切入一个热点统计分析模块。
代理模块：负责热key读写分离，
  1.读:热key出现后二级缓存热key和自动扩容，
  2.写:通过路由分配写入redis.同步策略，要看业务是要求强一致性，还是要求高可用性再做决定。
热点统计分析模块： 负责统计分析预测热key出现。
  1. redis本身还需要做好热点预警和节点自动扩充。
```

---

### 跳表 -skipelist
跳表是分层的有序链表，底层是真正保存的数据(传统的单链表)，而上层可以理解为下层节点的索引

#### 特点：
+ 使用二分查找法，复杂度O(logn);
+ 支持区间搜索
+ 第k层的索引个数为n/(2^k) -->n为链表节点数

#### 结构图：
例如 有数据[1,12,20,32,45,56,100] n=7 
 + 第1层索引个数7/2=3
 + 第2层索引个数7/4=1
```markdown
第2层： head  —————————————————>32————————————————————>tail
        |                       |                     |
第1层： head  ————————>12———————>32—————————>56———————>tail
        |             |         |          |          |
第0层：  head ———>1———>12———>20——>32———>45——>56——>100——>ail
```
#### 链表操作crud流程
+ insert: 插入15，首先找到离15最近，比15小的节点12，然后将节点插入底层（单链表的插入操作）。判断是否需要分层，不需要返回。需要的话自底向上一层一层插入节点
+ remove: 删除15，首先找到离15最近，比15小的节点12，然后将节点在底层删除，如果有上层节点，继续删除。
+ get: 从顶层head查找，若下一节点<查找元素,则右移；若>查找元素或者nil，则下移一层；重复直到找到12

#### golang代码实现

#### 应用场景：
+ redis 有序集合zset
  
  ``HashMap和跳跃表(SkipList)来保证数据的存储和有序，
  HashMap里放的是成员到score的映射，
  跳跃表里存放的 是所有的成员 ``
+ leveldb
+ HBase MemStore 内部存储
+ Lucene, elasticSearch


