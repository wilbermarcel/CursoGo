package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 2
const delay = 3

func main() {
	exibeIntroducao()
	for {
		exibeMenu()
		trataMenu(leComando())
	}
}

func exibeIntroducao() {
	var (
		nome   = "Wilber"
		versao = 1.1
	)
	fmt.Println("Olá Sr.", nome)
	fmt.Println("A versão do programa é", versao)
}

func exibeMenu() {
	fmt.Println()
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)
	return comando
}

func trataMenu(comando int) {
	switch comando {
	case 1:
		iniciarMonitoramento()
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Comando inválido")
		os.Exit(-1)
	}
}

func iniciarMonitoramento() {
	sites := make([]string, 2)
	fmt.Println()
	fmt.Println("Monitorando...")
	sites[0] = "https://www.alura.com.br"
	sites[1] = "https://random-status-code.herokuapp.com/"
	sites = append(sites, "https://www.caelum.com.br")
	for i := 0; i < monitoramentos; i++ {
		fmt.Println("Monitoramento", i+1)
		for _, site := range sites {
			testaSite(site)
		}
		if i < (monitoramentos - 1) {
			time.Sleep(delay * time.Second)
			fmt.Println()
		}
	}
}

func testaSite(site string) {
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}
