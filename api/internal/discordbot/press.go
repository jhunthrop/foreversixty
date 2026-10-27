package discordbot

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

// The sign-up flow's links and member-text bounds.
const (
	setupPath     = "/setup"
	signupAnchor  = "#signup"
	memberTextMax = 100
	memberTextOut = 200
)

// signupAsk is one sign-up asked from Discord: the account, its server's guild, the raid and
// the state, and the character when it is already chosen. Update answers by replacing the
// message the component sits on (the character select).
type signupAsk struct {
	GuildID, RaidID, UserID int64
	State, CharacterKey     string
	Update                  bool
}

// reply is one ephemeral sentence with at most the given link buttons.
func (a signupAsk) reply(text string, buttons ...Component) Reply {
	return pressReply(text, a.Update, buttons...)
}

func pressReply(text string, update bool, buttons ...Component) Reply {
	return Reply{Message: Message{Content: text, Ephemeral: true, Components: rowOf(buttons...)}, Update: update}
}

// fill replaces each {marker} in template with its value: fill(t, "{Name}", "Zulmara", …).
func fill(template string, pairs ...string) string {
	return strings.NewReplacer(pairs...).Replace(template)
}

// memberText is member-written text (a raid title, a character name) made to render literally.
func memberText(raw string) string {
	return notify.EscapeWithin(notify.Collapse(raw), memberTextMax, memberTextOut)
}

// stateWord is the sign-up state as the post's buttons say it.
func stateWord(state string) string {
	for _, s := range signUpStates {
		if s.State == state {
			return s.Label
		}
	}
	return state
}

func roleWord(role string) string {
	if w, ok := roleWords[role]; ok {
		return w
	}
	return roleWords[guilds.RoleDPS]
}

// press answers a raid post's In, Late, Bench or Out (custom id raid:{raid}:{state}).
func (s *Service) press(ctx context.Context, in Interaction) (Reply, error) {
	raidID, state, ok := parseRaidCustomID(in.CustomID())
	if !ok {
		return pressReply(pressCopy.Failure, false), nil
	}
	a, refused, ok := s.resolveAsker(ctx, in, false)
	if !ok {
		return refused, nil
	}
	a.RaidID, a.State = raidID, state
	return s.runFlow(ctx, in, a), nil
}

// resolveAsker is the flow's steps 1 and 2: the presser's linked account, then the server's
// guild. ok is false with the answer when either is missing.
func (s *Service) resolveAsker(ctx context.Context, in Interaction, update bool) (signupAsk, Reply, bool) {
	userID, linked, err := s.viewer(ctx, in)
	if err != nil {
		return signupAsk{}, s.saveFailed(in, update, fmt.Errorf("discordbot: signup viewer: %w", err)), false
	}
	if !linked {
		return signupAsk{}, pressReply(pressCopy.NotLinked, update, s.linkButton()), false
	}
	guildID, bound, err := s.serverGuild(ctx, in)
	if err != nil {
		return signupAsk{}, s.saveFailed(in, update, fmt.Errorf("discordbot: signup server: %w", err)), false
	}
	if !bound {
		return signupAsk{}, pressReply(commonCopy.Unbound, update), false
	}
	return signupAsk{GuildID: guildID, UserID: userID, Update: update}, Reply{}, true
}

// runFlow is signupFlow with an unexpected failure answered by the failure sentence and
// logged by ids only.
func (s *Service) runFlow(ctx context.Context, in Interaction, a signupAsk) Reply {
	rep, err := s.signupFlow(ctx, a)
	if err != nil {
		return s.saveFailed(in, a.Update, err)
	}
	return rep
}

func (s *Service) saveFailed(in Interaction, update bool, err error) Reply {
	s.log.Error("discordbot", "op", "discord_signup", "outcome", "failed", "interaction", in.ID, "err", err)
	return pressReply(pressCopy.Failure, update)
}

