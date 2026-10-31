package discordbot

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

// Parts of a /signup card besides the four state words.
const (
	partSignupRaid = "raid"
	partSignupChar = "char"
)

// signupGuildTimeLayout is the Guild time in a raid option's description, as the post shows it.
const signupGuildTimeLayout = "3:04 pm MST"

// signupPicks is what a /signup card has chosen: the raid and the presser's own character.
type signupPicks struct {
	RaidID       int64
	CharacterKey string
}

// signupPicksOf reads the card state `{raid}:{key}`; anything unreadable reads as empty, and
// every use re-checks the picks against the presser's own raids and characters.
func signupPicksOf(state []string) signupPicks {
	if len(state) != 2 {
		return signupPicks{}
	}
	id, err := strconv.ParseInt(state[0], 10, 64)
	if err != nil || id <= 0 {
		return signupPicks{}
	}
	return signupPicks{RaidID: id, CharacterKey: state[1]}
}

func signupCardID(part string, p signupPicks) string {
	return cardCustomID(prefixSignup, part, strconv.FormatInt(p.RaidID, 10), p.CharacterKey)
}

// signupData is what a card is drawn from, read fresh on every draw and press.
type signupData struct {
	raids    []guilds.Raid
	verified []guilds.SignupCandidate
}

// signupGate is the member gate and the planned raids; refused is the answer when the card
// cannot be drawn (not a member, waiting for approval, or no raid).
func (s *Service) signupGate(ctx context.Context, a signupAsk) (data signupData, refused Reply, ok bool, err error) {
	cands, err := s.freshCandidates(ctx, a.GuildID, a.UserID)
	if err != nil {
		return signupData{}, Reply{}, false, err
	}
	verified := verifiedOf(cands)
	if len(verified) == 0 {
		refused, err = s.notSignable(ctx, a, cands)
		return signupData{}, refused, false, err
	}
	raids, err := s.signableRaids(ctx, a.GuildID)
	if err != nil {
		return signupData{}, Reply{}, false, err
	}
	if len(raids) == 0 {
		return signupData{}, a.reply(pressCopy.NoRaid), false, nil
	}
	return signupData{raids: raids, verified: verified}, Reply{}, true, nil
}

// signupStart is a /signup with no status: the card for the typed raid and character, else the
// next raid and the main.
func (s *Service) signupStart(ctx context.Context, in Interaction, a signupAsk) (Reply, error) {
	data, refused, ok, err := s.signupGate(ctx, a)
	if err != nil || !ok {
		return refused, err
	}
	p := signupPicks{RaidID: data.raids[0].ID, CharacterKey: mainOf(data.verified).CharacterKey}
	if ref, _ := in.Option(optRaid); ref != "" {
		id, rep, found := s.typedRaid(ctx, a, data.raids, ref)
		if !found {
			return rep, nil
		}
		p.RaidID = id
	}
	if ref, _ := in.Option(optCharacter); ref != "" {
		c, found := chosen(data.verified, ref)
		if !found {
			return a.reply(pressCopy.Failure), nil
		}
		p.CharacterKey = c.CharacterKey
	}
	return s.signupCard(ctx, a, data.raids, data.verified, p, false)
}

// typedRaid is the raid a typed option names among the signable ones; otherwise the answer is
// the raid's refusal sentence (cancelled, ended, locked) or the failure sentence.
func (s *Service) typedRaid(ctx context.Context, a signupAsk, raids []guilds.Raid, ref string) (int64, Reply, bool) {
	id, err := strconv.ParseInt(ref, 10, 64)
	if err != nil || id <= 0 {
		return 0, a.reply(pressCopy.Failure), false
	}
	if _, listed := raidOf(raids, id); listed {
		return id, Reply{}, true
	}
	raid, found, err := s.d.Guilds.GetRaid(ctx, a.GuildID, id)
	if err == nil && found {
		if refusal := s.raidRefusal(raid); refusal != "" {
			return 0, a.reply(refusal), false
		}
	}
	return 0, a.reply(pressCopy.Failure), false
}

// mainOf is the main among the verified characters, else the first (the service puts the
// main first).
func mainOf(verified []guilds.SignupCandidate) guilds.SignupCandidate {
	for _, c := range verified {
		if c.Main {
			return c
		}
	}
	return verified[0]
}

