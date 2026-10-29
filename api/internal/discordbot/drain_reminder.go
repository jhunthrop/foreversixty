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

// reminds reports whether a raid in state still takes its reminder: an open one is edited
// into Tomorrow; a moved one keeps Moved and is reminded by the reply.
func reminds(state notify.RaidPostState) bool {
	return state == notify.PostOpen || state == notify.PostMoved
}

// remind sends the due reminder with the post's own send (§P.5). The post is drawn as
// Tomorrow (Moved stays Moved). When the channel moved on since the post, or the post cannot
// show Tomorrow, a reply under it carries the reminder line and the post's sign-up row. A
// post created by this very drain is the newest message, so it takes no reply. It answers
// the post's message id and look and records the reminder; the caller finishes the edit.
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
	d.write(ctx, "reminder_finish", p.RaidID, func(ctx context.Context) error {
		return d.Store.FinishReminder(ctx, p, replyID, now)
	})
	return messageID, look, nil
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
// raid already read Moved stays. Discord's 10008 counts as deleted.
func (d *Drainer) withdraw(ctx context.Context, p Post, state notify.RaidPostState) error {
	if p.ReminderMessageID == nil || !withdrawn(state, p) {
		return nil
	}
	if err := d.REST.DeleteMessage(ctx, p.ChannelID, *p.ReminderMessageID); err != nil {
		return err
	}
	d.logger().Info("discordbot", "op", "reminder_withdrawn", "raid_id", p.RaidID, "state", state)
	d.write(ctx, "reminder_clear", p.RaidID, func(ctx context.Context) error {
		return d.Store.ClearReminder(ctx, p, state == notify.PostMoved)
	})
	return nil
}

// withdrawn reports whether a reminder reply up for p no longer stands in state.
func withdrawn(state notify.RaidPostState, p Post) bool {
	switch state {
	case notify.PostCancelled, notify.PostLocked:
		return true
	}
	return p.RemindedAt == nil
}

// sendPost brings the post up to date in its look (lookOf) and answers that look: a due
// reminder rides with it, and a reminder that no longer stands is withdrawn. A reminder due
// on a raid that takes none (locked, cancelled, ended, done) is dropped unsent.
func (d *Drainer) sendPost(ctx context.Context, p Post, facts guilds.RaidPostFacts, state notify.RaidPostState, now time.Time) (string, notify.RaidPostState, error) {
	if p.RemindDue && reminds(state) {
		return d.remind(ctx, p, facts, state, now)
	}
	look := lookOf(state, p, facts.Post.StartsAt, now)
	messageID, err := d.send(ctx, p, postFor(facts, look, p.MovedFrom), facts, now)
	if err == nil {
		err = d.withdraw(ctx, p, state)
	}
	if err == nil && p.RemindDue {
		d.write(ctx, "reminder_drop", p.RaidID, func(ctx context.Context) error {
			return d.Store.FinishReminder(ctx, p, nil, now)
		})
	}
	return messageID, look, err
}
