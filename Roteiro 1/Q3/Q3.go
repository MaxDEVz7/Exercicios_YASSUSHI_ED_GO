package main

import (
 "fmt"
)

func ClassicarNota(nota float32) (string){

	switch {
		case nota >= 9.0:
			return "Excelente"

		case nota >= 7.0:
			return "Bom"

		case nota >= 5:
			return "Regular"

		case nota >= 3:
			return "Ruim"

		case nota <= 2:
			return "Muito Ruim"

		default:
			return "Nota inválida"
 }
}

func main() {

 var nota float32

 fmt.Println("Escreva sua nota:")

 fmt.Scanln(&nota)

 fmt.Println(ClassicarNota(nota))
}