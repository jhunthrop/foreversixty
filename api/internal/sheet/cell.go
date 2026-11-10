package sheet

import (
	"image"
	"image/color"
	"slices"
	"strings"
	"sync"

	"github.com/jhunthrop/foreversixty/api/internal/brand"
)

// socketMasks are the socket's outer shape and the icon's clip, drawn once.
var socketMasks = sync.OnceValues(func() (*image.Alpha, *image.Alpha) {
	return brand.RoundRectMask(int(d(sockW)), int(d(sockW)), d(sockR)), brand.RoundRectMask(int(d(iconW)), int(d(iconW)), d(sockR-ringW))
})

// fitSlack is how far, in logical px, a line may run past the 196 px text box and still fit:
// under one pixel of a browser's sub-pixel layout, as the mock's own fit script measures it.
const fitSlack = 1

// cellText is a cell's words after the fit rule (spec §B.3): the name's lines and style, the
// meta lines chosen, and the gold line.
type cellText struct {
	name      []string
	nameStyle style
	nameColor color.RGBA
	nameLH    float64 // the name block's line height, logical
	meta      []string
	gold      string
	label     [][]word // an empty slot's "{Slot} · {status}", wrapped
	goldInk   color.RGBA
	goldStyle style
}

// cell is a laid-out cell: its words and its socket.
type cell struct {
	cellText
	icon  image.Image // nil for an empty socket
	ring  color.RGBA
	empty bool
}

// fitCell applies the fit rule to one slot's words.
func (r *renderer) fitCell(kind string, sl Slot) (cellText, error) {
	if sl.Empty {
		return r.fitEmpty(kind, sl)
	}
	ct, err := r.fitName(sl)
	if err != nil {
		return ct, err
	}
	var candidates []string
	if kind == KindChar {
		if ct.gold, ct.goldInk, ct.goldStyle, err = r.gainLine(sl); err != nil {
			return ct, err
		}
		candidates = labelled(SlotLabel(sl.Slot), itemLevelMeta(sl.ItemLevel))
	} else {
		candidates = labelled(SlotLabel(sl.Slot), metaCandidates(sl.Meta)...)
	}
	lines := 2
	if len(ct.name) > 1 {
		lines = 1
	}
	if ct.gold != "" {
		lines--
	}
	ct.meta, err = r.fitMeta(candidates, lines)
	return ct, err
}

// fitEmpty is an empty socket's words, muted: "{Slot} · Empty" (/char), "{Slot} · No pick"
// (/bis) and "Off hand · Two-hand in main hand". A /char slot with an upgrade waiting adds the
// pick in gold under it ("BiS: Royal Seal of Eldre'Thalas"), or the panel's status when the pick
// has no name.
func (r *renderer) fitEmpty(kind string, sl Slot) (cellText, error) {
	status := CopyNone
	if sl.TwoHandOff {
		status = CopyTwoHand
	} else if kind == KindChar {
		status = CopyEmpty
	}
	lead, err := r.faces.text(styleMeta)
	if err != nil {
		return cellText{}, err
	}
	words := spanWords(lead, colorMuted, SlotLabel(sl.Slot)+metaJoin+status)
	if sl.TwoHandOff {
		words = keepLastTogether(words)
	}
	ct := cellText{label: wrapWords(words, d(textW+fitSlack))}
	if kind == KindChar && sl.Gain != "" && !sl.TwoHandOff {
		gold, err := r.faces.text(styleGold)
		if err != nil {
			return cellText{}, err
		}
		ct.gold, ct.goldInk, ct.goldStyle = firstFitting(gold, d(textW+fitSlack), named(CopyBisPick, sl.Pick), sl.Gain), colorGold, styleGold
	}
	return ct, nil
}

// named is prefix+name, or "" without a name.
func named(prefix, name string) string {
	if name == "" {
		return ""
	}
	return prefix + name
}

// firstFitting is the first non-empty line that draws within maxW; the last one when none does.
// Nothing is cut.
func firstFitting(t text, maxW float64, lines ...string) string {
	last := ""
	for _, l := range lines {
		if l == "" {
			continue
		}
		if t.width(l) <= maxW {
			return l
		}
		last = l
	}
	return last
}

// gainLine is a worn /char cell's last line, its ink and style: a modelled gain in the gain
// green, naming the pick when it fits ("+11.0 DPS · Dawn Armor"); a pick with no figure in the
// meta's muted regular ("Upgrade: Dawn Armor", else "Gain not modelled yet"); "" with no upgrade.
func (r *renderer) gainLine(sl Slot) (string, color.RGBA, style, error) {
	if sl.Gain == "" {
		return "", color.RGBA{}, styleGold, nil
	}
	st, ink := styleGold, colorGain
	lines := []string{named(sl.Gain+metaJoin, sl.Pick), sl.Gain}
	if !isFigure(sl.Gain) {
		st, ink = styleMeta, colorMuted
		lines = []string{named(CopyUpgradePick, sl.Pick), sl.Gain}
	}
	t, err := r.faces.text(st)
	if err != nil {
		return "", color.RGBA{}, st, err
	}
	return firstFitting(t, d(textW+fitSlack), lines...), ink, st, nil
}

// fitName draws the name on one line at 20 px, shrinking 1 px at a time to 14 px; a name still
// wider takes two lines at 14/17, word-wrapped. Nothing is cut.
func (r *renderer) fitName(sl Slot) (cellText, error) {
	col := colorText
	if sl.known() {
		col = qualityText[sl.Quality]
	}
	maxW := d(textW + fitSlack)
	for px := nameMaxPx; px >= nameMinPx; px-- {
		t, err := r.faces.text(nameStyle(px))
		if err != nil {
			return cellText{}, err
		}
		if t.width(sl.Name) <= maxW {
			return cellText{name: []string{sl.Name}, nameStyle: nameStyle(px), nameColor: col, nameLH: nameLH}, nil
		}
	}
	t, err := r.faces.text(styleNameTwo)
	if err != nil {
		return cellText{}, err
	}
	lines, _ := t.wrap(sl.Name, maxW)
	return cellText{name: lines, nameStyle: styleNameTwo, nameColor: col, nameLH: nameTwoLH}, nil
}

