/*在Go语言中，程序在编译时可能捕获到一些错误。有些错误只能在运行时出现
例如，数组访问越界、空指针引用等，这些运行时出现的错误都会引起宕机（panic）
当宕机发生时，程序就会停止运行，编译器输出对应的报错信息
如果在宕机后想让程序继续执行，则可以使用宕机恢复（recover）机制

panic()是Go语言的内置函数。它类似于其他编程语言中抛出异常的throw语句
panic()一般用在函数内部
*/

package main
import "fmt"
/*当调用函数执行到panic()时，不执行panic()后面的代码，如果在panic函数
前面有defer语句则正常执行该语句，之后返回调用函数，执行每一层的defer语句
直到所有正在执行的函数都被终止为止
*/
func test(){
	defer func(){
		fmt.Println("exit func test")
	}()
	panic("Program crash")//触发宕机
}
func main(){

	//手动触发宕机
//	panic("Program crash")
	defer func(){
		fmt.Println("exit func main")
	}()
	test()
	fmt.Println("不会执行")
}

