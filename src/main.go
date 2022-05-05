package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    values, options := parseArguments()

    for _, value := range values {
        text, err := InfographicsTextFromString(value, options)
        if err != nil {
            Error(err.Error())
            continue
        }
        fmt.Printf("%v => %v\n", value, text)
    }
}

/**
 *
 * コマンドライン引数を解析し, (values, options) を返す.
 *
 * values には変換対象の数値文字列 (Ex. "123") が含まれる.
 *
 */
func parseArguments() ([]string, *InfographicsTextOptions) {
    values := []string{}
    options := NewInfographicsTextOptions()

    for i := 1; i < len(os.Args); i++ {
        option := os.Args[i]

        if option == "--" {
            values = append(values, os.Args[i + 1:]...)
            break
        }
        if option == "-s" || option == "--short" {
            options.short = true
            continue
        }
        if strings.HasPrefix(option, "--delimiter=") {
            delimiter := strings.Replace(option, "--delimiter=", "", 1)
            options.delimiter = delimiter
            continue
        }
        if option == "-d" || option == "--delimiter" {
            i++
            options.delimiter = os.Args[i]
            continue
        }
        if strings.HasPrefix(option, "-") {
            Warn("invalid options: %v", option)
            continue
        }
        values = append(values, option)
    }
    return values, options
}

