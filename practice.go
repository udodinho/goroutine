package main

import (
	"fmt"
	"time"
)

//func main() {
//	runtime.GOMAXPROCS(2)
//
//	var wg sync.WaitGroup
//	wg.Add(2)
//
//	fmt.Println("Starting Go Routines")
//	go func() {
//		defer wg.Done()
//
//		for char := 'a'; char < 'a'+26; char++ {
//			fmt.Printf("%c ", char)
//		}
//}()
//
//	go func() {
//				defer  wg.Done()
//				time.Sleep(1 * time.Microsecond)
//				for number := 1; number < 27; number++ {
//					fmt.Printf("%d ", number)
//				}
//
//			}()
//	fmt.Println("Waiting To Finish")
//		wg.Wait()
//		fmt.Println("\nTerminating Program")
//
//}


//var Wait sync.WaitGroup
//var Counter  = 0
//var Lock sync.Mutex
//
//func main() {
//	for routine := 1; routine <= 2; routine++ {
//		Wait.Add(1)
//		go Routine(routine)
//
//	}
//	Wait.Wait()
//	fmt.Printf("Final counter: %d\n", Counter)
//
//}
//
//func Routine(id int) {
//	for count := 0; count < 2; count++ {
//		Lock.Lock()
//		value := Counter
//		time.Sleep(1 * time.Nanosecond)
//		value++
//		Counter = value
//
//		Lock.Unlock()
//	}
//	Wait.Done()
//}


type Ball struct {
	hits int
}
func main() {
	table := make(chan *Ball)
	go Player("Ping", table)
	go Player("Pong", table)

	table <- new(Ball)
	time.Sleep(1 * time.Second)
	<- table
}

func Player(name string, table chan *Ball) {
	for {
		ball := <- table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}