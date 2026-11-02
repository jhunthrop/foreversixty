package main

// pairMate is what a paired slot's other item rules out: the same item (by
// id), another quality of it (by name), and, from round 9, anything in the
// same client limit category. The client's ItemLimitCategory rows that hold
// equippable items in 1.60.1.70707 all have Quantity 1 ("Unique-Equipped:
// Seal of Battle (1)"), so a shared category means one of the two.
type pairMate struct {
	ID            int
	Name          string
	LimitCategory int
}

// mateOf is the pair mate a picked item makes; nil is no mate.
func mateOf(item *scored) pairMate {
	if item == nil {
		return pairMate{}
	}
	return pairMate{ID: item.ID, Name: item.Name, LimitCategory: item.LimitCategory}
}

// excludes reports whether c may not be worn beside the mate.
func (m pairMate) excludes(c candidate) bool {
	if m.ID != 0 && c.ID == m.ID {
		return true
	}
	if m.Name != "" && c.Name == m.Name {
		return true
	}
	return m.LimitCategory != 0 && c.LimitCategory == m.LimitCategory
}
