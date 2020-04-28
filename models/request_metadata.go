package models

import "net/url"

// RequestMetadata contains the raw address of the place that the user
// wants to send a message as well as the fully parsed version of it
type RequestMetadata struct {
	rawAddress string
	URL url.URL
}
