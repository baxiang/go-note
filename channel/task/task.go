package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
// 任务
type Task struct {
	id      int
	Content int64
}
type Result struct {
	task   Task
	result int64
}

var tasks = make(chan Task, 10)
var taskResults = make(chan Result, 10)

func work(num int64) int64 {
	// 模拟工作
	randTime := rand.Intn(10)
	time.Sleep(time.Duration(randTime)*time.Second)
	return num+time.Now().Unix()
}
func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		result := Result{task, work(task.Content)}
		taskResults <- result
	}
}

func makeWorkerPool(numOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < numOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	//当所有工作结束时
	close(taskResults)
}
func allocate(numOfTasks int) {
	for i := 0; i < numOfTasks; i++ {
		task := Task{i, time.Now().Unix()}
		tasks <- task
	}
	close(tasks)
}
func getResult(done chan bool) {
	for result := range taskResults {
		fmt.Printf("Task id %d, randnum %d , sum %d\n", result.task.id, result.task.Content, result.result)
	}
	done <- true
}
func main() {
	startTime := time.Now()
	numOfWorkers := 20
	numOfTasks := 200

	var done = make(chan bool)
	go allocate(numOfTasks)
	go makeWorkerPool(numOfWorkers)

	go getResult(done)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
