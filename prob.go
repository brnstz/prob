package prob

import "math/big"

// ReplaceOrdered computes the number of possible uniquely ordered outcomes
// when k choices are made from n possibilities, allowing for duplicate
// values from n.
func ReplaceOrdered(n, k *big.Int) *big.Int {
	result := big.NewInt(0)

	result.Exp(n, k, nil)

	return result
}
