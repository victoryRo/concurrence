package mutex

import (
	"fmt"
	"sync"
	"time"
)

type account struct {
	name    string
	balance int
}

func transfer(amount int, source, dest *account) {
	if source.balance < amount {
		fmt.Printf("❌: %s\n", fmt.Sprintf("%v %v", source, dest))
		return
	}
	time.Sleep(time.Second)

	dest.balance += amount
	source.balance -= amount
	fmt.Printf("✅: %s\n", fmt.Sprintf("%v %v", source, dest))
}

func CheckMutex() {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	wg.Add(2)

	juan := account{"Juan", 900}
	pedro := account{"Pedro", 1200}

	for _, value := range []int{500, 500} {
		go func(quantity int) {
			mu.Lock()
			transfer(quantity, &juan, &pedro)
			mu.Unlock()
			wg.Done()
		}(value)
	}

	wg.Wait()
}
