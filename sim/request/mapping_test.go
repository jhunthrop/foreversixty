package request

import (
	"errors"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/internal/strcase"
	"github.com/jhunthrop/foreversixty/sim/specs"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

// The slot table is positional and nothing at run time complains when it
// is wrong, so it is checked against the engine's own enum rather than
// trusted. proto.ItemSlot's names are the same names in upper camel.
func TestSlotOrderMatchesTheEngineEnum(t *testing.T) {
	if len(proto.ItemSlot_name) != SlotCount {
		t.Fatalf("proto.ItemSlot has %d values, SlotCount is %d", len(proto.ItemSlot_name), SlotCount)
	}
	// SlotCount is a constant rather than len(slotOrder), so that an
	// importer cannot assign to it and hand the engine a short
	// equipment array. This is what keeps the two in step.
	if len(slotOrder) != SlotCount {
		t.Fatalf("slotOrder has %d entries, SlotCount is %d", len(slotOrder), SlotCount)
	}
	for i, want := range slotOrder {
		got := strcase.Snake(strings.TrimPrefix(proto.ItemSlot(i).String(), "ItemSlot"))
		if got != want {
			t.Errorf("slot %d is %q in the engine and %q here", i, got, want)
		}
	}
}

// The level the envelope insists on is the engine's own, not a number
// of ours: sim/core builds every character at core.CharacterMaxLevel
// and proto.Player has no level field, so nothing else can be run.
func TestTheEnvelopesSimLevelIsTheEngines(t *testing.T) {
	if api.SimLevel != core.CharacterMaxLevel {
		t.Errorf("api.SimLevel = %d, the engine builds characters at %d", api.SimLevel, core.CharacterMaxLevel)
	}
}

// The engine has no per-player level, so a request for anything but its
// own level must fail rather than be answered with a level-60 sim. The
// envelope's own validation is where that refusal lives, so there is
// one level rule in the module rather than two.
func TestBuildRejectsALevelTheEngineCannotSimulate(t *testing.T) {
	req := fury()
	req.Character.Level = 40
	_, err := Build(req)
	if err == nil {
		t.Fatal("a level-40 character was built without error")
	}
	if !strings.Contains(err.Error(), "character.level") {
		t.Errorf("error %q does not name the level", err)
	}
}

// A buff id is the engine's field name, and it lands in exactly one of
// the four buff messages: the most specific one that carries it.
func TestBuffIdsLandInTheRightMessage(t *testing.T) {
	got, err := buffsFor([]string{"battle_shout", "blessing_of_wisdom", "sunder_armor", "innervates", "mana_tide_totems"})
	if err != nil {
		t.Fatal(err)
	}
	if got.Raid.BattleShout != proto.TristateEffect_TristateEffectRegular {
		t.Errorf("RaidBuffs.BattleShout = %v, want the plain version", got.Raid.BattleShout)
	}
	if got.Individual.BlessingOfWisdom != proto.TristateEffect_TristateEffectRegular {
		t.Errorf("IndividualBuffs.BlessingOfWisdom = %v; the personal blessing wins over the raid-wide field of the same name", got.Individual.BlessingOfWisdom)
	}
	if got.Raid.BlessingOfWisdom != proto.TristateEffect_TristateEffectMissing {
		t.Error("blessing_of_wisdom was applied twice, once per message")
	}
	if !got.Debuffs.SunderArmor {
		t.Error("Debuffs.SunderArmor is not set")
	}
	if got.Individual.Innervates != 1 {
		t.Errorf("IndividualBuffs.Innervates = %d, want 1", got.Individual.Innervates)
	}
	if got.Party.ManaTideTotems != 1 {
		t.Errorf("PartyBuffs.ManaTideTotems = %d, want 1", got.Party.ManaTideTotems)
	}
}

func TestUnknownBuffFails(t *testing.T) {
	if _, err := buffsFor([]string{"blessing_of_vulpera"}); !errors.Is(err, ErrUnknownBuff) {
		t.Fatalf("buffsFor with an unknown id returned %v, want ErrUnknownBuff", err)
	}
}

// A consumable id is an enum value name; the field it belongs in is
// found rather than tabulated.
func TestConsumeIdsMapOntoTheirField(t *testing.T) {
	got, err := consumes([]string{"elixir_of_the_mongoose", "flask_of_the_titans", "dragon_breath_chili"}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if got.AgilityElixir != proto.AgilityElixir_ElixirOfTheMongoose {
		t.Errorf("AgilityElixir = %v", got.AgilityElixir)
	}
	if got.Flask != proto.Flask_FlaskOfTheTitans {
		t.Errorf("Flask = %v", got.Flask)
	}
	if !got.DragonBreathChili {
		t.Error("DragonBreathChili is not set; a consumable that is its own flag must still map")
	}
}

// Both weapon imbue slots hold the same enum, so a bare imbue id says
// nothing about which hand and must be refused rather than guessed.
func TestAmbiguousConsumeMustBeQualified(t *testing.T) {
	_, err := consumes([]string{"shadow_oil"}, nil)
	if !errors.Is(err, ErrAmbiguousConsume) {
		t.Fatalf("a bare weapon imbue returned %v, want ErrAmbiguousConsume", err)
	}
	if !strings.Contains(err.Error(), "main_hand_imbue") || !strings.Contains(err.Error(), "off_hand_imbue") {
		t.Errorf("the error does not name both slots: %v", err)
	}
	got, err := consumes([]string{"off_hand_imbue:shadow_oil"}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if got.OffHandImbue != proto.WeaponImbue_ShadowOil {
		t.Errorf("OffHandImbue = %v, want ShadowOil", got.OffHandImbue)
	}
	if got.MainHandImbue != proto.WeaponImbue_WeaponImbueUnknown {
		t.Errorf("MainHandImbue = %v, want it untouched", got.MainHandImbue)
	}
}

func TestUnknownConsumeFails(t *testing.T) {
	if _, err := consumes([]string{"elixir_of_vulpera"}, nil); !errors.Is(err, ErrUnknownConsume) {
		t.Fatalf("consumes with an unknown id returned %v, want ErrUnknownConsume", err)
	}
}

// Every spec the builder carries must attach both halves the engine
// needs: its options and a rotation.
func TestEverySpecAttachesOptionsAndARotation(t *testing.T) {
	for slug := range specOptions {
		t.Run(slug, func(t *testing.T) {
			req := fury()
			req.Spec = slug
			// The class has to follow the spec: they are cross-checked,
			// so a warrior cannot be handed a mage's rotation.
			req.Character.Class = specs.ByKey[slug].ClassSlug
			req.Character.Talents = ""
			req.Character.Gear = nil
			got, err := Build(req)
			if err != nil {
				t.Fatal(err)
			}
			p := got.Raid.Parties[0].Players[0]
			if p.Spec == nil {
				t.Error("no spec options")
			}
			if p.Rotation == nil || len(p.Rotation.PriorityList) == 0 {
				t.Error("no APL rotation; the engine would run an empty priority list")
			}
		})
	}
}

// A spec on the canonical list that no agent exists for yet, and a
// slug that is not a spec at all, are different refusals: the first is
// "not yet", the second is "never".
func TestBuildRejectsASpecItDoesNotCarry(t *testing.T) {
	req := fury()
	req.Spec = "shaman-enhancement"
	req.Character.Class = "shaman"
	req.Character.Talents = ""
	req.Character.Gear = nil
	if _, err := Build(req); !errors.Is(err, ErrUnsupportedSpec) {
		t.Fatalf("Build with an unsupported spec returned %v, want ErrUnsupportedSpec", err)
	}

	req = fury()
	req.Spec = "warrior-berserker"
	if _, err := Build(req); !errors.Is(err, ErrUnknownSpec) {
		t.Fatalf("Build with a spec that is not on the canonical list returned %v, want ErrUnknownSpec", err)
	}
}

// A warrior carrying spec "mage-frost" used to build: the player came
// from Character.Class and the agent from Spec, so the engine was
// handed a warrior running a frost mage's priority list and answered
// with a number. sim/specs is the authoritative pairing.
func TestBuildRejectsASpecFromAnotherClass(t *testing.T) {
	req := fury()
	req.Spec = "mage-frost"
	_, err := Build(req)
	if !errors.Is(err, ErrSpecClassMismatch) {
		t.Fatalf("Build accepted a warrior running mage-frost: %v", err)
	}
	if !strings.Contains(err.Error(), "mage") || !strings.Contains(err.Error(), "warrior") {
		t.Errorf("the error names neither side of the mismatch: %v", err)
	}
}

// The envelope's own validation runs before any mapping, so a malformed
// request never reaches the engine's protobufs.
func TestBuildValidatesTheEnvelopeFirst(t *testing.T) {
	var empty api.SimRequest
	if _, err := Build(empty); err == nil {
		t.Fatal("an empty request was built without error")
	}
}
