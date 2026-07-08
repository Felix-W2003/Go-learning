package main

import (
	"fmt"
)

//自定义函数，计算长方形的面积
func area(length float64, width float64) (float64, error) {
	if length <= 0 {
		return -1, fmt.Errorf("长方形的长是%v，长不能小于或等于0", length)
	}
	if width <= 0 {
		return -1, fmt.Errorf("长方形的宽是%v，宽不能小于或等于0", width)
	}
	return length * width, nil
}
func main() {
	result, err := area(6, -3)               //调用函数
	if err != nil {
		fmt.Println(err)                    //输出错误信息
	} else {
		fmt.Printf("长方形的面积为：%v\n", result)  //输出计算结果
	}
	result, err = area(6, 3)                //调用函数
	if err != nil {
		fmt.Println(err)                    //输出错误信息
	} else {
		fmt.Printf("长方形的面积为：%v\n", result)
	}
}