package guilds

import (
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

// siteSlotWordsPath is the site's slot words, the table clientSlotWords copies.
const siteSlotWordsPath = "../../../web/src/lib/items/slot-words.ts"

// siteSlotWords reads slot-words.ts SLOT_WORDS: every `n: 'Word',` line of its object literal.
func siteSlotWords(t *testing.T) map[int]string {
	t.Helper()
	raw, err := os.ReadFile(siteSlotWordsPath)
	if err != nil {
		t.Fatal(err)
	}
	_, body, found := strings.Cut(string(raw), "const SLOT_WORDS")
	if !found {
		t.Fatal("slot-words.ts has no SLOT_WORDS")
	}
	body, _, _ = strings.Cut(body, "};")
	out := map[int]string{}
	for _, m := range regexp.MustCompile(`(\d+):\s*'([^']*)'`).FindAllStringSubmatch(body, -1) {
		n, err := strconv.Atoi(m[1])
		if err != nil {
			t.Fatal(err)
		}
		out[n] = m[2]
	}
	return out
}

func TestClientSlotWordsAreTheSites(t *testing.T) {
	site := siteSlotWords(t)
	if len(site) == 0 {
		t.Fatal("read no slot words from slot-words.ts")
	}
	if len(site) != len(clientSlotWords) {
		t.Fatalf("Go has %d slot words, the site %d", len(clientSlotWords), len(site))
	}
	for n, word := range site {
		if clientSlotWords[n] != word {
			t.Fatalf("inventory type %d: Go %q, the site %q", n, clientSlotWords[n], word)
		}
	}
}

func TestEnchantSlotsReadInTheClientsWords(t *testing.T) {
	want := map[string]string{"main_hand": "Main Hand", "chest": "Chest", "back": "Back", "feet": "Feet"}
	for _, slot := range enchantableSlots {
		if enchantableSlotLabels[slot] != want[slot] {
			t.Fatalf("%s reads %q, want %q", slot, enchantableSlotLabels[slot], want[slot])
		}
	}
}
