package main

import (
	"fmt"
	"sync"
)

//WaitGroup不能是负数
//如果需要引用传递则要传递指针
func waitgroupPrint() {
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
