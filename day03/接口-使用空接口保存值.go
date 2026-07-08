package main

import "fmt"

// noFunctions 是空接口类型（无任何方法），等价于 interface{}
type noFunctions interface{}

func main() {
	var nf noFunctions
	fmt.Printf("空接口保存的值：%v，空接口的类型：%T\n", nf, nf)
	num := 711
	nf = num
	fmt.Printf("空接口保存的值：%v，空接口的类型：%T\n", nf, nf)
	str := "hello, go"
	nf = str
	fmt.Printf("空接口保存的值：%v，空接口的类型：%T\n", nf, nf)
	b := true
	nf = b
	fmt.Printf("空接口保存的值：%v，空接口的类型：%T\n", nf, nf)
}