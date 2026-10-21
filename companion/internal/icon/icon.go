// companion/internal/icon/icon.go
// Package icon carries the Forever Sixty seal, rendered by
// tools/icons/render.mjs from web/public/favicon.svg and embedded here:
// the tray icon on every platform, the ICO the Windows tray wants, and
// the ICO and ICNS the Windows resource and the macOS bundle are built
// from.
package icon

import (
	"embed"
	"fmt"
)

//go:embed assets/seal-*.png assets/seal.ico assets/seal.icns assets/seal-tray-*.ico
var assets embed.FS

// Sizes are the PNG edges shipped, in pixels.
var Sizes = []int{16, 24, 32, 48, 64, 128, 256, 512}

func read(name string) []byte {
	b, err := assets.ReadFile("assets/" + name)
	if err != nil {
		panic(err) // the pattern above guarantees the file; a test pins the set
	}
	return b
}

// PNG is the 32 px seal, the tray icon on macOS and Linux.
func PNG() []byte { return read("seal-32.png") }

// PNGSize is the seal at one of Sizes, or nil for any other size.
func PNGSize(size int) []byte {
	for _, s := range Sizes {
		if s == size {
			return read(fmt.Sprintf("seal-%d.png", size))
		}
	}
	return nil
}

// ICO is the multi-size icon for Windows (16 to 256 px).
func ICO() []byte { return read("seal.ico") }

// Tones are the headline tones that carry a dot on the tray icon. The
// fourth, off, is the plain seal.
var Tones = []string{"on", "warn", "bad"}

func hasDot(tone string) bool {
	for _, t := range Tones {
		if t == tone {
			return true
		}
	}
	return false
}

// TrayICO is the Windows tray icon for a headline tone: the seal with a
// green (on), amber (warn) or red (bad) dot at its lower right, 16 to 48
// px. Tone off, and any tone this does not know, is the plain seal.
func TrayICO(tone string) []byte {
	if !hasDot(tone) {
		return ICO()
	}
	return read("seal-tray-" + tone + ".ico")
}

// TrayPNG is the same at 32 px for macOS and Linux.
func TrayPNG(tone string) []byte {
	if !hasDot(tone) {
		return PNG()
	}
	return read("seal-tray-" + tone + ".png")
}

// ICNS is the icon for the macOS app bundle (16 to 512 px with @2x).
func ICNS() []byte { return read("seal.icns") }
