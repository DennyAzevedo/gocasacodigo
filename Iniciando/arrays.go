package main

import "fmt"

func main() {
	var a [3]int
	numeros := [5]int{1, 2, 3, 4, 5}
	primos := [...]int{2, 3, 5, 7, 11, 13}
	nomes := [2]string{}

	fmt.Printf("Arrays - a: %v, numeros: %v, primos: %v, nomes: %v\n",
		a,
		numeros,
		primos,
		nomes,
	)
	fmt.Printf("Tamanho dos Arrays - a: %d, numeros: %d, primos: %d, nomes: %d\n",
		len(a),
		len(numeros),
		len(primos),
		len(nomes),
	)

	var multiA[2][2]int
	multiA[0][0], multiA[0][1] = 3, 5
	multiA[1][0], multiA[1][1] = 7, -2

	multiB := [2][2]int{{2, 13}, {-1, 6}}
	fmt.Printf("Matrizes:\nmultiA: %v\nmultiB: %v\n", multiA, multiB)
}
