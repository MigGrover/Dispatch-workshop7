package patternpc

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func createProducer(wg *sync.WaitGroup) <-chan int {
	c := make(chan int)
	wg.Add(1)
	go func() {
		defer close(c)
		defer wg.Done()
		for i := 0; i < 5; i++ {
			duration := rand.Intn(1e3)
			time.Sleep(time.Duration(duration) * time.Millisecond)
			c <- i
		}
	}()
	return c
}

func createConsumer(wg *sync.WaitGroup, c <-chan int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for value := range c {
			fmt.Println(value)
		}
	}()
}

func ExecutePCPattern() {
	wg := &sync.WaitGroup{}
	c := createProducer(wg)
	createConsumer(wg, c)
	wg.Wait()
}
