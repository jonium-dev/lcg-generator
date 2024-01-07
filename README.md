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