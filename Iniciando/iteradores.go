package main

import "fmt"

func main() {
	fmt.Println("Iteradores em Go")
	fmt.Println("Iteradores são usados para percorrer coleções de dados.")
	fmt.Println("Go só possui um tipo de iterador, o for.")
	fmt.Println("Exemplo de uso de iteradores.")
	fmt.Println("for como while")
	a, b := 0, 10
	for a < b {
		a += 1
		fmt.Println(a)
	}
	fmt.Println("for como for, tem o mesmo comportamento do for de C/C++")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// Não podemos imprimir o valor de i fora do escopo do for
	//fmt.Println(i)
	fmt.Println("for range, usado para iterar sobre slices, arrays, maps e strings")
	numeros := []int{1, 2, 3, 4, 5}
	fmt.Printf("Slice de números (inicial): %v\n", numeros)
	for n := range numeros {
		numeros[n] *= 2
	}
	fmt.Printf("Slice de números (final): %v\n", numeros)
	fmt.Println("Usando for range, com indice e valor")
	for i, n := range numeros {
		fmt.Printf("Índice: %d, Valor: %d\n", i, n)
	}
	fmt.Println("Usando for range, com apenas o valor")
	for _, n := range numeros {
		fmt.Printf("Valor: %d\n", n)
	}
}
