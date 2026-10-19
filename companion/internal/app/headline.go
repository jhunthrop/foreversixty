// companion/internal/app/headline.go
package app

// Headline is the one line the Status page opens with, and the tray
// repeats: "is my raid going to be uploaded?" in words and a colour.
type Headline struct {
	// Key names the state: connecting, not_connected, no_folder,
	// advanced_off, ready or logging.
	Key string `json:"key"`
	// Tone is the dot's colour: bad (red), warn (amber), off (grey)
	// or on (green).
	Tone  string `json:"tone"`
	Title string `json:"title"`
	Sub   string `json:"sub"`
	// Action is the one optional button.
	Action *HeadlineAction `json:"action,omitempty"`
}

// HeadlineAction is a button that goes to another page.
type HeadlineAction struct {
	Label string `json:"label"`
	Page  string `json:"page"`
}

// headlineFor picks the state, first match wins: waiting on a
// sign-in, not paired, no game folder, any install with advanced
// logging off, paired and not logging, logging.
func headlineFor(paired, connecting, logging bool, detected string, installs []InstallStatus) Headline {
	switch {
	case connecting && !paired:
		return Headline{Key: "connecting", Tone: "off", Title: CopyConnectingTitle,
			Sub:    CopyConnectingSub,
			Action: &HeadlineAction{Label: CopyConnectingBtn, Page: "device"}}
	case !paired:
		return Headline{Key: "not_connected", Tone: "bad", Title: CopyNotConnectedTitle,
			Sub:    CopyNotConnectedSub,
			Action: &HeadlineAction{Label: CopyNotConnectedBtn, Page: "device"}}
	case len(installs) == 0:
		return Headline{Key: "no_folder", Tone: "bad", Title: CopyNoFolderTitle,
			Sub:    CopyNoFolderSub,
			Action: &HeadlineAction{Label: CopyNoFolderBtn, Page: "settings"}}
	}
	for _, in := range installs {
		if in.AdvancedKnown && !in.Advanced {
			return Headline{Key: "advanced_off", Tone: "warn",
				Title: CopyAdvancedOffTitle, Sub: CopyAdvancedOffSub}
		}
	}
	if logging {
		sub := CopyLoggingSub
		if detected != "" {
			sub = CopyLoggingSubAs(detected)
		}
		return Headline{Key: "logging", Tone: "on", Title: CopyLoggingTitle, Sub: sub}
	}
	return Headline{Key: "ready", Tone: "off", Title: CopyReadyTitle, Sub: CopyReadySub}
}
