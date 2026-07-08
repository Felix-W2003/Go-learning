package main
import "fmt"

type carInMotion interface{
	move(speed int)		//为行驶中的汽车声明接口
	brake()(int,int)	//刹车函数。返回值分别是刹车前的速度和刹车后的速度
	park()				//泊车函数
	//消耗燃油函数。参数分别为剩余燃油量和燃油平均消耗，返回值是汽车
	//还能行驶的距离
	consumeOil(fuelLeft float64, aver_consumption float64)(distance float64)

}

type car struct{	//表示汽车的结构体
	color string 	//车身颜色
}


func (c *car)move(speed int){
	fmt.Printf("一辆%v的汽车在以%vkm/h的速度匀速行驶\n",c.color,speed)
}


func (c *car)brake()(int,int){
	fmt.Printf("这辆%v的汽车开始刹车\n",c.color)
	speedBeforeBrake := 60
	speedAfterBrake := 0
	return speedBeforeBrake,speedAfterBrake
}

func (c *car)park(){
	fmt.Printf("这辆%v的汽车停在路边的车位里\n",c.color)

}

//实现接口中的consumeOil()函数
func  (c *car)consumeOil(fuelLeft float64,aver_consumption float64)(distance float64){
	fmt.Printf("这辆%v的汽车剩余燃油%vL,燃油平均消耗%vL/100km\n",c.color,fuelLeft,aver_consumption)
	return fuelLeft/aver_consumption *100
}

func main(){

	var cim carInMotion
	cr := car{color:"红色"}
	cim = &cr
	cim.move(60)
	speedBeforeBrake,speedAfterBrake := cim.brake()
	fmt.Printf("刹车前的速度是%vkm/h,刹车后的速度是%vkm/h\n",speedBeforeBrake,speedAfterBrake)
	cim.park()
	distance := cim.consumeOil(27,6.3)
	fmt.Printf("还能继续行驶%.2fkm\n",distance)
}