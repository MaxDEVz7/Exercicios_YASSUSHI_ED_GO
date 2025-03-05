package main

import (

 "fmt"

)

func somaN(num int) (int){

    var soma int

    for i := 0; i <= num; i++{
        soma += i
    }

    return soma
}



func main() {

    fmt.Println(somaN(5))

}

