package coach

import (
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

const (
	ranksUnitRank  = "rank"
	ranksUnitCasts = "casts"
	ranksPercent   = 100.0
	ranksCannotKey = "cannot_level"
)

// lowRankUse is one ability family's casts below the highest rank trainable at the
// player's level.
type lowRankUse struct {
	family   string
	top      TrainableRank // the highest rank trainable at Row.Level
	byRank   map[int64]int // spell id -> casts, for ranks below top
	casts    int           // casts of every rank of the family
	lowCasts int           // casts below top
	worst    TrainableRank // the below-top rank cast most
}

// ranks is F5 (§6.2): a lower rank cast where a higher one is trainable (damage and
// tanks), or a heal rank whose overheal runs well over the top logs' for the same rank.
func ranks(in Input, _ Counted) outcome {
	if in.Row.Role == RoleHealer {
		return healRanks(in)
	}
	return lowRanks(in)
}

func lowRanks(in Input) outcome {
	if in.Row.Level <= 0 {
		return outcome{CouldNotCheck: []string{Drafts[ranksCannotKey]}}
	}
	var worst *lowRankUse
	for _, u := range lowRankUses(in) {
		fires := u.lowCasts >= RanksFiresCasts ||
			float64(u.lowCasts)/float64(u.casts)*ranksPercent >= RanksFiresPct
		if fires && (worst == nil || u.lowCasts > worst.lowCasts) {
			worst = u
		}
	}
	if worst == nil {
		return outcome{}
	}
	f := lowRankFinding(in, *worst)
	f.Anchors = []string{"#0"}
	return outcome{Finding: f, Anchors: []Anchor{{Kind: AnchorCast, SpellID: worst.worst.ID}}}
}

// lowRankUses groups the player's casts by family, in family-name order.
func lowRankUses(in Input) []*lowRankUse {
	uses := map[string]*lowRankUse{}
	for _, row := range in.Fight.Casts {
		if row.OwnerGUID != in.Row.GUID || row.Succeeded <= 0 {
			continue
		}
		family, ok := in.Tables.Rotation.Family(row.SpellID)
		if !ok || downrankListed(in.Tables.Rules, family) {
			continue
		}
		top, cast, ok := trainableRankOf(in, family, row.SpellID)
		if !ok {
			continue
		}
		u := uses[family]
		if u == nil {
			u = &lowRankUse{family: family, top: top, byRank: map[int64]int{}}
			uses[family] = u
		}
		u.casts += int(row.Succeeded)
		if cast.Rank < top.Rank {
			u.lowCasts += int(row.Succeeded)
			u.byRank[cast.ID] += int(row.Succeeded)
		}
	}
	out := make([]*lowRankUse, 0, len(uses))
	for _, u := range uses {
		u.worst = mostCastRank(in, *u)
		out = append(out, u)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].family < out[j].family })
	return out
}

func downrankListed(r SpecRules, family string) bool {
	return slices.ContainsFunc(r.Downranks, func(e RuleEntry) bool { return e.Family == family })
}

// trainableRankOf is the highest rank of family trainable at the player's level and the
// trainable entry of spellID; ok is false for a family with no ranks to choose between.
func trainableRankOf(in Input, family string, spellID int64) (top, cast TrainableRank, ok bool) {
	var found bool
	for _, r := range in.Tables.Trainable[family] {
		if r.Rank <= 0 {
			return TrainableRank{}, TrainableRank{}, false
		}
		if int64(r.Level) <= in.Row.Level && r.Rank > top.Rank {
			top = r
		}
		if r.ID == spellID {
			cast, found = r, true
		}
	}
	return top, cast, found && top.Rank > 0
}

// mostCastRank is the below-top rank with the most casts (the lowest rank on a tie).
func mostCastRank(in Input, u lowRankUse) TrainableRank {
	var best TrainableRank
	for _, r := range in.Tables.Trainable[u.family] {
		if _, ok := u.byRank[r.ID]; ok && (best.ID == 0 || u.byRank[r.ID] > u.byRank[best.ID]) {
			best = r
		}
	}
	return best
}

func lowRankFinding(in Input, u lowRankUse) *Finding {
	n := u.byRank[u.worst.ID]
	level := strconv.FormatInt(in.Row.Level, 10)
	fields := map[string]string{
		"Spell": u.family, "n": strconv.Itoa(n),
		"r": strconv.Itoa(u.worst.Rank), "next": strconv.Itoa(u.top.Rank),
	}
	return &Finding{
		Code: CodeRanks, Severity: SeverityMinor,
		Title:  Fill(TitleLowRank, fields),
		Detail: Fill(DetailLowRank, fields),
		Basis: Basis{Kind: BasisSpellData, Label: Fill(BasisRanks, map[string]string{
			"level": level, "class": strings.ToLower(in.Row.Class)})},
		Facts: Facts{
			Numbers: []NumberFact{
				{V: float64(n), Text: fields["n"], Unit: ranksUnitCasts, Of: u.family + " casts of the lower rank"},
				{V: float64(u.worst.Rank), Text: fields["r"], Unit: ranksUnitRank, Of: "rank cast"},
				{V: float64(u.top.Rank), Text: fields["next"], Unit: ranksUnitRank, Of: "highest rank trainable at your level"},
				{V: float64(in.Row.Level), Text: level, Unit: "level", Of: "your level"},
			},
			Spells: []SpellFact{
				{ID: u.worst.ID, Name: u.family, Rank: u.worst.Rank},
				{ID: u.top.ID, Name: u.family, Rank: u.top.Rank},
			},
		},
	}
}

