package main

import "sync"

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


var Wait sync.WaitGroup
var Counter  = 0
var Lock sync.Mutex

func main() {
	for routine := 1; routine <= 2; routine++ {
		Wait.Add(1)
		go Routine(routine)

	}