LCG Parameter Generator

---

Generates parameters for LCG type PRNGs.

Currently, implements fixed period generator parameter search.

```
go run ./cmd/generate_lcg.go 25000                               
Factors: [2 5]
Params 1:
 a: 21
 c: 3
 m: 25000
Params 2:
 a: 42
 c: 3
 m: 25000

```

Uses the info in the [c ≠ 0 section on LCGs](https://en.wikipedia.org/wiki/Linear_congruential_generator) in Wikipedia:

> When c ≠ 0, correctly chosen parameters allow a period equal to m, for all seed values. This will occur if and only if:
> 1. m and c are relatively prime,
> 2. a-1 is divisible by all prime factors of m,
> 3. a-1 is divisible by 4 if m is divisible by 4

You can also generate primes with the provided scripts under `/cmd`.
The `primes` module contains the pregenerated primes up to 3 million.

The modules are as follows:

- CMD: Scripts
- Generator: LCG code
- Primes: Generating primes, prime related utils, and pregenerated primes
- Search: Parameter search for LCGs
