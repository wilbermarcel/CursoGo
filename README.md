# Curso de Golang
Repositório do curso de Go (Alura)

<p align="center">
  <a href="https://golang.org/">
    <img
      alt="Golang"
      src="https://blog.remontti.com.br/wp-content/uploads/2019/10/Golang.png"
      width="400"
    />
  </a>
</p>

# Minhas anotações do curso
## Parte 1

    https://golang.org/ 
    https://golang.org/dl/

    Gopher -> mascote do Go

    go version

    compilação:
    go build hello.go
    ./hello

    go run hello.go

    //primeiro programa
    package main

    import "fmt"

    func main() {
	    fmt.Println("Olá Mundo")
    }

    //Declarando variáveis
    var nome string = "Wilber"
    var nome = "Wilber"
    nome := "Wilber"

    //Comando de leitura
    fmt.Scan(&<variavel>)

    https://www.alura.com.br/artigos/variaves-com-go-lang

    fmt.Println("\033[2J") <- Limpando a tela

    //Fechando um programa
    import "os"
    os.Exit(0) <- ocorreu tudo bem
    os.Exit(-1) <- ocorreu algum erro

    https://golang.org/pkg/

    //Devolvendo mais de um valor no Go
    func devolveNomeEIdade() (string, int) {
      nome := "Wilber"
      idade := 30
      return nome, idade
    }

    nome, idade := devolveNomeEIdade()
    ou
    _, idade := devolveNomeEIdade() OU nome, _ := devolveNomeEIdade() -> Para ignorar um dos retornos use o underline

    https://golang.org/pkg/os/#pkg-constants
    https://golang.org/src/time/format.go
