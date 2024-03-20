package syncro

import (
	"fmt"
	"sync"
	"time"
)

// -------------------------------------------------------------------- secuential

func timer(duration []int) {
	for _, t := range duration {
		executeTime(t)
	}
}

func executeTime(undTime int) {
	fmt.Printf("time duration %v\n", undTime)
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
		go func() {
			executeTime(tt)
			wg.Done()
		}()
	}
	wg.Wait()
}

// ExecPkg syncro
func ExecPkg() {
	times := []int{1, 2, 5}

	// timer(times)
	timerConcurrent(times)
}
