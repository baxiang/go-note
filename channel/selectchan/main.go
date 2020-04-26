package main

import (
	"fmt"
	"math/rand"
)

func main() {
	intChannels := [3]chan int {
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}
	// 随机选择一个通道，并向它发送元素子
	index := rand.Intn(3)
	fmt.Printf("The index:%d \n",index)
	intChannels[index] <-index
	// 哪一个通道中有可取的元素值，哪个对应的分支就被执行。
	select {
	case<-intChannels[0]:
		fmt.Println("The first candidate case is selected.")
	case<-intChannels[1]:
		fmt.Println("The second candidate case is selected.")
	case<-intChannels[2]:
		fmt.Println("The third candidate case is selected.")
	default:
		//fmt.Println("No candidate case is selected")
	}
}
