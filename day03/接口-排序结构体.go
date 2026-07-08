package main
import ("fmt"
		"sort")


type Car struct {
	Brand string
	Color string
}

type Cars []*Car //把Car指针类型的切片声明为Cars类型

func (c Cars)Less(i,j int) bool {
	return c[i].Brand < c[j].Brand
}

func (c Cars)Len()int {
	return len(c)
}

func (c Cars) Swap(i,j int){
	c[i],c[j]=c[j],c[i]
}
func main() {
	cars := Cars{
		&Car{"Volkswagen", "black"},
		&Car{"Toyota", "red"},
		&Car{"Honda", "white"},
		&Car{"LEXUS", "silver"},
		&Car{"BMW", "king"},
		&Car{"FORD", "blue"},
	}
	sort.Sort(cars)
	for _, v := range cars {
		fmt.Printf("%v\n", v)
	}
}