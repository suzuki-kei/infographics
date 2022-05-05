package main

import (
    "testing"
)

func TestCopySlice(t *testing.T) {
    slice := []int {1, 2, 3}

    copiedSlice := CopySlice(slice)
    assertDeepEquals(t, []int {1, 2, 3}, copiedSlice)

    slice[0] = 10
    slice[1] = 20
    slice[2] = 30
    assertDeepEquals(t, []int {1, 2, 3}, copiedSlice)
}

func TestMapSlice(t *testing.T) {
    values := []int {1, 2, 3}
    mapper := func(x int) int { return x + 1 }
    mappedValues := MapSlice(values, mapper)
    expectedValues := []int {2, 3, 4}
    assertDeepEquals(t, expectedValues, mappedValues)
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
    assertDeepEquals(t, expected, actual)
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
    assertDeepEquals(t, expected, product)
}

