package main

import (
	"fmt"
	"github.com/joniumGit/lcg-generator/primes"
	"os"
	"strconv"
	"strings"
)

// Generates primes up to Args[1].
// Primes are printed to output or written to Args[2].
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

	var pkg string
	if len(os.Args) > 3 {
		pkg = os.Args[3]
	} else {
		pkg = "main"
	}

	var machine primes.PrimeMachine

	content := fmt.Sprintf("package %s\n\nvar PRIMES = [LENGTH]int{\n", pkg)

	fmt.Printf("Starting generation\n")
	progress := 0.00
	for i := 2; i <= target; i++ {
		if machine.IsPrime(i) {
			content += fmt.Sprintf("    %d,\n", i)
		}
		t := (float64(i) / float64(target)) * 100
		if t-progress > 5 {
			fmt.Printf("%.2f\r", t)
		}
	}
	fmt.Printf("\nDone\n")

	content += "}\n"
	content = strings.Replace(content, "LENGTH", fmt.Sprintf("%d", len(machine.Memory)), 1)

	if len(os.Args) > 2 {
		file, err := os.Create(os.Args[2])
		if err != nil {
			panic(err)
		}

		_, err = file.WriteString(content)
		if err != nil {
			panic(err)
		}

		err = file.Close()
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Print(content)
	}
}
