package main

const (
	maxJobsInQueue = 100
)

type Dispatcher struct {
	taskChan chan Task
}

// TODO: Init Dispatcher
func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		taskChan: make(chan Task, maxJobsInQueue),
	}
}

// Start starts to consume tasks
func (d *Dispatcher) Start() {
	// TODO: Implement Dispatcher Run
}

func (d *Dispatcher) AddTask(task Task) {
	d.taskChan <- task
}
