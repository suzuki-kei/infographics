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
        infographicsText, success := InfographicsTextFromString(stringValue, " ")
        if !success {
            fmt.Println("ERROR: " + stringValue)
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

func parseArguments() Options {
    var options Options

    // TODO: Options.delimiter も解釈する.

    for _, arg := range os.Args[1:] {
        if arg == "-s" || arg == "--short" {
            options.short = true
            continue
        }
        if strings.HasPrefix(arg, "-") {
            fmt.Println("WARN: Invalid option [" + arg + "]")
            continue
        }
        options.values = append(options.values, arg)
    }

    return options
}

