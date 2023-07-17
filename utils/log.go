package utils

import (
	"io"
	"os"

	config "github.com/delta/arcadia-backend/config"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *logrus.Logger

func InitLogger() {
	config := config.GetConfig()

	var (
		fileName = config.Log.FileName
		maxSize  = config.Log.MaxSize
		logLevel = config.Log.Level
	)

	if config.Log.FileName == "" {
		fileName = "./log.log"
	}

	if config.Log.MaxSize == 0 {
		maxSize = 50
	}

	if config.Log.Level == "" {
		logLevel = "info"
	}

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		panic(err)
	}

	var Out io.Writer

	if config.AppEnv == "DEV" {
		Out = io.Writer(os.Stdout)
	} else {
		Out = &lumberjack.Logger{
			Filename:   fileName,
			MaxSize:    maxSize, // in megabytes
			MaxBackups: 3,
		}
	}

	var Formatter logrus.Formatter

	if config.AppEnv == "DEV" {
		Formatter = &logrus.TextFormatter{
			DisableTimestamp: true,
		}
	} else {
		Formatter = &logrus.JSONFormatter{
			// Time stamp in DD-MM-YYYY HH:MM:SS format
			TimestampFormat: "02-01-2006 15:04:05",
		}
	}

	Logger = &logrus.Logger{
		Out:       Out,
		Level:     level,
		Formatter: Formatter,
	}

	Logger.Info("Logger started")
}

func GetLogger() *logrus.Logger {
	return Logger
}

func NewLogger(fileName string) *logrus.Logger {
	return &logrus.Logger{
		Out: &lumberjack.Logger{
			Filename:   fileName,
			MaxSize:    50, // in megabytes
			MaxBackups: 3,
		},
		Level: logrus.InfoLevel,
		Formatter: &logrus.JSONFormatter{
			// Time stamp in DD-MM-YYYY HH:MM:SS format
			TimestampFormat: "02-01-2006 15:04:05",
		},
	}
}

func GetControllerLogger(controller string) *logrus.Entry {
	return Logger.WithFields(logrus.Fields{
		"controller": controller,
	})
}

func GetControllerLoggerWithFields(controller string, fields map[string]interface{}) *logrus.Entry {
	return Logger.WithFields(logrus.Fields{
		"controller": controller,
		"param":      fields,
	})
}

func GetFunctionLogger(function string) *logrus.Entry {
	return Logger.WithFields(logrus.Fields{
		"function": function,
	})
}

func GetFunctionLoggerWithFields(function string, fields map[string]interface{}) *logrus.Entry {
	return Logger.WithFields(logrus.Fields{
		"function": function,
		"param":    fields,
	})
}
