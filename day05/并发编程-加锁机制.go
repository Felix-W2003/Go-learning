/* 
在Go语言提供提供sync包中有两种锁类型。即互斥锁（sync.Mutex)和
读写互斥锁(sync.RWMutex)

其中，互斥锁比较简单，当一个goroutine获得互斥锁后，其他goroutine只能
等这个goroutine释放互斥锁

读写互斥锁则是经典的单写多读模型。读锁被占用时仅阻止写。但不阻止读。在写锁
被占用的情况下，既阻止写，也阻止读
*/

package main
import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex //声明互斥锁的全局变量

func printData(){
	mutex.Lock()
	for i:=3;i<6;i++{
		time.Sleep(1*time.Second)
		fmt.Println(i)
	}
	mutex.Unlock()
}

func main(){

	go printData()
	
	mutex.Lock()
	for i:=0;i<3;i++{
		time.Sleep(1*time.Second)
		fmt.Println(i)
	}
	mutex.Unlock()
	time.Sleep(6*time.Second)

}