package discordbot

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

// shareIDPrefix is the Share in channel button's custom id prefix: share:{token}.
const shareIDPrefix = "share"

// shareLife is how long an answer's Share button works.
const shareLife = time.Hour

// followUpRetryAfter is the pause before the one retry of a follow-up Discord refused with a
// 404: the acknowledgement it needs may not have landed yet.
const followUpRetryAfter = time.Second

// lookupAnswer computes a lookup's answer for viewer (nil: a signed-out visitor) from its
// resolved options. found is false when there is nothing to show that viewer. The message
// carries its link row; the Share button is the caller's.
type lookupAnswer func(ctx context.Context, guildID int64, viewer *int64, opts map[string]string) (Message, bool, error)

// lookups are the answers Share in channel can recompute, by command (spec §3.3.4).
func (s *Service) lookups() map[string]lookupAnswer {
	return map[string]lookupAnswer{"char": s.answerChar, "bis": s.answerBis, "report": s.answerReport}
}

// visitorAnswer is command's answer as a signed-out visitor sees it; nil for a command that
// has no Share.
func (s *Service) visitorAnswer(command string) func(ctx context.Context, guildID int64, opts map[string]string) (Message, bool, error) {
	answer, ok := s.lookups()[command]
	if !ok {
		return nil
	}
	return func(ctx context.Context, guildID int64, opts map[string]string) (Message, bool, error) {
		return answer(ctx, guildID, nil, opts)
	}
}

// storedShare is a discord_shares row's options column: the answer's resolved options and the
// Discord user it was answered to, the only one whose press shares it.
type storedShare struct {
	By      string            `json:"by"`
	Options map[string]string `json:"options"`
}

// offerShare stores the answer's resolved options under a fresh token for an hour, bound to
// the Discord user it answered, and returns its Share in channel button. The token names no
// account.
func (s *Service) offerShare(ctx context.Context, in Interaction, guildID int64, command string, opts map[string]string) (Component, error) {
	token := auth.NewSessionID()
	raw, err := json.Marshal(storedShare{By: in.DiscordUserID, Options: opts})
	if err != nil {
		return Component{}, fmt.Errorf("discordbot: share options: %w", err)
	}
	if _, err := s.d.Store.Pool.Exec(ctx, `insert into discord_shares (token, guild_id, command, options, expires_at)
		values ($1, $2, $3, $4, $5)`, token, guildID, command, raw, s.now().Add(shareLife)); err != nil {
		return Component{}, fmt.Errorf("discordbot: save share: %w", err)
	}
	return shareButton(token, false), nil
}

func shareButton(token string, shared bool) Component {
	label := lookupCopy.Share
	if shared {
		label = lookupCopy.Shared
	}
	return Component{Kind: KindButton, Style: StyleSecondary, Label: label, CustomID: shareIDPrefix + ":" + token, Disabled: shared}
}

// withShare is msg with button first on its first row.
func withShare(msg Message, button Component) Message {
	rows := append([]ActionRow(nil), msg.Components...)
	if len(rows) == 0 {
		rows = []ActionRow{{}}
	}
	rows[0] = append(ActionRow{button}, rows[0]...)
	msg.Components = rows
	return msg
}

type shareRow struct {
	command string
	by      string
	opts    map[string]string
	used    bool
}

// share answers a press of Share in channel: the answer recomputed as a visitor goes to the
// channel once, and the ephemeral's button turns into a disabled Shared.
func (s *Service) share(ctx context.Context, in Interaction) (Reply, error) {
	token := strings.TrimPrefix(in.CustomID(), shareIDPrefix+":")
	guildID, bound, err := s.serverGuild(ctx, in)
	if err != nil || !bound {
		return failureReply(), err
	}
	row, found, err := s.shareRow(ctx, token, guildID)
	if err != nil {
		return Reply{}, err
	}
	visitor := s.visitorAnswer(row.command)
	if !found || visitor == nil || row.by != in.DiscordUserID {
		return failureReply(), nil
	}
	row.opts = withSourcePage(in, row)
	if row.used {
		return s.sharedReply(ctx, in, guildID, token, row, Message{})
	}
	public, post, ok, err := s.publicAnswer(ctx, row.command, guildID, row.opts)
	if err != nil {
		return Reply{}, fmt.Errorf("discordbot: share recompute: %w", err)
	}
	if !ok {
		return pressReply(lookupCopy.NoLongerPublic, false), nil
	}
	claimed, err := s.claimShare(ctx, token)
	if err != nil || !claimed {
		if err != nil {
			return Reply{}, err
		}
		return s.sharedReply(ctx, in, guildID, token, row, public)
	}
	rep, err := s.sharedReply(ctx, in, guildID, token, row, public)
	if err != nil {
		s.releaseShare(ctx, token)
		return Reply{}, err
	}
	public.Ephemeral = false
	restored := withLiveShare(rep.Message, token)
	rep.After = func(ctx context.Context) {
		if answerDropped(ctx) { // the Shared answer never reached Discord: free the token
			s.releaseShare(ctx, token)
			return
		}
		s.postShare(ctx, in, token, s.publicShare(post, public, in), restored, row.command)
	}
	return rep, nil
}

