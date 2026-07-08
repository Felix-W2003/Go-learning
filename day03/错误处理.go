/* 错误处理在每个编程语言中都是重要的内容，在Go语言中也不例外
Go语言通过内置的错误接口提供非常简单的错误处理机制。error类型是一个
接口类型，Go语言的开发者可以在编码中通过实现error接口类型生成错误信息
error处理过程类似C语言中的错误码，可逐层返回，直到被处理
*/


package main
import ("fmt"
		"os")


func main(){
	f,err := os.Open("./book.txt")
	if err != nil{
		fmt.Print(err)
	} else {
		fmt.Print(f)
	}
}