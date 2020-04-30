package models

import "fmt"

// Message contains some data that needs to be sent to a
// location on the network
type Message struct {
	Data    interface{}
	Address NetworkLocation
}

// NewMessage constructs a new instance of Message
func NewMessage(data interface{}, address string) Message {
	loc := NewNetworkLocation(address)
	return Message{Data: data, Address: loc}
}

func (m Message) String() string {
	return fmt.Sprintf("Message to %s\n with content\n%s", m.Address.rawAddress, m.Data)
}
