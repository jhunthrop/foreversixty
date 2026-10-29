package discordbot

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
	"github.com/jhunthrop/foreversixty/api/internal/raidspec"
)

// prefixSpec is the "Which spec tonight?" card (spec 2026-10-28 §P.10).
const prefixSpec = "gspc"

// customIDMax is Discord's custom id bound; a longer id is never drawn.
const customIDMax = 100

// The spec card's parts. Ruling 8 keeps `ask` for a button that opens a modal, so the ask's
// select is `pick` (gspc:1:pick:{raid}:{state}:{key}); `chg` is the Change spec button
// (gspc:1:chg:{raid}:{key}) and `set` the select it opens (gspc:1:set:{raid}:{key}).
const (
	partSpecPick   = "pick"
	partSpecChange = "chg"
	partSpecSet    = "set"
)

// roleDescription is a spec's role as the option says it (bis.role).
func roleDescription(role string) string {
	switch role {
	case raidspec.RoleTank:
		return guideCopy.BisRole[1]
	case raidspec.RoleHealer:
		return guideCopy.BisRole[2]
	}
	return guideCopy.BisRole[0]
}

// needsSpec reports a sign-up that asks the spec first: In, Late or Bench, a class with specs,
// and no spec ever picked for this guild's raids.
func needsSpec(state string, c guilds.SignupCandidate) bool {
	return state != guilds.SignupOut && c.RaidSpec == "" && len(raidspec.ForClass(c.Class)) > 0
}

// changeSpecButtons is `Change spec` under an In, Late or Bench outcome; none for Out, for a
// class with no specs, or for an id Discord would refuse.
func changeSpecButtons(state string, raidID int64, c guilds.SignupCandidate) []Component {
	id := cardCustomID(prefixSpec, partSpecChange, strconv.FormatInt(raidID, 10), c.CharacterKey)
	if state == guilds.SignupOut || len(raidspec.ForClass(c.Class)) == 0 || utf8.RuneCountInString(id) > customIDMax {
		return nil
	}
	return []Component{{Kind: KindButton, Style: StyleSecondary, Label: guideCopy.SpecChange, CustomID: id}}
}

// specCardIn is what one spec card is drawn for.
type specCardIn struct {
	raid   guilds.Raid
	zone   string
	char   guilds.SignupCandidate
	state  string // the asked state; "" for a Change spec card
	update bool
}

// specAsk answers "Which spec tonight?" in place of the sign-up: a new message on a press of
// the post or the reminder, in place after the character select or on the /signup card.
func (s *Service) specAsk(ctx context.Context, a signupAsk, c guilds.SignupCandidate) (Reply, error) {
	raid, zone, ok, err := s.specRaid(ctx, a.GuildID, a.RaidID)
	if err != nil {
		return Reply{}, err
	}
	if !ok {
		return a.reply(pressCopy.Failure), nil
	}
	return s.specCard(specCardIn{raid: raid, zone: zone, char: c, state: a.State, update: a.Update})
}

// specRaid is the raid and the guild's raid zone; ok is false for a raid that is not the guild's.
func (s *Service) specRaid(ctx context.Context, guildID, raidID int64) (guilds.Raid, string, bool, error) {
	raid, found, err := s.d.Guilds.GetRaid(ctx, guildID, raidID)
	if err != nil || !found {
		return guilds.Raid{}, "", false, wrapSpec("raid", err)
	}
	zone, err := s.d.Guilds.GuildRaidTimezone(ctx, guildID)
	if err != nil {
		return guilds.Raid{}, "", false, wrapSpec("zone", err)
	}
	return raid, zone, true, nil
}

func wrapSpec(op string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("discordbot: spec card %s: %w", op, err)
}

