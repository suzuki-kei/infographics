package main

import (
    "testing"
)

func TestInfographicsTextFromString(t *testing.T) {
    testCases := [][]string {
        {"0", "零"},
        {"1", "一"},
        {"10", "十"},
        {"12345", "一万 千 千 百 百 百 十 十 十 十 一 一 一 一 一"},
        {"70000", "一万 一万 一万 一万 一万 一万 一万"},
        {"100000", "十万"},
        {"17500000000", "百億 十億 十億 十億 十億 十億 十億 十億 一億 一億 一億 一億 一億"},
    }
    options := NewInfographicsTextOptions()

    for _, testCase := range testCases {
        value := testCase[0]
        expected := testCase[1]
        infographicsText, success := InfographicsTextFromString(value, options)
        assertEquals(t, true, success)
        assertEquals(t, expected, infographicsText)
    }
}

