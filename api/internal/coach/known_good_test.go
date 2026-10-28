package coach

import "testing"

// TestKnownGood_RealForeverPull runs a real Forever boss pull through Review with the
// shipped rule files (§20.6): the Fury warrior's F1 facts carry the Bloodthirst cast
// count read from the log; the flask holder has no flask finding; an avoidable death
// (when the pull has one) is the F7 finding. The values are pinned with a comment, and a
// change fails until a person explains it (tenet 8).
func TestKnownGood_RealForeverPull(t *testing.T) {
	t.Skip("waits on the owner-provided real Forever pull; its values are pinned from that pull")
}
