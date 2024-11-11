package httpx

import "net/http"

var _ http.RoundTripper = &UserAgentTransport{}

// NewUserAgentTransport creates a new UserAgentTransport, which wraps a given
// http.RoundTripper and sets a custom User-Agent header for all HTTP requests.
func NewUserAgentTransport(base http.RoundTripper, userAgent string) *UserAgentTransport {
	if base == nil {
		base = http.DefaultTransport
	}
	return &UserAgentTransport{
		Base:      base,
		UserAgent: userAgent,
	}
}

// UserAgentTransport wraps a given http.RoundTripper and sets a custom
// User-Agent header for all HTTP requests.
type UserAgentTransport struct {
	// Base is the base RoundTripper used to make HTTP requests.
	// If nil, http.DefaultTransport is used.
	Base http.RoundTripper

	// UserAgent is the value of the User-Agent header
	UserAgent string
}

// RoundTrip executes a single HTTP transaction and sets the User-Agent header.
//
// Per the http.RoundTripper contract, it clones the provided request
// to ensure the original request is unmodified.
func (u *UserAgentTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r2 := r.Clone(r.Context())

	r2.Header.Set("User-Agent", u.UserAgent)

	base := u.Base
	if u.Base == nil {
		base = http.DefaultTransport
	}
	return base.RoundTrip(r2)
}
