// api/internal/guilds/me_checklist_prep.go
//
// The checklist's flask-and-food and fire resistance items (D14), read from the export's
// supplies= and school= sections through raid_prep.go.
package guilds

// flaskFoodCheck is ok when the bags hold the raid leader's minimum (a flask or a raid
// elixir, and buff food), else a fail naming what is missing. Not tracked for a leveling
// character, without the bags share, or for an export that predates supplies=.
func flaskFoodCheck(in checklistInput) ChecklistItem {
	it := ChecklistItem{Key: CheckFlaskFood}
	switch {
	case in.readiness.Leveling:
		it.State, it.Reason = CheckNotTracked, reasonLeveling
		return it
	case in.row.Consent != "gear_bags":
		it.State, it.Reason = CheckNotTracked, reasonBagsNotShared
		return it
	}
	r, ok := suppliesOf(in.row)
	if !ok {
		it.State, it.Reason = CheckNotTracked, reasonNeedsConsumables
		return it
	}
	it.Supplies, it.State = r.Stacks, CheckOK
	if !r.ready() {
		it.State, it.Missing = CheckFail, r.Missing
	}
	return it
}

// fireResistanceCheck reports the fire resistance total, always ok when read: no floor is
// set, since no Forever raid on the schedule asks every raider for a number, and the raid
// leader reads the tanks' totals on the Readiness board.
func fireResistanceCheck(row RosterRow) ChecklistItem {
	it := ChecklistItem{Key: CheckFireResistance}
	switch r := resistancesOf(row); {
	case !hasGearConsent(row.Consent):
		it.State, it.Reason = CheckNotTracked, reasonGearNotShared
	case r == nil:
		it.State, it.Reason = CheckNotTracked, reasonNeedsResistance
	default:
		fire := r.Fire
		it.State, it.Resistance = CheckOK, &fire
	}
	return it
}
