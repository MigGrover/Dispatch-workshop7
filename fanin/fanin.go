package fanin

import (
	"fmt"
	"math/rand"
	"time"
)

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string, 2)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

func createWorker(id int) <-chan string {
	c := make(chan string)
	go func() {
		for {
			duration := rand.Intn(1e3)
			time.Sleep(time.Duration(duration) * time.Millisecond)
			fmt.Printf("Worker %v did something.\n", id)
			c <- fmt.Sprintf("Hello from worker %v\n", id)
		}
	}()

	return c
}

func ExecuteFanIn() {
	rand.Seed(time.Now().UnixNano())
	w1c := createWorker(1)
	w2c := createWorker(2)
	fIc := fanIn(w1c, w2c)
	for message := range fIc {
		fmt.Println(message)
	}
}
