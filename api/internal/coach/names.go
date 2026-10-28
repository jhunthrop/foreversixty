package coach

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/consumables"
)

// maxNameWords is the longest dictionary name Find looks for.
const maxNameWords = 5

const (
	deprecatedMarker = "deprecated"
	minPluralLen     = 3 // a last word this short or shorter is not tried as a plural
)

// Directories and files of one build's data the dictionary reads.
const (
	spellConstDir  = "spellconst"
	trainablesDir  = "trainables"
	talentsDir     = "talents"
	itemNamesFile  = "itemnames.json"
	spellsFile     = "spells.json"
	itemsFile      = "items.json"
	jsonGlobSuffix = "*.json"
)

var wordRE = regexp.MustCompile(`[\p{L}\p{N}][\p{L}\p{N}'’-]*`)

// NameDict holds the build's spell, talent, item and boss names, keyed by their lowercased
// words, for n-gram lookup up to maxNameWords words.
type NameDict struct {
	names map[string]string // lowercased words joined by one space -> the name as written
}

// nameSpan is one dictionary name found in prose, by byte offsets.
type nameSpan struct {
	Start, End int
	Name       string
}

// LoadNames reads the spell names (spellconst, trainables and talents of every class) and
// item names of the build in buildDir, the boss names of the mechanics tables and the
// consumable and raid buff names of the consumables catalogue.
func LoadNames(buildDir string, mechanicsTables []mechanics.Table) (*NameDict, error) {
	d := &NameDict{names: map[string]string{}}
	loaders := []func(string) ([]string, error){
		spellNames, talentNames, itemNames, namedRows(spellsFile), namedRows(itemsFile),
	}
	for _, load := range loaders {
		names, err := load(buildDir)
		if err != nil {
			return nil, err
		}
		d.add(names...)
	}
	for _, t := range mechanicsTables {
		d.add(t.Name)
	}
	d.add(consumableNames(consumables.Load())...)
	if len(d.names) == 0 {
		return nil, fmt.Errorf("coach names: no names found under %s", buildDir)
	}
	return d, nil
}

// add puts names in the dictionary. Placeholder names ("DEPRECATED", "UNUSED") are skipped.
func (d *NameDict) add(names ...string) {
	for _, n := range names {
		words := splitWords(n)
		if len(words) == 0 || len(words) > maxNameWords || isPlaceholder(n) {
			continue
		}
		d.names[joinKey(words)] = strings.TrimSpace(n)
	}
}

// isPlaceholder is true for the build's dummy names: marked deprecated, or all capitals.
func isPlaceholder(name string) bool {
	return strings.Contains(strings.ToLower(name), deprecatedMarker) || name == strings.ToUpper(name)
}

// Find returns every dictionary name in s, as the dictionary writes it, in the order they
// appear. Only capitalised n-grams count, and the longest match at a position wins.
func (d *NameDict) Find(s string) []string {
	spans := d.spans(s)
	out := make([]string, len(spans))
	for i, sp := range spans {
		out[i] = sp.Name
	}
	return out
}

// spans is Find with positions.
func (d *NameDict) spans(s string) []nameSpan {
	locs := wordRE.FindAllStringIndex(s, -1)
	var out []nameSpan
	for i := 0; i < len(locs); {
		sp, words, ok := d.matchAt(s, locs[i:])
		if !ok {
			i++
			continue
		}
		out = append(out, sp)
		i += words
	}
	return out
}

// matchAt tries the longest name first at the first word of locs.
func (d *NameDict) matchAt(s string, locs [][]int) (nameSpan, int, bool) {
	for n := min(maxNameWords, len(locs)); n >= 1; n-- {
		words := make([]string, n)
		for j := range words {
			words[j] = s[locs[j][0]:locs[j][1]]
		}
		canonical, form, ok := d.lookup(words)
		if ok && countsAsName(canonical, form) {
			return nameSpan{Start: locs[0][0], End: locs[n-1][1], Name: canonical}, n, true
		}
	}
	return nameSpan{}, 0, false
}

// countsAsName decides whether prose that matches a dictionary name is a use of that name.
// A phrase of two or more words counts in any case ("flask of the titans"); a single word
// counts only when capitalised as the dictionary writes it, so the common words that are
// also spell names ("rage", "charge", "unused") are prose when lowercase.
func countsAsName(canonical string, prose []string) bool {
	if len(prose) > 1 {
		return true
	}
	return startsUpper(prose[0]) && capitalsAgree(canonical, prose)
}

// lookup finds the name for words, allowing the last word to be possessive ("Sinister
// Strike's") or plural ("Aimed Shots"). It returns the name as written in the dictionary
// and the prose words as they matched it, without the possessive or plural ending.
func (d *NameDict) lookup(words []string) (string, []string, bool) {
	last := words[len(words)-1]
	for _, form := range []string{last, trimSuffixes(last, "'s", "’s"), trimSuffixes(last, "s")} {
		if form == "" || (form != last && len(last) <= minPluralLen) {
			continue
		}
		candidate := append(append([]string{}, words[:len(words)-1]...), form)
		if canonical, ok := d.names[joinKey(candidate)]; ok {
			return canonical, candidate, true
		}
	}
	return "", nil, false
}

