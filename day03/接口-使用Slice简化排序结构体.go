package main

import (
	"fmt"
	"sort"
)

type Car struct {
	Brand string
	Color string
}

func main() {
	cars := []*Car{
		{"Volkswagen", "black"},
		{"Toyota", "red"},
		{"Honda", "white"},
		{"LEXUS", "silver"},
		{"BMW", "king"},
		{"FORD", "blue"},
	}
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].Brand < cars[j].Brand
	})
	for _, v := range cars {
		fmt.Printf("%v\n", v)
	}
}