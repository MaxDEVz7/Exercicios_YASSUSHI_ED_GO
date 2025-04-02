package main
import (
  "fmt"
  
)

func main(){
  var a int = 10
  var b int = 20

  fmt.Println("Valor de 'a' antes da modificação:", a)
  fmt.Println("Valor de 'b' antes da modificação:", b)
  
  swap(&a, &b)
  
  fmt.Println("valor de a:", a)
  fmt.Println("valor de b:", b)

}

func swap(a, b *int){
  var backup int = *a
  *a = *b
  *b = backup
}