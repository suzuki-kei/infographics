package application

import (
    "fmt"
    "os"
    "src/infographics"
    "src/logging"
    "strings"
)

func Run() {
    values, options := parseArguments(os.Args[1:])

    for _, value := range values {
        text, err := infographics.TextFromString(value, options)
        if err != nil {
            logging.Error(err.Error())
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
func parseArguments(arguments []string) ([]string, *infographics.Options) {
    values := []string{}
    options := infographics.NewOptions()

    for i := 0; i < len(arguments); i++ {
        option := arguments[i]

        if option == "--" {
            values = append(values, arguments[i + 1:]...)
            break
        }
        if option == "-s" || option == "--short" {
            options.Short = true
            continue
        }
        if strings.HasPrefix(option, "--delimiter=") {
            delimiter := strings.Replace(option, "--delimiter=", "", 1)
            options.Delimiter = delimiter
            continue
        }
        if option == "-d" || option == "--delimiter" {
            i++
            options.Delimiter = arguments[i]
            continue
        }
        if strings.HasPrefix(option, "-") {
            logging.Warn("invalid options: %v", option)
            continue
        }
        values = append(values, option)
    }
    return values, options
}