func raidOf(raids []guilds.Raid, id int64) (guilds.Raid, bool) {
	for _, r := range raids {
		if r.ID == id {
			return r, true
		}
	}
	return guilds.Raid{}, false
}

// signupCard draws the card for picks p. Picks that are not the presser's own signable raid
// and verified character answer the failure sentence, so a forged id changes nothing.
func (s *Service) signupCard(ctx context.Context, a signupAsk, raids []guilds.Raid, verified []guilds.SignupCandidate, p signupPicks, update bool) (Reply, error) {
	raid, okRaid := raidOf(raids, p.RaidID)
	char, okChar := verifiedKey(verified, p.CharacterKey)
	if !okRaid || !okChar {
		return pressReply(pressCopy.Failure, update), nil
	}
	views, err := s.signupViews(ctx, a, raids, raid)
	if err != nil {
		return Reply{}, err
	}
	facts, found, err := s.d.Guilds.RaidPostFacts(ctx, raid.ID)
	if err != nil {
		return Reply{}, fmt.Errorf("discordbot: signup card facts: %w", err)
	}
	if !found {
		return pressReply(pressCopy.Failure, update), nil
	}
	movedFrom, err := s.d.Store.MovedFrom(ctx, raid.ID)
	if err != nil {
		return Reply{}, fmt.Errorf("discordbot: signup card moved: %w", err)
	}
	embed := s.raidAnswerEmbed(ctx, facts, movedFrom, ownOf(views[raid.ID], char.CharacterKey), verified)
	rows := signupRows(raids, views, verified, p, facts.RaidURL)
	return Reply{Message: Message{Embeds: []notify.Embed{embed}, Components: rows, Ephemeral: true}, Update: update}, nil
}

// verifiedKey is the verified character with exactly this key.
func verifiedKey(verified []guilds.SignupCandidate, key string) (guilds.SignupCandidate, bool) {
	for _, c := range verified {
		if c.CharacterKey == key {
			return c, true
		}
	}
	return guilds.SignupCandidate{}, false
}

// signupViews are the presser's member views of the raids a card shows: the chosen one, and
// every listed one when the raid select needs their counts.
func (s *Service) signupViews(ctx context.Context, a signupAsk, raids []guilds.Raid, picked guilds.Raid) (map[int64]guilds.RaidView, error) {
	listed := []guilds.Raid{picked}
	if len(raids) > 1 {
		listed = raids[:min(len(raids), maxSelectOptions)]
	}
	views := make(map[int64]guilds.RaidView, len(listed))
	for _, r := range listed {
		v, found, err := s.d.Guilds.MemberRaidView(ctx, a.GuildID, a.UserID, r.ID)
		if err != nil {
			return nil, fmt.Errorf("discordbot: signup card view: %w", err)
		}
		if found {
			views[r.ID] = v
		}
	}
	return views, nil
}

// ownOf is the character's own state in the raid view; nil when it has not signed.
func ownOf(v guilds.RaidView, key string) *guilds.NextRaidOwn {
	for _, c := range v.Viewer.Characters {
		if c.CharacterKey == key && c.State != nil {
			return &guilds.NextRaidOwn{CharacterKey: key, State: *c.State, Role: c.Role}
		}
	}
	return nil
}

func signupRows(raids []guilds.Raid, views map[int64]guilds.RaidView, verified []guilds.SignupCandidate, p signupPicks, raidURL string) []ActionRow {
	var rows []ActionRow
	if len(raids) > 1 {
		rows = append(rows, ActionRow{raidSelect(raids, views, p)})
	}
	if len(verified) > 1 {
		rows = append(rows, ActionRow{signupCharSelect(verified, p)})
	}
	buttons := make(ActionRow, 0, len(signUpStates)+1)
	for _, st := range signUpStates {
		buttons = append(buttons, answerButton(st.Label, signupCardID(st.State, p)))
	}
	return append(rows, rowOf(append(buttons, linkButton(postCopy.RaidPage, raidURL))...)...)
}

func raidSelect(raids []guilds.Raid, views map[int64]guilds.RaidView, p signupPicks) Component {
	options := make([]SelectOption, 0, min(len(raids), maxSelectOptions))
	for _, r := range raids[:min(len(raids), maxSelectOptions)] {
		options = append(options, raidOption(r, views[r.ID], r.ID == p.RaidID))
	}
	return Component{Kind: KindStringSelect, CustomID: signupCardID(partSignupRaid, p),
		Placeholder: guideCopy.SignupPlaceholderRaid, Options: options}
}

