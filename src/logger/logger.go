package logger

import (
	"fmt"
	"io"
	"time"
)

type LogLevel int

const (
	Info    LogLevel = 0
	Warning LogLevel = 1
	Error   LogLevel = 2
	Disable LogLevel = 3
)

type Message struct {
	Date    string
	Level   string
	Message string
	Step    string
}

type logger struct {
	out   io.Writer
	level LogLevel
}

type Logger interface {
	LogInfo(message string, step string)
	LogError(message string, step string)
	LogWarning(message string, step string)
}

func NewLogger(out io.Writer, level LogLevel) Logger {
	return &logger{
		out:   out,
		level: level,
	}
}

func log(out io.Writer, message Message) {
	fmt.Fprintf(out, "%+v\n", message)
}

func (l *logger) LogInfo(message string, step string) {
	if l.level > Info {
		return
	}

	log(l.out, Message{Date: time.Now().Local().String(), Level: "Info", Message: message, Step: step})
}

func (l *logger) LogError(message string, step string) {
	if l.level > Error {
		return
	}

	log(l.out, Message{Date: time.Now().Local().String(), Level: "Error", Message: message, Step: step})
}

func (l *logger) LogWarning(message string, step string) {
	if l.level > Warning {
		return
	}

	log(l.out, Message{Date: time.Now().Local().String(), Level: "Warning", Message: message, Step: step})
}
