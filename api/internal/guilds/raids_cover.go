// api/internal/guilds/raids_cover.go
//
// The buff and debuff cover a raid leader checks before the pull (owner decision D15): for each
// row of the curated catalogue (data/curated/raid-cover.json, Go twin raids_cover_catalogue.go),
// how many signed In, Late and Bench can provide it, against how many the raid wants. Pure: the
// raid view hands it the sign-ups it already built.
package guilds

import (
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/raidspec"
)

//go:generate go run ../../cmd/genraidcover -build 1.60.1.70790

// coverEntry is one catalogue row. Icon is the spell's icon stem in the build
// (/data/{build}/icons/{stem}.webp).
type coverEntry struct {
	Key       string          `json:"key"`
	Name      string          `json:"name"`
	Kind      string          `json:"kind"`
	SpellID   int             `json:"spell_id"`
	Icon      string          `json:"-"`
	Need      coverNeed       `json:"need"`
	Pool      string          `json:"pool"`
	Providers []coverProvider `json:"providers"`
}

// coverNeed is how many providers the raid wants: Count for the raid, or Count for every
// coverGroupSize raiders (In or Late) who fight in melee or at range, never fewer than Count.
type coverNeed struct {
	Count int    `json:"count"`
	Per   string `json:"per"`
}

// The coverNeed.Per values.
const (
	coverPerRaid        = "raid"
	coverPerMeleeGroup  = "melee_group"
	coverPerCasterGroup = "caster_group"
)

// coverGroupSize is a party's size: a party buff reaches its own group of five.
const coverGroupSize = 5

// coverProvider is who can provide a row: a class, optionally only one spec (a specs.json key)
// and a talent by its client name. Spell is the id the provider learns when it is not the row's.
type coverProvider struct {
	Class  string `json:"class"`
	Spec   string `json:"spec"`
	Talent string `json:"talent"`
	Spell  int    `json:"spell"`
}

// The CoverRow.State values: enough In or Late, some but too few, none.
const (
	CoverOK      = "ok"
	CoverShort   = "short"
	CoverMissing = "missing"
)

// CoverRow is one row of RaidView.Cover. ProvidersIn, ProvidersLate and ProvidersBench are the
// players signed so who can provide it; Likely is how many of the In and Late providers are
// counted on a talent their sign-up cannot show (an MM hunter is likely to have Trueshot Aura).
// State reads In plus Late against Need.
type CoverRow struct {
	Key            string `json:"key"`
	Name           string `json:"name"`
	Kind           string `json:"kind"`
	SpellID        int    `json:"spell_id"`
	Icon           string `json:"icon"`
	Need           int    `json:"need"`
	Per            string `json:"per"`
	ProvidersIn    int    `json:"providers_in"`
	ProvidersLate  int    `json:"providers_late"`
	ProvidersBench int    `json:"providers_bench"`
	Likely         int    `json:"likely"`
	State          string `json:"state"`
}

// coverSignup is what the cover reads off a sign-up. Talents nil means not known.
type coverSignup struct {
	Class   string
	SpecKey string
	Group   string
	State   string
	Talents []string
}

// raidCover is the raid view's cover: nil for a dungeon or an event.
func raidCover(zoneKind string, signups []SignupView) []CoverRow {
	if zoneKind != ZoneRaid {
		return nil
	}
	in := make([]coverSignup, len(signups))
	for i, sg := range signups {
		in[i] = coverSignup{Class: sg.Class, Group: sg.Group, State: sg.State}
		if sg.SpecKey != nil {
			in[i].SpecKey = *sg.SpecKey
		}
	}
	return computeCover(raidCoverCatalogue, in)
}

// computeCover is the cover of catalogue for signups. Rows sharing a pool draw on the same
// players in catalogue order: a lone paladin covers the first blessing, not all three.
func computeCover(catalogue []coverEntry, signups []coverSignup) []CoverRow {
	melee, caster := coverAudience(signups)
	pools := map[string]*coverTally{}
	out := make([]CoverRow, 0, len(catalogue))
	for _, e := range catalogue {
		need := e.Need.resolve(melee, caster)
		tally := tallyProviders(e, signups)
		if e.Pool != "" {
			left, seen := pools[e.Pool]
			if !seen {
				fresh := tally
				left = &fresh
				pools[e.Pool] = left
			}
			tally = left.take(need)
		}
		out = append(out, coverRow(e, need, tally))
	}
	return out
}

