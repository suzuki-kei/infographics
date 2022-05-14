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
 * IEC 単位系を用いてインフォグラフィックを生成する.
 *
 */
type IecInfographicsGenerator struct {

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
func (this IecInfographicsGenerator) Generate(value *big.Int) (string, error) {
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
func (this IecInfographicsGenerator) generateLongText(value *big.Int) (string, error) {
    if value.Cmp(big.NewInt(0)) < 0 {
        return "", fmt.Errorf("must be (value >= 0): %v", value)
    }
    if value.Cmp(big.NewInt(0)) == 0 {
        return "0", nil
    }

    numeralUnits := this.createNumeralUnits([][]string{
        {"1", "1"},
        {"1024", "Ki"},                         // kibi (2^10)
        {"1048576", "Mi"},                      // mebi (2^20)
        {"1073741824", "Gi"},                   // gibi (2^30)
        {"1099511627776", "Ti"},                // tebi (2^40)
        {"1125899906842624", "Pi"},             // pebi (2^50)
        {"1152921504606846976", "Ei"},          // exbi (2^60)
        {"1180591620717411303424", "Zi"},       // zebi (2^70)
        {"1208925819614629174706176", "Yi"},    // yobi (2^80)
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
func (this IecInfographicsGenerator) generateShortText(value *big.Int) (string, error) {
    if value.Cmp(big.NewInt(0)) < 0 {
        return "", fmt.Errorf("must be (value >= 0): %v", value)
    }
    if value.Cmp(big.NewInt(0)) == 0 {
        return "0", nil
    }

    numeralUnits := this.createNumeralUnits([][]string{
        {"1", ""},
        {"1024", "Ki"},                         // kibi (2^10)
        {"1048576", "Mi"},                      // mebi (2^20)
        {"1073741824", "Gi"},                   // gibi (2^30)
        {"1099511627776", "Ti"},                // tebi (2^40)
        {"1125899906842624", "Pi"},             // pebi (2^50)
        {"1152921504606846976", "Ei"},          // exbi (2^60)
        {"1180591620717411303424", "Zi"},       // zebi (2^70)
        {"1208925819614629174706176", "Yi"},    // yobi (2^80)
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
func (this IecInfographicsGenerator) createNumeralUnits(unitAndNames [][]string) []NumeralUnit {
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

