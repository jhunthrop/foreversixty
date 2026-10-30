package sheet

import (
	"image/color"
	"image/draw"
	"strings"
)

// word is one word of a line that mixes styles: its text, face and colour.
type word struct {
	text string
	t    text
	col  color.RGBA
}

// spanWords splits s into words drawn in t and col. A trailing space in s is kept as the break
// before the next span.
func spanWords(t text, col color.RGBA, s string) []word {
	fields := strings.Fields(s)
	out := make([]word, len(fields))
	for i, f := range fields {
		out[i] = word{text: f, t: t, col: col}
	}
	return out
}

// lineWidth is the device width of words joined by single spaces, each space in the face of
// the word after it.
func lineWidth(words []word) float64 {
	w := 0.0
	for i, wd := range words {
		if i > 0 {
			w += wd.t.width(" ")
		}
		w += wd.t.width(wd.text)
	}
	return w
}

// wrapWords breaks words into lines no wider than maxW device px, greedily.
func wrapWords(words []word, maxW float64) [][]word {
	var lines [][]word
	var line []word
	for _, wd := range words {
		next := append(append([]word(nil), line...), wd)
		if len(line) > 0 && lineWidth(next) > maxW {
			lines = append(lines, line)
			line = []word{wd}
			continue
		}
		line = next
	}
	if len(line) > 0 {
		lines = append(lines, line)
	}
	return lines
}

// drawWords draws one line of mixed words in the box from device x to x+w.
func drawWords(dst draw.Image, words []word, x, w, top float64, a align) {
	switch a {
	case alignRight:
		x += w - lineWidth(words)
	case alignCentre:
		x += (w - lineWidth(words)) / 2
	}
	for i, wd := range words {
		if i > 0 {
			x += wd.t.width(" ")
		}
		wd.t.draw(dst, wd.col, x, top, wd.text)
		x += wd.t.width(wd.text)
	}
}
