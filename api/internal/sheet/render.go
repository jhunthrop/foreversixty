package sheet

import (
	"fmt"
	"image"
	"image/color"
	"strconv"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/brand"
)

// The geometry in logical px (round 5, design/specs/2026-11-09-directions.md); the PNG is drawn
// at scale 2. The sheet is the site's character page in small: the blue title band across the
// top, then the brown ground with the gear on one card, the card's frame and corners marking it
// as the sheet's one lead panel, and the footer line on the ground under it.
const (
	scale = 2

	// Width and Height are the PNG's size.
	Width  = sheetW * scale
	Height = sheetH * scale

	sheetW = edge*2 + inset*2 + cellW*2 + centre // 660
	sheetH = footerY + footerLH + edge

	// Discord shows the sheet about 550 px wide on desktop and 350 px on a phone, so the type
	// is set large against the canvas: a 18 px meta line is about 15 px on desktop.
	bandH   = 80 // the title band
	edge    = 10 // the ground around the card
	inset   = 26 // the card's padding: the corner ornaments sit in it
	cardTop = bandH + edge
	cornerW = 44 // the ornaments at the site's small size

	pad    = edge + inset // the slots' left edge
	row0   = cardTop + inset
	pitch  = 80
	cellW  = 256
	cellH  = 72
	sockW  = 36
	sockR  = 4
	ringW  = 2
	iconW  = 32
	gap    = 10
	textW  = cellW - sockW - gap - 4 // 206: 4 px clear of the centre column
	centre = 76

	nameTopNudge = -3 // the name's line box sits 3 px above the cell
	nameMaxPx    = 23
	nameMinPx    = 16
	nameLH       = 27
	nameTwoLH    = 19
	metaLH       = 23

	weaponsRow = row0 + pitch*8
	rangedRow  = row0 + pitch*9
	cardBottom = rangedRow + cellH + inset
	rightX     = sheetW - pad - cellW
	rangedX    = (sheetW - cellW) / 2
	centreX    = (sheetW - centre) / 2
	centreH    = pitch*7 + cellH
	footerY    = cardBottom + 10
	footerLH   = 22

	bandCrest = 48 // the class crest leading the band
	bandSeal  = 44 // the seal closing it, as the site's nav bar draws it
	bandPadX  = 18
)

// The frame's columns without Shirt and Tabard (spec §B.2), both top-aligned as in the game:
// the left column's six sit on the right column's rows 1 to 6.
var (
	leftColumn  = []string{"head", "neck", "shoulder", "back", "chest", "wrist"}
	rightColumn = []string{"hands", "waist", "legs", "feet", "finger1", "finger2", "trinket1", "trinket2"}
)

// The text styles: Marcellus for the title and the big figure, Open Sans for every other word,
// with no letter spacing (round 5's tracking.label is 0 on the web).
var (
	styleKicker  = style{face: brand.OpenSans400, size: 16, lineHeight: 20}
	styleTitle   = style{face: brand.Marcellus400, size: 30, lineHeight: 34}
	styleBrand   = style{face: brand.Marcellus400, size: 19, lineHeight: 22}
	styleMark    = style{face: brand.OpenSans400, size: 14, lineHeight: 18}
	styleFooter  = style{face: brand.OpenSans400, size: 16, lineHeight: footerLH}
	styleMeta    = style{face: brand.OpenSans400, size: 18, lineHeight: metaLH}
	styleGold    = style{face: brand.OpenSans600, size: 18, lineHeight: metaLH}
	styleNameTwo = style{face: brand.OpenSans600, size: nameMinPx, lineHeight: nameTwoLH}
	styleLevel   = style{face: brand.Marcellus400, size: 30, lineHeight: 34}
	styleRange   = style{face: brand.OpenSans400, size: 14, lineHeight: 18}
	styleILvlLab = style{face: brand.OpenSans600, size: 14, lineHeight: 18}
	styleILvl    = style{face: brand.Marcellus400, size: 34, lineHeight: 38}
	styleGainSum = style{face: brand.Marcellus400, size: 26, lineHeight: 30}
	styleFaction = style{face: brand.OpenSans600, size: 14, lineHeight: 18}
)

func nameStyle(px int) style {
	return style{face: brand.OpenSans600, size: float64(px), lineHeight: nameLH}
}

// d is a logical length in device px.
func d(v float64) float64 { return v * scale }

// at is a logical position in device px.
func at(x, y float64) (float64, float64) { return d(x), d(y) }

