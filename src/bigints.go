package main

import (
    "math/big"
)

func BigIntFromString(value string) (*big.Int, bool) {
    return new(big.Int).SetString(value, 10)
}

