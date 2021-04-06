package api

import (
	"encoding/json"
	"log"
	"net/http"
	"server/repository"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan *repository.Message)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/chats/{chatId:[0-9]+}/messages", messageHandler).Methods("POST")
	router.HandleFunc("/ws", wsHandler)
	go webSocketResponse()

	log.Fatal(http.ListenAndServe(":8000", router))
}

func writer(newMessage *repository.Message) {
	broadcast <- newMessage
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	var newMessage repository.Message
	if err := json.NewDecoder(r.Body).Decode(&newMessage); err != nil {
		log.Printf("ERROR: %s", err)
		http.Error(w, "Bad request", http.StatusTeapot)
		return
	}
	defer r.Body.Close()
	go writer(&newMessage)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	clients[ws] = true
}

func webSocketResponse() {
	for {
		val := <-broadcast
		//нужно получить айди собеседника чтобы ему отправить сообщение

		for client := range clients {
			err := client.WriteJson(val)
			if err != nil {
				log.Printf("Websocket error: %s", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
