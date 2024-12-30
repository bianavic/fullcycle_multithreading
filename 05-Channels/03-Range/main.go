package main

import (
	"fmt"
	"sync"
)

// thread 1
func main() {

	// canal VAZIO
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10) // temos 10 'creditos'
	go publish(ch)
	reader(ch, &wg)
	wg.Wait() // ao finalizar a leitura dos 'çreditos' é avisado que 'wait' terminou
}

/*
funcao que le um canal de inteiros
assim q le o valor, esvazia o canal
*/
func reader(ch chan int, wg *sync.WaitGroup) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
		wg.Done() // conforme le, avisa que a iteracao acabou
	}
}

/*
funcao que recebe 1 canal de inteiros

	cada vez que o for roda, faz um i e manda os dados para o canal
*/
func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) // evitar deadlock fechando o canal
}
