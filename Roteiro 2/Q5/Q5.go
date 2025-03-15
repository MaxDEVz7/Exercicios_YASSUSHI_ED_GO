package main

import("fmt")

func main (){		
	var matriz [3][3]string

	matriz[0][0] = "1"

	matriz[1][0] = "2"

	matriz[2][0] = "3"

	fmt.Println("Insira o ra do aluno 1")

	fmt.Scan(&matriz[0][1]);

	fmt.Println("Insira o nome do aluno 1")

	fmt.Scan(&matriz[0][2]);

	fmt.Println("Insira o ra do aluno 2")

	fmt.Scan(&matriz[1][1]);

	fmt.Println("Insira o nome do aluno 2")

	fmt.Scan(&matriz[1][2]);

	fmt.Println("Insira o ra do aluno 3")

	fmt.Scan(&matriz[2][1]);

	fmt.Println("Insira o nome do aluno 3")

	fmt.Scan(&matriz[2][2]);

	fmt.Println("Indice RA Nome")

	for i:=0; i<3; i++{
		for j:=0; j<3; j++{
			fmt.Print(matriz[i][j] + " ");

		}
		fmt.Println();
	}
}