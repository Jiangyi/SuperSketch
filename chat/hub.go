package chat

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"github.com/kvnxiao/supersketch/auth"
	"strings"
	"strconv"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Hub struct {
	lobby map[string]*Room
}

func (h *Hub) CloseRoom(roomId string) {
	delete(h.lobby, roomId)
}

func (h *Hub) GetRoom(roomId string) *Room {
	if _, ok := h.lobby[roomId]; !ok {
		h.lobby[roomId] = NewRoom(roomId)
	}
	return h.lobby[roomId]
}

func (h *Hub) HandleWs(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(auth.SESSION_ID); if err != nil {
		auth.SendFlashMessage(w, r, err.Error(), "/login")
	}
	decoded, err := auth.DecodeSessionID(cookie); if err != nil {
		auth.SendFlashMessage(w, r, err.Error(), "/login")
	}
	log.Println("Original cookie: " + cookie.Value)
	log.Println("Decoded cookie: " + decoded)
	idString := decoded[strings.Index(decoded, ":") + 1:]
	log.Println("ID: " + idString)
	id, err := strconv.Atoi(idString); if err != nil {
		auth.SendFlashMessage(w, r, err.Error(), "/login")
	}

	params := r.URL.Query()
	roomId := params.Get("rid")
	if roomId == "" {
		h.handleLobbyWs(w, r, id)
	} else {
		h.handleRoomWs(w, r, roomId, id)
	}
}

func (h *Hub) handleLobbyWs(w http.ResponseWriter, r *http.Request, userId int) {

}

func (h *Hub) handleRoomWs(w http.ResponseWriter, r *http.Request, roomId string, userId int) {
	log.Println("HANDLING ROOM WEBSOCKET")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("err")
		return
	}

	// Handle websocket connection
	defer conn.Close()
	room := h.GetRoom(roomId)

	room.AddConn(conn, userId)

	// Creates a new goroutine for each websocket connection
	// Handles messages sent and received for this connection and its associated room
	go room.HandleMsg(userId)

	log.Println(room.clients[userId])
	room.clients[userId].ReadLoop()
	room.RemoveConn(userId)
	if room.IsCloseable() {
		room.Close()
		h.CloseRoom(roomId)
	}
}

func NewHub() *Hub {
	hub := &Hub{
		lobby: make(map[string]*Room),
	}
	return hub
}
