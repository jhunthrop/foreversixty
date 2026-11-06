package coach

import (
	"context"
	"testing"
	"time"
)

// presetPrep is furyPrep's raid of 40 with the player made class/spec at level, its
// consumables expectation read through TablesFor as production reads it, and nothing
// drunk: no flask, no food, no potion.
func presetPrep(t *testing.T, class, spec string, level int64) Input {
	t.Helper()
	in := furyPrep(t, 40, nil)
	in.Fight = withLevel(withRole(in.Fight, in.Row.GUID, class, spec, RoleDPS), in.Row.GUID, level)
	in.Row = rosterRowOf(t, in.Fight, in.Row.GUID)
	ft := &FileTables{BuildsDir: repoBuilds, CoachDir: t.TempDir(), Specs: fakeSpecs{}}
	tb, err := ft.TablesFor(context.Background(), in.Row, furyEncounter, time.Now())
	if err != nil {
		t.Fatal(err)
	}
	in.Tables.Consumables = tb.Consumables
	return in
}

func assertNoFlaskClaim(t *testing.T, in Input) {
	t.Helper()
	o := preparation(in, CountedTime(in))
	if o.MissingFlask || (o.Finding != nil && o.Finding.Title == TitleNoFlask) {
		t.Fatalf("%s %s at %d: the raid preset carries no flask, so no flask is never missing: %+v",
			in.Row.Class, in.Row.Spec, in.Row.Level, o.Finding)
	}
}

func assertFlaskClaim(t *testing.T, in Input) {
	t.Helper()
	o := preparation(in, CountedTime(in))
	if !o.MissingFlask || o.Finding == nil || o.Finding.Title != TitleNoFlask {
		t.Fatalf("%s %s at %d: want the no-flask finding: %+v", in.Row.Class, in.Row.Spec, in.Row.Level, o.Finding)
	}
}

// DRUID-R26-6: the feral cat's raid preset carries no flask at 60.
func TestPreparation_FeralCatAt60IsNotToldNoFlask(t *testing.T) {
	assertNoFlaskClaim(t, presetPrep(t, "Druid", "Feral", 60))
}

// No flask can be drunk below 50: a level 30 Balance druid's preset carries none.
func TestPreparation_Level30BalanceIsNotToldNoFlask(t *testing.T) {
	assertNoFlaskClaim(t, presetPrep(t, "Druid", "Balance", 30))
}

// A spec whose preset carries a flask at its level still hears it.
func TestPreparation_FireMageAt60WithoutAFlaskIsToldNoFlask(t *testing.T) {
	assertFlaskClaim(t, presetPrep(t, "Mage", "Fire", 60))
}

// A log with no level keeps the role's expectation.
func TestPreparation_NoLevelKeepsTheRoleExpectation(t *testing.T) {
	assertFlaskClaim(t, presetPrep(t, "Druid", "Feral", 0))
}
