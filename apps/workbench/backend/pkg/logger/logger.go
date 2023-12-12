package logger

import (
	"log"
	"os"
	"path"
	"time"

	"workbench/pkg/consts"

	"github.com/sirupsen/logrus"
)

var LoggersObj *logrus.Logger

func init() {
	if LoggersObj != nil {
		src, _ := setOutputFile()
		LoggersObj.Out = src
		return
	}
	logger := logrus.New()
	src, _ := setOutputFile()
	logger.Out = src
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	LoggersObj = logger
}

func setOutputFile() (*os.File, error) {

	now := time.Now()

	wd, err := os.Getwd()
	if err != nil {
		LoggersObj.Errorf("Get wd failed, err: %v", err)
	}

	logFilePath := path.Join(wd, consts.LoggerFilePath)
	_, err = os.Stat(logFilePath)

	if os.IsNotExist(err) {
		if err := os.MkdirAll(logFilePath, 0777); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	logFileName := now.Format(consts.TimeFormatToDate) + consts.LoggerFileSuffix
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return src, nil
}
