package bulk

// Options are the inputs the planner takes rather than embeds.
//
// There is one, and its zero value is the right answer: the enchant
// table is embedded in sim/internal/simdb with the items (contract
// A9), so Expand needs nothing at runtime. Options exists so a test
// can substitute a small table, and so a caller holding a different
// build's rows has a door - not because anything in production passes
// one.

import "github.com/jhunthrop/foreversixty/sim/internal/simdb"

// EnchantTable answers whether an enchant may go somewhere. The
// embedded build satisfies it through embeddedEnchants.
type EnchantTable interface {
	Fits(effectID int, slot string, item simdb.Item, class string) error
}

// Options are ExpandWith's inputs.
type Options struct {
	// Enchants overrides the embedded table. Nil is the embedded one.
	Enchants EnchantTable
}

// enchants is the table this call should use.
func (o Options) enchants() EnchantTable {
	if o.Enchants != nil {
		return o.Enchants
	}
	return embeddedEnchants{}
}

// embeddedEnchants is the active build's table, as shipped.
type embeddedEnchants struct{}

func (embeddedEnchants) Fits(effectID int, slot string, item simdb.Item, class string) error {
	return simdb.EnchantFits(effectID, slot, item, class)
}
