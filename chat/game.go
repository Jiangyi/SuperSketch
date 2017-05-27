package chat

type Game struct {
	inProgress bool
	turn int // user id specify who's turn it is
	word string // word to guess
}
