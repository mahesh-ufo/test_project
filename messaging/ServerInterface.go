package messaging

import (
	taskManager "github.com/mahesh-ufo/test_project/task_manager"
)

//ServerInterface : Server interface
type ServerInterface interface {
	SetTaskManager(taskManager taskManager.TMInterface) error
	Start(ip string, port string) error
}

// GRPCServer is used to implement ServerInterface and TestMessagingServer
type GRPCServer struct {
	taskManager taskManager.TMInterface
}
