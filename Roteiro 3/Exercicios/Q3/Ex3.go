package main
import (
  "fmt"
)

func Ex3(){

  // 1. Declare uma variável b do tipo float64 e um ponteiro p2 que aponte
  // para b.

  var b float64;
  var p2 *float64 = &b;

  fmt.Print(p2);

  // 1.1) Modifique o valor de b através do ponteiro p2 e imprima o resultado.

  *p2 = 10.0;
  fmt.Println(b);

}