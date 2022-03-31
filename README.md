### 停止groutine
+ channel close机制
+ channel 轮询-select 
+ channel + context
```
More @ https://geektutu.com/post/hpg-timeout-goroutine.html
https://www.cnblogs.com/kcxg/p/15064297.html
```

epoll和select区别
  (1)select==>无差别轮询，时间复杂度O(n)
     它仅仅知道了，有I/O事件发生了，却并不知道是哪那几个流（可能有一个，多个，甚至全部），我们只能无差别轮询所有流，找出能读出数据，或者写入数据的流，对他们进行操作。

  (2)poll==>时间复杂度O(n)
     poll本质上和select没有区别，它将用户传入的数组拷贝到内核空间，然后查询每个fd对应的设备状态， 但是它没有最大连接数的限制，原因是它是基于链表来存储的.

  (3)epoll==>时间复杂度O(1)
     map底层数据结构，时间复杂度。事件驱动（每个事件关联上fd）的
	 poll只要一次拷贝，利用mmap()文件映射内存加速与内核空间的消息传递
