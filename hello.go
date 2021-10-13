package main //pacote principal e deve começar por ele

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	//"io/ioutil"
	"bufio"
	"net/http"
	"os"
	"strconv"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	exibeIntroducao() // chama na função principal so pra exibir
	//leSiteDoArquivo()
	registraLog("site-falso", false)
	for {

		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:

			inicarMonitoramento()

		case 2:
			fmt.Println("Exibindo Logs...")
			imprimiLogs()

		case 0:
			fmt.Println("Saindo do Programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando")
			os.Exit(-1) // -1 indica que o ocorreu algum erro
		}
	}

}

func exibeIntroducao() {
	nome := "Edrielle"
	versao := 1.1
	fmt.Println("Ola sra.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func leComando() int { // o int depois do (), é para informar o tipo do retorno

	var comandoLido int
	fmt.Scan(&comandoLido) // como se fosse o input, o comando scan ele pode colocar ali sem o scanf, e tirar o %d
	fmt.Println("O comando escolhido foi", comandoLido)
	fmt.Println("")

	return comandoLido
}

func exibeMenu() {

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

}

func inicarMonitoramento() {
	fmt.Println("Monitorando...")
	/*sites := []string{
		"https://random-status-code.herokuapp.com",
		"https://www.alura.com.br",
		"https://www.caelum.com.br",
	}*/
	sites := leSiteDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando o site", i, site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")

}

func testaSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu o erro:", err.Error())
	}
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true) // aqui a função vai analisar o parametro site recebido se é verdadeiro, para registrar
	} else {
		fmt.Println("O site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false) //aqui a função vai analisar o parametro site recebido se é falso, para registrar
	}

}

func leSiteDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")
	//arquivo, err := ioutil.ReadFile("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu erro:", err.Error())

	}
	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n') // tem que ser o \n que é pulada de linha, na aspa simples, dupla da como string é um delimitador
		linha = strings.TrimSpace(linha)      // remove quebra de linha, espaços, tabs etc
		fmt.Println(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
			//if err != nil {
			//fmt.Println("O erro ocorreu:", err.Error())
		}
	}
	arquivo.Close() // fecha arquivo.
	//fmt.Println(sites)
	//fmt.Println(string(arquivo))
	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) // ela da opção de abrir arquivo, as flags informa o que pode fazer com aquele arquivo

	if err != nil {
		fmt.Println("Deu erro", err.Error())
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	//fmt.Println(arquivo)

	arquivo.Close() // o + concatena
}

func imprimiLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu o erro: ", err.Error())
	}
	fmt.Println(string(arquivo)) // joga o array de bytes para string
}