// signupFlow runs experience §3.3.2 from step 2 on: the raid must be the guild's; then its
// state; then the account's characters; then the site's sign-up.
func (s *Service) signupFlow(ctx context.Context, a signupAsk) (Reply, error) {
	raid, found, err := s.d.Guilds.GetRaid(ctx, a.GuildID, a.RaidID)
	if err != nil {
		return Reply{}, fmt.Errorf("discordbot: signup raid: %w", err)
	}
	if !found {
		return a.reply(pressCopy.Failure), nil
	}
	if refusal := s.raidRefusal(raid); refusal != "" {
		return a.reply(refusal), nil
	}
	cands, err := s.d.Guilds.SignupCandidates(ctx, a.GuildID, a.UserID)
	if err != nil {
		return Reply{}, fmt.Errorf("discordbot: signup candidates: %w", err)
	}
	verified := verifiedOf(cands)
	if len(verified) == 0 {
		return s.notSignable(ctx, a, cands)
	}
	if a.CharacterKey != "" {
		c, ok := chosen(verified, a.CharacterKey)
		if !ok {
			return a.reply(pressCopy.Failure), nil
		}
		return s.signUp(ctx, a, c)
	}
	if len(verified) == 1 {
		return s.signUp(ctx, a, verified[0])
	}
	return characterSelect(a, raid, verified), nil
}

// raidRefusal is the sentence of a raid nobody can sign up for from Discord: cancelled, ended
// (Ruling 11; a done raid has ended too) or locked. "" when it is open.
func (s *Service) raidRefusal(raid guilds.Raid) string {
	switch {
	case raid.Status == guilds.RaidCancelled:
		return pressCopy.Cancelled
	case raid.Status == guilds.RaidDone || raid.Ended(s.now()):
		return pressCopy.Ended
	case raid.Status == guilds.RaidLocked:
		return pressCopy.Locked
	}
	return ""
}

func verifiedOf(cands []guilds.SignupCandidate) []guilds.SignupCandidate {
	out := make([]guilds.SignupCandidate, 0, len(cands))
	for _, c := range cands {
		if c.Verified {
			out = append(out, c)
		}
	}
	return out
}

// chosen is the verified character a select value or a /signup option names: by key, or by
// its name typed in full.
func chosen(verified []guilds.SignupCandidate, ref string) (guilds.SignupCandidate, bool) {
	for _, c := range verified {
		if c.CharacterKey == ref {
			return c, true
		}
	}
	for _, c := range verified {
		if strings.EqualFold(c.Name, ref) {
			return c, true
		}
	}
	return guilds.SignupCandidate{}, false
}

// notSignable answers an account with no verified character in the guild: none at all, one
// waiting for approval, or several.
func (s *Service) notSignable(ctx context.Context, a signupAsk, cands []guilds.SignupCandidate) (Reply, error) {
	switch len(cands) {
	case 0:
		return s.notMember(ctx, a)
	case 1:
		return a.reply(fill(pressCopy.WaitingOne, "{Name}", memberText(cands[0].Name))), nil
	}
	return a.reply(pressCopy.WaitingMany), nil
}

func (s *Service) notMember(ctx context.Context, a signupAsk) (Reply, error) {
	name, err := s.d.Guilds.GuildName(ctx, a.GuildID)
	if err != nil {
		return Reply{}, fmt.Errorf("discordbot: signup guild name: %w", err)
	}
	setup := linkButton(pressCopy.SetupButton, s.d.PublicBaseURL+setupPath)
	return a.reply(fill(pressCopy.NotMember, "{Guild}", memberText(name)), setup), nil
}

