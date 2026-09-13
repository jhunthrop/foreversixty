package builds

import (
	"fmt"

	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// pointsPerTier is how many points a tree needs before its tier n unlocks.
const pointsPerTier = 5

// Validate checks in against the tree data for its tree_version and returns
// one message per offending field, keyed the way the contract specifies
// ("class_id", "race_id", "tree_version", "title", "point_order",
// "point_order[7]", "gear.head"). It returns nil when the input is valid.
//
// Points are counted even when the point that spends them is itself
// rejected, so one bad point does not cascade into a message on every later
// point: the build is refused either way, and the first message is the one
// that tells the planner what to fix.
func Validate(data *trees.Data, in Input) map[string]string {
	in = in.Normalize()
	fields := map[string]string{}

	b, ok := data.Build(in.TreeVersion)
	if !ok {
		fields["tree_version"] = fmt.Sprintf("No talent data for tree version %s", in.TreeVersion)
		return fields
	}

	class, classOK := b.Class(in.ClassID)
	if !classOK {
		fields["class_id"] = fmt.Sprintf("Unknown class %d", in.ClassID)
	}
	race, raceOK := b.Race(in.RaceID)
	if !raceOK {
		fields["race_id"] = fmt.Sprintf("Unknown race %d", in.RaceID)
	}
	if classOK && raceOK && !b.ComboAllowed(in.RaceID, in.ClassID) {
		fields["race_id"] = fmt.Sprintf("%s %s is not a legal combination", race.Name, class.Name)
	}
	if classOK {
		validatePoints(b, in, class, fields)
		validateGear(b, in, class, fields)
	}
	if len([]rune(in.Title)) > MaxTitleLen {
		fields["title"] = fmt.Sprintf("Title is at most %d characters", MaxTitleLen)
	}

	if len(fields) == 0 {
		return nil
	}
	return fields
}

// validatePoints implements rules 2 to 5.
func validatePoints(b *trees.Build, in Input, class trees.Class, fields map[string]string) {
	if len(in.PointOrder) > MaxPoints {
		fields["point_order"] = fmt.Sprintf("A build has at most %d points", MaxPoints)
	}
	perTree := map[int]int{}
	perTalent := map[int]int{}
	for i, id := range in.PointOrder {
		key := fmt.Sprintf("point_order[%d]", i)
		t, ok := b.Talent(in.ClassID, id)
		if !ok {
			fields[key] = fmt.Sprintf("Talent %d is not a %s talent", id, class.Name)
			continue
		}
		need := pointsPerTier * t.Tier
		switch {
		case perTree[t.TreeID] < need:
			fields[key] = fmt.Sprintf("Tier %d of %s needs %d points in %s first", t.Tier, t.TreeName, need, t.TreeName)
		case t.PrereqTalentID != nil && perTalent[*t.PrereqTalentID] < prereqRank(t):
			prereq, _ := b.Talent(in.ClassID, *t.PrereqTalentID)
			fields[key] = fmt.Sprintf("%s needs %d points in %s first", t.Name, prereqRank(t), prereq.Name)
		case perTalent[id] >= t.MaxRank:
			fields[key] = fmt.Sprintf("%s has only %d ranks", t.Name, t.MaxRank)
		}
		perTree[t.TreeID]++
		perTalent[id]++
	}
}

// prereqRank is how many points a talent's prerequisite needs. The data
// files set prereq_rank alongside prereq_talent_id; when only the id is
// present, one point is the weakest requirement that still means something.
func prereqRank(t trees.TalentRef) int {
	if t.PrereqRank == nil {
		return 1
	}
	return *t.PrereqRank
}

// validateGear implements rule 6. Two rings or two trinkets may hold the
// same item id only when that item is not unique.
func validateGear(b *trees.Build, in Input, class trees.Class, fields map[string]string) {
	for slot, itemID := range in.Gear {
		key := "gear." + slot
		if !slotSet[slot] {
			fields[key] = fmt.Sprintf("Unknown gear slot %s", slot)
			continue
		}
		item, ok := b.Item(in.ClassID, itemID)
		if !ok {
			fields[key] = fmt.Sprintf("Item %d is not available to %s", itemID, class.Name)
			continue
		}
		if item.Slot != ItemSlot(slot) {
			fields[key] = fmt.Sprintf("%s cannot go in the %s slot", item.Name, slot)
		}
	}
	for _, pair := range [][2]string{{"finger1", "finger2"}, {"trinket1", "trinket2"}} {
		first, firstOK := in.Gear[pair[0]]
		second, secondOK := in.Gear[pair[1]]
		if !firstOK || !secondOK || first != second {
			continue
		}
		if item, ok := b.Item(in.ClassID, first); ok && item.Unique {
			fields["gear."+pair[1]] = fmt.Sprintf("Only one %s can be equipped", item.Name)
		}
	}
}
