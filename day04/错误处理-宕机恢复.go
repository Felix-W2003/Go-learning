package main 
import "fmt"

/* 使用recover可以在宕机后让程序继续执行。recover是Go语言的内置函数
该函数可以捕获panic信息，类似于其他编程语言中用于捕获异常的try...catch语句
recover通常在使用defer语句的函数中执行*/
func test(){
	defer func(){
		err := recover()
		if err != nil{
			fmt.Println(err)
		}
		fmt.Println("恢复执行")
	}()
	panic("程序崩溃")//触发宕机
}


func main(){

	fmt.Println("程序开始")
	test()
	fmt.Println("程序结束")
}
/* 由运行结果可以看出，通过recover函数可以获取panic信息，后面的程序可以正常
按顺序执行
*/