# fullcycle_multithreading

1- Go Routines

2- Wait Groups

3- Problema simples de concorrencia

4- Entendendo Mutex e Operacoes Atomicas

verificar race condition
```GO
go run -race main.go
```

soma atomica - incremento utilizando pacote prorprio do GO
```GO
atomic.AddUint64(&number, 1)
```

5- Channels
- Iniciando com Channels
- Forever
- For/Range
- Range with WaitGroup
- Directions: canais receive-only e send-only
  - receive-only: direcao de receber informacoes (chan<-)
  ```go
    func recebe(nome string, hello chan<- string) {
    hello <- nome
    }
  ```
    - send-only: direcao de enviar informacoes (<-chan)
  ```go
    func ler(data <-chan string) {
    fmt.Println(data)
    }
  ```
- Load Balancer: balanceamento de carga com workers

    ex: num servidor web
  - aqui receberiamos as requests
  ```go
    for i := 0; i < 100; i++ {
       data <- i // a var data RECEBE request
	 }
  ```
  - enqto worker PROCESSARIA as requisicoes
  ```go
  func worker(workerId int, data chan int) { // PROCESSA as requests
    for x := range data {
    println("Worker id %d got %d\n", workerId, x)
    time.Sleep(time.Second)
     }
  }
  ```
- Select: escolher o canal que esta pronto para ser lido ou escrito