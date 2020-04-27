package layers

// PresentationLayer represents the highest layer of the OSI stack
// It accepts data directly from the user
type PresentationLayer struct {
	name LayerName
	up   Layer
	down Layer
}

// NewPresentationLayer creates a new instance of PresentationLayer
func NewPresentationLayer(up, down Layer) *PresentationLayer {
	return &PresentationLayer{Presentation, up, down}
}

// Name returns the name of the layer
func (pl PresentationLayer) Name() LayerName {
	return pl.name
}

// SendData is a function that the layer uses to move data up or down
func (pl PresentationLayer) SendData(to Layer, data interface{}) bool {
	return to.ReceiveData(pl, data)
}

// ReceiveData is a function that the layer uses to get data from other layers
func (pl PresentationLayer) ReceiveData(from Layer, data interface{}) bool {
	if from.Name() == Application {
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
	return true
}
