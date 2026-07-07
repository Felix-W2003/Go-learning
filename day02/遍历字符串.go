package main
import "fmt"
import "strings"

func strEncry(r rune) rune {
	return r+1

}
//go语言中的strings.Map()函数可以将字符串中
// 的每个字符都传入一个函数中进行处理，并返回一个新的字符串。

func main(){
	str := "hello,go"
	fmt.Println("原始字符串:", str)
	//将字符串转换为rune切片
	newStr := strings.Map(strEncry, str)
	fmt.Println("加密后的字符串:", newStr)
}