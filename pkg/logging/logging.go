package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
)

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() Logger {
	return Logger{e}
}

func (l *Logger) GetLoggerWithField(k string, v interface{}) Logger {
	return Logger{l.WithField(k, v)}
}

func Init() {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s:%d", filename, f.Line), fmt.Sprintf("%s()", f.Function)
		},
		DisableColors: false,
		FullTimestamp: true,
	}

	_, err := os.Stat("logs")
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.Mkdir("logs", 0755); err != nil {
				panic(any(fmt.Sprintf("[Message]: %s", err)))
			}
		}
	}

	allFile, err := os.OpenFile("logs/all.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(any(fmt.Sprintf("[Message]: %s", err)))
	}

	l.SetOutput(allFile)

	l.SetLevel(logrus.TraceLevel)

	e = logrus.NewEntry(l)
}
