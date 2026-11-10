// companion/internal/icon/icon.go
// Package icon carries the tray's Forever Sixty seal v4, rendered by
// tools/icons/render.py from design/assets/logo/lx-seal-v4*.png. Each
// platform's binary embeds only its own tray images (tray_windows.go:
// the ICOs; tray_other.go: the PNGs). The app icons live beside them in
// assets/ and are never embedded: seal.ico goes into the Windows exe as a
// resource (cmd/foreversixty-companion/winres) and seal.icns into the
// macOS bundle (tools/package-mac-companion.sh).
package icon

import "fmt"

// Tones are the headline tones that carry a dot on the tray icon. The
// fourth, off, is the plain seal.
var Tones = []string{"on", "warn", "bad"}

// Off is the tone of the plain seal, and any tone this does not know.
const Off = "off"

// toneOf is tone when it carries a dot, else Off.
func toneOf(tone string) string {
	for _, t := range Tones {
		if t == tone {
			return t
		}
	}
	return Off
}

// Tray is the tray icon for a headline tone: the seal, with a green (on),
// amber (warn) or red (bad) dot at its lower right. On Windows it is an
// ICO at 16 to 48 px; on macOS and Linux a 32 px PNG (16 pt at 2x in the
// menu bar), in colour: a template image is drawn from its alpha alone
// and would erase both the seal and the dot.
func Tray(tone string) []byte {
	b, err := tray.ReadFile(fmt.Sprintf("assets/seal-tray-%s.%s", toneOf(tone), trayExt))
	if err != nil {
		panic(err) // the embed pattern guarantees every tone; a test pins the set
	}
	return b
}
