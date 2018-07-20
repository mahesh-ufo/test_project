package main

//protoc *.proto --go_out=plugins=grpc:./

import (
	"fmt"

	config "github.com/mahesh-ufo/test_project/config"
	messaging "github.com/mahesh-ufo/test_project/messaging"
	taskManager "github.com/mahesh-ufo/test_project/task_manager"
)

func main() {
	var appConfig config.IniConfig
	var err = appConfig.SetConfig("../resources/testConfig.ini", 5)

	if err != nil {
		fmt.Println("Failed to read config file")
		return
	}

	var logLevel = appConfig.GetLogLevel()
	fmt.Printf("Log level is %d \n", logLevel)

	var mySerialTaskManager taskManager.TMInterface
	mySerialTaskManager = &taskManager.SerialTaskManager{Name: "My Serial Task Manager"}
	mySerialTaskManager.StartTaskManager()

	var server messaging.GRPCServer
	server.SetTaskManager(mySerialTaskManager)
	server.Start("", ":50001")
}
