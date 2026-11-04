package request

import (
	"fmt"
	"slices"
	"sort"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
	"github.com/jhunthrop/foreversixty/sim/internal/enginetalents"
	"github.com/jhunthrop/foreversixty/sim/internal/spellranks"
	"github.com/jhunthrop/foreversixty/sim/leveling"
	"github.com/jhunthrop/foreversixty/sim/specs"
	engine "github.com/wowsims/classic/sim"
	"github.com/wowsims/classic/sim/core/proto"
	"path/filepath"
)

// rankRegistrationLevels are the band floors and tops a character is built
// at: every rank learned by one of these levels is checked at the first of
// them it is learned by.
var rankRegistrationLevels = []int{10, 19, 20, 29, 30, 39, 40, 49, 50, 59, 60}

// Round 19 (hunter lane): the engine registered only Serpent Sting's top
// rank, so a downranked line resolved to an id the engine did not know and
// dropped out of the rotation with no error ("sting downranks lose" in round
// 17 was the sting no longer being cast). Every rank of every ranked spell a
// curated rotation or a rank-choice row casts below its top rank
// (downrankedFamilies) must be a registered engine spell at every band that
// has learned it.
//
// For each written spec and each level (band floors and tops), the guide build's ladder talents at
// that level (TestRotationLadder's character) cast, through one probe line
// each, the id sim/internal/spellranks.SpellIDOfRank resolves for every
// non-book rank learned by that level. The engine's own validation
// (warnedActions, as the smoke test reads it) names every probe it cannot
// resolve. A family the talents have not taught yet at that level is
// skipped (its talent spell ids, as the ladder's rule 1 reads them).
func TestEveryLearnedRankOfACastSpellIsRegistered(t *testing.T) {
	registerEngine.Do(engine.RegisterAll)
	build := activeBuild(t)
	ranks, err := loadSpellRanks(repoRoot, build)
	if err != nil {
		t.Fatal(err)
	}
	all, err := embeddedRankChoices()
	if err != nil {
		t.Fatal(err)
	}
	var missing []string
	for _, rot := range writtenRotations(t) {
		spec := specs.ByKey[rot.spec]
		missing = append(missing, unregisteredRanks(t, build, spec, ranks, all[rot.spec])...)
	}
	sort.Strings(missing)
	for _, m := range missing {
		if reason, known := knownUnregisteredRanks[m]; known {
			t.Logf("known: %s (%s)", m, reason)
			continue
		}
		t.Error(m)
	}
}

// knownUnregisteredRanks are findings reported to the class lanes (round 19),
// recorded so the guard is green while they are fixed in the engine, and
// deleted when they are, the way smokeBuildWarnings carries its FINDINGs.
var knownUnregisteredRanks = map[string]string{
	"druid: druid-feral-bear Primal Bite rank 1 (id 407995) is not a registered engine spell at levels [50 59 60]": "FINDING: the engine registers only the highest Primal Bite rank learned; the bear rotation casts rank 3, which resolves to rank 2 at 50-59, so nothing is cast at rank 1 there today. Every rank costs the same 20 rage on the same cooldown, so a lower rank never wins (round 19 druid)",
	"druid: druid-feral-bear Primal Bite rank 2 (id 1238069) is not a registered engine spell at levels [60]":      "FINDING: as rank 1; at 60 the rotation casts rank 3",
	"druid: druid-feral-bear Primal Bite rank 4 (id 1238073) is not a registered engine spell at levels [60]":      "FINDING, not a bug (round 19 druid): 1238073 has no SkillLineAbility row, so no trainer, talent or item teaches it, and trainables/druid.json omits it; its client rows equal rank 3 (77 plus 100% weapon damage, 6 s cooldown, same category and class mask). spellranks.json lists it by its 'Rank 4' subtext only",
}

// probeFamily is one ranked spell a spec casts: its name and an id of it.
type probeFamily struct {
	name     string
	authored int32
	tiers    []rankTier
}

// downrankedFamilies is every ranked family the spec casts below its top
// rank: one a rank-choice row names, or one the curated rotation casts at an
// authored rank under the family's highest. The engine registers a spell's
// highest rank learned at the character's level for most spells, so only a
// family cast at a lower rank needs its other ranks registered.
func downrankedFamilies(t *testing.T, spec string, abilities classAbilities, choices rankChoices) []probeFamily {
	t.Helper()
	curated, err := loadLadderCurated(repoRoot, spec)
	if err != nil {
		t.Fatal(err)
	}
	casts, err := authoredCastSpellIDs(curated.Rotation)
	if err != nil {
		t.Fatal(err)
	}
	var out []probeFamily
	for name, tiers := range abilities.Tiers {
		top := tiers[len(tiers)-1].Rank
		var authored int32
		downranked := false
		for _, tier := range tiers {
			for _, id := range tier.IDs {
				_, chosen := choices[int32(id)]
				if !casts[id] && !chosen {
					continue
				}
				if authored == 0 {
					authored = int32(id)
				}
				downranked = downranked || chosen || tier.Rank < top
			}
		}
		if downranked {
			out = append(out, probeFamily{name: name, authored: authored, tiers: tiers})
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].name < out[j].name })
	return out
}

