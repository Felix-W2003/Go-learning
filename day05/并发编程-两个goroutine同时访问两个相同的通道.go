/*select 语法格式与switch语法格式几乎相同，都是由一系列的case语句
和一个可选的default语句组成，但是select只适用于通道
每个case语句都指定一次通信。指定一次通信是指使用一个通道发送或接收数据。
当触发一个case语句时，将执行这个case语句；如果触发多个case语句，则随机
执行一个case语句

当两个goroutine同时访问这两个相同的通道时，在两个goroutine中必须以
相同顺序访问这两个通道

*/

package main 
import "fmt"

func main(){
	ch1 := make(chan int)
	ch2 := make(chan int)
	//为匿名函数创建goroutine
	go func(){
		v := 1
		ch1 <- v
		v2 := <-ch2
		fmt.Println(v,v2)
		fmt.Println("zhixing2")
	}()

	v := 2
	var v2 int
	select {
	case ch2 <- v: 
		fmt.Println("zhixing1")	//使用通道ch2发送变量v
	case v2 = <-ch1: //使用变量v2接受通道ch1传输的数据
		fmt.Println("zhixing3")
	
	}
	fmt.Println("zhixing4")
	fmt.Println(v,v2)
}