package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("CPU Begin",runtime.NumCPU())
	fmt.Println("GOROUTINE Begin",runtime.NumGoroutine())
	wg.Add(2)
	go hear()
	go speak()
	fmt.Println("CPU middle",runtime.NumCPU())
	fmt.Println("GOROUTINE middle",runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("CPU end",runtime.NumCPU())
	fmt.Println("GOROUTINE end",runtime.NumGoroutine())
}

