package rating

import (
	"context"
	"reflect"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/consumables"
	ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// part is one component of a test card: its name, score and the weight the card applied;
// a negative score marks it excluded.
type part struct {
	name          string
	score, weight float64
}

// lossCard is a scored card from parts, in the engine's own component order.
func lossCard(parts ...part) ratingengine.Card {
	var c ratingengine.Card
	for i, p := range parts {
		c.Components[i] = ratingengine.Component{Name: p.name, Score: p.score, Weight: p.weight, Excluded: p.score < 0}
	}
	return c
}

func TestBiggestLossNamesTheComponentThatCostTheMostPoints(t *testing.T) {
	const player = "p1"
	row := summary.RosterRow{GUID: player, Name: "Lossy", ActiveMS: 150000}
	alive := summary.Summary{DurationMS: 180000, Roster: []summary.RosterRow{row}}
	diedAt := func(atMS int64) summary.Summary {
		f := alive
		f.Deaths = []summary.Death{{GUID: player, AtMS: atMS}}
		return f
	}
	// every scored, survival scored 100 and carrying a death moment named cause
	withCause := func(cause string, survival float64) ratingengine.Card {
		c := lossCard(part{"output", 90, 0.4}, part{"survival", survival, 0.2}, part{"mechanics", 100, 0.1},
			part{"utility", 100, 0.1}, part{"preparation", 100, 0.1}, part{"activity", 100, 0.1})
		c.Components[1].Moments = []ratingengine.Moment{{Kind: "death", AtMS: 1, SpellName: cause}}
		return c
	}

	for _, c := range []struct {
		name        string
		fight       summary.Summary
		card        ratingengine.Card
		phrase      string
		lossSeconds int
	}{
		{"downtime: the seconds alive and idle",
			alive,
			lossCard(part{"output", 95, 0.4}, part{"survival", 100, 0.2}, part{"mechanics", 100, 0.1},
				part{"utility", 100, 0.1}, part{"preparation", 100, 0.1}, part{"activity", 40, 0.1}),
			"downtime", 30},
		{"downtime is the activity component's own idle time, rounded",
			diedAt(160000),
			lossCard(part{"output", 95, 0.4}, part{"survival", 100, 0.2}, part{"mechanics", 100, 0.1},
				part{"utility", 100, 0.1}, part{"preparation", 100, 0.1}, part{"activity", 40, 0.1}),
			"downtime", 10},
		{"a death in the first half is dead early",
			diedAt(60000), withCause("Lava Burst", 20), "dead early", 120},
		{"a later death names its cause, lower-cased",
			diedAt(150000), withCause("Lava Burst", 20), "died to lava burst", 30},
		{"a death with no cause is a death, never Unknown cause",
			diedAt(150000), withCause("Unknown cause", 20), "a death", 30},
		{"survival with no death is avoidable damage",
			alive, withCause("", 20), "standing in fire", 0},
		{"preparation",
			alive,
			lossCard(part{"output", 100, 0.4}, part{"survival", 100, 0.2}, part{"mechanics", 100, 0.1},
				part{"utility", 100, 0.1}, part{"preparation", 0, 0.1}, part{"activity", 100, 0.1}),
			"missing consumables", 0},
		{"mechanics, with the seconds dead",
			diedAt(170000),
			lossCard(part{"output", 100, 0.4}, part{"survival", 95, 0.2}, part{"mechanics", 10, 0.1},
				part{"utility", 100, 0.1}, part{"preparation", 100, 0.1}, part{"activity", 100, 0.1}),
			"missed mechanics", 10},
		{"output",
			alive,
			lossCard(part{"output", 50, 0.4}, part{"survival", 60, 0.2}, part{"mechanics", 100, 0.1},
				part{"utility", 100, 0.1}, part{"preparation", 100, 0.1}, part{"activity", 100, 0.1}),
			"low output", 0},
		{"utility, the sixth component, by its own name",
			alive,
			lossCard(part{"output", 100, 0.4}, part{"survival", 100, 0.2}, part{"mechanics", 100, 0.1},
				part{"utility", 0, 0.1}, part{"preparation", 100, 0.1}, part{"activity", 100, 0.1}),
			"utility", 0},
		{"an excluded component costs nothing",
			alive,
			lossCard(part{"output", -1, 0}, part{"survival", 100, 0.5}, part{"mechanics", 100, 0.1},
				part{"utility", 100, 0.1}, part{"preparation", 90, 0.2}, part{"activity", 100, 0.1}),
			"missing consumables", 0},
		{"a weighted shortfall beats a bigger unweighted one",
			alive,
			lossCard(part{"output", 80, 0.5}, part{"survival", 100, 0.2}, part{"mechanics", 100, 0.1},
				part{"utility", 100, 0.1}, part{"preparation", 100, 0.05}, part{"activity", 0, 0.05}),
			"low output", 0},
		{"a perfect card has no loss",
			alive,
			lossCard(part{"output", 100, 0.4}, part{"survival", 100, 0.2}, part{"mechanics", 100, 0.1},
				part{"utility", 100, 0.1}, part{"preparation", 100, 0.1}, part{"activity", 100, 0.1}),
			"", 0},
	} {
		t.Run(c.name, func(t *testing.T) {
			// The Activity component's idle time (ratingengine.ActivityIdleMS): the downtime
			// cases give their own; every other case gets 7 s it must not report.
			idle, ok := map[string]int64{
				"downtime: the seconds alive and idle":                        30000,
				"downtime is the activity component's own idle time, rounded": 9600,
			}[c.name]
			if !ok {
				idle = 7000
			}
			phrase, seconds := biggestLoss(c.fight, row, c.card, idle)
			if phrase != c.phrase || seconds != c.lossSeconds {
				t.Fatalf("biggestLoss = %q, %d; want %q, %d", phrase, seconds, c.phrase, c.lossSeconds)
			}
		})
	}
}

func TestConsumablesMissedNamesThePlainCategories(t *testing.T) {
	dps, _ := consumables.Load().For("dps")
	row := summary.RosterRow{GUID: "g", Name: "Prepared"}
	missed := func(auras []summary.AuraTrack) []string {
		return consumablesMissed(summary.Summary{Combatants: []summary.CombatantRow{preparedCombatant("g")}, Auras: auras}, row)
	}
	if got := missed(preparedPull("g")); got == nil || len(got) != 0 {
		t.Errorf("fully prepared = %#v, want an empty list, not nil", got)
	}
	if got := missed(preparedPull("g", dps.Flask[0].SpellID)); !reflect.DeepEqual(got, []string{"flask"}) {
		t.Errorf("no flask = %v", got)
	}
	if got := missed(preparedPull("g", dps.Flask[0].SpellID, dps.Food[0].SpellID)); !reflect.DeepEqual(got, []string{"flask", "food"}) {
		t.Errorf("neither = %v", got)
	}
	// Weapon stones and oils stay out until a beta log shows how they appear.
	if got := missed(preparedPull("g", dps.WeaponEnchant[0].SpellID)); len(got) != 0 {
		t.Errorf("no weapon stone = %v, want nothing missed", got)
	}
	if got := consumablesMissed(summary.Summary{Auras: preparedPull("g")}, row); got == nil || len(got) != 0 {
		t.Errorf("no combatant info = %#v, want an empty list", got)
	}
}

// RateFight stores the last-raid facts on every player's row, from the card it scored.
func TestRateFightStoresTheLastRaidFacts(t *testing.T) {
	pool := testPool(t)
	store := &Store{Pool: pool}
	ctx := context.Background()
	dps, _ := consumables.Load().For("dps")
	row := summary.RosterRow{GUID: "lr-g", Name: "Lastraider", Class: "Warrior", Spec: "Arms", Role: "dps",
		ActiveMS: 100000}
	f := fightFixture("last-raid-facts", true, row)
	f.Summary.Deaths = []summary.Death{{GUID: "lr-g", AtMS: 150000}}
	f.Summary.Combatants = []summary.CombatantRow{preparedCombatant("lr-g")}
	f.Summary.Auras = preparedPull("lr-g", dps.Flask[0].SpellID)
	t.Cleanup(func() { pool.Exec(context.Background(), `delete from rating_scores where report_id = $1`, f.ReportID) })
	if err := store.RateFight(ctx, f); err != nil {
		t.Fatal(err)
	}

	var (
		deaths, lossSeconds *int
		biggestLoss         *string
		missed              []string
	)
	if err := pool.QueryRow(ctx,
		`select deaths, biggest_loss, loss_seconds, consumables_missed from rating_scores where report_id = $1`,
		f.ReportID).Scan(&deaths, &biggestLoss, &lossSeconds, &missed); err != nil {
		t.Fatal(err)
	}
	if deaths == nil || *deaths != 1 {
		t.Errorf("deaths = %v, want 1", deaths)
	}
	if !reflect.DeepEqual(missed, []string{"flask"}) {
		t.Errorf("consumables_missed = %v, want [flask]", missed)
	}
	// The card itself depends on whatever digests the shared database holds; the stored phrase
	// must be one biggestLoss can produce, and its seconds the ones that phrase carries.
	if biggestLoss == nil || lossSeconds == nil {
		t.Fatalf("biggest_loss = %v, loss_seconds = %v: a death, no flask and idle time must cost something",
			biggestLoss, lossSeconds)
	}
	wantSeconds := map[string]int{
		"downtime": 50, "a death": 30, "missing consumables": 30, "missed mechanics": 30,
		"low output": 30, "utility": 30,
	}
	if want, ok := wantSeconds[*biggestLoss]; !ok || *lossSeconds != want {
		t.Errorf("biggest_loss = %q with %d seconds, want a known phrase with its seconds", *biggestLoss, *lossSeconds)
	}
}
