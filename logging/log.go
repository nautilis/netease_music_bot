package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/nautilis/netease_music_bot/file"
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

// Setup initialize the log instance
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
func Debug(f string, v ...interface{}) {
	setPrefix(DEBUG)
	logger.Printf(f, v...)
}

func Debugf(f string, v ...interface{}) {
	setPrefix(DEBUG)
	logger.Printf(f, v...)
}

// Info output logs at info level
func Info(f string, v ...interface{}) {
	setPrefix(INFO)
	logger.Printf(f, v...)
}

func Infof(f string, v ...interface{}) {
	setPrefix(INFO)
	logger.Printf(f, v...)
}

// Warn output logs at warn level
func Warn(f string, v ...interface{}) {
	setPrefix(WARNING)
	logger.Printf(f, v...)
}

// Error output logs at error level
func Error(f string, v ...interface{}) {
	setPrefix(ERROR)
	logger.Printf(f, v...)
}

func Errorf(f string, v ...interface{}) {
	setPrefix(ERROR)
	logger.Printf(f, v...)
}

// Fatal output logs at fatal level
func Fatal(f string, v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalf(f, v...)
}

func Fatalf(f string, v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalf(f, v...)
}

// setPrefix set the prefix of the log output
func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}
