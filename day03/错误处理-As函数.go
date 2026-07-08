package main

import (
	"errors"
	"fmt"
)

type TestError struct {
	err string
}

func (t TestError) Error() string {
	return t.err
}

func main() {
	err := TestError{"num 不是数值类型"}
	err1 := fmt.Errorf("数据类型错误: %w", err)
	var test TestError
	if errors.As(err1, &test) {
		fmt.Println("错误信息:", test.err)
	}
}