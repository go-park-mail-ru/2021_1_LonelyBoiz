package repository

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var clients = make(map[int]*websocket.Conn)
var messagesChan = make(chan *Message)
var chatsChan = make(chan *Chat)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/chats/{chatId:[0-9]+}/messages", messageHandler).Methods("POST")
	router.HandleFunc("/likes", likesHandler).Methods("POST")
	router.HandleFunc("/ws", wsHandler)
	go webSocketResponse()

	log.Fatal(http.ListenAndServe(":8000", router))
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	ctx := r.Context()
	id, ok := ctx.Value(model.CtxUserId).(int)
	if !ok {
		log.Println("error: get id from context")
	}

	clients[id] = ws
}

func messagesWriter(newMessage *Message) {
	messagesChan <- newMessage
}

func chatsWriter(newChat *Chat) {
	chatsChan <- newChat
}

func likesHandler(w http.ResponseWriter, r *http.Request) {}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, ok := ctx.Value(model.CtxUserId).(int)
	if !ok {
		log.Println("error: get id from context")
	}

	var newMessage Message
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newMessage)
	defer r.Body.Close()
	if err != nil {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		responseWithJson(w, 400, response)
		return
	}

	if id != newMessage.AuthorId {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Отказно в доступе"}
		responseWithJson(w, 401, response)
		return
	}

	if len(newMessage.Text) > 250 {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Ошибка валидации"}
		response.Description["text"] = "Слишком длинный текст"
		responseWithJson(w, 400, response)
		return
	}

	newMessage, err = AddMessage(newMessage.AuthorId, newMessage.ChatId, newMessage.Text)
	if err != nil {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: err}
		responseWithJson(w, 400, response)
		return
	}

	go messagesWriter(&newMessage)
}

func webSocketResponse() {
	for {
		newMessage := <-messagesChan
		partnerId, err := GetPartnerId(newMessage.ChatId, newMessage.AuthorId)
		if err != nil {
			//надо залогировать ошибку
			continue
		}

		client := clients[partnerId]
		err = client.WriteJSON(newMessage)
		if err != nil {
			//тут тоже залогировать
			client.Close()
			delete(clients, partnerId)
		}
	}
}
