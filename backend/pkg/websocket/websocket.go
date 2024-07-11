package websocket

import (
	// "fmt"
	// "io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)


var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool {return true} , //allows for connection from anywhere
}

// a reader function that will take a pointer to the websocket connection 

// func Reader(conn *websocket.Conn) {
// 	for {
// 		// read the message or parse the message 
// 		messageType, message, err := conn.ReadMessage()

// 		if err != nil {
// 			log.Println(err)
// 		}

// 		fmt.Println(string(message))

// 		if err := conn.WriteMessage(messageType, message); err != nil{
// 			log.Println(err)
// 			return
// 		}
// 	}

// }

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
	}

	return ws, err
}

// func Writer(conn * websocket.Conn) {
// 	for {
// 		fmt.Println("Sending message")
// 		messageType, r, err := conn.NextReader()

// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		w, err := conn.NextWriter(messageType)

// 		if err != nil {
// 			fmt.Println(err)
// 			return 
// 		}

// 		if _, err := io.Copy(w,r); err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		if err := w.Close(); err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 	}
// }