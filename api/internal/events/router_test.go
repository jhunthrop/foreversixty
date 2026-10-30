package events

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"
)

type fakeStore struct {
	mu        sync.Mutex
	users     []int64
	prefs     map[int64]Prefs
	prefsErr  error
	claimErr  error
	claims    map[string]bool
	unclaimed []string
	audiences []Audience
}

func claimKey(uid int64, key string, ch Channel) string {
	return fmt.Sprintf("%s|%s|%d", ch, key, uid)
}

func (f *fakeStore) Recipients(_ context.Context, a Audience) ([]int64, error) {
	f.audiences = append(f.audiences, a)
	if a.Kind == ToUsers {
		return a.UserIDs, nil
	}
	return f.users, nil
}

func (f *fakeStore) PrefsOf(context.Context, []int64) (map[int64]Prefs, error) {
	return f.prefs, f.prefsErr
}

func (f *fakeStore) Claim(_ context.Context, uid int64, key string, ch Channel) (bool, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.claimErr != nil {
		return false, f.claimErr
	}
	if f.claims == nil {
		f.claims = map[string]bool{}
	}
	k := claimKey(uid, key, ch)
	if f.claims[k] {
		return false, nil
	}
	f.claims[k] = true
	return true, nil
}

func (f *fakeStore) Unclaim(_ context.Context, uid int64, key string, ch Channel) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	k := claimKey(uid, key, ch)
	delete(f.claims, k)
	f.unclaimed = append(f.unclaimed, k)
	return nil
}

type sent struct {
	uid  int64
	kind Kind
}

type fakeSink struct {
	got []sent
	err error
}

func (s *fakeSink) Deliver(_ context.Context, uid int64, e Event) error {
	s.got = append(s.got, sent{uid, e.Kind})
	return s.err
}

func newRouter(st *fakeStore) (*Router, *fakeSink, *fakeSink) {
	push, discord := &fakeSink{}, &fakeSink{}
	return &Router{Store: st, Sinks: map[Channel]Sink{ChannelPush: push, ChannelDiscord: discord}}, push, discord
}

func raidPosted(raidID int64) Event {
	return Event{Kind: RaidPosted, Audience: Audience{Kind: ToGuildMembers, GuildID: 7}, DedupeKey: fmt.Sprintf("raid.posted:%d", raidID)}
}

func TestPublishSendsEveryRecipientOnEveryDefaultChannel(t *testing.T) {
	st := &fakeStore{users: []int64{1, 2}}
	r, push, discord := newRouter(st)
	if err := r.Publish(context.Background(), raidPosted(1)); err != nil {
		t.Fatal(err)
	}
	if len(push.got) != 2 || len(discord.got) != 2 {
		t.Fatalf("push %v discord %v", push.got, discord.got)
	}
	if st.audiences[0].GuildID != 7 {
		t.Fatalf("audience %+v", st.audiences)
	}
}

func TestPublishIsIdempotentPerUserKeyAndChannel(t *testing.T) {
	st := &fakeStore{users: []int64{1}}
	r, push, _ := newRouter(st)
	for range 3 {
		if err := r.Publish(context.Background(), raidPosted(1)); err != nil {
			t.Fatal(err)
		}
	}
	if len(push.got) != 1 {
		t.Fatalf("push sent %d times", len(push.got))
	}
}

func TestPublishFollowsPreferences(t *testing.T) {
	st := &fakeStore{users: []int64{1, 2}, prefs: map[int64]Prefs{
		1: Prefs{}.With(RaidPosted, ChannelPush, false),
	}}
	r, push, discord := newRouter(st)
	if err := r.Publish(context.Background(), raidPosted(1)); err != nil {
		t.Fatal(err)
	}
	if len(push.got) != 1 || push.got[0].uid != 2 || len(discord.got) != 2 {
		t.Fatalf("push %v discord %v", push.got, discord.got)
	}
	// Off by default: nobody opted in.
	digest := Event{Kind: BisChanged, Audience: Users(1, 2), DedupeKey: "bis.changed:w1"}
	push.got, discord.got = nil, nil
	if err := r.Publish(context.Background(), digest); err != nil || len(push.got)+len(discord.got) != 0 {
		t.Fatalf("digest sent: %v %v %v", push.got, discord.got, err)
	}
}

