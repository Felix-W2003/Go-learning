package main
import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
	Sex  string
}


func Unmarshal(){
	b1 := []byte(`{"Name":"David","Age":26,"Sex":"Male"}`)
	var m Person
	json.Unmarshal(b1,&m)
	fmt.Printf("m:%v\n",m)
}


func main(){
	Unmarshal()
}