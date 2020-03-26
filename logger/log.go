package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	home "github.com/mitchellh/go-homedir"

	"github.com/Eldius/tools-setup-go/config"
)

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
Debug used to log debug massage
*/
func Debug(msg string) {
	cfg := config.LoadSetupSpecsConfig()
	if cfg.Verbose {
		f, err := os.OpenFile(getLogFile(), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(err.Error())
		}
		defer f.Close()

		log.SetOutput(f)
		log.Println(msg)
	}
}

/*
DebugInterface used to log debug info
*/
func DebugInterface(obj interface{}) {
	cfg := config.LoadSetupSpecsConfig()
	if cfg.Verbose {
		f, err := os.OpenFile(getLogFile(), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(err.Error())
		}
		defer f.Close()

		log.SetOutput(f)
		jsonb, err := json.Marshal(obj)
		if err != nil {
			panic(err.Error())
		}
		log.Println(string(jsonb))
		fmt.Println("i -> ", jsonb)
	}
}

/*
Info log info
*/
func Info(msg ...interface{}) {
	logFile := getLogFile()
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(f)
	}
	defer f.Close()

	log.Println(msg...)
	fmt.Println("i -> ", msg)
}

func getLogFile() string {
	logFile := config.GetLogFile()
	f, _ := os.Create(logFile)
	f.Close()

	return logFile
}
