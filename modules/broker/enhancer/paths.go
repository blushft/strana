package enhancer

import (
	"net/url"
	"strings"
)

func cleanPathname(p string) string {
	return "/" + strings.TrimLeft(p, "/")
}

func cleanHostname(h string) string {
	u, err := url.Parse(h)
	if err != nil {
		return ""
	}

	return u.Scheme + "://" + u.Host
}
