package main
import ("fmt"
)

//定义结构体
type divide struct {
	dividend float64
	divisor	 float64
}



//定义结构体Error()函数
func (d divide)Error()string{
	//自定义错误信息
	err := `
	被除数：%v
	除数：0
	错误：除数不能为0
	`
	return fmt.Sprintf(err,d.dividend)

}

func divideBy(dividend float64,divisor float64)(float64,error){
	if divisor == 0{
		return -1,divide{
			dividend:dividend,
			divisor:divisor,
		}
	}
	return dividend/divisor,nil
}

func main(){
	result,err:= divideBy(6,0)

	if err!= nil{
		fmt.Print(err)
	} else {
		fmt.Print(result)
	}
}