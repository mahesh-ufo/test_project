package config

//configParams : Application Configuration
type configParams struct {
	LogLevel  int16
	TestConf1 int16
}

//ConfigurationReader interface which should be implemented
type ConfigurationReader interface {
	SetConfig(confFileName string, logLevel int16) error
	GetLogLevel() int16
	GetTestConf1() int16
}

//IniConfig : config to read file from INI type config file
//Method are implemented in IniConfigreader.go
type IniConfig struct {
	configParams
}
