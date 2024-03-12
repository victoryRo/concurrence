package canal

import "fmt"

func example1() {
	msg := make(chan string)

	// child gorutine
	go func() {
		msg <- "Thanks..."
	}()

	//listen gorutine main
	fmt.Println(<-msg)
}

// ----------------------------------------------------

func unidirectionalChannel() {
	quantity := make(chan int)

	go receive(quantity)
	send(quantity)
}

func receive(num <-chan int) {
	fmt.Println(<-num)
}

func send(num chan<- int) {
	num <- 33
}

// ----------------------------------------------------

func chanWithBuffer() {
	quantity := make(chan int, 2)
	signal := make(chan struct{})

	go receiver(signal, quantity)
	sender(quantity)

	<-signal
}

func receiver(signal chan<- struct{}, num <-chan int) {
	fmt.Println(<-num)
	fmt.Println(<-num)
	fmt.Println(<-num)

	signal <- struct{}{}
}

func sender(num chan<- int) {
	num <- 71
	num <- 21
	num <- 11
}

//----------------------------------------------------

func SendChannels() {
	//example1()
	//unidirectionalChannel()
	chanWithBuffer()
}
