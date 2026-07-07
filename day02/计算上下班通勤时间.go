package main
import "fmt"

var totalTime int 

func getTotalTime(time_s int ,time_w int) int{
		totalTime = time_s + time_w
		return totalTime
}

func main() {
	fmt.Println("总通勤时间:", getTotalTime(30, 45))


	var str = `总通勤时间: 
	地铁时间: 30分钟
	步行时间: 45分钟`
	fmt.Println(str)
}