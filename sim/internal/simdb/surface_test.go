package simdb

import "testing"

// The parity contract's section 5 widens SimItem with the four fields
// sim/bulk's expansion rules read. A build whose rows leave them all at
// zero is a stale simdb.bin, and Expand would then quietly stop
// enforcing unique-equipped and required level - the two rules whose
// absence looks like a correct answer. So the embedded table is asked
// whether anything at all sets them.
func TestTheEmbeddedTableCarriesTheFieldsExpansionNeeds(t *testing.T) {
	db, err := load()
	if err != nil {
		t.Fatal(err)
	}
	var unique, level, faction, suffixes int
	for _, it := range db.Items {
		if it.Unique {
			unique++
		}
		if it.RequiredLevel > 0 {
			level++
		}
		if it.FactionRestriction != 0 {
			faction++
		}
		if len(it.RandomSuffixOptions) > 0 {
			suffixes++
		}
	}
	t.Logf("items=%d unique=%d required_level=%d faction=%d suffixes=%d",
		len(db.Items), unique, level, faction, suffixes)
	for _, c := range []struct {
		name  string
		count int
	}{
		{"unique", unique},
		{"required_level", level},
		{"faction_restriction", faction},
		{"random_suffix_options", suffixes},
	} {
		if c.count == 0 {
			t.Errorf("no item sets %s; re-run the data lane's `python -m pipeline genproto simdb` for the active build", c.name)
		}
	}
}
