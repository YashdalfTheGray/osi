package layers

import "github.com/YashdalfTheGray/osi/models"

// DataLinkLayer represents the highest layer of the OSI stack
// It accepts data directly from the user
type DataLinkLayer struct {
	name models.LayerName
	up   models.Layer
	down models.Layer
}

// NewDataLinkLayer creates a new instance of DataLinkLayer
func NewDataLinkLayer(up, down models.Layer) *DataLinkLayer {
	return &DataLinkLayer{models.Presentation, up, down}
}

// Name returns the name of the layer
func (sl DataLinkLayer) Name() models.LayerName {
	return sl.name
}

// SendData is a function that the layer uses to move data up or down
func (sl DataLinkLayer) SendData(to models.Layer, data interface{}) bool {
	return true
}

// ReceiveData is a function that the layer uses to get data from other layers
func (sl DataLinkLayer) ReceiveData(from models.Layer, data interface{}) bool {
	return true
}
