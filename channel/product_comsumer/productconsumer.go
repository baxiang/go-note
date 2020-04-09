package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Product struct {
	Id string
}

type Producer struct {
	Id int
}

func (p *Producer) work(b chan <-Product){
	for {
		product := Product{ Id: fmt.Sprintf("%d",time.Now().Unix()) }
		b <- product
		fmt.Printf("生产者-%d 生产了产品-%s\n", p.Id, product.Id)
		// 休息1秒
		randTime := rand.Intn(10)
		time.Sleep(time.Duration(randTime)*time.Second)
	}
}

type Consumer struct {
	Id int
}

func (c *Consumer)consumer(b <-chan Product){
	for{
		product:=<-b
		fmt.Printf("消费者-%d 消费了产品-%s\n", c.Id, product.Id)
	}
}

func main() {
	var buffer = make(chan Product,10) //队列
	//创建Producer
	for i := 1; i <= 2; i++ {
		p := &Producer{ Id: i }
		go p.work(buffer)
	}
   // 创建Consumer
	for i := 1; i <= 5; i++ {
		c := &Consumer{ Id: i }
		go c.consumer(buffer)
	}
	wait := make(chan bool)
	<-wait
}
