package pflagx

import (
	"net/url"
)

// URL allows a url.URL to be added as a pflag.
type URL struct {
	*url.URL
}

func (u *URL) String() string {
	if u.URL == nil {
		return ""
	}
	return u.URL.String()
}

func (u *URL) Set(s string) error {
	newURL, err := url.Parse(s)
	if err != nil {
		return err
	}
	u.URL = newURL
	return nil
}

func (*URL) Type() string {
	return "string"
}
