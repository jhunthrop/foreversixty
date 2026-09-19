package request

// The client's consumable table.
//
// data/builds/<build>/simconsumes.json is the data lane's list of every
// consumable item the client carries: {id, name, quality,
// required_level, spell_ids}. The settings bar offers those items, so a
// request names a consumable by its item id, and this table is how an
// item id reaches the engine's Consumes message.
//
// The join is by name: the engine names its enum values after the items
// themselves, so "Elixir of the Mongoose" is AgilityElixir's
// ElixirOfTheMongoose. Nothing hand-written pairs the two, and nothing
// needs regenerating when the engine gains a flask - the pairing is
// recomputed from whichever build's file is loaded. An item the engine
// has no value for (most of the 1,579 rows are food and potions no sim
// models) simply does not resolve, and saying so at the boundary is the
// point.
//
// The file is not embedded. It belongs to a build, the site already
// syncs data/builds/<build>/ for the planner and the report's tooltips,
// and the wasm fetches from the same place; embedding it would pin one
// build into the binary. The api lane loads it once at startup and
// passes it in Options. sim/request/testdata/simconsumes.excerpt.json is
// a ten-row excerpt of the real file, for the tests.

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// itemPrefix marks a consumable id that names a client item rather than
// an engine enum value: "item:13452".
const itemPrefix = "item"

// Consumables is a build's consumable item table, as loaded from
// data/builds/<build>/simconsumes.json. The zero value is unusable; use
// LoadConsumables.
type Consumables struct {
	// names maps an item id onto the lookup key its name normalises to,
	// which is the same key an engine enum value name normalises to.
	names map[int64]string
}

// simConsume is one row of simconsumes.json. Only the id and the name
// take part in the join; the rest is the planner's.
type simConsume struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// LoadConsumables reads a build's simconsumes.json.
func LoadConsumables(r io.Reader) (*Consumables, error) {
	var rows []simConsume
	if err := json.NewDecoder(r).Decode(&rows); err != nil {
		return nil, fmt.Errorf("request: simconsumes.json is not readable: %w", err)
	}
	if len(rows) == 0 {
		return nil, fmt.Errorf("request: simconsumes.json carries no consumables")
	}
	c := &Consumables{names: make(map[int64]string, len(rows))}
	for _, row := range rows {
		if row.ID == 0 || row.Name == "" {
			return nil, fmt.Errorf("request: simconsumes.json has a row with no id or no name: %+v", row)
		}
		c.names[row.ID] = normalizeName(row.Name)
	}
	return c, nil
}

// Len is how many consumable items the table carries.
func (c *Consumables) Len() int { return len(c.names) }

// IDs lists every consumable id this table adds to the vocabulary,
// sorted. Only the items the engine has a value for are listed, because
// only those can be simulated.
func (c *Consumables) IDs() []string {
	out := make([]string, 0, len(c.names))
	for id, key := range c.names {
		if len(consumeFields(consumesDescriptor(), "", key)) == 0 {
			continue
		}
		out = append(out, itemPrefix+":"+strconv.FormatInt(id, 10))
	}
	sort.Strings(out)
	return out
}

// key returns the lookup key for an item id.
func (c *Consumables) key(id int64) (string, bool) {
	if c == nil {
		return "", false
	}
	k, ok := c.names[id]
	return k, ok
}

// normalizeName turns an item name into the key an engine enum value
// name normalises to: "Elixir of the Mongoose" and ElixirOfTheMongoose
// both become elixir_of_the_mongoose. Punctuation is dropped, because
// the engine's names carry none.
func normalizeName(name string) string {
	var b strings.Builder
	b.Grow(len(name))
	pendingSep := false
	for _, r := range name {
		switch {
		case unicode.IsLetter(r) || unicode.IsDigit(r):
			if pendingSep && b.Len() > 0 {
				b.WriteByte('_')
			}
			pendingSep = false
			b.WriteRune(unicode.ToLower(r))
		case r == ' ' || r == '-' || r == '_':
			pendingSep = true
		default:
			// An apostrophe or a full stop joins the word it splits:
			// "Dirge's Kick" is DirgesKick to the engine.
		}
	}
	return b.String()
}
