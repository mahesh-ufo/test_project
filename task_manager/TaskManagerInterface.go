package taskmanager

//TMInterface : Task manager interface
type TMInterface interface {
	StartTaskManager() error
	RunTask(Task) (TaskReply, error)
}

//Task to perform
type Task interface {
	Execute() (TaskReply, error)
}

//TaskReply : reply
type TaskReply interface {
}
