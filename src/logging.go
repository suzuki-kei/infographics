package main

import (
    "fmt"
    "io"
    "os"
)

/**
 *
 * ERROR レベルのログを出力する.
 *
 */
func Error(message string, values ...any) {
    log(os.Stderr, "ERROR", message, values...)
}

/**
 *
 * WARN レベルのログを出力する.
 *
 */
func Warn(message string, values ...any) {
    log(os.Stderr, "WARN", message, values...)
}

/**
 *
 * ログ出力の共通処理.
 *
 */
func log(writer io.Writer, level string, message string, values ...any) {
    format := fmt.Sprintf("[%s] %s\n", level, message)
    fmt.Fprintf(writer, format, values...)
}

