package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/GeorgeHN666/go-websocket/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Handle() {

	router := mux.NewRouter()

	router.HandleFunc("/ws", routers.WebsocketConn)

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	h := cors.AllowAll().Handler(router)

	log.Printf("Server Listening in PORT:%s", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, h))

}
