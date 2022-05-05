package main

import (
    "testing"
)

func TestBigIntFromString(t *testing.T) {
    stringValues := []string{
        "0",
        "1",
        "4294967296", // 2^32
        "18446744073709551616", // 2^64
        "1267650600228229401496703205376", // 2^100
    }
    for _, stringValue := range stringValues {
        bigintValue, err := BigIntFromString(stringValue)
        AssertSuccess(t, err)
        AssertEquals(t, stringValue, bigintValue.String())
    }
}

