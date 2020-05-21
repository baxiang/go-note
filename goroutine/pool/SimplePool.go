package main

import (
	"fmt"
	"sync"
	"time"
)

type SimplePool struct {
	wg sync.WaitGroup
	work chan func()
}

func NewSimplePoll(works int)*SimplePool{
	p := &SimplePool{
		wg:   sync.WaitGroup{},
		work: make(chan func()),
	}
	p.wg.Add(works)
	for i:=0;i<works;i++{
		go func() {
			defer func() {
				if err:=recover();err!=nil{
					fmt.Println(err)
					p.wg.Done()
				}
				for fn :=range p.work{
					fn()
				}
				p.wg.Done()
			}()
		}()
	}
	return p
}

func(p *SimplePool)Add(fn func()){
	p.work<-fn
}

func(p *SimplePool)Run(){
	close(p.work)
	p.wg.Wait()
}

func main() {
	p := NewSimplePoll(20)
	for i:=0;i<100;i++{
		p.Add(taskExec(i))
	}
	p.Run()
}
func taskExec(i int)func()  {
	return func() {
		time.Sleep(time.Second*1)
		fmt.Println("exec finish",i)
	}
}