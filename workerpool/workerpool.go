package workerpool

import (
	"log"
	"sync"
	"time"
)

type iworkerPool interface {
	run()
	addTask(task func())
	close()
}

type workerPool struct {
	maxWorker   int
	queuedTaskC chan func()
	waitGroup   *sync.WaitGroup
}

func newWorkerPool(maxWorker int, wg *sync.WaitGroup) iworkerPool {
	wp := &workerPool{
		maxWorker:   maxWorker,
		queuedTaskC: make(chan func()),
		waitGroup:   wg,
	}

	return wp
}

func (wp *workerPool) run() {
	for i := 0; i < wp.maxWorker; i++ {
		wID := i + 1
		log.Printf("[WorkerPool] Worker %d has been spawned", wID)
		wp.waitGroup.Add(1)
		go func(workerID int) {
			defer wp.waitGroup.Done()
			for task := range wp.queuedTaskC {
				log.Printf("[WorkerPool] Worker %d start processing task", wID)
				task()
				log.Printf("[WorkerPool] Worker %d finish processing task", wID)
			}
		}(wID)
	}
}

func (wp *workerPool) addTask(task func()) {
	wp.queuedTaskC <- task
}

func (wp *workerPool) close() {
	close(wp.queuedTaskC)
}

func ExecuteWorkerPool() {

	wg := &sync.WaitGroup{}

	totalWorker := 4
	wp := newWorkerPool(totalWorker, wg)
	wp.run()

	type result struct {
		id    int
		value int
	}

	totalTask := 20
	resultC := make(chan result, totalTask)
	defer close(resultC)

	for i := 0; i < totalTask; i++ {
		id := i + 1
		wp.addTask(func() {
			log.Printf("[main] Starting task %d", id)
			time.Sleep(2 * time.Second)
			resultC <- result{id, id * 2}
		})
	}

	for i := 0; i < totalTask; i++ {
		res := <-resultC
		log.Printf("[main] Task %d has been finished with result %d:", res.id, res.value)
	}

	wp.close()
	wg.Wait()
}
