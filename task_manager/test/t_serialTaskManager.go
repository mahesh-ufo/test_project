package main

import (
	"fmt"
	"time"

	taskManager "github.com/mahesh-ufo/test_project/task_manager"
)

type myTaskReply struct {
	data string
}

func (tr *myTaskReply) getData() string {
	return tr.data

}

type myTask struct {
	Name     string
	CanSleep bool
}

func (mt *myTask) Execute() (taskManager.TaskReply, error) {

	fmt.Println("Executing task with name : [" + mt.Name + "]")

	if mt.CanSleep {
		time.Sleep(time.Second * 5)
	}

	var rep taskManager.TaskReply
	rep = myTaskReply{data: "Reply for task[" + mt.Name + "]"}
	return rep, nil
}

type application struct {
	mySerialTaskManager taskManager.TMInterface
}

func (app *application) taskThread(taskName string, canSleep bool) {
	var mTask = myTask{Name: taskName, CanSleep: canSleep}

	fmt.Println("Task thread started")

	var taskReply, err = app.mySerialTaskManager.RunTask(&mTask)

	if err != nil {
		fmt.Println("Failed to run the task")
		fmt.Println(err)
	} else {

		var mReply = taskReply.(myTaskReply)
		fmt.Println("Reply to thread is :" + mReply.getData())
		fmt.Println("***********************")
	}
}

func (app *application) startApplication() {

	app.mySerialTaskManager = &taskManager.SerialTaskManager{Name: "My Serial Task Manager"}
	app.mySerialTaskManager.StartTaskManager()

	go app.taskThread("Task1", true)
	go app.taskThread("Task2", false)
	go app.taskThread("Task3", true)
	go app.taskThread("Task4", false)
	go app.taskThread("Task5", true)
	go app.taskThread("Task6", false)
	go app.taskThread("Task7", true)
	go app.taskThread("Task8", false)
	go app.taskThread("Task9", true)

	for true {

	}
}

func main() {

	var myApp application
	myApp.startApplication()

}
