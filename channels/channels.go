package main

import (
	"fmt"
)

func main() {
	// ch := make(chan int)
	// go func(a int, b int) {
	// 	c := a + b
	// 	ch <- c
	// }(1, 2)
	// r := <-ch
	// fmt.Println(r)
	// declaration of the channel
	//var ch2 chan int
	ch2 := make(chan int)

	//for send the values in the channels, we need to use the <- arrow on the right side of the channel varibale
	// ch2 <- go multiply(5,5)
	go func() {
		ch2 <- multiply(5, 5)
	}()
	//so we cant add direct result into the channel, we can only add by using the go routine

	//if we have send a value there must be reciver also for that channel or else it will give an error

	channelReciver := <-ch2
	fmt.Println(channelReciver)

}
func multiply(a int, b int) int {
	c := a * b
	return c
}
