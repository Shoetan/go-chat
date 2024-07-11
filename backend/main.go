package main

import (
	"fmt"
	"net/http"

	"github.com/go-chat/pkg/websocket"
)




func serveWebsocket(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.Host)

	//make websocket connection 
	conn, err := websocket.Upgrade(w,r)

	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}

	client := &websocket.Client{
		Conn:conn,
		Pool:pool,

	}

	pool.Register <- client

	client.Read()

}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Running server")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request){
		serveWebsocket(pool, w, r)
	})
}

func main() {
	fmt.Println("Websocket server v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)

	
}