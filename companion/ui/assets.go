// companion/ui/assets.go
// Package ui is the companion's window: three files, embedded in the
// binary and served over loopback to a native webview. It cannot
// import the site's stylesheet, so the design tokens it needs are
// copied into styles.css and kept in step with
// web/src/styles/tokens.css by hand.
package ui

import "embed"

// Assets is index.html and what it loads.
//
//go:embed index.html app.js styles.css
var Assets embed.FS
