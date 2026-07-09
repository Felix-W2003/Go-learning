package main 
import (
		"fmt"
		"time"
)



func getOff(names []string){
	// 遍历存储下车乘客的姓名的切片中的元素
	for i,name := range names{
		fmt.Println("第",i+1,"位乘客",name,"正在下车")
		//延时1s
		time.Sleep(1*time.Second)
	}
}

func main(){
	var offNames = [5]string{"David","Leon","Steven","James","Tom"}
	go getOff(offNames[:])

	var onNames = [5]string{"张三","李四","王五","赵六","周七"}

	for i,name := range onNames{
		fmt.Println("第",i+1,"位乘客",name,"正在上车")
		//延时1s
		time.Sleep(1*time.Second)
	}
}