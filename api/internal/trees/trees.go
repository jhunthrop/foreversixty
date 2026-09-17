// Package trees holds the client's talent, item, class, and race data for
// every build under TREE_DATA_DIR. It is read once at startup and then only
// read from, so every accessor is safe for concurrent use.
package trees

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type Rank struct {
	SpellID     int    `json:"spell_id"`
	Description string `json:"description"`
}

type Talent struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Icon           string `json:"icon"`
	MaxRank        int    `json:"max_rank"`
	Tier           int    `json:"tier"`
	Column         int    `json:"column"`
	PrereqTalentID *int   `json:"prereq_talent_id"`
	PrereqRank     *int   `json:"prereq_rank"`
	Ranks          []Rank `json:"ranks"`
	// SpellID is the spell the client writes when the talent is learned,
	// which is what a combat log's COMBATANT_INFO carries. The ID above is
	// the client's trait node id, and the two are different numbers.
	SpellID int `json:"spell_id"`
}

type Tree struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Position int      `json:"position"`
	Talents  []Talent `json:"talents"`
	// Background names the processed art at /data/<build>/trees/<background>.webp.
	Background string `json:"background"`
}

type classTalents struct {
	Build     string `json:"build"`
	ClassID   int    `json:"class_id"`
	ClassSlug string `json:"class_slug"`
	Trees     []Tree `json:"trees"`
}

type Item struct {
	ID            int            `json:"id"`
	Name          string         `json:"name"`
	Icon          string         `json:"icon"`
	Slot          string         `json:"slot"`
	Quality       int            `json:"quality"`
	RequiredLevel int            `json:"required_level"`
	ItemLevel     int            `json:"item_level"`
	Armor         int            `json:"armor"`
	Stats         map[string]int `json:"stats"`
	SetID         *int           `json:"set_id"`
	Unique        bool           `json:"unique"`
}

type classItems struct {
	Build     string `json:"build"`
	ClassSlug string `json:"class_slug"`
	Items     []Item `json:"items"`
}

type SetBonus struct {
	Pieces      int    `json:"pieces"`
	Description string `json:"description"`
}

type Set struct {
	ID      int        `json:"id"`
	Name    string     `json:"name"`
	ItemIDs []int      `json:"item_ids"`
	Bonuses []SetBonus `json:"bonuses"`
}

type Class struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Slug  string `json:"slug"`
	Color string `json:"color"`
}

type Race struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Slug    string `json:"slug"`
	Faction string `json:"faction"`
}

type Combo struct {
	RaceID       int  `json:"race_id"`
	ClassID      int  `json:"class_id"`
	NewInForever bool `json:"new_in_forever"`
}

// TalentRef is a talent plus the tree it lives in, which the validator needs
// for tier and prerequisite messages and the page needs for the point split.
type TalentRef struct {
	Talent
	TreeID       int
	TreeName     string
	TreePosition int
}

// Build is one client build's data, indexed for lookup by id.
type Build struct {
	Version string

	classes        map[int]Class
	races          map[int]Race
	combos         map[[2]int]Combo
	trees          map[int][]Tree
	talents        map[int]map[int]TalentRef
	talentsBySpell map[int]map[int]TalentRef
	items          map[int]map[int]Item
	sets           []Set
}

func (b *Build) Class(id int) (Class, bool) { c, ok := b.classes[id]; return c, ok }

func (b *Build) Race(id int) (Race, bool) { r, ok := b.races[id]; return r, ok }

func (b *Build) ComboAllowed(raceID, classID int) bool {
	_, ok := b.combos[[2]int{raceID, classID}]
	return ok
}

func (b *Build) Talent(classID, talentID int) (TalentRef, bool) {
	byID, ok := b.talents[classID]
	if !ok {
		return TalentRef{}, false
	}
	t, ok := byID[talentID]
	return t, ok
}

// TalentBySpellID resolves the spell id the client writes to the talent it
// belongs to. A build emitted before talents carried a spell id indexes every
// talent under 0, so a lookup for 0 is refused rather than answering with an
// arbitrary talent.
func (b *Build) TalentBySpellID(classID, spellID int) (TalentRef, bool) {
	if spellID == 0 {
		return TalentRef{}, false
	}
	bySpell, ok := b.talentsBySpell[classID]
	if !ok {
		return TalentRef{}, false
	}
	t, ok := bySpell[spellID]
	return t, ok
}

func (b *Build) Item(classID, itemID int) (Item, bool) {
	byID, ok := b.items[classID]
	if !ok {
		return Item{}, false
	}
	it, ok := byID[itemID]
	return it, ok
}

// Trees returns the class's trees in client order (by position).
func (b *Build) Trees(classID int) []Tree { return b.trees[classID] }

func (b *Build) Sets() []Set { return b.sets }

