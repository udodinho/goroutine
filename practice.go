package main

import (
	"fmt"
	"sync"
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


//type Ball struct {
//	hits int
//}
//func main() {
//	table := make(chan *Ball)
//	go Player("Ping", table)
//	go Player("Pong", table)
//
//	table <- new(Ball)
//	time.Sleep(1 * time.Second)
//	<- table
//}
//
//func Player(name string, table chan *Ball) {
//	for {
//		ball := <- table
//		ball.hits++
//		fmt.Println(name, ball.hits)
//		time.Sleep(100 * time.Millisecond)
//		table <- ball
//	}
//}

// Channels


//func waitForTask() {
//	ch := make(chan string)
//
//	go func() {
//	time.Sleep(1 * time.Second)
//	ch <- "paper"
//	}()
//	p := <- ch
//	fmt.Println(p)
//}
//
//func fanOut() {
//	emps := 20
//	ch := make(chan string, emps)
//
//	for e := 0; e < emps; e++ {
//		go func() {
//			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
//			ch <- "paper"
//		}()
//
//		for emps > 0 {
//			p := <-ch
//			fmt.Println(p)
//			emps--
//		}
//	}
//}

//func main() {
//	ch := make(chan string)
//	go sendData(ch)
//	go getData(ch)
//	time.Sleep(1e9)
//
//}
//
//func sendData(ch chan string) {
//	ch <- "Washington"
//	ch <- "London"
//	ch <- "Tripoli"
//	ch <- "Beijing"
//	ch <- "Tokyo"
//}
//
//func getData(ch chan string) {
//var input string
//for {
//	input= <- ch
//	fmt.Printf("%s\n", input)
//	}
//	close(ch)
//}

//func main() {
//	ch1 := make(chan int)
//	go pump(ch1)
//	go suck(ch1)
//	time.Sleep(1e9)
//}
//
//func pump(ch chan int) {
//	for i := 0; ; i++ {
//		ch <- i
//	}
//}
//
//func suck(ch chan  int) {
//	for {
//		fmt.Println(<- ch)
//	}
//}

// Using range and close  method in channels


var wg sync.WaitGroup

func main() {
	ch := make(chan int)

	wg.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			
		}
	}()

}