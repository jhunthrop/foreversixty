//go:build !js

package main

// main is a no-op stand-in for the real entrypoint in main.go, which
// only builds under GOOS=js GOARCH=wasm.
//
// Before exports.go existed, this package had no host-buildable file at
// all, so `go build ./...` skipped it silently. exports.go's JSON
// bodies are deliberately host-testable (no build tag), which makes
// `package main` exist on the host for the first time - with no
// func main() there, `go build ./...` fails with "function main is
// undeclared in the main package". This file is that function, so the
// routine host build stays green; it is never what runs in the
// browser.
func main() {}
