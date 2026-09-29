package main

import "testing"

func TestEligible(t *testing.T) {
	cases := []struct {
		name    string
		c       candidate
		class   string
		level   int
		faction string
		want    bool
	}{
		{
			name:    "a required level above the character's is refused",
			c:       candidate{RequiredLevel: 25, EffectiveRequiredLevel: 25},
			class:   "hunter",
			level:   20,
			faction: "horde",
			want:    false,
		},
		{
			name:    "a required level at or below the character's is allowed",
			c:       candidate{RequiredLevel: 20, EffectiveRequiredLevel: 20},
			class:   "hunter",
			level:   20,
			faction: "horde",
			want:    true,
		},
		{
			// 2026-09-28 quest-levels lane: eligible() gates on
			// EffectiveRequiredLevel, not RequiredLevel - a quest
			// reward's own required_level (0 here, the Deadhead Blade
			// shape) is not the real gate once a quest floors it higher.
			name:    "a quest-floored effective level above the character's is refused even when required_level is 0",
			c:       candidate{RequiredLevel: 0, EffectiveRequiredLevel: 60},
			class:   "hunter",
			level:   10,
			faction: "horde",
			want:    false,
		},
		{
			name:    "mail before 40 is refused for a hunter",
			c:       candidate{ClassID: armorClassID, SubclassID: armorSubmailID},
			class:   "hunter",
			level:   39,
			faction: "horde",
			want:    false,
		},
		{
			name:    "mail from 40 is allowed for a hunter",
			c:       candidate{ClassID: armorClassID, SubclassID: armorSubmailID},
			class:   "hunter",
			level:   40,
			faction: "horde",
			want:    true,
		},
		{
			name:    "leather is allowed at any level for a hunter",
			c:       candidate{ClassID: armorClassID, SubclassID: armorSubleatherID},
			class:   "hunter",
			level:   10,
			faction: "horde",
			want:    true,
		},
		{
			name:    "cloth is allowed at any level for a hunter",
			c:       candidate{ClassID: armorClassID, SubclassID: armorSubclothID},
			class:   "hunter",
			level:   10,
			faction: "horde",
			want:    true,
		},
		{
			name:    "a non-armor item (class_id != 4) skips the armor gate entirely",
			c:       candidate{ClassID: 2, SubclassID: armorSubmailID, RequiredLevel: 1},
			class:   "hunter",
			level:   10,
			faction: "horde",
			want:    true,
		},
		{
			name:    "a faction-restricted item matching the character's faction is allowed",
			c:       candidate{FactionRestriction: "horde"},
			class:   "hunter",
			level:   20,
			faction: "horde",
			want:    true,
		},
		{
			name:    "a faction-restricted item for the other faction is refused",
			c:       candidate{FactionRestriction: "alliance"},
			class:   "hunter",
			level:   20,
			faction: "horde",
			want:    false,
		},
		{
			name:    "an item with no faction restriction is allowed for either faction",
			c:       candidate{FactionRestriction: ""},
			class:   "hunter",
			level:   20,
			faction: "alliance",
			want:    true,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := eligible(tc.c, tc.class, tc.level, tc.faction); got != tc.want {
				t.Errorf("eligible(%+v, %q, %d, %q) = %v, want %v", tc.c, tc.class, tc.level, tc.faction, got, tc.want)
			}
		})
	}
}

// items.json spells a restriction "alliance_only"/"horde_only"; the bands
// are ranked for "alliance"/"horde". The candidate loader maps one onto
// the other, and eligible() then admits the item for its own faction only.
func TestFactionRestrictionMatchesTheBandFaction(t *testing.T) {
	if got := factionOfRestriction("alliance_only"); got != "alliance" {
		t.Fatalf("alliance_only -> %q, want alliance", got)
	}
	if got := factionOfRestriction(""); got != "" {
		t.Fatalf("empty restriction -> %q, want empty", got)
	}
	tunic := candidate{ID: 2041, RequiredLevel: 14, ClassID: 4, SubclassID: 2, FactionRestriction: factionOfRestriction("alliance_only")}
	if !eligible(tunic, "hunter", 20, "alliance") {
		t.Fatal("an alliance-only quest reward must be eligible for an alliance band")
	}
	if eligible(tunic, "hunter", 20, "horde") {
		t.Fatal("an alliance-only quest reward must not be eligible for a horde band")
	}
}
