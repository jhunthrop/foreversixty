package talents

import (
	"errors"
	"path/filepath"
	"testing"
)

func load(t *testing.T) *Layouts {
	t.Helper()
	l, err := Load("testdata")
	if err != nil {
		t.Fatal(err)
	}
	return l
}

func TestPointsBecomeThePositionalString(t *testing.T) {
	l := load(t)
	for _, c := range []struct {
		name   string
		points []int64
		want   string
	}{
		{"nothing spent", nil, "-"},
		{
			// Three ranks of the first Arms talent, one of Deflection.
			"one tree", []int64{105958, 105958, 105958, 105957},
			"31-",
		},
		{
			// Order does not matter: the string is positional.
			"the order points were spent in does not matter",
			[]int64{105957, 105958, 105958, 105958},
			"31-",
		},
		{
			// A gap is a zero, and trailing zeros are trimmed per tree.
			"a gap in the middle", []int64{105958, 105956},
			"101-",
		},
		{
			"both trees", []int64{105958, 105900, 105900, 105901},
			"1-21",
		},
	} {
		t.Run(c.name, func(t *testing.T) {
			got, err := l.String("warrior", c.points)
			if err != nil {
				t.Fatal(err)
			}
			if got != c.want {
				t.Fatalf("String(%v) = %q, want %q", c.points, got, c.want)
			}
		})
	}
}

func TestTheStringIsWhatTheEngineParses(t *testing.T) {
	// The engine splits on "-" and reads one digit per talent, in the
	// tree's own order; a tree with nothing in it is an empty segment.
	// Two trees means exactly one separator.
	got, err := load(t).String("warrior", []int64{105901})
	if err != nil {
		t.Fatal(err)
	}
	if got != "-01" {
		t.Fatalf("second tree only = %q, want %q", got, "-01")
	}
}

func TestAnUnknownClassOrTalentIsAnError(t *testing.T) {
	l := load(t)
	if _, err := l.String("deathknight", nil); !errors.Is(err, ErrUnknownClass) {
		t.Fatalf("err = %v, want ErrUnknownClass", err)
	}
	if _, err := l.String("warrior", []int64{999999}); !errors.Is(err, ErrUnknownTalent) {
		t.Fatalf("err = %v, want ErrUnknownTalent", err)
	}
	// Four points in a three-rank talent is a corrupt input, not a
	// four in the string: the engine would read it as a rank that
	// does not exist.
	over := []int64{105958, 105958, 105958, 105958}
	if _, err := l.String("warrior", over); !errors.Is(err, ErrTooManyRanks) {
		t.Fatalf("err = %v, want ErrTooManyRanks", err)
	}
}

func TestLoadRejectsADirectoryWithNoLayouts(t *testing.T) {
	if _, err := Load(t.TempDir()); err == nil {
		t.Fatal("an empty directory is not a talent layout")
	}
	if _, err := Load(filepath.Join(t.TempDir(), "nope")); err == nil {
		t.Fatal("a missing directory is not a talent layout")
	}
}

// TestTheRealBuildLoads runs against whatever client build the
// repository carries, without naming one: the newest directory under
// data/builds that has a talents/ in it.
func TestTheRealBuildLoads(t *testing.T) {
	dirs, err := filepath.Glob(filepath.Join("..", "..", "data", "builds", "*", "talents"))
	if err != nil {
		t.Fatal(err)
	}
	if len(dirs) == 0 {
		t.Skip("no client build in this checkout")
	}
	dir := dirs[len(dirs)-1]
	l, err := Load(dir)
	if err != nil {
		t.Fatalf("loading %s: %v", dir, err)
	}
	for _, class := range l.Classes() {
		// Spending nothing is a valid build in every class, and it
		// exercises the tree count and the ordering.
		s, err := l.String(class, nil)
		if err != nil {
			t.Errorf("%s: %v", class, err)
			continue
		}
		if s == "" {
			t.Errorf("%s: the empty build has no separators", class)
		}
		// Every node in the layout is spendable to one rank.
		for _, id := range l.nodes(class) {
			if _, err := l.String(class, []int64{id}); err != nil {
				t.Errorf("%s: one point in %d: %v", class, id, err)
			}
		}
	}
}
