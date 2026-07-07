package main
import "fmt"
/*结构体是一种数据类型，所以他也可以作为匿名成员。这时，访问嵌套
的结构体成员只需要使用“结构体实例名.成员名”即可*/
type exam struct {
	subject string
	score int
}
type info struct {
	name string 
	age int
	exam //匿名成员
}

func main (){
	//实例化结构体
	e := exam{subject:"机械制图",score:80}
	i := info{name:"江南",age:23,exam:e}
	//访问结构体成员
	fmt.Printf("姓名：%v\n",i.name)
	fmt.Printf("年龄：%v\n",i.age)
	fmt.Printf("考试科目：%v\n",i.subject)
	fmt.Printf("考试分数:%v\n",i.score)
	fmt.Println(i)
}
/*在上述代码中，通过结构体实例i可以直接访问exam结构体的成员subject和
score，不必使用i.exam.subject和i.exam.score

*/