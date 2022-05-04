package main

/**
 *
 * スライスをコピーする.
 *
 */
func CopySlice[T any](slice []T) []T {
    return append([]T{}, slice...)
}

/**
 *
 * スライスを変換する.
 *
 */
func MapSlice[From any, To any](slice []From, mapper func(From) To) []To {
    values := make([]To, len(slice))

    for i, value := range slice {
        values[i] = mapper(value)
    }
    return values
}

/**
 *
 * スライスの直積を生成し, 戻り値で返す.
 *
 */
func ProductSlices[T any](slices ...[]T) [][]T {
    product := [][]T {}

    ProductSlicesCallback(
        func(values []T) {
            product = append(product, values)
        },
        slices...)
    return product
}

/**
 *
 * スライスの直積を生成し, コールバックを呼び出す.
 *
 */
func ProductSlicesCallback[T any](callback func([]T), slices ...[]T) {
    productSlicesCallback_Loop(slices, callback)

    // NOTE: ループ版と再帰版の両方を実装したかっただけ.
    // productSlicesCallback_Tailrec([]T {}, slices, callback)
}

// ProductSlicesCallback のループによる実装.
func productSlicesCallback_Loop[T any](slices [][]T, callback func([]T)) {
    i := 0
    indices := make([]int, len(slices))
    leadingValues := make([]T, len(slices))

    for {
        if i >= len(slices) {
            callback(CopySlice(leadingValues))
            i--
            indices[i]++
            continue
        }
        if indices[i] >= len(slices[i]) {
            if i == 0 {
                break
            }
            indices[i] = 0
            i--
            indices[i]++
            continue
        }
        leadingValues[i] = slices[i][indices[i]]
        i++
    }
}

// ProductSlicesCallback の再帰による実装.
func productSlicesCallback_Tailrec[T any](leadingValues []T, slices [][]T, callback func([]T)) {
    if len(slices) == 0 {
        callback(leadingValues)
        return
    }
    for _, value := range slices[0] {
        productSlicesCallback_Tailrec(
            append(leadingValues, value),
            slices[1:],
            callback)
    }
}

