package main 
import "fmt"

func main() {

	var num *int32 = new(int32)
	*num = 10
	fmt.Println("指针变量num指向的内存地址：",num)
	fmt.Println("指针变量num指向的内存地址的值：",*num)
//实际开发中并不经常使用内置函数new（）初始化指针变量，这是因为通过内置函数new()得到的是指向
//某个 类型的指针 ，如果不为这个指针赋值，那么这个指针的值就是这个
//类型的默认值



//以下是初始化指针变量的常用方式
var number int =10
var ptr *int
ptr = &number
fmt.Println("指针变量ptr指向的内存地址：",ptr)
fmt.Println("指针变量ptr指向的内存地址的值：",*ptr)
//使用短变量声明的语法格式可以简化为
number1 := 10
ptr1 :=&number1
fmt.Println("指针变量ptr1指向的内存地址：",ptr1)
fmt.Println("指针变量ptr1指向的内存地址的值：",*ptr1)



str := "知识就是力量"
ptr2 := &str
fmt.Printf("指针变量ptr2类型:%T\n",ptr2)
fmt.Printf("指针变量ptr2指向的内存地址:%p\n",ptr2)
str_value := *ptr2
fmt.Printf("变量str_value的类型是：%T\n",str_value)
fmt.Printf("变量str_value的值是：%s\n",str_value)

}