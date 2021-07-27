package logging

import (
	"log"
	"os"
)

// Custom logging with dev/prod tunning
// Inspired by: https://www.honeybadger.io/blog/golang-logging/

const (
	environmentEnvKey   = "env"
	logsFilePathEnvKey  = "logging.absoluteFilePath"
	infoLoggerPrefix    = "INFO: "
	warningLoggerPrefix = "WARNING: "
	errorLoggerPrefix   = "ERROR: "
	loggerFlag          = log.Ldate | log.Ltime | log.Lshortfile | log.Lmicroseconds
)

var (
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
)

func init() {

	environment := os.Getenv(environmentEnvKey)

	if environment == "production" {

		loggingFilePath := os.Getenv(logsFilePathEnvKey)

		file, err := os.OpenFile(loggingFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0770)

		if err != nil {
			panic("Could not open logging file")
		}

		infoLogger = log.New(file, infoLoggerPrefix, loggerFlag)
		warningLogger = log.New(file, warningLoggerPrefix, loggerFlag)
		errorLogger = log.New(file, errorLoggerPrefix, loggerFlag)
	} else {
		infoLogger = log.New(os.Stdout, infoLoggerPrefix, loggerFlag)
		warningLogger = log.New(os.Stdout, warningLoggerPrefix, loggerFlag)
		errorLogger = log.New(os.Stderr, errorLoggerPrefix, loggerFlag)
	}
}

func LogInfo(msg string) {
	infoLogger.Println(msg)
}

func LogWarning(msg string) {
	warningLogger.Println(msg)
}

func LogError(msg string) {
	errorLogger.Println(msg)
}
