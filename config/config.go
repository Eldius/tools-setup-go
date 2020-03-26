package config

import (
	"fmt"
	"os"

	home "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

/*
AppConfig is the configuration struct
*/
type AppConfig struct {
	Verbose   bool
	BinFolder string
}

const (
	logFolder = "~/.tools-setup/"
)

func init() {
	if configDir, err := home.Expand(logFolder); err != nil {
		panic(err.Error())
	} else {
		os.MkdirAll(configDir, os.ModePerm)
	}
}

/*
ExpandedBinFolder returns the parsed/expanded bin folder path
*/
func (s *AppConfig) ExpandedBinFolder() string {
	result, err := home.Expand(s.BinFolder)
	if err != nil {
		panic(err.Error())
	}
	return result
}

/*
LoadSetupSpecsConfig loads setup from file
*/
func LoadSetupSpecsConfig() AppConfig {
	fmt.Println("Using config file [2]:", viper.ConfigFileUsed())
	logFile, err := home.Expand(fmt.Sprintf("%s/execution.log", logFolder))
	if err != nil {
		panic(err.Error())
	}
	viper.SetDefault("verbose", false)
	viper.SetDefault("logfile", logFile)

	var appConfig AppConfig
	err = viper.Unmarshal(&appConfig)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(fmt.Sprintf("config: %v", appConfig))
	return appConfig
}

/*
GetLogFile returns the log file path
*/
func GetLogFile() string {
	return fmt.Sprintf("%s/execution.log", logFolder)
}
