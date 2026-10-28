package coach

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"unicode"
)

// PersonaStatus says how finished a persona file is (§8.3).
type PersonaStatus string

const (
	PersonaFull    PersonaStatus = "full"
	PersonaStarter PersonaStatus = "starter"
)

const (
	personaVersionPrefix = "persona-"
	drillIDPrefix        = "drill."
	maxDrillChars        = 240
	frontMatterFence     = "---"
	sectionPrefix        = "## "
	bulletPrefix         = "- "
)

// Section names, in the order a persona file holds them (§8.1).
const (
	sectionWho    = "Who they are"
	sectionKnow   = "What they know"
	sectionDrills = "Drills"
	sectionVoice  = "Voice"
)

var personaSections = []string{sectionWho, sectionKnow, sectionDrills, sectionVoice}

var frontMatterKeys = []string{"spec", "role", "status", "version", "sources", "reviewed"}

// Persona is one spec's coach persona: the system-prompt prose and the drill catalogue.
type Persona struct {
	Spec     string
	Status   PersonaStatus
	Version  string // personaVersionPrefix + hex sha256 of the file's bytes (§8.3)
	Sources  []string
	Reviewed string
	Body     string            // the four sections, as the system prompt carries them
	Drills   map[string]string // "drill.core_casts" -> text
}

type frontMatter struct {
	spec, role, status, version, reviewed string
	sources                               []string
}

// ParsePersona reads and checks one persona file (§8).
func ParsePersona(data []byte) (Persona, error) {
	fm, body, err := splitFrontMatter(string(data))
	if err != nil {
		return Persona{}, err
	}
	sections, err := splitSections(body)
	if err != nil {
		return Persona{}, err
	}
	drills, err := parseDrills(sections[sectionDrills], fm.role)
	if err != nil {
		return Persona{}, err
	}
	if err := checkProse(sections); err != nil {
		return Persona{}, err
	}
	sum := sha256.Sum256(data)
	return Persona{
		Spec:     fm.spec,
		Status:   PersonaStatus(fm.status),
		Version:  personaVersionPrefix + hex.EncodeToString(sum[:]),
		Sources:  fm.sources,
		Reviewed: fm.reviewed,
		Body:     strings.TrimSpace(body),
		Drills:   drills,
	}, nil
}

// LoadPersona reads dir/<spec>.md and checks that it describes that spec.
func LoadPersona(dir, spec string) (Persona, error) {
	path := filepath.Join(dir, spec+".md")
	data, err := os.ReadFile(path)
	if err != nil {
		return Persona{}, fmt.Errorf("read persona %s: %w", path, err)
	}
	p, err := ParsePersona(data)
	if err != nil {
		return Persona{}, fmt.Errorf("persona %s: %w", path, err)
	}
	if p.Spec != spec {
		return Persona{}, fmt.Errorf("persona %s: front matter spec %q does not match the file name", path, p.Spec)
	}
	return p, nil
}

// Personas caches loaded personas by spec; a changed file on disk is a new version only
// after a restart.
type Personas struct {
	Dir   string
	mu    sync.Mutex
	cache map[string]Persona
}

// For returns the persona for spec, loading it on first use.
func (p *Personas) For(spec string) (Persona, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if cached, ok := p.cache[spec]; ok {
		return cached, nil
	}
	loaded, err := LoadPersona(p.Dir, spec)
	if err != nil {
		return Persona{}, err
	}
	if p.cache == nil {
		p.cache = map[string]Persona{}
	}
	p.cache[spec] = loaded
	return loaded, nil
}

func splitFrontMatter(text string) (frontMatter, string, error) {
	lines := strings.Split(text, "\n")
	if strings.TrimSpace(lines[0]) != frontMatterFence {
		return frontMatter{}, "", fmt.Errorf("no front matter")
	}
	for i := 1; i < len(lines); i++ {
		if strings.TrimSpace(lines[i]) == frontMatterFence {
			fm, err := parseFrontMatter(lines[1:i])
			return fm, strings.Join(lines[i+1:], "\n"), err
		}
	}
	return frontMatter{}, "", fmt.Errorf("front matter is not closed")
}

func parseFrontMatter(lines []string) (frontMatter, error) {
	scalars := map[string]string{}
	var sources []string
	inSources := false
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		switch {
		case trimmed == "":
		case inSources && strings.HasPrefix(trimmed, bulletPrefix):
			sources = append(sources, strings.TrimSpace(strings.TrimPrefix(trimmed, bulletPrefix)))
		default:
			key, value, ok := strings.Cut(line, ":")
			if !ok {
				return frontMatter{}, fmt.Errorf("front matter line %q is not key: value", line)
			}
			key = strings.TrimSpace(key)
			inSources = key == "sources"
			scalars[key] = strings.TrimSpace(value)
		}
	}
	return buildFrontMatter(scalars, sources)
}

