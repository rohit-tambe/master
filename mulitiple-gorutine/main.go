package main

import (
	"fmt"
	"runtime"
	"sync"
	// "time"
)
var wg sync.WaitGroup
func numbers() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
		// time.Sleep(100 * time.Millisecond)
	}
	// wg.Done()
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		fmt.Printf(" %c ", i)
		// time.Sleep(100 * time.Millisecond)
	}
	wg.Done()
}
func main() {
	fmt.Println("OS", runtime.GOOS)
	fmt.Println("Architechure ", runtime.GOARCH)
	fmt.Println("CPU ", runtime.NumCPU())
	fmt.Println("Total GoRutines ", runtime.NumGoroutine())
	wg.Add(1)
	go numbers()
	 alphabets()
	fmt.Println("Total GoRutines After ", runtime.NumGoroutine())
	fmt.Println("CPU After", runtime.NumCPU())
	// time.Sleep(30000 * time.Microsecond)
	fmt.Println("main terminated")
	wg.Wait()
}
