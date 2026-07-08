package main
import (
		"errors"
		"fmt"
		"math"
)

func Sqrt(num float64)(float64 ,error){
	if num <0{
		//自定义错误信息
		return -1,errors.New("错误：负数没有平方根！")
	}
	return math.Sqrt(num),nil
}


func main(){

	result,err := Sqrt(-2)
	if err!= nil{
		fmt.Print(err)
	} else {
		fmt.Print(result)
	}
}