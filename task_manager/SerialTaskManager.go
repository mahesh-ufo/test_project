package taskmanager

import (
	"fmt"
	"time"
)

type taskResponse struct {
	TaskReply
}

func (tr *taskResponse) setTaskReply(myTaskReply TaskReply) error {
	tr.TaskReply = myTaskReply
	return nil
}

func (tr *taskResponse) getTaskReply() TaskReply {

	return tr.TaskReply
}

type taskRequest struct {
	responseChannel        chan taskResponse
	responseChannelCreated bool
	Task
}

func (tr *taskRequest) setTask(myTask Task) error {
	tr.Task = myTask
	return nil
}

func (tr *taskRequest) getResponseChannel() chan taskResponse {

	if !tr.responseChannelCreated {
		tr.responseChannel = make(chan taskResponse)
		tr.responseChannelCreated = true
		fmt.Println("crating response channel")
	}

	fmt.Println(" response channel already created")
	return tr.responseChannel
}

//SerialTaskManager : Implements TMInterface interface
type SerialTaskManager struct {
	Name          string
	workerChannel chan taskRequest
}

func (taskManager *SerialTaskManager) startWorker() {
	for true {
		var reqTask taskRequest
		fmt.Println("Listening for Request in task channel")

		reqTask = <-taskManager.workerChannel
		fmt.Println("Request came to worker task channel")
		time.Sleep(time.Millisecond * 50)

		taskReply, err := reqTask.Task.Execute()

		if err != nil {
			fmt.Println("error while processing task")
		}

		responseChannel := reqTask.getResponseChannel()
		var tResp taskResponse
		tResp.setTaskReply(taskReply)

		responseChannel <- tResp

	}

}

//StartTaskManager : starts task manager
func (taskManager *SerialTaskManager) StartTaskManager() error {
	taskManager.workerChannel = make(chan taskRequest)
	go taskManager.startWorker()

	return nil
}

//RunTask : Queue up the task serial execution and wait for response
func (taskManager *SerialTaskManager) RunTask(myTaskToRun Task) (TaskReply, error) {
	fmt.Printf("Executing task for [" + taskManager.Name + "]")

	var t taskRequest
	t.setTask(myTaskToRun)
	var responseChannel = t.getResponseChannel()
	taskManager.workerChannel <- t

	fmt.Println("Data send to serial task worker channel")
	var resp taskResponse
	resp = <-responseChannel

	return resp.getTaskReply(), nil
}
