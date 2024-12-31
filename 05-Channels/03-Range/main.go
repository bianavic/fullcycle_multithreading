package main

import "fmt"

// thread 1
func main() {
	hello := make(chan string)
	go recebe("Hello", hello)
	ler(hello)
}

/*
RECEIVE-ONLY
ao executar a funcao estamos:

  - o canal esta recebendo um valor em hello

  - estamos pegando o valor do nome

  - o q tb significa que estamos PUBLICANDO informacoes no nosso canal

    porque a funcao apenas RECEBE dados, podemos setar na assinatura a DIRECAO que o canal tem
    isso diz que este CANAL ira apenas RECEBER informacao
*/
func recebe(nome string, hello chan<- string) { // DIRECAO de receber informacoes (seta do lado direito de chan) que o canal tem
	hello <- nome
}

/*
SEND-ONLY
funcao que APENAS le dados:

	ela esvazia o canal -
	 recebe dados
	 entrega dados

esta DIRECAO é de ESVAZIAR (seta do lado ESQUERDO de chan) -> sua funcao é de entregar resultado
*/
func ler(data <-chan string) {
	fmt.Println(<-data)
}
