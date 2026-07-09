/*通道是go语言在两个或多个goroutine之间的一种通信方式。
通道可以让一个goroutine给另一个goroutine发送消息。当需要在goroutine之间
共享一个数据资源，通道是确保同步交换数据资源的方法

多个goroutine为了争抢数据资源，势必会降低执行效率。为了保证执行效率不降低
goroutine之间通过通道进行通信，确保同一时刻只有一个goroutine访问通道，
并执行发送和接受数据的操作

通道就像队列一样，遵循先入先出的规则，保证发送和接受数据的顺序

在go语言中，通道是一种特殊的数据类型，通道需要一个类型修饰
这样通道的类型就是在其内部传输的数据的类型


使用通道接受和发送数据具有以下特性：
如果没有接受通道传输的数据，那么发送数据操作将被持续阻塞
如果接受了通道传输的数据，但尚未执行发送数据操作，那么接受数据操作将持续被阻塞

通道一次只能传输一个数据

*/


package main 
import (

	"fmt"
	"time"

)

func main(){
	chel := make(chan int)
	//为匿名函数创建goroutine
	go func(){
		for i:=0;i<6;i++{
			chel <- i //使用通道发送数字0~5
			time.Sleep(1*time.Second)
		}
	}()
	//遍历通道传输的数据
	for data := range chel {
		fmt.Println(data)
		if data == 5{
			break
		}
	}
}