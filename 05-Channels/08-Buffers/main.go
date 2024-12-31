package main

func main() {
	ch := make(chan string, 2) // adiciona 2 para dar tempo de ler os 2, evitando deadlock
	ch <- "hello"
	ch <- "world"

	println(<-ch)
	println(<-ch)
}
