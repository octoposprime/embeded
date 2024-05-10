package logger

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

var loggerInstance *Logger

func init() {
	GetLoggerInstance()
}

type Logger struct {
	Debug bool
}

func GetLoggerInstance() *Logger {
	if loggerInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if loggerInstance == nil {
			loggerInstance = &Logger{}
		}
	}
	return loggerInstance
}

func (l *Logger) SetDebug(debug bool) {
	l.Debug = debug
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	if l.Debug {
		fmt.Printf(format, args...)
	}
}

func (l *Logger) Infof(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
