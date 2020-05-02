package layers

import "github.com/YashdalfTheGray/osi/models"

// PhysicalLayer represents the highest layer of the OSI stack
// It accepts data directly from the user
type PhysicalLayer struct {
	name models.LayerName
	up   models.Layer
	down models.Layer
}

// NewPhysicalLayer creates a new instance of PhysicalLayer
func NewPhysicalLayer(up, down models.Layer) *PhysicalLayer {
	return &PhysicalLayer{models.Presentation, up, down}
}

// Name returns the name of the layer
func (sl PhysicalLayer) Name() models.LayerName {
	return sl.name
}

// SendData is a function that the layer uses to move data up or down
func (sl PhysicalLayer) SendData(to models.Layer, data interface{}) bool {
	return true
}

// ReceiveData is a function that the layer uses to get data from other layers
func (sl PhysicalLayer) ReceiveData(from models.Layer, data interface{}) bool {
	return true
}
