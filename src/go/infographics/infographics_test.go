package infographics

import (
    "testing"
)

import (
    "localhost/assert"
)

func TestGenerate(t *testing.T) {
    type TestCase struct {
        value string
        short bool
        separator string
        expected string
    }
    testCases := []TestCase{
        {"0", true, " ", "0"},
        {"0", false, " ", "零"},
        {"1", true, " ", "1"},
        {"1", false, " ", "一"},
        {"100000", true, " ", "10万"},
        {"100000", false, " ", "十万"},
        {"7050301", true, " ", "700万 5万 300 1"},
        {"7050301", false, " ", "百万 百万 百万 百万 百万 百万 百万 一万 一万 一万 一万 一万 百 百 百 一"},
    }
    for _, testCase := range testCases {
        value := testCase.value
        short := testCase.short
        separator := testCase.separator
        systemOfUnit := ChineseNumeral
        expected := testCase.expected

        text, err := Generate(value, short, separator, systemOfUnit)
        assert.Success(t, err)
        assert.Equals(t, expected, text)
    }
}

