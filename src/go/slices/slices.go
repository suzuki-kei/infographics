package slices

/**
 *
 * スライスをコピーする.
 *
 */
func Copy[T any](slice []T) []T {
    return append([]T{}, slice...)
}

/**
 *
 * スライスを変換する.
 *
 */
func Map[From any, To any](slice []From, mapper func(From) To) []To {
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
func Product[T any](slices ...[]T) [][]T {
    product := [][]T{}

    ProductCallback(
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
func ProductCallback[T any](callback func([]T), slices ...[]T) {
    productCallback_Loop(slices, callback)

    // NOTE: ループ版と再帰版の両方を実装したかっただけ.
    // productsCallback_Tailrec([]T{}, slices, callback)
}

// ProductCallback のループによる実装.
func productCallback_Loop[T any](slices [][]T, callback func([]T)) {
    i := 0
    indices := make([]int, len(slices))
    leadingValues := make([]T, len(slices))

    for {
        if i >= len(slices) {
            callback(Copy(leadingValues))
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

// ProductCallback の再帰による実装.
func productCallback_Tailrec[T any](leadingValues []T, slices [][]T, callback func([]T)) {
    if len(slices) == 0 {
        callback(leadingValues)
        return
    }
    for _, value := range slices[0] {
        productCallback_Tailrec(
            append(leadingValues, value),
            slices[1:],
            callback)
    }
}

/**
 *
 * 指定した値 value を n 個持つスライスを生成する.
 *
 */
func Repeat[T any](value T, n int) []T {
    if n < 0 {
        return []T{}
    }

    values := make([]T, n)
    {
        for i := 0; i < n; i++ {
            values[i] = value
        }
    }

    return values
}

