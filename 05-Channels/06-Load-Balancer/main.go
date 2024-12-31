package main

import (
	"fmt"
	"time"
)

/*
funcao realiza processamento
sera a maquina que vai processar os dados

	OBS: depende da qtdade de CPU para comportar a qtdade de workers/threads
*/
func worker(workerId int, data chan int) {
	// cada vez q o canal encher, itero para fazer o processamento
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

// solucao 2: criar 10000 workers (threads) com o FOR loop
func main() {
	// criar canal de data
	data := make(chan int)
	QtddWorkers := 10000 // 10000 threads

	// o FOR inicializa workers
	for i := 0; i < QtddWorkers; i++ {
		go worker(i, data)
	}

	// enviar infos para o canal, joga os dados
	// mas temos 10 workers, entao o processamento será 10 x mais rapido
	for i := 0; i < 100000; i++ {
		data <- i // data RECEBE infos com o valor de i
	}
}

//func main() {
//	// criar canal de data
//	data := make(chan int)
//
//	// inicializa worker
//	go worker(1, data) // apenas 1 worker -> lentidao, chamadas sao enfileiradas, leitura de forma serializada
//	/*
//		// solucao 1 para aumentar velocidade:
//			enqto 1 espera o time.Sleep, o worker 2 pega a prox msg e comeca a processar
//			enqto worker 2 esta esperando processar a info, o worker 1 terminou de processar e ja esta pronto para pegar a prox info
//	*/
//	go worker(2, data) // adiciona worker para aumentar a velocidade
//
//	// enviar infos para o canal
//	for i := 0; i < 100; i++ {
//		data <- i // data RECEBE infos com o valor de i
//	}
//}
