package discordbot

import (
	"context"
	"slices"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/notify"
	"github.com/jhunthrop/foreversixty/api/internal/upgrades"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

// bisPicks is the /bis card's state: what the member has chosen so far.
type bisPicks struct{ Class, Spec, Range, Slot string }

// The positions of a /bis card's state fields in its custom id.
const (
	bisStateClass = iota
	bisStateSpec
	bisStateRange
	bisStateSlot
)

// state is the picks as card-id fields: class, spec, range, slot.
func (p bisPicks) state() []string { return []string{p.Class, p.Spec, p.Range, p.Slot} }

// bisPicksOf reads a card's state. It trusts nothing: a class, spec, range or slot that does
// not exist (or a spec of another class) reads as empty, a bad range as the Phase 1 raid,
// and a missing field as empty.
func bisPicksOf(state []string) bisPicks {
	at := func(i int) string {
		if i < len(state) {
			return state[i]
		}
		return ""
	}
	p := bisPicks{Class: at(bisStateClass), Spec: at(bisStateSpec), Range: at(bisStateRange), Slot: at(bisStateSlot)}
	if !choiceHas(classChoices, p.Class) {
		p.Class = ""
	}
	if _, ok := specByKey(p.Class + "-" + p.Spec); !ok {
		p.Spec = ""
	}
	if _, ok := bandLabels[p.Range]; !ok {
		p.Range = phaseRaid
	}
	if !choiceHas(slotChoices, p.Slot) {
		p.Slot = ""
	}
	return p
}

func choiceHas(choices []Choice, value string) bool {
	return value != "" && slices.ContainsFunc(choices, func(c Choice) bool { return c.Value == value })
}

// specKey is the picks' spec key ("hunter-beast-mastery"); ok is false without a spec.
func (p bisPicks) specKey() (string, bool) {
	if p.Class == "" || p.Spec == "" {
		return "", false
	}
	return p.Class + "-" + p.Spec, true
}

// sameBuild is whether two picks name the same class, spec and range.
func (p bisPicks) sameBuild(o bisPicks) bool {
	return p.Class == o.Class && p.Spec == o.Spec && p.Range == o.Range
}

// bisMain is the viewer's main character as /bis picks: its name and class, spec and range.
type bisMain struct {
	Name  string
	Picks bisPicks
}

// mainPicks is the viewer's main as /bis picks; nil without a main.
func (s *Service) mainPicks(ctx context.Context, guildID int64, viewer *int64) (*bisMain, error) {
	main, ok, err := s.mainOf(ctx, guildID, viewer)
	if err != nil || !ok {
		return nil, err
	}
	rng := phaseRaid
	if band, known := upgrades.LevelBand(main.level); known {
		rng = phaseOfBand(band)
	}
	return &bisMain{Name: main.name, Picks: bisPicks{Class: main.spec.ClassSlug, Spec: main.spec.SpecSlug, Range: rng}}, nil
}

// bisCard is everything the /bis card draws.
type bisCard struct {
	Picks  bisPicks
	Lead   string
	Main   *bisMain
	Linked bool
}

// bisLeadFor is the card's lead line for its picks.
func bisLeadFor(p bisPicks, main *bisMain, linked bool) string {
	switch {
	case main != nil && p.Spec != "" && p.sameBuild(main.Picks):
		return fill(guideCopy.BisLeadMain, "{Name}", memberText(main.Name))
	case p.Spec != "":
		return guideCopy.BisLeadReady
	case p.Class != "":
		return fill(guideCopy.BisLeadClass, "{Class}", classLabel(p.Class))
	case p.Slot != "":
		return fill(guideCopy.BisLeadSlot, "{Slot}", slotLabel(p.Slot))
	case !linked:
		return guideCopy.BisLeadUnlinked
	}
	return guideCopy.BisLeadNone
}

// bisCardMessage is the /bis card: its heading and lead in a container, and four rows, only
// the member sees it. It is Components V2 from its first frame (spec 2026-10-29 §A.7).
func (s *Service) bisCardMessage(c bisCard) Message {
	rows := []ActionRow{
		{bisClassSelect(c.Picks)},
		{bisSpecSelect(c.Picks)},
		{bisRangeSelect(c.Picks)},
		s.bisButtons(c),
	}
	card := Message{Embeds: []notify.Embed{s.bisCardEmbed(c)}, Components: rows, Ephemeral: true}
	return asComponentsV2(card, crestAlt(c.Picks.Class))
}

func (s *Service) bisCardEmbed(c bisCard) notify.Embed {
	p := c.Picks
	e := notify.Embed{Author: guideCopy.BisAuthor, Description: c.Lead, Color: goldColor, Title: guideCopy.BisTitleNone}
	if p.Class == "" {
		return e
	}
	e.Color = classColorOf(p.Class)
	e.Thumbnail = s.d.PublicBaseURL + "/icons/hd/crests/" + p.Class + ".webp"
	e.Title = fill(guideCopy.BisTitleClass, "{Class}", classLabel(p.Class))
	if key, ok := p.specKey(); ok {
		sp, _ := specByKey(key)
		e.Title = fill(guideCopy.BisTitleReady, "{Spec}", sp.Name, "{Class}", classLabel(p.Class), "{band label}", bandLabels[p.Range])
	}
	return e
}

func bisClassSelect(p bisPicks) Component {
	opts := make([]SelectOption, len(classChoices))
	for i, c := range classChoices {
		opts[i] = SelectOption{Label: c.Name, Value: c.Value, Description: specNamesOf(c.Value), Default: c.Value == p.Class}
	}
	return Component{Kind: KindStringSelect, CustomID: cardCustomID(prefixBis, "c", p.state()...),
		Placeholder: guideCopy.BisPlaceholderClass, Options: opts}
}

// specNamesOf is a class's spec names in catalogue order, joined by ", ".
func specNamesOf(class string) string {
	var names []string
	for _, sp := range specs.All {
		if sp.ClassSlug == class {
			names = append(names, sp.Name)
		}
	}
	return strings.Join(names, ", ")
}

func bisSpecSelect(p bisPicks) Component {
	id := cardCustomID(prefixBis, "s", p.state()...)
	if p.Class == "" {
		return lockedSelect(id, guideCopy.BisPlaceholderSpecLocked)
	}
	var opts []SelectOption
	for _, sp := range specs.All {
		if sp.ClassSlug == p.Class {
			opts = append(opts, SelectOption{Label: sp.Name, Value: sp.SpecSlug,
				Description: roleWord(sp.Role), Default: sp.SpecSlug == p.Spec})
		}
	}
	return Component{Kind: KindStringSelect, CustomID: id, Placeholder: guideCopy.BisPlaceholderSpec, Options: opts}
}

func bisRangeSelect(p bisPicks) Component {
	opts := make([]SelectOption, len(phaseChoices))
	for i, c := range phaseChoices {
		opts[i] = SelectOption{Label: c.Name, Value: c.Value, Description: guideCopy.BisRangeDesc[i], Default: c.Value == p.Range}
	}
	return Component{Kind: KindStringSelect, CustomID: cardCustomID(prefixBis, "r", p.state()...),
		Placeholder: guideCopy.BisPlaceholderRange, Options: opts}
}

// bisButtons are Show BiS, Use {Name} when a pick differs from the main's, and Link Discord
// for a viewer who is not linked.
func (s *Service) bisButtons(c bisCard) ActionRow {
	row := ActionRow{{Kind: KindButton, Style: StyleSuccess, Label: guideCopy.BisShow,
		CustomID: cardCustomID(prefixBis, "go", c.Picks.state()...)}}
	if c.Main != nil && !c.Picks.sameBuild(c.Main.Picks) {
		row = append(row, Component{Kind: KindButton, Style: StyleSecondary,
			Label:    fill(guideCopy.BisUseMain, "{Name}", c.Main.Name),
			CustomID: cardCustomID(prefixBis, "me", c.Picks.state()...)})
	}
	if !c.Linked {
		row = append(row, s.linkButton())
	}
	return row
}

// bisCardFor is the card a /bis without a usable spec opens with: the typed options filled
// in, the rest from the viewer's main (unless the typed spec was unknown), and the lead.
func (s *Service) bisCardFor(ctx context.Context, in Interaction, guildID int64, viewer *int64, unknownSpec bool) (Message, error) {
	main, err := s.mainPicks(ctx, guildID, viewer)
	if err != nil {
		return Message{}, err
	}
	class, _ := in.Option(optClass)
	phase, _ := in.Option(optPhase)
	slot, _ := in.Option(optSlot)
	p := bisPicksOf([]string{class, "", phase, slot})
	if _, typed := bandLabels[phase]; !typed {
		p.Range = phaseRaid
	}
	if main != nil && !unknownSpec {
		p = fromMain(p, class, phase, main.Picks)
	}
	lead := bisLeadFor(p, main, viewer != nil)
	if unknownSpec {
		lead = unknownSpecLead(p.Class)
	}
	return s.bisCardMessage(bisCard{Picks: p, Lead: lead, Main: main, Linked: viewer != nil}), nil
}

// fromMain fills a card's picks the member did not type from their main: the class, the
// spec when the class is the main's, and the range when no valid one was typed.
func fromMain(p bisPicks, typedClass, typedPhase string, main bisPicks) bisPicks {
	if typedClass == "" || p.Class == "" {
		p.Class = main.Class
	}
	if p.Class != main.Class {
		return p
	}
	p.Spec = main.Spec
	if _, typed := bandLabels[typedPhase]; !typed {
		p.Range = main.Range
	}
	return p
}

func unknownSpecLead(class string) string {
	if class == "" {
		return guideCopy.BisLeadUnknownSpecNoClass
	}
	return fill(guideCopy.BisLeadUnknownSpec, "{Class}", classLabel(class))
}
