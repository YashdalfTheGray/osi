package layers

// ApplicationLayer represents the highest layer of the OSI stack
// It accepts data directly from the user
type ApplicationLayer struct {
	name string
	up   *Layer
	down *Layer
}

// NewApplicationLayer creates a new instance of ApplicationLayer
func NewApplicationLayer(up, down *Layer) *ApplicationLayer {
	return &ApplicationLayer{"Application", up, down}
}

// GetName returns the name of the layer
func (ap ApplicationLayer) GetName() string {
	return ap.name
}

// SendData is a function that the layer uses to move data up or down
func (ap *ApplicationLayer) SendData(to *Layer, data interface{}) bool {
	return true
}

// ReceiveData is a function that the layer uses to get data from other layers
func (ap *ApplicationLayer) ReceiveData(from *Layer, data interface{}) bool {
	return true
}
