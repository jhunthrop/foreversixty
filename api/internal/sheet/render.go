package sheet

import (
	"image"
	"image/color"
	"strconv"
)

// The geometry of spec §B.3, in logical px; the PNG is drawn at scale 2.
const (
	scale = 2

	// Width and Height are the PNG's size.
	Width  = sheetW * scale
	Height = sheetH * scale

	sheetW = 600
	sheetH = 844
	border = 1 // the surface's 1 px line; positioned content starts inside it, as in CSS
	pad    = 16
	row0   = 92
	pitch  = 72
	cellW  = 240
	cellH  = 64
	sockW  = 36
	sockR  = 4
	ringW  = 2
	iconW  = 32
	gap    = 8
	textW  = cellW - sockW - gap // 196
	centre = 72

	nameTopNudge = -2 // the name's line box sits 2 px above the cell (the mock's margin-top)
	nameMaxPx    = 20
	nameMinPx    = 14
	nameLH       = 24
	nameTwoLH    = 17
	metaLH       = 20

	weaponsRow = row0 + pitch*8 // 668
	rangedRow  = row0 + pitch*9 // 740
	rightX     = sheetW - pad - cellW
	rangedX    = (sheetW - cellW) / 2
	centreX    = (sheetW - centre) / 2
	centreH    = pitch*7 + cellH // 568: y 92 to 660
	ruleY      = 76
	footerY    = sheetH - 30
)

// The frame's columns without Shirt and Tabard (spec §B.2): the left column's six sit on the
// right column's rows 2 to 7, so both centre on the crest.
var (
	leftColumn  = []string{"head", "neck", "shoulder", "back", "chest", "wrist"}
	rightColumn = []string{"hands", "waist", "legs", "feet", "finger1", "finger2", "trinket1", "trinket2"}
)

// The text styles of spec §B.3.
var (
	styleKicker   = style{face: barlow600, size: 12, lineHeight: 15, trackingEm: 0.12, upper: true}
	styleTitle    = style{face: cinzel700, size: 24, lineHeight: 30, trackingEm: 0.01}
	styleMark     = style{face: barlow400, size: 13, lineHeight: 16}
	styleFooter   = style{face: barlow400, size: 13, lineHeight: 16}
	styleMeta     = style{face: barlow400, size: 16, lineHeight: metaLH}
	styleGold     = style{face: barlow600, size: 16, lineHeight: metaLH}
	styleNameTwo  = style{face: barlow600, size: nameMinPx, lineHeight: nameTwoLH}
	styleLevel    = style{face: barlow600, size: 14, lineHeight: 17}
	styleRange    = style{face: barlow400, size: 13, lineHeight: 16}
	styleILvlLab  = style{face: barlow600, size: 11, lineHeight: 13, trackingEm: 0.12, upper: true}
	styleILvl     = style{face: cinzel700, size: 22, lineHeight: 26}
	styleFaction  = style{face: barlow600, size: 13, lineHeight: 16}
	styleDimLabel = styleMeta
)

func nameStyle(px int) style { return style{face: barlow600, size: float64(px), lineHeight: nameLH} }

// d is a logical length in device px.
func d(v float64) float64 { return v * scale }

// at is a logical position inside the border, in device px.
func at(x, y float64) (float64, float64) { return d(x + border), d(y + border) }

// renderer is one render's faces and art.
type renderer struct {
	faces *faceSet
	icons IconSource
	img   *image.RGBA
}

func newRenderer(icons IconSource) (*renderer, error) {
	fs, err := newFaceSet()
	if err != nil {
		return nil, err
	}
	return &renderer{faces: fs, icons: icons}, nil
}

func (r *renderer) draw(s Sheet) (*image.RGBA, error) {
	r.img = image.NewRGBA(image.Rect(0, 0, Width, Height))
	r.surface()
	steps := []func(Sheet) error{r.header, r.cells, r.centreColumn, r.footer}
	for _, step := range steps {
		if err := step(s); err != nil {
			return nil, err
		}
	}
	return r.img, nil
}

// surface is --surface-card: the gold wash fading from 7.5% to 1.2% over the top 45%, on
// #0d0f15, inside the 1 px warm line.
func (r *renderer) surface() {
	stop := float64(Height) * 0.45
	for y := range Height {
		a := 0.012
		if fy := float64(y) + 0.5; fy < stop {
			a = 0.075 + (0.012-0.075)*fy/stop
		}
		fillRect(r.img, image.Rect(0, y, Width, y+1), over(colorSurface, colorGold, a))
	}
	b := border * scale
	fillRect(r.img, image.Rect(0, 0, Width, b), colorLineWarm)
	fillRect(r.img, image.Rect(0, Height-b, Width, Height), colorLineWarm)
	fillRect(r.img, image.Rect(0, 0, b, Height), colorLineWarm)
	fillRect(r.img, image.Rect(Width-b, 0, Width, Height), colorLineWarm)
}

