package generator

type LCG struct {
	M int
	C int
	A int
}

// NextValue Creates the next value for these parameters
func (parameters *LCG) NextValue(current int) int {
	return (parameters.A*current + parameters.C) % parameters.M
}
