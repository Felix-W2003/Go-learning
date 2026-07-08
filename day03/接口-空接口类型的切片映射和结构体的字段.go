package main

import "fmt"

func main() {
	// 初始化空接口类型切片，可存放任意类型元素
	s := []interface{}{123, "hello", true}
	fmt.Printf("切片中的数据：%v\n", s)

	// 初始化value为空接口的map
	m := map[string]interface{}{}
	m["product"] = "cap"
	m["price"] = 15
	fmt.Printf("集合中的数据：%v\n", m)

	// 定义内部包含空接口字段的结构体变量
	var car struct {
		price interface{}
	}
	// 赋值int类型
	car.price = 159800
	fmt.Printf("结构体中的字段数据：%v，其数据类型：%T\n", car.price, car.price)
	// 重新赋值string类型
	car.price = "壹拾伍万玖仟捌佰圆整"
	fmt.Printf("结构体中的字段数据：%v，其数据类型：%T\n", car.price, car.price)
}