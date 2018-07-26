package server

import (
	"net/http"
	"log"
	"golang.org/x/net/websocket"
	"github.com/JILeXanDR/parser/ws"
)

func InitServer(port string) {

	ws.WebSocketClients = make(ws.Clients)

	http.Handle("/", http.FileServer(http.Dir("static")))
	http.Handle("/io", websocket.Handler(ws.WebSocketHandler))
	//http.HandleFunc("/api/currencies", currenciesHandler)
	//http.HandleFunc("/api/exchange-rates", currenciesHandler)

	log.Println("Start server at http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
