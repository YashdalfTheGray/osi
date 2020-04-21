package layers

// Layer defines the basic structure of an OSI layer
type Layer interface {
	GetName() string
	SendData(to *Layer, data interface{}) bool
	ReceiveData(from *Layer, data interface{}) bool
}