// specCard draws the card: the class's colour and crest, the lead, and one select of the class's
// specs in catalogue order, the saved one chosen.
func (s *Service) specCard(in specCardIn) (Reply, error) {
	specs := raidspec.ForClass(in.char.Class)
	if len(specs) == 0 {
		return pressReply(pressCopy.Failure, in.update), nil
	}
	raidRef := strconv.FormatInt(in.raid.ID, 10)
	lead := fill(guideCopy.SpecLeadChange, s.raidPairs(in)...)
	id := cardCustomID(prefixSpec, partSpecSet, raidRef, in.char.CharacterKey)
	if in.state != "" {
		lead = fill(guideCopy.SpecLead, append(s.raidPairs(in), "{State}", stateWord(in.state))...)
		id = cardCustomID(prefixSpec, partSpecPick, raidRef, in.state, in.char.CharacterKey)
	}
	if utf8.RuneCountInString(id) > customIDMax {
		return pressReply(pressCopy.Failure, in.update), nil
	}
	options := make([]SelectOption, len(specs))
	for i, sp := range specs {
		options[i] = SelectOption{Label: sp.Name, Value: sp.Key, Description: roleDescription(sp.Role), Default: sp.Key == in.char.RaidSpec}
	}
	embed := notify.Embed{
		Title: guideCopy.SpecTitle, Description: lead, Color: classColors[in.char.Class],
		Thumbnail: s.d.PublicBaseURL + "/icons/hd/crests/" + in.char.Class + ".webp",
		Footer:    cardFooter, Timestamp: s.now(),
	}
	sel := Component{Kind: KindStringSelect, CustomID: id, Placeholder: guideCopy.SpecPlaceholder, Options: options}
	msg := Message{Embeds: []notify.Embed{embed}, Components: []ActionRow{{sel}}, Ephemeral: true}
	return Reply{Message: msg, Update: in.update}, nil
}

// raidPairs are the lead's and the sentences' markers for the card's raid and character.
func (s *Service) raidPairs(in specCardIn) []string {
	return []string{"{Name}", memberText(in.char.Name), "{Title}", memberText(in.raid.Title),
		"{Date}", in.raid.StartsAt.In(raidLocation(in.zone)).Format(raidChoiceDate)}
}

// specPress answers the presses on a spec card. A card of an older shape answers spec.stale.
func (s *Service) specPress() Handler {
	parts := map[string]cardPart{
		partSpecPick: s.specPicked, partSpecChange: s.specChangeOpened, partSpecSet: s.specChanged,
	}
	return func(ctx context.Context, in Interaction) (Reply, error) {
		p, ok := parseCardID(in.CustomID())
		if ok && p.Version != cardVersion {
			return Reply{Update: true, Message: Message{Content: guideCopy.SpecStale, Ephemeral: true}}, nil
		}
		part := parts[p.Part]
		if !ok || part == nil {
			return pressReply(pressCopy.Failure, true), nil
		}
		return part(ctx, in, p)
	}
}

// specPressRef is a spec card id's raid, state ("" on Change spec) and character key; ok is
// false for any other shape.
func specPressRef(p cardPress) (raidID int64, state, key string, ok bool) {
	fields := p.State
	if p.Part == partSpecPick {
		if len(fields) != 3 || !knownSignUpState(fields[1]) || fields[1] == guilds.SignupOut {
			return 0, "", "", false
		}
		state, fields = fields[1], []string{fields[0], fields[2]}
	}
	if len(fields) != 2 || fields[1] == "" {
		return 0, "", "", false
	}
	raidID, err := strconv.ParseInt(fields[0], 10, 64)
	if err != nil || raidID <= 0 {
		return 0, "", "", false
	}
	return raidID, state, fields[1], true
}

// specSeat is a spec press re-checked: the presser, the raid still open, and the presser's own
// verified character. refused is the answer in place when any check fails.
type specSeat struct {
	ask  signupAsk
	in   specCardIn
	char guilds.SignupCandidate
}

func (s *Service) specSeatOf(ctx context.Context, in Interaction, p cardPress) (specSeat, Reply, bool, error) {
	raidID, state, key, ok := specPressRef(p)
	if !ok {
		return specSeat{}, pressReply(pressCopy.Failure, true), false, nil
	}
	a, refused, ok := s.resolveAsker(ctx, in, true)
	if !ok {
		return specSeat{}, refused, false, nil
	}
	a.RaidID, a.State, a.CharacterKey = raidID, state, key
	raid, zone, found, err := s.specRaid(ctx, a.GuildID, raidID)
	if err != nil || !found {
		return specSeat{}, pressReply(pressCopy.Failure, true), false, err
	}
	if refusal := s.raidRefusal(raid); refusal != "" {
		return specSeat{}, a.reply(refusal), false, nil
	}
	cands, err := s.freshCandidates(ctx, a.GuildID, a.UserID)
	if err != nil {
		return specSeat{}, Reply{}, false, err
	}
	char, known := candidateByKey(cands, key)
	switch {
	case !known:
		return specSeat{}, a.reply(pressCopy.Failure), false, nil
	case !char.Verified:
		return specSeat{}, a.reply(fill(pressCopy.WaitingOne, "{Name}", memberText(char.Name))), false, nil
	}
	return specSeat{ask: a, char: char, in: specCardIn{raid: raid, zone: zone, char: char, state: state, update: true}}, Reply{}, true, nil
}

