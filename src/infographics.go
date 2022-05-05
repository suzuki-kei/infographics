package main

import (
    "fmt"
    "math/big"
    "sort"
    "strings"
)

type InfographicsTextOptions struct {
    // true の場合は短縮表現を生成する
    short bool

    // 区切り文字
    delimiter string
}

func NewInfographicsTextOptions() *InfographicsTextOptions {
    options := new(InfographicsTextOptions)
    options.short = false
    options.delimiter = " "
    return options
}

func InfographicsTextFromString(
        value string, options *InfographicsTextOptions) (string, bool) {
    bigintValue, success := BigIntFromString(value)
    if !success {
        return "", false
    }

    text, success := InfographicsTextFromBigInt(bigintValue, options)
    if !success {
        return "", false
    }

    return text, true
}

type UnitToNamePair struct {
    unit *big.Int
    name string
}

func InfographicsTextFromBigInt(
        value *big.Int, options *InfographicsTextOptions) (string, bool) {
    if value.Cmp(big.NewInt(0)) < 0 {
        return "", false
    }

    if value.Cmp(big.NewInt(0)) == 0 {
        if options.short {
            return "0", true
        } else {
            return "零", true
        }
    }

    unitToNameMap := CreateUnitToNameMap()
    var unitToNamePairs []UnitToNamePair
    for unit, name := range unitToNameMap {
        unitToNamePairs = append(unitToNamePairs, UnitToNamePair{unit, name})
    }
    sort.Slice(unitToNamePairs, func(i, j int) bool {
        return !(unitToNamePairs[i].unit.Cmp(unitToNamePairs[j].unit) < 0)
    })

    var texts []string
    for _, pair := range unitToNamePairs {
        quotient, remainder := new(big.Int).DivMod(value, pair.unit, new(big.Int))
        quotientInt := int(quotient.Int64())

        if options.short {
            if quotientInt > 0 {
                name := pair.name
                name = strings.ReplaceAll(name, "千", "000")
                name = strings.ReplaceAll(name, "百", "00")
                name = strings.ReplaceAll(name, "十", "0")
                name = strings.ReplaceAll(name, "一", "")
                texts = append(texts, fmt.Sprintf("%d%s", quotientInt, name))
            }
        } else {
            for i := 0; i < quotientInt; i++ {
                texts = append(texts, pair.name)
            }
        }
        value = remainder
    }

    return strings.Join(texts, options.delimiter), true
}

func CreateUnitToNameMap() map[*big.Int]string {
    prefixes := []string {
        "一", "十", "百", "千",
    }
    suffixes := []string {
        "", "万", "億", "兆", "京", "垓", "𥝱", "穣", "溝", "澗", "正",
        "載", "極", "恒河沙", "阿僧祇", "那由他", "不可思議", "無量大数",
    }
    names := MapSlice(
        ProductSlices(suffixes, prefixes),
        func(pair []string) string {
            return pair[1] + pair[0]
        })

    unitToNameMap := make(map[*big.Int]string)
    for i, name := range names {
        base := big.NewInt(10)
        exponent := big.NewInt(int64(i))
        unit := new(big.Int).Exp(base, exponent, nil)
        unitToNameMap[unit] = name
    }
    return unitToNameMap
}

