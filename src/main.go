package main

import (
    "fmt"
)

func main() {
    stringValues := []string {
        "0",
        "12345",
        "70000",
        "100000",
        "1000000",
        "1800000",
        "10000000000",
        "17500000000",
    }

    for _, stringValue := range stringValues {
        infographicsText, success := InfographicsTextFromString(stringValue, " ")
        if !success {
            fmt.Println("ERROR: " + stringValue)
            continue
        }
        fmt.Println(stringValue + " => " + infographicsText)
    }
}