// trimSuffixes removes the first of suffixes that word ends with; word itself when none do.
func trimSuffixes(word string, suffixes ...string) string {
	for _, suffix := range suffixes {
		if strings.HasSuffix(word, suffix) {
			return strings.TrimSuffix(word, suffix)
		}
	}
	return word
}

// capitalsAgree is true when every word the dictionary capitalises is capitalised in the
// prose too; "Flask of the Titans" matches "Flask of the Titans" but not "flask of the titans".
func capitalsAgree(canonical string, prose []string) bool {
	canon := splitWords(canonical)
	for i, w := range canon {
		if startsUpper(w) && !startsUpper(prose[i]) {
			return false
		}
	}
	return true
}

func startsUpper(w string) bool {
	r, _ := utf8.DecodeRuneInString(w)
	return unicode.IsUpper(r)
}

func splitWords(s string) []string { return wordRE.FindAllString(s, -1) }

func joinKey(words []string) string {
	return strings.ToLower(strings.ReplaceAll(strings.Join(words, " "), "’", "'"))
}

// readJSONFiles reads every *.json file in dir under buildDir; a missing directory is an
// error because a build without them cannot name its spells.
func readJSONFiles(buildDir, dir string) ([][]byte, error) {
	paths, err := filepath.Glob(filepath.Join(buildDir, dir, jsonGlobSuffix))
	if err != nil || len(paths) == 0 {
		return nil, fmt.Errorf("coach names: no %s files under %s (%v)", dir, buildDir, err)
	}
	out := make([][]byte, 0, len(paths))
	for _, p := range paths {
		data, err := os.ReadFile(p)
		if err != nil {
			return nil, fmt.Errorf("coach names: %w", err)
		}
		out = append(out, data)
	}
	return out, nil
}

func spellNames(buildDir string) ([]string, error) {
	var names []string
	files, err := readJSONFiles(buildDir, spellConstDir)
	if err != nil {
		return nil, err
	}
	for _, data := range files {
		spells, err := ParseSpellConst(data)
		if err != nil {
			return nil, err
		}
		for _, s := range spells {
			names = append(names, s.Name)
		}
	}
	files, err = readJSONFiles(buildDir, trainablesDir)
	if err != nil {
		return nil, err
	}
	for _, data := range files {
		families, err := ParseTrainables(data)
		if err != nil {
			return nil, err
		}
		for family := range families {
			names = append(names, family)
		}
	}
	return names, nil
}

// talentFile is the part of data/builds/<build>/talents/<class>.json the dictionary reads.
type talentFile struct {
	Trees []struct {
		Talents []struct {
			Name string `json:"name"`
		} `json:"talents"`
	} `json:"trees"`
}

func talentNames(buildDir string) ([]string, error) {
	files, err := readJSONFiles(buildDir, talentsDir)
	if err != nil {
		return nil, err
	}
	var names []string
	for _, data := range files {
		var f talentFile
		if err := json.Unmarshal(data, &f); err != nil {
			return nil, fmt.Errorf("coach names: talents: %w", err)
		}
		for _, tree := range f.Trees {
			for _, t := range tree.Talents {
				names = append(names, t.Name)
			}
		}
	}
	return names, nil
}

func itemNames(buildDir string) ([]string, error) {
	data, err := os.ReadFile(filepath.Join(buildDir, itemNamesFile))
	if err != nil {
		return nil, fmt.Errorf("coach names: %w", err)
	}
	var f struct {
		Names map[string]string `json:"names"`
	}
	if err := json.Unmarshal(data, &f); err != nil {
		return nil, fmt.Errorf("coach names: itemnames: %w", err)
	}
	names := make([]string, 0, len(f.Names))
	for _, n := range f.Names {
		names = append(names, n)
	}
	return names, nil
}

// consumableNames is every flask, food, enchant, potion, world buff and raid buff name the
// catalogue lists, for every role.
func consumableNames(c consumables.Catalogue) []string {
	var names []string
	for _, rc := range c.Roles {
		for _, category := range []string{
			consumables.CategoryFlask, consumables.CategoryFood, consumables.CategoryWeaponEnchant,
			consumables.CategoryWorldBuffs, consumables.CategoryCombatPotion,
		} {
			for _, e := range rc.Entries(category) {
				names = append(names, e.Name)
			}
		}
	}
	for _, b := range c.RaidBuffs {
		names = append(names, b.Name)
	}
	return names
}

// namedRows reads a build file that is an array of rows with a "name" (spells.json,
// items.json) and returns every row's name.
func namedRows(file string) func(string) ([]string, error) {
	return func(buildDir string) ([]string, error) {
		data, err := os.ReadFile(filepath.Join(buildDir, file))
		if err != nil {
			return nil, fmt.Errorf("coach names: %w", err)
		}
		var rows []struct {
			Name string `json:"name"`
		}
		if err := json.Unmarshal(data, &rows); err != nil {
			return nil, fmt.Errorf("coach names: %s: %w", file, err)
		}
		names := make([]string, len(rows))
		for i, r := range rows {
			names[i] = r.Name
		}
		return names, nil
	}
}
