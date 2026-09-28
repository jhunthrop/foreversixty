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
			c:       candidate{RequiredLevel: 25},
			class:   "hunter",
			level:   20,
			faction: "horde",
			want:    false,
		},
		{
			name:    "a required level at or below the character's is allowed",
			c:       candidate{RequiredLevel: 20},
			class:   "hunter",
			level:   20,
			faction: "horde",
			want:    true,
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
