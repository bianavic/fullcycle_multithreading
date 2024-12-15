package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

// Thread 1 (main tb é uma thread)
func main() {
	// 1- tarefa A e B sao executadas de forma concorrente
	//task("A") // Thread 2
	//task("B") // Thread 3

	// 2- como reduzir o tempo pela metade? acontecer de forma simultanea?
	go task("A") // Thread 2
	go task("B") // Thread 3
	// funcao anonima
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(1 * time.Second)
		}
	}()

	// NADA AQUI -> qdo main chegar aqui ele SAI, o programa encerra, as threads morrem.
	time.Sleep(15 * time.Second) // evita que o programa encerre antes de executar tudo
}

// funcoes anonimas
