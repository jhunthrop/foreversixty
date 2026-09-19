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
