// api/cmd/seedguild/gear.go
//
// Gear and the FS1 export string every mock character needs so addon_exports.export
// parses everywhere it is read (api/internal/fs1 decodes it for GET /v1/me; the web
// planner decodes its own copy the same way). Gear comes from the real BiS tables under
// data/builds/<build>/bis/*.json when a spec has one; a spec with no BiS file (every tank
// and healer spec here — the catalogue only ever simmed DPS) gets no gear rather than
// invented item ids, per the seed's own rule.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
)

// fs1DataBuild is the data build every mock export declares, matching the BiS tables this
// tool reads gear from (data/builds/1.60.1.70009).
const fs1DataBuild = "1.60.1.70009"

// fs1Level is every mock character's level. Forever's beta caps leveling at 30 as this
// tool is written (2026-10), but the guild page this seed exists to let the owner design
// against is the launch raid picture — a 24-raider, level-60 Barrow Deeps/Hyjal Summit/
// Onyxia's Lair roster (Forever's real first raid tier, see raid.go's own header) —
// so the mock characters are seeded at level 60 regardless of what the beta itself allows
// today.
const fs1Level = 60

// bisBuild is the data build directory the BiS tables live under.
const bisBuild = "1.60.1.70009"

// repoRoot finds the repository root from this source file's own location, so the BiS
// loader works under `go run` regardless of the caller's working directory. This file
// lives at <root>/api/cmd/seedguild/gear.go.
func repoRoot() (string, error) {
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("seedguild: could not resolve this source file's own path")
	}
	return filepath.Abs(filepath.Join(filepath.Dir(thisFile), "..", "..", ".."))
}

// bisDir is the directory a spec's BiS json lives in.
func bisDir() (string, error) {
	root, err := repoRoot()
	if err != nil {
		return "", err
	}
	return filepath.Join(root, "data", "builds", bisBuild, "bis"), nil
}

// bisAlternative is one alternative item a BiS slot lists - only the field this tool
// reads.
type bisAlternative struct {
	ItemID int `json:"item_id"`
}

// bisSlot is one slot of one band/faction entry in a BiS json file — only the fields this
// tool reads.
type bisSlot struct {
	Slot         string           `json:"slot"`
	ItemID       int              `json:"item_id"`
	Alternatives []bisAlternative `json:"alternatives"`
}

type bisBand struct {
	Band    int       `json:"band"`
	Faction string    `json:"faction"`
	Slots   []bisSlot `json:"slots"`
}

type bisFile struct {
	Bands []bisBand `json:"bands"`
}

// loadBisSlots reads file's alliance, band-60 entry whole (the full launch-raid BiS
// picture every spec file in this catalogue carries) — the top pick and every real
// alternative per slot, for loadBisGear and the readiness variants (readiness.go) to pick
// from. file is a basename under the BiS directory ("warrior-fury.json"); an empty file,
// or one with no such band, returns a nil slice rather than an error, so a tank/healer
// spec with no BiS table at all (this catalogue only ever simmed DPS specs) simply seeds
// with no gear.
func loadBisSlots(dir, file string) ([]bisSlot, error) {
	if file == "" {
		return nil, nil
	}
	raw, err := os.ReadFile(filepath.Join(dir, file))
	if err != nil {
		return nil, fmt.Errorf("seedguild: read BiS file %s: %w", file, err)
	}
	var parsed bisFile
	if err := json.Unmarshal(raw, &parsed); err != nil {
		return nil, fmt.Errorf("seedguild: parse BiS file %s: %w", file, err)
	}
	for _, band := range parsed.Bands {
		if band.Band == 60 && band.Faction == "alliance" {
			return band.Slots, nil
		}
	}
	return nil, nil
}

