package infographics

import (
    "fmt"
)

import (
    "localhost/bigints"
)

type SystemOfUnit int

const (
    // 漢数字.
    // Ex. 一, 十, 百, 千, 万, ...
    ChineseNumeral SystemOfUnit = iota

    // SI 単位系.
    // Ex. kilo (10^3), mega (10^6), giga (10^9), ...
    SI SystemOfUnit = iota

    // IEC 単位系.
    // Ex. kibi (2^10), mebi (2^20), gibi (2^30), ...
    IEC SystemOfUnit = iota
)

/**
 *
 * インフォグラフィック文字列を生成する.
 *
 */
func Generate(value string, short bool, separator string, systemOfUnit SystemOfUnit) (string, error) {
    bigintValue, err := bigints.FromString(value)
    if err != nil {
        return "", err
    }

    switch systemOfUnit {
        case ChineseNumeral:
            generator := ChineseNumeralInfographicsGenerator{
                short: short,
                separator: separator,
            }
            return generator.Generate(bigintValue)
        case SI:
            generator := SiInfographicsGenerator{
                short: short,
                separator: separator,
            }
            return generator.Generate(bigintValue)
        case IEC:
            generator := IecInfographicsGenerator{
                short: short,
                separator: separator,
            }
            return generator.Generate(bigintValue)
        default:
            return "", fmt.Errorf("invalid SystemOfUnit: %v", systemOfUnit)
    }
}

