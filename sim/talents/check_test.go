package talents

import (
	"errors"
	"strings"
	"testing"
)

// The Demonology tree of the 70753 warlock, in the string's (tier, column)
// order: 0 Improved Health Funnel (2), 1 Improved Imp (3), 2 Demonic
// Embrace (5), 3 Unholy Power (5), 4 Demonic Aegis (2), 5 Improved
// Voidwalker (3), 6 Fel Vitality (3), 7 Demonic Energies (2), 8 Improved
// Sayaad (3), 9 Demonic Sacrifice (1), 10 Master Summoner (2), 11
// Decimation (2), 12 Fel Domination (1, needs Master Summoner 2), 13
// Demonic Brand (3), 14 Improved Felhunter (3), 15 Soul Link (1, needs
// Demonic Sacrifice), 16 Demonic Knowledge (3), 17 Master Demonologist
// (5), 18 Demonic Pact (1, needs Soul Link).
func warlock(t *testing.T) *Layouts {
	t.Helper()
	l, err := Load("testdata/legality")
	if err != nil {
		t.Fatal(err)
	}
	return l
}

func TestALegalBuildPasses(t *testing.T) {
	l := warlock(t)
	for _, c := range []struct {
		name    string
		talents string
		level   int
	}{
		{"nothing spent", "--", 60},
		{"nothing spent below level 10", "", 5},
		// 5 Demonic Embrace, 5 Unholy Power: row 2 opens.
		{"two first-row talents at level 19", "-0055", 19},
		// 2+3+5+5 = 15 in rows 1 and 2, 3 in row 3, 2 Master Summoner,
		// then Fel Domination (row 4) with its prerequisite and 15 above.
		{"a prerequisite met", "-2355000030201", 60},
		// The same 21 points, spent exactly at level 30.
		{"exactly the level's points", "-2355000030201", 30},
	} {
		if err := l.Check("warlock", c.talents, c.level); err != nil {
			t.Errorf("%s: %q at level %d refused: %v", c.name, c.talents, c.level, err)
		}
	}
}

func TestAnIllegalBuildIsRefusedAndSaysWhy(t *testing.T) {
	l := warlock(t)
	for _, c := range []struct {
		name    string
		talents string
		level   int
		want    string
	}{
		// Decimation sits in row 4 and needs 15 points above it.
		{"a row not yet open", "-00000000000200", 60, "Decimation needs 15 points in Demonology above its row"},
		// 15 above, but Master Summoner has 0 of the 2 Fel Domination needs.
		{"a prerequisite missing", "-2355000000001", 60, "Fel Domination needs 2 points in Master Summoner"},
		{"more points than the level gives", "-2355000030201", 29, "21 points; a level 29 character has 20"},
		{"more than 51 points", "25522323135211351-2355233231221331351-", 60, "96 points; a level 60 character has 51"},
		{"a talent past its last rank", "3", 60, "Improved Life Tap has 3 of 2 points"},
		{"more digits than the tree has", "-00000000000000000000", 60, "Demonology has 19 talents"},
		{"four trees", "---", 60, "4 trees"},
		{"a character, not a digit", "x", 60, "not a digit"},
	} {
		err := l.Check("warlock", c.talents, c.level)
		if err == nil {
			t.Errorf("%s: %q at level %d accepted", c.name, c.talents, c.level)
			continue
		}
		if !errors.Is(err, ErrIllegalBuild) {
			t.Errorf("%s: %v is not ErrIllegalBuild", c.name, err)
		}
		if !strings.Contains(err.Error(), c.want) {
			t.Errorf("%s: %q does not say %q", c.name, err, c.want)
		}
	}
}

func TestCheckRefusesAnUnknownClass(t *testing.T) {
	if err := warlock(t).Check("monk", "", 60); !errors.Is(err, ErrUnknownClass) {
		t.Fatalf("got %v, want ErrUnknownClass", err)
	}
}
