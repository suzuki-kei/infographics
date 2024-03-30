package infographics

import (
    "testing"
)

import (
    "localhost/assert"
    "localhost/bigints"
)

func TestIecInfographicsGeneratorGenerate(t *testing.T) {
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
        {"10240", true, " ", "10Ki"},
        {"10240", false, " ", "Ki Ki Ki Ki Ki Ki Ki Ki Ki Ki"},
        {"5245953", true, " ", "5Mi 3Ki 1"},
        {"5245953", false, " ", "Mi Mi Mi Mi Mi Ki Ki Ki 1"},
    }
    for _, testCase := range testCases {
        generator := IecInfographicsGenerator{
            short: testCase.short,
            delimiter: testCase.delimiter,
        }
        value, _ := bigints.FromString(testCase.value)
        text, err := generator.Generate(value)
        assert.Success(t, err)
        assert.Equals(t, testCase.expected, text)
    }
}

