package generator

import (
	"fmt"
	"math/rand"
	"time"
)

func createGenerator(message string) <-chan bool {
	c := make(chan bool)
	go func() {
		defer close(c)
		duration := rand.Intn(1e3)
		time.Sleep(time.Duration(duration) * time.Millisecond)
		fmt.Printf("I did something in parallel and got '%s' as message and got me %vms to finish.\n", message, duration)
		c <- true
	}()

	return c
}

func ExecuteGenerator() {
	rand.Seed(time.Now().UnixNano())
	generator1 := createGenerator("1")
	generator2 := createGenerator("2")
	generator3 := createGenerator("3")
	generator4 := createGenerator("4")
	if <-generator3 && <-generator4 {
		fmt.Println("Generators 3 and 4 finished")
	}
	if <-generator1 {
		fmt.Println("Generator 1 finished")
	}
	if <-generator2 {
		fmt.Println("Generator 2 finished")
	}
}
