package WorkerPool

import (
	"ProgressBar"
	"strconv"
	"sync"
)

type Worker struct {
	Id int
}

type WorkerPool struct {
	Workers          []Worker
	WorkerTasksIndex chan int
	WorkerTasks      chan string
	WorkerError      chan error
}

type Works struct {
	Tasks []string
	Wg    sync.WaitGroup
}

// initilize workers in worker pool
func (p *WorkerPool) New(workerNum int) []Worker {
	var workers []Worker
	for i := 1; i <= workerNum; i++ {
		worker := Worker{}
		worker.Id = i
		workers = append(workers, worker)
	}
	return workers
}

// A worker process a task
func (w Worker) Run(pool WorkerPool, wg *sync.WaitGroup, progress ProgressBar.Progress) {
	for i := range pool.WorkerTasks {
		temp, _ := strconv.Atoi(i)
		temp = temp * 2
		progress.Add(<-pool.WorkerTasksIndex, progress)
		wg.Done()
	}
}

// dispatch tasks to worker in the worker pool and the concurrency len is the length of worker pool
func Dispatcher(works Works, workerNum int, taskName string) {
	progress := ProgressBar.Progress{}
	progress.TaskName = taskName
	progress.TotalNum = len(works.Tasks)
	progress.Init(progress)
	pool := WorkerPool{}
	pool.Workers = pool.New(workerNum)
	pool.WorkerTasks = make(chan string, workerNum)
	pool.WorkerTasksIndex = make(chan int, workerNum)
	for _, j := range pool.Workers {
		go j.Run(pool, &works.Wg, progress)
	}

	for _, i := range works.Tasks {
		pool.WorkerTasks <- i
		val, _ := strconv.Atoi(i)
		pool.WorkerTasksIndex <- val + 1
		works.Wg.Add(1)
	}
	close(pool.WorkerTasks)
	works.Wg.Wait()
	progress.Add(len(works.Tasks), progress)
}
