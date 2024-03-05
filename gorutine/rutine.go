package rutine

import (
	"fmt"
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