// loadBisGear reads slot -> item id out of file's alliance, band-60 BiS picks (the top
// pick only; see loadBisSlots for the alternatives a readiness variant swaps in).
// item_id 0 is the table's own "nothing equipped here" sentinel (a two-hander with no
// off-hand, a caster with no ranged weapon) - not a real item id, so it is left out
// rather than written as gear.
func loadBisGear(dir, file string) (map[string]int, error) {
	slots, err := loadBisSlots(dir, file)
	if err != nil {
		return nil, err
	}
	gear := make(map[string]int, len(slots))
	for _, s := range slots {
		if s.ItemID > 0 {
			gear[s.Slot] = s.ItemID
		}
	}
	return gear, nil
}

// talentString is a plausible, fully-spent FS1 talent field ("<t1>/<t2>/<t3>") for role: a
// tank or healer puts most of a level-60 character's 51 talent points in its own tree
// with a handful in a secondary utility tree, a DPS spends nearly all of them in one
// tree. The exact digits are flavour — fs1.Decode only reads the base-36 digit sum per
// tree — so nothing here needs to match a real Classic talent build, only to decode
// validly the way a genuine export does.
func talentString(role string) string {
	switch role {
	case roleTank:
		return "0023500000000000000/05000/0"
	case roleHealer:
		return "0000000000000000000/235500/05"
	default:
		return "0000000000000000000/0/5235500"
	}
}

// talentStringUnspent is talentString with several points pulled back out - the two
// UnspentTalents roster entries (roster.go), a readiness gap: a level-60 character who
// has not finished spending a respec.
func talentStringUnspent(role string) string {
	switch role {
	case roleTank:
		return "0023500000000000000/00000/0"
	case roleHealer:
		return "0000000000000000000/205500/00"
	default:
		return "0000000000000000000/0/5205000"
	}
}

// talentStringFor picks the full or the unspent variant for c.
func talentStringFor(c mockCharacter) string {
	if c.UnspentTalents {
		return talentStringUnspent(c.Role)
	}
	return talentString(c.Role)
}

// gearPiece is one equipped item: the item id, and the enchant id on it (0 for none).
type gearPiece struct {
	ItemID  int
	Enchant int
}

// encodeGearPiece matches addon/ForeverSixty/Codec.lua's own encodeItemParts: a bare item
// id with no enchant, "item:enchant" with one.
func encodeGearPiece(p gearPiece) string {
	if p.Enchant > 0 {
		return fmt.Sprintf("%d:%d", p.ItemID, p.Enchant)
	}
	return fmt.Sprintf("%d", p.ItemID)
}

// buildFS1 assembles one mock character's export string: the head (data build, class,
// race, talents, gear) plus the level, professions and bags sections, in the real
// addon's own fixed section order (Codec.lua's encodeFS1: level, bags, bank, sets,
// loadouts, professions, guild, who — bank/sets/loadouts/guild/who are not written here,
// since this tool has nothing honest to put in them). Gear keys and bag items are
// written in a fixed order so the same character always produces byte-identical output
// across runs (dry-run output and a real apply must agree, and a rerun's rows must not
// look like a "different" export to anything that diffs them).
func buildFS1(classSlug, raceSlug, talents string, gear map[string]gearPiece, professions []string, bags []int) string {
	slots := make([]string, 0, len(gear))
	for slot := range gear {
		slots = append(slots, slot)
	}
	sort.Strings(slots)
	gearParts := make([]string, 0, len(slots))
	for _, slot := range slots {
		gearParts = append(gearParts, slot+"="+encodeGearPiece(gear[slot]))
	}
	head := strings.Join([]string{"FS1", fs1DataBuild, classSlug, raceSlug, talents}, ":") +
		":" + strings.Join(gearParts, ",")

	sections := []string{fmt.Sprintf("level=%d", fs1Level)}
	if len(bags) > 0 {
		items := make([]string, len(bags))
		for i, id := range bags {
			items[i] = fmt.Sprint(id)
		}
		sections = append(sections, "bags="+strings.Join(items, ","))
	}
	if len(professions) > 0 {
		sections = append(sections, "professions="+strings.Join(professions, ","))
	}
	return head + "|" + strings.Join(sections, "|")
}
