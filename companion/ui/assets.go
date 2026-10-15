// companion/ui/assets.go
// Package ui is the companion's window: a page, its script, its words
// (copy.js), its helpers (format.js) and its stylesheet, embedded in the
// binary and served over loopback to a native webview. It cannot
// import the site's stylesheet, so the design tokens it needs are
// copied into styles.css and kept in step with
// web/src/styles/tokens.css by hand.
package ui

import "embed"

// Assets is index.html and what it loads.
//
//go:embed index.html app.js copy.js format.js styles.css
var Assets embed.FS
