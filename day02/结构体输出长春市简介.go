package main
import ("fmt"
"encoding/json")

type spots struct {
// 定义结构体，成员为切片类型，切片元素同样为结构体
	Spot []struct {
		Name string
		Level string
		Ticket int
	}
}

type city struct {
	name string 
	code string
	tourism spots
}

func main() {
	var s spots//实例化spots结构体
	//实例化city结构体
	c:=city{name:"长春市",code:"0431",tourism:s}
	//JSON格式数据
	j:=`{"spot":[
				{"name":"长影世纪城","level":"AAAAA","ticket":120},
				{"name":"伪满皇宫博物院","level":"AAAA","ticket":60},
				{"name":"长春世界雕塑公园","level":"AAAA","ticket":0}]
				}`
	json.Unmarshal([]byte(j),&s)//JSON转换为结构体
	fmt.Printf("城市: %v\n", c.name)
	fmt.Printf("电话区号: %v\n", c.code)
	fmt.Println("旅游景点信息如下：")
	fmt.Println("====================================")
	for _,v := range s.Spot{
		fmt.Printf("景点名称: %v\n", v.Name)
		fmt.Printf("景点等级: %v\n", v.Level)
		fmt.Printf("景点门票: %v\n", v.Ticket)
	}
}
