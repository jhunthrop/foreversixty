package main

import "testing"

func TestMockRosterHasTwentyFourUniqueCharacters(t *testing.T) {
	roster := mockRoster()
	if len(roster) != 24 {
		t.Fatalf("len(roster) = %d, want 24", len(roster))
	}
	seen := map[string]bool{}
	for _, c := range roster {
		if seen[c.Name] {
			t.Fatalf("duplicate name %q", c.Name)
		}
		seen[c.Name] = true
	}
}

func TestMockRosterCoversAllNineClasses(t *testing.T) {
	want := []string{"warrior", "paladin", "hunter", "rogue", "priest", "shaman", "mage", "warlock", "druid"}
	got := classes(mockRoster())
	if len(got) != len(want) {
		t.Fatalf("classes = %v, want exactly %v", got, want)
	}
	for _, c := range want {
		if !got[c] {
			t.Fatalf("missing class %q", c)
		}
	}
}

func TestMockRosterRoleCounts(t *testing.T) {
	roster := mockRoster()
	tanks := countBy(roster, func(c mockCharacter) bool { return c.Role == roleTank })
	healers := countBy(roster, func(c mockCharacter) bool { return c.Role == roleHealer })
	dps := countBy(roster, func(c mockCharacter) bool { return c.Role == roleDPS })
	if tanks != 1 {
		t.Errorf("tanks = %d, want 1", tanks)
	}
	if healers != 6 {
		t.Errorf("healers = %d, want 6", healers)
	}
	if dps != 17 {
		t.Errorf("dps = %d, want 17", dps)
	}
	if tanks+healers+dps != len(roster) {
		t.Errorf("roles don't sum to roster size: %d+%d+%d != %d", tanks, healers, dps, len(roster))
	}
}

func TestMockRosterOfficersAndUnverified(t *testing.T) {
	roster := mockRoster()
	officers := countBy(roster, func(c mockCharacter) bool { return c.Officer })
	unverified := countBy(roster, func(c mockCharacter) bool { return c.Unverifed })
	if officers != 2 {
		t.Errorf("officers = %d, want 2", officers)
	}
	if unverified != 3 {
		t.Errorf("unverified = %d, want 3", unverified)
	}
}

// Live-fix round, defect 2: this fixture's two Skyborne characters now carry the
// canonical slug races.json actually names ("high-order-skyborne" - the Alliance half of
// the neutral race, since every other race on this roster is Alliance), not the bare
// "skyborne" that factionForRace (api/internal/guilds/bis_lookup.go) could never resolve
// to a faction - which left every Skyborne character's own gear_gap stuck null.
func TestMockRosterSkyborneCount(t *testing.T) {
	skyborne := countBy(mockRoster(), func(c mockCharacter) bool { return c.RaceSlug == "high-order-skyborne" })
	if skyborne != 2 {
		t.Errorf("skyborne = %d, want 2 (one or two is the spec, this fixture picks two)", skyborne)
	}
}

func TestMockRosterConsentDistribution(t *testing.T) {
	roster := mockRoster()
	gearBags := countBy(roster, func(c mockCharacter) bool { return c.Consent == "gear_bags" })
	rosterConsent := countBy(roster, func(c mockCharacter) bool { return c.Consent == "roster" })
	gear := countBy(roster, func(c mockCharacter) bool { return c.Consent == "gear" })
	if gearBags != 1 {
		t.Errorf("gear_bags consent = %d, want 1", gearBags)
	}
	if rosterConsent == 0 {
		t.Errorf("roster consent = 0, want a few")
	}
	if gear <= rosterConsent+gearBags {
		t.Errorf("gear consent (%d) should be most of the roster, got roster=%d gear_bags=%d", gear, rosterConsent, gearBags)
	}
}

func TestMockRosterTwoAccountsShareTwoCharactersEach(t *testing.T) {
	roster := mockRoster()
	byAccount := map[string][]string{}
	for _, c := range roster {
		byAccount[c.accountKey()] = append(byAccount[c.accountKey()], c.Name)
	}
	shared := 0
	for key, names := range byAccount {
		if len(names) < 2 {
			continue
		}
		if len(names) != 2 {
			t.Errorf("account %q has %d characters, want exactly 2", key, len(names))
		}
		shared++
	}
	if shared != 2 {
		t.Fatalf("accounts with a main and an alt = %d, want 2", shared)
	}
	if len(byAccount) != 22 {
		t.Fatalf("distinct accounts = %d, want 22", len(byAccount))
	}
}

