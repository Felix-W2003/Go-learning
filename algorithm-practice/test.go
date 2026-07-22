package main

import "fmt"

type Student struct {
	name  string
	score int64
}

func main() {
	var a = []int{}
	stu := Student{
		name:  "张三",
		score: 60,
	}
	fmt.Println(len(a))
	// a = append(a, 1, 2, 3, 4)
	// fmt.Println(a)
	// fmt.Println(a[0])
	stus := []Student{}
	fmt.Println(stu)
	fmt.Println(stus)
	stus = append(stus, Student{
		name:  "lisi",
		score: 99,
	}, stu)
	stus2 := []Student{{
		name:  "王五",
		score: 90,
	}}
	stus = append(stus, stus2...)
	fmt.Println(stus)
}
