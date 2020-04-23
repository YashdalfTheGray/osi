package layers

// PresentationLayer represents the highest layer of the OSI stack
// It accepts data directly from the user
type PresentationLayer struct {
	name LayerName
	up   *Layer
	down *Layer
}

// NewPresentationLayer creates a new instance of PresentationLayer
func NewPresentationLayer(up, down *Layer) *PresentationLayer {
	return &PresentationLayer{Presentation, up, down}
}

// GetName returns the name of the layer
func (ap PresentationLayer) GetName() string {
	return string(ap.name)
}

// SendData is a function that the layer uses to move data up or down
func (ap *PresentationLayer) SendData(to *Layer, data interface{}) bool {
	return true
}

// ReceiveData is a function that the layer uses to get data from other layers
func (ap *PresentationLayer) ReceiveData(from *Layer, data interface{}) bool {
	return true
}
