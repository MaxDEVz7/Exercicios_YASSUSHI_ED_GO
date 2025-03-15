package main

import (

"fmt"

)

func main() {

	var vetor [6]int

	vetor[0] = 1

	vetor[1] = 2

	vetor[2] = 4

	vetor[3] = 5

	vetor[4] = 6

	vetor[5] = 4


	for i := 0; i < 6; i++{

		fmt.Println(vetor[i]);
	}
}