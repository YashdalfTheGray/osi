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
