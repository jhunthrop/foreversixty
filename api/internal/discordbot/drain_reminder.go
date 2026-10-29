package discordbot

import (
	"context"
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

// lookOf is how the post is drawn in state: an open raid reminded before its start shows
// Tomorrow until the start, then its open look again. Any other state outranks Tomorrow
// (Ruling 4: a moved post reads Moved).
func lookOf(state notify.RaidPostState, p Post, startsAt, now time.Time) notify.RaidPostState {
	if state == notify.PostOpen && p.RemindedAt != nil && now.Before(startsAt) {
		return notify.PostTomorrow
	}
	return state
}

// reminds reports whether a raid in state still takes its reminder before its start: an
// open one is edited into Tomorrow; a moved one keeps Moved and is reminded by the reply.
func reminds(state notify.RaidPostState, startsAt, now time.Time) bool {
	return (state == notify.PostOpen || state == notify.PostMoved) && now.Before(startsAt)
}

// sendPost withdraws a reminder that no longer stands, then brings the post up to date in its
// look (lookOf) and answers that look. A due reminder rides with the send; one due on a raid
// that takes none (locked, cancelled, ended, done, or past its start) is dropped unsent.
func (d *Drainer) sendPost(ctx context.Context, p Post, facts guilds.RaidPostFacts, state notify.RaidPostState, now time.Time) (string, notify.RaidPostState, error) {
	p, err := d.withdraw(ctx, p, state)
	if err != nil {
		return "", state, err
	}
	if p.RemindDue && reminds(state, facts.Post.StartsAt, now) {
		return d.remind(ctx, p, facts, state, now)
	}
	look := lookOf(state, p, facts.Post.StartsAt, now)
	messageID, err := d.send(ctx, p, postFor(facts, look, p.MovedFrom), facts, now)
	if err == nil && p.RemindDue {
		err = d.record(ctx, func(ctx context.Context) error { return d.Store.FinishReminder(ctx, p, nil, now) })
	}
	return messageID, look, err
}

// remind sends the due reminder with the post's own send (§P.5). The post is drawn as
// Tomorrow (Moved stays Moved). When the channel moved on since the post, or the post cannot
// show Tomorrow, a reply under it carries the reminder line and the post's sign-up row. A
// post created by this very drain is the newest message, so it takes no reply. It answers
// the post's message id and look. The reminder is recorded before the caller finishes the
// edit; a failed record is returned, so the edit is not finished as if it had been sent.
func (d *Drainer) remind(ctx context.Context, p Post, facts guilds.RaidPostFacts, state notify.RaidPostState, now time.Time) (string, notify.RaidPostState, error) {
	look := notify.PostMoved
	if state == notify.PostOpen {
		look = notify.PostTomorrow
	}
	newest := ""
	if p.MessageID != nil && look == notify.PostTomorrow {
		latest, err := d.REST.LatestMessageID(ctx, p.ChannelID)
		if err != nil {
			return "", look, err
		}
		newest = latest
	}
	messageID, err := d.send(ctx, p, postFor(facts, look, p.MovedFrom), facts, now)
	if err != nil {
		return "", look, err
	}
	var replyID *string
	if p.MessageID != nil && (look != notify.PostTomorrow || newest != messageID) {
		id, err := d.REST.CreateMessage(ctx, p.ChannelID, reminderReply(p, facts, messageID))
		if err != nil {
			return "", look, err
		}
		replyID = &id
	}
	d.logger().Info("discordbot", "op", "post_reminded", "raid_id", p.RaidID, "reply", replyID != nil)
	err = d.record(ctx, func(ctx context.Context) error { return d.Store.FinishReminder(ctx, p, replyID, now) })
	return messageID, look, err
}

// reminderReply is the reminder as a reply to the post: the line, and the post's own open
// row, so a press there is a press on the post. Its nonce keeps a retried reply single.
func reminderReply(p Post, facts guilds.RaidPostFacts, postID string) Message {
	open := postFor(facts, notify.PostOpen, nil)
	return Message{
		Content:    notify.ReminderLine(facts.Post.RaidPosted),
		Components: postRows(open, facts.RaidURL, facts.NightURL),
		ReplyTo:    postID,
		Nonce:      reminderNonce(p, facts.Post.StartsAt),
	}
}

// reminderNonce names one reminder: q{raid}.{start, unix s}, base 36, inside Discord's 25
// characters for any id. A raid moved and reminded again has a new start, so a new nonce.
func reminderNonce(p Post, startsAt time.Time) string {
	return "q" + strconv.FormatInt(p.RaidID, 36) + "." + strconv.FormatInt(startsAt.Unix(), 36)
}

// withdraw deletes a reminder reply that no longer stands: the raid was cancelled or locked,
// or moved after it went out (the move cleared reminded_at, Ruling 4). A reply sent while the
// raid already read Moved stays. Discord's 10008 counts as deleted. It answers the post as
// the store now holds it, so a reminder due in the same claim starts from the cleared reply.
func (d *Drainer) withdraw(ctx context.Context, p Post, state notify.RaidPostState) (Post, error) {
	if p.ReminderMessageID == nil || !withdrawn(state, p) {
		return p, nil
	}
	if err := d.REST.DeleteMessage(ctx, p.ChannelID, *p.ReminderMessageID); err != nil {
		return p, err
	}
	d.logger().Info("discordbot", "op", "reminder_withdrawn", "raid_id", p.RaidID, "state", state)
	rearm := state == notify.PostMoved
	if err := d.record(ctx, func(ctx context.Context) error { return d.Store.ClearReminder(ctx, p, rearm) }); err != nil {
		return p, err
	}
	cleared := p
	cleared.ReminderMessageID = nil
	if rearm {
		cleared.RemindedAt = nil
	}
	return cleared, nil
}

// withdrawn reports whether a reminder reply up for p no longer stands in state.
func withdrawn(state notify.RaidPostState, p Post) bool {
	switch state {
	case notify.PostCancelled, notify.PostLocked:
		return true
	}
	return p.RemindedAt == nil
}

// record runs one store write the row's finish depends on, on its own short context like
// write, and answers its error instead of only logging it: the caller then takes the row's
// error path rather than finishing the edit as if the write had landed.
func (d *Drainer) record(ctx context.Context, do func(context.Context) error) error {
	ctx, cancel := detached(ctx, writeBudget)
	defer cancel()
	return do(ctx)
}
