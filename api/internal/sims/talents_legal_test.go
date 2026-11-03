package sims

import (
	"errors"
	"path/filepath"
	"strings"
	"testing"

	simapi "github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/talents"
)

// activeLayouts is the 70753 talent layout, the active build's.
func activeLayouts(t *testing.T) *talents.Layouts {
	t.Helper()
	l, err := talents.Load(filepath.Join("..", "..", "..", "data", "builds", "1.60.1.70753", "talents"))
	if err != nil {
		t.Fatal(err)
	}
	return l
}

func demonologyRun(build string) simapi.SimRequest {
	c := aCharacter("warlock", "orc")
	c.Talents = build
	return simapi.SimRequest{Spec: "warlock-demonology", Character: c}
}

// A server run sims only a build the game lets the character hold. The
// envelope checks a talent string's shape, never its rules, so a request
// with Decimation (row 4) and nothing above it ran as sent.
func TestAServerRunRefusesABuildTheGameRefuses(t *testing.T) {
	l := activeLayouts(t)
	for _, c := range []struct {
		name string
		req  simapi.SimRequest
		want string
	}{
		{"the character's own build", demonologyRun("-00000000000200"), "Decimation needs 15 points"},
		{"a talent loadout to compare", func() simapi.SimRequest {
			r := demonologyRun("-2355000030201")
			r.Bulk = &simapi.BulkSpec{Mode: simapi.KindTalents, Talents: []simapi.TalentLoadout{
				{Name: "kept", Talents: "-2355000030201"},
				{Name: "rushed", Talents: "-2355000000001"},
			}}
			return r
		}(), `"rushed": talents: illegal build: Fel Domination needs 2 points in Master Summoner`},
	} {
		err := talentsLegal(l, c.req)
		if !errors.Is(err, talents.ErrIllegalBuild) || !strings.Contains(err.Error(), c.want) {
			t.Errorf("%s: %v, want an illegal build saying %q", c.name, err, c.want)
		}
	}
}

func TestAServerRunWithALegalBuildPasses(t *testing.T) {
	if err := talentsLegal(activeLayouts(t), demonologyRun("-2355000030201")); err != nil {
		t.Fatal(err)
	}
	// No layout loaded is no check: the deployment logs that at startup.
	if err := talentsLegal(nil, demonologyRun("-00000000000200")); err != nil {
		t.Fatal(err)
	}
}
