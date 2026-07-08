package main
import "fmt"

type runner  interface {
	run()
}

type person struct {
	legs int
}

type cat struct{
	legs int
}


func (p *person)run(){
	fmt.Printf("hunman use %v legs to run\n",p.legs)
}

func (c *cat)run(){
	fmt.Printf("Cat use %v legs to run\n",c.legs)
}


func main(){

	var r runner 
	pn := person{legs:2}
	ct := cat{legs:4}
	r = &pn
	value_pn,ok_pn := r.(*person)
	fmt.Printf("the value of r is %v,the type of person is %v",value_pn,ok_pn)
	r = &ct
	value_ct,ok_ct := r.(*person)
	fmt.Printf("the value of r is %v,the type of person is %v",value_ct,ok_ct)
}