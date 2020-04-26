package layers

// SessionLayer represents the highest layer of the OSI stack
// It accepts data directly from the user
type SessionLayer struct {
	name LayerName
	up   Layer
	down Layer
}

// NewSessionLayer creates a new instance of SessionLayer
func NewSessionLayer(up, down Layer) *SessionLayer {
	return &SessionLayer{Presentation, up, down}
}

// Name returns the name of the layer
func (sl SessionLayer) Name() LayerName {
	return sl.name
}

// SendData is a function that the layer uses to move data up or down
func (sl SessionLayer) SendData(to Layer, data interface{}) bool {
	return to.ReceiveData(sl, data)
}

// ReceiveData is a function that the layer uses to get data from other layers
func (sl SessionLayer) ReceiveData(from Layer, data interface{}) bool {
	return true
}
