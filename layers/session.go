package layers

import "github.com/YashdalfTheGray/osi/models"

// SessionLayer represents the highest layer of the OSI stack
// It accepts data directly from the user
type SessionLayer struct {
	name models.LayerName
	up   models.Layer
	down models.Layer
}

// NewSessionLayer creates a new instance of SessionLayer
func NewSessionLayer(up, down models.Layer) *SessionLayer {
	return &SessionLayer{models.Presentation, up, down}
}

// Name returns the name of the layer
func (sl SessionLayer) Name() models.LayerName {
	return sl.name
}

// SendData is a function that the layer uses to move data up or down
func (sl SessionLayer) SendData(to models.Layer, data interface{}) bool {
	return to.ReceiveData(sl, data)
}

// ReceiveData is a function that the layer uses to get data from other layers
func (sl SessionLayer) ReceiveData(from models.Layer, data interface{}) bool {
	if from.Name() == models.Presentation {
		return sl.receiveDataFromPresentation(data)
	}
	return sl.receiveDataFromTransport(data)
}

func (sl SessionLayer) receiveDataFromPresentation(data interface{}) bool {
	return true
}

func (sl SessionLayer) receiveDataFromTransport(data interface{}) bool {
	return true
}
