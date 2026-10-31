package discordbot

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/bis"
	"github.com/jhunthrop/foreversixty/api/internal/sheet"
	"github.com/jhunthrop/foreversixty/api/internal/upgrades"
)

// SheetRenderer draws the character sheet a share posts (spec 2026-10-29 Part B);
// *sheet.Renderer is the production one.
type SheetRenderer interface {
	Render(s sheet.Sheet) ([]byte, error)
}

// The shared sheet's attachment (spec §B.4).
const (
	sheetFileName    = "sheet.png"
	sheetContentType = "image/png"
	sheetDateLayout  = "2 Jan 2006"
)

// sheetPost is a public share drawn as the character sheet: the model, and the words around
// the image.
type sheetPost struct {
	model  sheet.Sheet
	accent int
	head   string
	link   Component
}

// publicAnswer is what Share in channel posts for command, from one read: the visitor's
// answer, and for /bis without a slot and /char the sheet drawn from the same data. ok is
// false when there is nothing public to show. A sheet that cannot be modelled leaves the
// answer to post alone (spec §B.5), logged sheet_render.
func (s *Service) publicAnswer(ctx context.Context, command string, guildID int64, opts map[string]string) (Message, *sheetPost, bool, error) {
	switch command {
	case "bis":
		d, found, err := s.loadBis(ctx, guildID, opts)
		if err != nil || !found {
			return Message{}, nil, false, err
		}
		panel, ok, err := s.bisPanelOf(ctx, d, opts)
		if err != nil || !ok || opts[optSlot] != "" {
			return panel, nil, ok, err
		}
		return panel, s.bisSheet(d), true, nil
	case "char":
		d, found, err := s.loadChar(ctx, guildID, nil, opts)
		if err != nil || !found {
			return Message{}, nil, false, err
		}
		post := s.charSheet(d)
		if post == nil {
			s.log.Warn("discordbot", "op", "sheet_render", "outcome", "no_model", "class", d.class, "faction", d.faction)
		}
		return s.charPanel(ctx, d), post, true, nil
	}
	m, ok, err := s.visitorAnswer(command)(ctx, guildID, opts)
	return m, nil, ok, err
}

// publicShare is what Share in channel posts: the sheet with its words when it renders, else
// panel (the visitor's answer, its link row only). A render failure is logged sheet_render.
func (s *Service) publicShare(post *sheetPost, panel Message, in Interaction) Message {
	if post == nil {
		return panel
	}
	if s.d.Sheets == nil {
		s.log.Warn("discordbot", "op", "sheet_render", "outcome", "no_renderer", "interaction", in.ID)
		return panel
	}
	start := s.now()
	png, err := s.d.Sheets.Render(post.model)
	if err != nil {
		s.log.Warn("discordbot", "op", "sheet_render", "interaction", in.ID, "err", err)
		return panel
	}
	s.log.Info("discordbot", "op", "sheet_render", "outcome", "ok", "ms", s.now().Sub(start).Milliseconds(), "bytes", len(png))
	return sheetMessage(post, png)
}

// sheetMessage is the shared sheet's message: the heading, the sheet in a gallery with its
// alt text, and the one link (spec §B.4). The /char gold legend is on the image alone.
func sheetMessage(post *sheetPost, png []byte) Message {
	alt := sheet.Alt(post.model)
	body := []Node{textNode(post.head), galleryNode("attachment://"+sheetFileName, alt)}
	return Message{
		ComponentsV2: true,
		Layout:       []Node{containerNode(post.accent, body...)},
		Components:   []ActionRow{{post.link}},
		Files:        []Attachment{{Name: sheetFileName, Description: alt, ContentType: sheetContentType, Data: png}},
	}
}

// bisSheet is the /bis sheet, all 17 picks of the band whichever page was shared, and the
// words of its shared message.
func (s *Service) bisSheet(d bisData) *sheetPost {
	model := sheet.Sheet{Kind: sheet.KindBis, Class: d.spec.ClassSlug, Spec: d.spec.Name, Faction: d.faction,
		Title: fill(strings.TrimPrefix(panelCopy.BisHeadTitle, heading("")), d.who()...), Kicker: sheet.CopyBisKicker,
		Subtitle: sheetRange(d), Footer: sheetDate(panelCopy.SheetUpdated, d.band.GeneratedAt), Level: d.level, Band: d.label}
	for _, c := range slotChoices {
		sl, has := d.bySlot[c.Value]
		switch {
		case has && sl.EmptyReason == bis.EmptyTwoHand:
			model.TwoHand = true
		case !has || sl.ItemID <= 0:
			model.Slots = append(model.Slots, sheet.Slot{Slot: c.Value, Empty: true})
		default:
			cell := s.sheetCell(d.build, d.spec.ClassSlug, c.Value, sl.ItemID, sl.ItemName)
			cell.Meta = s.sourceLine(sl.ItemID, sl.SourceKind, d.faction, sl.Source)
			model.Slots = append(model.Slots, cell)
		}
	}
	return &sheetPost{model: model, accent: classColorOf(d.spec.ClassSlug), head: fill(panelCopy.BisHeadTitle, d.who()...) + "\n" + d.headLine(),
		link: Component{Kind: KindButton, Style: StyleLink, Label: lookupCopy.BisAllSlots, URL: d.page}}
}

