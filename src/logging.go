package main

import (
    "fmt"
    "io"
    "os"
)

func log(writer io.Writer, level string, message string, values ...any) {
    format := fmt.Sprintf("[%s] %s\n", level, message)
    fmt.Fprintf(writer, format, values...)
}

func error(message string, values ...any) {
    log(os.Stderr, "ERROR", message, values...)
}

func warn(message string, values ...any) {
    log(os.Stderr, "WARN", message, values...)
}

