package config

import (
	"log"
	"strconv"

	"github.com/knq/ini"
)

//SetConfig :
func (conf *IniConfig) SetConfig(confFileName string, logLevel int16) error {
	conf.configParams.LogLevel = logLevel

	confFile, err := ini.LoadFile(confFileName)

	if err != nil {
		log.Println(err)
		return err
	}

	testConf1, err := strconv.ParseInt(confFile.GetKey("General.GetTestConf1"), 10, 16)

	if err != nil {
		log.Println(err)
		return err
	}

	conf.configParams.TestConf1 = int16(testConf1)
	return nil
}

//GetLogLevel : get the log level
func (conf *IniConfig) GetLogLevel() int16 {
	return conf.configParams.LogLevel
}

//GetTestConf1 : get the GetTestConf1
func (conf *IniConfig) GetTestConf1() int16 {
	return conf.configParams.TestConf1
}
