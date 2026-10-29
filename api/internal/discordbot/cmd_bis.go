package discordbot

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/bis"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

// /bis's options.
const (
	optClass = "class"
	optSpec  = "spec"
	optSlot  = "slot"
	optPhase = "phase"
)

// phaseRaid is the Phase 1 raid choice: level 60 with the raid-ready preset.
const phaseRaid = "p1"

// defaultFaction is the faction a guild with none resolved reads, the site's own default.
const defaultFaction = guilds.FactionAlliance

// topSlots is how many slots /bis lists without a slot; slotAlternatives how many
// alternatives follow a slot's pick.
const (
	topSlots         = 5
	slotAlternatives = 2
)

// cmdBis answers /bis: one slot's pick and two alternatives, or the five slots where the pick
// gains most over the next best, ephemeral, with Share in channel. Omitted options come from
// the linked viewer's main character (spec §3.3.3, ruling 9).
func (s *Service) cmdBis(ctx context.Context, in Interaction) (Reply, error) {
	guildID, viewer, bound, err := s.lookupCaller(ctx, in)
	if err != nil || !bound {
		return pressReply(commonCopy.Unbound, false), err
	}
	specText, _ := in.Option(optSpec)
	if specText == "" {
		return s.bisCardReply(ctx, in, guildID, viewer, false)
	}
	opts, ok, err := s.resolveBis(ctx, in, guildID, viewer)
	if err != nil {
		return Reply{}, err
	}
	if !ok {
		return s.bisCardReply(ctx, in, guildID, viewer, true)
	}
	return s.bisAnswer(ctx, in, guildID, viewer, opts, false)
}

// bisCardReply opens the /bis card.
func (s *Service) bisCardReply(ctx context.Context, in Interaction, guildID int64, viewer *int64, unknownSpec bool) (Reply, error) {
	msg, err := s.bisCardFor(ctx, in, guildID, viewer, unknownSpec)
	return Reply{Message: msg}, err
}

// mainFacts is what /bis defaults from: the viewer's main character in the guild.
type mainFacts struct {
	name  string
	spec  specs.Spec
	level *int
}

// resolveBis turns the options into a spec key, a phase and the slot. ok is false when no
// spec can be named.
func (s *Service) resolveBis(ctx context.Context, in Interaction, guildID int64, viewer *int64) (map[string]string, bool, error) {
	class, _ := in.Option(optClass)
	specText, _ := in.Option(optSpec)
	slot, _ := in.Option(optSlot)
	phase, _ := in.Option(optPhase)
	main, hasMain, err := s.mainOf(ctx, guildID, viewer)
	if err != nil {
		return nil, false, err
	}
	chosen, ok := resolveSpec(class, specText)
	if !ok {
		return nil, false, nil
	}
	if _, valid := bandLabels[phase]; !valid {
		phase = phaseRaid
		if band, known := levelBand(main.level); hasMain && known && main.spec.ClassSlug == chosen.ClassSlug {
			phase = phaseOfBand(band)
		}
	}
	return map[string]string{optSpec: chosen.Spec, optPhase: phase, optSlot: slot}, true, nil
}

// phaseOfBand is a level band's phase choice; level 60 reads the Phase 1 raid (the site's
// default preset at 60 is raid-ready).
func phaseOfBand(band int) string {
	if band == levelBands[len(levelBands)-1] {
		return phaseRaid
	}
	return fmt.Sprint(band)
}

