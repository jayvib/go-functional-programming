package chanfunc

import (
	"math/rand"
	"time"
)

func IntGenerator(size int) <-chan int {
	rand.Seed(time.Now().Unix())
	intChan := make(chan int, size)
	go func() {
		defer close(intChan)
		for i := 0; i < size; i++ {
			intChan <- rand.Intn(1000)
		}
	}()
	return intChan
}
