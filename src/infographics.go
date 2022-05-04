package main

import (
    "math/big"
    "sort"
    "strings"
)

func InfographicsTextFromString(value string, delimiter string) (string, bool) {
    bigintValue, success := BigIntFromString(value)
    if !success {
        return "", false
    }

    infographicsText, success := InfographicsTextFromBigInt(bigintValue, delimiter)
    if !success {
        return "", false
    }

    return infographicsText, true
}

type UnitToNamePair struct {
    unit *big.Int
    name string
}

func InfographicsTextFromBigInt(value *big.Int, delimiter string) (string, bool) {
    if value.Cmp(big.NewInt(0)) < 0 {
        return "", false
    }
    if value.Cmp(big.NewInt(0)) == 0 {
        return "零", true
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
        for i := 0; i < int(quotient.Int64()); i++ {
            texts = append(texts, pair.name)
        }
        value = remainder
    }

    return strings.Join(texts, delimiter), true
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
        ProductSlices(prefixes, suffixes),
        func(pair []string) string {
            return strings.Join(pair, "")
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

func BigIntFromString(value string) (*big.Int, bool) {
    return new(big.Int).SetString(value, 10)
}

