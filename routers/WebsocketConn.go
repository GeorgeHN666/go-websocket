package routers

import (
	"log"
	"net/http"

	"github.com/GeorgeHN666/go-websocket/models"
	"github.com/gorilla/websocket"
)

var Clients = make(map[*websocket.Conn]bool)

var BroadCast = make(chan models.Message)

var Upgrader = websocket.Upgrader{}

func WebsocketConn(w http.ResponseWriter, r *http.Request) {

	ws, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()

	Clients[ws] = true

	for {

		var msg models.Message

		// We Read a new Message that comes like JSON and then we map it to a Message Object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error::%v", err)
			delete(Clients, ws)
			break
		}

		// We send the information of the message to the broadcaster
		BroadCast <- msg
	}

	go ListenMessages()

}

func ListenMessages() {
	for {
		// Grab the next message for the Broadcast Channel
		msg := <-BroadCast

		for client := range Clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error::%v", err)
				client.Close()
				delete(Clients, client)
			}
		}

	}
}
