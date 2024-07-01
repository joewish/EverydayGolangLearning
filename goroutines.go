package main

import (
	"fmt"
	"time"
)

func main() {
	go printGoRoutine("first")
	fmt.Println("hello world")
	printGoRoutine("second")
	time.Sleep(time.Second)
}

func printGoRoutine(number string) {
	fmt.Println("this go routine " + number)
}
