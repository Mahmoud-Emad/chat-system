package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	initMySQLDB()
	// initRedisDB()

	r := mux.NewRouter()
	chatsRoute := "/applications/{application_token}/chats/"
	messagesRoute := "/applications/{application_token}/chats/{chat_number}/messages/"

	r.HandleFunc(chatsRoute, createChat).Methods("POST")
	r.HandleFunc(messagesRoute, createMessage).Methods("POST")

	http.Handle("/", r)

	fmt.Printf("Golang Server is Listening on port: %d\n", 8001)
	fmt.Printf("Exposed chat route is: %s\n", chatsRoute)
	fmt.Printf("Exposed message route is: %s\n", messagesRoute)
	log.Fatal(http.ListenAndServe(":8001", nil))
}
