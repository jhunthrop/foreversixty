// api/internal/bnetapi/static.go
package bnetapi

import "regexp"

// staticNamespaceRe matches a versioned static namespace inside any link a
// Blizzard response carries, e.g. "namespace=static-1.60.1_70399-forever-us".
var staticNamespaceRe = regexp.MustCompile(`namespace=(static-[0-9][0-9._]*-[a-z0-9]+-[a-z]{2})\b`)

// StaticNamespace discovers the versioned static namespace from a response
// body's own links. Forever's static namespace is versioned and the plain
// "static-forever-{region}" answers 403 (probed 2026-11-03), and the version
// Blizzard serves lags the client build (1.60.1_70755 answered 403 while
// 1.60.1_70399 answered 200), so it is never derived from the client build:
// every profile response links its class, race and items under the
// namespace Blizzard is serving, and this reads it from there. Empty when
// the body links no static namespace.
//
// Static data is validation only: the client-derived tables stay the
// source of truth for the pipeline.
func StaticNamespace(body []byte) string {
	m := staticNamespaceRe.FindSubmatch(body)
	if m == nil {
		return ""
	}
	return string(m[1])
}
