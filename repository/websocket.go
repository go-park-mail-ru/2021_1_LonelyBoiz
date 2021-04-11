package repository

/*import (
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

type Resp struct {
	ResponseType string      `json:"type"`
	Object       interface{} `json:"obj"`
}

func main() {
	router := mux.NewRouter()
	//надо обернуть в проверку куки
	router.HandleFunc("/chats/{chatId:[0-9]+}/messages", messageHandler).Methods("POST")
	router.HandleFunc("/likes", likesHandler).Methods("POST")
	router.HandleFunc("/ws", wsHandler).Methods("Get")

	go webSocketMessageResponse()
	go webSocketChatResponse()

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

func likesHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, ok := ctx.Value(model.CtxUserId).(int)
	if !ok {
		//ответить что сервер не смог взять адйишникиз контекста
		log.Println("error: get id from context")
	}

	var like Like
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err := decoder.Decode(&like)
	if err != nil {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		responseWithJson(w, 400, response)
		return
	}

	if like.Reaction != "like" || like.Reaction != "skip" {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		response.Description["like"] = "неправильный формат реацкции ожидается skip или like"
		responseWithJson(w, 400, response)
		return
	}

	rowsAffected, err := DB.Rating(id, like.UserId, like.Reaction)
	if err != nil {
		//залогировать ошибку
		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		responseWithJson(w, 500, response)
		return
	}
	if rowsAffected == -1 {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Отказано в доступе"}
		response.Description["userID"] = "Пытаешься поставить лайк человеку не со своей ленты"
		responseWithJson(w, 403, response)
		return
	}

	reciprocity, err := DB.CheckReciprocity(like.UserId, id)
	if err != nil {
		//залогировать ошибку
		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Не удалось проверить взаимность"}
		responseWithJson(w, 500, response)
		return
	}
	if reciprocity == false {
		w.WriteHeader(204)
		return
	}

	var newChat Chat
	newChat.ChatId, err = DB.CreateChat(id, like.UserId)
	if err != nil {
		//залогировать ошибку
		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Не удалось создать чат"}
		responseWithJson(w, 500, response)
		return
	}

	newChat, err = DB.GetNewChat(newChat.ChatId, like.UserId)
	if err != nil {
		//залогировать ошибку
		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Не удалось получить новый чат из базы"}
		responseWithJson(w, 500, response)
		return
	}

	go chatsWriter(&newChat)

	responseWithJson(w, 200, newChat)
}

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

	newMessage, err = DB.AddMessage(newMessage.AuthorId, newMessage.ChatId, newMessage.Text)
	if err != nil {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: err}
		responseWithJson(w, 400, response)
		return
	}

	go messagesWriter(&newMessage)

	responseWithJson(w, 200, newMessage)
}

func webSocketMessageResponse() {
	for {
		newMessage := <-messagesChan
		partnerId, err := DB.GetPartnerId(newMessage.ChatId, newMessage.AuthorId)
		if err != nil {
			//надо залогировать ошибку
			continue
		}

		response := Resp{ResponseType: "message", Object: newMessage}

		client, ok := clients[partnerId]
		if !ok {
			client.Close()
			delete(clients, partnerId)
			continue
		}

		err = client.WriteJSON(response)
		if err != nil {
			//тут тоже залогировать
			client.Close()
			delete(clients, partnerId)
		}
	}
}

func webSocketChatResponse() {
	for {
		newChat := <-chatsChan
		partnerId := newChat.PartnerId

		response := Resp{ResponseType: "chat", Object: newChat}

		client, ok := clients[partnerId]
		if !ok {
			client.Close()
			delete(clients, partnerId)
			continue
		}

		err := client.WriteJSON(response)
		if err != nil {
			//тут тоже залогировать
			client.Close()
			delete(clients, partnerId)
		}
	}
}
*/
