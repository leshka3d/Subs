package messages

// Message - sublitle message
type Message struct {
	ID      string
	Time    string
	Message string
}

//MessageMan - rule the messages
type MessageMan struct {
	mes     []Message
	current int
	end     int
}
