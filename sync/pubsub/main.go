package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type (
	subscriber chan interface{}
	topicFunc func(v interface{}) bool
)

type Publisher struct {
	m sync.RWMutex
	buffer int
	timeout time.Duration
	subscribers map[subscriber] topicFunc
}
func NewPublisher(timeout time.Duration,buffer int)*Publisher{
	return &Publisher{
		buffer:      buffer,
		timeout:     timeout,
		subscribers: make(map[subscriber]topicFunc),
	}
}
func (p *Publisher)Publish(v interface{})  {
	p.m.RLock()
	defer p.m.RUnlock()
	var wg sync.WaitGroup

	for sub,topic :=range p.subscribers{
		wg.Add(1)
		go p.sendTopic(sub,topic,v,&wg)
	}
	wg.Wait()
}

func(p *Publisher)sendTopic(
	sub chan interface{}, topic topicFunc, v interface{}, wg *sync.WaitGroup,
	){
	defer wg.Done()
	if topic!=nil&&!topic(v){
		return
	}

	select {
	case sub<-v:
	case <-time.After(p.timeout):
	}
}

func(p *Publisher)SubscribeTopic(topic topicFunc)chan interface{}{
	ch :=make(chan interface{},p.buffer)
	p.m.Lock()
	p.subscribers[ch] = topic
	p.m.Unlock()
	return ch
}

func(p *Publisher)SubAll() chan interface{}{
	 return p.SubscribeTopic(nil)
}

func(p *Publisher)Close(){
	p.m.Lock()
	defer p.m.Unlock()
	for sub :=range p.subscribers{
		delete(p.subscribers,sub)
		close(sub)
	}
}

func main() {
	p := NewPublisher(100*time.Millisecond,10)
	defer p.Close()

	all :=p.SubAll()
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s,ok:=v.(string);ok {
			return strings.Contains(s,"golang")
		}
		return false
	})

	p.Publish("Hello world!")
	p.Publish("Hello golang!")
	go func() {
		for msg :=range all{
			fmt.Println("all",msg)
		}
	}()

	go func() {
		for msg :=range golang{
			fmt.Println("golang:",msg)
		}
	}()
	select {
	case <-time.After(3*time.Second):
	}
}
