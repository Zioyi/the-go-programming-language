# Go中的并发
https://tech.bytedance.net/articles/6943848264026292237

## Channel
通道是协程间沟通的正确方式

[image](https://tech-proxy.bytedance.net/tos/images/1616742446977_59f89b095fa77d4dc5f6a7e6658052e2)

特性：
- 本身是线程安全的，这也是多线程通信的保证；
- 数据结构类似FIFO队列，保证收发数据的顺序；


## Select
select与for一起使用，达到在循环中监听channel的效果。
```go
for {
    select {
    case <-done:
        return
    default:
        // Do your work
    }
}
```
        
使用协程时注意协程泄露的问题
