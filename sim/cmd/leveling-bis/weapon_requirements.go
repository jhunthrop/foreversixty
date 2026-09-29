package main

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
)

// This lane's brief (bis-ranker-integrity, 2026-09-29), item 5's second
// half: rogue-subtlety band 40's main_hand publishes "Ardent Custodian"
// (item 868) - a real, legally-equippable Mace1H (icon inv_mace_13,
// stats {defense: 5}, subclass_id 4 per items.json - a rogue can wear
// it, eligible.go's own doc: weapon proficiency is delegated entirely
// to the per-class item file, and Classic rogues really do have Mace
// proficiency, so proficiency.py did nothing wrong here). But
// rogue-subtlety's own rotation (data/curated/apl/rogue-subtlety.json)
// opens from Premeditation into Ambush and leans on Backstab as a
// builder - both Classic abilities the client itself refuses to cast
// without a Dagger equipped in the main hand, full stop, regardless of
// the item's own raw weapon DPS or how legally a rogue could otherwise
// wear it. This is not a proficiency gap (a mace is not "illegal" for
// a rogue in general - Assassination/Combat's own bandwidth for maces
// is real) - it is a PER-ROTATION requirement score()/eligible() have
// no way to see, so it is enforced here, from the rotation itself,
// rather than a hand-maintained "rogue-subtlety needs a dagger" fact
// that could drift from what the APL actually casts.

// daggerRequiredSpellIDs is Backstab's and Ambush's own spell ids
// across every rank this build's spellconst carries
// (data/builds/1.60.1.70009/spellconst/rogue.json: spells[<id>].name ==
// "Backstab" or "Ambush"). Both are Classic's only two dagger-only
// rogue abilities; a rotation naming either one, at any rank, cannot
// function without a dagger equipped, so this table is keyed by spell
// id rather than by spec - a future spec whose rotation also happens
// to cast Backstab/Ambush is covered the same way with no code change.
var daggerRequiredSpellIDs = map[int]bool{
	// Backstab
	53: true, 2589: true, 2590: true, 2591: true, 8721: true,
	11279: true, 11280: true, 11281: true, 25300: true,
	462709: true, 462710: true, 462711: true, 462712: true, 462713: true,
	462714: true, 462715: true, 462716: true, 462717: true,
	// Ambush
	8676: true, 8724: true, 8725: true, 11267: true, 11268: true,
	11269: true, 24337: true,
	462718: true, 462719: true, 462720: true, 462721: true, 462722: true, 462723: true,
}

// aplRotationRequiresDagger reports whether
// data/curated/apl/<spec>.json's own rotation (prepullActions and
// priorityList together - collectSpellIDs walks the whole decoded
// document, not just one of them) names a spell id in
// daggerRequiredSpellIDs anywhere in its tree. A spec with no apl file
// at all (not yet curated - writtenSpecs' own doc calls this "skipped,
// not an error") is reported as false, nil, the same tolerant reading
// writtenSpecs gives a missing file elsewhere in this package: runSpec
// runs for a single named spec regardless of its apl state (only -all
// gates on writtenSpecs), so a spec with no rotation yet must not fail
// to rank at all over a requirement that simply does not apply to it
// yet.
func aplRotationRequiresDagger(repoRoot, spec string) (bool, error) {
	path := filepath.Join(repoRoot, "data", "curated", "apl", spec+".json")
	b, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return false, nil
		}
		return false, err
	}
	var doc any
	if err := json.Unmarshal(b, &doc); err != nil {
		return false, err
	}
	ids := map[int]bool{}
	collectSpellIDs(doc, ids)
	for id := range ids {
		if daggerRequiredSpellIDs[id] {
			return true, nil
		}
	}
	return false, nil
}

// collectSpellIDs recursively gathers every numeric value found under
// a JSON key literally named "spellId", anywhere in v's arbitrarily
// nested structure, into out. The APL's own encoding nests it two ways
// - a bare {"spellId": 1787} and a ranked {"spellId": {"spellId": 1787,
// "rank": 4}} - so this recurses into every value regardless of what
// it turns out to be, rather than assuming one fixed shape.
func collectSpellIDs(v any, out map[int]bool) {
	switch t := v.(type) {
	case map[string]any:
		for k, val := range t {
			if k == "spellId" {
				if n, ok := val.(float64); ok {
					out[int(n)] = true
				}
			}
			collectSpellIDs(val, out)
		}
	case []any:
		for _, item := range t {
			collectSpellIDs(item, out)
		}
	}
}

// daggerSubclassID is the client's own ItemSubclassWeapon id for
// Dagger (data/builds/1.60.1.70009/items.json: every item this build
// calls a dagger by name carries subclass_id 15).
const daggerSubclassID = 15

// restrictToDaggers returns a copy of list with every weapon candidate
// that is not a Dagger dropped - a non-weapon (should not appear in a
// main_hand/off_hand list at all, but left alone rather than assumed
// impossible) passes through unchanged. Used only for a spec
// aplRotationRequiresDagger found true, right after candidatesBySlot
// builds bySlot["main_hand"]/["off_hand"] and before pick() or any
// later pass ever sees them - pick()'s own off_hand case (a dual-
// wielder's off hand pool is main_hand's one-handers merged in) then
// inherits the restriction for free, with no second filter needed.
func restrictToDaggers(list []scored) []scored {
	out := make([]scored, 0, len(list))
	for _, c := range list {
		if c.ClassID == itemClassWeapon && c.SubclassID != daggerSubclassID {
			continue
		}
		out = append(out, c)
	}
	return out
}
