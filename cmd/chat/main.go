package main

import (
	"log"
	"net/http"
	"os"

	"golang.org/x/net/websocket"

	"github.com/kavehmz/garbage/005_chat/chat"
)

func main() {
	http.HandleFunc("/githubcb", chat.GithubCallback)
	http.Handle("/chat", websocket.Handler(chat.Chat))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
