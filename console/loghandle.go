package console


import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
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
	Ptype string
}
type jsonlog struct {
	Time string
	Level string
	Caller string
	Msg string
}

//type Logger logger
var Logger logger

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
func (l logger) SetLevel(lv ,logType string) logger {
	l.Level = LogLevel(lv)
	l.Ptype = logType
	return l
}

func colorPrint(levelStr, msg string)  {
	switch levelStr {
	case "DEBUG":
		color.Cyan(msg)
		//color.White(msg)
	case "INFO":
		color.Green(msg)
	case "WARN":
		color.Yellow(msg)
	case "ERROR":
		color.Red(msg)
	default:
		color.Black(msg)
	}
}

func (l *logger) log(levelstr, format string, arg ...interface{}) {
	if l.Level <= LogLevel(levelstr) {

		switch strings.ToUpper(l.Ptype) {
		case "JSON":
			logJson := jsonlog{
				Time: time.Now().Format("2006-01-02 15:04:05"),
				Level: levelstr,
				Caller: getFuncName(),
				Msg: fmt.Sprintf(format, arg...),
			}
			b, err := json.Marshal(logJson)
			if err != nil {
				fmt.Printf("json Marshal err: %v\n",err)
				return
			}
			colorPrint(levelstr,string(b))
			//fmt.Printf("%s\n",b)

		default:
			s := fmt.Sprintf("%s [%s] [%s] %s", time.Now().Format("2006-01-02 15:04:05"), levelstr,getFuncName() ,fmt.Sprintf(format, arg...))
			//fmt.Println(s)
			colorPrint(levelstr,s)
		}
	}
}

func (l *logger) DEBUG(format string, arg ...interface{}) {
	l.log("DEBUG", format, arg...)
}

func (l *logger) INFO(format string, arg ...interface{}) {
	l.log("INFO", format, arg...)
}

func (l *logger) WARN(format string, arg ...interface{}) {
	l.log("WARN", format, arg...)
}

func (l *logger) ERROR(format string, arg ...interface{}) {
	l.log("ERROR", format, arg...)
}

func (l *logger) FATAL(format string, arg ...interface{}) {
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
	return fmt.Sprintf("%s → %s ,line %d",filename,funcname,line)
}

