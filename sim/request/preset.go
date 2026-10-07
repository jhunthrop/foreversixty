package request

// Named sim presets: a set of raid buffs, target debuffs and consumables
// a request can carry besides the character's own kit.
//
// data/curated/presets.json states them in this module's own id
// vocabulary (vocabulary.go). LoadPresets refuses a file naming an id the
// request layer cannot map, or a buff filed as a debuff, so a typo fails
// when the file is read and not as a sim that quietly ran without the
// buff; Resolve turns a preset name into the ids one spec carries; and
// Layer puts the class kit (sim/leveling KitBuffs and KitConsumes) back
// on top, so the kit still wins any consumable slot it holds.

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/jhunthrop/foreversixty/sim/leveling"
	"github.com/jhunthrop/foreversixty/sim/specs"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// RaidPreset names the Phase 1 raid context in data/curated/presets.json.
const RaidPreset = "raid"

// BarePreset names the character with only its class kit applied.
const BarePreset = "bare"

var (
	// ErrUnknownPreset is returned for a preset name the curated file
	// does not state.
	ErrUnknownPreset = errors.New("request: unknown preset")
	// ErrPresetMisfiled is returned for a buff listed among the debuffs,
	// or a debuff listed among the buffs.
	ErrPresetMisfiled = errors.New("request: preset entry is in the wrong list")
	// ErrPresetNoGroup is returned when no consumable group of a preset
	// covers a spec.
	ErrPresetNoGroup = errors.New("request: preset has no consumable group for this spec")
)

// debuffOwner is the buff message vocabulary.go files every debuff under.
const debuffOwner = "Debuffs"

// PresetEntry is one id a preset applies and the label a player reads.
type PresetEntry struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

// presetEntry is PresetEntry as the curated file states it, with the
// reason the entry is there.
type presetEntry struct {
	PresetEntry
	Reason string `json:"reason"`
}

// consumeGroup is the consumables one role carries: every spec whose
// class is listed, else whose reference stat is listed.
type consumeGroup struct {
	Classes        []string      `json:"classes"`
	ReferenceStats []string      `json:"reference_stats"`
	IDs            []presetEntry `json:"ids"`
	// DualWieldIDs are added for the specs leveling.DualWieldSpecs names.
	DualWieldIDs []presetEntry `json:"dual_wield_ids"`
}

type presetConsumes struct {
	Groups  map[string]consumeGroup  `json:"groups"`
	ByClass map[string][]presetEntry `json:"by_class"`
	BySpec  map[string][]presetEntry `json:"by_spec"`
}

type presetDef struct {
	Label    string         `json:"label"`
	Notes    string         `json:"notes"`
	Buffs    []presetEntry  `json:"buffs"`
	Debuffs  []presetEntry  `json:"debuffs"`
	Consumes presetConsumes `json:"consumes"`
}

// Presets is the curated file's presets, by name.
type Presets map[string]presetDef

// ResolvedPreset is one preset as one spec carries it. Its JSON is the
// shape the BiS files publish under "presets".
type ResolvedPreset struct {
	Label    string        `json:"label"`
	Buffs    []PresetEntry `json:"buffs"`
	Debuffs  []PresetEntry `json:"debuffs"`
	Consumes []PresetEntry `json:"consumes"`
	Notes    string        `json:"notes"`

	slots map[string]string // consumable id -> the Consumes field it fills
}

// LoadPresets reads and validates the curated presets file at path.
func LoadPresets(path string) (Presets, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("request: reading presets: %w", err)
	}
	var file struct {
		Presets Presets `json:"presets"`
	}
	if err := json.Unmarshal(raw, &file); err != nil {
		return nil, fmt.Errorf("request: decoding %s: %w", path, err)
	}
	for name, def := range file.Presets {
		if err := def.validate(); err != nil {
			return nil, fmt.Errorf("request: preset %q: %w", name, err)
		}
	}
	return file.Presets, nil
}

func (d presetDef) validate() error {
	owners := buffOwners()
	if err := validateBuffList(d.Buffs, owners, false); err != nil {
		return err
	}
	if err := validateBuffList(d.Debuffs, owners, true); err != nil {
		return err
	}
	for _, group := range d.Consumes.Groups {
		if err := validateConsumeList(append(slices.Clone(group.IDs), group.DualWieldIDs...)); err != nil {
			return err
		}
	}
	for _, lists := range []map[string][]presetEntry{d.Consumes.ByClass, d.Consumes.BySpec} {
		for _, list := range lists {
			if err := validateConsumeList(list); err != nil {
				return err
			}
		}
	}
	return nil
}

// buffOwners maps every buff id onto the message that carries it.
func buffOwners() map[string]string {
	vocabulary := buffVocabulary()
	owners := make(map[string]string, len(vocabulary))
	for _, entry := range vocabulary {
		owners[entry.id] = entry.owner
	}
	return owners
}

func validateBuffList(entries []presetEntry, owners map[string]string, wantDebuffs bool) error {
	for _, entry := range entries {
		if err := entry.validateLabel(); err != nil {
			return err
		}
		owner, ok := owners[entry.ID]
		if !ok {
			return fmt.Errorf("%w: %q", ErrUnknownBuff, entry.ID)
		}
		if (owner == debuffOwner) != wantDebuffs {
			return fmt.Errorf("%w: %q lands in %s", ErrPresetMisfiled, entry.ID, owner)
		}
	}
	return nil
}

