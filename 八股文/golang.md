## 同步原语
Golang的提供的同步机制有sync模块下的Mutex、WaitGroup以及语言自身提供的chan等。 这些同步的方法都是以runtime中实现的底层同步机制（cas、atomic、spinlock、sem）为基础的。
### 一、mutex
互互斥锁，并发程序中对共享资源进行访问控制的主要手段，由状态和信号量组成。方法Lock()和Unlock()分别用于加锁和解锁

##### 1.mutex结构
状态(互斥锁)+ 信号量组成
```go
type Mutex struct {
    state int32
    sema  uint32
}
```
#### 2.Mutex 状态
   + Locked: 表示该Mutex是否已被锁定，0：没有锁定 1：已被锁定。
   + Woken: 表示是否有协程已被唤醒，0：没有协程唤醒 1：已有协程唤醒，正在加锁过程中。
   + Starving：表示该Mutex是否处理饥饿状态， 0：没有饥饿 1：饥饿状态，说明有协程阻塞了超过1ms。
   + Waiter: 表示阻塞等待锁的协程个数，协程解锁时根据此值来判断是否需要释放信号量

#### 3.mutex 工作模式
 + 正常模式：锁的等待者会按照先进先出的顺序获取锁
 + 饥饿模式： 互斥锁会直接交给等待队列最前面的 Goroutine。新的 Goroutine 在该状态下不能获取锁、也不会进入自旋状态，它们只会在队列的末尾等待。如果一个 Goroutine 获得了互斥锁并且它在队列的末尾或者它等待的时间少于 1ms，那么当前的互斥锁就会切换回正常模式。

#### 4.加锁和解锁。
+ 加锁
```markdown
如果互斥锁处于初始化状态，会通过置位 mutexLocked 加锁；
如果互斥锁处于 mutexLocked 状态并且在普通模式下工作，会进入自旋，执行 30 次 PAUSE 指令消耗 CPU 时间等待锁的释放；
如果当前 Goroutine 等待锁的时间超过了 1ms，互斥锁就会切换到饥饿模式；
互斥锁在正常情况下会通过 runtime.sync_runtime_SemacquireMutex 将尝试获取锁的 Goroutine 切换至休眠状态，等待锁的持有者唤醒；
如果当前 Goroutine 是互斥锁上的最后一个等待的协程或者等待的时间小于 1ms，那么它会将互斥锁切换回正常模式；

```

+ 解锁
```markdown
当互斥锁已经被解锁时，调用 sync.Mutex.Unlock会直接抛出异常；
当互斥锁处于饥饿模式时，将锁的所有权交给队列中的下一个等待者，等待者会负责设置 mutexLocked 标志位；
当互斥锁处于普通模式时，如果没有 Goroutine 等待锁的释放或者已经有被唤醒的 Goroutine 获得了锁，会直接返回；在其他情况下会通过 sync.runtime_Semrelease唤醒对应的 Goroutine；
```

#### 5.读写锁 RWMutex
细粒度的互斥锁
```go
type RWMutex struct {
    w           Mutex   //复用互斥锁提供的能力；
    writerSem   uint32  //读等待
    readerSem   uint32  //写等待
    readerCount int32   //正在进行读操作的数量
    readerWait  int32   //写操作被阻塞时等待的读操作个数
}
```
工作原理
 + 调用 sync.RWMutex.Lock尝试获取写锁时；
    + 每次 sync.RWMutex.RUnlock都会将 readerCount 其减一，当它归零时该 Goroutine 会获得写锁；
    + 将 readerCount 减少 rwmutexMaxReaders 个数以阻塞后续的读操作；
 + 调用 sync.RWMutex.Unlock 释放写锁时，会先通知所有的读操作，然后才会释放持有的互斥锁；

### 二、WaitGroup
常用场景： 批量发出 RPC 或者 HTTP 请求
sync.WaitGroup将原本顺序执行的代码在多个 Goroutine 中并发执行，提高程序运行效率
#### 1.结构体
```go
type WaitGroup struct {
    noCopy noCopy //保证 sync.WaitGroup不会被开发者通过再赋值的方式拷贝
    state1 [3]uint32 //状态和信号量
}
```
#### 2.接口
 + sync.WaitGroup.Add
 + sync.WaitGroup.Wait
 + sync.WaitGroup.Done (add中传-1)

### 自旋锁
自旋锁是指当一个线程在获取锁的时候，如果锁已经被其他线程获取，那么该线程将循环等待，然后不断地判断是否能够被成功获取，直到获取到锁才会退出循环。
获取锁的线程一直处于活跃状态 Golang中的自旋锁用来实现其他类型的锁,与互斥锁类似，不同点在于，它不是通过休眠来使进程阻塞，而是在获得锁之前一直处于活跃状态(自旋)。

