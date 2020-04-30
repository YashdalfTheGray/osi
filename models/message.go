package models

// Message contains some data that needs to be sent to a
// location on the network
type Message struct {
	data    interface{}
	address NetworkLocation
}

// NewMessage constructs a new instance of Message
func NewMessage(data interface{}, address string) Message {
	loc := NewNetworkLocation(address)
	return Message{data: data, address: loc}
}
