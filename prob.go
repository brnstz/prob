package prob

import "math/big"

// ReplaceOrdered computes the possible number of uniquely ordered outcomes
// when k choices are made from n possibilities, allowing for duplicate
// values from n. Mathmetically, this is: n^k
func ReplaceOrdered(n, k *big.Int) *big.Int {
	result := big.NewInt(0)

	result.Exp(n, k, nil)

	return result
}

// NoReplaceOrdered computes the possible number of ordered outcomes when k
// choices are made from n possibilities, where values in n may only be chosen
// once. Mathematically this is: n * (n - 1) * ... (n - k + 1)
func NoReplaceOrdered(n, k *big.Int) *big.Int {
	var result big.Int
	var multiplier big.Int
	var lastvalue big.Int

	one := big.NewInt(1)

	// Initialize result to n. This is effectively the first iteration.
	result.Set(n)

	// Initialize multiplier to n - 1. This will be the second iteration,
	// assuming that n is large enough.
	multiplier.Sub(n, one)

	// Initialize lastvalue to n - k + 1
	lastvalue.Sub(n, k)
	lastvalue.Add(&lastvalue, one)

	// Multiple values until multiplier is less than (n - k + 1)
	for {
		if multiplier.Cmp(&lastvalue) == -1 {
			break
		}

		// Multiply the result by the multiplier
		result.Mul(&result, &multiplier)

		// Decrement the multiplier by 1
		multiplier.Sub(&multiplier, one)
	}

	return &result
}
