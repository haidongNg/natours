package logging

import (
	"fmt"
	"log"
	"natours/pkg/file"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func Setup() {
	var err error
	filePath := getLogFilePath()
	fileName := getLogFileName()
	F, err = file.MustOpen(fileName, filePath)

	if err != nil {
		log.Fatalf("logging.Setup err: %v", err)
	}
	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

// Debug output logs at debug level
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Printf("%s", v)
}

// Info output logs at debug level
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Printf("%s", v)
}

// Warming output logs at debug level
func Warming(v ...interface{}) {
	setPrefix(WARNING)
	logger.Printf("%s", v)
}

// Error output logs at debug level
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Printf("%s", v)
}

// Fatal output logs at debug level
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Printf("%s", v)
}

// setPrefix set the prefix of log output
func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)

	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}
