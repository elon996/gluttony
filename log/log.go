package log

import (
	"github.com/gluttony/color"
)

var Log = New()

type Logger struct {
	Level int
}

func New() *Logger {
	return &Logger{
		Level:        DebugLevel,
	}
}


var ErrorLevel = 4
var WarnLevel = 3
var InfoLevel = 2
var DebugLevel = 1


func (logger *Logger) SetLevel(level int) {
	logger.Level = level
}

func (logger *Logger) Debug(format string, args ...interface{}) {
	if logger.Level > 1 {
		return
	}
	color.Cyanf("[Debu] "+format+"\n", args ...)
}

func (logger *Logger) Info(format string, args ...interface{}) {
	if logger.Level > 2 {
		return
	}
	color.Greenf("[Info] "+format+"\n", args ...)
}

func (logger *Logger) Warn(format string, args ...interface{}) {
	if logger.Level > 3 {
		return
	}
	color.Yellowf("[Warn] "+format+"\n", args ...)
}

func (logger *Logger) Error(format string, args ...interface{}) {
	if logger.Level > 4 {
		return
	}
	color.Redf("[Erro] "+format+"\n", args ...)
}