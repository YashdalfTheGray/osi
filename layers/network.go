package layers

import "github.com/YashdalfTheGray/osi/models"

// NetworkLayer represents the highest layer of the OSI stack
// It accepts data directly from the user
type NetworkLayer struct {
	name models.LayerName
	up   models.Layer
	down models.Layer
}

// NewNetworkLayer creates a new instance of NetworkLayer
func NewNetworkLayer(up, down models.Layer) *NetworkLayer {
	return &NetworkLayer{models.Presentation, up, down}
}

// Name returns the name of the layer
func (sl NetworkLayer) Name() models.LayerName {
	return sl.name
}

// SendData is a function that the layer uses to move data up or down
func (sl NetworkLayer) SendData(to models.Layer, data interface{}) bool {
	return true
}

// ReceiveData is a function that the layer uses to get data from other layers
func (sl NetworkLayer) ReceiveData(from models.Layer, data interface{}) bool {
	return true
}
