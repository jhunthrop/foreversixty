package builds

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// AddonCodePrefix is the addon code format's version prefix, from the
// Phase 2 addon design (docs/superpowers/specs/2026-09-13-phase-2-addon-design.md).
const AddonCodePrefix = "FSB1"

// base36Max is the largest value AddonCode can encode as one base-36
// character. A tab index, talent tier, or talent column past this would
// need a second character and break the fixed-width triples Follow.lua
// decodes without a lookup table.
const base36Max = 35

// AddonCode renders b as the addon's FSB1 code:
// FSB1:<data-build>:<class-slug>:<order>:<gear>.
//
// <order> is one base-36 triple per point spent - tab index, tier,
// column - concatenated in spend order, so the addon can walk it
// against Data.lua without knowing a talent id. <gear> is the planned
// item per slot, in the planner's own slot order, each with its item
// stats, so the addon's gear scoring needs no site lookup at all.
//
// AddonCode needs the build's own data-build tree to resolve a class
// slug and a talent's tab/tier/column, so a tree version, class, or
// talent the tree data does not have is an error rather than a
// half-rendered code the addon cannot decode.
func AddonCode(data *trees.Data, b Build) (string, error) {
	tb, ok := buildData(data, b.TreeVersion)
	if !ok {
		return "", fmt.Errorf("builds: addon code: no tree data for build %s", b.TreeVersion)
	}
	class, ok := tb.Class(b.ClassID)
	if !ok {
		return "", fmt.Errorf("builds: addon code: unknown class %d in tree data %s", b.ClassID, b.TreeVersion)
	}
	order, err := addonOrder(tb, b)
	if err != nil {
		return "", err
	}
	return strings.Join([]string{
		AddonCodePrefix, b.TreeVersion, class.Slug, order, addonGear(tb, b),
	}, ":"), nil
}

// addonOrder renders one base-36 (tab, tier, column) triple per point in
// b.PointOrder, concatenated in the order the points were spent.
func addonOrder(tb *trees.Build, b Build) (string, error) {
	var sb strings.Builder
	for _, id := range b.PointOrder {
		t, ok := tb.Talent(b.ClassID, id)
		if !ok {
			return "", fmt.Errorf("builds: addon code: unknown talent %d in tree data %s", id, tb.Version)
		}
		for _, v := range [3]int{t.TreePosition, t.Tier, t.Column} {
			digit, err := base36Digit(v)
			if err != nil {
				return "", err
			}
			sb.WriteByte(digit)
		}
	}
	return sb.String(), nil
}

// base36Digit encodes v as the one base-36 character the format's
// fixed-width triples need. A tab index, tier, or column that does not
// fit is bad tree data, not a build to silently mis-render.
func base36Digit(v int) (byte, error) {
	if v < 0 || v > base36Max {
		return 0, fmt.Errorf("builds: addon code: %d does not fit one base-36 digit", v)
	}
	return strconv.FormatInt(int64(v), 36)[0], nil
}

// addonGear renders b's gear in the planner's own slot order:
// <slot>=<item-id>:<stat>=<value>;... pairs separated by commas. A slot
// whose item the tree data does not have - deleted, or from an older
// data build - still names the item id without stats: the addon shows
// a planned item it cannot score rather than losing the slot.
func addonGear(tb *trees.Build, b Build) string {
	entries := make([]string, 0, len(b.Gear))
	for _, slot := range Slots {
		itemID, ok := b.Gear[slot]
		if !ok {
			continue
		}
		entry := fmt.Sprintf("%s=%d", slot, itemID)
		if item, ok := tb.Item(b.ClassID, itemID); ok && len(item.Stats) > 0 {
			entry += ":" + statPairs(item.Stats)
		}
		entries = append(entries, entry)
	}
	return strings.Join(entries, ",")
}

// statPairs renders an item's stats as <stat>=<value>;... in a fixed
// order: map iteration is random, and the same build must render the
// same addon code every time.
func statPairs(stats map[string]int) string {
	keys := make([]string, 0, len(stats))
	for k := range stats {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	pairs := make([]string, len(keys))
	for i, k := range keys {
		pairs[i] = fmt.Sprintf("%s=%d", k, stats[k])
	}
	return strings.Join(pairs, ";")
}
