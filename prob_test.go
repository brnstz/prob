package prob_test

import (
	"math/big"
	"testing"

	"github.com/brnstz/prob"
)

type test struct {
	n      int64
	k      int64
	result string
}

type Tester func(n, k *big.Int) *big.Int

func runTests(t *testing.T, tests []test, tester Tester, name string) {
	for _, test := range tests {
		// Get the result as a base10 string
		actual := tester(
			big.NewInt(test.n), big.NewInt(test.k),
		).Text(10)

		// Check result matches or signal fatal error
		if actual != test.result {
			t.Fatalf("expected %s(%d, %d) == %v but got %v",
				name, test.n, test.k, test.result, actual,
			)
		}
	}
}

func TestReplaceOrdered(t *testing.T) {
	// tests is a list to run with the expectation that
	// prob.ReplaceOrdered(n, k) == result (where n and k are converted to
	// big.Int values and the result is tested as a base10 string)
	tests := []test{
		{5, 4, "625"},
		{10, 2, "100"},
		{0, 21, "0"},
	}

	runTests(t, tests, prob.ReplaceOrdered, "prob.ReplaceOrdered")
}

func TestNoReplaceOrdered(t *testing.T) {
	// tests is a list to run with the expectation that
	// prob.NoReplaceOrdered(n, k) == result (where n and k are converted to
	// big.Int values and the result is tested as a base10 string)
	tests := []test{
		{100, 3, "970200"},
		{5, 6, "0"},
	}

	runTests(t, tests, prob.NoReplaceOrdered, "prob.NoReplaceOrdered")
}

func TestNoReplaceUnordered(t *testing.T) {

}
