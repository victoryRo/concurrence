package canal

import (
	"fmt"
	"sync"
	"time"
)

// -------------------------------------------------------------------- sequential

func timer(duration []int) {
	for _, t := range duration {
		executeTime(t)
	}
}

func executeTime(undTime int) {
	//fmt.Printf("time duration %v\n", undTime)
	for {
		if undTime <= 0 {
			break
		} else {
			time.Sleep(1 * time.Second)
			undTime--
		}
	}
}

//-------------------------------------------------------------------- concurrent

func timerConcurrent(duration []int) {
	var wg sync.WaitGroup
	wg.Add(len(duration))

	for _, t := range duration {
		tt := t
		go func(tm int) {
			executeTime(tm)
			wg.Done()
		}(tt)
	}
	wg.Wait()
}

//-------------------------------------------------------------------- channels

func timerChannelCSP(duration []int) {
	signal := make(chan struct{})

	for _, t := range duration {
		go func(tm int) {
			executeTime(tm)
			signal <- struct{}{}
		}(t)
	}

	for range duration {
		<-signal
		fmt.Println("signal")
	}
}

//-------------------------------------------------------------------- concurrent cancellation

func timerChannelCancellation(duration []int) {
	done := make(chan struct{})

	for _, t := range duration {
		go func(tm int) {
			executeTime(tm)
			fmt.Printf("Execute %d secods\n", tm)
			select {
			case <-done:
				return
			}
		}(t)
	}

	select {
	case <-time.After(2 * time.Second):
		close(done)
	}
}

// ChannelPkg syncro
func ChannelPkg() {
	times := []int{3, 7, 2, 1, 2, 5}

	//timer(times)
	//timerConcurrent(times)
	//timerChannelCSP(times)
	timerChannelCancellation(times)
}
