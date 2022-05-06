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
func TextFromString(value string, options *Options) (string, error) {
    bigintValue, err := bigints.FromString(value)
    if err != nil {
        return "", err
    }

    return textFromBigInt(bigintValue, options)
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
 * big.Int からインフォグラフィック文字列を生成する.
 *
 */
func textFromBigInt(value *big.Int, options *Options) (string, error) {
    if options.Short {
        return shortTextFromBigInt(value, options.Delimiter)
    } else {
        return longTextFromBigInt(value, options.Delimiter)
    }
}

/**
 *
 * big.Int からインフォグラフィック文字列の長いバージョンを生成する.
 *
 */
func longTextFromBigInt(value *big.Int, delimiter string) (string, error) {
    if value.Cmp(big.NewInt(0)) < 0 {
        return "", fmt.Errorf("must be (value >= 0): %v", value)
    }
    if value.Cmp(big.NewInt(0)) == 0 {
        return "零", nil
    }

    numeralUnits := createNumeralUnits(
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

    return strings.Join(texts, delimiter), nil
}

/**
 *
 * big.Int からインフォグラフィック文字列の短いバージョンを生成する.
 *
 */
func shortTextFromBigInt(value *big.Int, delimiter string) (string, error) {
    if value.Cmp(big.NewInt(0)) < 0 {
        return "", fmt.Errorf("must be (value >= 0): %v", value)
    }
    if value.Cmp(big.NewInt(0)) == 0 {
        return "0", nil
    }

    numeralUnits := createNumeralUnits(
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

    return strings.Join(texts, delimiter), nil
}

/**
 *
 * prefixes と suffixes から NumeralUnit のスライスを生成する.
 *
 * 戻り値は NumeralUnit.unit の降順に整列されている.
 *
 */
func createNumeralUnits(prefixes []string, suffixes []string) []NumeralUnit {
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