// fitMeta is the first candidate that fits the lines available; none fits, none. Each candidate
// is tried broken between its " · " parts first (the separator is dropped at the break, so the
// slot label stands alone rather than leaving an orphaned "58" or "Bothunk"), then word-wrapped
// as a browser would, before the next, shorter candidate.
func (r *renderer) fitMeta(candidates []string, lines int) ([]string, error) {
	if lines <= 0 {
		return nil, nil
	}
	t, err := r.faces.text(styleMeta)
	if err != nil {
		return nil, err
	}
	maxW := d(textW + fitSlack)
	for _, c := range candidates {
		if wrapped, fits := wrapParts(t, c, maxW); fits && len(wrapped) <= lines {
			return wrapped, nil
		}
		if wrapped, fits := t.wrap(c, maxW); fits && len(wrapped) <= lines {
			return wrapped, nil
		}
	}
	return nil, nil
}

// wrapParts lays s out greedily by its " · " parts, word-wrapping a part only when it cannot
// fit a line alone.
func wrapParts(t text, s string, maxW float64) (lines []string, fits bool) {
	fits = true
	var line string
	for _, part := range strings.Split(s, metaJoin) {
		next := part
		if line != "" {
			next = line + metaJoin + part
		}
		if t.width(next) <= maxW {
			line = next
			continue
		}
		if line != "" {
			lines = append(lines, line)
		}
		if t.width(part) <= maxW {
			line = part
			continue
		}
		words, ok := t.wrap(part, maxW)
		fits = fits && ok
		lines = append(lines, words[:len(words)-1]...)
		line = words[len(words)-1]
	}
	if line != "" {
		lines = append(lines, line)
	}
	return lines, fits
}

// layoutCell is the slot's words and its socket's art.
func (r *renderer) layoutCell(kind string, sl Slot) (cell, error) {
	ct, err := r.fitCell(kind, sl)
	if err != nil {
		return cell{}, err
	}
	if sl.Empty {
		return cell{cellText: ct, empty: true}, nil
	}
	ring := colorEmptyRing
	icon := sl.Icon
	if sl.known() {
		ring = qualityRing[sl.Quality]
	} else {
		icon = QuestionMark
	}
	img, err := iconOrQuestionMark(r.icons, icon)
	if err != nil {
		return cell{}, err
	}
	return cell{cellText: ct, icon: img, ring: ring}, nil
}

// drawCell draws a cell at logical (x, y): the socket on the outer edge, the text beside it,
// right-aligned in the right column.
func (r *renderer) drawCell(c cell, x, y int, right bool) error {
	sockX, textX, a := x, x+sockW+gap, alignLeft
	if right {
		sockX, textX, a = x+cellW-sockW, x+cellW-sockW-gap-textW, alignRight
	}
	r.socket(c, sockX, y)
	top := float64(y + nameTopNudge)
	for _, line := range c.name {
		nameT, err := r.faces.text(c.nameStyle)
		if err != nil {
			return err
		}
		tx, ty := at(float64(textX), top)
		nameT.drawAligned(r.img, c.nameColor, tx, d(textW), ty, a, line)
		top += c.nameLH
	}
	for _, line := range c.label {
		tx, ty := at(float64(textX), top)
		drawWords(r.img, line, tx, d(textW), ty, a)
		top += metaLH
	}
	return r.drawMetaAndGold(c, textX, top, a)
}

func (r *renderer) drawMetaAndGold(c cell, textX int, top float64, a align) error {
	meta, err := r.faces.text(styleMeta)
	if err != nil {
		return err
	}
	for _, line := range c.meta {
		tx, ty := at(float64(textX), top)
		meta.drawAligned(r.img, colorMuted, tx, d(textW), ty, a, line)
		top += metaLH
	}
	if c.gold == "" {
		return nil
	}
	gold, err := r.faces.text(c.goldStyle)
	if err != nil {
		return err
	}
	tx, ty := at(float64(textX), top)
	gold.drawAligned(r.img, c.goldInk, tx, d(textW), ty, a, c.gold)
	return nil
}

// socket is the 36 x 36 socket: a 2 px quality ring with radius 4 around the icon at 32 x 32
// (its native 64 px), or the empty socket's fill and ring.
func (r *renderer) socket(c cell, x, y int) {
	px, py := at(float64(x), float64(y))
	p := image.Pt(int(px), int(py))
	outer, inner := socketMasks()
	inset := int(d(ringW))
	if c.empty {
		brand.FillMask(r.img, p, outer, colorEmptyRing)
		brand.FillMask(r.img, p.Add(image.Pt(inset, inset)), inner, colorEmptyFill)
		return
	}
	brand.FillMask(r.img, p, outer, c.ring)
	brand.FillMask(r.img, p.Add(image.Pt(inset, inset)), inner, colorBlack)
	brand.PlaceImage(r.img, p.Add(image.Pt(inset, inset)), int(d(iconW)), int(d(iconW)), c.icon, inner)
}

// keepLastTogether joins the last two words into one unbreakable word, so "main hand" never
// leaves "hand" alone on a line.
func keepLastTogether(words []word) []word {
	if len(words) < 2 {
		return words
	}
	n := len(words)
	joined := words[n-2]
	joined.text += " " + words[n-1].text
	return append(slices.Clone(words[:n-2]), joined)
}
