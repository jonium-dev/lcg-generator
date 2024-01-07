package primes

import "testing"

func TestPrimeFactorsUsesPrimes(t *testing.T) {
	factors := PrimeFactors(10)
	if factors[0] != 2 || factors[1] != 5 {
		t.Fatalf("Failed to produce correct primes (%d)", factors)
	}
}
