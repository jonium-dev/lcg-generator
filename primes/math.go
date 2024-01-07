package primes

func PrimeFactors(value int) (factors []int) {
	for _, prime := range PRIMES {
		if value%prime == 0 {
			factors = append(factors, prime)
		}
	}
	return factors
}
