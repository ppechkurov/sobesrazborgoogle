package channels

import (
	"fmt"
	"time"
)

func worker(done chan<- struct{}) {
	fmt.Println("woring...")
	time.Sleep(time.Second)
	fmt.Println("done!")

	done <- struct{}{}
}

func Run() error {
	done := make(chan struct{}, 1)
	defer close(done)

	go worker(done)

	<-done

	return nil
}
