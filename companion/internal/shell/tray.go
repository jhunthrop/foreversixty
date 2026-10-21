// companion/internal/shell/tray.go
// What the tray says about the companion, and when it is rewritten. The
// words are the Status headline's title, word for word; the icon is the
// seal with a dot in the headline's tone. Kept apart from the systray
// calls (shell_gui.go) so the mapping is testable without a tray.
package shell

import "unicode/utf8"

// TooltipMax is where Windows cuts a tray tooltip off: 127 characters.
const TooltipMax = 127

// tooltipPrefix leads the tooltip; the menu's first line is the title alone.
const tooltipPrefix = "Forever Sixty · "

// trayView is everything the tray shows that can change: the headline's
// title (the tooltip and the menu's first line) and its tone (the icon).
type trayView struct{ Line, Tone string }

// tooltip is "Forever Sixty · Logging", cut to what Windows will show.
func tooltip(line string) string {
	t := tooltipPrefix + line
	if utf8.RuneCountInString(t) <= TooltipMax {
		return t
	}
	r := []rune(t)
	return string(r[:TooltipMax])
}

// trayApply is what the tray does with a change.
type trayApply struct {
	Tooltip func(string)
	Title   func(string)
	Icon    func(tone string)
}

// applyTray brings the tray from prev to now, calling only what changed:
// the words are rewritten when they differ and the icon only when the
// tone does, so a two-second poll that finds nothing new touches nothing.
// A nil prev is the first call, which sets everything.
func applyTray(prev *trayView, now trayView, a trayApply) {
	if prev == nil || prev.Line != now.Line {
		a.Tooltip(tooltip(now.Line))
		a.Title(now.Line)
	}
	if prev == nil || prev.Tone != now.Tone {
		a.Icon(now.Tone)
	}
}
