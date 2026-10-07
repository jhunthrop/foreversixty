package main

import (
	"os"
	"path/filepath"
	"testing"
)

// writeEngineFile drops one Go source file under engineDir/sim/<class>/,
// the layout modeledGoNames walks.
func writeEngineFile(t *testing.T, engineDir, class, name, src string) {
	t.Helper()
	dir := filepath.Join(engineDir, "sim", class)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, name), []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
}

// TestModeledGoNamesFollowsAnAliasedReceiver is the mage-fire bug: the
// scan must not require the literal receiver name "Talents" - an
// aliased local (`t := mage.Talents`) reading `t.FirePower` has to
// count exactly as "mage.Talents.FirePower" would.
func TestModeledGoNamesFollowsAnAliasedReceiver(t *testing.T) {
	dir := t.TempDir()
	writeEngineFile(t, dir, "mage", "talents.go", `package mage

func (mage *Mage) applyDeclarativeTalents() {
	t := mage.Talents
	if t.FirePower > 0 {
		doSomething(t.FirePower)
	}
}
`)
	names, err := modeledGoNames(dir, "mage")
	if err != nil {
		t.Fatal(err)
	}
	if !names["FirePower"] {
		t.Errorf("names = %v, want FirePower modeled via the t := mage.Talents alias", names)
	}
}

// TestModeledGoNamesCountsAChainedRead is the druid Furor bug: a read
// through the Talents field itself (druid.Talents.Furor) is a read, and the
// aliased-receiver pattern alone consumed "druid.Talents" and never saw it.
func TestModeledGoNamesCountsAChainedRead(t *testing.T) {
	dir := t.TempDir()
	writeEngineFile(t, dir, "druid", "talents.go", `package druid

func (druid *Druid) applyFuror() {
	if druid.Talents.Furor == 0 {
		return
	}
	cost := 100 - 10*druid.Talents.NaturalShapeshifter
	_ = cost
	_ = druid.Talents.NaturesFocus
	_, _ = druid.Talents.FeralCharge, druid.Talents.PrimalBite
}
`)
	names, err := modeledGoNames(dir, "druid")
	if err != nil {
		t.Fatal(err)
	}
	for _, n := range []string{"Furor", "NaturalShapeshifter"} {
		if !names[n] {
			t.Errorf("names[%q] = false, want it read through druid.Talents", n)
		}
	}
	// Only the discarded `_ = druid.Talents.X` acknowledgements stay out.
	for _, n := range []string{"NaturesFocus", "FeralCharge", "PrimalBite"} {
		if names[n] {
			t.Errorf("names[%q] = true, want a blank-assigned acknowledgement not to count as modeling", n)
		}
	}
}

// TestModeledGoNamesSkipsUncompiledFiles: the go tool ignores files and
// directories that start with "_", so code parked there models nothing.
func TestModeledGoNamesSkipsUncompiledFiles(t *testing.T) {
	dir := t.TempDir()
	writeEngineFile(t, dir, "druid", "_maul.go", "package druid\n\nfunc retired() { _ = druid.Talents.RendAndTear * 2 }\n")
	writeEngineFile(t, dir, "druid/_tank", "retired.go", "package tank\n\nfunc retired() { x := druid.Talents.KingOfTheJungle; use(x) }\n")
	writeEngineFile(t, dir, "druid", "live.go", "package druid\n\nfunc live() { use(druid.Talents.Furor) }\n")
	names, err := modeledGoNames(dir, "druid")
	if err != nil {
		t.Fatal(err)
	}
	if !names["Furor"] {
		t.Error("names[Furor] = false, want the compiled file read")
	}
	for _, n := range []string{"RendAndTear", "KingOfTheJungle"} {
		if names[n] {
			t.Errorf("names[%q] = true, want the uncompiled file skipped", n)
		}
	}
}

// TestModeledGoNamesStillSkipsCommentsTestsAndGenerated checks the
// generalized regex did not widen what counts as a read: a name that
// only appears in a comment, a _test.go file or a _auto_gen.go file is
// still not modeled.
func TestModeledGoNamesStillSkipsCommentsTestsAndGenerated(t *testing.T) {
	dir := t.TempDir()
	writeEngineFile(t, dir, "mage", "talents.go", `package mage

// t.CommentOnly is not a real read.
func (mage *Mage) applyDeclarativeTalents() {
}
`)
	writeEngineFile(t, dir, "mage", "talents_test.go", `package mage

func testOnly() {
	t := mage.Talents
	_ = t.TestFileOnly
}
`)
	writeEngineFile(t, dir, "mage", "talents_auto_gen.go", `package mage

func generatedOnly() {
	t := mage.Talents
	_ = t.GeneratedOnly
}
`)
	names, err := modeledGoNames(dir, "mage")
	if err != nil {
		t.Fatal(err)
	}
	for _, n := range []string{"CommentOnly", "TestFileOnly", "GeneratedOnly"} {
		if names[n] {
			t.Errorf("names[%q] = true, want it skipped (comment/test/generated)", n)
		}
	}
}

// TestClassifyPrecedence is the precedence rule between the static
// scan and the probe: the probe wins when it moves DPS beyond the
// combined error, whatever the scan found; short of that, the scan's
// "modeled" is at worst "modeled, no damage"; unmodeled only when
// neither signal finds anything.
func TestClassifyPrecedence(t *testing.T) {
	for _, c := range []struct {
		name          string
		staticModeled bool
		credit        credit
		want          talentClass
	}{
		{
			name:          "scan says unmodeled, probe beyond error -> damage wins",
			staticModeled: false,
			credit:        credit{Diff: 33.1, Err: 8.0},
			want:          classDamage,
		},
		{
			name:          "scan says modeled, probe beyond error -> still damage",
			staticModeled: true,
			credit:        credit{Diff: 33.1, Err: 8.0},
			want:          classDamage,
		},
		{
			name:          "scan says modeled, probe within error -> modeled no damage",
			staticModeled: true,
			credit:        credit{Diff: 1.0, Err: 8.0},
			want:          classModeledNoDamage,
		},
		{
			name:          "scan says unmodeled, probe within error -> unmodeled",
			staticModeled: false,
			credit:        credit{Diff: 1.0, Err: 8.0},
			want:          classUnmodeled,
		},
		{
			name:          "negative diff beyond error still counts as damage",
			staticModeled: false,
			credit:        credit{Diff: -12.0, Err: 3.0},
			want:          classDamage,
		},
		{
			name:          "diff exactly at the error boundary is not beyond it",
			staticModeled: false,
			credit:        credit{Diff: 8.0, Err: 8.0},
			want:          classUnmodeled,
		},
	} {
		if got := classify(c.staticModeled, c.credit); got != c.want {
			t.Errorf("%s: classify(%v, %+v) = %v, want %v", c.name, c.staticModeled, c.credit, got, c.want)
		}
	}
}

// TestTalentClassStringMatchesReportLabels pins the display strings
// report.go's table and prose depend on.
func TestTalentClassStringMatchesReportLabels(t *testing.T) {
	for class, want := range map[talentClass]string{
		classUnmodeled:       "unmodeled",
		classModeledNoDamage: "modeled, no damage",
		classDamage:          "damage",
	} {
		if got := class.String(); got != want {
			t.Errorf("%v.String() = %q, want %q", class, got, want)
		}
	}
}
