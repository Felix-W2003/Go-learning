package main
import ("fmt"
		"strconv")

type goods struct{
	name 	[]string
	price 	[]int
	num		[]int
}




func (g *goods)show()string{
	var info string
	for k,v :=range g.name{
	fmt.Printf("k:%v\n",k)
	fmt.Printf("v:%v\n",v)
	info += "商品名称:"+v
	info += "商品单价" + strconv.Itoa(g.price[k])+"元"
	info += "商品数量" + strconv.Itoa(g.num[k])+"\n"
	}
	return info
}





//定义结构体方法，计算商品总价
func (g *goods)count()int{
	var total int 
	for k,v := range g.price{
		total += v*g.num[k]
	}
	return total
}

func (g *goods)init(name []string,price []int, num []int){
	//初始化结构体成员
	g.name = name 
	g.price = price 
	g.num = num
}

func main(){
	g := &goods{}
	g.init([]string{"品牌手机","品牌计算机"},[]int{1699,2399},[]int{2,3})
	info := g.show()
	result := g.count()
	fmt.Println("-------------------------------")
	fmt.Print(info)
	fmt.Printf("商品总价：%v元",result)
}