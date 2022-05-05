package main

import (
    "testing"
)

func TestBigIntFromString(t *testing.T) {
    testCases := []string {
        "0",
        "1",
        "4294967296", // 2^32
        "18446744073709551616", // 2^64
        "1267650600228229401496703205376", // 2^100
    }
    for _, stringValue := range testCases {
        bigintValue, success := BigIntFromString(stringValue)
        assertEquals(t, true, success)
        assertEquals(t, stringValue, bigintValue.String())
    }
}

