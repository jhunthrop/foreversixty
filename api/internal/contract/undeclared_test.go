package contract

import (
	"slices"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestUndeclared(t *testing.T) {
	const spec = `
components:
  schemas:
    Envelope:
      type: object
      properties: { ok: {}, data: { description: any }, request_id: {} }
    Row:
      type: object
      properties:
        id: {}
        tags: { type: array, items: { $ref: '#/components/schemas/Tag' } }
        stats: { type: object, additionalProperties: { type: integer } }
        extra: { type: object, additionalProperties: true }
        loose: { type: object }
        maybe: { oneOf: [{ $ref: '#/components/schemas/Tag' }, { type: "null" }] }
    Tag:
      type: object
      properties: { name: {} }
    Answer:
      allOf:
        - $ref: '#/components/schemas/Envelope'
        - { type: object, properties: { data: { $ref: '#/components/schemas/Row' } } }
`
	var root map[string]any
	if err := yaml.Unmarshal([]byte(spec), &root); err != nil {
		t.Fatal(err)
	}
	refs := localRefs(root)
	value := map[string]any{
		"ok": true, "request_id": "r", "stray": 1,
		"data": map[string]any{
			"id": 1, "unknown": 2,
			"tags":  []any{map[string]any{"name": "a"}, map[string]any{"name": "b", "colour": "x"}},
			"stats": map[string]any{"agility": 3},
			"extra": map[string]any{"anything": map[string]any{"deep": 1}},
			"loose": map[string]any{"whatever": 1},
			"maybe": map[string]any{"name": "n", "nope": 1},
		},
	}
	got := undeclared(refs("#/components/schemas/Answer"), value, refs)
	want := []string{"$.data.maybe.nope", "$.data.tags[1].colour", "$.data.unknown", "$.stray"}
	slices.Sort(got)
	if !slices.Equal(got, want) {
		t.Errorf("undeclared = %v, want %v", got, want)
	}
}
