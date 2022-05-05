package main

import (
    "reflect"
    "testing"
)

func assertEquals[T comparable](t *testing.T, expected T, actual T) {
    if expected != actual {
        t.Errorf("not equals.\n\texpected = %v\n\tactual = %v", expected, actual)
    }
}

func assertDeepEquals[T any](t *testing.T, expected T, actual T) {
    if !reflect.DeepEqual(expected, actual) {
        t.Errorf("not deep equals.\n\texpected = %v\n\tactual = %v", expected, actual)
    }
}

