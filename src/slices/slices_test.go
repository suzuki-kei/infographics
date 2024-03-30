package slices

import (
    "testing"
)

import (
    "localhost/assert"
)

func TestCopy(t *testing.T) {
    slice := []int{1, 2, 3}

    copiedSlice := Copy(slice)
    assert.DeepEquals(t, []int{1, 2, 3}, copiedSlice)

    slice[0] = 10
    slice[1] = 20
    slice[2] = 30
    assert.DeepEquals(t, []int{1, 2, 3}, copiedSlice)
}

func TestMap(t *testing.T) {
    values := []int{1, 2, 3}
    mapper := func(x int) int { return x + 1 }
    mappedValues := Map(values, mapper)
    expectedValues := []int{2, 3, 4}
    assert.DeepEquals(t, expectedValues, mappedValues)
}

func TestProduct(t *testing.T) {
    actual := Product(
        []string{"A", "B", "C"},
        []string{"1", "2", "3"})
    expected := [][]string{
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
    assert.DeepEquals(t, expected, actual)
}

func TestProductCallback(t *testing.T) {
    product := [][]int{
    }
    callback := func(values []int) {
        product = append(product, values)
    }
    slices := [][]int{
        {1, 2, 3},
        {4, 5, 6},
    }
    expected := [][]int{
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
    ProductCallback(callback, slices...)
    assert.DeepEquals(t, expected, product)
}

func TestRepeat(t *testing.T) {
    type TestCase struct {
        value string
        n int
        expected []string
    }
    testCases := []TestCase{
        // n 要素のスライスが生成される.
        {"A", 1, []string{"A"}},
        {"B", 2, []string{"B", "B"}},
        {"C", 3, []string{"C", "C", "C"}},

        // n <= 0 の場合は空のスライスが生成される.
        {"A", 0, []string{}},
        {"A", -1, []string{}},
    }
    for _, testCase := range testCases {
        value := testCase.value
        n := testCase.n
        expected := testCase.expected
        actual := Repeat(value, n)
        assert.DeepEquals(t, expected, actual)
    }
}

