package chat

import (
	"github.com/gorilla/websocket"
	"log"
	"github.com/kvnxiao/supersketch/db"
)

type User struct {
	ID        int
	Username  string
	PassHash  string
	FirstName string
	LastName  string
}

type Client struct {
	// Client id
	user User
	// Websocket connection
	conn *websocket.Conn
	// Send / output channel for writing messages from server to client
	out  chan Message
}

func NewClient(conn *websocket.Conn, userId int) *Client {
	username, firstName, lastName, err := db.GetUserByID(userId); if err != nil {
		return nil
	}

	user := &User{
		ID: userId,
		Username: username,
		FirstName: firstName,
		LastName: lastName,
	}
	client := &Client{
		conn: conn,
		out:  make(chan Message),
		user: *user,
	}
	return client
}

// Reads from websocket connection and pumps to out channel
func (c *Client) ReadLoop() {
	defer close(c.out)
	for {
		m := Message{}
		err := c.conn.ReadJSON(&m)
		if err != nil {
			break
		}
		message := Message{
			UserId: m.UserId,
			SendType: Broadcast,
			Message:   m.Message,
		}
		c.out <- message
	}
	log.Println("Read loop done")
}

func (c *Client) WriteMessage(msg interface{}) {
	log.Println(msg)
	mutex.Lock()
	defer mutex.Unlock()
	err := c.conn.WriteJSON(msg)
	if err != nil {
		log.Println("Error in writing message!", err)
	}
}
