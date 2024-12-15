package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

// qtdade visitas que o site teve
var number uint64 = 0

/*
Encontrar cenario de race condition:
go run -race main.go

Resolver race condition
Lock() e Unlock()
*/

// MUTEX
//func main() {
//	m := sync.Mutex{}
//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		m.Lock()
//		number++ // no momento que tentarem acessar o valor da variavel, ele ira bloquear
//		m.Unlock()
//		time.Sleep(300 * time.Millisecond)
//		w.Write([]byte(fmt.Sprintf("Você é o visitante número %d", number)))
//	})
//	http.ListenAndServe(":3000", nil)
//}

// ATOMIC
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddUint64(&number, 1) // internamente faz lock e unlock
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Você é o visitante número %d", number)))
	})
	http.ListenAndServe(":3000", nil)
}
