package prob

import "math/big"

var (
	// In many cases we want to compare with one and zero. Have some private
	// read-only variables for doing that.
	zero = big.NewInt(0)
	one  = big.NewInt(1)
)

// ReplaceOrdered computes the possible number of uniquely ordered outcomes
// when k choices are made from n possibilities, allowing for duplicate
// values from n. Mathmetically, this is: n^k
func ReplaceOrdered(n, k *big.Int) *big.Int {
	var result big.Int

	result.Exp(n, k, nil)

	return &result
}

// NoReplaceOrdered computes the possible number of ordered outcomes when k
// choices are made from n possibilities, where values in n may only be chosen
// once. Mathematically this is: n * (n - 1) * ... (n - k + 1)
func NoReplaceOrdered(n, k *big.Int) *big.Int {
	var result, multiplier, lastvalue big.Int

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

// factorial computes the factorial of n
func factorial(n *big.Int) *big.Int {
	var result, multiplier big.Int

	// special case, factorial of 0 is 1
	if n.Cmp(zero) == 0 {
		return big.NewInt(1)
	}

	// Initialize result to starting value
	result.Set(n)

	// Initialize multiplier to first value to be multiplied, which is
	// n - 1
	multiplier.Sub(&result, one)

	for {
		// If we're less than 1, it's time to stop
		if multiplier.Cmp(one) == -1 {
			break
		}

		result.Mul(&result, &multiplier)

		multiplier.Sub(&multiplier, one)
	}

	return &result
}

// NoReplaceUnordered computes the number of outcomes when k choices
// are made from n possibilities, without regard for order and where
// values in n may only be chosen once. This is the same as the binomial
// coefficient. This computation is also provided by math/big.Int.Binomial() but
// is re-implemented here for fun. Mathemtically, this is:
// n! / (k! * (n - k)!)
func NoReplaceUnordered(n, k *big.Int) *big.Int {
	var result, denom big.Int

	// Start out the denominator with n - k factorial
	denom.Sub(n, k)
	denom.Set(factorial(&denom))

	// Finish the denominator by multiplying k factorial by previous value
	denom.Mul(factorial(k), &denom)

	// Compute final result with division
	result.Div(factorial(n), &denom)

	return &result
}

// ReplaceUnordered computes the number of outcomes when k choices are made
// from n possibilities, without regard for order and allowing for duplicate
// values from n. This is also known as the Bose-Einstein value. Mathemtically,
// this is bionomial(n + k - 1, k)
func ReplaceUnordered(n, k *big.Int) *big.Int {
	var n2 big.Int

	// Set up first part of bionomail as n + k -1
	n2.Add(n, k)
	n2.Sub(&n2, one)

	return NoReplaceUnordered(&n2, k)
}
