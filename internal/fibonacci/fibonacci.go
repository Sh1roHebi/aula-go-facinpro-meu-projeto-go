package fibonacci

import "fmt"

func Fibonacci(n int) int {
	if n < 0 {
		panic("n não pode ser negativo")
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func Sequence(n int) []int {
	if n < 0 {
		panic("n não pode ser negativo")
	}
	seq := make([]int, n+1)
	if n >= 0 {
		seq[0] = 0
	}
	if n >= 1 {
		seq[1] = 1
	}
	for i := 2; i <= n; i++ {
		seq[i] = seq[i-1] + seq[i-2]
	}
	return seq
}

func PrintSequence(n int) {
	seq := Sequence(n)
	fmt.Printf("Sequência de Fibonacci até F(%d): %v\n", n, seq)
}