// mainOf is the viewer's main character in the guild (else its first), with its spec and level.
func (s *Service) mainOf(ctx context.Context, guildID int64, viewer *int64) (mainFacts, bool, error) {
	if viewer == nil {
		return mainFacts{}, false, nil
	}
	cands, err := s.d.Guilds.SignupCandidates(ctx, guildID, *viewer)
	if err != nil {
		return mainFacts{}, false, fmt.Errorf("discordbot: bis main: %w", err)
	}
	verified := verifiedOf(cands)
	if len(verified) == 0 {
		return mainFacts{}, false, nil
	}
	main := verified[0] // the main sorts first
	key, ok := bis.FileSlugFor(main.Class, main.Spec)
	if !ok {
		return mainFacts{}, false, nil
	}
	sp, _ := specByKey(key)
	out := mainFacts{spec: sp, name: main.Name}
	if parts := strings.SplitN(main.CharacterKey, "/", 3); len(parts) == 3 && s.d.Characters != nil {
		card, found, err := s.d.Characters.CharacterCard(ctx, parts[0], parts[1], parts[2])
		if err != nil {
			return mainFacts{}, false, fmt.Errorf("discordbot: bis main level: %w", err)
		}
		if found {
			out.level = card.Character.Level
		}
	}
	return out, true, nil
}

// resolveSpec reads the spec option: a spec key from the autocomplete ("warrior-fury"), or a
// spec's name or slug, unique within class when one is given.
func resolveSpec(class, text string) (specs.Spec, bool) {
	text = strings.TrimSpace(text)
	if sp, ok := specByKey(strings.ToLower(text)); ok {
		return sp, class == "" || sp.ClassSlug == class
	}
	var found []specs.Spec
	for _, sp := range specs.All {
		if (class == "" || sp.ClassSlug == class) && (strings.EqualFold(sp.Name, text) || strings.EqualFold(sp.SpecSlug, text)) {
			found = append(found, sp)
		}
	}
	if len(found) != 1 {
		return specs.Spec{}, false
	}
	return found[0], true
}

// answerBis is the /bis answer for resolved options. It is the same for every viewer: the
// BiS lists are public.
func (s *Service) answerBis(ctx context.Context, guildID int64, _ *int64, opts map[string]string) (Message, bool, error) {
	sp, ok := specByKey(opts[optSpec])
	label, known := bandLabels[opts[optPhase]]
	if !ok || !known {
		return Message{}, false, nil
	}
	g, err := s.d.Guilds.BotGuildFacts(ctx, guildID)
	if err != nil {
		return Message{}, false, err
	}
	faction := g.Faction
	if faction == "" {
		faction = defaultFaction
	}
	band, preset := phaseBand(opts[optPhase])
	build, _ := s.d.Guilds.ActiveBuild()
	data, found, err := bis.LoadPresetBand(s.d.Guilds.BisDataDir(), s.d.Guilds.BisBuild(), sp.Spec, band, faction, preset)
	if err != nil || !found {
		return Message{}, false, err
	}
	page := s.bisPageURL(sp, faction, band, preset)
	view := bisView{s: s, spec: sp, build: build, page: page, faction: faction}
	var embeds []notify.Embed
	lead, button := "", ""
	if slot := opts[optSlot]; slot != "" {
		pick, has := data.BySlot()[slot]
		if !has || pick.ItemID <= 0 {
			return Message{}, false, nil
		}
		lead, button, embeds = fill(lookupCopy.BisSlotLead, "{Slot}", slotLabel(slot)), lookupCopy.BisOpen, view.slotEmbeds(pick)
	} else {
		lead, button, embeds = lookupCopy.BisTopLead, lookupCopy.BisAllSlots, view.topEmbeds(data)
	}
	who := []string{"{Spec}", sp.Name, "{Class}", classLabel(sp.ClassSlug), "{band label}", label}
	closeEmbeds(embeds, fill(lookupCopy.BisFooter, who...), s.now())
	link := Component{Kind: KindButton, Style: StyleLink, Label: button, URL: page}
	return Message{Content: fill(lead, who...), Embeds: embeds, Components: []ActionRow{{link}}}, true, nil
}

// phaseBand is a phase choice's level band and preset: the Phase 1 raid is level 60
// raid-ready; the bare levels the bare preset.
func phaseBand(phase string) (int, string) {
	if phase == phaseRaid {
		return levelBands[len(levelBands)-1], bis.RaidPreset
	}
	var band int
	_, _ = fmt.Sscan(phase, &band)
	return band, bis.BarePreset
}

