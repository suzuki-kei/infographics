package main

import (
    "fmt"
    "io"
    "os"
)

func Error(message string, values ...any) {
    log(os.Stderr, "ERROR", message, values...)
}

func Warn(message string, values ...any) {
    log(os.Stderr, "WARN", message, values...)
}

func log(writer io.Writer, level string, message string, values ...any) {
    format := fmt.Sprintf("[%s] %s\n", level, message)
    fmt.Fprintf(writer, format, values...)
}

