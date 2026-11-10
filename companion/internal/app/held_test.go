package app

import (
	"testing"

	"github.com/jhunthrop/foreversixty/companion/internal/pipeline"
)

// A held queue turns Logging amber, headline and tray alike; a title that
// already names a problem keeps it, and nothing but Logging changes.
func TestAHeldQueueTurnsTheLoggingHeadlineAmber(t *testing.T) {
	held := &pipeline.Held{Kind: "fight", Fight: 4, Attempts: 2}
	logging := Headline{Key: "logging", Tone: "on", Title: CopyLoggingTitle}
	if got := withHeld(logging, held); got.Tone != "warn" || got.Title != CopyLoggingTitleHeld {
		t.Errorf("held logging = %+v", got)
	}
	if got := withHeld(logging, nil); got != logging {
		t.Errorf("unheld logging changed: %+v", got)
	}
	notCredited := Headline{Key: "logging", Tone: "warn", Title: CopyLoggingTitleNotCredited}
	if got := withHeld(notCredited, held); got.Title != CopyLoggingTitleNotCredited || got.Tone != "warn" {
		t.Errorf("held, not credited = %+v", got)
	}
	ready := Headline{Key: "ready", Tone: "off", Title: CopyReadyTitle}
	if got := withHeld(ready, held); got != ready {
		t.Errorf("held ready changed: %+v", got)
	}
}
