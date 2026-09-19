package request

import (
	"errors"
	"os"
	"slices"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/internal/strcase"
	"github.com/wowsims/classic/sim/core/proto"
)

// Everything the vocabulary lists must resolve, or the settings bar
// built from it sends ids that fail at the boundary.
func TestKnownBuffsAllResolve(t *testing.T) {
	ids := KnownBuffs()
	if len(ids) == 0 {
		t.Fatal("KnownBuffs is empty")
	}
	if !slices.IsSorted(ids) {
		t.Error("KnownBuffs is not sorted; the list is an interface and must be stable")
	}
	for _, id := range ids {
		if _, err := buffsFor([]string{id}); err != nil {
			t.Errorf("KnownBuffs lists %q, which does not resolve: %v", id, err)
		}
	}
	// One from each of the four messages, so a message dropped from the
	// walk is a failure rather than a quietly shorter list.
	for _, want := range []string{"sunder_armor", "blessing_of_kings", "mana_tide_totems", "battle_shout"} {
		if !slices.Contains(ids, want) {
			t.Errorf("KnownBuffs does not list %q", want)
		}
	}
}

func TestKnownConsumablesAllResolve(t *testing.T) {
	ids := KnownConsumables()
	if len(ids) == 0 {
		t.Fatal("KnownConsumables is empty")
	}
	if !slices.IsSorted(ids) {
		t.Error("KnownConsumables is not sorted")
	}
	for _, id := range ids {
		if _, err := consumes([]string{id}, nil); err != nil {
			t.Errorf("KnownConsumables lists %q, which does not resolve: %v", id, err)
		}
	}
	for _, want := range []string{"elixir_of_the_mongoose", "flask_of_the_titans", "main_hand_imbue:shadow_oil", "off_hand_imbue:shadow_oil"} {
		if !slices.Contains(ids, want) {
			t.Errorf("KnownConsumables does not list %q", want)
		}
	}
	// A bare ambiguous id is not listed, because it does not resolve.
	if slices.Contains(ids, "shadow_oil") {
		t.Error("KnownConsumables lists the bare imbue id, which is ambiguous")
	}
}

func TestKnownProfessionsAllResolve(t *testing.T) {
	ids := KnownProfessions()
	if len(ids) == 0 {
		t.Fatal("KnownProfessions is empty")
	}
	if !slices.IsSorted(ids) {
		t.Error("KnownProfessions is not sorted; the list is an interface and must be stable")
	}
	for _, id := range ids {
		if _, ok := ParseProfession(id); !ok {
			t.Errorf("KnownProfessions lists %q, which does not resolve", id)
		}
		// And it resolves through the boundary too, not only through the
		// parser: a slug the map knows but Build refuses is still a trap.
		if _, _, err := professionsFor([]string{id}); err != nil {
			t.Errorf("KnownProfessions lists %q, which Build refuses: %v", id, err)
		}
	}
	if _, ok := ParseProfession("cooking"); ok {
		t.Error("an unlisted profession slug resolved")
	}
}

// The slugs are the site's own interface, so they are a map rather than
// a descriptor walk - but the map must still name every profession the
// engine has, or a character could carry one the sim cannot be told
// about, and must name no value the engine dropped.
func TestTheProfessionMapMatchesTheEngineEnum(t *testing.T) {
	values := proto.Profession(0).Descriptor().Values()
	fromEngine := map[string]bool{}
	for i := 0; i < values.Len(); i++ {
		v := values.Get(i)
		if v.Number() == 0 { // ProfessionUnknown is "none", not a profession
			continue
		}
		fromEngine[strcase.Snake(string(v.Name()))] = true
	}
	for slug := range fromEngine {
		if _, ok := ParseProfession(slug); !ok {
			t.Errorf("the engine has profession %q and the slug map does not; add it and regenerate IDS.md", slug)
		}
	}
	for _, slug := range KnownProfessions() {
		if !fromEngine[slug] {
			t.Errorf("the slug map has profession %q the engine's enum does not", slug)
		}
	}
	if len(fromEngine) != len(KnownProfessions()) {
		t.Errorf("the engine has %d professions, the slug map %d", len(fromEngine), len(KnownProfessions()))
	}
}

