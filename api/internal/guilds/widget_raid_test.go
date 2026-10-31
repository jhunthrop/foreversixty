package guilds

import (
	"strconv"
	"testing"
	"time"
)

func (w raidWorld) widgetRaid(t *testing.T, uid int64, key string) WidgetRaid {
	t.Helper()
	wr, err := w.h.svc.WidgetRaidFor(t.Context(), uid, key, time.Now())
	if err != nil {
		t.Fatal(err)
	}
	return wr
}

func TestWidgetRaidNoGuild(t *testing.T) {
	w := newRaidWorld(t, true)
	if wr := w.widgetRaid(t, w.member, "us/hardcore/nobody"); wr.Guild != nil || wr.NextRaid != nil {
		t.Fatalf("widget = %+v, want no guild", wr)
	}
	// Another account's character is not the caller's guild either.
	if wr := w.widgetRaid(t, w.member, member2Key); wr.Guild != nil {
		t.Fatalf("widget = %+v for someone else's character", wr)
	}
}

func TestWidgetRaidNothingPlanned(t *testing.T) {
	w := newRaidWorld(t, true)
	wr := w.widgetRaid(t, w.member, memberKeyB)
	if wr.Guild == nil || wr.Guild.ID != w.gid || wr.Guild.Name != "Olympus" {
		t.Fatalf("guild = %+v", wr.Guild)
	}
	if wr.NextRaid != nil || wr.LastRaid != nil {
		t.Fatalf("widget = %+v, want nothing planned and no last raid", wr)
	}
	w.raidAt(t, raidZoneOny, time.Now().Add(-72*time.Hour))
	if last := w.widgetRaid(t, w.member, memberKeyB).LastRaid; last == nil || last.ZoneName != "Onyxia's Lair" {
		t.Fatalf("last raid = %+v", last)
	}
}

func TestWidgetRaidOwnAnswerAndCount(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)

	next := w.widgetRaid(t, w.member, memberKeyB).NextRaid
	if next == nil || next.RaidID != raid.ID || next.Status != WidgetRaidOpen || !next.CanSignUp {
		t.Fatalf("next = %+v", next)
	}
	if next.My.State != nil || next.SignedUp != 0 || next.InviteAt != nil || !next.PullAt.Equal(next.StartsAt) {
		t.Fatalf("unsigned next = %+v", next)
	}
	if next.Size == nil || *next.Size != 40 || next.GuildName != "Olympus" || next.ZoneName != "Onyxia's Lair" {
		t.Fatalf("raid facts = %+v", next)
	}

	w.as(w.member)
	w.signup(t, raid.ID, memberKeyB, `{"state":"in"}`)
	w.as(w.member2)
	w.signup(t, raid.ID, member2Key, `{"state":"late","role":"healer"}`)

	next = w.widgetRaid(t, w.member, memberKeyB).NextRaid
	if next.My.State == nil || *next.My.State != SignupIn || next.My.Role != RoleDPS {
		t.Fatalf("my = %+v", next.My)
	}
	if next.SignedUp != 2 {
		t.Fatalf("signed_up = %d, want In + Late = 2", next.SignedUp)
	}
	// The other character of the same account has no answer of its own.
	if other := w.widgetRaid(t, w.member, memberKeyA).NextRaid; other.My.State != nil {
		t.Fatalf("other character's state = %v", *other.My.State)
	}
}

func TestWidgetRaidSpecAndUnverified(t *testing.T) {
	w := newRaidWorld(t, true)
	w.createRaid(t)
	if mine := w.widgetRaid(t, w.member, memberKeyA).NextRaid.My; mine.Spec != nil {
		t.Fatalf("spec = %q before any pick, want null", *mine.Spec)
	}
	if _, err := w.saveSpec(0, w.member, memberKeyA, "warrior-protection"); err != nil {
		t.Fatal(err)
	}
	mine := w.widgetRaid(t, w.member, memberKeyA).NextRaid.My
	if mine.Spec == nil || *mine.Spec != "Protection" || mine.Role != RoleTank {
		t.Fatalf("my = %+v after the pick", mine)
	}
	if next := w.widgetRaid(t, w.member, memberKeyC).NextRaid; next == nil || next.CanSignUp {
		t.Fatalf("unverified character can sign up: %+v", next)
	}
}

func TestWidgetRaidLockedAndCancelled(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	w.exec(t, "update guild_raids set status = 'locked' where id = "+strconv.FormatInt(raid.ID, 10))
	if next := w.widgetRaid(t, w.member, memberKeyB).NextRaid; next.Status != WidgetRaidLocked || next.CanSignUp {
		t.Fatalf("locked next = %+v", next)
	}
	w.exec(t, "update guild_raids set status = 'cancelled' where id = "+strconv.FormatInt(raid.ID, 10))
	if next := w.widgetRaid(t, w.member, memberKeyB).NextRaid; next == nil || next.Status != WidgetRaidCancelled || next.CanSignUp {
		t.Fatalf("cancelled next = %+v", next)
	}
}

func TestLiveRaidEndIsForMembersOnly(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	end, found, err := w.h.svc.LiveRaidEnd(t.Context(), w.member, raid.ID)
	if err != nil || !found || !end.Equal(raid.StartsAt.Add(time.Duration(raid.DurationMinutes)*time.Minute)) {
		t.Fatalf("member: %v %v %v", end, found, err)
	}
	stranger := seedUser(t, w.h.pool, "stranger@example.com")
	if _, found, err := w.h.svc.LiveRaidEnd(t.Context(), stranger, raid.ID); err != nil || found {
		t.Fatalf("stranger found the raid: %v %v", found, err)
	}
	if _, found, err := w.h.svc.LiveRaidEnd(t.Context(), w.member, raid.ID+999); err != nil || found {
		t.Fatalf("unknown raid found: %v %v", found, err)
	}
}
