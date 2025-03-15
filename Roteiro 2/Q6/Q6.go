//Neste arquivo contém as respostas das questões 6 a 8.

package main

import "fmt"

type Produto struct{
	nome string;
	preco float64;
	quantidade int;
}

func main (){

	Notebook := Produto{
		nome: "Notebook",
		preco: 75.99,
		quantidade: 12,
	}
	
	fmt.Println("Nome do produto padrão do sistema cadastrado: ", Notebook.nome, "	Preço: ", Notebook.preco, "	quantidade no sistema: ", Notebook.quantidade)
}