// The resolver accepts nothing the vocabulary does not list: an id that
// works but is undocumented is a trap for the web lane.
func TestTheVocabularyAndTheResolverAgree(t *testing.T) {
	buffs := KnownBuffs()
	for _, e := range buffVocabulary() {
		if !slices.Contains(buffs, e.id) {
			t.Errorf("the buff walk resolves %q, which KnownBuffs does not list", e.id)
		}
	}
	consumables := KnownConsumables()
	for _, e := range consumeVocabulary() {
		if !slices.Contains(consumables, e.id) {
			t.Errorf("the consumable walk resolves %q, which KnownConsumables does not list", e.id)
		}
	}
	professions := KnownProfessions()
	for _, e := range professionVocabulary() {
		if !slices.Contains(professions, e.id) {
			t.Errorf("the profession walk resolves %q, which KnownProfessions does not list", e.id)
		}
	}
	// And an id off the list still fails.
	if _, err := buffsFor([]string{"blessing_of_vulpera"}); !errors.Is(err, ErrUnknownBuff) {
		t.Errorf("an unlisted buff id resolved: %v", err)
	}
	if _, _, err := professionsFor([]string{"cooking"}); !errors.Is(err, ErrUnknownProfession) {
		t.Errorf("an unlisted profession slug resolved: %v", err)
	}
}

// The descriptor walk must give the same answer every time, or two
// identical requests would build two different protobufs.
func TestTheWalkIsDeterministic(t *testing.T) {
	for i := 0; i < 20; i++ {
		if !slices.Equal(KnownBuffs(), KnownBuffs()) {
			t.Fatal("KnownBuffs differs between calls")
		}
		if !slices.Equal(KnownConsumables(), KnownConsumables()) {
			t.Fatal("KnownConsumables differs between calls")
		}
		// The profession list is built by ranging a map, which Go
		// deliberately randomises, so the sort is what makes it an
		// interface rather than a coin flip.
		if !slices.Equal(KnownProfessions(), KnownProfessions()) {
			t.Fatal("KnownProfessions differs between calls")
		}
	}
}

// IDS.md is generated, so the committed file and the generator cannot
// drift: a flask the engine gains shows up here rather than in a
// settings bar nobody can build.
func TestIDsMarkdownIsCommitted(t *testing.T) {
	want, err := os.ReadFile("IDS.md")
	if err != nil {
		t.Fatalf("%v (run `go run ./internal/genids` from sim/)", err)
	}
	if got := IDsMarkdown(); got != string(want) {
		t.Error("IDS.md is out of date; run `go run ./internal/genids` from sim/ and commit the result")
	}
}

// A graded buff is the engine's TristateEffect: value one is the plain
// version and value two the talented one. The settings bar needs both,
// so every tristate field answers to "<id>" and "<id>:improved".
func TestGradedBuffIds(t *testing.T) {
	got, err := buffsFor([]string{"battle_shout:improved", "power_word_fortitude"})
	if err != nil {
		t.Fatal(err)
	}
	if got.Raid.BattleShout != proto.TristateEffect_TristateEffectImproved {
		t.Errorf("battle_shout:improved set %v, want the improved version", got.Raid.BattleShout)
	}
	if got.Raid.PowerWordFortitude != proto.TristateEffect_TristateEffectRegular {
		t.Errorf("the plain id set %v", got.Raid.PowerWordFortitude)
	}

	// A field that is not graded has no improved form, and accepting
	// one would silently apply the plain buff under a name that
	// promised more.
	if _, err := buffsFor([]string{"blessing_of_kings:improved"}); err == nil {
		t.Error("an improved form of a boolean buff was accepted")
	}
	if _, err := buffsFor([]string{"no_such_buff:improved"}); err == nil {
		t.Error("an improved form of an unknown buff was accepted")
	}
}

func TestKnownBuffsListsBothForms(t *testing.T) {
	ids := KnownBuffs()
	for _, want := range []string{"battle_shout", "battle_shout:improved"} {
		if !slices.Contains(ids, want) {
			t.Errorf("KnownBuffs does not list %q", want)
		}
	}
	if slices.Contains(ids, "blessing_of_kings:improved") {
		t.Error("KnownBuffs lists an improved form for a buff that has none")
	}
	for _, id := range ids {
		if _, err := buffsFor([]string{id}); err != nil {
			t.Errorf("KnownBuffs lists %q, which does not resolve: %v", id, err)
		}
	}
}

