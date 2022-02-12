package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()

	go func(a int, b int){
		fmt.Println(a, b)
	}(1,2)
	time.Sleep(time.Second * 1)
	//
}
