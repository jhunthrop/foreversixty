package parse

import (
	"context"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
)

const hookURL = "https://discord.com/api/webhooks/1/token"

type allowAll struct{}

func (allowAll) Can(context.Context, int64, entitlements.Feature) (bool, entitlements.Reason, error) {
	return true, entitlements.ReasonEntitled, nil
}

// notifying wires a saved webhook for the report's owner into d.
func notifying(t *testing.T, d *Deps, reportID string) *notify.FakePoster {
	t.Helper()
	return notifyingWith(t, d, reportID, map[notify.EventName]bool{notify.EventSimFinished: true, notify.EventUploadParsed: true})
}

// notifyingWith is notifying with the webhook's events chosen.
func notifyingWith(t *testing.T, d *Deps, reportID string, events map[notify.EventName]bool) *notify.FakePoster {
	t.Helper()
	rep, err := d.Reports.Get(context.Background(), reportID)
	if err != nil {
		t.Fatal(err)
	}
	sealer, err := notify.NewSealer(make([]byte, 32))
	if err != nil {
		t.Fatal(err)
	}
	box, err := sealer.Seal(hookURL)
	if err != nil {
		t.Fatal(err)
	}
	st := &notify.Store{Pool: d.Reports.Pool}
	if err := st.Save(context.Background(), notify.SaveInput{UserID: *rep.OwnerID, Kind: notify.KindDiscord, ChannelID: "1", TokenBox: box, Events: events, Source: notify.SourcePasted}); err != nil {
		t.Fatal(err)
	}
	poster := &notify.FakePoster{}
	d.Notify = &notify.Deliverer{
		Store: st, Sealer: sealer, Entitlements: allowAll{}, Poster: poster,
		PublicBaseURL: "https://foreversixty.gg",
	}
	return poster
}

func TestACompletedUploadPostsOneEmbed(t *testing.T) {
	d, reportID, _, _ := fixtureUpload(t, false)
	poster := notifying(t, &d, reportID)
	if err := Report(context.Background(), d, reportID); err != nil {
		t.Fatal(err)
	}
	posts := poster.Posts()
	if len(posts) != 1 {
		t.Fatalf("posts = %+v", posts)
	}
	e := posts[0].Embed
	if e.Title != "Tuesday finished parsing" || e.URL != "https://foreversixty.gg/reports/"+reportID {
		t.Fatalf("embed = %+v", e)
	}
	last := e.Fields[len(e.Fields)-1]
	if last.Name != "Visible to" || last.Value != "Everyone" {
		t.Fatalf("fields = %+v", e.Fields)
	}
}

func TestAFailedParsePostsNothing(t *testing.T) {
	d, reportID, _, _ := fixtureUpload(t, false)
	poster := notifying(t, &d, reportID)
	d.Rank = failingRanker{}
	if err := Report(context.Background(), d, reportID); err == nil {
		t.Fatal("want the failure")
	}
	if n := len(poster.Posts()); n != 0 {
		t.Fatalf("%d posts for a failed parse", n)
	}
}

func TestALiveReportPostsNothing(t *testing.T) {
	d, _, _, _ := fixtureUpload(t, false)
	ctx := context.Background()
	owner, err := (&auth.Store{Pool: d.Reports.Pool}).UpsertEmailUser(ctx, "live@example.com")
	if err != nil {
		t.Fatal(err)
	}
	live := auth.NewReportID()
	if _, err := d.Reports.Create(ctx, reports.Report{
		ID: live, OwnerID: &owner.ID, Title: "Live", Visibility: reports.Public, Status: reports.StatusLive,
	}); err != nil {
		t.Fatal(err)
	}
	poster := notifying(t, &d, live)
	if err := Report(ctx, d, live); err == nil {
		t.Fatal("a live report has no upload to parse")
	}
	if n := len(poster.Posts()); n != 0 {
		t.Fatalf("%d posts for a live report", n)
	}
}

func TestAPrivateReportPostsOnlyWhenPrivateLogsAreOn(t *testing.T) {
	for _, c := range []struct {
		name    string
		private bool
		posts   int
	}{{"flag off", false, 0}, {"flag on", true, 1}} {
		t.Run(c.name, func(t *testing.T) {
			d, reportID, _, _ := fixtureUpload(t, false)
			if _, err := d.Reports.Pool.Exec(context.Background(),
				`update reports set visibility = 'private' where id = $1`, reportID); err != nil {
				t.Fatal(err)
			}
			poster := notifyingWith(t, &d, reportID, map[notify.EventName]bool{
				notify.EventUploadParsed: true, notify.EventUploadPrivate: c.private,
			})
			if err := Report(context.Background(), d, reportID); err != nil {
				t.Fatal(err)
			}
			posts := poster.Posts()
			if len(posts) != c.posts {
				t.Fatalf("%d posts, want %d", len(posts), c.posts)
			}
			if c.posts == 1 {
				last := posts[0].Embed.Fields[len(posts[0].Embed.Fields)-1]
				if last.Name != "Visible to" || last.Value != "Only you" {
					t.Fatalf("fields = %+v", posts[0].Embed.Fields)
				}
			}
		})
	}
}
