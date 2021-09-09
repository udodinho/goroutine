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
