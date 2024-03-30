package application

import (
    "fmt"
    "os"
    "strings"
)

import (
    "localhost/infographics"
    "localhost/logging"
)

/**
 *
 * インフォグラフィック文字列を生成するときのオプション.
 *
 */
type Options struct {
    // true の場合は短縮表現を生成する.
    short bool

    // 区切り文字.
    delimiter string

    // 単位系.
    systemOfUnit infographics.SystemOfUnit
}

/**
 *
 * デフォルト値で初期化した Options を生成する.
 *
 */
func newOptions() *Options {
    return &Options{
        short: false,
        delimiter: " ",
        systemOfUnit: infographics.ChineseNumeral,
    }
}

func Run() {
    values, options := parseArguments(os.Args[1:])
    short := options.short
    delimiter := options.delimiter
    systemOfUnit := options.systemOfUnit

    for _, value := range values {
        text, err := infographics.Generate(value, short, delimiter, systemOfUnit)
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
func parseArguments(arguments []string) ([]string, *Options) {
    values := []string{}
    options := newOptions()

    for i := 0; i < len(arguments); i++ {
        option := arguments[i]

        if option == "--" {
            values = append(values, arguments[i + 1:]...)
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
            options.delimiter = arguments[i]
            continue
        }
        if option == "--chinese-numeral" {
            options.systemOfUnit = infographics.ChineseNumeral
            continue
        }
        if option == "--si" {
            options.systemOfUnit = infographics.SI
            continue
        }
        if option == "--iec" {
            options.systemOfUnit = infographics.IEC
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

