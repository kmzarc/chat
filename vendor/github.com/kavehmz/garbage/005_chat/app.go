package main

import (
	"log"
	"net/http"

	"github.com/kavehmz/garbage/005_chat/chat"
	"golang.org/x/net/websocket"
)

func main() {
	http.HandleFunc("/githubcb", chat.GithubCallback)
	http.Handle("/chat", websocket.Handler(chat.Chat))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
