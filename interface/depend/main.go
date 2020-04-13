package main

import "fmt"

//------  抽象层 -----
type Card interface{
	Display()
}

type Memory interface {
	Storage()
}

type CPU interface {
	Calculate()
}

type Computer struct {
	cpu CPU
	mem Memory
	card Card
}

func NewComputer(cpu CPU, mem Memory, card Card) *Computer{
	return &Computer{
		cpu:cpu,
		mem:mem,
		card:card,
	}
}

func (this *Computer) DoWork() {
	this.cpu.Calculate()
	this.mem.Storage()
	this.card.Display()
}

//------  实现层 -----
//intel
type IntelCPU struct {}

func (i *IntelCPU) Calculate() {
	fmt.Println("Intel CPU 开始计算了...")
}

type AMDCPU struct {}
func (a *AMDCPU) Calculate() {
	fmt.Println("AMD CPU 开始计算了...")
}
type IntelMemory struct {}

func (i *IntelMemory) Storage() {
	fmt.Println("Intel Memory 开始存储了...")
}

type IntelCard struct {}

func (i *IntelCard) Display() {
	fmt.Println("Intel Card 开始显示了...")
}

//kingston
type KingstonMemory struct {}

func (k *KingstonMemory) Storage() {
	fmt.Println("Kingston memory storage...")
}

//nvidia
type NvidiaCard struct {}

func (k *NvidiaCard) Display() {
	fmt.Println("Nvidia card display...")
}



//------  业务逻辑层 -----
func main() {
	//intel系列的电脑
	com1 := NewComputer(&IntelCPU{}, &IntelMemory{}, &IntelCard{})
	com1.DoWork()

	//组合电脑
	com2 := NewComputer(&AMDCPU{}, &KingstonMemory{}, &NvidiaCard{})
	com2.DoWork()
}
