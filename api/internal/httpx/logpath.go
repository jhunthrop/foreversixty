package httpx

import (
	"path"
	"strings"
)

// secretPathPrefixes are the routes whose last path segment is a bearer secret: anyone holding
// the full path can read what it guards, so a log line must never carry it.
var secretPathPrefixes = []string{
	// The calendar link (guilds.CalendarFeedPrefix): the path is the token.
	"/v1/calendar/",
}

// redactedSegment stands in for a secret path segment in the logs.
const redactedSegment = ":redacted"

// LogPath is p as a log line may carry it: a secret route's remainder replaced. The loggers
// run before the mux cleans a path (and answers a non-canonical one with a redirect), so the
// match is on the cleaned path and anywhere in it: "//v1/calendar/x", "/v1/./calendar/x" and
// "/a/../v1/calendar/x" are all redacted.
func LogPath(p string) string {
	cleaned := path.Clean("/" + p)
	for _, prefix := range secretPathPrefixes {
		if i := strings.Index(cleaned, prefix); i >= 0 && len(cleaned) > i+len(prefix) {
			return cleaned[:i] + prefix + redactedSegment
		}
		if i := strings.Index(p, prefix); i >= 0 && len(p) > i+len(prefix) {
			return p[:i] + prefix + redactedSegment
		}
	}
	return p
}
