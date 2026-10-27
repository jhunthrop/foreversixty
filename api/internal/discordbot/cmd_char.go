package discordbot

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/bis"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
	"github.com/jhunthrop/foreversixty/api/internal/rankings"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

// optName is /char's character option.
const optName = "name"

// maxUpgrades is how many upgrade item embeds follow the card.
const maxUpgrades = 3

// tankRole is the spec role whose upgrades are ranked by score, not DPS.
const tankRole = "tank"

// cmdChar answers /char: the public character card and its top upgrades, ephemeral, with
// Share in channel. A private or unknown name answers the one sentence.
func (s *Service) cmdChar(ctx context.Context, in Interaction) (Reply, error) {
	guildID, viewer, bound, err := s.lookupCaller(ctx, in)
	if err != nil || !bound {
		return pressReply(commonCopy.Unbound, false), err
	}
	name, _ := in.Option(optName)
	opts := map[string]string{optName: name}
	msg, found, err := s.answerChar(ctx, guildID, viewer, opts)
	if err != nil {
		return Reply{}, err
	}
	if !found {
		return pressReply(lookupCopy.NoCharacter, false), nil
	}
	button, err := s.offerShare(ctx, guildID, "char", opts)
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

// answerChar is the /char answer for viewer (nil: a visitor), resolving the name within the
// guild's region and ruleset. Readiness and Synced are the guild's own reads: only a verified
// member of the guild sees them.
func (s *Service) answerChar(ctx context.Context, guildID int64, viewer *int64, opts map[string]string) (Message, bool, error) {
	g, err := s.d.Guilds.BotGuildFacts(ctx, guildID)
	if err != nil {
		return Message{}, false, err
	}
	card, found, err := s.d.Characters.CharacterCard(ctx, g.Region, g.Ruleset, opts[optName])
	if err != nil || !found {
		return Message{}, false, err
	}
	spec, _ := specByKey(deref(card.GearSpec))
	head := s.cardEmbed(card, spec)
	member, err := s.verifiedMember(ctx, guildID, viewer)
	if err != nil {
		return Message{}, false, err
	}
	if member {
		if head.Fields, err = s.memberFields(ctx, guildID, card.Character.Key, head.Fields); err != nil {
			return Message{}, false, err
		}
	}
	head.Fields = append(head.Fields, upgradesField(card, spec))
	embeds := append([]notify.Embed{head}, s.upgradeEmbeds(card, spec, g.Faction)...)
	closeEmbeds(embeds, lookupCopy.Footer, s.now())
	link := Component{Kind: KindButton, Style: StyleLink, Label: lookupCopy.CharacterPage, URL: s.characterURL(card.Character.Key)}
	return Message{Embeds: embeds, Components: []ActionRow{{link}}}, true, nil
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

// cardEmbed is the card: name, class colour, crest, the level line and the item level.
func (s *Service) cardEmbed(card rankings.Character, spec specs.Spec) notify.Embed {
	class := strings.ToLower(card.Character.Class)
	e := notify.Embed{
		Title: memberText(card.Character.Name), URL: s.characterURL(card.Character.Key), Color: classColors[class],
		Description: cardLine(card, spec, class),
	}
	if _, known := classColors[class]; known {
		e.Thumbnail = s.d.PublicBaseURL + "/icons/hd/crests/" + class + ".webp"
	}
	if level, ok := averageItemLevel(card.Gear); ok {
		e.Fields = append(e.Fields, notify.EmbedField{Name: lookupCopy.ItemLevel, Value: level, Inline: true})
	}
	return e
}

// cardLine is "Level 60 Marksmanship Hunter · Olympus XXVII", each part only when known.
func cardLine(card rankings.Character, spec specs.Spec, class string) string {
	who, guildPart, _ := strings.Cut(lookupCopy.CharLine, lineJoin)
	level := ""
	if card.Character.Level != nil {
		level = fmt.Sprint(*card.Character.Level)
	} else {
		who = strings.TrimPrefix(who, "Level {n}")
	}
	line := strings.Join(strings.Fields(fill(who, "{n}", level, "{Spec}", spec.Name, "{Class}", classLabel(class))), " ")
	if card.Character.Guild != nil && card.Character.Guild.Name != "" {
		line += lineJoin + fill(guildPart, "{Guild}", memberText(card.Character.Guild.Name))
	}
	return line
}

// averageItemLevel is the mean item level of the worn items the data knows, one decimal.
func averageItemLevel(gear []rankings.GearSlot) (string, bool) {
	total, count := 0, 0
	for _, g := range gear {
		if g.Worn != nil && g.Worn.ItemLevel != nil {
			total += *g.Worn.ItemLevel
			count++
		}
	}
	if count == 0 {
		return "", false
	}
	return fmt.Sprintf("%.1f", float64(total)/float64(count)), true
}

// memberFields adds Readiness and Synced, the readiness board's word and the export's time.
func (s *Service) memberFields(ctx context.Context, guildID int64, key string, fields []notify.EmbedField) ([]notify.EmbedField, error) {
	facts, err := s.d.Guilds.CharacterReadinessFacts(ctx, guildID, key)
	if err != nil {
		return nil, fmt.Errorf("discordbot: char readiness: %w", err)
	}
	if !facts.OnRoster {
		return fields, nil
	}
	fields = append(fields, notify.EmbedField{Name: lookupCopy.Readiness, Value: facts.Word, Inline: true})
	if facts.SyncedAt != nil {
		fields = append(fields, notify.EmbedField{Name: lookupCopy.Synced, Value: fmt.Sprintf("<t:%d:R>", facts.SyncedAt.Unix()), Inline: true})
	}
	return fields, nil
}

func upgradesField(card rankings.Character, spec specs.Spec) notify.EmbedField {
	value := lookupCopy.UpgradesByDPS
	switch {
	case len(card.Gear) == 0:
		value = lookupCopy.NoGear
	case spec.Role == tankRole:
		value = lookupCopy.UpgradesByScore
	}
	return notify.EmbedField{Name: lookupCopy.TopUpgrades, Value: value}
}

// upgradeEmbeds are the gear panel's three largest gaps, each the pick's item embed with its
// gain: the DPS the BiS data measured over the worn item when it did, else the score points.
func (s *Service) upgradeEmbeds(card rankings.Character, spec specs.Spec, faction string) []notify.Embed {
	slots := upgradeSlots(card.Gear)
	if len(slots) == 0 {
		return nil
	}
	build, _ := s.d.Guilds.ActiveBuild()
	measured := s.measuredBand(card, spec, faction, build)
	class := strings.ToLower(card.Character.Class)
	out := make([]notify.Embed, 0, len(slots))
	for _, g := range slots {
		delta := signedPoints(*g.Gap)
		if pick, ok := measured[g.Slot]; ok && spec.Role != tankRole {
			if v := bis.VerdictFor(pick, wornID(g), g.Worn != nil); v.GainDps != nil {
				delta = signedDPS(*v.GainDps)
			}
		}
		url := s.characterURL(card.Character.Key) + "#gear-slot-" + g.Slot
		facts := s.itemFacts(build, class, g.Pick.ID, g.Pick.Name, g.Pick.Source, url)
		out = append(out, ItemEmbed(facts, slotLabel(g.Slot), delta, ""))
	}
	return out
}

// upgradeSlots are the slots with a pick that gains, largest gap first, at most three.
func upgradeSlots(gear []rankings.GearSlot) []rankings.GearSlot {
	var out []rankings.GearSlot
	for _, g := range gear {
		if g.Pick != nil && g.Gap != nil && *g.Gap > 0 && (g.Worn == nil || g.Worn.ID != g.Pick.ID) {
			out = append(out, g)
		}
	}
	sort.SliceStable(out, func(i, j int) bool { return *out[i].Gap > *out[j].Gap })
	if len(out) > maxUpgrades {
		out = out[:maxUpgrades]
	}
	return out
}

func wornID(g rankings.GearSlot) int {
	if g.Worn == nil {
		return 0
	}
	return g.Worn.ID
}

// measuredBand is the character's bare band by slot, the one the gear panel's picks come from;
// empty when any part is unknown.
func (s *Service) measuredBand(card rankings.Character, spec specs.Spec, faction, build string) map[string]bis.Slot {
	band, ok := levelBand(card.Character.Level)
	if spec.Spec == "" || faction == "" || build == "" || !ok {
		return nil
	}
	dir := s.d.Guilds.BisDataDir()
	b, found, err := bis.LoadPresetBand(dir, bis.SourceBuild(dir, build), spec.Spec, band, faction, bis.BarePreset)
	if err != nil {
		s.log.Warn("discordbot", "op", "char_band", "spec", spec.Spec, "err", err)
	}
	if !found {
		return nil
	}
	return b.BySlot()
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
