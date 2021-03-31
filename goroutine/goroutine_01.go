package goroutine

import (
	"fmt"
	"time"
)

func calculate(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

func GoRoutine() {
	fmt.Println("Concurrency in golang")
	go calculate(5)
	go calculate(5)
	fmt.Scanln()
}
