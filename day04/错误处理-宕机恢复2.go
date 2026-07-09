/* 如果panic函数和recover函数一起使用，并且程序中的函数调用比较复杂，则在执行完对应的defer
语句后，程序退出当前函数并返回到调用处继续执行*/
package main
import "fmt"

func first(){
	fmt.Println("first函数开始")
	second()
	fmt.Println("first函数结束")
}


func second(){
	defer func(){
		recover()		//宕机恢复
	}()
	fmt.Println("second函数开始")
	third()
	fmt.Println("second函数结束")
}

func third(){
	fmt.Println("third函数开始")
	panic("Program crash")
	fmt.Println("third函数结束")
}

func main(){
	fmt.Println("程序开始")
	first()
	fmt.Println("程序结束")
}

/* 由运行结果可以看出，在third函数中触发宕机，程序执行second函数中的
defer语句，在执行完该语句后退出second函数，并返回调用该函数的first函数
继续执行*/