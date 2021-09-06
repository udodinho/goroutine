package main


func main() {
	ch := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		done <- true
	}()

	go func() {
		<- done
		<- done
		close(ch)
	}()

	for n := range ch {

	}
}
