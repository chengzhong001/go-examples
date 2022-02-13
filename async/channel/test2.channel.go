package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)
	fmt.Println("len(c) =", len(c), ", cap(c) =", cap(c))
	go func() {
		defer fmt.Println("sub go is running")
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("sub go is running: i value =", i,  "len(c) =", len(c), ", cap(c) =", cap(c))
		}
	}()
	time.Sleep(time.Second*1)
	for i := 0; i < 4; i++ {
		num := <- c;
		fmt.Println(num)
	}
	fmt.Println("main finish")
}
