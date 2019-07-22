package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	MaxWorkers     = 10
)

var (
	dispatcher = NewDispatcher()
)
func main() {
	// start running dispatcher
	dispatcher.Start()

	// start serve the endpoint for adding messages to job queue
	http.HandleFunc("/", payloadHandler)
	http.ListenAndServe(":4000", nil)
}


func payloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Read body
	b, err := ioutil.ReadAll(r.Body)
	//defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var payload Payload
	err = json.Unmarshal(b, &payload)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// let's create a task with the payload
	task := Task{Payload: payload}

	// Push the work onto the queue.
	fmt.Println("[payloadHandler]", task, "-> TaskQueue")
	dispatcher.AddTask(task)


	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}