package guilds

import (
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/raidspec"
)

// A small catalogue: one raid buff, a melee party buff, a talent-gated spec buff and a pool.
var testCover = []coverEntry{
	{Key: "fort", Name: "Power Word: Fortitude", Kind: "buff", SpellID: 10938, Icon: "spell_holy_wordfortitude",
		Need: coverNeed{Count: 1, Per: coverPerRaid}, Providers: []coverProvider{{Class: "priest"}}},
	{Key: "wf", Name: "Windfury Totem", Kind: "buff", SpellID: 10614, Icon: "spell_nature_windfury",
		Need: coverNeed{Count: 1, Per: coverPerMeleeGroup}, Providers: []coverProvider{{Class: "shaman"}}},
	{Key: "tsa", Name: "Trueshot Aura", Kind: "buff", SpellID: 20906, Icon: "ability_trueshot",
		Need:      coverNeed{Count: 1, Per: coverPerRaid},
		Providers: []coverProvider{{Class: "hunter", Spec: "hunter-marksmanship", Talent: "Trueshot Aura"}}},
	{Key: "kings", Name: "Blessing of Kings", Kind: "buff", SpellID: 20217, Icon: "spell_magic_magearmor",
		Need: coverNeed{Count: 1, Per: coverPerRaid}, Pool: "bless", Providers: []coverProvider{{Class: "paladin"}}},
	{Key: "might", Name: "Blessing of Might", Kind: "buff", SpellID: 19838, Icon: "spell_holy_fistofjustice",
		Need: coverNeed{Count: 1, Per: coverPerRaid}, Pool: "bless", Providers: []coverProvider{{Class: "paladin"}}},
}

func signup(class, spec, group, state string) coverSignup {
	return coverSignup{Class: class, SpecKey: spec, Group: group, State: state}
}

func melee(n int, state string) []coverSignup {
	out := make([]coverSignup, n)
	for i := range out {
		out[i] = signup("rogue", "rogue-combat", raidspec.GroupMelee, state)
	}
	return out
}

type coverWant struct {
	need, in, late, bench, likely int
	state                         string
}

func TestComputeCover(t *testing.T) {
	cases := []struct {
		name    string
		signups []coverSignup
		want    map[string]coverWant
	}{
		{"nobody signed", nil, map[string]coverWant{
			"fort": {need: 1, state: CoverMissing}, "wf": {need: 1, state: CoverMissing},
			"kings": {need: 1, state: CoverMissing},
		}},
		{"a priest In covers fort; one on the bench counts apart", []coverSignup{
			signup("priest", "priest-holy", raidspec.GroupHealer, SignupIn),
			signup("Priest", "priest-shadow", raidspec.GroupRanged, SignupBench),
			signup("priest", "priest-holy", raidspec.GroupHealer, SignupOut),
		}, map[string]coverWant{"fort": {need: 1, in: 1, bench: 1, state: CoverOK}}},
		{"late counts toward the need", []coverSignup{
			signup("priest", "priest-holy", raidspec.GroupHealer, SignupLate),
		}, map[string]coverWant{"fort": {need: 1, late: 1, state: CoverOK}}},
		{"bench alone is missing", []coverSignup{
			signup("priest", "priest-holy", raidspec.GroupHealer, SignupBench),
		}, map[string]coverWant{"fort": {need: 1, bench: 1, state: CoverMissing}}},
		{"eleven melee want three windfury; one shaman is short", append(melee(10, SignupIn),
			signup("warrior", "warrior-protection", raidspec.GroupTank, SignupLate),
			signup("shaman", "shaman-restoration", raidspec.GroupHealer, SignupIn),
		), map[string]coverWant{"wf": {need: 3, in: 1, state: CoverShort}}},
		{"an MM hunter is likely, a BM hunter is not, an unknown spec is not", []coverSignup{
			signup("hunter", "hunter-marksmanship", raidspec.GroupRanged, SignupIn),
			signup("hunter", "hunter-beast-mastery", raidspec.GroupRanged, SignupIn),
			signup("hunter", "", raidspec.GroupRanged, SignupIn),
		}, map[string]coverWant{"tsa": {need: 1, in: 1, likely: 1, state: CoverOK}}},
		{"one paladin gives one blessing, the first", []coverSignup{
			signup("paladin", "paladin-holy", raidspec.GroupHealer, SignupIn),
		}, map[string]coverWant{
			"kings": {need: 1, in: 1, state: CoverOK}, "might": {need: 1, state: CoverMissing},
		}},
		{"a pool fills In, then Late, then Bench", []coverSignup{
			signup("paladin", "paladin-holy", raidspec.GroupHealer, SignupLate),
			signup("paladin", "paladin-holy", raidspec.GroupHealer, SignupBench),
			signup("paladin", "paladin-holy", raidspec.GroupHealer, SignupBench),
		}, map[string]coverWant{
			"kings": {need: 1, late: 1, state: CoverOK}, "might": {need: 1, bench: 1, state: CoverMissing},
		}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			rows := computeCover(testCover, c.signups)
			if len(rows) != len(testCover) {
				t.Fatalf("%d rows", len(rows))
			}
			for _, r := range rows {
				w, ok := c.want[r.Key]
				if !ok {
					continue
				}
				got := coverWant{r.Need, r.ProvidersIn, r.ProvidersLate, r.ProvidersBench, r.Likely, r.State}
				if got != w {
					t.Errorf("%s = %+v, want %+v", r.Key, got, w)
				}
			}
		})
	}
}

func TestCoverTalentsWhenKnown(t *testing.T) {
	with := signup("hunter", "hunter-marksmanship", raidspec.GroupRanged, SignupIn)
	with.Talents = []string{"Trueshot Aura"}
	without := signup("hunter", "hunter-marksmanship", raidspec.GroupRanged, SignupIn)
	without.Talents = []string{}
	rows := computeCover(testCover, []coverSignup{with, without})
	if r := rows[2]; r.ProvidersIn != 1 || r.Likely != 0 || r.State != CoverOK {
		t.Fatalf("tsa = %+v", r)
	}
}

func TestRaidCoverIsForRaidsOnly(t *testing.T) {
	if raidCover(ZoneDungeon, nil) != nil || raidCover(ZoneEvent, nil) != nil {
		t.Fatal("a dungeon or event has no cover")
	}
	key := "priest-holy"
	rows := raidCover(ZoneRaid, []SignupView{{Class: "priest", SpecKey: &key, Group: raidspec.GroupHealer, State: SignupIn}})
	if len(rows) != len(raidCoverCatalogue) {
		t.Fatalf("%d rows", len(rows))
	}
	for _, r := range rows {
		if r.Key == "fortitude" && (r.ProvidersIn != 1 || r.State != CoverOK || r.Icon == "") {
			t.Fatalf("fortitude = %+v", r)
		}
	}
}
