package layers

// Layer defines the basic structure of an OSI layer
type Layer interface {
	GetName() string
	SendData(to *Layer, data interface{}) bool
	ReceiveData(from *Layer, data interface{}) bool
}

// LayerName is used to represent a layer
type LayerName string

// This "enum" is here to make the layer names into constants
const (
	Physical     LayerName = "Physical"
	DataLink     LayerName = "DataLink"
	Network      LayerName = "Network"
	Transport    LayerName = "Transport"
	Session      LayerName = "Session"
	Presentation LayerName = "Presentation"
	Application  LayerName = "Application"
)
