package logger

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type LogLevel int32
type LogLabel string

const (
	Undefined LogLevel = iota
	_
	Loggers
	_
	Debuggers
	_
	Diebuggers
	_
	Alerts
	_
	Errors
)

const (
	Logger LogLabel = "LOG"
	Debug  LogLabel = "DEBUG"
	Diebug LogLabel = "DIEBUG"
	Alert  LogLabel = "ALERT"
	Error  LogLabel = "ERROR"
)

type Logging interface {
	NewLog(level string, msg string, data ...interface{})
}

type GenLog interface {
	GenerateLog(level LogLabel, message string, data ...interface{})
}

type logs struct{}

func NewLogging() Logging {
	return logs{}
}

func (l logs) NewLog(level string, msg string, data ...interface{}) {

	switch filterLogginglevel(level) {
	case Loggers:
		l.GenerateLog(Logger, msg, data)
	case Debuggers:
		l.GenerateLog(Debug, msg, data)
	case Diebuggers:
		l.GenerateLog(Diebug, msg, data)
	case Errors:
		l.GenerateLog(Error, msg, data)
	case Alerts:
		l.GenerateLog(Alert, msg, data)
	default:
		l.Error(fmt.Errorf("\nlog level not found, try to use: log, debug, alert or error\n"))
	}
}

func (l logs) GenerateLog(level LogLabel, message string, data ...interface{}) {
	var outReturn = make([]byte, 1024)
	for d := range data {
		if data == nil || d == 0 {
			out, err := exec.Command("echo", "-e", fmt.Sprintf(`\e%s[%s]\e[0m message: %v`, colorByLabel(level), level, message)).Output()
			if err != nil {
				l.Error(err)
			}
			outReturn = out
		} else {
			out, err := exec.Command("echo", "-e", fmt.Sprintf(`\e%s[%s]\e[0m message: %v | %v`, colorByLabel(level), level, message, data)).Output()
			if err != nil {
				l.Error(err)
			}
			outReturn = out
		}
	}
	if level == Diebug {
		log.Fatal("\n", string(outReturn), "\n")
	}
	log.Print("\n", string(outReturn), "\n")
}

func (l logs) Error(err error) {
	out, err := exec.Command("echo", "-e", fmt.Sprintf(`\e[31m[ERROR]\e[0m error: %v`, err)).Output()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Fatal("\n", string(out), "\n")
}

func colorByLabel(level LogLabel) string {
	switch level {
	case "LOG":
		return "[32m"
	case "DEBUG":
		return "[36m"
	case "DIEBUG":
		return "[36m"
	case "ALERT":
		return "[33m"
	case "ERROR":
		return "[31m"
	default:
		return "[37m"
	}
}

func filterLogginglevel(level string) LogLevel {
	switch {
	case strings.ToLower(level) == "log":
		return 2
	case strings.ToLower(level) == "debug":
		return 4
	case strings.ToLower(level) == "diebug":
		return 6
	case strings.ToLower(level) == "alert":
		return 8
	case strings.ToLower(level) == "error":
		return 10
	default:
		return 0
	}
}
