package primes

import (
	"testing"
)

func Test277IsPrime(t *testing.T) {
	value := 277
	prime := IsPrime(value, 2)
	if !prime {
		t.Fatalf(`prime was not a prime (%d)`, value)
	}
}

func Test1000IsNotPrime(t *testing.T) {
	value := 1000
	prime := IsPrime(value, 2)
	if prime {
		t.Fatalf(`Non-prime detected as a prime (%d)`, value)
	}
}

func Test277IsPrimeInMachine(t *testing.T) {
	value := 277
	prime := IsPrime(value, 2)
	if !prime {
		t.Fatalf(`prime was not a prime (%d)`, value)
	}
}

func Test1000IsNotPrimeInMachineWithoutExistingPrime(t *testing.T) {
	value := 1000
	machine := PrimeMachine{Memory: []int{}}
	prime := machine.IsPrime(value)
	if prime {
		t.Fatalf(`Non-prime detected as a prime (%d)`, value)
	}
}

func Test1000IsNotPrimeInMachineWithExistingPrime(t *testing.T) {
	value := 1000
	machine := PrimeMachine{Memory: []int{2}}
	prime := machine.IsPrime(value)
	if prime {
		t.Fatalf(`Non-prime detected as a prime (%d)`, value)
	}
}

func TestMachineLargestPrimeDefaultsTo2(t *testing.T) {
	machine := PrimeMachine{Memory: []int{}}
	if machine.LargestSeenPrime() != 2 {
		t.Fatalf("Wrong default for largest prime (%d)", machine.LargestSeenPrime())
	}
}

func TestMachineLargestPrimeReturnsLast(t *testing.T) {
	machine := PrimeMachine{Memory: []int{2, 3, 5}}
	if machine.LargestSeenPrime() != 5 {
		t.Fatalf("Wrong element for largest prime (%d)", machine.LargestSeenPrime())
	}
}

func TestMachineMemoizesPrime(t *testing.T) {
	value := 277
	machine := PrimeMachine{Memory: []int{2}}
	if prime := machine.IsPrime(value); !prime {
		t.Fatalf(`prime was not a prime (%d)`, value)
	}
	if memoryLength := len(machine.Memory); memoryLength != 2 {
		t.Fatalf(`Unexpected Memory length (%d)`, memoryLength)
	}
	if machine.LargestSeenPrime() != value {
		t.Fatalf("Failed to append a prime to Memory")
	}
}

func TestMachineNext(t *testing.T) {
	machine := PrimeMachine{Memory: []int{2}}
	if prime := machine.Next(); prime != 3 {
		t.Fatalf("Received unexpected prime on next (%d)", prime)
	}
	if machine.LargestSeenPrime() != 3 {
		t.Fatalf("Failed to append a prime to Memory")
	}
}
