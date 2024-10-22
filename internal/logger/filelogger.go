package logger

import (
	"os"
)

type fileLogger struct {
	file *os.File
}

func (f *fileLogger) Log(content string) {
	_, err := f.file.WriteString(content)
	if err != nil {
		panic(err)
	}
}

func FileLogger(dest string) fileLogger {
	file, err := os.OpenFile(dest, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	return fileLogger{file}
}
