package httpx

import "strings"

// secretPathPrefixes are the routes whose last path segment is a bearer secret: anyone holding
// the full path can read what it guards, so a log line must never carry it.
var secretPathPrefixes = []string{
	// The calendar link (guilds.CalendarFeedPrefix): the path is the token.
	"/v1/calendar/",
}

// redactedSegment stands in for a secret path segment in the logs.
const redactedSegment = ":redacted"

// LogPath is path as a log line may carry it: a secret route's remainder replaced.
func LogPath(path string) string {
	for _, prefix := range secretPathPrefixes {
		if strings.HasPrefix(path, prefix) && len(path) > len(prefix) {
			return prefix + redactedSegment
		}
	}
	return path
}
