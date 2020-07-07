1、并行和并发
并行(parallel)：指在同一时刻，有多条指令在多个处理器上同时执行。

并发(concurrency)：指在同一时刻只能有一条指令执行，但多个进程指令被快速的轮换执行，使得在宏观上具有多个进程同时执行的效果，但在微观上并不是同时执行的，只是把时间分成若干段，通过cpu时间片轮转使多个进程快速交替的执行

并行和并发的区别：
- 并行是两个队列同时使用两台咖啡机 （真正的多任务）
- 并发是两个队列交替使用一台咖啡机 （ 假 的多任务）

2、常用的并发编程技术
（1）、进程并发
- 程序和进程
程序，是指编译好的二进制文件，在磁盘上，不占用系统资源(cpu、内存、打开的文件、设备、锁....)
进程，是一个抽象的概念，与操作系统原理联系紧密。进程是活跃的程序，占用系统资源。在内存中执行。(程序运行起来，产生一个进程)

- 进程状态 
进程基本的状态有5种。分别为初始态，就绪态，运行态，挂起态与终止态。

3、Go并发
Go语言中的并发程序主要使用两种手段来实现。goroutine和channel。

- Goroutine的创建

package main

import("fmt" "time")

func newTask(){
    i := 0
    for{
        i++
        fmt.Printf("new goroutine:i= = %d\n",i)
        time.Sleep(1 * time.Second)
    }
}
func main(){
    go newTask
    i := 0
    for {
        i++
        fmt.Printf("main goroutine:i= %d\n",i)
        time.Sleep(1 * time.Second)
    }
}

- runtime包
runtime.Gosched( )时会暂停向下执行，直到其它协程执行完后，再回到该位置，主协程继续向下执行。

runtime.Goexit() 将立即终止当前 goroutine 执⾏，调度器确保所有已注册 defer延迟调用被执行。

runtime.GOMAXPROCS() 用来设置可以并行计算的CPU核数的最大值，并返回之前的值。

- channel

channel是一个数据类型，主要用来解决协程的同步问题以及协程之间数据共享（数据传递）的问题。

（1）定义channel变量

 make(chan Type)  //等价于make(chan Type, 0)
 make(chan Type, capacity)

 channel非常像生活中的管道，一边可以存放东西，另一边可以取出东西。channel通过操作符 <- 来接收
 和发送数据，发送和接收数据语法：
    channel <- value      //发送value到channel
    <-channel             //接收并将其丢弃
    x := <-channel        //从channel中接收数据，并赋值给x
    x, ok := <-channel    //功能同上，同时检查通道是否已关闭或者是否为空

默认情况下，channel接收和发送数据都是阻塞的，除非另一端已经准备好，这样就使得goroutine同步变的
更加的简单，而不需要显式的lock。

func main(){
    c := make(chan int)
    go func(){
        defer fmt.Println("子携程结束")
        fmt.Println("子携程正在运行")
        c <- 666
    }()
    num := <- c
}

(2)、无缓冲的channel（会阻塞其他线程）
缓冲的通道（unbuffered channel）是指在接收前没有能力保存任何值的通道。

无缓冲的channel创建格式：
make(chan Type) // //等价于make(chan Type, 0)

package main

import(
    "fmt"
    "time"
)
func mian(){
    c := make(chan int,0) //创建无缓冲的通道
    //内置函数 len 返回未被读取的缓冲元素数量，cap 返回缓冲区大小
    fmt.Printf("len(c) =%d,cap(c)=%d\n",len(c),cap(c))
    go func(){
       defer fmt.Println("子协程结束")
       for i := 0; i < 3 ;i++{
           c <- 1
            fmt.Printf("子协程正在运行[%d]: len(c)=%d, cap(c)=%d\n", i, len(c), cap(c))
       }
    }()
    time.Sleep(2 * time.Second)
    for i:=0;i < 3;i ++{
        num <- c
        fmt.Println("num = ",num)
    }
    fmt.Println("main协程结束")
}

（3）、有缓冲的channel
如果给定了一个缓冲区容量，通道就是异步的。只要缓冲区有未使用空间用于发送数据，或还包含可以接收的数据，
那么其通信就会无阻塞地进行。

func main(){
    c := make(chan int,3)//带缓冲的通道
     //内置函数 len 返回未被读取的缓冲元素数量， cap 返回缓冲区大小
    fmt.Printf("len(c)=%d, cap(c)=%d\n", len(c), cap(c))
    go func(){
        defer fmt.Println("子携程结束")

        for i:= 0; i <3;i++{
            c <- i
            fmt.Printf("子协程正在运行[%d]: len(c)=%d, cap(c)=%d\n", i, len(c), cap(c))
        }
    }()
    time.Sleep(2* time.Sleep)

    for i:= 0;i<3;i++{
        num := <- c
        fmt.Println("num =", num)
    }
    fmt.Println("main协程结束")
}
3、关闭channel


