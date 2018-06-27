package domain

import (
	tld "github.com/jpillora/go-tld"
)

//Parse extracts host of root from domain or subdomain
func Parse(URL string) (string, error) {

	u, err := tld.Parse(URL)

	if err != nil {
		return "", err
	}

	domain := u.Domain + "." + u.TLD

	return domain, nil

}
