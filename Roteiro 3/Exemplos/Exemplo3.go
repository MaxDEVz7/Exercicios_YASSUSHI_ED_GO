package aula_3

import "fmt"

func increment(val int) {
	val++
	fmt.Println("Dentro da função increment:", val)
}
func Exemplo3() {
	x := 10
	increment(x)
	fmt.Println("Fora da função increment:", x)
}
