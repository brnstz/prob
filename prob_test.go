package prob_test

import (
	"math/big"
	"testing"

	"github.com/brnstz/prob"
)

func TestReplaceOrdered(t *testing.T) {
	// tests is a list to run with the expectation that
	// prob.ReplaceOrdered(n, k) == result (where n and k are converted to
	// big.Int values and the result is tested as a base10 string)
	tests := []struct {
		n      int64
		k      int64
		result string
	}{
		{5, 4, "625"},
		{10, 2, "100"},
	}

	// Go through each test
	for _, test := range tests {

		// Get the result as a base10 string
		actual := prob.ReplaceOrdered(
			big.NewInt(test.n), big.NewInt(test.k),
		).Text(10)

		// Check result matches or signal fatal error
		if actual != test.result {
			t.Fatal("expected prob.ReplaceOrdered(%d, %d) == %v but got %v",
				test.n, test.k, test.result, actual,
			)
		}
	}
}
