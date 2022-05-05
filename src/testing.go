package main

import (
    "reflect"
    "strings"
    "testing"
)

/**
 *
 * エラーが発生しなかったことを確認する.
 *
 */
func AssertSuccess(t *testing.T, err error) {
    if err == nil {
        return
    }
    errorf(t, "error occured.", nil, err)
}

/**
 *
 * 二つの値が == を用いて等しいことを確認する.
 *
 */
func AssertEquals[T comparable](t *testing.T, expected T, actual T) {
    if expected == actual {
        return
    }
    errorf(t, "not equal.", expected, actual)
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
    errorf(t, "not deep equal.", expected, actual)
}

/**
 *
 * testing.T.Errorf() を用いて検証に失敗したことを報告する.
 *
 */
func errorf(t *testing.T, description string, expected any, actual any) {
    format := strings.Join([]string{
        description,
        "expected = %v",
        "actual = %v",
    }, "\n")
    t.Errorf(format, expected, actual)
}

