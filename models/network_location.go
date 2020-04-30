package models

import "net/url"

// NetworkLocation contains the raw address of the place that the user
// wants to send a message as well as the fully parsed version of it
type NetworkLocation struct {
	rawAddress string
	URL        *url.URL
}

// NewNetworkLocation constructs a new instance of NetworkLocation.
// This constructor will fail (and panic) if the URL parsing
// is not successful.
func NewNetworkLocation(rawURL string) NetworkLocation {
	parsed, err := url.Parse(rawURL)

	if err != nil {
		panic(err)
	}

	return NetworkLocation{rawAddress: rawURL, URL: parsed}

}
