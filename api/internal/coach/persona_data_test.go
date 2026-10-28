package coach

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// personaDir is where the shipped persona files live, relative to this package.
const personaDir = "../../../data/curated/coach"

// repoRoot is the repository root, relative to this package; persona sources are repo paths.
const repoRoot = "../../.."

// fullPersonaSpecs are the specs whose personas ship full at launch (§8.3).
var fullPersonaSpecs = []string{"warrior-fury", "priest-holy", "warrior-protection"}

const reviewPending = "pending"

func TestFullPersonas_ParseAndAreFull(t *testing.T) {
	for _, spec := range fullPersonaSpecs {
		p, err := LoadPersona(personaDir, spec)
		if err != nil {
			t.Fatalf("%s: %v", spec, err)
		}
		if p.Status != PersonaFull {
			t.Errorf("%s: status %q, want %q", spec, p.Status, PersonaFull)
		}
		if len(p.Sources) == 0 {
			t.Errorf("%s: no sources listed", spec)
		}
		for _, src := range p.Sources {
			if _, err := os.Stat(filepath.Join(repoRoot, src)); err != nil {
				t.Errorf("%s: source %s does not exist", spec, src)
			}
		}
	}
}

// TestFullPersonas_Reviewed fails on purpose until wow-player has read each file
// and the controller has recorded `reviewed: wow-player, <date>` (§8.2).
func TestFullPersonas_Reviewed(t *testing.T) {
	for _, spec := range fullPersonaSpecs {
		p, err := LoadPersona(personaDir, spec)
		if err != nil {
			t.Fatalf("%s: %v", spec, err)
		}
		if p.Reviewed == "" || p.Reviewed == reviewPending {
			t.Errorf("%s: reviewed is %q; a wow-player read must be recorded before it ships", spec, p.Reviewed)
		}
	}
}

func TestPersonas_RoleIsTheSpecsJSONRole(t *testing.T) {
	roles := map[string]string{}
	for _, s := range loadSpecs(t) {
		roles[s.Spec] = s.Role
	}
	for _, spec := range fullPersonaSpecs {
		want, ok := roles[spec]
		if !ok {
			t.Fatalf("%s: not in specs.json", spec)
		}
		got, err := personaFileRole(filepath.Join(personaDir, spec+".md"))
		if err != nil {
			t.Fatalf("%s: %v", spec, err)
		}
		if got != want {
			t.Errorf("%s: front-matter role %q, specs.json says %q", spec, got, want)
		}
	}
}

// personaFileRole reads the `role:` line of a persona file's front matter.
// Persona carries no role field (Task 22), so the data test reads it directly.
func personaFileRole(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read persona %s: %w", path, err)
	}
	lines := strings.Split(string(data), "\n")
	if len(lines) == 0 || strings.TrimSpace(lines[0]) != "---" {
		return "", fmt.Errorf("persona %s: no front matter", path)
	}
	for _, line := range lines[1:] {
		if strings.TrimSpace(line) == "---" {
			break
		}
		if role, ok := strings.CutPrefix(line, "role:"); ok {
			return strings.TrimSpace(role), nil
		}
	}
	return "", fmt.Errorf("persona %s: front matter has no role", path)
}

// TestPersonas_NoRealPeople needs Task 27's dictionaries; Task 27 adds it to this file (F7).

const starterStatus = "starter"

// TestAllPersonas_Parse: every spec in specs.json has a persona that loads, whose
// front-matter role is specs.json's role; three are full and the rest starters.
func TestAllPersonas_Parse(t *testing.T) {
	full := map[string]bool{}
	for _, spec := range fullPersonaSpecs {
		full[spec] = true
	}
	var fullCount, starterCount int
	for _, s := range loadSpecs(t) {
		p, err := LoadPersona(personaDir, s.Spec)
		if err != nil {
			t.Fatalf("%s: %v", s.Spec, err)
		}
		role, err := personaFileRole(filepath.Join(personaDir, s.Spec+".md"))
		if err != nil || role != s.Role {
			t.Errorf("%s: front-matter role %q (err %v), specs.json says %q", s.Spec, role, err, s.Role)
		}
		switch {
		case full[s.Spec] && p.Status == PersonaFull:
			fullCount++
		case !full[s.Spec] && string(p.Status) == starterStatus:
			starterCount++
		default:
			t.Errorf("%s: status %q does not fit its place in the launch set", s.Spec, p.Status)
		}
	}
	if fullCount != len(fullPersonaSpecs) || starterCount != 25 {
		t.Errorf("got %d full and %d starter personas, want %d and 25", fullCount, starterCount, len(fullPersonaSpecs))
	}
}

// TestStarterPersonas_Reviewed fails on purpose until wow-player has read each starter
// and the controller has recorded the read in `reviewed` (§8.2).
func TestStarterPersonas_Reviewed(t *testing.T) {
	for _, s := range loadSpecs(t) {
		p, err := LoadPersona(personaDir, s.Spec)
		if err != nil {
			t.Fatalf("%s: %v", s.Spec, err)
		}
		if string(p.Status) == starterStatus && (p.Reviewed == "" || p.Reviewed == reviewPending) {
			t.Errorf("%s: reviewed is %q; a wow-player read must be recorded before it ships", s.Spec, p.Reviewed)
		}
	}
}