func TestMockRosterPairedAccountsShareConsent(t *testing.T) {
	byAccount := map[string][]mockCharacter{}
	for _, c := range mockRoster() {
		byAccount[c.accountKey()] = append(byAccount[c.accountKey()], c)
	}
	for key, chars := range byAccount {
		if len(chars) < 2 {
			continue
		}
		for _, c := range chars[1:] {
			if c.Consent != chars[0].Consent {
				t.Errorf("account %q: consent disagrees between %s (%s) and %s (%s)",
					key, chars[0].Name, chars[0].Consent, c.Name, c.Consent)
			}
		}
	}
}

func TestMockRosterAttendanceWithinEighteenToTwentyTwoPerNight(t *testing.T) {
	roster := mockRoster()
	for _, tag := range []string{"barrow", "hyjal", "ony1", "ony2"} {
		n := countBy(roster, func(c mockCharacter) bool { return c.attends(tag) })
		if n < 18 || n > 22 {
			t.Errorf("%s: %d attendees, want 18-22", tag, n)
		}
	}
}

func TestMockRosterExactlyTwoAttendOnlyOneNight(t *testing.T) {
	tags := []string{"barrow", "hyjal", "ony1", "ony2"}
	oneNightOnly := countBy(mockRoster(), func(c mockCharacter) bool {
		attended := 0
		for _, tag := range tags {
			if c.attends(tag) {
				attended++
			}
		}
		return attended == 1
	})
	if oneNightOnly != 2 {
		t.Fatalf("characters attending exactly one night = %d, want 2", oneNightOnly)
	}
}

func TestMockRosterExactlyOneAbsentFromTheLastWeek(t *testing.T) {
	// ony1 and ony2 are this roster's two last-7-days reports (raid_test.go asserts that
	// independently); a character absent from both has no attendance in the last week.
	absentFromBothRecent := countBy(mockRoster(), func(c mockCharacter) bool {
		return !c.attends("ony1") && !c.attends("ony2")
	})
	if absentFromBothRecent != 1 {
		t.Fatalf("characters absent from both recent nights = %d, want 1", absentFromBothRecent)
	}
}

func TestMockRosterDeathProneCount(t *testing.T) {
	n := countBy(mockRoster(), func(c mockCharacter) bool { return c.DeathProne })
	if n == 0 || n > 4 {
		t.Fatalf("DeathProne count = %d, want a small handful (1-4)", n)
	}
}

func TestMockRosterReadinessVarietyCounts(t *testing.T) {
	roster := mockRoster()
	nonBis := countBy(roster, func(c mockCharacter) bool { return c.NonBisSlots > 0 })
	if nonBis != 4 {
		t.Fatalf("characters with non-BiS slots = %d, want 4", nonBis)
	}
	for _, c := range roster {
		if c.NonBisSlots > 0 && (c.NonBisSlots < 2 || c.NonBisSlots > 4) {
			t.Errorf("%s: NonBisSlots = %d, want 2-4", c.Name, c.NonBisSlots)
		}
	}
	missingEnchant := countBy(roster, func(c mockCharacter) bool { return c.MissingEnchantSlot != "" })
	if missingEnchant != 2 {
		t.Fatalf("characters with a missing enchant = %d, want 2", missingEnchant)
	}
	unspent := countBy(roster, func(c mockCharacter) bool { return c.UnspentTalents })
	if unspent != 2 {
		t.Fatalf("characters with unspent talents = %d, want 2", unspent)
	}
}

func TestMockRosterItemLevelsInRange(t *testing.T) {
	for _, c := range mockRoster() {
		if c.ItemLevel < 55 || c.ItemLevel > 70 {
			t.Errorf("%s: item level %d outside 55-70", c.Name, c.ItemLevel)
		}
	}
}
