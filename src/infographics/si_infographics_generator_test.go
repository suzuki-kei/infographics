package infographics

import (
    "src/assert"
    "src/bigints"
    "testing"
)

func TestSiInfographicsGeneratorGenerate(t *testing.T) {
    type TestCase struct {
        value string
        short bool
        delimiter string
        expected string
    }
    testCases := []TestCase{
        {"0", true, " ", "0"},
        {"0", false, " ", "0"},
        {"1", true, " ", "1"},
        {"1", false, " ", "1"},
        {"100000", true, " ", "100k"},
        {"100000", false, " ", "k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k"},
        {"7050301", true, " ", "7M 50k 3h 1"},
        {"7050301", false, " ", "M M M M M M M k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k k h h h 1"},
    }
    for _, testCase := range testCases {
        generator := SiInfographicsGenerator{
            short: testCase.short,
            delimiter: testCase.delimiter,
        }
        value, _ := bigints.FromString(testCase.value)
        text, err := generator.Generate(value)
        assert.Success(t, err)
        assert.Equals(t, testCase.expected, text)
    }
}

