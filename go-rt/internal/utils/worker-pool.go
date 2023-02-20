package utils

import "sync"

type Task func()

type WorkerPool struct {
	maxConcurrency int
	taskChannel    chan Task
	waitGroup      sync.WaitGroup
}

func NewWorkerPool(workerCount int) WorkerPool {
	return WorkerPool{maxConcurrency: workerCount, taskChannel: make(chan Task)}
}

func (this *WorkerPool) Run() {
	for i := 0; i < this.maxConcurrency; i++ {
		go this.work()
	}
}

func (this *WorkerPool) Wait() {
	this.waitGroup.Wait()
}

func (this *WorkerPool) Close() {
	close(this.taskChannel)
	this.waitGroup.Wait()
}

func (this *WorkerPool) AddTask(task Task) {
	this.waitGroup.Add(1)
	this.taskChannel <- task
}

func (this *WorkerPool) work() {
	for task := range this.taskChannel {
		task()
		this.waitGroup.Done()
	}
}
