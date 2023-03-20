package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Logger struct {
	logWriter io.Writer
}

func (l *Logger) Log(message string) error {
	_, err := l.logWriter.Write([]byte(message))
	return err
}

type LoggerDecorator struct {
	Logger
}

func (ld *LoggerDecorator) Log(message string) error {
	message = fmt.Sprintf("[%s] %s", time.Now().Format(time.RFC3339), message)
	return ld.Logger.Log(message)
}

func main() {

	a := Logger{os.Stdout}
	err := a.Log("10\n")
	fmt.Println(err)

}
