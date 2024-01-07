package primes

import "math"

type PrimeMachine struct {
	Memory []int
}

func IsPrime(value int, checkStartingFrom int) bool {
	for i := checkStartingFrom; i <= int(math.Sqrt(float64(value))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return true
}

func (machine *PrimeMachine) LargestSeenPrime() int {
	if len(machine.Memory) > 0 {
		return machine.Memory[len(machine.Memory)-1]
	} else {
		return 2
	}
}

func (machine *PrimeMachine) isDivisibleByMemoizedPrimes(value int) bool {
	for _, prime := range machine.Memory {
		if value%prime == 0 {
			return true
		}
	}
	return false
}

func (machine *PrimeMachine) IsPrime(value int) bool {
	result := !machine.isDivisibleByMemoizedPrimes(value) && IsPrime(value, machine.LargestSeenPrime())
	if result {
		machine.Memory = append(machine.Memory, value)
	}
	return result
}

func (machine *PrimeMachine) Next() int {
	for i := machine.LargestSeenPrime(); !machine.IsPrime(i); i++ {
	}
	return machine.LargestSeenPrime()
}
