package main

import "fmt"

//在Go语言中，把没有名称的变量称做匿名变量，并且使用空白标识符（_）来表示匿名变量。
// 匿名变量可以用来接收函数的返回值，但不会被使用。
func getYearDatas() (int, int) {
	return 12, 24
}

func main() {
	monthNums, _ := getYearDatas()      //获取一年中的月份数量
	_, solarTermsNums := getYearDatas() //获取一年中的节气的数量
	fmt.Println("一年中有", monthNums, "个月份和", solarTermsNums, "个节气")
}
