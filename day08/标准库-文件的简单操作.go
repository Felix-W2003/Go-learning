package main
import (
	"fmt"
	"os"
)


//创建文件
func createFile(){
	f,err := os.Create("test02.txt")
	if err != nil{
		fmt.Printf("err:%v\n",err)
	} else {
		fmt.Printf("f:%v\n",f)
	}
}


//重命名文件
func renameFile(){
	err :=  os.Rename("test.txt","test01.txt")
	if err != nil{
		fmt.Printf("err:%v\n",err)
	}
}



//写文件
func writeFile(){
	s := "hello world"
	os.WriteFile("test03.txt",[]byte(s),os.ModePerm)

}



func main(){

	createFile()
	renameFile()
	writeFile()
}