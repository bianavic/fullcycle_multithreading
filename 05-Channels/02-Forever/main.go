package main

/*
SITUACAO 1: fatal error: all goroutines are asleep - deadlock!
Nao foi adicionada uma thread que possa enxer o canal, e o processo esta preso. por isso o deadlock
// thread 1
func main() {

	// canal VAZIO
	forever := make(chan bool)

	// aguarda canal cheio para que ele o esvazie
	// segura o processo do main
	<-forever
}
*/

/*
SITUACAO 2: fatal error: all goroutines are asleep - deadlock!
No exemplo abaixo tb nao existe perspectiva de preencher o canal.
func main() {

	// canal VAZIO
	forever := make(chan bool)

	// thread 2 rodando em background
	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}
	}()

	// aguarda canal cheio para que ele o esvazie
	// segura o processo do main
	<-forever
}
*/

// thread 1
func main() {

	// canal VAZIO
	forever := make(chan bool)

	// thread 2: ao concluir o laço ira encher o valor
	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}
		forever <- true
	}()

	// aguarda canal cheio para que ele o esvazie
	// segura o processo do main
	<-forever
}
