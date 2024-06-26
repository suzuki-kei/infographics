package application

import (
    "fmt"
    "io"
    "log"
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
    separator string

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
        short: true,
        separator: " ",
        systemOfUnit: infographics.ChineseNumeral,
    }
}

func Run() {
    options, values := getOptionsAndValues()
    short := options.short
    separator := options.separator
    systemOfUnit := options.systemOfUnit

    for _, value := range values {
        text, err := infographics.Generate(value, short, separator, systemOfUnit)
        if err != nil {
            logging.Error(err.Error())
            continue
        }
        fmt.Printf("%v => %v\n", value, text)
    }
}

/**
 *
 * オプションと変換対象の値を読み込む.
 *
 */
func getOptionsAndValues() (*Options, []string) {
    values, options := parseArguments(os.Args[1:])
    if len(values) != 0 {
        return options, values
    }

    lines := readLines(os.Stdin)
    return options, lines
}

/**
 *
 * io.Reader から全てのデータをテキストとして読み込む.
 *
 */
func readLines(r io.Reader) []string {
    bytes, err := io.ReadAll(r)
    if err != nil {
        log.Fatal(err)
    }

    text := strings.TrimSuffix(string(bytes), "\n")
    return strings.Split(text, "\n")
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
        if option == "-l" || option == "--long" {
            options.short = false
            continue
        }
        if strings.HasPrefix(option, "--separator=") {
            separator := strings.Replace(option, "--separator=", "", 1)
            options.separator = separator
            continue
        }
        if option == "-s" || option == "--separator" {
            i++
            options.separator = arguments[i]
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

