package search

import (
	"fmt"
	"github.com/joniumGit/lcg-generator/generator"
	"github.com/joniumGit/lcg-generator/primes"
)

// FixedPeriodAdjustC Finds the lowest C that is coprime to M for the given parameters and lowerBound
func FixedPeriodAdjustC(parameters *generator.LCG, lowerBound int) error {
	for ; lowerBound < parameters.M; lowerBound++ {
		if GCD(lowerBound, parameters.M) == 1 {
			parameters.C = lowerBound
			return nil
		}
	}
	return fmt.Errorf("failed to find suitable C for lower bound of %d", lowerBound)
}

// FixedPeriodFindSmallestA Finds the smallest value of A that fulfills all criteria
//
// - A-1 needs to be divisible by all prime factors of M
//
// - If M is divisible by 4 the A-1 needs to be divisible by 4
func FixedPeriodFindSmallestA(parameters *generator.LCG) {
	aMinus := 1

	// A-1 needs to be divisible by all prime factors of M
	for _, factor := range primes.PrimeFactors(parameters.M) {
		aMinus *= factor
	}

	// If M is divisible by 4 the A-1 needs to be divisible by 4
	if parameters.M%4 == 0 && aMinus%4 != 0 {
		if aMinus%2 == 0 {
			aMinus *= 2
		} else {
			aMinus *= 4
		}
	}

	// If M was a prime then this will likely be over M so do a modulo
	// This will make A equal to 1 always as 0 < A < M and m !/ 4
	// I leave this here as a failsafe
	parameters.A = (aMinus + 1) % parameters.M
}

// FixedPeriodAdjustA Finds the smallest A for the given parameters and lowerBound
func FixedPeriodAdjustA(parameters *generator.LCG, lowerBound int) error {
	// Initialize to baseline
	FixedPeriodFindSmallestA(parameters)

	if parameters.A > lowerBound {
		return nil
	}

	// Set up a lookup for primes that are already in A
	lookup := map[int]bool{}
	for _, i := range primes.PrimeFactors(parameters.A) {
		lookup[i] = true
	}

	// Let's try first to find a new prime to multiply with
	for _, prime := range primes.PRIMES {
		if !lookup[prime] {
			t := parameters.A * prime
			if t >= parameters.M {
				break
			}
			if t > lowerBound {
				parameters.A = t
				return nil
			}
		}
	}

	// Otherwise just run up the number line
	for i := 2; i*parameters.A < parameters.M; i++ {
		if t := parameters.A * i; t > lowerBound {
			parameters.A = t
			return nil
		}
	}

	return fmt.Errorf("failed to find suitable A for lower bound %d", lowerBound)
}

// FindFixedPeriodParameters Finds parameters for an LCG of the given period
//
// Do note that a prime period will produce an absolutely horrible LCG
// with nearly no randomness at all (A=1, C=2, M=prime).
func FindFixedPeriodParameters(period int) (parameters generator.LCG, err error) {
	parameters.M = period

	err = FixedPeriodAdjustC(&parameters, 2)
	if err != nil {
		return parameters, err
	}

	err = FixedPeriodAdjustA(&parameters, 0)
	if err != nil {
		return parameters, err
	}

	return parameters, nil
}
