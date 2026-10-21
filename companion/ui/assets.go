// companion/ui/assets.go
// Package ui is the companion's window: a page, its script, its words
// (copy.js), its helpers (format.js), its stylesheet, the seal and the
// site's typefaces, embedded in the binary and served over loopback to
// a native webview. It cannot import the site's stylesheet, so the
// design tokens it needs are copied into styles.css and kept in step
// with web/src/styles/tokens.css by hand; the fonts are the site's own
// subset files, copied byte for byte (fonts/README.md).
package ui

import (
	"embed"
	"mime"
)

// Assets is index.html and what it loads.
//
//go:embed index.html app.js copy.js format.js styles.css seal.svg fonts/*.woff2
var Assets embed.FS

// Go's built-in table of extensions has no .woff2, and on Windows the
// registry is consulted too and may not either; without a type the
// server would sniff the font as application/octet-stream and the
// webview would refuse to preload it.
func init() {
	_ = mime.AddExtensionType(".woff2", "font/woff2")
}
