package infographics

import (
    "src/bigints"
)

/**
 *
 * インフォグラフィック文字列を生成する.
 *
 */
func Generate(value string, short bool, delimiter string) (string, error) {
    bigintValue, err := bigints.FromString(value)
    if err != nil {
        return "", err
    }

    generator := ChineseNumeralInfographicsGenerator{
        short: short,
        delimiter: delimiter,
    }
    return generator.Generate(bigintValue)
}

