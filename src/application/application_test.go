package application

import (
    "src/assert"
    "testing"
)

func TestParseArguments(t *testing.T) {
    type TestCase struct {
        arguments []string
        expectedValues []string
        expectedOptions *Options
    }
    testCases := []TestCase{
        {[]string{}, []string{}, &Options{false, " "}},
        {[]string{"123"}, []string{"123"}, &Options{false, " "}},
        {[]string{"-s", "123"}, []string{"123"}, &Options{true, " "}},
        {[]string{"--short", "123"}, []string{"123"}, &Options{true, " "}},
        {[]string{"-d", ",", "123"}, []string{"123"}, &Options{false, ","}},
        {[]string{"--delimiter", ",", "123"}, []string{"123"}, &Options{false, ","}},
        {[]string{"--delimiter=,", "123"}, []string{"123"}, &Options{false, ","}},
        {[]string{"123", "-s", "456", "-d", ",", "789"}, []string{"123", "456", "789"}, &Options{true, ","}},
    }
    for _, testCase := range testCases {
        values, options := parseArguments(testCase.arguments)
        assert.DeepEquals(t, testCase.expectedValues, values)
        assert.DeepEquals(t, testCase.expectedOptions, options)
    }
}

