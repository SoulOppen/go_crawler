package main

import (
	"fmt"
	"net/url"
)

func normalizeURL(s string) (string, error) {
	out := ""
	u, err := url.Parse(s)
	if err != nil {
		return out, err
	}

	// userinfo
	if u.User != nil {
		out = fmt.Sprint(u.User.String(), "@")
	}

	// host
	if !u.OmitHost && u.Host != "" {
		out = fmt.Sprint(out, u.Host)
	}

	// path
	out = fmt.Sprint(out, u.Path)

	// query
	if u.ForceQuery || u.RawQuery != "" {
		out = fmt.Sprint(out, "?", u.RawQuery)
	}

	// fragment
	if u.Fragment != "" {
		out = fmt.Sprint(out, "#", u.Fragment)
	}
	return out, nil
}
