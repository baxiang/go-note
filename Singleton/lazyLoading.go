package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"unsafe"
)

type singleton struct {

}
var instance *singleton
var initialized uint32
var lock sync.Mutex

func GetInstance()*singleton{
	if atomic.LoadUint32(&initialized)==1 {
		return instance
	}
	lock.Lock()
	defer lock.Unlock()
	if initialized == 0{
		instance = &singleton{}
		atomic.StoreUint32(&initialized,1)
	}
	return instance
}

func main() {
	a := GetInstance()
	b := GetInstance()
	fmt.Println(unsafe.Pointer(a))
	fmt.Println(unsafe.Pointer(b))
}
