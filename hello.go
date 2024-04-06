package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	hello()

	for {
		menu()
		comando := lerComando()

		switch comando {
		case 0:
			fmt.Println("Saindo, bye")
			os.Exit(0)
		case 1:
			iniciarMonitoramento()

		case 2:
			fmt.Println("Exibindo logs")
		default:
			fmt.Println("Comando desconhecido!")
			os.Exit(-1)
		}
	}

}

func hello() {
	nome := "Matheus"
	var versao = 1.1
	var idade int

	fmt.Println("hello world!", nome, "!")
	fmt.Println("versao:", versao, "!")
	fmt.Println("idade:", idade)

	fmt.Println(reflect.TypeOf(idade))
	fmt.Println(reflect.TypeOf(versao))
	fmt.Println(reflect.TypeOf(nome))
}

func lerComando() int8 {
	var comando int8
	// fmt.Scanf("%d", &comando)
	fmt.Scan(&comando)

	fmt.Println("Endereço de memoria", &comando)
	fmt.Println("O comando escolhido foi", comando)
	fmt.Println("")
	return comando
}

func menu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	// var sites [4]string // tamanho fixo
	sites := []string{"https://os.redezone.com.br", "https://energia.redezone.com.br", "https://iot.redezone.com.br"}

	for i := 0; i < monitoramentos; i++ {

		for i, site := range sites {

			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}

		time.Sleep(delay * time.Second)
		fmt.Println("")
		fmt.Println("")
	}

	fmt.Println("")

}

func testaSite(site string) {

	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}
