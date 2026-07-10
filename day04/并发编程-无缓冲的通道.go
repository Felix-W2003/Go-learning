
/* 在Go语言中，无缓冲的通道是指在接受数据前没有能力保存任何值的通道
为了能够让无缓冲的通道完成发送、接收数据的操作，执行发送数据操作的goroutine
和执行接受数据操作的goroutine需要同时准备就绪。也就是说，对于无缓冲的通道
执行发送和接受数据的操作是同步的，其中的任意一个操作都无法离开另一个操作单独存在
因此，无缓冲的通道又称同步通道
*/

package main
import "fmt"
/*
表示阅读的函数，参数有两个：
bookName: 图书的名称
chel: 通道的名称
*/

func read(bookName string, chel chan bool){
	fmt.Println("我正在读"+bookName)
	chel <- true
}

func main(){
	chel := make(chan bool)
	//执行并发程序
	go read("《Go语言从入门到精通》",chel)
	<-chel  //通道传输的数据被接受，但是接收到的数据将被忽略
}