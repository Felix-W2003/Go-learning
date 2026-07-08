package main
import "fmt"

func main(){
	var a,b interface{}
	a = 711
	b = "实践出真知"
	value_a,ok_a := a.(int)
	fmt.Printf("a的值为%v,数据类型是int类型:%v\n",value_a,ok_a)
	value_b,ok_b := b.(int)
	fmt.Printf("b的值为%v,数据类型是int类型:%v\n",value_b,ok_b)
}