// withSourcePage keeps a /bis answer on the page it showed when Share was pressed: the
// presser's own copy is redrawn at that page. Share stores no page; the sheet shows all 17.
func withSourcePage(in Interaction, row shareRow) map[string]string {
	if row.command != "bis" {
		return row.opts
	}
	if page := sourcePage(in.Source.Components); page != "" {
		return withOption(row.opts, optPage, page)
	}
	return row.opts
}

// sharedReply re-renders the presser's own answer with the disabled Shared button; when it
// no longer renders for them, the visitor's answer stands in.
func (s *Service) sharedReply(ctx context.Context, in Interaction, guildID int64, token string, row shareRow, fallback Message) (Reply, error) {
	viewer, linked, err := s.viewer(ctx, in)
	if err != nil {
		return Reply{}, fmt.Errorf("discordbot: share viewer: %w", err)
	}
	var who *int64
	if linked {
		who = &viewer
	}
	msg, ok, err := s.lookups()[row.command](ctx, guildID, who, row.opts)
	if err != nil {
		return Reply{}, fmt.Errorf("discordbot: share render: %w", err)
	}
	if ok {
		msg = decorateOwn(row.command, row.opts, msg)
		if msg, err = s.withGuidedControls(ctx, row.command, row.opts, guildID, who, msg); err != nil {
			return Reply{}, fmt.Errorf("discordbot: share controls: %w", err)
		}
	} else {
		msg = fallback
	}
	msg = withShare(msg, shareButton(token, true))
	msg.Ephemeral = true
	return Reply{Message: msg, Update: true}, nil
}

// withLiveShare is a Shared answer with its button live again, saying the share failed: as
// its content line, or on a V2 answer as a Text Display above the panel (spec §A.7). The
// restore edits the Shared answer, which the bot already drew in the answer's encoding.
func withLiveShare(shared Message, token string) Message {
	rows := append([]ActionRow(nil), shared.Components...)
	rows[0] = append(ActionRow{shareButton(token, false)}, rows[0][1:]...)
	shared.Components = rows
	if shared.componentsV2() {
		shared.Layout = append([]Node{textNode(commonCopy.Failure)}, shared.Layout...)
		shared.sourceV2 = true
		return shared
	}
	shared.Content = commonCopy.Failure
	return shared
}

// postShare sends the public follow-up. A failure frees the token and puts the live button
// back on the ephemeral, so another press can try again.
func (s *Service) postShare(ctx context.Context, in Interaction, token string, public, restored Message, command string) {
	err := s.d.REST.FollowUp(ctx, in.Token, public)
	var api *APIError
	if errors.As(err, &api) && api.Status == http.StatusNotFound {
		select {
		case <-time.After(followUpRetryAfter):
			err = s.d.REST.FollowUp(ctx, in.Token, public)
		case <-ctx.Done():
		}
	}
	if err == nil {
		return
	}
	s.log.Warn("discordbot", "op", "share_follow_up", "interaction", in.ID, "command", command, "err", describe(err))
	s.releaseShare(ctx, token)
	if editErr := s.d.REST.EditOriginal(ctx, in.Token, restored); editErr != nil {
		s.log.Warn("discordbot", "op", "share_follow_up", "outcome", "edit_failed", "interaction", in.ID, "err", describe(editErr))
	}
}

func (s *Service) shareRow(ctx context.Context, token string, guildID int64) (shareRow, bool, error) {
	var (
		row shareRow
		raw []byte
	)
	err := s.d.Store.Pool.QueryRow(ctx, `select command, options, used_at is not null from discord_shares
		where token = $1 and guild_id = $2 and expires_at > $3`, token, guildID, s.now()).Scan(&row.command, &raw, &row.used)
	if errors.Is(err, pgx.ErrNoRows) {
		return shareRow{}, false, nil
	}
	if err != nil {
		return shareRow{}, false, fmt.Errorf("discordbot: read share: %w", err)
	}
	var stored storedShare
	if err := json.Unmarshal(raw, &stored); err != nil {
		return shareRow{}, false, fmt.Errorf("discordbot: share options: %w", err)
	}
	row.by, row.opts = stored.By, stored.Options
	return row, true, nil
}

// claimShare marks the token used; false when another press already did.
func (s *Service) claimShare(ctx context.Context, token string) (bool, error) {
	tag, err := s.d.Store.Pool.Exec(ctx, `update discord_shares set used_at = $2 where token = $1 and used_at is null`, token, s.now())
	if err != nil {
		return false, fmt.Errorf("discordbot: claim share: %w", err)
	}
	return tag.RowsAffected() == 1, nil
}

// shareKeep is how long past its expiry a Share token's row is kept before the sweep.
const shareKeep = 24 * time.Hour

// SweepShares deletes Share tokens a day past their expiry. The post drainer's loop runs it
// hourly (Drainer.Run), the bot's one periodic loop.
func (s *Store) SweepShares(ctx context.Context, now time.Time) error {
	return s.exec(ctx, `delete from discord_shares where expires_at < $1`, now.Add(-shareKeep))
}

func (s *Service) releaseShare(ctx context.Context, token string) {
	if _, err := s.d.Store.Pool.Exec(ctx, `update discord_shares set used_at = null where token = $1`, token); err != nil {
		s.log.Error("discordbot", "op", "release_share", "err", err)
	}
}
