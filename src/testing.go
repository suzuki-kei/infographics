package main

import (
    "reflect"
    "strings"
    "testing"
)

func assertEquals[T comparable](t *testing.T, expected T, actual T) {
    if expected == actual {
        return
    }
    format := strings.Join([]string{
        "not equal.",
        "expected = %v",
        "actual= %v",
    }, "\n")
    t.Errorf(format, expected, actual)
}

func assertDeepEquals[T any](t *testing.T, expected T, actual T) {
    if reflect.DeepEqual(expected, actual) {
        return
    }
    format := strings.Join([]string{
        "not deep equal.",
        "expected = %v",
        "actual= %v",
    }, "\n")
    t.Errorf(format, expected, actual)
}

