package coach

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
)

// RuleEntry is one ability in a spec's rule file: its family name, every rank's spell id,
// whether it is a talent (assessed only when cast that fight, §6.1), and the stamp of the
// client build it was checked against.
type RuleEntry struct {
	Family   string  `json:"family"`
	SpellIDs []int64 `json:"spell_ids"`
	Talent   bool    `json:"talent"`
	Verified string  `json:"verified"`
}

// AuraRule is one owned aura F2 judges. SpellID is the aura's canonical (stamped) rank,
// the key its uptime is reported under; SpellIDs, when set, is every rank's id, and an
// application of any of them is the aura (a lower rank is still the aura).
type AuraRule struct {
	SpellID  int64   `json:"spell_id"`
	SpellIDs []int64 `json:"spell_ids,omitempty"`
	Name     string  `json:"name"`
	Kind     string  `json:"kind"` // AuraDot, AuraDebuff, AuraSelfBuff or AuraDefensive
	Verified string  `json:"verified"`
}

// IDs is every spell id that counts as this aura: SpellIDs, or SpellID alone when the
// rule lists no ranks.
func (a AuraRule) IDs() []int64 {
	if len(a.SpellIDs) > 0 {
		return a.SpellIDs
	}
	return []int64{a.SpellID}
}

// The aura kinds a rule file may name.
const (
	AuraDot       = "dot"
	AuraDebuff    = "debuff"
	AuraSelfBuff  = "self_buff"
	AuraDefensive = "defensive"
)

var auraKinds = map[string]bool{AuraDot: true, AuraDebuff: true, AuraSelfBuff: true, AuraDefensive: true}

// SpecRules is data/curated/coach/<spec>.json: the spec's core cooldown abilities, owned
// auras, deliberate downranks, defensive cooldowns and efficiency cooldowns (§6.1). A
// missing list sends its findings to "Couldn't check".
type SpecRules struct {
	Spec       string      `json:"spec"`
	Build      string      `json:"build"`
	Core       []RuleEntry `json:"core"`
	Auras      []AuraRule  `json:"auras"`
	Downranks  []RuleEntry `json:"downranks"`
	Defensives []RuleEntry `json:"defensives"`
	Efficiency []RuleEntry `json:"efficiency"`
}

// ParseSpecRules reads one rule file. It refuses an unknown field, a file with no spec,
// and any entry with no verified stamp, no spell id or an unknown aura kind, naming it.
func ParseSpecRules(data []byte) (SpecRules, error) {
	var r SpecRules
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&r); err != nil {
		return SpecRules{}, fmt.Errorf("coach rules: %w", err)
	}
	if r.Spec == "" {
		return SpecRules{}, errors.New("coach rules: no spec")
	}
	lists := map[string][]RuleEntry{"core": r.Core, "downranks": r.Downranks, "defensives": r.Defensives, "efficiency": r.Efficiency}
	for list, entries := range lists {
		for _, e := range entries {
			if err := checkRuleEntry(e); err != nil {
				return SpecRules{}, fmt.Errorf("coach rules %s: %s: %w", r.Spec, list, err)
			}
		}
	}
	for _, a := range r.Auras {
		if err := checkAuraRule(a); err != nil {
			return SpecRules{}, fmt.Errorf("coach rules %s: auras: %w", r.Spec, err)
		}
	}
	return r, nil
}

func checkRuleEntry(e RuleEntry) error {
	switch {
	case e.Family == "":
		return errors.New("an entry with no family")
	case e.Verified == "":
		return fmt.Errorf("%s has no verified stamp", e.Family)
	case len(e.SpellIDs) == 0:
		return fmt.Errorf("%s has no spell id", e.Family)
	}
	for _, id := range e.SpellIDs {
		if id <= 0 {
			return fmt.Errorf("%s has spell id %d", e.Family, id)
		}
	}
	return nil
}

func checkAuraRule(a AuraRule) error {
	switch {
	case a.Name == "":
		return fmt.Errorf("aura %d has no name", a.SpellID)
	case a.Verified == "":
		return fmt.Errorf("%s has no verified stamp", a.Name)
	case a.SpellID <= 0:
		return fmt.Errorf("%s has no spell id", a.Name)
	case !auraKinds[a.Kind]:
		return fmt.Errorf("%s has unknown kind %q", a.Name, a.Kind)
	case len(a.SpellIDs) > 0 && !slices.Contains(a.SpellIDs, a.SpellID):
		return fmt.Errorf("%s lists ranks without its spell id %d", a.Name, a.SpellID)
	}
	for _, id := range a.SpellIDs {
		if id <= 0 {
			return fmt.Errorf("%s has spell id %d", a.Name, id)
		}
	}
	return nil
}

// LoadSpecRules reads <dir>/<spec>.json. ok is false, with no error, when the file does
// not exist; a file whose spec field is not spec is an error.
func LoadSpecRules(dir, spec string) (SpecRules, bool, error) {
	data, err := os.ReadFile(filepath.Join(dir, spec+".json"))
	if errors.Is(err, fs.ErrNotExist) {
		return SpecRules{}, false, nil
	}
	if err != nil {
		return SpecRules{}, false, fmt.Errorf("coach rules %s: %w", spec, err)
	}
	r, err := ParseSpecRules(data)
	if err != nil {
		return SpecRules{}, false, err
	}
	if r.Spec != spec {
		return SpecRules{}, false, fmt.Errorf("coach rules %s: the file says spec %q", spec, r.Spec)
	}
	return r, true, nil
}