func TestQuietHoursHoldBackPushButNotRaidStarting(t *testing.T) {
	quiet := Prefs{Quiet: QuietHours{Enabled: true, Start: 0, End: 23*60 + 59}}
	st := &fakeStore{prefs: map[int64]Prefs{1: quiet}}
	r, push, discord := newRouter(st)
	r.Now = func() time.Time { return time.Date(2026, 11, 4, 12, 0, 0, 0, time.UTC) }
	posted := Event{Kind: RaidPosted, Audience: Users(1), DedupeKey: "a"}
	if err := r.Publish(context.Background(), posted); err != nil {
		t.Fatal(err)
	}
	if len(push.got) != 0 || len(discord.got) != 1 {
		t.Fatalf("quiet: push %v discord %v", push.got, discord.got)
	}
	starting := Event{Kind: RaidStarting, Audience: Users(1), DedupeKey: "b"}
	if err := r.Publish(context.Background(), starting); err != nil {
		t.Fatal(err)
	}
	if len(push.got) != 1 || push.got[0].kind != RaidStarting {
		t.Fatalf("raid.starting held back: %v", push.got)
	}
}

func TestAFailedDeliveryIsUnclaimedAndReturned(t *testing.T) {
	st := &fakeStore{}
	r, push, discord := newRouter(st)
	push.err = errors.New("apns down")
	err := r.Publish(context.Background(), Event{Kind: SimFinished, Audience: Users(1), DedupeKey: "sim:1"})
	if err == nil || len(st.unclaimed) != 1 || len(discord.got) != 1 {
		t.Fatalf("err %v unclaimed %v discord %v", err, st.unclaimed, discord.got)
	}
	push.err = nil
	if err := r.Publish(context.Background(), Event{Kind: SimFinished, Audience: Users(1), DedupeKey: "sim:1"}); err != nil {
		t.Fatal(err)
	}
	if len(push.got) != 2 || len(discord.got) != 1 {
		t.Fatalf("retry: push %v discord %v", push.got, discord.got)
	}
}

func TestUnreadablePrefsAndClaimsStillDeliver(t *testing.T) {
	st := &fakeStore{prefsErr: errors.New("db"), claimErr: errors.New("db")}
	r, push, discord := newRouter(st)
	if err := r.Publish(context.Background(), Event{Kind: SimFinished, Audience: Users(1), DedupeKey: "k"}); err != nil {
		t.Fatal(err)
	}
	if len(push.got) != 1 || len(discord.got) != 1 {
		t.Fatalf("push %v discord %v", push.got, discord.got)
	}
}

func TestPublishRefusesWhatItCannotRoute(t *testing.T) {
	r, _, _ := newRouter(&fakeStore{})
	bad := []Event{
		{Kind: "nope", Audience: Users(1), DedupeKey: "k"},
		{Kind: SimFinished, Audience: Users(1)},
		{Kind: RaidPosted, Audience: Audience{Kind: ToGuildMembers}, DedupeKey: "k"},
		{Kind: RaidStarting, Audience: Audience{Kind: ToRaidInLate, GuildID: 1}, DedupeKey: "k"},
		{Kind: RaidShort, Audience: Users(1), DedupeKey: "k"},
		{Kind: SimFinished, Audience: Audience{Kind: "x"}, DedupeKey: "k"},
	}
	for _, e := range bad {
		if r.Publish(context.Background(), e) == nil {
			t.Errorf("%+v accepted", e)
		}
	}
}

func TestRouterWithoutSinksSendsNothing(t *testing.T) {
	st := &fakeStore{}
	r := &Router{Store: st}
	if err := r.Publish(context.Background(), raidPosted(1)); err != nil || len(st.audiences) != 0 {
		t.Fatalf("err %v audiences %v", err, st.audiences)
	}
}

func TestQuietlyIsNilSafe(t *testing.T) {
	Quietly(context.Background(), nil, nil, raidPosted(1))
	r, push, _ := newRouter(&fakeStore{})
	Quietly(context.Background(), r, nil, Event{Kind: "nope"})
	if len(push.got) != 0 {
		t.Fatal("sent")
	}
}
