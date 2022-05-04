package main

import (
    "reflect"
    "testing"
)

func TestCopySlice(t *testing.T) {
    values1 := []int {1, 2, 3}
    values2 := CopySlice(values1)

    if !reflect.DeepEqual(values2, []int{1, 2, 3}) {
        t.Errorf("not expected result: %v", values2)
    }

    values2[0] = 10
    values2[1] = 20
    values2[2] = 30

    if !reflect.DeepEqual(values1, []int{1, 2, 3}) {
        t.Errorf("source slice was unintentionally modified: %v", values1)
    }
}

func TestMapSlice(t *testing.T) {
    values := []int {1, 2, 3}
    mapper := func(x int) int { return x + 1 }
    mappedValues := MapSlice(values, mapper)
    expectedValues := []int {2, 3, 4}

    if !reflect.DeepEqual(mappedValues, expectedValues) {
        t.Errorf("not expected result: %v", mappedValues)
    }
}

func TestProductSlices(t *testing.T) {
    actual := ProductSlices(
        []string {"A", "B", "C"},
        []string {"1", "2", "3"})
    expected := [][]string {
        {"A", "1"},
        {"A", "2"},
        {"A", "3"},
        {"B", "1"},
        {"B", "2"},
        {"B", "3"},
        {"C", "1"},
        {"C", "2"},
        {"C", "3"},
    }
    if !reflect.DeepEqual(actual, expected) {
        t.Errorf("not expected result: %v", actual)
    }
}

func TestProductSlicesCallback(t *testing.T) {
    product := [][]int {
    }
    callback := func(values []int) {
        product = append(product, values)
    }
    slices := [][]int {
        {1, 2, 3},
        {4, 5, 6},
    }
    expected := [][]int {
        {1, 4},
        {1, 5},
        {1, 6},
        {2, 4},
        {2, 5},
        {2, 6},
        {3, 4},
        {3, 5},
        {3, 6},
    }

    ProductSlicesCallback(callback, slices...)

    if !reflect.DeepEqual(product, expected) {
        t.Errorf("not expected result: %v", product)
    }
}

