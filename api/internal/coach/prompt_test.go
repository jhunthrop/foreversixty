package coach

import (
	"encoding/json"
	"flag"
	"os"
	"reflect"
	"strings"
	"testing"
)

var updatePrompt = flag.Bool("update-prompt", false, "rewrite the golden prompt")

const (
	goldenPromptFile  = "testdata/prompt_fury.golden"
	goldenVersionFile = "testdata/prompt_fury.version"
	fixtureModelSlug  = "deepseek/x"
)

func promptText(t *testing.T) string {
	t.Helper()
	msgs := BuildPrompt(examplePack(), fixturePersona(t))
	return "SYSTEM\n" + msgs[0].Content + "\nUSER\n" + msgs[1].Content + "\n"
}

func TestBuildPrompt_SystemIsRulesThenPersona(t *testing.T) {
	msgs := BuildPrompt(examplePack(), fixturePersona(t))
	if len(msgs) != 2 || msgs[0].Role != "system" || msgs[1].Role != "user" {
		t.Fatalf("%+v", msgs)
	}
	body := strings.TrimSpace(fixturePersona(t).Body)
	if !strings.HasPrefix(msgs[0].Content, rulesText) || !strings.HasSuffix(strings.TrimSpace(msgs[0].Content), body) {
		t.Fatal("system = rules then persona")
	}
}

func TestBuildPrompt_UserIsThePackWithoutAllowed(t *testing.T) {
	var got map[string]any
	if err := json.Unmarshal([]byte(BuildPrompt(examplePack(), fixturePersona(t))[1].Content), &got); err != nil {
		t.Fatal(err)
	}
	if _, has := got["allowed"]; has {
		t.Fatal("allowed is the validator's")
	}
	if got["findings"] == nil || got["couldNotCheck"] == nil {
		t.Fatal("the pack is the user content")
	}
}

func TestBuildPrompt_PackCarriesNoIdentity(t *testing.T) {
	user := strings.ToLower(BuildPrompt(examplePack(), fixturePersona(t))[1].Content)
	for _, banned := range []string{"account", "guild", "realm", "battlenet", "battle.net", "battletag"} {
		if strings.Contains(user, banned) {
			t.Errorf("pack JSON mentions %q", banned)
		}
	}
}

