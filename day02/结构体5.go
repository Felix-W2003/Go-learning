/*某些情况下，要使用指针类型的变量作为方法的接受者。一种是在
调用方法时需要修改变量值，还有一种是结构体成员较多，占用内存较大

*/

package main
import( "fmt" 
"strconv")

type profile struct {
	name string 
	age int
	interest[] string
	evaluate string

}

//定义指针接收者方法
func (p *profile)get_info(subject string,score int) string{
	info := "姓名:"+p.name +"\n年龄："+strconv.Itoa(p.age)
	var hobby string
	for _,v:=range p.interest{
		hobby +=v+""
	}
	info +="\n兴趣爱好:"+hobby+"\n自我评价："+p.evaluate
	info +="\n考试科目："+subject+"\n考试分数："+strconv.Itoa(score)
	return info
}



func main (){
	p := &profile{
		name : "张三",
		age : 20,
		interest: []string{"看电影","唱歌","打飞机"},
		evaluate:"圣人一个",
	}

	result := p.get_info("综合知识",76)
	fmt.Printf(result)
}
