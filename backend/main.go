package main

import (
	"fmt"
	"net/http"
	"github.com/zamirszn/go_react_chat_app/pkg/websocket"
)


func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r,)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}
	go websocket.Writer(ws)

	websocket.Reader(ws)
}

func setupRoutes() {
	
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("chat app v1")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}


