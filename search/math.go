package search

func GCD(a int, b int) int {
	for b > 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
