package application

import (
    "src/assert"
    "src/infographics"
    "testing"
)

func newOptions(short bool, delimiter string) *infographics.Options {
    options := infographics.NewOptions()
    options.Short = short
    options.Delimiter = delimiter
    return options
}

func TestParseArguments(t *testing.T) {
    type TestCase struct {
        arguments []string
        expectedValues []string
        expectedOptions *infographics.Options
    }
    testCases := []TestCase{
        {[]string{}, []string{}, newOptions(false, " ")},
        {[]string{"123"}, []string{"123"}, newOptions(false, " ")},
        {[]string{"-s", "123"}, []string{"123"}, newOptions(true, " ")},
        {[]string{"--short", "123"}, []string{"123"}, newOptions(true, " ")},
        {[]string{"-d", ",", "123"}, []string{"123"}, newOptions(false, ",")},
        {[]string{"--delimiter", ",", "123"}, []string{"123"}, newOptions(false, ",")},
        {[]string{"--delimiter=,", "123"}, []string{"123"}, newOptions(false, ",")},
        {[]string{"123", "-s", "456", "-d", ",", "789"}, []string{"123", "456", "789"}, newOptions(true, ",")},
    }
    for _, testCase := range testCases {
        values, options := parseArguments(testCase.arguments)
        assert.DeepEquals(t, testCase.expectedValues, values)
        assert.DeepEquals(t, testCase.expectedOptions, options)
    }
}

