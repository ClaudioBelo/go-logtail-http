package logtail

import (
	"fmt"
	"strings"
)

type Logger struct {
	LogTail *LogTail
}

func NewLogger(tail *LogTail) *Logger {
	return &Logger{
		LogTail: tail,
	}
}

func (l *Logger) Error(msg string) error {
	return l.logType("Error", msg)
}

func (l *Logger) Info(msg string) error {
	return l.logType("Info", msg)
}

func (l *Logger) Warning(msg string) error {
	return l.logType("Warning", msg)
}

func (l *Logger) Logf(fStr string, params ...interface{}) error {
	if strings.Count(fStr, "%") != len(params) {
		formatedMsg := fmt.Sprintf(fStr, params...)
		return l.LogTail.sendMessage(formatedMsg)
	} else {
		return fmt.Errorf("Logger.Logf: Number of format specifiers (%v) does not match number of parameters (%v)", strings.Count(fStr, "%"), len(params))
	}
}

func (l *Logger) logType(logType, msg string) error {
	formatedMsg := fmt.Sprintf("%v: %v", logType, msg)
	return l.LogTail.sendMessage(formatedMsg)
}
