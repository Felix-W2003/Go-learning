package main
import "fmt"

//结构体也可以应用在在结构体的嵌套中。在嵌套
//结构体中使用匿名结构体的好处是可以简化定义多个结构体的步骤，使代码更简洁

type exam struct {
	subject string
	score int

}

type info struct{
	name string
	age int
	record exam
	//结构体嵌套中使用匿名结构体
	job struct {
		position string
		year int
	}
}

func main(){
	//实例化结构体
	e := exam{subject:"计算机技术与软件专业技术资格",score:100}
	i := info{name:"Felix",age:23,record:e}
	i.job.position="运维"
	i.job.year=1
	fmt.Printf("姓名:%v\n",i.name)
	fmt.Printf("年龄:%v\n",i.age)
	fmt.Printf("职位:%v\n",i.job.position)
	fmt.Printf("工作年限:%v\n",i.job.year)
	fmt.Printf("考试科目:%v\n",i.record.subject)
	fmt.Printf("考试分数:%v\n",i.record.score)
}