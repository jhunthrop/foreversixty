package main

import (
	"testing"

	"github.com/jhunthrop/foreversixty/sim/internal/inproc"
)

// A bare-band Balance druid has no Innervate below 40 and no consumables,
// and its Starfire filler spent half of a 180 s fight out of mana (about 90 s
// at band 30). Improved Wrath halves Wrath's cost and cast time, so the
// rotation falls back to Wrath while the pool is behind the fight's pace and
// Innervate is cooling down.
func TestBareBalanceDruidsDoNotSpendHalfTheFightOutOfMana(t *testing.T) {
	for _, band := range []int{20, 30, 40, 50} {
		o := options{
			repoRoot: "../../..", spec: "druid-balance", preset: "bare", band: band,
			faction: "horde", iterations: 200, seed: 1,
		}
		req, err := buildRequest(o)
		if err != nil {
			t.Fatal(err)
		}
		est, player, err := inproc.PlainRunWithRotation(req, nil)
		if err != nil {
			t.Fatal(err)
		}
		r := newReport(req, fightSecondsOf(req), est, player, nil)
		if r.SecondsOOM > 0.3*r.FightSeconds {
			t.Errorf("band %d: %.1f s of %.0f s out of mana, want under 30%% of the fight", band, r.SecondsOOM, r.FightSeconds)
		}
	}
}
