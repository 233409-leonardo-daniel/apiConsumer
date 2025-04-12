package repositories

import (
	"net/http"
)

type IWebSocket interface {
	HandleConnections(w http.ResponseWriter, r *http.Request)
	HandleMessages()
	BroadcastMessage(message string)
}
