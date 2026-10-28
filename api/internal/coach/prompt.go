package coach

import (
	"crypto/sha256"
	_ "embed" // the rules and the schema are embedded files
	"encoding/hex"
	"encoding/json"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/llm"
)

// The shared coach rules (§9.2) and the output schema (§9.3). Neither names a model.
var (
	//go:embed prompt_rules.txt
	rulesTemplate string
	//go:embed schema.json
	outputSchemaText string
	// rulesText is the template with the validator's connectives filled in, so the
	// prompt and the validator read one list.
	rulesText = strings.ReplaceAll(rulesTemplate, connectivesPlaceholder, strings.Join(RelationConnectives, ", "))
)

const (
	connectivesPlaceholder = "{{connectives}}"
	roleSystem             = "system"
	roleUser               = "user"
	// personaTemplate sits between the rules and the persona in the system message. It is
	// part of the prompt template, so it is part of model_version.
	personaTemplate = "\nThe persona you write as follows. Its drill ids are the only ones you may use.\n\n"
	// versionHexChars is how much of the sha256 model_version keeps (§14).
	versionHexChars = 12
	versionSep      = "."
	hashSep         = "\x00"
)

// promptPack is the pack as the model sees it: every field but Allowed, which is the
// validator's set and not prompt material.
type promptPack struct {
	Version       string     `json:"pack_version"`
	Status        Status     `json:"status"`
	NotReviewable string     `json:"not_reviewable,omitempty"`
	Context       Context    `json:"context"`
	Findings      []Finding  `json:"findings"`
	Strengths     []Strength `json:"strengths"`
	Links         []Link     `json:"links"`
	CouldNotCheck []string   `json:"couldNotCheck"`
	Anchors       []Anchor   `json:"anchors"`
	BasisLabels   []string   `json:"basis_labels"`
}

func promptPackOf(p Pack) promptPack {
	return promptPack{
		Version: p.Version, Status: p.Status, NotReviewable: p.NotReviewable, Context: p.Context,
		Findings: p.Findings, Strengths: p.Strengths, Links: p.Links, CouldNotCheck: p.CouldNotCheck,
		Anchors: p.Anchors, BasisLabels: p.BasisLabels,
	}
}

// BuildPrompt is the system message (the coach rules, then the persona) and the user
// message (the pack as JSON).
func BuildPrompt(p Pack, persona Persona) []llm.Message {
	// A promptPack holds only strings, numbers and slices of them: Marshal cannot fail.
	packJSON, _ := json.Marshal(promptPackOf(p))
	system := rulesText + personaTemplate + strings.TrimSpace(persona.Body)
	return []llm.Message{
		{Role: roleSystem, Content: system},
		{Role: roleUser, Content: string(packJSON)},
	}
}

// ModelVersion is the model slug and the first 12 hex of the sha256 of the slug, the
// coach rules, the output schema and the prompt template (§14). It moves when any of
// them changes, and not otherwise.
func ModelVersion(model string) string {
	return modelVersionOf(model, rulesText, outputSchemaText, personaTemplate)
}

func modelVersionOf(model string, inputs ...string) string {
	sum := sha256.Sum256([]byte(strings.Join(append([]string{model}, inputs...), hashSep)))
	return model + versionSep + hex.EncodeToString(sum[:])[:versionHexChars]
}

// OutputSchema is the §9.3 shape as JSON Schema.
func OutputSchema() json.RawMessage { return json.RawMessage(outputSchemaText) }
