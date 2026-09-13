// Package site server-renders the shared build page and its preview card.
package site

import (
	_ "embed"
	"html/template"
	"strings"
)

// mainMarker separates the header from the footer in chrome.html; the page
// content goes where it sits.
const mainMarker = "<!--MAIN-->"

//go:embed chrome.html
var chromeHTML string

// Header and Footer are the site's own header and footer markup. They are a
// copy, kept honest by chrome_test.go, which fails when the site's chrome
// changes and the copy does not.
var Header, Footer = splitChrome(chromeHTML)

// splitChrome cuts a chrome document into its header and footer halves. A
// chrome file without the marker is a build-time mistake in an embedded
// asset, not a runtime condition, so it panics rather than degrading.
func splitChrome(s string) (template.HTML, template.HTML) {
	header, footer, ok := strings.Cut(s, mainMarker)
	if !ok {
		panic("site: chrome.html must contain " + mainMarker + " between the header and the footer")
	}
	return template.HTML(strings.TrimSpace(header)), template.HTML(strings.TrimSpace(footer))
}
