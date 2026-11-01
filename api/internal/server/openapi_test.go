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
		"/v1/auth/app/start", "/v1/auth/app/continue", "/v1/auth/app/exchange", "/v1/auth/app/logout",
		"/v1/auth/email", "/v1/auth/email/callback", "/v1/auth/logout", "/v1/sessions", "/v1/me",
		"/v1/devices/pair", "/v1/devices/claim", "/v1/devices/connect", "/v1/devices/connect/{connect_id}", "/v1/devices/connect/lookup", "/v1/devices/connect/approve", "/v1/devices/connect/deny", "/v1/devices", "/v1/devices/self", "/v1/devices/{id}",
		"/v1/reports", "/v1/reports/{id}", "/v1/guilds/{id}/reports", "/v1/reports/{id}/visibility",
		"/v1/reports/{id}/access", "/v1/reports/{id}/files/{path}",
		"/v1/reports/{id}/fights/{n}", "/v1/reports/{id}/fights/{n}/live",
		"/v1/reports/{id}/raw", "/v1/reports/{id}/complete",
		"/v1/uploads", "/v1/uploads/{upload_id}/complete",
		"/v1/rankings", "/v1/rankings/percentile", "/v1/rankings/guilds",
		"/v1/characters/search", "/v1/characters/{region}/{ruleset}/{name}", "/v1/guilds/{region}/{ruleset}/{name}",
		"/v1/addon/exports", "/v1/addon/inbox", "/reports/{id}/card.png",
		"/v1/sims", "/v1/sims/{id}", "/v1/sims/{id}/progress", "/v1/sims/run",
		"/v1/sims/{id}/seen", "/v1/sims/{id}/rerun",
		"/v1/specs", "/v1/characters/{region}/{ruleset}/{name}/sim-input",
		"/v1/characters/{region}/{ruleset}/{name}/history",
		"/v1/phases",
		"/v1/me/widget", "/v1/me/live-activities", "/v1/me/live-activities/{id}",
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
		Paths      map[string]map[string]any `yaml:"paths"`
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

	path, _ := doc.Paths["/v1/characters/{region}/{ruleset}/{name}"]["get"].(map[string]any)
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
		if !nullable(m) || !strings.Contains(desc(m), "ull") {
			t.Errorf("%s must be nullable with its null rule: %v", field, m)
		}
	}
	pick, _ := p["pick"].(map[string]any)
	chance, _ := props(pick)["chance"].(map[string]any)
	if !nullable(chance) || desc(chance) == "" {
		t.Errorf("pick.chance must be nullable with its null rule: %v", chance)
	}
	worn, _ := p["worn"].(map[string]any)
	for _, field := range []string{"name", "item_level"} {
		m, _ := props(worn)[field].(map[string]any)
		if !nullable(m) {
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

// Task L6.8: the guild roster's attendance is defined in plain words (the window is the
// guild's last 8 non-private reports with a boss fight, presence is read by character key),
// and roster_best rows name ExecutionScore like every other row that carries it.
func TestOpenAPIDocumentsGuildAttendanceAndRosterBestExecutionScore(t *testing.T) {
	b, err := os.ReadFile("../../openapi.yaml")
	if err != nil {
		t.Fatal(err)
	}
	var doc struct {
		Paths      map[string]map[string]any `yaml:"paths"`
		Components struct {
			Schemas map[string]map[string]any `yaml:"schemas"`
		} `yaml:"components"`
	}
	if err := yaml.Unmarshal(b, &doc); err != nil {
		t.Fatal(err)
	}
	words := func(v any) string {
		m, _ := v.(map[string]any)
		s, _ := m["description"].(string)
		return strings.ToLower(strings.Join(strings.Fields(s), " "))
	}

	home, ok := doc.Paths["/v1/guilds/{id}/home"]["get"].(map[string]any)
	if !ok || !strings.Contains(words(home), "guildattendance") {
		t.Errorf("/v1/guilds/{id}/home must be listed and point roster[].attendance at GuildAttendance: %v", home)
	}
	att, ok := doc.Components.Schemas["GuildAttendance"]
	if !ok {
		t.Fatal("no GuildAttendance schema")
	}
	// Fix round 1, items 2 and 4: the window is non-private, unflagged reports someone is
	// present in, and the home path lists its cursor and its 400 and 404.
	for _, want := range []string{"unflagged", "ranked or rated row"} {
		if !strings.Contains(words(att), want) {
			t.Errorf("GuildAttendance's description lacks %q: %q", want, words(att))
		}
	}
	params, _ := home["parameters"].([]any)
	hasCursor := false
	for _, p := range params {
		if m, _ := p.(map[string]any); m["name"] == "cursor" && m["in"] == "query" {
			hasCursor = true
		}
	}
	if !hasCursor {
		t.Errorf("/v1/guilds/{id}/home must list its cursor query parameter: %v", params)
	}
	responses, _ := home["responses"].(map[string]any)
	for _, code := range []string{"400", "404"} {
		if _, ok := responses[code]; !ok {
			t.Errorf("/v1/guilds/{id}/home must list its %s: %v", code, responses)
		}
	}
	for _, want := range []string{"last 8", "non-private", "character key", "guids"} {
		if !strings.Contains(words(att), want) {
			t.Errorf("GuildAttendance's description lacks %q: %q", want, words(att))
		}
	}
	props, _ := att["properties"].(map[string]any)
	for _, field := range []string{"present", "nights"} {
		if words(props[field]) == "" {
			t.Errorf("GuildAttendance.%s has no plain-words definition", field)
		}
	}

	if !strings.Contains(words(doc.Components.Schemas["ExecutionScore"]), "roster_best") {
		t.Errorf("ExecutionScore must name the guild page's roster_best rows: %q", words(doc.Components.Schemas["ExecutionScore"]))
	}
	guild, _ := doc.Paths["/v1/guilds/{region}/{ruleset}/{name}"]["get"].(map[string]any)
	if !strings.Contains(words(guild), "executionscore") {
		t.Errorf("the guild path's roster_best description must point execution_score at ExecutionScore: %q", words(guild))
	}
}

// nullable reports whether an OpenAPI 3.1 schema admits null: "null" among its types.
func nullable(schema map[string]any) bool {
	switch t := schema["type"].(type) {
	case []any:
		for _, v := range t {
			if v == "null" {
				return true
			}
		}
	case string:
		return t == "null"
	}
	return false
}
