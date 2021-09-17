package main

import "fmt"

//var wg sync.WaitGroup
//
//func main() {
//	fmt.Println("CPU Begin",runtime.NumCPU())
//	fmt.Println("GOROUTINE Begin",runtime.NumGoroutine())
//	wg.Add(2)
//	go hear()
//	go speak()
//	fmt.Println("CPU middle",runtime.NumCPU())
//	fmt.Println("GOROUTINE middle",runtime.NumGoroutine())
//	wg.Wait()
//	fmt.Println("CPU end",runtime.NumCPU())
//	fmt.Println("GOROUTINE end",runtime.NumGoroutine())
//}
//
//func speak()  {
//	fmt.Println("Hello you are there?")
//	wg.Done()
//}
//
//func hear()  {
//	fmt.Println("Yes am here!")
//	wg.Done()
//}


type Human interface {
	Speak()
}

type Person struct {
	name string
}

func (p *Person) Speak()  {
	fmt.Println("Hello are you there?")
}

func saySomething(h Human) {
	h.Speak()
}

func main() {
	p1 := Person{
		name: "Goodness",
	}
	p1.Speak()
	saySomething(&p1)
}