// renderer is one render's faces and art.
type renderer struct {
	faces *faceSet
	icons IconSource
	img   *image.RGBA
}

func newRenderer(icons IconSource) *renderer {
	return &renderer{faces: newFaceSet(), icons: icons}
}

func (r *renderer) draw(s Sheet) (*image.RGBA, error) {
	r.img = image.NewRGBA(image.Rect(0, 0, Width, Height))
	steps := []func(Sheet) error{r.surface, r.header, r.cells, r.centreColumn, r.footer}
	for _, step := range steps {
		if err := step(s); err != nil {
			return nil, err
		}
	}
	return r.img, nil
}

// surface is the ground, the band and the card.
func (r *renderer) surface(Sheet) error {
	if err := brand.Ground(r.img, r.img.Bounds()); err != nil {
		return err
	}
	if err := brand.Band(r.img, image.Rect(0, 0, Width, int(d(bandH)))); err != nil {
		return err
	}
	card := image.Rect(int(d(edge)), int(d(cardTop)), int(d(sheetW-edge)), int(d(cardBottom)))
	return brand.Panel(r.img, card, scale, brand.PanelStyle{Ornate: true, CornerSide: cornerW})
}

// header is the band's words: the class crest, the title over the kicker, and the site's seal and
// name at the right.
func (r *renderer) header(s Sheet) error {
	crest, err := r.icons.Crest(s.Class)
	if err != nil {
		return err
	}
	top := (bandH - bandCrest) / 2.0
	brand.PlaceImage(r.img, image.Pt(int(d(bandPadX)), int(d(top))), int(d(bandCrest)), int(d(bandCrest)), crest, nil)
	title, err := r.faces.text(styleTitle)
	if err != nil {
		return err
	}
	kicker, err := r.faces.text(styleKicker)
	if err != nil {
		return err
	}
	textX := bandPadX + bandCrest + 12.0
	block := styleTitle.lineHeight + styleKicker.lineHeight
	y := (bandH - block) / 2
	titleColor := colorCream
	if s.Kind == KindChar {
		titleColor = classColors[s.Class]
	}
	maxW := d(sheetW - textX - 170)
	tx, ty := at(textX, y)
	title.draw(r.img, titleColor, tx, ty, fitWidth(title, s.Title, maxW))
	tx, ty = at(textX, y+styleTitle.lineHeight)
	kicker.draw(r.img, colorText, tx, ty, fitWidth(kicker, s.Kicker, maxW))
	return r.mark()
}

// mark is the site's seal and name at the band's right: "Forever Sixty" over "foreversixty.gg".
func (r *renderer) mark() error {
	sealImg, err := seal()
	if err != nil {
		return err
	}
	right := float64(sheetW - bandPadX)
	top := (bandH - bandSeal) / 2.0
	brand.PlaceImage(r.img, image.Pt(int(d(right-bandSeal)), int(d(top))), int(d(bandSeal)), int(d(bandSeal)), sealImg, nil)
	name, err := r.faces.text(styleBrand)
	if err != nil {
		return err
	}
	url, err := r.faces.text(styleMark)
	if err != nil {
		return err
	}
	textRight := d(right - bandSeal - 10)
	y := (bandH - styleBrand.lineHeight - styleMark.lineHeight) / 2
	name.draw(r.img, colorFooter, textRight-name.width(CopyBrand), d(y), CopyBrand)
	url.draw(r.img, colorText, textRight-url.width(CopyMark), d(y+styleBrand.lineHeight), CopyMark)
	return nil
}

// footer is the sheet's last line, on the ground under the card.
func (r *renderer) footer(s Sheet) error {
	t, err := r.faces.text(styleFooter)
	if err != nil {
		return err
	}
	x, y := at(pad, footerY)
	t.draw(r.img, colorFooter, x, y, fitWidth(t, s.Footer, d(sheetW-2*pad)))
	return nil
}

// fitWidth trims s with "..." until it draws within maxW device px.
func fitWidth(t text, s string, maxW float64) string {
	if t.width(s) <= maxW {
		return s
	}
	runes := []rune(s)
	for len(runes) > 0 {
		runes = runes[:len(runes)-1]
		cut := strings.TrimRight(string(runes), " ") + "..."
		if t.width(cut) <= maxW {
			return cut
		}
	}
	return ""
}

// placement is where a slot's cell sits, in logical px, and whether it is right-styled.
type placement struct {
	key   string
	x, y  int
	right bool
}