// candidateByKey is the account's character with exactly this key, verified or not.
func candidateByKey(cands []guilds.SignupCandidate, key string) (guilds.SignupCandidate, bool) {
	for _, c := range cands {
		if c.CharacterKey == key {
			return c, true
		}
	}
	return guilds.SignupCandidate{}, false
}

// classSpec is the picked value when it is one of the character's class's specs.
func classSpec(c guilds.SignupCandidate, value string) (raidspec.Spec, bool) {
	for _, sp := range raidspec.ForClass(c.Class) {
		if sp.Key == value {
			return sp, true
		}
	}
	return raidspec.Spec{}, false
}

// specPicked is the ask's pick: the sign-up with the spec's role, then the spec saved (so a
// refused sign-up saves nothing), answered by the outcome plus spec.kept, in place.
func (s *Service) specPicked(ctx context.Context, in Interaction, p cardPress) (Reply, error) {
	seat, refused, ok, err := s.specSeatOf(ctx, in, p)
	if err != nil || !ok {
		return refused, err
	}
	sp, valid := classSpec(seat.char, picked(in))
	if !valid {
		return seat.ask.reply(pressCopy.Failure), nil
	}
	rep, signed, err := s.commitSignup(ctx, seat.ask, seat.char, sp.Role)
	if err != nil || !signed {
		return rep, err // a refusal sentence: nothing was saved
	}
	if err := s.saveSpec(ctx, seat, sp); err != nil {
		return Reply{}, err
	}
	rep.Message.Content += fill(guideCopy.SpecKept, "{Spec}", sp.Name, "{Name}", memberText(seat.char.Name))
	return rep, nil
}

func (s *Service) saveSpec(ctx context.Context, seat specSeat, sp raidspec.Spec) error {
	_, err := s.d.Guilds.SaveRaidSpec(ctx, guilds.RaidSpecCommand{
		GuildID: seat.ask.GuildID, ActorID: seat.ask.UserID, RaidID: seat.ask.RaidID,
		CharacterKey: seat.char.CharacterKey, Spec: sp.Key,
	})
	return s.specSaveErr(err)
}

// specSaveErr is nil, or the error the card failure route answers with the card kept.
func (s *Service) specSaveErr(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("discordbot: save raid spec: %w", err)
}

// specChangeOpened is Change spec: the card in place, the saved spec chosen.
func (s *Service) specChangeOpened(ctx context.Context, in Interaction, p cardPress) (Reply, error) {
	seat, refused, ok, err := s.specSeatOf(ctx, in, p)
	if err != nil || !ok {
		return refused, err
	}
	return s.specCard(seat.in)
}

// specChanged is a pick on the Change spec card: the spec saved and the row's role moved,
// never its state. The answer is the row's Already sentence plus spec.kept, or spec.kept alone
// with no row or an Out row.
func (s *Service) specChanged(ctx context.Context, in Interaction, p cardPress) (Reply, error) {
	seat, refused, ok, err := s.specSeatOf(ctx, in, p)
	if err != nil || !ok {
		return refused, err
	}
	sp, valid := classSpec(seat.char, picked(in))
	if !valid {
		return seat.ask.reply(pressCopy.Failure), nil
	}
	res, err := s.d.Guilds.SaveRaidSpec(ctx, guilds.RaidSpecCommand{
		GuildID: seat.ask.GuildID, ActorID: seat.ask.UserID, RaidID: seat.ask.RaidID,
		CharacterKey: seat.char.CharacterKey, Spec: sp.Key,
	})
	if errors.Is(err, guilds.ErrCharacterNotVerified) {
		return seat.ask.reply(fill(pressCopy.WaitingOne, "{Name}", memberText(seat.char.Name))), nil
	}
	if err != nil {
		return Reply{}, s.specSaveErr(err)
	}
	kept := fill(guideCopy.SpecKept, "{Spec}", sp.Name, "{Name}", memberText(seat.char.Name))
	if res.State == "" || res.State == guilds.SignupOut {
		return seat.ask.reply(strings.TrimPrefix(kept, " ")), nil
	}
	pairs := append(s.raidPairs(seat.in), "{Role}", roleWord(sp.Role))
	return seat.ask.reply(alreadySentence(res.State, pairs)+kept,
		changeSpecButtons(res.State, seat.ask.RaidID, res.Character)...), nil
}
