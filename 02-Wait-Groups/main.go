package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
		wg.Done() // a cada iteracao, subtrai um valor da qtdade de operacoes
	}
}

// associa go routines com o recurso de sincronizar a execucao das operacoes
func main() { // Thread 1 (main tb é uma thread)
	waitGroup := sync.WaitGroup{} // criou um wg
	waitGroup.Add(25)             // Adicionar quantidade de tarefas / operacoes
	go task("A", &waitGroup)      // Thread 2
	go task("B", &waitGroup)      // Thread 3
	go func() {                   // Thread 4
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(1 * time.Second)
			waitGroup.Done() // cada passagem pelo loop, conclui.
		}
	}()
	waitGroup.Wait() // esperar ate que as operacoes sejam finalizadas
}
