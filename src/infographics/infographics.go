package infographics

import (
    "fmt"
    "src/bigints"
)

type SystemOfUnit int

const (
    // 漢数字.
    ChineseNumeral SystemOfUnit = iota
)

/**
 *
 * インフォグラフィック文字列を生成する.
 *
 */
func Generate(value string, short bool, delimiter string, systemOfUnit SystemOfUnit) (string, error) {
    bigintValue, err := bigints.FromString(value)
    if err != nil {
        return "", err
    }

    switch systemOfUnit {
        case ChineseNumeral:
            generator := ChineseNumeralInfographicsGenerator{
                short: short,
                delimiter: delimiter,
            }
            return generator.Generate(bigintValue)
        default:
            return "", fmt.Errorf("invalid SystemOfUnit: %v", systemOfUnit)
    }
}

