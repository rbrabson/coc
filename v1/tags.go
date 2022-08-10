package coc

import "net/url"

// fmtTag formats the tag for use in a URL or in a query parameter.
func fmtTag(tag string) string {
	// If the tag doesn't have a '#' character at the front, add one
	if len(tag) == 0 {
		return tag
	}
	if tag[0] == '#' {
		return url.QueryEscape(tag)
	}
	return url.QueryEscape("#" + tag)

}
