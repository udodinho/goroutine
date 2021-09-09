package main

import "fmt"

// Using Semaphore method
//func main() {
//	ch := make(chan int)
//	done := make(chan bool)
//
//	go func() {
//		for i := 0; i < 10; i++ {
//			ch <- i
//		}
//		done <- true
//	}()
//
//	go func() {
//		for i := 0; i < 10; i++ {
//			ch <- i
//		}
//		done <- true
//	}()
//
//	go func() {
//		<- done
//		<- done
//		close(ch)
//	}()
//
//	for n := range ch {
//		fmt.Println(n)
//	}
//}

func main() {
	ch := make(chan int)
	done := make(chan bool)
	n := 100
	go func() {
		for i := 0; i < n; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		for n := range ch {
			fmt.Println(n)
		}
		done <- true
	}()
	go func() {
		for n := range ch {
			fmt.Println(n)
		}
		done <- true
	}()

	<- done
	<- done
}
