package config

import (
	"fmt"
	"os"
	"path/filepath"

	home "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

/*
AppConfig is the configuration struct
*/
type AppConfig struct {
	Verbose   bool
	BinFolder string
	DbFolder  string
}

const (
	configDir = "~/.tools-setup"
)

func init() {
	configFolder := GetAppConfigFolder()
	os.MkdirAll(configFolder, os.ModePerm)
}

func createBinFolder(binFolder string) {
	os.MkdirAll(binFolder, os.ModePerm)
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
	logFile := filepath.Join(GetAppConfigFolder(), "execution.log")
	viper.SetDefault("verbose", false)
	viper.SetDefault("logfile", logFile)
	viper.SetDefault("binfolder", filepath.Join(GetAppConfigFolder(), "bin"))
	viper.SetDefault("dbfolder", filepath.Join(GetAppConfigFolder()))

	var appConfig AppConfig
	err := viper.Unmarshal(&appConfig)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(fmt.Sprintf("config: %v", appConfig))
	createBinFolder(appConfig.BinFolder)
	return appConfig
}

/*
GetLogFile returns the log file path
*/
func GetLogFile() string {
	logFile := filepath.Join(GetAppConfigFolder(), "execution.log")
	os.Create(logFile)
	return logFile
}

/*
GetDBFile returns the log file path
*/
func GetDBFile() string {
	return filepath.Join(GetAppConfigFolder(), "specs.db")
}

/*
GetAppConfigFolder returns the config folder
*/
func GetAppConfigFolder() string {
	if dir, err := home.Expand(configDir); err != nil {
		panic(err.Error())
	} else {
		return dir
	}
}
