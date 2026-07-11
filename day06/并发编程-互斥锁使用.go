package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var numRand int       //声明表示随机数的全局变量
var rw sync.RWMutex   //声明读写互斥的全局变量
var wg sync.WaitGroup //声明同步等待组的全局变量

func read(i int) {
	rw.RLock()
	time.Sleep(time.Duration(i) * time.Second)
	fmt.Printf("执行读操作,数据:%d\n", numRand)
	rw.RUnlock()
	wg.Done()
}

func write(i int) {
	rw.Lock()
	numRand = rand.Intn(100)
	time.Sleep(time.Duration(i) * time.Second)
	fmt.Printf("执行写操作.数据：%d\n", numRand)
	rw.Unlock()
	wg.Done()
}

func main() {
	wg.Add(6)
	for i := 1; i < 4; i++ {
		go write(i)
	}
	for i := 1; i < 4; i++ {
		go read(i)
	}
	wg.Wait()
}

//在使用RLock()和RUnlock()、Lock()和Unlock()等函数
//它们必须成对出现
