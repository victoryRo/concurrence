package canal

import (
	"fmt"
)

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
func send(num chan<- int) {
	num <- 33
}

func receive(num <-chan int) {
	fmt.Println(<-num)
}

// ----------------------------------------------------

func chanWithBuffer() {
	quantity := make(chan int, 2)
	signal := make(chan struct{})

	go receiver(signal, quantity)
	sender(quantity)

	<-signal
}

func sender(num chan<- int) {
	num <- 71
	num <- 21
	num <- 11
}

func receiver(signal chan<- struct{}, num <-chan int) {
	fmt.Println(<-num)
	fmt.Println(<-num)
	fmt.Println(<-num)

	signal <- struct{}{}
}

//----------------------------------------------------

// chanWithBuffer2 patterns for close()
func chanWithBuffer2() {
	quantity := make(chan int, 2)
	signal := make(chan struct{})

	go receiver2(signal, quantity)
	sender2(quantity)

	<-signal
}

func sender2(num chan<- int) {
	num <- 1
	num <- 2
	num <- 3
	num <- 4
	num <- 5
	num <- 6

	close(num)
}

func receiver2(signal chan<- struct{}, num <-chan int) {
	for v := range num {
		fmt.Println(v)
	}

	signal <- struct{}{}
}

//----------------------------------------------------

// chanWithBuffer3 patterns for select{}
func chanWithBuffer3() {
	number := make(chan int)
	signal := make(chan struct{})

	go receiver3(signal, number)
	sender3(number)

	signal <- struct{}{}
}

func sender3(number chan<- int) {
	number <- 1
	number <- 2
	number <- 3
	number <- 4
	number <- 5
	number <- 6
}

func receiver3(signal <-chan struct{}, number <-chan int) {
	for {
		select {
		case v := <-number:
			fmt.Println(v)
		case <-signal:
			return
		default:
			fmt.Println("🤔")
		}
	}
}

//----------------------------------------------------

func SendChannels() {
	//example1()
	//unidirectionalChannel()
	// chanWithBuffer()
	//chanWithBuffer2()
	chanWithBuffer3()
}
