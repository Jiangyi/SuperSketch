package chat

type SendType int

const (
	Broadcast SendType = iota // 0
	Emit                   // 1
)

type Point struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Colour int `json:"col"`
}

type Message struct {
	UserId    int    `json:"id"`
	SendType  SendType  `json:"sendType"`
	Message   string `json:"msg"`
	OtherPlayers map[int]string `json:"players,omitempty"`
	State     int    `json:"state"`
	FirstName string `json:"fName,omitempty"`
	LastName  string `json:"lName,omitempty"`
	Point     *Point `json:"point,omitempty"`
}

func CreateMessage(message string, userId int, firstName string, lastName string, sendType SendType) (Message) {
	msg := Message{
		Message: message,
		UserId: userId,
		FirstName: firstName,
		LastName: lastName,
		SendType: sendType,
	}
	return msg
}