// The world buffs are a section of IndividualBuffs, and the engine
// marks them only with a comment. So the list is written down here and
// held to the message: the engine's world-buff fields are numbered from
// worldBuffFirstField up, and a new one that nobody added to the list
// fails rather than going missing from the settings bar.
func TestWorldBuffsCoverTheEnginesWorldBuffSection(t *testing.T) {
	names := map[string]bool{}
	for _, id := range WorldBuffs() {
		names[id] = true
	}
	desc := (&proto.IndividualBuffs{}).ProtoReflect().Descriptor()
	fields := desc.Fields()
	var expected int
	for i := 0; i < fields.Len(); i++ {
		fd := fields.Get(i)
		if fd.Number() < worldBuffFirstField {
			continue
		}
		expected++
		if !names[string(fd.Name())] {
			t.Errorf("IndividualBuffs.%s is a world buff and WorldBuffs() does not list it", fd.Name())
		}
	}
	if len(names) != expected {
		t.Errorf("WorldBuffs() has %d entries and IndividualBuffs has %d world-buff fields", len(names), expected)
	}
	if !slices.IsSorted(WorldBuffs()) {
		t.Error("WorldBuffs() is not sorted; the list is an interface and must be stable")
	}
}

// Stat ids are the engine's enum, so the weights panel is built from
// the engine rather than from a table that would drift.
func TestKnownStatsAreTheEnginesEnum(t *testing.T) {
	ids := KnownStats()
	if !slices.IsSorted(ids) {
		t.Error("KnownStats is not sorted")
	}
	if len(ids) != len(proto.Stat_name) {
		t.Errorf("KnownStats has %d entries, the enum has %d", len(ids), len(proto.Stat_name))
	}
	for _, id := range ids {
		if _, ok := ParseStat(id); !ok {
			t.Errorf("KnownStats lists %q, which does not resolve", id)
		}
	}
	for _, want := range []string{"agility", "attack_power", "crit", "hit", "spell_power", "spell_haste", "melee_haste"} {
		if !slices.Contains(ids, want) {
			t.Errorf("KnownStats does not list %q", want)
		}
	}
	// Section 1.4 says "haste"; the engine has two, and 10.8 settles
	// it. Likewise there is one hit and one crit, not four.
	for _, absent := range []string{"haste", "melee_crit", "spell_crit", "melee_hit", "spell_hit", "m_p5"} {
		if _, ok := ParseStat(absent); ok {
			t.Errorf("ParseStat resolved %q, which is not in the pinned vocabulary", absent)
		}
	}
	if got, _ := ParseStat("crit"); got != proto.Stat_StatCrit {
		t.Errorf("ParseStat(crit) = %v", got)
	}
	if got, _ := ParseStat("mp5"); got != proto.Stat_StatMP5 {
		t.Errorf("ParseStat(mp5) = %v; strcase.Snake would spell it m_p5 and statSpellings is what stops it", got)
	}
}

// The stat ids are an interface: the page builds the weights panel
// from them, the addon's Pawn string names them, and specs.json's
// reference_stat is one of them. So the whole list is pinned to the
// contract's section 10.8 verbatim rather than being whatever the
// derivation happens to produce.
func TestKnownStatsMatchThePinnedList(t *testing.T) {
	pinned := []string{
		"strength", "agility", "stamina", "intellect", "spirit",
		"spell_power", "arcane_power", "fire_power", "frost_power",
		"holy_power", "nature_power", "shadow_power", "mp5", "hit",
		"crit", "spell_haste", "spell_penetration", "attack_power",
		"melee_haste", "armor_penetration", "expertise", "mana",
		"energy", "rage", "armor", "ranged_attack_power", "defense",
		"block", "block_value", "dodge", "parry", "health",
		"arcane_resistance", "fire_resistance", "frost_resistance",
		"nature_resistance", "shadow_resistance", "bonus_armor",
		"healing_power", "spell_damage", "feral_attack_power",
	}
	slices.Sort(pinned)
	got := KnownStats()
	if !slices.Equal(got, pinned) {
		// Name the difference both ways: a stat the engine gained is
		// a decision to publish, and a stat it lost is a page control
		// that would send an id nothing resolves.
		for _, id := range got {
			if !slices.Contains(pinned, id) {
				t.Errorf("KnownStats lists %q, which the contract's 10.8 does not", id)
			}
		}
		for _, id := range pinned {
			if !slices.Contains(got, id) {
				t.Errorf("the contract's 10.8 lists %q and KnownStats does not", id)
			}
		}
	}
}
