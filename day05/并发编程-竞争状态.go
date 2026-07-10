/* 两个或多个goroutine在没有同步的情况下，访问某个共享的资源，这些goroutine
将处于竞争状态
为了让这些goroutine不进入竞争状态，可以使用同步goroutine机制对共享资源加锁
Go语言通过atomic包和sync包里的函数对共享资源执行加锁操作
*/
package main
import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg 		sync.WaitGroup	//声明同步等待组对象wg
)

func main(){
	wg.Add(2)
	go count(1)
	go count(2)
	wg.Wait()
	fmt.Println(counter)
}


func count(id int){
	for count := 0;count <2; count++{
		atomic.AddInt64(&counter,1)
		runtime.Gosched()
	}
	/*
		通知主程序当前并发程序已执行完毕，对wg执行减1操作
		当wg为0时，解除主程序的阻塞等待
	*/
	wg.Done()
}