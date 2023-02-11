package config

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

var folderName string

func Loggers(param string, msg interface{}) {
	now := time.Now() //or time.Now().UTC()
	folderName = "storage/log/" + param + "/"
	logFileName := folderName + now.Format("2006-01-02") + ".log" //now.Format("2006-01-02") = now

	f, err := os.OpenFile(logFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to create logfile" + logFileName)
		panic(err)
	}
	defer f.Close()

	InfoLogger := &logrus.Logger{
		// Log into f file handler and on os.Stdout
		Out:   io.MultiWriter(f, os.Stdout),
		Level: logrus.DebugLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "[%lvl%]: %time% - %msg%\n",
		},
	}

	if param == "info" {
		InfoLogger.Info("Data => ", msg)
	}
	if param == "warning" {
		InfoLogger.Warning("Data => ", msg)
	}
	if param == "error" {
		InfoLogger.Error("Data => ", msg)
	}
	if param == "fatal" {
		InfoLogger.Fatal("Data => ", msg)
	}
}
