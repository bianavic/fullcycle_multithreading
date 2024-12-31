package main

import "fmt"

// thread 1
func main() {

	// canal VAZIO
	ch := make(chan int)
	go publish(ch) // cuidado com deadlock, thread morrer enqto tentar ler, pq o reader pode ficar infinito
	/*
		o reader permanece em loop quase infinito, sempre aguardando o canal ser preenchido
		nao roda em background, por isso nao possui go na frente,
		se tiver, o processo morre antes de rodar
	*/
	reader(ch)

	// OPCAO 2: remove a funcao reader(ch)
	//for x := range ch {
	//	fmt.Printf("Received %d\n", x)
	//}
}

/*
funcao que le um canal de inteiros
lê o canal, joga o valor para x e imprime
assim q le o valor, ele esvazia o canal
*/
func reader(ch chan int) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
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
	close(ch) // fecha o canal, indicando que nada mais vai entrar no canal. evitando deadlock qdo concluir o processo i < 10
}