func TestBuildPrompt_GoldenFury(t *testing.T) {
	got := promptText(t)
	if *updatePrompt {
		if err := os.WriteFile(goldenPromptFile, []byte(got), 0o600); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(goldenVersionFile, []byte(ModelVersion(fixtureModelSlug)+"\n"), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	want, err := os.ReadFile(goldenPromptFile)
	if err != nil {
		t.Fatal(err)
	}
	if got != string(want) {
		t.Fatal("the prompt moved: review the diff, then run with -update-prompt")
	}
	wantVersion, err := os.ReadFile(goldenVersionFile)
	if err != nil {
		t.Fatal(err)
	}
	if v := ModelVersion(fixtureModelSlug); v+"\n" != string(wantVersion) {
		t.Fatalf("model_version moved: %s, pinned %s", v, wantVersion)
	}
}

func TestRulesText_HasTheCapsAndRules(t *testing.T) {
	for _, s := range []string{"280", "60", "320", "200", "240", "160", "rank-1", "word for word", "heal more", "openingLinks", "Use only facts in the evidence pack"} {
		if !strings.Contains(rulesText, s) {
			t.Errorf("rules lack %q", s)
		}
	}
	for _, w := range bannedTone {
		if !strings.Contains(rulesText, w) {
			t.Errorf("rules lack the banned word %q", w)
		}
	}
}

func TestRulesText_NamesNoModel(t *testing.T) {
	for _, m := range []string{"deepseek", "kimi", "gpt", "claude", "glm"} {
		if strings.Contains(strings.ToLower(rulesText), m) {
			t.Errorf("rules name %q", m)
		}
	}
}

func TestModelVersion_MovesWithItsInputs(t *testing.T) {
	base := modelVersionOf(fixtureModelSlug, "rules", "schema", "template")
	if !strings.HasPrefix(base, "deepseek/x.") || len(base) != len("deepseek/x.")+versionHexChars {
		t.Fatalf("shape: %s", base)
	}
	if base != modelVersionOf(fixtureModelSlug, "rules", "schema", "template") {
		t.Fatal("not stable")
	}
	hash := func(v string) string { return strings.SplitN(v, ".", 2)[1] }
	for name, v := range map[string]string{
		"slug":     modelVersionOf("other/y", "rules", "schema", "template"),
		"rules":    modelVersionOf(fixtureModelSlug, "rules2", "schema", "template"),
		"schema":   modelVersionOf(fixtureModelSlug, "rules", "schema2", "template"),
		"template": modelVersionOf(fixtureModelSlug, "rules", "schema", "template2"),
	} {
		if hash(v) == hash(base) {
			t.Errorf("hash did not move with %s", name)
		}
	}
	if ModelVersion(fixtureModelSlug) != modelVersionOf(fixtureModelSlug, rulesText, outputSchemaText, personaTemplate) {
		t.Fatal("ModelVersion uses the embedded inputs")
	}
}

func schemaNode(t *testing.T) map[string]any {
	t.Helper()
	var s map[string]any
	if err := json.Unmarshal(OutputSchema(), &s); err != nil {
		t.Fatal(err)
	}
	return s
}

func TestOutputSchema_MatchesReviewDoc(t *testing.T) {
	s := schemaNode(t)
	props := s["properties"].(map[string]any)
	assertFields(t, "ReviewDoc", reflect.TypeOf(ReviewDoc{}), s)
	assertFields(t, "FixItem", reflect.TypeOf(FixItem{}), props["fixFirst"].(map[string]any))
	assertFields(t, "ThenItem", reflect.TypeOf(ThenItem{}), props["thenFix"].(map[string]any)["items"].(map[string]any))
	assertFields(t, "WellItem", reflect.TypeOf(WellItem{}), props["doingWell"].(map[string]any)["items"].(map[string]any))
}

// assertFields: every JSON field of the type is a schema property and the other way
// round; every field but ReviewDoc's optional openingLinks is required.
func assertFields(t *testing.T, name string, typ reflect.Type, node map[string]any) {
	t.Helper()
	props := node["properties"].(map[string]any)
	required := map[string]bool{}
	for _, r := range node["required"].([]any) {
		required[r.(string)] = true
	}
	seen := map[string]bool{}
	for i := range typ.NumField() {
		field := strings.Split(typ.Field(i).Tag.Get("json"), ",")[0]
		seen[field] = true
		if _, ok := props[field]; !ok {
			t.Errorf("%s.%s is not in the schema", name, field)
		}
		if !required[field] && field != "openingLinks" {
			t.Errorf("%s.%s is not required", name, field)
		}
	}
	for p := range props {
		if !seen[p] {
			t.Errorf("schema property %s.%s has no field", name, p)
		}
	}
}

func maxLen(node map[string]any, path ...string) int {
	for _, p := range path {
		node = node[p].(map[string]any)
	}
	v, _ := node["maxLength"].(float64)
	return int(v)
}

func maxItems(node map[string]any) int {
	v, _ := node["maxItems"].(float64)
	return int(v)
}

func TestOutputSchema_CapsMatchTheParser(t *testing.T) {
	s := schemaNode(t)
	props := s["properties"].(map[string]any)
	then := props["thenFix"].(map[string]any)
	well := props["doingWell"].(map[string]any)
	checks := map[string][2]int{
		"opening":         {maxLen(props, "opening"), CapOpeningChars},
		"thenFix items":   {maxItems(then), CapThenItems},
		"doingWell items": {maxItems(well), CapWellItems},
		"fix title":       {maxLen(props, "fixFirst", "properties", "title"), CapTitleChars},
		"fix why":         {maxLen(props, "fixFirst", "properties", "why"), CapFixWhyChars},
		"fix drillText":   {maxLen(props, "fixFirst", "properties", "drillText"), CapDrillChars},
		"then title":      {maxLen(then, "items", "properties", "title"), CapTitleChars},
		"then why":        {maxLen(then, "items", "properties", "why"), CapThenWhyChars},
		"well text":       {maxLen(well, "items", "properties", "text"), CapWellChars},
	}
	for name, c := range checks {
		if c[0] != c[1] {
			t.Errorf("%s: schema %d, parser %d", name, c[0], c[1])
		}
	}
}

func TestOutputSchema_WorkedExampleParses(t *testing.T) {
	raw, err := json.Marshal(exampleReview())
	if err != nil {
		t.Fatal(err)
	}
	if _, err := ParseOutput(string(raw)); err != nil {
		t.Fatal(err)
	}
}
