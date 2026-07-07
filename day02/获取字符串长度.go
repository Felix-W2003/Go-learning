package main
import "fmt"

func main() {

	str1 := "hello,go"
	fmt.Println(len(str1)) //获取字符串的长度
	fmt.Println(str1[0])
	fmt.Println(len([]rune(str1))) //获取字符串的长度
	fmt.Println("--------------------")

	str2 := "hello,世界"
	fmt.Println(len(str2))
	fmt.Println(str2[0])
	fmt.Println(len([]rune(str2)))
}