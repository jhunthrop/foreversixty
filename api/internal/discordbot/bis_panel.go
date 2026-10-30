package discordbot

import (
	"context"
	"encoding/json"
	"math"
	"strconv"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/bis"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

// optPage is a /bis answer's panel page, "1" (Head to Legs) or "2" (Feet to Ranged). The page
// button carries it; Share never stores it, since the sheet shows all 17 slots.
const optPage = "page"

// The panel's pages (spec 2026-10-29 §A.5.1): the character frame's 17 slots, nine then eight.
const (
	slotsPerPage = 9
	firstPage    = 1
	lastPage     = 2
)

// minShownGain is the smallest gain a row states; a smaller one is a tie and shows none.
const minShownGain = 0.05

// textBudget is the bot's guard on a V2 message's text, under Discord's 4,000 (spec §A.6);
// maxV2Components its component ceiling, two under Discord's 40.
const (
	textBudget      = 3800
	maxV2Components = 38
)

// partPage is the page button's card part.
const partPage = "pg"

// bisData is a /bis answer's band as the panel and the sheet read it.
type bisData struct {
	spec    specs.Spec
	label   string
	level   int
	phase   string
	faction string
	build   string
	page    string
	band    bis.Band
	bySlot  map[string]bis.Slot
}

// loadBis reads the band a /bis answer's options name; found is false when any part is unknown.
func (s *Service) loadBis(ctx context.Context, guildID int64, opts map[string]string) (bisData, bool, error) {
	sp, ok := specByKey(opts[optSpec])
	label, known := bandLabels[opts[optPhase]]
	if !ok || !known {
		return bisData{}, false, nil
	}
	g, err := s.d.Guilds.BotGuildFacts(ctx, guildID)
	if err != nil {
		return bisData{}, false, err
	}
	faction := g.Faction
	if faction == "" {
		faction = defaultFaction
	}
	level, preset := phaseBand(opts[optPhase])
	build, _ := s.d.Guilds.ActiveBuild()
	band, found, err := bis.LoadPresetBand(s.d.Guilds.BisDataDir(), s.d.Guilds.BisBuild(), sp.Spec, level, faction, preset)
	if err != nil || !found {
		return bisData{}, false, err
	}
	return bisData{spec: sp, label: label, level: level, phase: opts[optPhase], faction: faction, build: build,
		page: s.bisPageURL(sp, faction, level, preset), band: band, bySlot: band.BySlot()}, true, nil
}

// answerBis is the /bis answer for resolved options: the slot panel with a slot, else the page
// panel. It is the same for every viewer: the BiS lists are public. Its one row is the link;
// the presser's own copy gains the page button and Change picks (decorateOwn).
func (s *Service) answerBis(ctx context.Context, guildID int64, _ *int64, opts map[string]string) (Message, bool, error) {
	d, found, err := s.loadBis(ctx, guildID, opts)
	if err != nil || !found {
		return Message{}, false, err
	}
	if slot := opts[optSlot]; slot != "" {
		return s.bisSlotPanel(ctx, d, slot)
	}
	return s.bisPagePanel(ctx, d, pageOf(opts)), true, nil
}

// pageOf is the options' page; anything but "2" is page 1.
func pageOf(opts map[string]string) int {
	if opts[optPage] == strconv.Itoa(lastPage) {
		return lastPage
	}
	return firstPage
}

// pageSlots are a page's slot keys in the character frame's order.
func pageSlots(page int) []Choice {
	if page == lastPage {
		return slotChoices[slotsPerPage:]
	}
	return slotChoices[:slotsPerPage]
}

// figure is a gain in the band's unit: DPS, or Forever Score points for a tank band.
func (d bisData) figure(v float64) string {
	if d.band.ScoreUnit == bis.TankScoreUnit {
		return signedPoints(v)
	}
	return signedDPS(v)
}

// headLine is "Best in slot · {band label} · {Faction}".
func (d bisData) headLine() string {
	return fill(panelCopy.BisHeadLine, "{band label}", d.label, "{Faction}", classLabel(d.faction))
}

func (d bisData) who() []string {
	return []string{"{Spec}", d.spec.Name, "{Class}", classLabel(d.spec.ClassSlug)}
}

// facts is an item of the band with the BiS page's own source line for its kind.
func (s *Service) bisFacts(d bisData, id int, name, kind, raw string) ItemFacts {
	return s.itemFacts(d.build, d.spec.ClassSlug, id, name, s.sourceLine(id, kind, d.faction, raw), "")
}

// pickGain is a pick's line-2 figure: Gain not modelled yet for an unmodelled effect, else
// its gain over the next best when there is one.
func (d bisData) pickGain(sl bis.Slot) string {
	if sl.EffectUnmodelled {
		return panelCopy.NotModelled
	}
	if len(sl.Alternatives) == 0 || math.Abs(sl.Alternatives[0].DpsDelta) < minShownGain {
		return ""
	}
	return d.figure(-sl.Alternatives[0].DpsDelta)
}

// bisPagePanel is one page of the /bis panel: the header, then each slot's row.
func (s *Service) bisPagePanel(ctx context.Context, d bisData, page int) Message {
	slots := pageSlots(page)
	header := strings.Join([]string{fill(panelCopy.BisHeadTitle, d.who()...), d.headLine(),
		fill(panelCopy.BisHeadPage, "{Page range}", panelCopy.BisPages[page-1], "{n}", strconv.Itoa(len(slots)))}, "\n")
	rows := make([]panelRow, 0, len(slots))
	for _, c := range slots {
		rows = append(rows, s.bisSlotRow(d, c))
	}
	body := s.guardedBody(ctx, d.build, []Node{sectionNode(header, s.crestThumb(d.spec.ClassSlug)), separatorNode()}, rows)
	link := Component{Kind: KindButton, Style: StyleLink, Label: lookupCopy.BisAllSlots, URL: d.page}
	return Message{Layout: []Node{containerNode(classColorOf(d.spec.ClassSlug), body...)}, Components: []ActionRow{{link}}}
}

// bisSlotRow is one slot of a page: the pick's row, or its one line when it has no pick.
func (s *Service) bisSlotRow(d bisData, c Choice) panelRow {
	sl, has := d.bySlot[c.Value]
	switch {
	case has && sl.EmptyReason == bis.EmptyTwoHand:
		main := d.bySlot["main_hand"]
		name := s.bisFacts(d, main.ItemID, main.ItemName, "", "").Name
		return panelRow{text: fill(panelCopy.BisOffHandTwoHand, "{main hand name}", memberText(name))}
	case !has || sl.ItemID <= 0:
		return panelRow{text: fill(panelCopy.BisSlotNone, "{Slot}", c.Name)}
	}
	return panelRow{item: &itemRow{item: s.bisFacts(d, sl.ItemID, sl.ItemName, sl.SourceKind, sl.Source), slot: c.Value, gain: d.pickGain(sl)}}
}

// bisSlotPanel is /bis slot:{slot}: the pick, then the next two with their deltas against it.
func (s *Service) bisSlotPanel(ctx context.Context, d bisData, slot string) (Message, bool, error) {
	pick, has := d.bySlot[slot]
	if !has || pick.ItemID <= 0 {
		return Message{}, false, nil
	}
	header := fill(panelCopy.BisSlotTitle, append(d.who(), "{Slot}", slotLabel(slot))...) + "\n" + d.headLine()
	rows := []panelRow{{item: &itemRow{item: s.bisFacts(d, pick.ItemID, pick.ItemName, pick.SourceKind, pick.Source), slot: slot, gain: d.pickGain(pick)}},
		{node: pointerTo(separatorNode())}, {text: panelCopy.BisSlotNext}}
	for i, alt := range pick.Alternatives {
		if i == slotAlternatives {
			break
		}
		gain := ""
		switch {
		case pick.EffectUnmodelled || alt.EffectUnmodelled:
			gain = panelCopy.NotModelled
		case math.Abs(alt.DpsDelta) >= minShownGain:
			gain = d.figure(alt.DpsDelta)
		}
		rows = append(rows, panelRow{item: &itemRow{item: s.bisFacts(d, alt.ItemID, alt.ItemName, alt.SourceKind, alt.Source), slot: slot, gain: gain}})
	}
	body := s.guardedBody(ctx, d.build, []Node{sectionNode(header, s.crestThumb(d.spec.ClassSlug)), separatorNode()}, rows)
	link := Component{Kind: KindButton, Style: StyleLink, Label: lookupCopy.BisOpen, URL: d.page}
	return Message{Layout: []Node{containerNode(classColorOf(d.spec.ClassSlug), body...)}, Components: []ActionRow{{link}}}, true, nil
}

func pointerTo[T any](v T) *T { return &v }

// panelRow is one entry of a panel's body: an item row, a line of text, or a fixed node.
type panelRow struct {
	item *itemRow
	text string
	node *Node
}

// guardedBody is head then rows, dropping line 3's source from item rows, last row first,
// while the body's text would pass textBudget (spec §A.6). head and fixed lines always stay.
func (s *Service) guardedBody(ctx context.Context, build string, head []Node, rows []panelRow) []Node {
	draw := func() []Node {
		out := append(make([]Node, 0, len(head)+len(rows)), head...)
		for _, r := range rows {
			switch {
			case r.item != nil:
				out = append(out, s.itemSection(ctx, build, *r.item))
			case r.node != nil:
				out = append(out, *r.node)
			default:
				out = append(out, textNode(r.text))
			}
		}
		return out
	}
	out := draw()
	for i := len(rows) - 1; i >= 0 && textLength(out) > textBudget; i-- {
		if rows[i].item != nil && !rows[i].item.noSource {
			rows[i].item.noSource = true
			out = draw()
		}
	}
	return out
}

// withPageButton puts the other page's button first on the answer's first row (before Change
// picks); a slot answer has none.
func withPageButton(msg Message, opts map[string]string, p bisPicks) Message {
	if p.Slot != "" || len(msg.Components) == 0 {
		return msg
	}
	other := lastPage
	if pageOf(opts) == lastPage {
		other = firstPage
	}
	button := Component{Kind: KindButton, Style: StyleSecondary, Label: panelCopy.BisPages[other-1],
		CustomID: cardCustomID(prefixBis, partPage, append(p.state(), strconv.Itoa(other))...)}
	rows := append([]ActionRow(nil), msg.Components...)
	rows[0] = append(ActionRow{button}, rows[0]...)
	msg.Components = rows
	return msg
}

// bisPageOf is the page a press's state names (its fifth field); page 1 otherwise.
func bisPageOf(state []string) string {
	if len(state) > bisStateSlot+1 && state[bisStateSlot+1] == strconv.Itoa(lastPage) {
		return strconv.Itoa(lastPage)
	}
	return strconv.Itoa(firstPage)
}

// bisPage answers the page button: the other page in place, with the message's own Share
// button copied unchanged (a token is minted only when the message carries none).
func (s *Service) bisPage(ctx context.Context, in Interaction, press cardPress) (Reply, error) {
	return s.bisPresser(false, func(ctx context.Context, in Interaction, seat bisSeat, p bisPicks) (Reply, error) {
		key, ok := p.specKey()
		if !ok {
			rep := failureReply()
			rep.Update = true
			return rep, nil
		}
		opts := bisOptions(key, p)
		msg, found, err := s.answerBis(ctx, seat.guildID, seat.viewer, withOption(opts, optPage, bisPageOf(press.State)))
		if err != nil {
			return Reply{}, err
		}
		if !found {
			rep := failureReply()
			rep.Update = true
			return rep, nil
		}
		share, has := sourceShare(in.Source.Components)
		if !has {
			if share, err = s.offerShare(ctx, in, seat.guildID, "bis", opts); err != nil {
				return Reply{}, err
			}
		}
		msg = withShare(decorateOwn("bis", withOption(opts, optPage, bisPageOf(press.State)), msg), share)
		msg.Ephemeral = true
		return Reply{Message: msg, Update: true}, nil
	})(ctx, in, press)
}

// withOption is opts with key set, as a copy.
func withOption(opts map[string]string, key, value string) map[string]string {
	out := make(map[string]string, len(opts)+1)
	for k, v := range opts {
		out[k] = v
	}
	out[key] = value
	return out
}

// wireButton is a button as a message's components carry it.
type wireButton struct {
	Type       ComponentKind     `json:"type"`
	Style      ButtonStyle       `json:"style"`
	Label      string            `json:"label"`
	CustomID   string            `json:"custom_id"`
	Disabled   bool              `json:"disabled"`
	Components []json.RawMessage `json:"components"`
	Accessory  json.RawMessage   `json:"accessory"`
}

// sourceShare is the Share in channel button a message carries, anywhere in its components.
func sourceShare(raw json.RawMessage) (Component, bool) {
	var list []json.RawMessage
	if json.Unmarshal(raw, &list) != nil {
		return Component{}, false
	}
	for _, item := range list {
		var b wireButton
		if json.Unmarshal(item, &b) != nil {
			continue
		}
		if b.Type == KindButton && strings.HasPrefix(b.CustomID, shareIDPrefix+":") {
			return Component{Kind: KindButton, Style: b.Style, Label: b.Label, CustomID: b.CustomID, Disabled: b.Disabled}, true
		}
		if nested, err := json.Marshal(b.Components); err == nil && len(b.Components) > 0 {
			if c, ok := sourceShare(nested); ok {
				return c, true
			}
		}
	}
	return Component{}, false
}

// sourcePage is the page a pressed /bis answer shows, read from its page button (which names
// the other page); "" when it has none.
func sourcePage(raw json.RawMessage) string {
	var list []json.RawMessage
	if json.Unmarshal(raw, &list) != nil {
		return ""
	}
	for _, item := range list {
		var b wireButton
		if json.Unmarshal(item, &b) != nil {
			continue
		}
		if p, ok := parseCardID(b.CustomID); ok && p.Prefix == prefixBis && p.Part == partPage {
			if bisPageOf(p.State) == strconv.Itoa(lastPage) {
				return strconv.Itoa(firstPage)
			}
			return strconv.Itoa(lastPage)
		}
		if nested, err := json.Marshal(b.Components); err == nil && len(b.Components) > 0 {
			if page := sourcePage(nested); page != "" {
				return page
			}
		}
	}
	return ""
}
