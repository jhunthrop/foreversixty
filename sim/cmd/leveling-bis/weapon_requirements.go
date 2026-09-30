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

// shootSpellID is Shoot's own spell id (data/builds/1.60.1.70009/
// spellconst/*.json: every class's spellconst carries the same single,
// unranked entry) - the client's generic ranged-weapon autoshot every
// class can use, not a class ability. Confirmed present, at this exact
// id, in every caster APL that actually wand-weaves (mage-arcane/-fire/
// -frost, priest-shadow, warlock-affliction/-demonology/-destruction)
// and absent from shaman-elemental's and druid-balance's own APLs -
// bis-ranker-integrity-4 lane, caster sweep item 2: those two specs'
// rotations never cast it, so their ranged slot gets no more credit for
// a wand's raw DPS than their main_hand/off_hand do for a dagger's.
const shootSpellID = 5019

// aplRotationCastsShoot reports whether data/curated/apl/<spec>.json's
// own rotation casts Shoot anywhere in its tree - collectSpellIDs
// (above) walks the same prepullActions+priorityList structure
// aplRotationRequiresDagger already does, so a spec that wand-weaves
// via a ranked or conditional Shoot line is caught the same tolerant
// way. A spec with no apl file at all reports false, nil, matching
// aplRotationRequiresDagger's own doc for exactly the same reason: a
// spec with no rotation yet must still rank, just without this credit.
func aplRotationCastsShoot(repoRoot, spec string) (bool, error) {
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
	return ids[shootSpellID], nil
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

// casterWandOnlyClasses is which class_slugs' own ranged slot proficiency
// is Wand and nothing else (Classic's own class skill list: Mage,
// Priest, Warlock) - this lane's brief (bis-ranker-integrity-6), item
// 9.
var casterWandOnlyClasses = map[string]bool{"mage": true, "priest": true, "warlock": true}

// rangedWeaponSkillClasses is which class_slugs carry Bows/Guns/
// Crossbows/Thrown Weapons proficiency (Classic's own class skill
// list: Hunter, Warrior, Rogue) - this lane's brief's own list.
// Paladin/Shaman/Druid are deliberately absent: their only ranged-slot
// proficiency is a relic (libram/idol/totem, ClassID armorClassID),
// which restrictRangedByProficiency below never touches at all - "already
// handled" per this lane's brief, by the per-class item file eligible.go's
// own doc already delegates class legality to (a paladin's item file
// was never going to carry a bow or a wand in the first place).
var rangedWeaponSkillClasses = map[string]bool{"hunter": true, "warrior": true, "rogue": true}

// rangedWeaponTypesForSkillClasses is the candidate.WeaponType values
// rangedWeaponSkillClasses' own classes may equip - every ranged
// weapon type except Wand.
var rangedWeaponTypesForSkillClasses = map[string]bool{"bow": true, "gun": true, "crossbow": true, "thrown": true}

// restrictRangedByProficiency returns a copy of list (candidatesBySlot's
// own bySlot["ranged"]) with every ranged weapon-class candidate this
// classSlug cannot actually wield dropped - this lane's brief
// (bis-ranker-integrity-6), item 9: score()'s own Shoot fallback
// (score.go) converts ANY ranged-slot item's raw DPS into a caster's
// score once its rotation casts Shoot, with nothing checking that the
// item is actually a wand - a thrown weapon (or a bow/gun/crossbow) in
// a caster's ranged slot would score identically to a real wand, with
// nothing in score() able to tell them apart. (This lane's own first
// pass named Torch of Light 279246 and Cold Snap 19130 as the repro -
// the controller's own direct note corrected that: both are real
// wands, subclass_id 19 per the client and classic-db's item_template,
// so they are legitimate caster picks; the underlying gap in score()
// was real regardless of those two items' own type.) The real fix is
// the same shape weapon_requirements.go already uses for a dagger-only
// rotation (restrictToDaggers, above): filter the CANDIDATE POOL
// itself, before pick()/rankSlotWithEffects/score() ever see a
// candidate this class cannot legally use as its ranged weapon, rather
// than patching every downstream consumer to second-guess a candidate
// that should never have reached them.
//
// A non-weapon-class candidate (a relic, ClassID armorClassID -
// paladin/shaman/druid's own ranged slot) passes through unchanged;
// this gate is only about weapon-class ranged items (bow/gun/
// crossbow/wand/thrown - ClassID itemClassWeapon).
//
// The gate itself, by classSlug:
//   - casterWandOnlyClasses (mage/priest/warlock): WeaponType must be
//     exactly "wand". Any other value - a real bow/gun/crossbow/
//     thrown, OR "" (unknown - candidate.WeaponType's own doc: every
//     ranged row until data/pipeline's classicdb-fidelity lane's own
//     rebuild lands) - is excluded. An unsourced-type weapon is never
//     a caster's wand (this lane's brief's own words): the absence of
//     proof is not proof of a wand.
//   - rangedWeaponSkillClasses (hunter/warrior/rogue): a RECOGNISED,
//     non-wand type (rangedWeaponTypesForSkillClasses) is kept; a
//     recognised type this class cannot use (a caster's wand) is
//     excluded; an EMPTY/unrecognised type is kept rather than
//     excluded - this lane's brief scopes the immediate "exclude
//     unknown" rule to caster ranged slots alone ("until the data
//     carries it... exclude it from CASTER ranged slots"), because
//     almost every real bow/gun/crossbow/thrown item a hunter/warrior/
//     rogue actually uses today carries this same empty WeaponType,
//     and excluding them too would empty every one of those classes'
//     ranged slots over the identical data gap, not fix a defect.
//   - every other classSlug (paladin/shaman/druid): excluded outright
//   - their per-class item file was never going to hand this
//     function a weapon-class ranged candidate at all (see
//     rangedWeaponSkillClasses' own doc), so one reaching here is a
//     data anomaly, not a legal pick.
func restrictRangedByProficiency(list []scored, classSlug string) []scored {
	out := make([]scored, 0, len(list))
	for _, c := range list {
		if c.ClassID != itemClassWeapon {
			out = append(out, c)
			continue
		}
		switch {
		case casterWandOnlyClasses[classSlug]:
			if c.WeaponType == "wand" {
				out = append(out, c)
			}
		case rangedWeaponSkillClasses[classSlug]:
			if c.WeaponType == "" || rangedWeaponTypesForSkillClasses[c.WeaponType] {
				out = append(out, c)
			}
		}
	}
	return out
}
