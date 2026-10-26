package notify

import (
	"context"
	"testing"
)

// closingPoster accepts the post, then takes the database away so the bookkeeping that follows
// the post fails.
type closingPoster struct {
	*FakePoster
	store *GuildStore
}

func (p closingPoster) Post(ctx context.Context, url string, embed Embed) error {
	err := p.FakePoster.Post(ctx, url, embed)
	p.store.Pool.Close()
	return err
}

func TestReportReadyPostedFlag(t *testing.T) {
	ctx := context.Background()
	t.Run("posted", func(t *testing.T) {
		r := newGuildRig(t)
		posted, err := r.d.ReportReady(ctx, r.gid, onyxiaReport())
		if err != nil || !posted || len(r.poster.Posts()) != 1 {
			t.Fatalf("posted %v, err %v, posts %d", posted, err, len(r.poster.Posts()))
		}
	})
	t.Run("no poster", func(t *testing.T) {
		var none *GuildDeliverer
		if posted, err := none.ReportReady(ctx, 1, onyxiaReport()); posted || err != nil {
			t.Fatalf("nil deliverer: posted %v, err %v", posted, err)
		}
	})
	t.Run("no webhook", func(t *testing.T) {
		r := newGuildRig(t)
		if err := r.store.Delete(ctx, r.gid); err != nil {
			t.Fatal(err)
		}
		if posted, err := r.d.ReportReady(ctx, r.gid, onyxiaReport()); posted || err != nil || len(r.poster.Posts()) != 0 {
			t.Fatalf("no channel: posted %v, err %v, posts %d", posted, err, len(r.poster.Posts()))
		}
	})
	t.Run("disabled", func(t *testing.T) {
		r := newGuildRig(t)
		r.poster.PostErr = &DiscordError{Status: 404, Body: "Unknown Webhook"}
		for i := 0; i < MaxFailures; i++ {
			r.d.RaidPosted(ctx, r.gid, onyxiaThursday())
		}
		r.poster.PostErr = nil
		before := len(r.poster.Posts())
		if posted, err := r.d.ReportReady(ctx, r.gid, onyxiaReport()); posted || err != nil || len(r.poster.Posts()) != before {
			t.Fatalf("switched off: posted %v, err %v", posted, err)
		}
	})
	t.Run("failed", func(t *testing.T) {
		r := newGuildRig(t)
		r.poster.PostErr = &DiscordError{Status: 500, Body: "boom"}
		if posted, err := r.d.ReportReady(ctx, r.gid, onyxiaReport()); posted || err != nil {
			t.Fatalf("failed post: posted %v, err %v", posted, err)
		}
	})
}

// A post that landed stays posted when saving the delivery record fails: an error would make
// the caller release its claim and post the report twice.
func TestReportReadyStaysPostedWhenTheRecordCannotBeSaved(t *testing.T) {
	r := newGuildRig(t)
	r.d.Poster = closingPoster{FakePoster: r.poster, store: r.store}
	posted, err := r.d.ReportReady(context.Background(), r.gid, onyxiaReport())
	if err != nil || !posted {
		t.Fatalf("posted %v, err %v, want posted with no error", posted, err)
	}
	if len(r.poster.Posts()) != 1 {
		t.Fatalf("posts = %d", len(r.poster.Posts()))
	}
}
