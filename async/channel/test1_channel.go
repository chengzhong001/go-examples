package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		defer fmt.Println("")
		fmt.Println("goroutine 正在运行...")
		c <- 666
	}()

	if num, ok := <- c; ok{
		fmt.Printf("num = %d\n",num)
	}

}
