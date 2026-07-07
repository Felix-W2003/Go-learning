package main
import "fmt"


/* 定义结构体使用type和struct关键字，语法如下
type name struct {
	member datatype
	member datatype
	...
}
定义结构体后还要实例化才能使用结构体的成员。
实例化就是为结构体中的成员赋值。在Go语言中，结构体有多种实例化方法
根据实际需要可以选用不同的方法
*/

type profile struct {
	name string
	age int
	interest []string
}
func main() {
	//1.在结构体名称后直接赋值
	p1 := profile{name:"小明",age:18,interest:[]string{"篮球","足球"}}
	fmt.Println(p1)
	fmt.Println(p1.name)
	fmt.Println(p1.age)
	fmt.Println(p1.interest)

	//2.定义结构体变量再赋值
	var p2 profile
	p2.name = "小红"
	p2.age = 20
	p2.interest = []string{"游泳","跑步"}
	fmt.Println(p2)
	fmt.Println(p2.name)
	fmt.Println(p2.age)
	fmt.Println(p2.interest)	

	//3.使用new()函数实例化结构体
	p3 := new(profile)
	p3.name = "小刚"
	p3.age = 22
	p3.interest = []string{"羽毛球","乒乓球"}
	fmt.Println(p3)
	fmt.Println(p3.name)
	fmt.Println(p3.age)
	fmt.Println(p3.interest)

	//4.使用&符号实例化结构体
	p4 := &profile{}
	p4.name = "小李"
	p4.age = 24
	p4.interest = []string{"滑雪","滑冰"}
	fmt.Println(p4)
	fmt.Println(p4.name)
	fmt.Println(p4.age)
	fmt.Println(p4.interest)	
}


