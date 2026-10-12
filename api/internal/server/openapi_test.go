// api/internal/server/openapi_test.go
package server

import (
	"os"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestOpenAPIListsEveryRoute(t *testing.T) {
	b, err := os.ReadFile("../../openapi.yaml")
	if err != nil {
		t.Fatal(err)
	}

	var doc map[string]any
	err = yaml.Unmarshal(b, &doc)
	if err != nil {
		t.Fatalf("openapi.yaml is not valid YAML: %v", err)
	}

	paths, ok := doc["paths"].(map[string]any)
	if !ok {
		t.Fatal("openapi.yaml missing or invalid 'paths' object")
	}

	requiredPaths := []string{
		"/health", "/version",
		"/v1/subscribe", "/v1/subscribe/confirm", "/v1/subscribe/unsubscribe",
		"/v1/builds", "/v1/builds/{id}", "/b/{id}", "/b/{id}/card.png",
		"/v1/auth/battlenet/start", "/v1/auth/battlenet/callback",
		"/v1/auth/email", "/v1/auth/email/callback", "/v1/auth/logout", "/v1/sessions", "/v1/me",
		"/v1/devices/pair", "/v1/devices/claim", "/v1/devices", "/v1/devices/{id}",
		"/v1/reports", "/v1/reports/{id}", "/v1/guilds/{id}/reports", "/v1/reports/{id}/visibility",
		"/v1/reports/{id}/access", "/v1/reports/{id}/files/{path}",
		"/v1/reports/{id}/fights/{n}", "/v1/reports/{id}/fights/{n}/live",
		"/v1/reports/{id}/raw", "/v1/reports/{id}/complete",
		"/v1/uploads", "/v1/uploads/{upload_id}/complete",
		"/v1/rankings", "/v1/rankings/percentile", "/v1/rankings/guilds",
		"/v1/characters/{region}/{ruleset}/{name}", "/v1/guilds/{region}/{ruleset}/{name}",
		"/v1/addon/exports", "/v1/addon/inbox", "/reports/{id}/card.png",
		"/v1/sims", "/v1/sims/{id}", "/v1/sims/{id}/progress", "/v1/sims/run",
		"/v1/specs", "/v1/characters/{region}/{ruleset}/{name}/sim-input",
		"/v1/phases",
	}
	for _, p := range requiredPaths {
		if _, ok := paths[p]; !ok {
			t.Errorf("openapi.yaml missing path %s", p)
		}
	}

	components, ok := doc["components"].(map[string]any)
	if !ok {
		t.Fatal("openapi.yaml missing 'components'")
	}
	schemas, ok := components["schemas"].(map[string]any)
	if !ok {
		t.Fatal("openapi.yaml missing 'components.schemas'")
	}
	for _, s := range []string{"Envelope", "Build", "BuildInput", "Report", "FightEntry",
		"MetricsRow", "RankingRow", "SimResult", "SimRequest", "SimRow", "SpecFidelity", "User"} {
		if _, ok := schemas[s]; !ok {
			t.Errorf("openapi.yaml missing schema %s", s)
		}
	}
}

// Task L6.7: the character page's gear panel and execution_score are documented, with every
// null rule.
func TestOpenAPIDocumentsTheGearPanelAndExecutionScore(t *testing.T) {
	b, err := os.ReadFile("../../openapi.yaml")
	if err != nil {
		t.Fatal(err)
	}
	var doc struct {
		Paths      map[string]map[string]map[string]any `yaml:"paths"`
		Components struct {
			Schemas map[string]map[string]any `yaml:"schemas"`
		} `yaml:"components"`
	}
	if err := yaml.Unmarshal(b, &doc); err != nil {
		t.Fatal(err)
	}
	desc := func(v any) string {
		m, _ := v.(map[string]any)
		s, _ := m["description"].(string)
		return s
	}
	props := func(schema map[string]any) map[string]any {
		p, _ := schema["properties"].(map[string]any)
		return p
	}

	path := doc.Paths["/v1/characters/{region}/{ruleset}/{name}"]["get"]
	for _, want := range []string{"gear", "gear_gap_total", "CharacterGear", "ExecutionScore"} {
		if !strings.Contains(desc(path), want) {
			t.Errorf("the character path's description does not name %q", want)
		}
	}

	gear, ok := doc.Components.Schemas["CharacterGear"]
	if !ok {
		t.Fatal("no CharacterGear schema")
	}
	total := desc(props(gear)["gear_gap_total"])
	for _, want := range []string{
		"score points the band's picks would add over what the character wears, summed over the slots",
		"not the Gear component's distance to its top, which is measured against the best set of the band across presets and factions",
		"null",
	} {
		if !strings.Contains(strings.Join(strings.Fields(total), " "), want) {
			t.Errorf("gear_gap_total's description lacks %q: %q", want, total)
		}
	}
	slot, ok := doc.Components.Schemas["CharacterGearSlot"]
	if !ok {
		t.Fatal("no CharacterGearSlot schema")
	}
	p := props(slot)
	for _, field := range []string{"worn", "pick", "gap"} {
		m, _ := p[field].(map[string]any)
		if m["nullable"] != true || !strings.Contains(desc(m), "ull") {
			t.Errorf("%s must be nullable with its null rule: %v", field, m)
		}
	}
	pick, _ := p["pick"].(map[string]any)
	chance, _ := props(pick)["chance"].(map[string]any)
	if chance["nullable"] != true || desc(chance) == "" {
		t.Errorf("pick.chance must be nullable with its null rule: %v", chance)
	}
	worn, _ := p["worn"].(map[string]any)
	for _, field := range []string{"name", "item_level"} {
		m, _ := props(worn)[field].(map[string]any)
		if m["nullable"] != true {
			t.Errorf("worn.%s must be nullable: %v", field, m)
		}
	}

	const plain = "share of what the same gear and talents do in the simulator"
	exec, ok := doc.Components.Schemas["ExecutionScore"]
	if !ok || !strings.Contains(strings.Join(strings.Fields(desc(exec)), " "), plain) {
		t.Errorf("ExecutionScore must say %q: %v", plain, exec)
	}
	row, _ := props(doc.Components.Schemas["RankingRow"])["execution_score"].(map[string]any)
	if row["$ref"] != "#/components/schemas/ExecutionScore" {
		t.Errorf("RankingRow.execution_score must $ref ExecutionScore: %v", row)
	}
	// Fix round 1, item 7: every null rule, in plain words.
	words := func(v any) string { return strings.ToLower(strings.Join(strings.Fields(desc(v)), " ")) }
	for _, want := range []string{"not been simulated yet", "healers and tanks", "not validated",
		"could not rebuild", "produced no damage", "no row for the player"} {
		if !strings.Contains(words(exec), want) {
			t.Errorf("ExecutionScore's null rules lack %q: %q", want, desc(exec))
		}
	}
	// Fix round 1, items 2, 3, 4, 5 and 9: the gap's pair and two-hander rules, its rounding,
	// the promoted pick's chance and the faction's fallback.
	gap, _ := p["gap"].(map[string]any)
	for _, want := range []string{"other slot of its pair", "two-handed", "as one unit", "under 0.05"} {
		if !strings.Contains(words(gap), want) {
			t.Errorf("gap's description lacks %q: %q", want, desc(gap))
		}
	}
	if !strings.Contains(words(pick), "battle.net") {
		t.Errorf("pick's description must name the faction's Battle.net fallback: %q", desc(pick))
	}
	if !strings.Contains(words(chance), "its own source") || strings.Contains(words(chance), "took the place") {
		t.Errorf("chance must say a promoted pick carries its own source's chance: %q", desc(chance))
	}
}