func (r *renderer) header(s Sheet) error {
	kicker, err := r.faces.text(styleKicker)
	if err != nil {
		return err
	}
	title, err := r.faces.text(styleTitle)
	if err != nil {
		return err
	}
	x, y := at(pad, 16)
	kicker.draw(r.img, colorGold, x, y, s.Kicker)
	x, y = at(pad, 34)
	title.draw(r.img, classColors[s.Class], x, y, s.Title)
	if err := r.mark(); err != nil {
		return err
	}
	rx, ry := at(pad, ruleY)
	fillRect(r.img, image.Rect(int(rx), int(ry), int(d(sheetW-border-pad)), int(ry)+scale), colorLineWarm)
	return nil
}

// mark is the site mark, right 16 and top 22: "foreversixty.gg" then the seal at 28 x 28.
func (r *renderer) mark() error {
	const sealW, top = 28, 22
	right := float64(sheetW - border - pad)
	sealImg, err := seal()
	if err != nil {
		return err
	}
	placeImage(r.img, image.Pt(int(d(right-sealW)), int(d(top+border))), int(d(sealW)), int(d(sealW)), sealImg, nil)
	mark, err := r.faces.text(styleMark)
	if err != nil {
		return err
	}
	textRight := d(right - sealW - gap)
	mark.draw(r.img, colorMuted, textRight-mark.width(CopyMark), d(top+border+(sealW-16)/2), CopyMark)
	return nil
}

func (r *renderer) footer(s Sheet) error {
	t, err := r.faces.text(styleFooter)
	if err != nil {
		return err
	}
	x, y := at(pad, footerY)
	w := d(sheetW - 2*border - 2*pad)
	t.draw(r.img, colorMuted, x, y, s.Footer)
	t.drawAligned(r.img, colorMuted, x, w, y, alignRight, s.FooterRight)
	return nil
}

// placement is where a slot's cell sits, in logical px, and whether it is right-styled.
type placement struct {
	key   string
	x, y  int
	right bool
}

// placements are the frame's cells (spec §B.3): the left and right columns, then the weapons
// rows; the off hand is left out when the main hand is two-handed.
func placements(twoHand bool) []placement {
	out := make([]placement, 0, len(slotLabels))
	for i, k := range leftColumn {
		out = append(out, placement{k, pad, row0 + pitch*(i+1), false})
	}
	for i, k := range rightColumn {
		out = append(out, placement{k, rightX, row0 + pitch*i, true})
	}
	out = append(out, placement{"main_hand", pad, weaponsRow, false})
	if !twoHand {
		out = append(out, placement{SlotOff, rightX, weaponsRow, true})
	}
	return append(out, placement{"ranged", rangedX, rangedRow, false})
}

// cells draws every placed slot.
func (r *renderer) cells(s Sheet) error {
	for _, p := range placements(s.TwoHand) {
		c, err := r.layoutCell(s.Kind, s.slot(p.key))
		if err != nil {
			return err
		}
		if err := r.drawCell(c, p.x, p.y, p.right); err != nil {
			return err
		}
	}
	return nil
}

// centreColumn is the crest, the band or the item level, and the faction, centred in the
// column between the slots (x 264, y 92 to 660).
func (r *renderer) centreColumn(s Sheet) error {
	type line struct {
		st    style
		col   color.RGBA
		text  string
		above float64 // margin-top, logical
	}
	var lines []line
	if s.Kind == KindBis {
		if s.Level > 0 {
			lines = append(lines, line{styleLevel, colorStrong, "Level " + strconv.Itoa(s.Level), 16})
		}
		if s.Subtitle != "" {
			lines = append(lines, line{styleRange, colorMuted, s.Subtitle, 0})
		}
	} else {
		lines = append(lines, line{styleILvlLab, colorMuted, CopyItemLevel, 18}, line{styleILvl, colorGold, s.Subtitle, 0})
	}
	const crestW, factionW, factionAbove, wordAbove = 64, 32, 20, 2
	total := float64(crestW + factionAbove + factionW + wordAbove + styleFaction.lineHeight)
	for _, l := range lines {
		total += l.above + l.st.lineHeight
	}
	y := row0 + (centreH-total)/2
	cx := float64(centreX)
	crest, err := r.icons.Crest(s.Class)
	if err != nil {
		return err
	}
	placeImage(r.img, image.Pt(int(d(cx+border+(centre-crestW)/2)), int(d(y+border))), int(d(crestW)), int(d(crestW)), crest, nil)
	y += crestW
	for _, l := range lines {
		t, err := r.faces.text(l.st)
		if err != nil {
			return err
		}
		y += l.above
		x, top := at(cx, y)
		t.drawAligned(r.img, l.col, x, d(centre), top, alignCentre, l.text)
		y += l.st.lineHeight
	}
	return r.faction(s.Faction, cx, y+factionAbove)
}

func (r *renderer) faction(faction string, cx, y float64) error {
	const markW, wordAbove = 32, 2
	mark, err := factionMark(faction)
	if err != nil {
		return err
	}
	placeImage(r.img, image.Pt(int(d(cx+border+(centre-markW)/2)), int(d(y+border))), int(d(markW)), int(d(markW)), mark, nil)
	word, col := "Alliance", colorAlliance
	if faction == "horde" {
		word, col = "Horde", colorHorde
	}
	t, err := r.faces.text(styleFaction)
	if err != nil {
		return err
	}
	x, top := at(cx, y+markW+wordAbove)
	t.drawAligned(r.img, col, x, d(centre), top, alignCentre, word)
	return nil
}
