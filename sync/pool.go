package main

import (
	"fmt"
	"sync"
)

func main() {
	sp := sync.Pool{
		//创建一个Pool，并且实现New()函数
		New: func() interface{}{
			return "hello world"
		},
	}
	// 第1次获取池中值
	v1 := sp.Get()
	fmt.Println("第1次值：", v1)
	//New()返回的是interface{}通过类型断言来转换
	if v1,ok := v1.(string);ok {
		v1= "hello golang"
		sp.Put(v1) // 获取池子中元素 然后修改值
	}
	// 第2次获取池中值
	v2 := sp.Get()
	fmt.Println("第2次值：", v2)
	// 第3次获取池中值
	v3 := sp.Get()
	//因为池中的对象已经没有了，所以又重新通过New()创建一个新对象，放入池中，然后返回
	fmt.Println("第3次值：", v3)
}
