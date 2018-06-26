package domain

import (
	"net/url"
	"strings"
)

//Parse extracts host of root from domain or subdomain
func Parse(URL string) (string, error) {
	u, err := url.Parse(URL)

	if err != nil {
		return "", err
	}
	parts := strings.Split(u.Hostname(), ".")

	domain := parts[len(parts)-2] + "." + parts[len(parts)-1]

	return domain, nil
}