// signUp calls the site's sign-up for c, with its spec's role, and words the outcome.
func (s *Service) signUp(ctx context.Context, a signupAsk, c guilds.SignupCandidate) (Reply, error) {
	res, err := s.d.Guilds.SignUp(ctx, guilds.SignupCommand{
		GuildID: a.GuildID, RaidID: a.RaidID, ActorID: a.UserID, CharacterKey: c.CharacterKey, State: a.State,
	})
	switch {
	case err == nil:
		return s.outcome(ctx, a, res), nil
	case errors.Is(err, guilds.ErrNotMember):
		return s.notMember(ctx, a)
	case errors.Is(err, guilds.ErrCharacterNotVerified):
		return a.reply(fill(pressCopy.WaitingOne, "{Name}", memberText(c.Name))), nil
	}
	if text, ok := signupRefusal(err); ok {
		return a.reply(text), nil
	}
	return Reply{}, fmt.Errorf("discordbot: signup: %w", err)
}

// signupRefusal is the sentence of a refusal the sign-up service answers.
func signupRefusal(err error) (string, bool) {
	switch {
	case errors.Is(err, guilds.ErrSignupRateLimited):
		return pressCopy.RateLimited, true
	case errors.Is(err, guilds.ErrRaidLocked):
		return pressCopy.Locked, true
	case errors.Is(err, guilds.ErrRaidCancelled):
		return pressCopy.Cancelled, true
	}
	return "", false
}

// outcome words what the sign-up now says. The post is not touched: the service's mark
// brings it up to date through the coalesced edit.
func (s *Service) outcome(ctx context.Context, a signupAsk, res guilds.SignupResult) Reply {
	name, title := memberText(res.Character.Name), memberText(res.Raid.Title)
	role := roleWord(res.Character.Role)
	if res.Unchanged {
		return a.reply(alreadySentence(a.State, name, role, title))
	}
	moved := ""
	if res.Moved != nil {
		moved = fill(pressCopy.Moved, "{Other}", memberText(*res.Moved))
	}
	switch a.State {
	case guilds.SignupIn:
		lead := fill(pressCopy.InSaved, "{Name}", name, "{Role}", role, "{Title}", title)
		if moved != "" {
			return a.reply(lead + moved)
		}
		return a.reply(lead + countClause(res.Raid))
	case guilds.SignupLate:
		text := fill(pressCopy.LateSaved, "{Name}", name, "{Role}", role, "{Title}", title) + moved
		return a.reply(text, linkButton(pressCopy.NoteButton, s.noteURL(ctx, a.RaidID)))
	case guilds.SignupBench:
		return a.reply(fill(pressCopy.BenchSaved, "{Name}", name, "{Title}", title))
	}
	return a.reply(fill(pressCopy.OutSaved, "{Name}", name, "{Title}", title))
}

func alreadySentence(state, name, role, title string) string {
	switch state {
	case guilds.SignupBench:
		return fill(pressCopy.AlreadyBench, "{Name}", name, "{Title}", title)
	case guilds.SignupOut:
		return fill(pressCopy.AlreadyOut, "{Name}", name, "{Title}", title)
	}
	return fill(pressCopy.Already, "{Name}", name, "{State}", stateWord(state), "{Role}", role, "{Title}", title)
}

// countClause is " 23 of 40 signed up." (In plus Late), or " 23 signed up." for a raid of
// unknown size.
func countClause(v guilds.RaidView) string {
	signed := strconv.Itoa(v.Counts.In + v.Counts.Late)
	if v.Size == nil {
		return fill(pressCopy.CountOpen, "{n}", signed)
	}
	return fill(pressCopy.Count, "{signed}", signed, "{size}", strconv.Itoa(*v.Size))
}

// noteURL is the raid page's sign-up section; "" (no button) when it cannot be read.
func (s *Service) noteURL(ctx context.Context, raidID int64) string {
	facts, found, err := s.d.Guilds.RaidPostFacts(ctx, raidID)
	if err != nil || !found || facts.RaidURL == "" {
		if err != nil {
			s.log.Warn("discordbot", "op", "discord_signup", "outcome", "note_link_failed", "raid_id", raidID, "err", err)
		}
		return ""
	}
	return facts.RaidURL + signupAnchor
}