// charSheet is the /char sheet: what the character wears, as a visitor sees it, with the
// gold gain of each slot the band's pick improves; nil for a character the sheet cannot draw
// (no class or faction it knows).
func (s *Service) charSheet(d charData) *sheetPost {
	if _, known := classColor(d.class); !known || (d.faction != "alliance" && d.faction != "horde") {
		return nil
	}
	level := 0
	if d.card.Character.Level != nil {
		level = *d.card.Character.Level
	}
	guild := ""
	if d.card.Character.Guild != nil {
		guild = d.card.Character.Guild.Name
	}
	average, _ := averageItemLevel(d.card.Gear)
	model := sheet.Sheet{Kind: sheet.KindChar, Class: d.class, Spec: d.spec.Name, Faction: d.faction,
		Title: d.card.Character.Name, Kicker: cardLineWith(d.card, d.spec, d.class, func(s string) string { return s }),
		Subtitle: average, Level: level, Guild: guild, Slots: s.charSheetCells(d)}
	model.TwoHand = s.wearsTwoHander(d)
	post := &sheetPost{model: model, accent: classColorOf(d.class), head: s.charHeader(d, false),
		link: Component{Kind: KindButton, Style: StyleLink, Label: lookupCopy.CharacterPage, URL: s.characterURL(d.card.Character.Key)}}
	if d.ups.band != "" {
		post.model.Footer = fill(panelCopy.SheetCharGold, "{band label}", d.ups.band)
	}
	return post
}

// charSheetCells are the worn items, each with its upgrade's gold line; an empty slot says
// whether an upgrade waits there.
func (s *Service) charSheetCells(d charData) []sheet.Slot {
	worn := wornBySlot(d.card.Gear)
	ups := d.ups.bySlot()
	var out []sheet.Slot
	for _, c := range slotChoices {
		u, upgraded := ups[c.Value]
		w := worn[c.Value]
		if w == nil {
			cell := sheet.Slot{Slot: c.Value, Empty: true}
			if upgraded {
				cell.Gain = sheet.CopyUpgrade
				if u.gain == panelCopy.NotModelled {
					cell.Gain = sheet.CopyNotModelled
				}
			}
			out = append(out, cell)
			continue
		}
		name, level := s.wornFacts(d, w)
		cell := s.sheetCell(d.build, d.class, c.Value, w.ID, name)
		if level > 0 {
			cell.ItemLevel = level
		}
		if upgraded {
			cell.Gain = u.gain
		}
		out = append(out, cell)
	}
	return out
}

// wearsTwoHander is a character whose main hand is two-handed and whose off hand is empty.
func (s *Service) wearsTwoHander(d charData) bool {
	worn := wornBySlot(d.card.Gear)
	main := worn["main_hand"]
	if main == nil || worn["off_hand"] != nil {
		return false
	}
	it, ok := s.d.Guilds.BuildItem(d.build, d.class, main.ID)
	return ok && it.TwoHand
}

// sheetCell is an item as the sheet draws it: the build's quality, item level and icon, the
// given name first (the data's), and the question mark's empty fields for an unknown item.
func (s *Service) sheetCell(build, class, slot string, id int, name string) sheet.Slot {
	cell := sheet.Slot{Slot: slot, ItemID: id, Name: name, Quality: -1}
	if it, ok := s.d.Guilds.BuildItem(build, class, id); ok {
		cell.Quality, cell.ItemLevel, cell.Icon = it.Quality, it.ItemLevel, it.Icon
		if cell.Name == "" {
			cell.Name = it.Name
		}
	}
	return cell
}

// sheetRange is the /bis sheet's range line: the Phase 1 raid, level 60 bare, or a lower
// band's levels.
func sheetRange(d bisData) string {
	switch {
	case d.phase == phaseRaid:
		return panelCopy.SheetRangeRaid
	case d.level == upgrades.TopBand:
		return panelCopy.SheetRangeBare
	}
	return fill(panelCopy.SheetRangeLevels, "{n}", strconv.Itoa(d.level), "{n+9}", strconv.Itoa(d.level+9))
}

// sheetDate is template with its date ("29 Oct 2026") from an RFC 3339 stamp; "" when the
// stamp is missing or unreadable.
func sheetDate(template, stamp string) string {
	at, err := time.Parse(time.RFC3339, stamp)
	if err != nil {
		return ""
	}
	return fill(template, "{d Mon yyyy}", at.UTC().Format(sheetDateLayout))
}
