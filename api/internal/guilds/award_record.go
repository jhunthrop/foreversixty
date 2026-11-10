// api/internal/guilds/award_record.go
//
// An award's decision record as the loot council kept it (migration 0073): main or off spec, the
// officer's reason, the yes and no votes and how many officers could vote. Every field is null on
// an award made before the council kept records. The night report's loot and the member profile's
// loot both carry it, so an officer can say why a drop went where it did.
package guilds

// AwardRecord is embedded in each award a read returns.
type AwardRecord struct {
	MsOs          *string `json:"ms_os"`
	Reason        *string `json:"reason"`
	VotesYes      *int    `json:"votes_yes"`
	VotesNo       *int    `json:"votes_no"`
	VotesOfficers *int    `json:"votes_officers"`
}

// awardRecordColumns are the loot_awards columns AwardRecord scans, in scanTargets' order; a query
// that names its loot_awards table `la` selects them with this.
const awardRecordColumns = `la.ms_os, la.reason, la.votes_yes, la.votes_no, la.votes_officers`

// scanTargets are AwardRecord's fields for rows.Scan, in awardRecordColumns' order.
func (r *AwardRecord) scanTargets() []any {
	return []any{&r.MsOs, &r.Reason, &r.VotesYes, &r.VotesNo, &r.VotesOfficers}
}
