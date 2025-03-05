package main

import "fmt"

func main() {

 var num int

 fmt.Println("Escreva um número: ")

 fmt.Scanln(&num)

 fmt.Println(parOrImpar(num))

}

func parOrImpar(num int) (string){

 if num % 2 == 0{

  return "par"

 }else{

  return "impar"

 }

}

