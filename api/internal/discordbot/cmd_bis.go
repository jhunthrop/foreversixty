package discordbot

import (
	"context"
	"fmt"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/bis"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/upgrades"
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

// slotAlternatives is how many alternatives follow a slot's pick.
const slotAlternatives = 2

// cmdBis answers /bis: one slot's pick and two alternatives, or the character frame's first
// page of picks, ephemeral, with Share in channel. Omitted options come from
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
		if band, known := upgrades.LevelBand(main.level); hasMain && known && main.spec.ClassSlug == chosen.ClassSlug {
			phase = phaseOfBand(band)
		}
	}
	return map[string]string{optSpec: chosen.Spec, optPhase: phase, optSlot: slot}, true, nil
}

// phaseOfBand is a level band's phase choice; level 60 reads the Phase 1 raid (the site's
// default preset at 60 is raid-ready).
func phaseOfBand(band int) string {
	if band == upgrades.TopBand {
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

// phaseBand is a phase choice's level band and preset: the Phase 1 raid is level 60
// raid-ready; the bare levels the bare preset.
func phaseBand(phase string) (int, string) {
	if phase == phaseRaid {
		return upgrades.TopBand, bis.RaidPreset
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
