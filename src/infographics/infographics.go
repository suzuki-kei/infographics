package infographics

import (
    "src/bigints"
)

/**
 *
 * インフォグラフィック文字列を生成するときのオプション.
 *
 */
type Options struct {
    // true の場合は短縮表現を生成する
    Short bool

    // 区切り文字
    Delimiter string
}

/**
 *
 * デフォルト値で初期化した Options を生成する.
 *
 */
func NewOptions() *Options {
    options := new(Options)
    options.Short = false
    options.Delimiter = " "
    return options
}

/**
 *
 * インフォグラフィック文字列を生成する.
 *
 */
func Generate(value string, options *Options) (string, error) {
    bigintValue, err := bigints.FromString(value)
    if err != nil {
        return "", err
    }

    generator := ChineseNumeralInfographicsGenerator{
        short: options.Short,
        delimiter: options.Delimiter,
    }
    return generator.Generate(bigintValue)
}