func unregisteredRanks(t *testing.T, build string, spec specs.Spec, ranks spellRanksFile, choices rankChoices) []string {
	t.Helper()
	class := spec.ClassSlug
	abilities := buildClassAbilities(ranks, class)
	families := downrankedFamilies(t, spec.Spec, abilities, choices)
	talentIDs, err := loadTalentSpellIDs(repoRoot, build, class)
	if err != nil {
		t.Fatal(err)
	}
	talentIDs = expandTalentSpellIDs(talentIDs, abilities)

	guideBuild, treeDigits, err := leveling.GuideBuildTalents(repoRoot, class, spec.SpecSlug)
	if err != nil {
		t.Fatal(err)
	}
	guideTrees, err := leveling.LoadTalentTrees(repoRoot, guideBuild, class)
	if err != nil {
		t.Fatal(err)
	}
	activeTrees, err := leveling.LoadTalentTrees(repoRoot, build, class)
	if err != nil {
		t.Fatal(err)
	}
	targets := leveling.GuideTalentTargets(guideTrees, treeDigits)
	engineDir, err := enginetalents.SourceDir(filepath.Join(repoRoot, "sim"))
	if err != nil {
		t.Fatal(err)
	}
	engineLayout, err := enginetalents.ForClass(engineDir, class)
	if err != nil {
		t.Fatal(err)
	}

	failedAt := map[int32][]int{}
	labels := map[int32]string{}
	for _, level := range rankRegistrationLevels {
		talents := leveling.LadderTalentString(activeTrees, targets, spec.TreeIndex, level, leveling.ForSpec(spec.Spec))
		engineTalents, err := engineLayout.Reposition(activeTrees, talents)
		if err != nil {
			t.Fatalf("%s level %d: %v", spec.Spec, level, err)
		}
		points, err := leveling.TalentRanksFromString(activeTrees, talents)
		if err != nil {
			t.Fatalf("%s level %d: %v", spec.Spec, level, err)
		}
		probes, names := rankProbes(class, level, families, talentIDs, points)
		if len(probes) == 0 {
			continue
		}
		req := api.SimRequest{
			EngineVersion: enginever.Version,
			Spec:          spec.Spec,
			Source:        api.CharacterSource{Kind: api.Sources[0]},
			Character: api.CharacterSpec{
				Name: spec.Spec, Race: smokeRace(t, class), Class: class, Level: level, Talents: engineTalents,
			},
			Encounter:  api.DefaultEncounter(),
			Iterations: smokeIterations,
			RandomSeed: smokeSeed,
		}
		built, err := BuildWith(req, Options{OpenIterations: true})
		if err != nil {
			t.Fatalf("%s level %d: building the request: %v", spec.Spec, level, err)
		}
		attachSmokeHealProfile(t, spec.Spec, built)
		built.Raid.Parties[0].Players[0].Rotation = probeRotation(probes)
		warned := warnedActions(t, built)
		for _, id := range probes {
			labels[id] = names[id]
			if slices.Contains(warned, spellAction(int(id))) {
				failedAt[id] = append(failedAt[id], level)
			}
		}
	}
	out := make([]string, 0, len(failedAt))
	for id, levels := range failedAt {
		out = append(out, fmt.Sprintf("%s: %s %s (id %d) is not a registered engine spell at levels %v",
			class, spec.Spec, labels[id], id, levels))
	}
	return out
}

// rankProbes is the id of every non-book rank of families learned by level,
// with its "Name rank N" label. A family whose ranks are
// talent spells is probed only once the ladder talents at level hold the
// talent.
func rankProbes(class string, level int, families []probeFamily, talentIDs map[int]int, points map[int]int) ([]int32, map[int32]string) {
	var probes []int32
	names := map[int32]string{}
	for _, f := range families {
		if node, isTalent := talentIDs[int(f.authored)]; isTalent && points[node] == 0 {
			continue
		}
		for _, tier := range f.tiers {
			if tier.Level > level {
				continue
			}
			id, ok := spellranks.SpellIDOfRank(class, f.authored, tier.Rank, level)
			if !ok {
				continue
			}
			probes = append(probes, id)
			names[id] = fmt.Sprintf("%s rank %d", f.name, tier.Rank)
		}
	}
	return probes, names
}

// probeRotation casts each id from its own priority line.
func probeRotation(ids []int32) *proto.APLRotation {
	rot := &proto.APLRotation{Type: proto.APLRotation_TypeAPL}
	for _, id := range ids {
		rot.PriorityList = append(rot.PriorityList, &proto.APLListItem{
			Action: &proto.APLAction{Action: &proto.APLAction_CastSpell{
				CastSpell: &proto.APLActionCastSpell{SpellId: &proto.ActionID{RawId: &proto.ActionID_SpellId{SpellId: id}}},
			}},
		})
	}
	return rot
}
