package coach

import (
	"context"
	"slices"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5"
)

// fakeExportRow is one addon_exports read: the export, or no row.
type fakeExportRow struct {
	code string
	none bool
}

func (r fakeExportRow) Scan(dest ...any) error {
	if r.none {
		return pgx.ErrNoRows
	}
	*dest[0].(*string) = r.code
	return nil
}

type fakeExports map[string]string

func (f fakeExports) QueryRow(_ context.Context, _ string, args ...any) pgx.Row {
	code, ok := f[args[0].(string)]
	return fakeExportRow{code: code, none: !ok}
}

// warriorExport is an FS1 export of a warrior at build with rank points in the Fury tree's
// talent named talent (and nothing else).
func warriorExport(t *testing.T, build, talent string, rank int) string {
	t.Helper()
	names, err := ParseTalentNames(buildFile(t, "talents", "warrior.json"))
	if err != nil {
		t.Fatal(err)
	}
	i := slices.Index(names[1], talent)
	if i < 0 {
		t.Fatalf("no %s in the Fury tree", talent)
	}
	fury := strings.Repeat("0", i) + string(rune('0'+rank))
	return "FS1:" + build + ":warrior:human:/" + fury + "/:"
}

func TestExportTalentsReadsTheStoredBuildByName(t *testing.T) {
	build := activeBuild(t)
	src := &ExportTalents{BuildsDir: buildsDir, DB: fakeExports{
		"eu-forever-war":     warriorExport(t, build, "Improved Whirlwind", 2),
		"eu-forever-bad":     warriorExport(t, "../../etc", "Improved Whirlwind", 2),
		"eu-forever-nobuild": warriorExport(t, "9.9.9.1", "Improved Whirlwind", 2),
	}}
	ctx := context.Background()
	got, err := src.TalentsOf(ctx, "eu-forever-war", "Warrior")
	if err != nil || got["Improved Whirlwind"] != 2 || len(got) != 1 {
		t.Fatalf("got %v, %v; want Improved Whirlwind 2", got, err)
	}
	for _, tc := range []struct{ key, class string }{
		{"eu-forever-war", "Rogue"},       // another class's export is not read
		{"eu-forever-bad", "Warrior"},     // a build that is not a version is never a path
		{"eu-forever-nobuild", "Warrior"}, // a build with no tree data
		{"eu-forever-none", "Warrior"},    // no stored build
	} {
		if got, err := src.TalentsOf(ctx, tc.key, tc.class); err != nil || got != nil {
			t.Errorf("%s as %s: got %v, %v; want unknown", tc.key, tc.class, got, err)
		}
	}
}

func TestDecodeTalentsRefusesAStringThatDoesNotFit(t *testing.T) {
	names := [][]string{{"A", "B"}, {"C"}, {"D"}}
	if got, ok := DecodeTalents([]string{"12", "", "z"}, names); !ok || got["A"] != 1 || got["B"] != 2 || got["D"] != 35 {
		t.Errorf("got %v %v", got, ok)
	}
	for _, trees := range [][]string{{"123", "", ""}, {"1", "#", ""}, {"1", ""}} {
		if got, ok := DecodeTalents(trees, names); ok {
			t.Errorf("%q decoded to %v", trees, got)
		}
	}
}
