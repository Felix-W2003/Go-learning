package main
import (
	"fmt"
)


//表示“接通电源”的函数，参数是只能发送数据的通道chel
func electrify(chel chan<-string){
	chel <- "接通电源！"
	close(chel)
}

/*
表示 “启动（发动机）”的函数，参数有两个：
ch1: 只能接受数据的通道
ch2: 只能发送数据的通道
*/
func start(ch1 <- chan string,ch2 chan<- string){
	data := <-ch1 //变量data接受通道ch1传输的数据
	ch2 <- data + "启动！"
	close(ch2)
}

//表示"驾驶"的函数，参数是只能发送数据的通道ch
func drive(ch <- chan string){
	for data:= range ch{
		fmt.Println(data+"准备就绪，开始行万里路.....")
	}
}

func main(){
	ch1 := make(chan string)
	ch2 := make(chan string)
	//执行并发程序
	go electrify(ch1)
	go start(ch1,ch2)
	//调用表示"驾驶"的函数
	drive(ch2)
}
