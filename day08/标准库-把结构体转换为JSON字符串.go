package main
import (
	"fmt"
	"encoding/json"
)

type Person struct{
	Name string
	Age  int 
	Sex  string
}

func Marshal(){
	p := Person{
		Name : "David",
		Age : 26,
		Sex : "Male",
	}

	b,_ := json.Marshal(p)
	fmt.Printf("b:%v",string(b))

	fmt.Printf("b:%T",b)
}

func main(){
	Marshal()
}