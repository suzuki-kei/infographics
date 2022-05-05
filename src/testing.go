package main

import (
    "reflect"
    "strings"
    "testing"
)

/**
 *
 * 二つの値が == を用いて等しいことを確認する.
 *
 */
func AssertEquals[T comparable](t *testing.T, expected T, actual T) {
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

/**
 *
 * 二つの値が reflect.DeepEqual を用いて等しいことを確認する.
 *
 */
func AssertDeepEquals[T any](t *testing.T, expected T, actual T) {
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

