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