// placements are the frame's cells (spec §B.3): the left and right columns, then the weapons
// rows. The off hand is always placed: under a two-handed main hand it is drawn empty with the
// words "Two-hand in main hand", as the client leaves that slot empty (paperdoll spec §2).
func placements() []placement {
	out := make([]placement, 0, len(slotLabels))
	for i, k := range leftColumn {
		out = append(out, placement{k, pad, row0 + pitch*i, false})
	}
	for i, k := range rightColumn {
		out = append(out, placement{k, rightX, row0 + pitch*i, true})
	}
	out = append(out, placement{"main_hand", pad, weaponsRow, false})
	out = append(out, placement{SlotOff, rightX, weaponsRow, true})
	return append(out, placement{"ranged", rangedX, rangedRow, false})
}

// cells draws every placed slot.
func (r *renderer) cells(s Sheet) error {
	for _, p := range placements() {
		c, err := r.layoutCell(s.Kind, s.cellFor(p.key))
		if err != nil {
			return err
		}
		if err := r.drawCell(c, p.x, p.y, p.right); err != nil {
			return err
		}
	}
	return nil
}

// centreLine is one text line of the centre column, above the faction.
type centreLine struct {
	st    style
	col   color.RGBA
	text  string
	above float64 // margin-top, logical
}

// centreLines are the centre column's text lines: /bis the level and the range; /char the item
// level, drawn only when every worn item is known (paperdoll spec §3.5) -- the caller leaves
// Subtitle empty otherwise, and the faction centres without it.
func centreLines(s Sheet) []centreLine {
	var lines []centreLine
	if s.Kind == KindBis {
		if s.Level > 0 {
			lines = append(lines, centreLine{styleILvlLab, colorMuted, CopyLevel, 0},
				centreLine{styleLevel, colorGold, strconv.Itoa(s.Level), 0})
		}
		if s.Subtitle != "" {
			lines = append(lines, centreLine{styleRange, colorText, s.Subtitle, 4})
		}
	} else {
		if s.Subtitle != "" {
			lines = append(lines, centreLine{styleILvlLab, colorMuted, CopyItemLevel, 0}, centreLine{styleILvl, colorGold, s.Subtitle, 0})
		}
		lines = append(lines, upgradeLines(s)...)
	}
	return lines
}

// upgradeLines are the /char summary under the item level: "Upgrades", the summed gain's
// figure in the gain green, then its unit and the slots it spans ("DPS · 6 slots").
func upgradeLines(s Sheet) []centreLine {
	figure, unit, ok := strings.Cut(s.UpgradeTotal, " ")
	if !ok || s.UpgradeSlots <= 0 {
		return nil
	}
	slots := CopySlotsOne
	if s.UpgradeSlots > 1 {
		slots = fmt.Sprintf(CopySlotsMany, s.UpgradeSlots)
	}
	return []centreLine{
		{styleILvlLab, colorMuted, CopyUpgrades, 22},
		{styleGainSum, colorGain, figure, 0},
		{styleRange, colorText, unit, 0},
		{styleRange, colorMuted, slots, 0},
	}
}

// centreColumn is the level or the item level over the faction, centred in the column between
// the slots.
func (r *renderer) centreColumn(s Sheet) error {
	lines := centreLines(s)
	const factionW, factionAbove, wordAbove = 44, 22, 4
	total := float64(factionW + wordAbove + styleFaction.lineHeight)
	if len(lines) > 0 {
		total += factionAbove
	}
	for _, l := range lines {
		total += l.above + l.st.lineHeight
	}
	y := row0 + (centreH-total)/2
	cx := float64(centreX)
	for _, l := range lines {
		t, err := r.faces.text(l.st)
		if err != nil {
			return err
		}
		y += l.above
		x, top := at(cx-2, y)
		t.drawAligned(r.img, l.col, x, d(centre+4), top, alignCentre, l.text)
		y += l.st.lineHeight
	}
	if len(lines) > 0 {
		y += factionAbove
	}
	return r.faction(s.Faction, cx, y)
}

func (r *renderer) faction(faction string, cx, y float64) error {
	const markW, wordAbove = 44, 4
	mark, err := factionMark(faction)
	if err != nil {
		return err
	}
	brand.PlaceImage(r.img, image.Pt(int(d(cx+(centre-markW)/2)), int(d(y))), int(d(markW)), int(d(markW)), mark, nil)
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