#### 与信号量区别
 + 自旋锁适合于保持时间非常短的情况，它可以在任何上下文使用；
 + 信号量适合于保持时间较长的情况，会只能在进程上下文使用。

#### 使用场景
我们在使用 Redis 对数据库中的数据进行缓存，发生缓存击穿时，大量的流量都会打到数据库上进而影响服务的尾延时。
```go
type service struct {
    requestGroup singleflight.Group
}

func (s *service) handleRequest(ctx context.Context, request Request) (Response, error) {
    v, err, _ := requestGroup.Do(request.Hash(), func() (interface{}, error) {
        rows, err := // select * from tables
        if err != nil {
            return nil, err
        }
        return rows, nil
    })
    if err != nil {
        return nil, err
    }
    return Response{
        rows: rows,
    }, nil
}
```
### chan

channel是一个数据类型，主要用来解决go程的同步问题以及协程之间数据共享（数据传递）的问题
内部实现同步，确保并发安全
#### chan 底层结构
```go
type hchan struct {
    qcount   uint           // total data in the queue
    dataqsiz uint           // 环形队列的长度，缓冲区大小
    buf      unsafe.Pointer // 缓冲的channel所特有的结构，用来存储缓存数据。是个循环链表
    elemsize uint16
    closed   uint32
    elemtype *_type // element type
    sendx    uint   // buf这个循环链表中的发送index
    recvx    uint   // buf这个循环链表中的接收的index
    recvq    waitq  // 发送(channel <- xxx)队列。是个双向链表
    sendq    waitq  // 接收(<-channel) 队列。是个双向链表
    lock mutex
}
```
#### chan 中mutex
   + 加锁
   + 把数据从goroutine中copy到“队列”中(或者从队列中copy到goroutine中）。
   + 释放锁

#### channel缓存满了
 + Go的调度模型GMP模型进行调度。让当前groutine等待，并从让出占用M，让其他G去使用
 + 同时将g1指针和send元素保存到sendq队列中等待被唤醒
 + 其他g2从缓存队列中取出数据，channel会将等待队列中的g1推出，将G1当时send的数据推到缓存中，然后调用Go的scheduler，唤醒G1，并把G1放到可运行的Goroutine队列中。

#### 有缓冲和无缓冲
 + 无缓冲：同一时刻，同时有 读、写两端把持 channel，否则会引起阻塞
 + 有缓冲：数据发送端，发送完数据，立即返回。数据接收端有可能立即读取，也可能延迟处理。
     + datasize>0, qcount==dataqsize表示buf已满

---
## GMP模型

--
## 垃圾回收
golang 使用三色标记法对栈或者指针进行全局扫描，进行垃圾回收，解决了引用计数缺点(循环引用)

### 回收过程中对象三种状态
初始化状态所有对象为白色
 + 灰色：对象还在标记队列中等待•
 + 黑色：对象已被标记，gcmarkBits 对应位为 1 -- 该对象不会在本次 GC 中被回收
 + 白色：对象未被标记，gcmarkBits 对应位为 0 -- 该对象将会在本次 GC 中被清理

golang 的垃圾回收算法属于 标记-清除，是需要 STW 的。为了缩短stw时间，引入写屏障和辅助GC

### 混合写屏障 和辅助GC
#### 混合写屏障
只需要在开始时并发扫描各个goroutine的栈，使其变黑并一直保持，这个过程不需要STW，而标记结束后，因为栈在扫描后始终是黑色的，也无需再进行re-scan操作了，减少了STW的时间

#### 辅助GC
为了防止内存分配过快，在 GC 执行过程中，GC 过程中 mutator 线程会并发运行，而 mutator assist 机制会协助 GC 做一部分的工作。

### 垃圾回收触发机制
1.内存分配量达到阈值：每次内存分配都会检查当前内存分配量是否达到阈值，如果达到阈值则触发 GC。即每当内存扩大一倍时启动 GC。
2.定时触发 GC：默认情况下，2分钟触发一次 GC，该间隔由 src/runtime/proc.go 中的 forcegcperiod 声明。
3.手动触发 GC：在代码中，可通过使用 runtime.GC() 手动触发 GC

### GC 优化建议
 + 减少对象分配个数， 
 + 采用对象复用、将小对象组合成大对象或 
 + 采用小数据类型（如使用 int8 代替 int）
