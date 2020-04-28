package models

import "net/url"

// RequestMetadata contains the raw address of the place that the user
// wants to send a message as well as the fully parsed version of it
type RequestMetadata struct {
	rawAddress string
	URL        *url.URL
}

// NewRequestMetadata constructs a new instance of RequestMetadata.
// This constructor will fail (and panic) if the URL parsing
// is not successful.
func NewRequestMetadata(rawURL string) RequestMetadata {
	parsed, err := url.Parse(rawURL)

	if err != nil {
		panic(err)
	}

	return RequestMetadata{rawAddress: rawURL, URL: parsed}

}
