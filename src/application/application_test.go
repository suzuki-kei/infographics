package application

import (
    "src/assert"
    "src/infographics"
    "testing"
)

func TestParseArguments(t *testing.T) {
    type TestCase struct {
        arguments []string
        expectedValues []string
        expectedOptions *Options
    }
    testCases := []TestCase{
        // 何も指定しない場合.
        {
            []string{},
            []string{},
            &Options{false, " ", infographics.ChineseNumeral},
        },
        // 値だけを指定する場合.
        {
            []string{"123"},
            []string{"123"},
            &Options{false, " ", infographics.ChineseNumeral},
        },
        // 各オプションを 1 つ指定する場合.
        {
            []string{"-s", "123"},
            []string{"123"},
            &Options{true, " ", infographics.ChineseNumeral},
        },
        {
            []string{"--short", "123"},
            []string{"123"},
            &Options{true, " ", infographics.ChineseNumeral},
        },
        {
            []string{"-d", ",", "123"},
            []string{"123"},
            &Options{false, ",", infographics.ChineseNumeral},
        },
        {
            []string{"--delimiter", ",", "123"},
            []string{"123"},
            &Options{false, ",", infographics.ChineseNumeral},
        },
        {
            []string{"--delimiter=,", "123"},
            []string{"123"},
            &Options{false, ",", infographics.ChineseNumeral},
        },
        {
            []string{"--chinese-numeral", "123"},
            []string{"123"},
            &Options{false, " ", infographics.ChineseNumeral},
        },
        // 複数のオプションを同時に指定する場合.
        {
            []string{"123", "-s", "456", "-d", ",", "789"},
            []string{"123", "456", "789"},
            &Options{true, ",", infographics.ChineseNumeral},
        },
    }
    for _, testCase := range testCases {
        values, options := parseArguments(testCase.arguments)
        assert.DeepEquals(t, testCase.expectedValues, values)
        assert.DeepEquals(t, testCase.expectedOptions, options)
    }
}

