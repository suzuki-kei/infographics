package infographics

import (
    "fmt"
    "math/big"
    "sort"
    "src/slices"
    "strings"
)

/**
 *
 * 漢数字を用いてインフォグラフィックを生成する.
 *
 */
type ChineseNumeralInfographicsGenerator struct {

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
func (this ChineseNumeralInfographicsGenerator) Generate(value *big.Int) (string, error) {
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
func (this ChineseNumeralInfographicsGenerator) generateLongText(value *big.Int) (string, error) {
    if value.Cmp(big.NewInt(0)) < 0 {
        return "", fmt.Errorf("must be (value >= 0): %v", value)
    }
    if value.Cmp(big.NewInt(0)) == 0 {
        return "零", nil
    }

    numeralUnits := this.createNumeralUnits(
        []string{
            "一", "十", "百", "千",
        },
        []string{
            "", "万", "億", "兆", "京", "垓", "𥝱", "穣", "溝", "澗", "正",
            "載", "極", "恒河沙", "阿僧祇", "那由他", "不可思議", "無量大数",
        },
    )

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
func (this ChineseNumeralInfographicsGenerator) generateShortText(value *big.Int) (string, error) {
    if value.Cmp(big.NewInt(0)) < 0 {
        return "", fmt.Errorf("must be (value >= 0): %v", value)
    }
    if value.Cmp(big.NewInt(0)) == 0 {
        return "0", nil
    }

    numeralUnits := this.createNumeralUnits(
        []string{
            "", "0", "00", "000",
        },
        []string{
            "", "万", "億", "兆", "京", "垓", "𥝱", "穣", "溝", "澗", "正",
            "載", "極", "恒河沙", "阿僧祇", "那由他", "不可思議", "無量大数",
        },
    )

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
 * 単位と名前 (Ex. {10000, "万"}).
 *
 */
type NumeralUnit struct {
    // 単位 (Ex. 10000)
    unit *big.Int

    // 名前 (Ex. "万")
    name string
}

/**
 *
 * prefixes と suffixes から NumeralUnit のスライスを生成する.
 *
 * 戻り値は NumeralUnit.unit の降順に整列されている.
 *
 */
func (this ChineseNumeralInfographicsGenerator) createNumeralUnits(prefixes []string, suffixes []string) []NumeralUnit {
    var names []string
    {
        for _, suffix := range suffixes {
            for _, prefix := range prefixes {
                name := fmt.Sprintf("%s%s", prefix, suffix)
                names = append(names, name)
            }
        }
    }

    createUnit := func(exponent int) *big.Int {
        bigintBase := big.NewInt(int64(10))
        bigintExponent := big.NewInt(int64(exponent))
        return new(big.Int).Exp(bigintBase, bigintExponent, nil)
    }

    var numeralUnits []NumeralUnit
    {
        for i, name := range names {
            unit := createUnit(i)
            numeralUnit := NumeralUnit{unit, name}
            numeralUnits = append(numeralUnits, numeralUnit)
        }
        sort.Slice(numeralUnits, func(i, j int) bool {
            return !(numeralUnits[i].unit.Cmp(numeralUnits[j].unit) < 0)
        })
    }

    return numeralUnits
}

