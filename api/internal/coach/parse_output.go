package coach

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
)

// Caps on the review's prose (§9.3). The prompt and the JSON schema state the same values.
const (
	CapOpeningChars     = 280
	CapOpeningSentences = 2
	CapTitleChars       = 60
	CapFixWhyChars      = 320
	CapThenWhyChars     = 200
	CapDrillChars       = 240
	CapWellChars        = 160
	CapThenItems        = 3
	CapWellItems        = 2
)

// ReviewDoc is the output shape of §9.3: what the model returns and what the rule-only
// rendering fills. FixFirst is null exactly when the pack has no rank-1 finding.
type ReviewDoc struct {
	Opening      string     `json:"opening"`
	OpeningLinks []string   `json:"openingLinks"`
	FixFirst     *FixItem   `json:"fixFirst"`
	ThenFix      []ThenItem `json:"thenFix"`
	DoingWell    []WellItem `json:"doingWell"`
	CouldntCheck []string   `json:"couldntCheck"`
}

// FixItem is the review's one thing to fix first, with its drill.
type FixItem struct {
	Finding   string   `json:"finding"`
	Title     string   `json:"title"`
	Why       string   `json:"why"`
	Drill     string   `json:"drill"`
	DrillText string   `json:"drillText"`
	Anchors   []string `json:"anchors"`
	Links     []string `json:"links"`
}

// ThenItem is one of the further findings worth fixing.
type ThenItem struct {
	Finding string   `json:"finding"`
	Title   string   `json:"title"`
	Why     string   `json:"why"`
	Anchors []string `json:"anchors"`
	Links   []string `json:"links"`
}

// WellItem is one thing the player did well, citing a pack strength.
type WellItem struct {
	Strength string   `json:"strength"`
	Text     string   `json:"text"`
	Anchors  []string `json:"anchors"`
	Links    []string `json:"links"`
}

const (
	jsonFence = "```json"
	endFence  = "```"
)

// requiredKeys are the top-level fields the model must send, even when null or empty.
var requiredKeys = []string{"opening", "fixFirst", "thenFix", "doingWell", "couldntCheck"}

// ParseOutput reads the model's content into a ReviewDoc. It tolerates a ```json fence
// and nothing else: prose around the JSON, an unknown field, a missing top-level field or
// a second value after the first is an error.
func ParseOutput(content string) (ReviewDoc, error) {
	body := stripFence(strings.TrimSpace(content))
	if err := requireKeys(body); err != nil {
		return ReviewDoc{}, err
	}
	dec := json.NewDecoder(strings.NewReader(body))
	dec.DisallowUnknownFields()
	var doc ReviewDoc
	if err := dec.Decode(&doc); err != nil {
		return ReviewDoc{}, fmt.Errorf("review JSON: %w", err)
	}
	if _, err := dec.Token(); !errors.Is(err, io.EOF) {
		return ReviewDoc{}, errors.New("review JSON: text after the JSON object")
	}
	return doc, nil
}

func stripFence(s string) string {
	inner, ok := strings.CutPrefix(s, jsonFence)
	if !ok {
		return s
	}
	inner, ok = strings.CutSuffix(strings.TrimSpace(inner), endFence)
	if !ok {
		return s
	}
	return strings.TrimSpace(inner)
}

func requireKeys(body string) error {
	var top map[string]json.RawMessage
	if err := json.NewDecoder(strings.NewReader(body)).Decode(&top); err != nil {
		return fmt.Errorf("review JSON: %w", err)
	}
	for _, k := range requiredKeys {
		if _, ok := top[k]; !ok {
			return fmt.Errorf("review JSON: field %q is missing", k)
		}
	}
	return nil
}
