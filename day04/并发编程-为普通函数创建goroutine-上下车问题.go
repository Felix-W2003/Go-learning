/* 在Go程序中，使用go关键字为普通函数创建goroutine。需要特别注意的是，可以为
一个普通函数创建多个goroutine；但一个goroutine必定对应一个普通函数
当使用go关键字创建goroutine时，将忽略被调用函数的返回值
*/

package main
import (
	"fmt"
	"time"
)

func getOff(){
	//循环5次
	for i:=5;i>0;i--{
		fmt.Println("还有",i,"位乘客下车")
		//延时1s
		time.Sleep(1*time.Second)
	}
}

func main(){

	//执行并发程序
	go getOff()
	//循环5次
	for i:=1;i<6;i++{
		fmt.Println("第",i,"位乘客上车")
		time.Sleep(1*time.Second)
	}

}