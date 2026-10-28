package coach

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"regexp"
	"slices"
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
	noLine               = 0
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

// specKeyRE is the shape of a specs.json key; it keeps path separators out of file names.
var specKeyRE = regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)*$`)

// PersonaError is a refused persona file: which file, which line, and why.
type PersonaError struct {
	File string // empty until LoadPersona fills it in
	Line int    // 1-based; noLine when the problem is the file as a whole
	Msg  string
}

func (e *PersonaError) Error() string {
	file := e.File
	if file == "" {
		file = "persona"
	}
	if e.Line == noLine {
		return fmt.Sprintf("%s: %s", file, e.Msg)
	}
	return fmt.Sprintf("%s:%d: %s", file, e.Line, e.Msg)
}

func personaErr(line int, format string, args ...any) *PersonaError {
	return &PersonaError{Line: line, Msg: fmt.Sprintf(format, args...)}
}

// Persona is one spec's coach persona: the system-prompt prose and the drill catalogue.
type Persona struct {
	Spec     string
	Status   PersonaStatus
	Version  string // personaVersionPrefix + hex sha256 of the normalised file (§8.3)
	Sources  []string
	Reviewed string
	Body     string            // the four sections, as the system prompt carries them
	Drills   map[string]string // "drill.core_casts" -> text
}

// clone copies the slice and map so a caller cannot mutate a cached persona.
func (p Persona) clone() Persona {
	p.Sources = slices.Clone(p.Sources)
	p.Drills = maps.Clone(p.Drills)
	return p
}

type frontMatter struct {
	spec, role, status, version, reviewed string
	sources                               []string
	lines                                 map[string]int // key -> line number
}

type numberedLine struct {
	n    int
	text string
}

type section struct {
	line  int // the "## " heading's line
	lines []numberedLine
}

// normalise makes the hashed bytes stable across checkouts: LF line endings, no
// trailing whitespace on a line, exactly one final newline.
func normalise(data []byte) string {
	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	for i, l := range lines {
		lines[i] = strings.TrimRight(l, " \t\r")
	}
	return strings.TrimRight(strings.Join(lines, "\n"), "\n") + "\n"
}

// ParsePersona reads and checks one persona file (§8). The checks on digits and
// exclamation marks cover the body only: front matter holds source paths with build
// numbers and a review date, and is not prose. Every refusal is a *PersonaError.
func ParsePersona(data []byte) (Persona, error) {
	text := normalise(data)
	fm, bodyLines, err := splitFrontMatter(text)
	if err != nil {
		return Persona{}, err
	}
	sections, err := splitSections(bodyLines)
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
	sum := sha256.Sum256([]byte(text))
	return Persona{
		Spec:     fm.spec,
		Status:   PersonaStatus(fm.status),
		Version:  personaVersionPrefix + hex.EncodeToString(sum[:]),
		Sources:  fm.sources,
		Reviewed: fm.reviewed,
		Body:     bodyText(bodyLines),
		Drills:   drills,
	}, nil
}

func bodyText(lines []numberedLine) string {
	texts := make([]string, len(lines))
	for i, l := range lines {
		texts[i] = l.text
	}
	return strings.TrimSpace(strings.Join(texts, "\n"))
}

// LoadPersona reads dir/<spec>.md and checks that it describes that spec.
func LoadPersona(dir, spec string) (Persona, error) {
	if !specKeyRE.MatchString(spec) {
		return Persona{}, fmt.Errorf("persona spec %q is not a specs.json key", spec)
	}
	path := filepath.Join(dir, spec+".md")
	data, err := os.ReadFile(path)
	if err != nil {
		return Persona{}, fmt.Errorf("read persona %s: %w", path, err)
	}
	p, fm, err := parseWithFrontMatter(data)
	if err == nil && p.Spec != spec {
		err = personaErr(fm.lines["spec"], "front matter spec %q does not match the file name %q", p.Spec, spec)
	}
	if err != nil {
		var pe *PersonaError
		if errors.As(err, &pe) {
			named := *pe
			named.File = path
			return Persona{}, &named
		}
		return Persona{}, err
	}
	return p, nil
}

// parseWithFrontMatter is ParsePersona that also returns the parsed front matter.
func parseWithFrontMatter(data []byte) (Persona, frontMatter, error) {
	p, err := ParsePersona(data)
	if err != nil {
		return Persona{}, frontMatter{}, err
	}
	fm, _, _ := splitFrontMatter(normalise(data))
	return p, fm, nil
}

// Personas caches loaded personas by spec; a changed file on disk is a new version only
// after a restart.
type Personas struct {
	Dir   string
	mu    sync.Mutex
	cache map[string]Persona
}

// For returns a copy of the persona for spec, loading it on first use. The disk read
// happens outside the lock; a racing load of the same spec keeps whichever stored first.
func (p *Personas) For(spec string) (Persona, error) {
	if cached, ok := p.cached(spec); ok {
		return cached, nil
	}
	loaded, err := LoadPersona(p.Dir, spec)
	if err != nil {
		return Persona{}, err
	}
	return p.store(spec, loaded), nil
}

func (p *Personas) cached(spec string) (Persona, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	cached, ok := p.cache[spec]
	return cached.clone(), ok
}

func (p *Personas) store(spec string, loaded Persona) Persona {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.cache == nil {
		p.cache = map[string]Persona{}
	}
	if existing, ok := p.cache[spec]; ok {
		return existing.clone()
	}
	p.cache[spec] = loaded
	return loaded.clone()
}

func splitFrontMatter(text string) (frontMatter, []numberedLine, error) {
	lines := strings.Split(text, "\n")
	if lines[0] != frontMatterFence {
		return frontMatter{}, nil, personaErr(1, "no front matter")
	}
	for i := 1; i < len(lines); i++ {
		if lines[i] == frontMatterFence {
			fm, err := parseFrontMatter(lines[1:i])
			return fm, numbered(lines[i+1:], i+2), err
		}
	}
	return frontMatter{}, nil, personaErr(1, "front matter is not closed")
}

func numbered(lines []string, first int) []numberedLine {
	out := make([]numberedLine, len(lines))
	for i, l := range lines {
		out[i] = numberedLine{n: first + i, text: l}
	}
	return out
}

func parseFrontMatter(lines []string) (frontMatter, error) {
	scalars := map[string]string{}
	at := map[string]int{}
	var sources []string
	inSources := false
	for i, line := range lines {
		n := i + 2 // line 1 is the opening fence
		trimmed := strings.TrimSpace(line)
		switch {
		case trimmed == "":
		case inSources && strings.HasPrefix(trimmed, bulletPrefix):
			sources = append(sources, strings.TrimSpace(strings.TrimPrefix(trimmed, bulletPrefix)))
		default:
			key, value, ok := strings.Cut(line, ":")
			if !ok {
				return frontMatter{}, personaErr(n, "front matter line %q is not key: value", line)
			}
			key = strings.TrimSpace(key)
			inSources = key == "sources"
			scalars[key] = strings.TrimSpace(value)
			at[key] = n
		}
	}
	return buildFrontMatter(scalars, at, sources)
}

func buildFrontMatter(scalars map[string]string, at map[string]int, sources []string) (frontMatter, error) {
	for _, key := range frontMatterKeys {
		if _, ok := scalars[key]; !ok {
			return frontMatter{}, personaErr(noLine, "front matter is missing %q", key)
		}
	}
	fm := frontMatter{
		spec: scalars["spec"], role: scalars["role"], status: scalars["status"],
		version: scalars["version"], reviewed: scalars["reviewed"], sources: sources, lines: at,
	}
	for _, key := range []string{"spec", "version", "reviewed"} {
		if scalars[key] == "" {
			return frontMatter{}, personaErr(at[key], "front matter %q is empty", key)
		}
	}
	if PersonaStatus(fm.status) != PersonaFull && PersonaStatus(fm.status) != PersonaStarter {
		return frontMatter{}, personaErr(at["status"], "unknown status %q", fm.status)
	}
	if _, ok := Precedence[fm.role]; !ok {
		return frontMatter{}, personaErr(at["role"], "unknown role %q", fm.role)
	}
	return fm, nil
}

// splitSections returns each section by name, requiring the four in order.
func splitSections(body []numberedLine) (map[string]section, error) {
	out := map[string]section{}
	var current string
	for _, l := range body {
		if name, ok := strings.CutPrefix(l.text, sectionPrefix); ok {
			name = strings.TrimSpace(name)
			want := ""
			if len(out) < len(personaSections) {
				want = personaSections[len(out)]
			}
			if name != want {
				return nil, personaErr(l.n, "section %q where %q is due; the order is %q", name, want, personaSections)
			}
			current = name
			out[name] = section{line: l.n}
			continue
		}
		if current == "" {
			if strings.TrimSpace(l.text) != "" {
				return nil, personaErr(l.n, "text before the first section")
			}
			continue
		}
		s := out[current]
		s.lines = append(s.lines, l)
		out[current] = s
	}
	if len(out) != len(personaSections) {
		return nil, personaErr(noLine, "has %d of the sections %q", len(out), personaSections)
	}
	return out, nil
}

func parseDrills(s section, role string) (map[string]string, error) {
	drills := map[string]string{}
	for _, l := range s.lines {
		if strings.TrimSpace(l.text) == "" {
			continue
		}
		id, text, err := parseDrillLine(l)
		if err != nil {
			return nil, err
		}
		if !isDrillCode(id) {
			return nil, personaErr(l.n, "drill id %q is not %s plus a finding code", id, drillIDPrefix)
		}
		if _, dup := drills[id]; dup {
			return nil, personaErr(l.n, "drill %q appears twice", id)
		}
		drills[id] = text
	}
	for _, code := range Precedence[role] {
		if id := drillIDPrefix + string(code); drills[id] == "" {
			return nil, personaErr(s.line, "missing drill %q for role %s", id, role)
		}
	}
	return drills, nil
}

func parseDrillLine(l numberedLine) (id, text string, err error) {
	rest, ok := strings.CutPrefix(l.text, bulletPrefix)
	if !ok {
		return "", "", personaErr(l.n, "drills line %q is not a bullet", l.text)
	}
	id, text, ok = strings.Cut(rest, ":")
	id, text = strings.TrimSpace(id), strings.TrimSpace(text)
	if !ok || text == "" {
		return "", "", personaErr(l.n, "drills line %q is not - drill.<code>: <text>", l.text)
	}
	if n := len([]rune(text)); n > maxDrillChars {
		return "", "", personaErr(l.n, "drill %q is %d characters, over %d", id, n, maxDrillChars)
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

// checkProse enforces §8.2 on the body: no digits outside Drills, no exclamation mark,
// and no banned tone word except in the Voice section's rule lines. The tone check reads
// each paragraph as one run too, so a phrase split across a line break is still caught.
func checkProse(sections map[string]section) error {
	for _, name := range personaSections {
		lines := sections[name].lines
		for _, l := range lines {
			if err := checkProseLine(name, l); err != nil {
				return err
			}
		}
		for _, p := range tonedParagraphs(name, lines) {
			if bad := toneProblems(p.text); len(bad) > 0 {
				return personaErr(p.n, "banned tone %q in section %q", bad, name)
			}
		}
	}
	return nil
}

// tonedParagraphs joins each run of consecutive non-blank lines into one line (numbered by
// its first line), leaving out the Voice section's rule lines, which may name a rule.
func tonedParagraphs(name string, lines []numberedLine) []numberedLine {
	var out []numberedLine
	var cur []string
	first := 0
	flush := func() {
		if len(cur) > 0 {
			out = append(out, numberedLine{n: first, text: strings.Join(cur, " ")})
			cur = nil
		}
	}
	for _, l := range lines {
		text := strings.TrimSpace(l.text)
		if text == "" || isVoiceRule(name, l) {
			flush()
			continue
		}
		if len(cur) == 0 {
			first = l.n
		}
		cur = append(cur, text)
	}
	flush()
	return out
}

func checkProseLine(name string, l numberedLine) error {
	if strings.Contains(l.text, "!") {
		return personaErr(l.n, "exclamation mark in section %q", name)
	}
	if name != sectionDrills && strings.IndexFunc(l.text, unicode.IsDigit) >= 0 {
		return personaErr(l.n, "digit in section %q; prose quotes no numbers", name)
	}
	if isVoiceRule(name, l) {
		return nil
	}
	if bad := toneProblems(l.text); len(bad) > 0 {
		return personaErr(l.n, "banned tone %q in section %q", bad, name)
	}
	return nil
}

// isVoiceRule reports a Voice section bullet, which may name a rule in its own words.
func isVoiceRule(name string, l numberedLine) bool {
	return name == sectionVoice && strings.HasPrefix(l.text, bulletPrefix)
}
