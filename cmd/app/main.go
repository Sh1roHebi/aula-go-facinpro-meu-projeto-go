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
}
