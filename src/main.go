package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    values, options := parseArguments()

    for _, value := range values {
        text, success := InfographicsTextFromString(value, options)
        if !success {
            error("invalid value: %v", value)
            continue
        }
        fmt.Printf("%v => %v\n", value, text)
    }
}

func parseArguments() ([]string, *InfographicsTextOptions) {
    values := []string {}
    options := NewInfographicsTextOptions()

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
        values = append(values, option)
    }
    return values, options
}

