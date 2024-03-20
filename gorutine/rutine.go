package rutine

import (
	"fmt"
	"sync"
	"time"
)

func hello() int {
	fmt.Println("Hello i am beging with gorutines")
	return 1
}

func SendAction() {
	//go hello()
	go func() {
		fmt.Println("I am an anonymous function")
	}()

	time.Sleep(time.Millisecond)
	fmt.Println("Hi Gophers")
}

// DataRace data race solution with WaitGroup and Mutex
func DataRace() {
	wg := sync.WaitGroup{}
	mx := sync.Mutex{}
	wg.Add(1)

	data := 1

	go func() {
		mx.Lock()
		data++
		mx.Unlock()
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(data)
}

func DataRaceInLoop() {
	wg := sync.WaitGroup{}
	wg.Add(5)

	for i := 0; i < 5; i++ {
		j := i
		go func() {
			fmt.Println(j)
			wg.Done()
		}()
	}

	wg.Wait()
}

func DataRaceMap() {
	var wg sync.WaitGroup
	var mx sync.Mutex
	wg.Add(2)

	courses := make(map[string]string)

	go func() {
		mx.Lock()
		courses["go from scratch"] = "Intermedio"
		mx.Unlock()
		wg.Done()
	}()

	go func() {
		mx.Lock()
		courses["go concurrent"] = "Avanzado"
		mx.Unlock()
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(courses)
}
