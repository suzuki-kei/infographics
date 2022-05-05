package main

import (
    "testing"
)

type TestCase struct {
    value string
    short bool
    delimiter string
    expected string
}

func TestInfographicsTextFromString(t *testing.T) {
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
        options := NewInfographicsTextOptions()
        options.short = testCase.short
        options.delimiter = testCase.delimiter
        text, success := InfographicsTextFromString(testCase.value, options)
        assertEquals(t, true, success)
        assertEquals(t, testCase.expected, text)
    }
}

