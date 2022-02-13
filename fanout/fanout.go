package fanout

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func createGenerator(wg *sync.WaitGroup) <-chan string {
	c := make(chan string)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(c)
		for i := 0; i < 100; i++ {
			c <- strconv.Itoa(i)
		}
	}()

	return c
}

func fanOut(wg *sync.WaitGroup, input <-chan string, numJobs int) {
	for i := 0; i < numJobs; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for message := range input {
				fmt.Printf("Worker %v, got message: %s\n", id, message)
				duration := rand.Intn(1e3)
				time.Sleep(time.Duration(duration) * time.Millisecond)
			}
		}(i)
	}
}

func ExecuteFanOut() {
	wg := &sync.WaitGroup{}
	c := createGenerator(wg)
	fanOut(wg, c, 5)
	wg.Wait()
}
