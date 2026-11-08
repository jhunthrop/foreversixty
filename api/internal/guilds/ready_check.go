// api/internal/guilds/ready_check.go
//
// The ready check (spec 2026-11-08-app-flows F2): one character's pre-pull list for raid night
// and Home's "2 things before you pull", from its last addon export. The checks are the D12
// checklist's (buildChecklist over computeReadiness, attunementOf and the supplies reading),
// worded as the app draws them: one line per thing to bring, each ok, miss or warn, with where
// to get what is missing. A check the export cannot answer is left out; the list never guesses.
package guilds

import (
	"context"
	"fmt"
	"strings"
	"time"
)

// The ready states.
const (
	ReadyOK   = "ok"
	ReadyMiss = "miss"
	ReadyWarn = "warn"
)

// readyCopy is the ready check's words.
var readyCopy = struct {
	LevelTitle, LevelLine, Attuned, NotAttuned, KeysUnseen, AttuneWhere                        string
	FlaskTitle, FoodTitle, NoneInBags, AuctionHouse, Enchanted, NoEnchant, EnchantWhere, Every string
}{
	LevelTitle:   "Level %d",
	LevelLine:    "You are level %d",
	Attuned:      "Attuned",
	NotAttuned:   "Not attuned",
	KeysUnseen:   "The addon has not seen your keys yet",
	AttuneWhere:  "The attunement quests",
	FlaskTitle:   "Flask or elixir",
	FoodTitle:    "Buff food",
	NoneInBags:   "None in your bags",
	AuctionHouse: "Auction house",
	Enchanted:    "%s not enchanted",
	NoEnchant:    "No enchant on it",
	EnchantWhere: "An enchanter or the auction house",
	Every:        "Every slot enchanted",
}

// enchantSlotWords is how a slot reads in "Boots not enchanted".
var enchantSlotWords = map[string]string{"main_hand": "Weapon", "chest": "Chest", "back": "Cloak", "feet": "Boots"}

// ReadyWhere is where to get what is missing; Link is a site or app link, null when none.
type ReadyWhere struct {
	Label string  `json:"label"`
	Link  *string `json:"link"`
}

// ReadyItem is one line of the ready check.
type ReadyItem struct {
	Key   string      `json:"key"`
	State string      `json:"state"`
	Title string      `json:"title"`
	Line  string      `json:"line"`
	Where *ReadyWhere `json:"where"`
}

// ReadyCheck is GET /v1/characters/{region}/{ruleset}/{name}/ready: the items and when the
// export they read was taken (null and no items without one).
type ReadyCheck struct {
	ExportedAt *time.Time  `json:"exported_at"`
	Items      []ReadyItem `json:"items"`
}

// ReadyCheckOf is key's ready check for zone (nil: no raid, so no attunement line).
func (s *Store) ReadyCheckOf(ctx context.Context, key string, zone *Zone) (ReadyCheck, error) {
	out := ReadyCheck{Items: []ReadyItem{}}
	p, found, err := s.characterPrepOf(ctx, key, "gear_bags")
	if err != nil || !found || p.exportedAt == nil {
		return out, err
	}
	out.ExportedAt = p.exportedAt
	checks, _ := buildChecklist(checklistInput{zone: zone, readiness: p.readiness, row: p.row, onRoster: true,
		band: p.band, hasBand: p.hasBand})
	out.Items = readyItems(checks, zone)
	return out, nil
}

// readyItems words the checklist's items as ready lines, in the checklist's order.
func readyItems(checks []ChecklistItem, zone *Zone) []ReadyItem {
	out := []ReadyItem{}
	for _, c := range checks {
		switch c.Key {
		case CheckLevel:
			if c.State == CheckFail && c.Level != nil && c.RequiredLevel != nil {
				out = append(out, ReadyItem{Key: "level", State: ReadyMiss,
					Title: fmt.Sprintf(readyCopy.LevelTitle, *c.RequiredLevel), Line: fmt.Sprintf(readyCopy.LevelLine, *c.Level)})
			}
		case CheckAttunement:
			if it, ok := attunementLine(c, zone); ok {
				out = append(out, it)
			}
		case CheckFlaskFood:
			out = append(out, supplyLines(c)...)
		case CheckEnchants:
			out = append(out, enchantLines(c)...)
		}
	}
	return out
}

func attunementLine(c ChecklistItem, zone *Zone) (ReadyItem, bool) {
	if zone == nil {
		return ReadyItem{}, false
	}
	_, slug, _ := strings.Cut(zone.Key, ":")
	it := ReadyItem{Key: "attune." + slug, Title: zone.Name}
	if c.KeyItem != nil {
		it.Title = c.KeyItem.Name
	}
	switch {
	case c.State == CheckOK:
		it.State, it.Line = ReadyOK, readyCopy.Attuned
	case c.State == CheckFail:
		it.State, it.Line, it.Where = ReadyMiss, readyCopy.NotAttuned, &ReadyWhere{Label: readyCopy.AttuneWhere}
	case c.Reason == reasonKeysNotInExport:
		it.State, it.Line = ReadyWarn, readyCopy.KeysUnseen
	default:
		return ReadyItem{}, false
	}
	return it, true
}

// supplyLines is the flask and the food lines, each ok with the stack that covers it.
func supplyLines(c ChecklistItem) []ReadyItem {
	if c.State == CheckNotTracked {
		return nil
	}
	line := func(key, title, missing string, kinds ...string) ReadyItem {
		for _, m := range c.Missing {
			if m == missing {
				return ReadyItem{Key: key, State: ReadyMiss, Title: title, Line: readyCopy.NoneInBags,
					Where: &ReadyWhere{Label: readyCopy.AuctionHouse}}
			}
		}
		for _, st := range c.Supplies {
			for _, k := range kinds {
				if st.Kind == k {
					return ReadyItem{Key: key, State: ReadyOK, Title: st.Name, Line: fmt.Sprintf("%d in your bags", st.Count)}
				}
			}
		}
		return ReadyItem{Key: key, State: ReadyOK, Title: title}
	}
	return []ReadyItem{
		line("flask", readyCopy.FlaskTitle, missingFlaskOrElixir, supplyFlask, supplyElixir),
		line("food", readyCopy.FoodTitle, missingFood, supplyFood),
	}
}

// enchantLines is one warn line per unenchanted slot, else one ok line.
func enchantLines(c ChecklistItem) []ReadyItem {
	switch c.State {
	case CheckOK:
		return []ReadyItem{{Key: "enchants", State: ReadyOK, Title: readyCopy.Every}}
	case CheckFail:
		out := make([]ReadyItem, 0, len(c.MissingSlots))
		for _, slot := range c.MissingSlots {
			words := enchantSlotWords[slot]
			if words == "" {
				words = slot
			}
			out = append(out, ReadyItem{Key: "enchant." + slot, State: ReadyWarn, Title: fmt.Sprintf(readyCopy.Enchanted, words),
				Line: readyCopy.NoEnchant, Where: &ReadyWhere{Label: readyCopy.EnchantWhere}})
		}
		return out
	}
	return nil
}
