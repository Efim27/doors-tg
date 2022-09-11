package requests

import (
	"net/url"
)

func JoinURL(baseURL, add string) (newURL string, err error) {
	base, err := url.Parse(baseURL)
	if err != nil {
		return
	}

	ref, err := url.Parse(add)
	if err != nil {
		return
	}

	newURL = base.ResolveReference(ref).String()
	return
}
