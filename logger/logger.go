package logger

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger
var dirLogs string = "logs/"

func init() {
	Log = logrus.New()

	if _, err := os.Stat(dirLogs); os.IsNotExist(err) {
		os.Mkdir(dirLogs, os.ModePerm)
	}

	logFileName := dirLogs + "sys_monitor_api_" + time.Now().Format("2006-01-02_15") + ".log"

	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Log.Fatal("Erro in creating log file: ", err)
	}

	Log.SetOutput(logFile)
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	Log.SetLevel(logrus.InfoLevel)
}
