package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Eldius/tools-setup-go/config"
)

/*
Debug used to log debug massage
*/
func Debug(msg string) {
	cfg := config.LoadSetupSpecsConfig()
	if cfg.Verbose {
		f := getLogFile()
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
		f := getLogFile()
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
	f := getLogFile()
	defer f.Close()

	log.Println(msg...)
	fmt.Println("i -> ", msg)
}

func getLogFile() *os.File {
	logFile := config.GetLogFile()
	f, openFileErr := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if openFileErr != nil {
		f, createFileErr := os.Create(logFile)
		if createFileErr != nil {
			panic(createFileErr.Error())
		}
		return f
	}

	return f
}
