package layers

// Layer defines the basic structure of an OSI layer
type Layer interface {
	GetName() string
	SendData(to *Layer, data interface{}) bool
	ReceiveData(from *Layer, data interface{}) bool
}

// LayerName is used to represent a layer
type LayerName int

// This "enum" is here to make the layer names into constants
const (
	Physical LayerName = iota + 1
	DataLink
	Network
	Transport
	Session
	Presentation
	Application
)
