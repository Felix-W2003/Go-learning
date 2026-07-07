/*GO语言中没有构造函数的概念，但是可以通过结构体初始化的过程
可以模拟构造函数，也就是以构造函数的方式实例化结构体，并在实例化过程中
为结构体的成员赋值。函数的返回值是结构体的实例，并以指针形式表示
*/
//不带参数的构造函数
package main
import "fmt"
type profile struct {
	name string
	age int
}

func NewProfile()*profile{
	//定义变量，作为结构体的成员值
	name := "Felix"
	age := 18
	//初始化结构体
	p:=profile{
		name:name,
		age:age,
	}
	return &p
}

func NewProfile1(name string,age int)*profile{
	//初始化结构体
	p:=profile{
		name:name,
		age:age,
	}
	return &p
}
func main(){

	//调用函数
	p := NewProfile()
	fmt.Println(p)
	fmt.Println(p.name)
	fmt.Println(p.age)


	g := NewProfile1("张三",29)
	fmt.Println(g)
	fmt.Println(g.name)
	fmt.Println(g.age)

}