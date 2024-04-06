package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	hello()
	leSitesDoArquivo()
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
	// sites := []string{"https://os.redezone.com.br", "https://energia.redezone.com.br", "https://iot.redezone.com.br"}
	sites := leSitesDoArquivo()

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

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro no site: ", site, "\n", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}

func leSitesDoArquivo() []string {

	var sites []string

	// retorna o endereço de memoria que contem o arquivo
	arquivo, err := os.Open("sites.txt") // &{0xc00007a1e0}

	// retorna um array de bytes
	// arquivo, err := ioutil.ReadFile("sites.txt") /*[104 116 116 112 115 58 47 47 111 115 46 114 101 100 101 122 111 110 101 46 99 111 109 46 98 114 10 104 116 116 112 115 58 47 47 101 110 101 114 103 105 97 46 114 101 100 101 122 111 110 101 46 99 111 109 46 98 114 32 10 104 116 116 112 115 58 47 47 105 111 116 46 114 101 100 101 122 111 110 101 46 99 111 109 46 98 114]*/

	if err != nil {
		fmt.Println("Ocorreu um erro: \n", err)
		return []string{}
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n') // ler ate a quebra de linha usar '', "eh string"
		// remover espaços e quebra linhas
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		fmt.Println(linha)
		if err == io.EOF {
			fmt.Println("Ocorreu um erro:", err)
			break
		}

	}
	arquivo.Close()

	// fmt.Println(string(arquivo))
	fmt.Println(sites)
	return sites
}
