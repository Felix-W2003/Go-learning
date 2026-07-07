package main
import ("fmt"
		"strconv")
// 方法
/*传统面向对象语言的方法是定义在类中，而结构体的方法是
定义在结构体之外的。通过将结构体和结构体方法分离，Go语言的代码更加
灵活。
和函数一样，结构体方法也是用func关键字定义。结构体方法和函数
的最大区别是结构体方法需要在func关键字和方法名之间使用小括号声明
一个变量作为方法的接收者。根据变量的类型，结构体的方法分为两种形式：
值接受者和方法和指针接受者方法
*/

//值接收者方法
type profile struct{
	name string
	age int

}

//定义结构体方法
func (p profile)get_info(subject string,score int)string{
	info :="姓名："+p.name+"\n年龄："+strconv.Itoa(p.age)
	info += "\n考试科目："+subject+"\n考试分数:"+strconv.Itoa(score)
	return info
}

func main(){
	//实例化
	p:=profile{name:"张三",age:20}
	//调用结构体方法
	result:=p.get_info("综合知识",76)
	fmt.Printf(result)
}

