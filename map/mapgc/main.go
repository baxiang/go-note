package main

import (
	"log"
	"runtime"
)

func printMemStats(mag string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("%v：分配的内存 = %vKB, GC的次数 = %v\n", mag, m.Alloc/1024, m.NumGC)
}

// go version go1.12.5 linux/amd64 不需要
func main() {
	m := make(map[int]int)
	for i := 0; i < 10000; i++ {
		m[i] = i
	}
	runtime.GC()
	printMemStats("add map")
	for i := 0; i < 10000; i++ {
		delete(m, i)
	}
	runtime.GC()
	printMemStats("clear map")
	m = nil
	runtime.GC()
	printMemStats("set nil")

}
