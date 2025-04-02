package aula_3

import "fmt"

func Exemplo1() {
	var a int = 42
	var p *int = &a
	fmt.Println("Valor de a:", a)
	fmt.Println("Endereço de a:", &a)
	fmt.Println("Valor de p (endereço de a):", p)
	fmt.Println("Valor apontado por p:", *p)
}
