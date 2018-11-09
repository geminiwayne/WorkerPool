package main

import (
	"WorkerPool"
	"strconv"
)

func main() {
	var test []string
	for i := 0; i < 50; i++ {
		test = append(test, strconv.Itoa(i))
	}
	works := WorkerPool.Works{}
	works.Tasks = test
	WorkerPool.Dispatcher(works, 10, "worker")
}
