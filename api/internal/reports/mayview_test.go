package reports

import (
	"context"
	"errors"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

type stubGuildRanker struct {
	member bool
	err    error
}

func (f stubGuildRanker) GuildRank(context.Context, int64, int64) (string, bool, error) {
	if f.member {
		return "member", true, f.err
	}
	return "", false, f.err
}

func TestMayView_TheRule(t *testing.T) {
	const ownerID, guildID = int64(7), int64(9)
	owner := auth.Actor{UserID: ownerID}
	stranger := auth.Actor{UserID: ownerID + 1}
	moderator := auth.Actor{UserID: ownerID + 2, Role: "moderator"}
	owned := func(vis string) Report {
		o, g := ownerID, guildID
		return Report{Visibility: vis, OwnerID: &o, GuildID: &g}
	}
	cases := []struct {
		name   string
		vis    string
		actor  auth.Actor
		member bool
		want   bool
	}{
		{"public to anyone", Public, auth.Actor{}, false, true},
		{"unlisted to anyone", Unlisted, auth.Actor{}, false, true},
		{"private to owner", Private, owner, false, true},
		{"private not to a stranger", Private, stranger, false, false},
		{"private to a moderator", Private, moderator, false, true},
		{"guild to a verified member", GuildTo, stranger, true, true},
		{"guild not to a non-member", GuildTo, stranger, false, false},
		{"signed out never sees private", Private, auth.Actor{}, false, false},
		{"signed out never sees guild", GuildTo, auth.Actor{}, true, false},
	}
	for _, c := range cases {
		if got := MayView(t.Context(), c.actor, owned(c.vis), stubGuildRanker{member: c.member}); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}

func TestMayView_RankerFailureDeniesAndNilRankerDenies(t *testing.T) {
	g := int64(9)
	rep := Report{Visibility: GuildTo, GuildID: &g}
	actor := auth.Actor{UserID: 3}
	if MayView(t.Context(), actor, rep, stubGuildRanker{member: true, err: errors.New("db down")}) {
		t.Error("a ranker error must deny")
	}
	if MayView(t.Context(), actor, rep, nil) {
		t.Error("a nil ranker must deny a guild report")
	}
}
