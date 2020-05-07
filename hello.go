package main

import (
	"fmt"
	// "reflect"
)

func main() {
	// var nome string = "Wilber"
	nome := "Wilber"
	var versao float32 = 1.1
	var comando int
	fmt.Println("Olá Sr.", nome)
	fmt.Println("A versão do programa é", versao)
	// fmt.Println(reflect.TypeOf(nome))
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)
}