// bisPageURL is the spec's BiS page at the band, faction and preset.
func (s *Service) bisPageURL(sp specs.Spec, faction string, band int, preset string) string {
	return fmt.Sprintf("%s/bis/%s/%s?faction=%s&preset=%s#band-%s-%d",
		s.d.PublicBaseURL, sp.ClassSlug, sp.SpecSlug, faction, preset, faction, band)
}

type bisView struct {
	s       *Service
	spec    specs.Spec
	build   string
	page    string
	faction string
}

// facts is an item of the band with the BiS page's own source line for its kind.
func (v bisView) facts(id int, name, kind, raw string) ItemFacts {
	return v.s.itemFacts(v.build, v.spec.ClassSlug, id, name, v.s.sourceLine(id, kind, v.faction, raw), v.page)
}

// slotEmbeds are the pick under "Best in slot" and its first two alternatives with their deltas.
func (v bisView) slotEmbeds(pick bis.Slot) []notify.Embed {
	label := slotLabel(pick.Slot)
	out := []notify.Embed{ItemEmbed(v.facts(pick.ItemID, pick.ItemName, pick.SourceKind, pick.Source), label, "", lookupCopy.BisAuthor)}
	for i, alt := range pick.Alternatives {
		if i == slotAlternatives {
			break
		}
		out = append(out, ItemEmbed(v.facts(alt.ItemID, alt.ItemName, alt.SourceKind, alt.Source), label, signedDPS(alt.DpsDelta), ""))
	}
	return out
}

// topEmbeds are the five slots whose pick beats its first alternative by the most DPS (ruling 4).
func (v bisView) topEmbeds(band bis.Band) []notify.Embed {
	var ranked []bis.Slot
	for _, sl := range band.Slots {
		if sl.ItemID > 0 && len(sl.Alternatives) > 0 {
			ranked = append(ranked, sl)
		}
	}
	sort.SliceStable(ranked, func(i, j int) bool {
		return ranked[i].Alternatives[0].DpsDelta < ranked[j].Alternatives[0].DpsDelta
	})
	if len(ranked) > topSlots {
		ranked = ranked[:topSlots]
	}
	out := make([]notify.Embed, 0, len(ranked))
	for _, sl := range ranked {
		gain := signedDPS(-sl.Alternatives[0].DpsDelta)
		out = append(out, ItemEmbed(v.facts(sl.ItemID, sl.ItemName, sl.SourceKind, sl.Source), slotLabel(sl.Slot), gain, ""))
	}
	return out
}

// autocompleteChar suggests the guild's public characters for /char's name.
func (s *Service) autocompleteChar(ctx context.Context, in Interaction) ([]Choice, error) {
	focused, typed := in.Focused()
	if focused != optName {
		return []Choice{}, nil
	}
	guildID, bound, err := s.serverGuild(ctx, in)
	if err != nil || !bound {
		return []Choice{}, err
	}
	names, err := s.d.Guilds.PublicRosterNames(ctx, guildID, typed, maxChoices)
	if err != nil {
		return nil, fmt.Errorf("discordbot: char autocomplete: %w", err)
	}
	out := []Choice{}
	for _, n := range names {
		if len([]rune(n)) <= maxChoiceRunes {
			out = append(out, Choice{Name: n, Value: n})
		}
	}
	return out, nil
}

// autocompleteSpec suggests /bis's spec, filtered by the class option when given; the value
// is the spec key, so a spec resolves its class.
func (s *Service) autocompleteSpec(_ context.Context, in Interaction) ([]Choice, error) {
	focused, typed := in.Focused()
	if focused != optSpec {
		return []Choice{}, nil
	}
	class, _ := in.Option(optClass)
	out := []Choice{}
	for _, sp := range specs.All {
		name := sp.Name + " " + classLabel(sp.ClassSlug)
		if (class == "" || sp.ClassSlug == class) && matches(name, typed) && len(out) < maxChoices {
			out = append(out, Choice{Name: name, Value: sp.Spec})
		}
	}
	return out, nil
}
