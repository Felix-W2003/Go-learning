/*
通过fmt包的Errorf()函数可以格式化创建描述性的错误信息。
*/


package main
import (
		"fmt"
		"time"
)

func main(){
	const name,id = "Tom",10

	//自定义错误信息
	err := fmt.Errorf("用户%q(id%d)不存在",name,id)
	fmt.Println(err)
	fmt.Printf("错误发生时间:%v",time.Now())
}
