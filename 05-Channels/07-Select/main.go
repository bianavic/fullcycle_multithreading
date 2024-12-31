package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	id  int64
	Msg string
}

//func main() {
//	c1 := make(chan int)
//	c2 := make(chan int)
//
//	// funcao anonima, roda em thread separada
//	go func() {
//		c1 <- 1 // joga o valor para o canal
//		time.Sleep(time.Second * 4)
//	}()
//	go func() {
//		time.Sleep(time.Second * 4)
//		c2 <- 2 // joga o valor para o canal
//	}()
//
//	//// select para o programa e imprime valor que chegou em c1 ou c2,
//	//// nao sabemos qual vai chegar primeiro, ele imprime na sequencia que chegar
//	//select {
//	//case msg1 := <-c1:
//	//	println("Received", msg1)
//	//case msg2 := <-c2:
//	//	println("Received", msg2)
//	//case <-time.After(time.Second * 3): // cria timeout: se ninguem retornar antes de X tempo, farei outra coisa
//	//	println("Timeout")
//	//default:
//	//	println("No one was ready") // se ninguem estiver pronto pq estao todos esperando, nao quero perder tempo: imprime o default
//	//}
//
//	/*
//		adicinando FOR:
//		laco roda 2x:
//			na primeira rodada, quem ganha é mgs 1
//			na segunda rodada, quem ganha default: a msg 1 esta vazia, as outras nem serao chamadas por causa dos segundos, cai no default
//	*/
//	//for i := 0; i < 2; i++ {
//	//	select {
//	//	case msg1 := <-c1:
//	//		println("Received", msg1)
//	//	case msg2 := <-c2:
//	//		println("Received", msg2)
//	//	case <-time.After(time.Second * 3):
//	//		println("Timeout")
//	//	default:
//	//		println("No one was ready")
//	//	}
//	//}
//
//	/*
//		LOOP INFINITO FOR:
//		imagina que trabalho com rabbitmq e kafka e quero ficar ouvindo infinitamente
//	*/
//	for {
//		select {
//		case msg1 := <-c1: // rabbitmq
//			println("Received", msg1)
//		case msg2 := <-c2: // kafka
//			println("Received", msg2)
//		case <-time.After(time.Second * 3):
//			println("Timeout")
//		default:
//			println("No one was ready")
//		}
//	}
//}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64 = 0 // cuidado com o i GLOBAL: concorrencia - se muitas mgs rodando igual, muita gente querendo escrever nessas mesmas msgs, o i pode ser afetado

	// RABBITMQ
	go func() {
		for {
			// i++ // pode ser afetado - concorrencia
			atomic.AddInt64(&i, 1) // correto - evita concorrencia, a soma sempre fai ter um lock
			time.Sleep(time.Second * 2)
			msg := Message{i, "Hello from RabbitMQ"}
			c1 <- msg
		}
	}()
	// KAFKA
	go func() {
		for {
			// i++ // pode ser afetado - concorrencia
			atomic.AddInt64(&i, 1) // correto - evita concorrencia, a soma sempre fai ter um lock
			time.Sleep(time.Second * 1)
			msg := Message{i, "Hello from Kafka"}
			c2 <- msg
		}
	}()

	for {
		select {
		case msg := <-c1: // rabbitmq
			fmt.Printf("Received from RabbitMQ: ID: %d - %s\n", msg.id, msg.Msg)
		case msg := <-c2: // kafka
			fmt.Printf("Received from Kafka: ID: %d - %s\n", msg.id, msg.Msg)
		case <-time.After(time.Second * 3):
			println("Timeout")
		}
	}
}
