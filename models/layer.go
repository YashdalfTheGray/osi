package models

// Layer defines the basic structure of an OSI layer
type Layer interface {
	Name() LayerName
	SendData(to Layer, message Message) bool
	ReceiveData(from Layer, message Message) bool
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
