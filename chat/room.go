package chat

import (
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	"log"
	"sync"
)

var (
	mutex = &sync.Mutex{}
	rw = &sync.RWMutex{}
)

type Room struct {
	name    string
	clients map[int]*Client
	uuid    uuid.UUID
	count   int

	game    Game
}

func (r *Room) AddConn(conn *websocket.Conn, userId int) {
	client := NewClient(conn, userId)
	r.clients[userId] = client
	r.count++
	otherPlayers := make(map[int]string)
	for id, client := range r.clients {
		if id != userId {
			otherPlayers[id] = client.user.FirstName
		}
	}
	r.SendTo(userId, Message{
		Message: "You have joined the room",
		UserId: userId,
		OtherPlayers: otherPlayers,
		SendType: Broadcast,
		State: 1,
		FirstName: client.user.FirstName,
		LastName: client.user.LastName,
	})
	r.SendEmit(userId, Message{
		Message: client.user.FirstName + " has joined the room",
		UserId: userId,
		SendType: Emit,
		State: 1,
		FirstName: client.user.FirstName,
		LastName: client.user.LastName,
	})
	log.Printf("New client joined room '%s'", r.name)
}

func (r *Room) RemoveConn(id int) {
	delete(r.clients, id)
	r.count--
	log.Printf("Client left room '%s'", r.name)
}

func (r *Room) IsCloseable() bool {
	return r.count == 0
}

// Send message to only a single client
func (r *Room) SendTo(targetId int, msg Message) {
	r.clients[targetId].WriteMessage(msg)
}

// Send message to all, including sender
func (r *Room) SendBroadcast(msg Message) {
	for _, client := range r.clients {
		client.WriteMessage(msg)
	}
}

// Send message to all, excluding sender
func (r *Room) SendEmit(senderId int, msg Message) {
	for id, client := range r.clients {
		if id != senderId {
			client.WriteMessage(msg)
		}
	}
}

// Goroutine message handler
func (r *Room) HandleMsg(id int) {
	for {
		if r.clients[id] == nil {
			break
		}
		out := <-r.clients[id].out
		if out.SendType == Broadcast {
			r.SendBroadcast(out)
		} else {
			r.SendEmit(id, out)
		}
	}
}

func (r *Room) Close() {
	for k := range r.clients {
		delete(r.clients, k)
	}
}

func NewRoom(name string) *Room {
	room := &Room{
		name:    name,
		clients: make(map[int]*Client),
		uuid:    uuid.NewV4(),
		count:   0,
		game:    Game{
			inProgress: false,
		},
	}
	return room
}