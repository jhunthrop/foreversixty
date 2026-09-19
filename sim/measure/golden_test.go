package measure

import (
	"flag"
	"os"
	"strings"
	"testing"
)

// update rewrites the golden file instead of comparing against it:
//
//	go test ./measure/ -run TestTable -update
var update = flag.Bool("update", false, "rewrite the golden table")

const goldenTable = "testdata/planted-v22.table"

// TestTableGolden pins the printed table. The table is the product for
// anyone reading it with a fresh log in hand, so a column that silently
// moves or a figure that silently changes shape is a regression even
// when every number behind it is still right.
func TestTableGolden(t *testing.T) {
	rep, err := Run(load(t))
	if err != nil {
		t.Fatal(err)
	}
	got := rep.Table()
	if *update {
		if err := os.WriteFile(goldenTable, []byte(got), 0o644); err != nil {
			t.Fatal(err)
		}
		t.Logf("wrote %s", goldenTable)
		return
	}
	want, err := os.ReadFile(goldenTable)
	if err != nil {
		t.Fatalf("%v; run the test with -update to create it", err)
	}
	if got != string(want) {
		t.Errorf("the printed table changed.\n--- got ---\n%s\n--- want ---\n%s", got, want)
	}
}

// A report whose floor nothing clears must print no numbers at all: the
// withholding is the feature, and it is easy to break by adding a column
// above the Enough check.
func TestTableWithholdsEveryFigureUnderTheFloor(t *testing.T) {
	in := load(t)
	in.MinSamples = 1000
	rep, err := Run(in)
	if err != nil {
		t.Fatal(err)
	}
	out := rep.Table()
	for _, forbidden := range []string{"2.000", "10.00%", "45s", "500.0"} {
		if strings.Contains(out, forbidden) {
			t.Errorf("the table printed %q from below the sample floor:\n%s", forbidden, out)
		}
	}
}
