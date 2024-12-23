package config

import (
	"net/url"
	"strings"
)

func NormalizeURL(address string) (*url.URL, error) {
	// Trim any leading or trailing whitespace
	address = strings.TrimSpace(address)

	// Parse the URL
	u, err := url.Parse(address)
	if err != nil {
		return nil, err
	}

	// Force https
	u.Scheme = "https"

	// Ensure there's a path
	if u.Path == "" {
		u.Path = "/"
	}

	// Remove username and password if present
	u.User = nil

	// Lowercase the host
	u.Host = strings.ToLower(u.Host)

	return u, nil
}
