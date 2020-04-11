package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var total struct{
	sync.Mutex
	value int
}

type Task struct {
	Id  int
	Sum int
}

func NewTask(id ,s int)*Task{
	return &Task{Sum:s,Id:id}
}
func(t *Task) worker(wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Printf("任务%d开始执行,内容是1-%d的和\n",t.Id,t.Sum)
	for i:=1;i<=t.Sum;i++{
		total.Lock()
		total.value+=i
		total.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	rand.Seed(time.Now().Unix())
	for i:=0;i<2;i++{
		wg.Add(1)
		t := NewTask(i+1,rand.Intn(100))
		go t.worker(&wg)
	}
	wg.Wait()
	fmt.Println(total.value)
}