func coverRow(e coverEntry, need int, t coverTally) CoverRow {
	in, late, bench := t[coverIn].total(), t[coverLate].total(), t[coverBench].total()
	state := CoverMissing
	switch {
	case in+late >= need:
		state = CoverOK
	case in+late > 0:
		state = CoverShort
	}
	return CoverRow{
		Key: e.Key, Name: e.Name, Kind: e.Kind, SpellID: e.SpellID, Icon: e.Icon,
		Need: need, Per: e.Need.Per, ProvidersIn: in, ProvidersLate: late, ProvidersBench: bench,
		Likely: t[coverIn].likely + t[coverLate].likely, State: state,
	}
}

// coverAudience counts the In and Late raiders who fight in melee (tanks too) and at range
// (healers too).
func coverAudience(signups []coverSignup) (melee, caster int) {
	for _, sg := range signups {
		if sg.State != SignupIn && sg.State != SignupLate {
			continue
		}
		switch sg.Group {
		case raidspec.GroupTank, raidspec.GroupMelee:
			melee++
		case raidspec.GroupRanged, raidspec.GroupHealer:
			caster++
		}
	}
	return melee, caster
}

func (n coverNeed) resolve(melee, caster int) int {
	people := 0
	switch n.Per {
	case coverPerMeleeGroup:
		people = melee
	case coverPerCasterGroup:
		people = caster
	default:
		return n.Count
	}
	groups := (people + coverGroupSize - 1) / coverGroupSize
	return n.Count * max(1, groups)
}

// The tally's buckets, in the order a pool fills them.
const (
	coverIn = iota
	coverLate
	coverBench
	coverBuckets
)

var coverBucket = map[string]int{SignupIn: coverIn, SignupLate: coverLate, SignupBench: coverBench}

// coverCount is a bucket's providers: sure, and likely (a talent the sign-up cannot show).
type coverCount struct{ sure, likely int }

func (c coverCount) total() int { return c.sure + c.likely }

type coverTally [coverBuckets]coverCount

func tallyProviders(e coverEntry, signups []coverSignup) coverTally {
	var t coverTally
	for _, sg := range signups {
		b, counted := coverBucket[sg.State]
		if !counted {
			continue
		}
		switch provides(e, sg) {
		case provideSure:
			t[b].sure++
		case provideLikely:
			t[b].likely++
		}
	}
	return t
}

// take hands a pooled row up to need providers, In first, then Late, then Bench, sure before
// likely, and leaves the rest for the pool's next row.
func (t *coverTally) take(need int) coverTally {
	var got coverTally
	for b := range t {
		sure := min(t[b].sure, need)
		t[b].sure -= sure
		need -= sure
		likely := min(t[b].likely, need)
		t[b].likely -= likely
		need -= likely
		got[b] = coverCount{sure: sure, likely: likely}
	}
	return got
}

type provision int

const (
	provideNone provision = iota
	provideLikely
	provideSure
)

// provides is the best of e's providers for sg.
func provides(e coverEntry, sg coverSignup) provision {
	best := provideNone
	for _, p := range e.Providers {
		best = max(best, p.matches(sg))
	}
	return best
}

// matches: the class, then the spec when the provider names one (an unknown spec is not
// counted), then the talent, likely when the sign-up's talents are unknown.
func (p coverProvider) matches(sg coverSignup) provision {
	if !strings.EqualFold(p.Class, sg.Class) || (p.Spec != "" && p.Spec != sg.SpecKey) {
		return provideNone
	}
	if p.Talent == "" {
		return provideSure
	}
	if sg.Talents == nil {
		return provideLikely
	}
	for _, t := range sg.Talents {
		if t == p.Talent {
			return provideSure
		}
	}
	return provideNone
}
