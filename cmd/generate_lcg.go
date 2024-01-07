package main

import (
	"fmt"
	"github.com/joniumGit/lcg-generator/primes"
	"github.com/joniumGit/lcg-generator/search"
	"os"
	"strconv"
)

func main() {
	var target int

	if len(os.Args) > 1 {
		t, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
		target = t
	} else {
		target = 10_000
	}

	parameters, err := search.FindFixedPeriodParameters(target)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Factors: %d\n", primes.PrimeFactors(target))
	fmt.Printf("Params 1:\n a: %d\n c: %d\n m: %d\n", parameters.A, parameters.C, parameters.M)

	err = search.FixedPeriodAdjustA(&parameters, parameters.A)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Params 2:\n a: %d\n c: %d\n m: %d\n", parameters.A, parameters.C, parameters.M)
}
