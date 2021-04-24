package log

import "github.com/gluttony/color"

var Log = New()

type Logger struct {
	Level int
}

type Logsend struct {
	Level int
	Format string
	Args   []interface{}
}

func New() *Logger {
	return &Logger{
		Level:        DebugLevel,
	}
}

var c = make(chan Logsend, 10)

var ErrorLevel = 4
var WarnLevel = 3
var InfoLevel = 2
var DebugLevel = 1

func init()  {
	go Sendfor()
}

func Sendfor()  {
	for  {
		t := <-c
		switch t.Level {
		case 1:
			color.Cyanf("[Debu] "+t.Format+"\n", t.Args ...)
			break
		case 2:
			color.Greenf("[Info] "+t.Format+"\n", t.Args ...)
			break
		case 3:
			color.Yellowf("[Warn] "+t.Format+"\n", t.Args ...)
			break
		case 4:
			color.Redf("[Erro] "+t.Format+"\n", t.Args ...)
			break
		}
	}
}


func (logger *Logger) SetLevel(level int) {
	logger.Level = level
}

func (logger *Logger) Debug(format string, args ...interface{}) {
	if logger.Level > 1 {
		return
	}
	c <- Logsend{DebugLevel ,format, args}
}

func (logger *Logger) Info(format string, args ...interface{}) {
	if logger.Level > 2 {
		return
	}
	c <- Logsend{InfoLevel ,format, args}

}

func (logger *Logger) Warn(format string, args ...interface{}) {
	if logger.Level > 3 {
		return
	}
	c <- Logsend{WarnLevel ,format, args}

}

func (logger *Logger) Error(format string, args ...interface{}) {
	if logger.Level > 4 {
		return
	}
	c <- Logsend{ErrorLevel ,format, args}
}
