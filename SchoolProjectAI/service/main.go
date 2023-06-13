package main

import (
	"alexalsing/websockets/response"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func main() {

	fmt.Print("Hello?")
	http.HandleFunc("/GPT3", func(w http.ResponseWriter, r *http.Request) {

		// Upgrade upgrades the HTTP server connection to the WebSocket protocol.
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade failed: ", err)
			return
		}
		defer conn.Close()

		// Continuosly read and write message
		history := ""
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read failed:", err)
				break
			}
			input := string(message)
			prompt := history + "You say: " + input + "\nFriend: "

			res := response.AiResponse(prompt)

			history = prompt + res + "\n"

			fmt.Print(history)

			message = []byte(res)
			//Write message(AI answer goes here)
			err = conn.WriteMessage(1, message)
			if err != nil {
				log.Println("write failed:", err)
				break
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
	})

	http.ListenAndServe(":8080", nil)
}
