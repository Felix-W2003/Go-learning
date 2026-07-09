/* 在go程序中，使用go关键字还可以为匿名函数创建goroutine。
注意，go关键字的额后面需要包含两个内容：一个是定义的匿名函数；另一个是
匿名函数的调用参数
*/

package main

import (
	"fmt"
	"time"
)
func main(){
	go func(){
		for i:=0;i<3;i++{
			fmt.Println("左手画圆")
			time.Sleep(1*time.Second)
		}
	}()//匿名函数被调用时所需设置的参数



	go func(color string){
		for i:=0;i<3;i++{
			fmt.Println("左手画圆，圆的颜色是"+color)
			time.Sleep(1*time.Second)
		}
	}("红色")


	for i:=0;i<3;i++{
		fmt.Println("右手画方")
		time.Sleep(1*time.Second)
	}
}