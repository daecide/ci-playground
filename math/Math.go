package math

type Math struct{}

func (ptr *Math) Add(a, b int) int {
	return a + b
}

func (ptr *Math) Deduct(a, b int) int {
	return a - b
}
