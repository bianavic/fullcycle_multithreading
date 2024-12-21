package main

import "fmt"

// thread 1
func main() {
	// jogar um valor do tipo string pra esse canal
	canal := make(chan string) // VAZIO

	// thread 2
	// funcao anonima com acesso ao canal
	go func() {
		// jogando o valor da string para dentro do canal
		canal <- "Hello, Stranger!" // CHEIO
	}()

	// thread 1: o valor que adicionei ao canal, jogo para a variavel
	msg := <-canal // canal ESVAZIA
	fmt.Println(msg)
}