// Classes lists the build's classes by id. The rankings need it to turn
// a class name the engine read out of a log into the class id the
// talent data is keyed by.
func (b *Build) Classes() []Class {
	out := make([]Class, 0, len(b.classes))
	for _, c := range b.classes {
		out = append(out, c)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out
}

// Data holds every build found under the tree data directory.
type Data struct {
	builds  map[string]*Build
	skipped []string
}

func (d *Data) Build(version string) (*Build, bool) {
	b, ok := d.builds[version]
	return b, ok
}

func (d *Data) Versions() []string {
	out := make([]string, 0, len(d.builds))
	for v := range d.builds {
		out = append(out, v)
	}
	sort.Strings(out)
	return out
}

// Skipped lists directories that were not loaded because they predate the
// Phase 1 pipeline outputs (no combos.json or no talents directory), and the
// root directory itself when it does not exist. Startup logs these: the API
// still serves health, version, and subscribe without any tree data, and
// every save then fails validation on tree_version, which is the right
// behaviour while the data pipeline is catching up.
func (d *Data) Skipped() []string { return d.skipped }

// Latest is the newest build's data, which is what the rankings infer
// specs from: talent ids are stable across client builds, and a report
// parsed today is best read against today's trees. The second return is
// false when no build loaded at all.
//
// "Newest" is a numeric comparison of the dot-separated version, not the
// lexicographic order Versions() returns: build directories are named
// like "1.15.9.69722", and the patch component is already double-digit,
// so a plain string sort would rank "1.9" above "1.10".
func (d *Data) Latest() (*Build, bool) {
	versions := d.Versions()
	if len(versions) == 0 {
		return nil, false
	}
	return d.builds[newestVersion(versions)], true
}

// isClientBuild reports whether a version looks like a client build string
// ("1.60.1.69893") rather than a named data set ("forever-prebeta"). Only
// client builds are comparable as version numbers, and Latest must not hand
// the rankings a data set that is not the newest client data.
func isClientBuild(version string) bool {
	segments := strings.Split(version, ".")
	if len(segments) < 2 {
		return false
	}
	for _, segment := range segments {
		if _, err := strconv.Atoi(segment); err != nil {
			return false
		}
	}
	return true
}

// newestVersion returns the numerically greatest of a non-empty list of
// dot-separated build versions. A version that looks like a client build
// ("1.60.1.69893") is preferred over one that looks like a named data set
// ("forever-prebeta"), regardless of how the two compare as strings.
func newestVersion(versions []string) string {
	best := versions[0]
	for _, v := range versions[1:] {
		bestIsBuild, vIsBuild := isClientBuild(best), isClientBuild(v)
		if bestIsBuild != vIsBuild {
			if vIsBuild {
				best = v
			}
			continue
		}
		if compareVersions(v, best) > 0 {
			best = v
		}
	}
	return best
}

// compareVersions orders two dot-separated version strings segment by
// segment: a pair of segments that both parse as non-negative integers is
// compared numerically ("9" < "10"), and any other pair is compared as
// plain text, so a non-numeric segment never panics, it just orders
// lexicographically. A version with fewer segments than the other is
// lower once the shared segments are equal, the way "1.15" sorts below
// "1.15.1". It returns -1, 0, or 1 the way strings.Compare does.
func compareVersions(a, b string) int {
	as := strings.Split(a, ".")
	bs := strings.Split(b, ".")
	n := len(as)
	if len(bs) > n {
		n = len(bs)
	}
	for i := 0; i < n; i++ {
		var sa, sb string
		if i < len(as) {
			sa = as[i]
		}
		if i < len(bs) {
			sb = bs[i]
		}
		if sa == sb {
			continue
		}
		na, aErr := strconv.Atoi(sa)
		nb, bErr := strconv.Atoi(sb)
		if aErr == nil && bErr == nil {
			if na == nb {
				continue
			}
			if na < nb {
				return -1
			}
			return 1
		}
		if sa < sb {
			return -1
		}
		return 1
	}
	return 0
}

// Load reads every build directory under dir. A directory that does not look
// like a Phase 1 build is skipped (see Skipped); a directory that does look
// like one but contains bad data fails the load.
func Load(dir string) (*Data, error) {
	d := &Data{builds: map[string]*Build{}}
	entries, err := os.ReadDir(dir)
	if errors.Is(err, fs.ErrNotExist) {
		d.skipped = append(d.skipped, dir+": directory does not exist")
		return d, nil
	}
	if err != nil {
		return nil, fmt.Errorf("trees: read %s: %w", dir, err)
	}
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		path := filepath.Join(dir, e.Name())
		if !looksLikeBuild(path) {
			d.skipped = append(d.skipped, path+": no combos.json or talents directory")
			continue
		}
		b, err := loadBuild(path, e.Name())
		if err != nil {
			return nil, err
		}
		d.builds[e.Name()] = b
	}
	return d, nil
}

