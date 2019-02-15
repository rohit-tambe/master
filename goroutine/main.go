package main

import "fmt"
import "time"
import "runtime"

func Hello() {
	fmt.Println("Hello function Goroutine")
}

func main() {
	fmt.Println("OS", runtime.GOOS)
	fmt.Println("Architechure ", runtime.GOARCH)
	fmt.Println("CPU ", runtime.NumCPU())
	fmt.Println("Total GoRutines ", runtime.NumGoroutine())
	go Hello()
	fmt.Println("Total GoRutines After ", runtime.NumGoroutine())
	fmt.Println("CPU After", runtime.NumCPU())
	time.Sleep(1 * time.Second)
	fmt.Println("main function")
}
