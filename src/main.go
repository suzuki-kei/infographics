package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    options := parseArguments()
    // TODO: options.short を考慮する.

    for _, stringValue := range options.values {
        infographicsText, success := InfographicsTextFromString(stringValue, options.delimiter)
        if !success {
            error("invalid value: %v", stringValue)
            continue
        }
        fmt.Println(stringValue + " => " + infographicsText)
    }
}

type Options struct {
    short bool
    delimiter string
    values []string
}

func NewOptions() *Options {
    options := new(Options)
    options.short = false
    options.delimiter = " "
    options.values = []string {}
    return options
}

func parseArguments() *Options {
    options := NewOptions()

    for i := 1; i < len(os.Args); i++ {
        option := os.Args[i]

        if option == "-s" || option == "--short" {
            options.short = true
            continue
        }
        if option == "-d" || option == "--delimiter" {
            i++
            options.delimiter = os.Args[i]
            continue
        }
        if strings.HasPrefix(option, "-") {
            warn("invalid options: %v", option)
            continue
        }
        options.values = append(options.values, option)
    }
    return options
}