func validateConsumeList(entries []presetEntry) error {
	for _, entry := range entries {
		if err := entry.validateLabel(); err != nil {
			return err
		}
		if _, err := consumeSlot(entry.ID); err != nil {
			return err
		}
	}
	return nil
}

func (e presetEntry) validateLabel() error {
	if strings.TrimSpace(e.Label) == "" {
		return fmt.Errorf("entry %q has no label", e.ID)
	}
	return nil
}

// consumeSlot names the Consumes field an id fills, so two ids that
// would fight over one field can be told apart. An id the request layer
// cannot map, or a bare imbue that names both hands, is an error.
func consumeSlot(id string) (string, error) {
	set, err := consumes([]string{id}, nil)
	if err != nil {
		return "", err
	}
	slot := ""
	set.ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, _ protoreflect.Value) bool {
		slot = string(fd.Name())
		return false
	})
	return slot, nil
}

// Resolve returns the preset as spec carries it: the shared buffs and
// debuffs, and the consumables of the spec's group, class and school.
func (p Presets) Resolve(name string, spec specs.Spec) (ResolvedPreset, error) {
	def, ok := p[name]
	if !ok {
		return ResolvedPreset{}, fmt.Errorf("%w: %q", ErrUnknownPreset, name)
	}
	group, err := def.Consumes.groupFor(spec)
	if err != nil {
		return ResolvedPreset{}, fmt.Errorf("preset %q: %w", name, err)
	}
	var entries []presetEntry
	entries = append(entries, group.IDs...)
	if leveling.DualWieldSpecs[spec.Spec] {
		entries = append(entries, group.DualWieldIDs...)
	}
	entries = append(entries, def.Consumes.ByClass[spec.ClassSlug]...)
	entries = append(entries, def.Consumes.BySpec[spec.Spec]...)

	resolved := ResolvedPreset{
		Label:    def.Label,
		Buffs:    publicEntries(def.Buffs),
		Debuffs:  publicEntries(def.Debuffs),
		Consumes: publicEntries(entries),
		Notes:    def.Notes,
		slots:    make(map[string]string, len(entries)),
	}
	for _, entry := range entries {
		slot, err := consumeSlot(entry.ID)
		if err != nil {
			return ResolvedPreset{}, err
		}
		resolved.slots[entry.ID] = slot
	}
	return resolved, nil
}

func (c presetConsumes) groupFor(spec specs.Spec) (consumeGroup, error) {
	names := make([]string, 0, len(c.Groups))
	for name := range c.Groups {
		names = append(names, name)
	}
	slices.Sort(names)
	for _, name := range names {
		if slices.Contains(c.Groups[name].Classes, spec.ClassSlug) {
			return c.Groups[name], nil
		}
	}
	for _, name := range names {
		if slices.Contains(c.Groups[name].ReferenceStats, spec.ReferenceStat) {
			return c.Groups[name], nil
		}
	}
	return consumeGroup{}, fmt.Errorf("%w: %s", ErrPresetNoGroup, spec.Spec)
}

func publicEntries(entries []presetEntry) []PresetEntry {
	out := make([]PresetEntry, len(entries))
	for i, entry := range entries {
		out[i] = entry.PresetEntry
	}
	return out
}

// Layer puts the class kit on top of the preset and returns the buff and
// consumable ids a character carries. A buff the kit names (in either its
// plain or its improved form) is carried once, in the kit's form; a
// preset consumable whose slot the kit already fills is dropped, so a
// rogue's poison or an enhancement shaman's Windfury Weapon is never
// displaced by the preset's weapon imbue.
func (r ResolvedPreset) Layer(kitBuffs, kitConsumes []string) (buffs, consumeIDs []string) {
	buffs = slices.Clone(kitBuffs)
	kitBases := make(map[string]bool, len(kitBuffs))
	for _, id := range kitBuffs {
		kitBases[buffBase(id)] = true
	}
	for _, entry := range append(slices.Clone(r.Buffs), r.Debuffs...) {
		if !kitBases[buffBase(entry.ID)] {
			buffs = append(buffs, entry.ID)
		}
	}

	consumeIDs = slices.Clone(kitConsumes)
	kitSlots := make(map[string]bool, len(kitConsumes))
	for _, id := range kitConsumes {
		slot, err := consumeSlot(id)
		if err != nil {
			slot = id
		}
		kitSlots[slot] = true
	}
	for _, entry := range r.Consumes {
		if !kitSlots[r.slots[entry.ID]] {
			consumeIDs = append(consumeIDs, entry.ID)
		}
	}
	return buffs, consumeIDs
}

func buffBase(id string) string {
	base, _ := strings.CutSuffix(id, improvedSuffix)
	return base
}

// ResolveFromFile reads the curated presets at path and resolves the named
// one for spec, for the search tools that sim a character under a preset.
// A name of "bare" (the character with only its class kit) resolves to nil.
func ResolveFromFile(path, name string, spec specs.Spec) (*ResolvedPreset, error) {
	if name == BarePreset {
		return nil, nil
	}
	presets, err := LoadPresets(path)
	if err != nil {
		return nil, err
	}
	resolved, err := presets.Resolve(name, spec)
	if err != nil {
		return nil, err
	}
	return &resolved, nil
}
