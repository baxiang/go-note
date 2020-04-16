package main
func allocateMemory(){
	_= make([]byte,1<<20)
}

func main() {
	for i:=1;i<10000;i++{
		allocateMemory()
	}
}
