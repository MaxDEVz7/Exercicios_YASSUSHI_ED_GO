package aula_3

import "fmt"

func incrementByPointer(val *int) {
	*val++
	fmt.Println("Dentro da função incrementByPointer:", *val)
}
func Exemplo4() {
	x := 10
	incrementByPointer(&x)
	fmt.Println("Fora da função incrementByPointer:", x)
}
