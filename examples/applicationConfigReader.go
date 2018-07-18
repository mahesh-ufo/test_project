package main

import (
	"fmt"

	config "github.com/mahesh-ufo/test_project/config"
)

func main() {
	appConfig := config.IniConfig{}
	err := appConfig.SetConfig("testConfig", 5)

	if err != nil {
		fmt.Println("Failed to read config file")
		return
	}

	logLevel := appConfig.GetLogLevel()
	fmt.Printf("Log level is %d \n", logLevel)
}
