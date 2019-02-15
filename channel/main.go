package main

import "fmt"
import "time"

func init(){
	fmt.Println("This first statement ")
}
func main() {
	var a chan int
	if a == nil {
		fmt.Println("Channel is nil ready to define")
		a = make(chan int)
		fmt.Printf("Type of channel %T \n", a)
		fmt.Println("Channel ", a)

	}
	/*
		a <- data // write to channel a
		data := <- a // read from channel a
	*/

	fmt.Println("+++++++++++++++++++++++++++++++++++++++++")
	done := make(chan bool)
	go Hello(done)
	<-done
	fmt.Println("Main function ")

	number := 123456789
	for number != 0 {
		digit := number % 10
		fmt.Println(digit, number)
		number /= 10
	}
}

func Hello(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true // write to channel
}
