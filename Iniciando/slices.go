package main

import "fmt"

func main() {
	var a []int
	primos := []int{2, 3, 5, 7, 11, 13}
	nomes := []string{}

	fmt.Printf("Arrays:\na - %v\nprimos - %v\nnomes - %v\n",
		a,
		primos,
		nomes,
	)

	b := make([]int, 10)
	fmt.Printf("Slice b: %v\nTamanho: %d, Capacidade: %d\n",
		b,
		len(b),
		cap(b),
	)

	c := make([]int, 10, 20)
	fmt.Printf("Slice c: %v\nTamanho: %d, Capacidade: %d\n",
		c,
		len(c),
		cap(c),
	)
}
