package main

import (
	"fmt"

	"github.com/seu-usuario/meu-projeto-go/internal/fibonacci"
	"github.com/seu-usuario/meu-projeto-go/internal/hello"
)

func main() {
	fmt.Println("🚀 Meu primeiro projeto em Go com estrutura de mercado!")

	hello.SayHello()

	n := 10
	valor := fibonacci.Fibonacci(n)
	fmt.Printf("F(%d) = %d\n", n, valor)

	fibonacci.PrintSequence(n)

	fmt.Println("Função anônima: ")
	multiplicacao := func(n1, n2 int) int{
		return n1 * n2
	}
	fmt.Println("Quais números serão multiplicados")

	n1:= 7 
	n2:= 7
	fmt.Printf("A multiplicação de %d x %d = %d", n1 , n2, multiplicacao(n1, n2))
}