func buildFrontMatter(scalars map[string]string, sources []string) (frontMatter, error) {
	for _, key := range frontMatterKeys {
		if _, ok := scalars[key]; !ok {
			return frontMatter{}, fmt.Errorf("front matter is missing %q", key)
		}
	}
	fm := frontMatter{
		spec: scalars["spec"], role: scalars["role"], status: scalars["status"],
		version: scalars["version"], reviewed: scalars["reviewed"], sources: sources,
	}
	if fm.spec == "" || fm.reviewed == "" || fm.version == "" {
		return frontMatter{}, fmt.Errorf("front matter has an empty spec, version or reviewed")
	}
	if PersonaStatus(fm.status) != PersonaFull && PersonaStatus(fm.status) != PersonaStarter {
		return frontMatter{}, fmt.Errorf("unknown status %q", fm.status)
	}
	if _, ok := Precedence[fm.role]; !ok {
		return frontMatter{}, fmt.Errorf("unknown role %q", fm.role)
	}
	return fm, nil
}

// splitSections returns each section's text by name, requiring the four in order.
func splitSections(body string) (map[string]string, error) {
	var names []string
	texts := map[string][]string{}
	for _, line := range strings.Split(body, "\n") {
		if name, ok := strings.CutPrefix(line, sectionPrefix); ok {
			names = append(names, strings.TrimSpace(name))
			continue
		}
		if len(names) == 0 {
			if strings.TrimSpace(line) != "" {
				return nil, fmt.Errorf("text before the first section")
			}
			continue
		}
		last := names[len(names)-1]
		texts[last] = append(texts[last], line)
	}
	if strings.Join(names, "|") != strings.Join(personaSections, "|") {
		return nil, fmt.Errorf("sections are %q, want %q in order", names, personaSections)
	}
	out := map[string]string{}
	for name, lines := range texts {
		out[name] = strings.Join(lines, "\n")
	}
	return out, nil
}

func parseDrills(section, role string) (map[string]string, error) {
	drills := map[string]string{}
	for _, line := range strings.Split(section, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		id, text, err := parseDrillLine(line)
		if err != nil {
			return nil, err
		}
		if !isDrillCode(id) {
			return nil, fmt.Errorf("drill id %q is not %s plus a finding code", id, drillIDPrefix)
		}
		if _, dup := drills[id]; dup {
			return nil, fmt.Errorf("drill %q appears twice", id)
		}
		drills[id] = text
	}
	for _, code := range Precedence[role] {
		if id := drillIDPrefix + string(code); drills[id] == "" {
			return nil, fmt.Errorf("missing drill %q for role %s", id, role)
		}
	}
	return drills, nil
}

func parseDrillLine(line string) (id, text string, err error) {
	rest, ok := strings.CutPrefix(line, bulletPrefix)
	if !ok {
		return "", "", fmt.Errorf("drills line %q is not a bullet", line)
	}
	id, text, ok = strings.Cut(rest, ":")
	id, text = strings.TrimSpace(id), strings.TrimSpace(text)
	if !ok || text == "" {
		return "", "", fmt.Errorf("drills line %q is not - drill.<code>: <text>", line)
	}
	if n := len([]rune(text)); n > maxDrillChars {
		return "", "", fmt.Errorf("drill %q is %d characters, over %d", id, n, maxDrillChars)
	}
	return id, text, nil
}

// isDrillCode reports whether id is drill. plus one of the eight finding codes.
func isDrillCode(id string) bool {
	code, ok := strings.CutPrefix(id, drillIDPrefix)
	if !ok {
		return false
	}
	for _, c := range Precedence[RoleDPS] { // the DPS list holds all eight codes
		if string(c) == code {
			return true
		}
	}
	return false
}

// checkProse enforces §8.2: no digits outside Drills, no exclamation mark anywhere, and
// no banned tone word except in the Voice section's rule lines.
func checkProse(sections map[string]string) error {
	for _, name := range personaSections {
		text := sections[name]
		if strings.Contains(text, "!") {
			return fmt.Errorf("section %q has an exclamation mark", name)
		}
		if name != sectionDrills && strings.IndexFunc(text, unicode.IsDigit) >= 0 {
			return fmt.Errorf("section %q has a digit; prose quotes no numbers", name)
		}
		if bad := toneProblems(toneChecked(name, text)); len(bad) > 0 {
			return fmt.Errorf("section %q uses banned tone %q", name, bad)
		}
	}
	return nil
}

// toneChecked is the part of a section the tone list applies to: all of it, except the
// Voice section's "- " rule lines.
func toneChecked(name, text string) string {
	if name != sectionVoice {
		return text
	}
	var kept []string
	for _, line := range strings.Split(text, "\n") {
		if !strings.HasPrefix(line, bulletPrefix) {
			kept = append(kept, line)
		}
	}
	return strings.Join(kept, "\n")
}
