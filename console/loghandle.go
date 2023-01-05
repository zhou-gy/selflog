package console

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

const (
	UNKNOW uint16 = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

type logger struct {
	Level uint16
}

type Logger logger

func LogLevel(level string) uint16 {
	level = strings.ToUpper(level)
	switch level {
	case "DEBUG":
		return DEBUG
	case "INFO":
		return INFO
	case "WARN":
		return WARN
	case "ERROR":
		return ERROR
	case "FATAL":
		return FATAL
	default:
		return UNKNOW
	}
}

// level => ["DEBUG", "INFO", "WARN", "ERROR", "FATAL" ]
func (l Logger) SetLevel(level string) Logger {
	l.Level = LogLevel(level)
	return l
}

func (l *Logger) log(levelstr, format string, arg ...interface{}) {
	if l.Level <= LogLevel(levelstr) {
		s := fmt.Sprintf("%s [%s] [%s] %s", time.Now().Format("2006-01-02 15:04:05"), levelstr,getFuncName() ,fmt.Sprintf(format, arg...))
		fmt.Println(s)
	}
}

func (l *Logger) DEBUG(format string, arg ...interface{}) {
	l.log("DEBUG", format, arg...)
}

func (l *Logger) INFO(format string, arg ...interface{}) {
	l.log("INFO", format, arg...)
}

func (l *Logger) WARN(format string, arg ...interface{}) {
	l.log("WARN", format, arg...)
}

func (l *Logger) ERROR(format string, arg ...interface{}) {
	l.log("ERROR", format, arg...)
}

func (l *Logger) FATAL(format string, arg ...interface{}) {
	l.log("FATAL", format, arg...)
}

func getFuncName() string{
	pc ,file ,line , ok := runtime.Caller(3)
	if !ok {
		fmt.Println("runtime.Caller() err")
		return ""
	}
	funcname := runtime.FuncForPC(pc).Name()
	filenameList := strings.Split(file,"/")
	filename := filenameList[len(filenameList)-1]
	return fmt.Sprintf("%s => %s ,line %d",filename,funcname,line)
}


