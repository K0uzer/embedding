package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type TimestampWriter struct {
    io.Writer
}

func (tw *TimestampWriter) Write(p []byte) (int, error) {
	ts := time.Now().Format("2006-01-02 15:04:05")
	prefix := []byte(fmt.Sprintf("[%s] ", ts))
	tw.Writer.Write(prefix)
	return tw.Writer.Write(p)
}  

func main() {
	tw := &TimestampWriter{Writer: os.Stdout}
	fmt.Fprintln(tw, "сервер запущен")
	// [2024-01-15 14:30:00] сервер запущен
}