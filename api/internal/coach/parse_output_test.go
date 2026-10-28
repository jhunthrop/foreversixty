package coach

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func exampleJSON(t *testing.T) string {
	t.Helper()
	data, err := json.Marshal(exampleReview())
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}

func TestParseOutput_PlainAndFenced(t *testing.T) {
	raw := exampleJSON(t)
	for name, content := range map[string]string{
		"plain":  raw,
		"fenced": "```json\n" + raw + "\n```",
		"padded": "\n  " + raw + "  \n",
	} {
		doc, err := ParseOutput(content)
		if err != nil || !reflect.DeepEqual(doc, exampleReview()) {
			t.Errorf("%s: %v %+v", name, err, doc)
		}
	}
}

func TestParseOutput_Rejects(t *testing.T) {
	raw := exampleJSON(t)
	cases := map[string]string{
		"prose before":    "Here is the review: " + raw,
		"prose after":     raw + "\nHope that helps.",
		"two objects":     raw + raw,
		"unclosed fence":  "```json\n" + raw,
		"unknown field":   strings.Replace(raw, `{"opening"`, `{"extra":1,"opening"`, 1),
		"missing field":   strings.Replace(raw, `"couldntCheck"`, `"couldnt"`, 1),
		"not JSON":        "no",
		"empty":           "",
		"wrong item type": strings.Replace(raw, `"thenFix":[`, `"thenFix":[7,`, 1),
	}
	for name, content := range cases {
		if _, err := ParseOutput(content); err == nil {
			t.Errorf("%s: accepted", name)
		}
	}
}

func TestParseOutput_NullFixFirstParses(t *testing.T) {
	doc, err := ParseOutput(`{"opening":"x","fixFirst":null,"thenFix":[],"doingWell":[],"couldntCheck":[]}`)
	if err != nil || doc.FixFirst != nil {
		t.Fatalf("%v %+v", err, doc)
	}
}