// raidOption is `{Title} · {Ddd d Mmm}` described by the guild time and the count. The label
// is plain text in a select, so the title is collapsed, not escaped.
func raidOption(r guilds.Raid, v guilds.RaidView, selected bool) SelectOption {
	zone := ""
	if v.GuildTimezone != nil {
		zone = *v.GuildTimezone
	}
	date := r.StartsAt.In(raidLocation(zone)).Format(raidChoiceDate)
	room := maxOptionRunes - len([]rune(fill(guideCopy.SignupRaidLabel, "{Title}", "", "{Ddd d Mmm}", date)))
	title := truncateRunes(notify.Collapse(r.Title), room)
	return SelectOption{
		Label: fill(guideCopy.SignupRaidLabel, "{Title}", title, "{Ddd d Mmm}", date), Value: strconv.FormatInt(r.ID, 10),
		Default: selected, Description: truncateRunes(raidOptionDescription(r.StartsAt, zone, v), maxOptionRunes),
	}
}

// raidOptionDescription is the guild time and the count; without a guild zone only the count.
func raidOptionDescription(starts time.Time, zone string, v guilds.RaidView) string {
	signed := strconv.Itoa(v.Counts.In + v.Counts.Late)
	if v.Size == nil {
		return fill(guideCopy.SignupRaidDescNoSize, "{n}", signed)
	}
	count := fill(guideCopy.SignupRaidDesc, "{signed}", signed, "{size}", strconv.Itoa(*v.Size))
	if zone == "" {
		_, count, _ = strings.Cut(count, lineJoin)
		return count
	}
	return strings.Replace(count, "{Guild time}", starts.In(raidLocation(zone)).Format(signupGuildTimeLayout), 1)
}

func signupCharSelect(verified []guilds.SignupCandidate, p signupPicks) Component {
	options := make([]SelectOption, 0, min(len(verified), maxSelectOptions))
	for _, c := range verified[:min(len(verified), maxSelectOptions)] {
		o := characterOption(c)
		o.Default = c.CharacterKey == p.CharacterKey
		options = append(options, o)
	}
	return Component{Kind: KindStringSelect, CustomID: signupCardID(partSignupChar, p),
		Placeholder: guideCopy.SignupPlaceholderCharacter, Options: options}
}

// signupPress answers the presses on a /signup card. Every press reads the presser's raids and
// verified characters again, so the key in an id is only ever a request.
func (s *Service) signupPress() Handler {
	parts := map[string]cardPart{
		partSignupRaid: s.signupPick(func(p signupPicks, v string) signupPicks {
			id, _ := strconv.ParseInt(v, 10, 64)
			return signupPicks{RaidID: id, CharacterKey: p.CharacterKey}
		}),
		partSignupChar: s.signupPick(func(p signupPicks, v string) signupPicks {
			return signupPicks{RaidID: p.RaidID, CharacterKey: v}
		}),
	}
	for _, st := range signUpStates {
		parts[st.State] = s.signupCommit(st.State)
	}
	return s.cardHandler("/signup", parts)
}

// signupPick re-renders the card with one pick changed.
func (s *Service) signupPick(apply func(signupPicks, string) signupPicks) cardPart {
	return func(ctx context.Context, in Interaction, press cardPress) (Reply, error) {
		a, refused, ok := s.resolveAsker(ctx, in, true)
		if !ok {
			return refused, nil
		}
		data, refused, ok, err := s.signupGate(ctx, a)
		if err != nil || !ok {
			return refused, err
		}
		return s.signupCard(ctx, a, data.raids, data.verified, apply(signupPicksOf(press.State), picked(in)), true)
	}
}

// signupCommit is a status press: the press flow with the raid and character of the id, which
// the flow checks against the presser's verified characters.
func (s *Service) signupCommit(state string) cardPart {
	return func(ctx context.Context, in Interaction, press cardPress) (Reply, error) {
		a, refused, ok := s.resolveAsker(ctx, in, true)
		if !ok {
			return refused, nil
		}
		p := signupPicksOf(press.State)
		if p.RaidID == 0 || p.CharacterKey == "" {
			return a.reply(pressCopy.Failure), nil
		}
		a.RaidID, a.State, a.CharacterKey = p.RaidID, state, p.CharacterKey
		return s.runFlow(ctx, in, a), nil
	}
}
