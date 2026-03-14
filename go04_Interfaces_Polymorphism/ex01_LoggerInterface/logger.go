package main

import "os"

type Logger interface {
	Log(level, message string)
}

type ConsoleLogger struct{}

type FileLogger struct {
	file *os.File
}

type MultiLogger struct {
	Loggers *[]Logger
}
