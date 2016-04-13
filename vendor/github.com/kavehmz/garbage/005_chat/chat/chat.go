package chat

import (
	"encoding/json"
	"fmt"

	"golang.org/x/net/websocket"
	"golang.org/x/oauth2"
)

type outMessage struct {
	User    string `json:"user"`
	Message string `json:"message"`
}

type inMessage struct {
	Message string `json:"message"`
}

type connectionsCount struct {
	Count int `json:"count"`
}

// Chat will handle all WS requests
func Chat(ws *websocket.Conn) {
	connID := on.add(ws)
	defer func() {
		on.remove(connID)
	}()
	url, _ := json.Marshal(oauthURL{URL: oauthConf.AuthCodeURL(connID, oauth2.AccessTypeOnline)})
	ws.Write(url)
	request := make([]byte, 1000)
	for {
		n, err := ws.Read(request)
		if err != nil {
			fmt.Println("Error", err.Error())
			return
		}

		var in outMessage
		if err := json.Unmarshal(request[:n], &in); err != nil {
			panic(err)
		}

		msg := outMessage{Message: in.Message, User: on.connections[connID].userID}
		chatJSON, _ := json.Marshal(msg)
		on.broadcast(string(chatJSON))
	}

}
