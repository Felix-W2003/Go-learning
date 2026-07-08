/*
Error 的嵌套就是在定义的一个error中嵌套另一个error。因为error可以嵌套
所以每次嵌套时都可以提供新的错误信息，并且保留原来error。
Go语言扩展了fmt包的Errorf()函数，为该函数增加%w占位符生成Error嵌套。
*/

package main
import (
	"errors"
	"fmt"
)

func main(){

	err := errors.New("结果集中没有行")
	err2:= fmt.Errorf("找不到数据，%w",err)
	fmt.Println(err2)

	/*既然error可以嵌套生成一个新的error，那他也可以被分解，通过errors包中
	的Unwrap()函数可以得到被嵌套的error
	*/
	err3 := errors.New("结果集中没有行")
	err4:= fmt.Errorf("找不到数据，%w",err3)
	fmt.Println(errors.Unwrap(err4))

	/*通过errors包中的Is()函数可以判断两个error是否是同一个error*/
	err5 := errors.New("结果集中没有行")
	err6:= fmt.Errorf("找不到数据，%w",err5)
	fmt.Println(errors.Is(err6,err5))
}