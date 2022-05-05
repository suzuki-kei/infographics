package main

import (
    "math/big"
)

/**
 *
 * 数値文字列を big.Int に変換し, (value, success) を返す.
 *
 * 変換に失敗した場合は success が false となる.
 *
 */
func BigIntFromString(value string) (*big.Int, bool) {
    return new(big.Int).SetString(value, 10)
}

