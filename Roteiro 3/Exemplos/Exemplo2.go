package aula_3

import "fmt"

func Exemplo2() {
	var a int = 42
	var p *int = &a
	fmt.Println("Valor de a antes da modificação:", a)
	*p = 100
	fmt.Println("Valor de a após a modificação:", a)
}
