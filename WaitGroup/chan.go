package main

import "fmt"

func chanPrint() {
	c := make(chan int, 100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			c <- i
		}(i)
	}
	for i := 0; i < 100; i++ {
		<-c
	}
}
