package layers

import (
	"fmt"

	"github.com/YashdalfTheGray/osi/models"
)

// ApplicationLayer represents the highest layer of the OSI stack
// It accepts data directly from the user
type ApplicationLayer struct {
	name models.LayerName
	up   models.Layer
	down models.Layer
}

// NewApplicationLayer creates a new instance of ApplicationLayer
func NewApplicationLayer(up, down models.Layer) ApplicationLayer {
	return ApplicationLayer{models.Application, up, down}
}

// Name returns the name of the layer
func (al ApplicationLayer) Name() models.LayerName {
	return al.name
}

// SendData is a function that the layer uses to move data up or down
func (al ApplicationLayer) SendData(to models.Layer, data interface{}) bool {
	return to.ReceiveData(al, data)
}

// ReceiveData is a function that the layer uses to get data from other layers
func (al ApplicationLayer) ReceiveData(from models.Layer, data interface{}) bool {
	if from.Name() == models.Presentation {
		return al.receiveDataFromPresentation(data)
	}
	return al.receiveDataFromUser(data)
}

func (al ApplicationLayer) receiveDataFromUser(data interface{}) bool {
	if message, ok := data.(models.Message); ok {
		return al.SendData(al.down, message.Data)
	}
	fmt.Println("Expected message with string data and string address from the user")
	return false
}

func (al ApplicationLayer) receiveDataFromPresentation(data interface{}) bool {
	if str, ok := data.(string); ok {
		fmt.Println(str)
		return true
	}
	fmt.Println("Got malformed data from the presentation layer")
	return false
}
