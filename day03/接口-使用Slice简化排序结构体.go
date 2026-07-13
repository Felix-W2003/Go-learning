package main

import (
	"fmt"
	"sort"
)

type Car struct {
	Brand string
	Color string
}

func main() {
	cars := []*Car{
		{"Volkswagen", "black"},
		{"Toyota", "red"},
		{"Honda", "white"},
		{"LEXUS", "silver"},
		{"BMW", "king"},
		{"FORD", "blue"},
	}
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].Brand < cars[j].Brand
	})
	for _, v := range cars {
		fmt.Printf("%v\n", v)
	}
}

//空接口类型是指在接口中没有任何函数，用"interface{}"表示
//因为在空接口中没有任何函数，所以任何类型的变量都可以实现空接口
//也正因如此，在实际开发中，开发者既可以使用空接口保存任何类型的值，
//又可以从空接口中获取任何类型的值
