package layers

import "github.com/YashdalfTheGray/osi/models"

// PresentationLayer represents the highest layer of the OSI stack
// It accepts data directly from the user
type PresentationLayer struct {
	name models.LayerName
	up   models.Layer
	down models.Layer
}

// NewPresentationLayer creates a new instance of PresentationLayer
func NewPresentationLayer(up, down models.Layer) *PresentationLayer {
	return &PresentationLayer{models.Presentation, up, down}
}

// Name returns the name of the layer
func (pl PresentationLayer) Name() models.LayerName {
	return pl.name
}

// SendData is a function that the layer uses to move data up or down
func (pl PresentationLayer) SendData(to models.Layer, data interface{}) bool {
	return to.ReceiveData(pl, data)
}

// ReceiveData is a function that the layer uses to get data from other layers
func (pl PresentationLayer) ReceiveData(from models.Layer, data interface{}) bool {
	if from.Name() == models.Application {
		return pl.receiveDataFromApplication(data)
	}
	return pl.receiveDataFromSession(data)
}

func (pl PresentationLayer) receiveDataFromApplication(data interface{}) bool {
	// this is where we would do stuff like encryption and encoding
	// but for the most part, we're just going to make it a 1:1 passthrough
	return pl.SendData(pl.down, data)
}

func (pl PresentationLayer) receiveDataFromSession(data interface{}) bool {
	// similar to receiveDataFromApplication, this is also basically a 1:1
	// passthrough because by the time the data gets to the session layer,
	// things should already be user readable. This is, if we supported it,
	// where we would do decoding and decryption
	return pl.SendData(pl.up, data)
}
