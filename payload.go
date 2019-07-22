package main

import (
	"math/rand"
	"time"
)

// Task represents the job to be run
type Task struct {
	Payload Payload
}

func (t *Task) Do() {
	randomInt := rand.Intn(20)

	<-time.After(time.Duration(randomInt) * time.Second)
}

type Payload struct {
	Task string `json:"task"`
}
