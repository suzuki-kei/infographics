package main

import (
    "fmt"
    "math/big"
    "sort"
    "strings"
)

/**
 *
 * インフォグラフィック文字列を生成するときのオプション.
 *
 */
type InfographicsTextOptions struct {
    // true の場合は短縮表現を生成する
    short bool

    // 区切り文字
    delimiter string
}

/**
 *
 * デフォルト値で初期化した InfographicsTextOptions を生成する.
 *
 */
func NewInfographicsTextOptions() *InfographicsTextOptions {
    options := new(InfographicsTextOptions)
    options.short = false
    options.delimiter = " "
    return options
}

/**
 *
 * インフォグラフィック文字列を生成する.
 *
 */
func InfographicsTextFromString(
        value string, options *InfographicsTextOptions) (string, error) {
    bigintValue, err := BigIntFromString(value)
    if err != nil {
        return "", err
    }

    return infographicsTextFromBigInt(bigintValue, options)
}

/**
 *
 * 単位と名称 (Ex. {10000, "万"}).
 *
 */
type UnitToNamePair struct {
    // 単位 (Ex. 10000)
    unit *big.Int

    // 名前 (Ex. "万")
    name string
}

/**
 *
 * big.Int からインフォグラフィック文字列を生成する.
 *
 */
func infographicsTextFromBigInt(
        value *big.Int, options *InfographicsTextOptions) (string, error) {
    if value.Cmp(big.NewInt(0)) < 0 {
        return "", fmt.Errorf("must be (value >= 0): %v", value)
    }

    if options.short {
        text := infographicsShortTextFromBigInt(value, options.delimiter)
        return text, nil
    } else {
        text := infographicsLongTextFromBigInt(value, options.delimiter)
        return text, nil
    }
}

/**
 *
 * big.Int からインフォグラフィック文字列の長いバージョンを生成する.
 *
 * TODO value が負数の場合を考慮する.
 *
 */
func infographicsLongTextFromBigInt(value *big.Int, delimiter string) string {
    if value.Cmp(big.NewInt(0)) == 0 {
        return "零"
    }

    unitToNameMap := createUnitToNameMap()
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
        texts = append(texts, Repeat(pair.name, quotientInt)...)
        value = remainder
    }

    return strings.Join(texts, delimiter)
}

/**
 *
 * big.Int からインフォグラフィック文字列の短いバージョンを生成する.
 *
 * TODO value が負数の場合を考慮する.
 *
 */
func infographicsShortTextFromBigInt(value *big.Int, delimiter string) string {
    if value.Cmp(big.NewInt(0)) == 0 {
        return "0"
    }

    unitToNameMap := createUnitToNameMap()
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

        if quotientInt > 0 {
            name := pair.name
            name = strings.ReplaceAll(name, "千", "000")
            name = strings.ReplaceAll(name, "百", "00")
            name = strings.ReplaceAll(name, "十", "0")
            name = strings.ReplaceAll(name, "一", "")
            texts = append(texts, fmt.Sprintf("%d%s", quotientInt, name))
        }
        value = remainder
    }

    return strings.Join(texts, delimiter)
}

/**
 *
 * 単位と名前のマッピングを生成する.
 *
 */
func createUnitToNameMap() map[*big.Int]string {
    prefixes := []string{
        "一", "十", "百", "千",
    }
    suffixes := []string{
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