// healRead is one heal rank's overheal against the top logs' for the same rank.
type healRead struct {
	ability summary.Ability
	rank    int
	name    string
	casts   int
	pct     float64
	ref     float64
	basis   Basis
}

func (r healRead) gap() float64 { return r.pct - r.ref }

func healRanks(in Input) outcome {
	var o outcome
	var worst *healRead
	for _, a := range playerHealing(in) {
		r, ok := readHealRank(in, a, &o)
		if ok && r.gap() >= HealRankFiresPoints && (worst == nil || r.gap() > worst.gap()) {
			worst = &r
		}
	}
	if worst == nil {
		return o
	}
	o.Anchors = append(o.Anchors, Anchor{Kind: AnchorHealing, SpellID: worst.ability.SpellID})
	o.Finding = healRankFinding(*worst)
	o.Finding.Anchors = []string{"#" + strconv.Itoa(len(o.Anchors)-1)}
	return o
}

// playerHealing is the player's own healing abilities, spell id order.
func playerHealing(in Input) []summary.Ability {
	var out []summary.Ability
	for _, actor := range in.Fight.Healing {
		if actor.GUID == in.Row.GUID {
			out = append(out, actor.Abilities...)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].SpellID < out[j].SpellID })
	return out
}

// readHealRank measures one ability; ok is false when it is no ranked heal cast often
// enough, or when no top-logs median answers (a thin pool adds its line to o).
func readHealRank(in Input, a summary.Ability, o *outcome) (healRead, bool) {
	sc := in.Tables.Spells[a.SpellID]
	casts := succeededCasts(in, a.SpellID)
	if a.Via != "" || a.Total <= 0 || sc.Rank <= 0 || casts < HealRankMinCasts {
		return healRead{}, false
	}
	ref, basis, thin := comparison(in, func(p Profile) (float64, bool) { return p.OverhealOf(a.SpellID) }, nil, false)
	if thin != "" && !slices.Contains(o.CouldNotCheck, thin) {
		o.CouldNotCheck = append(o.CouldNotCheck, thin)
	}
	if basis == nil {
		return healRead{}, false
	}
	name := sc.Name
	if name == "" {
		name = a.Name
	}
	pct := float64(a.Overheal) / float64(a.Total) * ranksPercent
	return healRead{ability: a, rank: sc.Rank, name: name, casts: casts, pct: pct, ref: ref, basis: *basis}, true
}

func succeededCasts(in Input, spellID int64) int {
	var n int64
	for _, row := range in.Fight.Casts {
		if row.OwnerGUID == in.Row.GUID && row.SpellID == spellID {
			n += row.Succeeded
		}
	}
	return int(n)
}

func healRankFinding(r healRead) *Finding {
	fields := map[string]string{
		"Spell": r.name, "r": strconv.Itoa(r.rank), "pct": PctNumber(r.pct), "n": strconv.Itoa(r.casts),
	}
	sev := SeverityMinor
	if r.gap() >= HealRankMajorPoints {
		sev = SeverityMajor
	}
	return &Finding{
		Code: CodeRanks, Severity: sev,
		Title:  Fill(TitleHealRank, fields),
		Detail: Fill(DetailHealRank, fields) + " " + LowerRanks,
		Basis: Basis{Kind: r.basis.Kind, N: r.basis.N,
			Label: r.basis.Label + Fill(BasisSameRank, map[string]string{"pct": PctNumber(r.ref)})},
		Facts: Facts{
			Numbers: []NumberFact{
				{V: r.pct, Text: PctNumber(r.pct), Unit: "%", Of: r.name + " overheal"},
				{V: r.ref, Text: PctNumber(r.ref), Unit: "%", Of: r.name + " overheal in the comparison"},
				{V: float64(r.casts), Text: fields["n"], Unit: ranksUnitCasts, Of: r.name + " casts"},
				{V: float64(r.rank), Text: fields["r"], Unit: ranksUnitRank, Of: "heal rank"},
				{V: float64(r.basis.N), Text: strconv.Itoa(r.basis.N), Unit: "players", Of: "top logs"},
			},
			Spells: []SpellFact{{ID: r.ability.SpellID, Name: r.name, Rank: r.rank}},
		},
	}
}
