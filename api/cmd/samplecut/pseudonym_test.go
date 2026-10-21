package main

import (
	"strings"
	"testing"
)

const sampleLines = `9/17/2026 15:54:38.383-5  SWING_DAMAGE,Creature-0-4621-2991-518-251160-00022C5381,"Al'Aketh Convert",0xa48,0x80000000,Player-4619-005746D8,"Xuen-ClassicBetaPvP-",0x518,0x80000000,Creature-0-4621-2991-518-251160-00022C5381,0000000000000000,18,55,4,0,17
9/17/2026 15:54:38.626-5  SPELL_HEAL,Player-4613-005E9B4A,"Spellshock-ClassicBetaPvP2-",0x518,0x80000000,Player-4619-005746D8,"Xuen-ClassicBetaPvP-",0x518,0x80000000,2060,"Heal",0x2,Player-4619-005746D8,0000000000000000,1,2
9/17/2026 15:54:39.000-5  SPELL_AURA_APPLIED,Player-4613-0057D2B3,"Mork-Ghul-ClassicBetaPvP2-",0x518,0x0,Player-4613-0057D2B3,"Mork-Ghul-ClassicBetaPvP2-",0x518,0x0,1126,"Mark of the Wild",0x8,BUFF
`

func TestPseudonymiserIsDeterministic(t *testing.T) {
	a := NewPseudonymiser("seed").Log([]byte(sampleLines))
	b := NewPseudonymiser("seed").Log([]byte(sampleLines))
	if string(a) != string(b) {
		t.Fatal("the same log under the same seed cut differently")
	}
	if c := NewPseudonymiser("other").Log([]byte(sampleLines)); string(c) == string(a) {
		t.Fatal("a different seed gave the same names")
	}
}

func TestPseudonymiserLeavesNoOriginalBehind(t *testing.T) {
	p := NewPseudonymiser("seed")
	out := string(p.Log([]byte(sampleLines)))
	for _, orig := range []string{"Xuen", "Spellshock", "Mork-Ghul", "4619", "4613", "005746D8", "005E9B4A", "0057D2B3", "ClassicBetaPvP2"} {
		if strings.Contains(out, orig) {
			t.Errorf("%q survived the pseudonymiser:\n%s", orig, out)
		}
	}
	if !strings.Contains(out, "Al'Aketh Convert") || !strings.Contains(out, "Mark of the Wild") {
		t.Errorf("an enemy name or a spell was rewritten:\n%s", out)
	}
	if got := len(p.Originals); got != 3 {
		t.Errorf("saw %d original units, want 3", got)
	}
}

func TestPseudonymiserKeepsOneNamePerPlayerAndNoTwoShareOne(t *testing.T) {
	p := NewPseudonymiser("seed")
	lines := strings.Split(string(p.Log([]byte(sampleLines))), "\n")
	// The second line's heal target is the first line's player: the same pseudonym and GUID.
	if g1, g2 := guidAfter(lines[0], "Player-"), strings.Count(lines[1], guidAfter(lines[0], "Player-")); g1 == "" || g2 != 2 {
		t.Errorf("the same player did not keep one GUID across lines: %q seen %d times", g1, g2)
	}
	seen := map[string]string{}
	for guid, u := range p.units {
		orig := guid
		if prev, ok := seen[u]; ok {
			t.Errorf("%q and %q share the pseudonym %q", prev, orig, u)
		}
		seen[u] = orig
		if !strings.HasSuffix(u, "-"+sampleRealm+"-") || strings.Count(u, "-") != 3 {
			t.Errorf("pseudonym %q is not First-Last-Realm-", u)
		}
	}
}

func guidAfter(line, prefix string) string {
	i := strings.Index(line, prefix)
	if i < 0 {
		return ""
	}
	return line[i : i+len("Player-4184-00000001")]
}
