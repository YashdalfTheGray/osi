package layers

import "github.com/YashdalfTheGray/osi/models"

// TransportLayer represents the highest layer of the OSI stack
// It accepts data directly from the user
type TransportLayer struct {
	name models.LayerName
	up   models.Layer
	down models.Layer
}

// NewTransportLayer creates a new instance of TransportLayer
func NewTransportLayer(up, down models.Layer) *TransportLayer {
	return &TransportLayer{models.Presentation, up, down}
}

// Name returns the name of the layer
func (sl TransportLayer) Name() models.LayerName {
	return sl.name
}

// SendData is a function that the layer uses to move data up or down
func (sl TransportLayer) SendData(to models.Layer, data interface{}) bool {
	return true
}

// ReceiveData is a function that the layer uses to get data from other layers
func (sl TransportLayer) ReceiveData(from models.Layer, data interface{}) bool {
	return true
}
