package bigints

import (
    "fmt"
    "math/big"
)

/**
 *
 * 数値文字列を big.Int に変換する.
 *
 */
func FromString(value string) (*big.Int, error) {
    bigintValue, success := new(big.Int).SetString(value, 10)
    if !success {
        return nil, fmt.Errorf("can not convert to big.Int: %v", value)
    }
    return bigintValue, nil
}

