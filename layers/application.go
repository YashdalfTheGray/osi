package layers

import "fmt"

// ApplicationLayer represents the highest layer of the OSI stack
// It accepts data directly from the user
type ApplicationLayer struct {
	name LayerName
	up   Layer
	down Layer
}

// NewApplicationLayer creates a new instance of ApplicationLayer
func NewApplicationLayer(up, down Layer) ApplicationLayer {
	return ApplicationLayer{Application, up, down}
}

// Name returns the name of the layer
func (al ApplicationLayer) Name() LayerName {
	return al.name
}

// SendData is a function that the layer uses to move data up or down
func (al ApplicationLayer) SendData(to Layer, data interface{}) bool {
	return true
}

// ReceiveData is a function that the layer uses to get data from other layers
func (al ApplicationLayer) ReceiveData(from Layer, data interface{}) bool {
	if from.Name() == Presentation {
		return al.receiveDataFromPresentation(data)
	}
	return al.receiveDataFromUser(data)
}

func (al ApplicationLayer) receiveDataFromUser(data interface{}) bool {
	return al.SendData(al.down, data)
}

func (al ApplicationLayer) receiveDataFromPresentation(data interface{}) bool {
	if str, ok := data.(string); ok {
		fmt.Println(str)
		return true
	}
	fmt.Println("Got malformed data from the presentation layer")
	return false
}
