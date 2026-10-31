package discordbot

import (
	"context"
	"fmt"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/rankings"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

// optName is /char's character option.
const optName = "name"

// cmdChar answers /char: the public character card and its top upgrades, ephemeral, with
// Share in channel. A private or unknown name answers the one sentence.
func (s *Service) cmdChar(ctx context.Context, in Interaction) (Reply, error) {
	guildID, viewer, bound, err := s.lookupCaller(ctx, in)
	if err != nil || !bound {
		return pressReply(commonCopy.Unbound, false), err
	}
	name, _ := in.Option(optName)
	if name == "" {
		return s.charStart(ctx, in, guildID, viewer)
	}
	opts := charOptions(name, false)
	msg, found, err := s.answerChar(ctx, guildID, viewer, opts)
	if err != nil {
		return Reply{}, err
	}
	if !found {
		own, err := s.ownChars(ctx, guildID, viewer)
		if err != nil {
			return Reply{}, err
		}
		return s.charUnknownCard(name, viewer, own, false), nil
	}
	button, err := s.offerShare(ctx, in, guildID, "char", opts)
	if err != nil {
		return Reply{}, err
	}
	msg = withShare(msg, button)
	msg.Ephemeral = true
	return Reply{Message: msg}, nil
}

// lookupCaller is the server's guild and the linked viewer (nil when not linked).
func (s *Service) lookupCaller(ctx context.Context, in Interaction) (guildID int64, viewer *int64, bound bool, err error) {
	guildID, bound, err = s.serverGuild(ctx, in)
	if err != nil || !bound {
		return 0, nil, false, err
	}
	userID, linked, err := s.viewer(ctx, in)
	if err != nil {
		return 0, nil, false, fmt.Errorf("discordbot: lookup viewer: %w", err)
	}
	if linked {
		viewer = &userID
	}
	return guildID, viewer, true, nil
}

// charData is a /char answer's character as the panel and the sheet read it. Readiness and
// Synced are the guild's own reads: only a verified member of the guild sees them.
type charData struct {
	card    rankings.Character
	spec    specs.Spec
	class   string
	faction string
	build   string
	facts   *guilds.ReadinessFacts
	ups     upgradeList
}

// loadChar resolves the name within the guild's region and ruleset for viewer (nil: a visitor).
func (s *Service) loadChar(ctx context.Context, guildID int64, viewer *int64, opts map[string]string) (charData, bool, error) {
	g, err := s.d.Guilds.BotGuildFacts(ctx, guildID)
	if err != nil {
		return charData{}, false, err
	}
	card, found, err := s.d.Characters.CharacterCard(ctx, g.Region, g.Ruleset, opts[optName])
	if err != nil || !found {
		return charData{}, false, err
	}
	spec, _ := specByKey(deref(card.GearSpec))
	build, _ := s.d.Guilds.ActiveBuild()
	d := charData{card: card, spec: spec, class: strings.ToLower(card.Character.Class),
		faction: s.characterFaction(card, g.Faction), build: build, ups: s.upgrades(card, spec, g.Faction)}
	member, err := s.verifiedMember(ctx, guildID, viewer)
	if err != nil || !member {
		return d, true, err
	}
	facts, err := s.d.Guilds.CharacterReadinessFacts(ctx, guildID, card.Character.Key)
	if err != nil {
		return charData{}, false, fmt.Errorf("discordbot: char readiness: %w", err)
	}
	if facts.OnRoster {
		d.facts = &facts
	}
	return d, true, nil
}

// answerChar is the /char panel (spec 2026-10-29 §A.5.3): the character, then its top
// upgrades. Its one row is the character page; the presser's copy gains the rest.
func (s *Service) answerChar(ctx context.Context, guildID int64, viewer *int64, opts map[string]string) (Message, bool, error) {
	d, found, err := s.loadChar(ctx, guildID, viewer, opts)
	if err != nil || !found {
		return Message{}, false, err
	}
	return s.charPanel(ctx, d), true, nil
}

// charPanel is answerChar over a character already read.
func (s *Service) charPanel(ctx context.Context, d charData) Message {
	head := []Node{sectionNode(s.charHeader(d, true), s.crestThumb(d.class)), separatorNode(), textNode(upgradesText(d.card, d.ups))}
	var rows []panelRow
	if len(d.card.Gear) > 0 {
		shown := d.ups.items[:min(len(d.ups.items), maxUpgrades)]
		for _, u := range shown {
			rows = append(rows, panelRow{item: pointerTo(s.upgradeRow(d, u))})
		}
		if more := moreText(len(d.ups.items) - len(shown)); more != "" {
			rows = append(rows, panelRow{text: more})
		}
	}
	link := Component{Kind: KindButton, Style: StyleLink, Label: lookupCopy.CharacterPage, URL: s.characterURL(d.card.Character.Key)}
	body := s.guardedBody(ctx, d.build, head, rows)
	return Message{ComponentsV2: true, Layout: []Node{charContainer(d.class, body)}, Components: []ActionRow{{link}}}
}

// charContainer is the panel's container in the class colour, or uncoloured for an unknown class.
func charContainer(class string, body []Node) Node {
	if color, ok := classColor(class); ok {
		return containerNode(color, body...)
	}
	return Node{Kind: KindContainer, Children: body}
}

// charHeader is the linked name and the level line; with facts, the item level, readiness and
// synced subtext, each clause only when known.
func (s *Service) charHeader(d charData, withFacts bool) string {
	lines := []string{heading("[" + memberText(d.card.Character.Name) + "](" + s.characterURL(d.card.Character.Key) + ")"),
		cardLine(d.card, d.spec, d.class)}
	if !withFacts {
		return strings.Join(lines, "\n")
	}
	clauses := strings.Split(strings.TrimPrefix(lookupCopy.CharFacts, subtext("")), lineJoin)
	var kept []string
	if level, ok := averageItemLevel(d.card.Gear); ok {
		kept = append(kept, fill(clauses[0], "{x.x}", level))
	}
	if d.facts != nil {
		kept = append(kept, fill(clauses[1], "{word}", d.facts.Word))
		if d.facts.SyncedAt != nil {
			kept = append(kept, fill(clauses[2], "{unix}", fmt.Sprint(d.facts.SyncedAt.Unix())))
		}
	}
	if len(kept) > 0 {
		lines = append(lines, subtext(strings.Join(kept, lineJoin)))
	}
	return strings.Join(lines, "\n")
}

// upgradeRow is one upgrade as a panel row: the pick, its gain, and what it replaces.
func (s *Service) upgradeRow(d charData, u upgrade) itemRow {
	facts := s.itemFacts(d.build, d.class, u.id, u.name, s.sourceLine(u.id, u.kind, d.faction, u.source), "")
	name, level := s.wornFacts(d, u.worn)
	return itemRow{item: facts, slot: u.slot, gain: u.gain, replaces: replacesLine(name, level)}
}

// wornFacts is a worn item's name and item level: the export's, else the build's table.
func (s *Service) wornFacts(d charData, w *rankings.GearWorn) (string, int) {
	if w == nil {
		return "", 0
	}
	name, level := deref(w.Name), 0
	if w.ItemLevel != nil {
		level = *w.ItemLevel
	}
	if it, ok := s.d.Guilds.BuildItem(d.build, d.class, w.ID); ok {
		if name == "" {
			name = it.Name
		}
		if level == 0 {
			level = it.ItemLevel
		}
	}
	return name, level
}

func (s *Service) verifiedMember(ctx context.Context, guildID int64, viewer *int64) (bool, error) {
	if viewer == nil {
		return false, nil
	}
	cands, err := s.d.Guilds.SignupCandidates(ctx, guildID, *viewer)
	if err != nil {
		return false, fmt.Errorf("discordbot: char member: %w", err)
	}
	return len(verifiedOf(cands)) > 0, nil
}

func (s *Service) characterURL(key string) string { return s.d.PublicBaseURL + "/character/" + key }

// cardLine is "Level 60 Marksmanship Hunter · Olympus XXVII", each part only when known.
func cardLine(card rankings.Character, spec specs.Spec, class string) string {
	return cardLineWith(card, spec, class, memberText)
}

// cardLineWith is cardLine with the guild's name through text: memberText for Discord, as it
// is for the sheet's image.
func cardLineWith(card rankings.Character, spec specs.Spec, class string, text func(string) string) string {
	who, guildPart, _ := strings.Cut(lookupCopy.CharLine, lineJoin)
	level := ""
	if card.Character.Level != nil {
		level = fmt.Sprint(*card.Character.Level)
	} else {
		who = strings.TrimPrefix(who, "Level {n}")
	}
	line := strings.Join(strings.Fields(fill(who, "{n}", level, "{Spec}", spec.Name, "{Class}", classLabel(class))), " ")
	if card.Character.Guild != nil && card.Character.Guild.Name != "" {
		line += lineJoin + fill(guildPart, "{Guild}", text(card.Character.Guild.Name))
	}
	return line
}

// averageItemLevel is the mean item level of every worn item, one decimal (paperdoll spec §3.5).
// An empty slot is not equipped, so it is not counted, nor is the empty off hand under a
// two-hander; the gear list carries no shirt or tabard. When any worn item is unknown to the
// data (no item level) there is no figure at all: the unknown items are the newest, so a mean
// of the rest would be wrong in exactly the case that matters. It is the card's own figure, not
// the guild roster's item_level: the card answers for any public character and reads the same
// worn items the character page's paperdoll draws, so the two agree.
func averageItemLevel(gear []rankings.GearSlot) (string, bool) {
	total, count := 0, 0
	for _, g := range gear {
		if g.Worn == nil {
			continue
		}
		if g.Worn.ItemLevel == nil {
			return "", false
		}
		total += *g.Worn.ItemLevel
		count++
	}
	if count == 0 {
		return "", false
	}
	return fmt.Sprintf("%.1f", float64(total)/float64(count)), true
}

// levelBands are the BiS data's level bands.
var levelBands = []int{20, 30, 40, 50, 60}

// levelBand is the band a level falls in (20 to 29 is 20); false under the first band.
func levelBand(level *int) (int, bool) {
	if level == nil || *level < levelBands[0] {
		return 0, false
	}
	return min(*level/10*10, levelBands[len(levelBands)-1]), true
}

// itemFacts is an item as the build's table knows it, with the BiS data's name and source.
func (s *Service) itemFacts(build, class string, id int, name, source, url string) ItemFacts {
	facts := ItemFacts{ID: id, Name: name, Quality: -1, Source: source, URL: url}
	it, ok := s.d.Guilds.BuildItem(build, class, id)
	if !ok {
		return facts
	}
	facts.Quality, facts.ItemLevel = it.Quality, it.ItemLevel
	if facts.Name == "" {
		facts.Name = it.Name
	}
	if it.Icon != "" {
		facts.Icon = s.d.PublicBaseURL + "/data/" + build + "/icons/" + it.Icon + ".webp"
	}
	return facts
}

func specByKey(key string) (specs.Spec, bool) {
	for _, sp := range specs.All {
		if sp.Spec == key {
			return sp, true
		}
	}
	return specs.Spec{}, false
}

func deref(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
