package infographics

import (
    "fmt"
    "math/big"
    "sort"
    "src/bigints"
    "src/slices"
    "strings"
)

/**
 *
 * SI 単位系を用いてインフォグラフィックを生成する.
 *
 */
type SiInfographicsGenerator struct {

    // true の場合は短縮表現を生成する.
    short bool

    // 区切り文字.
    delimiter string

}

/**
 *
 * インフォグラフィック文字列を生成する.
 *
 */
func (this SiInfographicsGenerator) Generate(value *big.Int) (string, error) {
    if this.short {
        return this.generateShortText(value)
    } else {
        return this.generateLongText(value)
    }
}

/**
 *
 * 長いバージョンのインフォグラフィック文字列を生成する.
 *
 */
func (this SiInfographicsGenerator) generateLongText(value *big.Int) (string, error) {
    if value.Cmp(big.NewInt(0)) < 0 {
        return "", fmt.Errorf("must be (value >= 0): %v", value)
    }
    if value.Cmp(big.NewInt(0)) == 0 {
        return "0", nil
    }

    numeralUnits := this.createNumeralUnits([][]string{
        {"1", "1"},
        {"10", "da"},                       // deca (10^1)
        {"100", "h"},                       // hecto (10^2)
        {"1000", "k"},                      // kilo (10^3)
        {"1000000", "M"},                   // mega (10^6)
        {"1000000000", "G"},                // giga (10^9)
        {"1000000000000", "T"},             // tera (10^12)
        {"1000000000000000", "P"},          // peta (10^15)
        {"1000000000000000000", "E"},       // exa (10^18)
        {"1000000000000000000000", "Z"},    // zetta (10^21)
        {"1000000000000000000000000", "Y"}, // yotta (10^24)
    })

    var texts []string
    {
        for _, numeralUnit := range numeralUnits {
            quotient, remainder := new(big.Int).DivMod(value, numeralUnit.unit, new(big.Int))
            quotientInt := int(quotient.Int64())
            texts = append(texts, slices.Repeat(numeralUnit.name, quotientInt)...)
            value = remainder
        }
    }

    return strings.Join(texts, this.delimiter), nil
}

/**
 *
 * 短いバージョンのインフォグラフィック文字列を生成する.
 *
 */
func (this SiInfographicsGenerator) generateShortText(value *big.Int) (string, error) {
    if value.Cmp(big.NewInt(0)) < 0 {
        return "", fmt.Errorf("must be (value >= 0): %v", value)
    }
    if value.Cmp(big.NewInt(0)) == 0 {
        return "0", nil
    }

    numeralUnits := this.createNumeralUnits([][]string{
        {"1", ""},
        {"10", "da"},                       // deca (10^1)
        {"100", "h"},                       // hecto (10^2)
        {"1000", "k"},                      // kilo (10^3)
        {"1000000", "M"},                   // mega (10^6)
        {"1000000000", "G"},                // giga (10^9)
        {"1000000000000", "T"},             // tera (10^12)
        {"1000000000000000", "P"},          // peta (10^15)
        {"1000000000000000000", "E"},       // exa (10^18)
        {"1000000000000000000000", "Z"},    // zetta (10^21)
        {"1000000000000000000000000", "Y"}, // yotta (10^24)
    })

    var texts []string
    {
        for _, numeralUnit := range numeralUnits {
            quotient, remainder := new(big.Int).DivMod(value, numeralUnit.unit, new(big.Int))
            quotientInt := int(quotient.Int64())

            if quotientInt > 0 {
                texts = append(texts, fmt.Sprintf("%d%s", quotientInt, numeralUnit.name))
            }
            value = remainder
        }
    }

    return strings.Join(texts, this.delimiter), nil
}

/**
 *
 * prefixes と suffixes から NumeralUnit のスライスを生成する.
 *
 * 戻り値は NumeralUnit.unit の降順に整列されている.
 *
 */
func (this SiInfographicsGenerator) createNumeralUnits(unitAndNames [][]string) []NumeralUnit {
    var numeralUnits []NumeralUnit
    {
        for _, unitAndName := range unitAndNames {
            unit, _ := bigints.FromString(unitAndName[0]) // TODO エラーハンドリング
            name := unitAndName[1]
            numeralUnit := NumeralUnit{unit, name}
            numeralUnits = append(numeralUnits, numeralUnit)
        }
        sort.Slice(numeralUnits, func(i, j int) bool {
            return !(numeralUnits[i].unit.Cmp(numeralUnits[j].unit) < 0)
        })
    }

    return numeralUnits
}

