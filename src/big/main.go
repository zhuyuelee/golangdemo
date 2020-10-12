package main

import (
	"fmt"
	"math/big"
)

func main() {
	maxMoney := big.NewInt(10)
	minMoney := big.NewInt(12)

	total := maxMoney.Add(maxMoney, minMoney)

	fmt.Printf("maxMoney %d, type %T; minMoney %d, type %T\n", maxMoney, minMoney, maxMoney, minMoney)
	fmt.Printf("totalMoney %d type %T", total, total)
}
