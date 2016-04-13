package chat

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"sync"
)

type connInfo struct {
	userID string
	io.Writer
}

type online struct {
	sync.Mutex
	connections map[string]connInfo
	count       int
}

var on online

func (o *online) broadcastCount() {
	cJSON, _ := json.Marshal(connectionsCount{Count: o.count})
	on.broadcast(string(cJSON))
}

func (o *online) add(ws io.Writer) string {
	o.Lock()
	id := connID()
	o.connections[id] = connInfo{"", ws}
	o.count++
	o.Unlock()
	on.broadcastCount()
	return id
}

func (o *online) remove(id string) {
	o.Lock()
	delete(o.connections, id)
	o.count--
	o.broadcastCount()
	o.Unlock()
}

func (o *online) broadcast(msg string) {
	for _, ws := range o.connections {
		ws.Write([]byte(msg))
	}

}

func connID() string {
	b := make([]byte, 64)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}
	return fmt.Sprintf("%x", md5.Sum(b))
}

func init() {
	on.connections = make(map[string]connInfo)
}
