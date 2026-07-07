package main
import "fmt"

func modifyValue(number *int){
	*number = 11
	
}


//使用指针交换两个变量的值
func  exchangeValue(a *int,b *int){
	var temp int
	temp = *a
	*a = *b
	*b = temp
}
func main(){
	var number int =7
	modifyValue(&number)
	fmt.Println("number的值是:",number)



	var a int =10
	var b int =20
	exchangeValue(&a,&b)
	fmt.Println("a的值是:",a)
	fmt.Println("b的值是:",b)
}