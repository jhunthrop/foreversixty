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
