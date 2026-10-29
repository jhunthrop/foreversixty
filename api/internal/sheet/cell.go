package sheet

import (
	"image"
	"image/color"
	"sync"
)

// socketMasks are the socket's outer shape and the icon's clip, drawn once.
var socketMasks = sync.OnceValues(func() (*image.Alpha, *image.Alpha) {
	return roundRectMask(int(d(sockW)), int(d(sockW)), d(sockR)), roundRectMask(int(d(iconW)), int(d(iconW)), d(sockR-ringW))
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
		ct.gold = sl.Gain
		if m := itemLevelMeta(sl.ItemLevel); m != "" {
			candidates = []string{m}
		}
	} else {
		candidates = metaCandidates(sl.Meta)
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

// fitEmpty is an empty socket's words: "Empty" (with its gold line) on /char, "No pick" on /bis.
func (r *renderer) fitEmpty(kind string, sl Slot) (cellText, error) {
	ct := cellText{nameStyle: styleDimLabel, nameColor: colorMuted, nameLH: styleDimLabel.lineHeight}
	if kind == KindChar {
		ct.name, ct.gold = []string{CopyEmpty}, sl.Gain
		return ct, nil
	}
	ct.name = []string{CopyNone}
	return ct, nil
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

// fitMeta is the first candidate that fits the lines available, word-wrapped; none fits, none.
func (r *renderer) fitMeta(candidates []string, lines int) ([]string, error) {
	if lines <= 0 {
		return nil, nil
	}
	t, err := r.faces.text(styleMeta)
	if err != nil {
		return nil, err
	}
	for _, c := range candidates {
		wrapped, fits := t.wrap(c, d(textW+fitSlack))
		if fits && len(wrapped) <= lines {
			return wrapped, nil
		}
	}
	return nil, nil
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
		sockX, textX, a = x+cellW-sockW, x, alignRight
	}
	r.socket(c, sockX, y)
	top := float64(y + nameTopNudge)
	nameT, err := r.faces.text(c.nameStyle)
	if err != nil {
		return err
	}
	for _, line := range c.name {
		tx, ty := at(float64(textX), top)
		nameT.drawAligned(r.img, c.nameColor, tx, d(textW), ty, a, line)
		top += c.nameLH
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
	gold, err := r.faces.text(styleGold)
	if err != nil {
		return err
	}
	tx, ty := at(float64(textX), top)
	gold.drawAligned(r.img, colorGold, tx, d(textW), ty, a, c.gold)
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
		fillMask(r.img, p, outer, colorEmptyRing)
		fillMask(r.img, p.Add(image.Pt(inset, inset)), inner, colorEmptyFill)
		return
	}
	fillMask(r.img, p, outer, c.ring)
	fillMask(r.img, p.Add(image.Pt(inset, inset)), inner, colorBlack)
	placeImage(r.img, p.Add(image.Pt(inset, inset)), int(d(iconW)), int(d(iconW)), c.icon, inner)
}
