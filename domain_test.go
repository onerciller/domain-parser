package domain

import (
	"testing"
)

func TestParse(t *testing.T) {

	domain, _ := Parse("http://test.com")

	if domain != "test.com" {
		t.Errorf("Expected domain of test.com, but it was %s instead", domain)
	}

	subdomain, _ := Parse("http://subdomain.test.com")

	if subdomain != "test.com" {
		t.Errorf("Expected domain of test.com, but it was %s instead", subdomain)
	}

	w, _ := Parse("http://www.test.com")

	if w != "test.com" {
		t.Errorf("Expected domain of test.com, but it was %s instead", w)
	}

	wSubdomain, _ := Parse("http://www.subdomain.test.com")

	if wSubdomain != "test.com" {
		t.Errorf("Expected domain of test.com, but it was %s instead", wSubdomain)
	}

	m, _ := Parse("http://m.test.com")

	if m != "test.com" {
		t.Errorf("Expected domain of test.com, but it was %s instead", m)
	}

	en, _ := Parse("http://m.test.co.uk")

	if en != "test.co.uk" {
		t.Errorf("Expected domain of test.com, but it was %s instead", en)
	}

}
