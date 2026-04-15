package expression

import (
	"fmt"
	"time"
)

// goroutine(旧)
// GPM模型
// M指的是系统级线程
// P指的是能够使若干个G在恰当的时机与M对接，并得以运行的中介

// 数据同步
// channel

func ChannelExp() {
	// go test() // 当使用go关键字创建goroutine时，将忽略被调用函数的返回值
	//
	// var ch chan int
	// 无缓冲
	// ch1 := make(chan int)
	// 有缓存
	// ch2 := make(chan int, 5)

	ch := make(chan int)
	go func() {
		for i := range 5 {
			ch <- i
		}
		close(ch) // 关闭通道
	}()
	for data := range ch {
		println(data)
	}

	// select
	ch1 := make(chan int)
	select {
	case data := <-ch1:
		fmt.Println(data)
	case <-time.After(2 * time.Second):
		fmt.Println("out time")
	}
}

func SyncExp() {
	// sync.WaitGroup
	// var wg sync.WaitGroup
	// wg.Add(2)
	// wg.Done()
	// wg.Wait()

	//sync.Mutex
	// var mux sync.Mutex
	// mux.Lock()
	// mux.UnLock()
}

func ContextExp() {
	// ctx, cancle := context.WithCancel(context.Background())
	// ctx, cancle := context.WithTimeout(context.Background(), time.Second)
}
