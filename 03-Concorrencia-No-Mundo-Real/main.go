package main

import (
	"fmt"
	"net/http"
	"time"
)

// qtdade visitas que o site teve
var number uint64 = 0

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		number++
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Você é o visitante número %d", number)))
	})
	http.ListenAndServe(":3000", nil)
}

/*
OBS: para cada requisicao, esse servidor gera 1 nova thread, gerando CONCORRENCIA na VARIAVEL number

APACHE BENCHMARK:
simular concorrencia de 10mil requisicoes com 100 pessoas acessando simultaneamente

ab -n 10000 -c -100 http://localhost:3000/
curl http://localhost:3000

os resultados mostram que requisicoes foram perdidas no meio do caminho. imagine se fosse dinheiro?
*/
