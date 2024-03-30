package bigints

import (
    "testing"
)

import (
    "localhost/assert"
)

func TestFromString(t *testing.T) {
    stringValues := []string{
        "0",
        "1",
        "4294967296", // 2^32
        "18446744073709551616", // 2^64
        "1267650600228229401496703205376", // 2^100
    }
    for _, stringValue := range stringValues {
        bigintValue, err := FromString(stringValue)
        assert.Success(t, err)
        assert.Equals(t, stringValue, bigintValue.String())
    }
}