func looksLikeBuild(dir string) bool {
	if _, err := os.Stat(filepath.Join(dir, "combos.json")); err != nil {
		return false
	}
	info, err := os.Stat(filepath.Join(dir, "talents"))
	return err == nil && info.IsDir()
}

func loadBuild(dir, version string) (*Build, error) {
	b := &Build{
		Version:        version,
		classes:        map[int]Class{},
		races:          map[int]Race{},
		combos:         map[[2]int]Combo{},
		trees:          map[int][]Tree{},
		talents:        map[int]map[int]TalentRef{},
		talentsBySpell: map[int]map[int]TalentRef{},
		items:          map[int]map[int]Item{},
	}

	var classes []Class
	if err := readJSON(filepath.Join(dir, "classes.json"), &classes); err != nil {
		return nil, err
	}
	var races []Race
	if err := readJSON(filepath.Join(dir, "races.json"), &races); err != nil {
		return nil, err
	}
	var combos []Combo
	if err := readJSON(filepath.Join(dir, "combos.json"), &combos); err != nil {
		return nil, err
	}
	for _, c := range classes {
		b.classes[c.ID] = c
	}
	for _, r := range races {
		b.races[r.ID] = r
	}
	for _, c := range combos {
		if _, ok := b.races[c.RaceID]; !ok {
			return nil, fmt.Errorf("trees: %s: combos.json references unknown race %d", version, c.RaceID)
		}
		if _, ok := b.classes[c.ClassID]; !ok {
			return nil, fmt.Errorf("trees: %s: combos.json references unknown class %d", version, c.ClassID)
		}
		b.combos[[2]int{c.RaceID, c.ClassID}] = c
	}

	for _, c := range classes {
		var ct classTalents
		if err := readJSON(filepath.Join(dir, "talents", c.Slug+".json"), &ct); err != nil {
			return nil, err
		}
		if ct.ClassID != c.ID {
			return nil, fmt.Errorf("trees: %s: talents/%s.json has class_id %d, want %d", version, c.Slug, ct.ClassID, c.ID)
		}
		sort.SliceStable(ct.Trees, func(i, j int) bool { return ct.Trees[i].Position < ct.Trees[j].Position })
		byTalentID := map[int]TalentRef{}
		bySpellID := map[int]TalentRef{}
		for _, tree := range ct.Trees {
			for _, t := range tree.Talents {
				if len(t.Ranks) != t.MaxRank {
					return nil, fmt.Errorf("trees: %s: talent %d has %d ranks, want max_rank %d", version, t.ID, len(t.Ranks), t.MaxRank)
				}
				if _, dup := byTalentID[t.ID]; dup {
					return nil, fmt.Errorf("trees: %s: duplicate talent id %d in class %s", version, t.ID, c.Slug)
				}
				byTalentID[t.ID] = TalentRef{Talent: t, TreeID: tree.ID, TreeName: tree.Name, TreePosition: tree.Position}
				if t.SpellID != 0 {
					if _, dup := bySpellID[t.SpellID]; dup {
						return nil, fmt.Errorf("trees: %s: duplicate talent spell id %d in class %s", version, t.SpellID, c.Slug)
					}
					bySpellID[t.SpellID] = byTalentID[t.ID]
				}
			}
		}
		b.trees[c.ID] = ct.Trees
		b.talents[c.ID] = byTalentID
		b.talentsBySpell[c.ID] = bySpellID

		var ci classItems
		err := readJSON(filepath.Join(dir, "items", c.Slug+".json"), &ci)
		switch {
		case errors.Is(err, fs.ErrNotExist):
			// The pipeline emits an items file only when the item table
			// normalizes cleanly for that class; absence is expected.
			b.items[c.ID] = map[int]Item{}
		case err != nil:
			return nil, err
		default:
			byItemID := make(map[int]Item, len(ci.Items))
			for _, it := range ci.Items {
				byItemID[it.ID] = it
			}
			b.items[c.ID] = byItemID
		}
	}

	if err := readJSON(filepath.Join(dir, "sets.json"), &b.sets); err != nil && !errors.Is(err, fs.ErrNotExist) {
		return nil, err
	}
	return b, nil
}

// readJSON decodes path into v. Unknown fields are allowed on purpose: the
// pipeline adds fields (forever_changes, sources) the API has no use for,
// and a new field there must not break the API.
func readJSON(path string, v any) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("trees: open %s: %w", path, err)
	}
	defer f.Close()
	if err := json.NewDecoder(f).Decode(v); err != nil {
		return fmt.Errorf("trees: decode %s: %w", path, err)
	}
	return nil
}
