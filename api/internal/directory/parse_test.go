// api/internal/directory/parse_test.go
package directory

import (
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/bnetapi"
)

func bnetSlug(name string) string { return bnetapi.CharacterSlug(name) }

func TestActiveSpecReadsOnlyTheActiveGroup(t *testing.T) {
	raw := []byte(`{"specialization_groups":[
		{"is_active":false,"specializations":[{"specialization_name":"Holy","spent_points":31}]},
		{"is_active":true,"specializations":[{"specialization_name":"Shadow","spent_points":11},{"specialization_name":"Discipline","spent_points":3}]}]}`)
	got, err := activeSpec(raw)
	if err != nil || got != "Shadow" {
		t.Fatalf("activeSpec = %q, %v", got, err)
	}
	none, err := activeSpec([]byte(`{"specialization_groups":[]}`))
	if err != nil || none != "" {
		t.Fatalf("no group = %q, %v", none, err)
	}
}

func TestSummarizeProfessionsReadsFlatSkillsAndTiers(t *testing.T) {
	raw := []byte(`{"primaries":[{"profession":{"name":"Mining"},"skill_points":120,"max_skill_points":150}],
		"secondaries":[{"profession":{"name":"Fishing"},"tiers":[{"skill_points":75,"max_skill_points":150}]}]}`)
	got, err := summarizeProfessions(raw)
	if err != nil || len(got) != 2 || got[0].Skill != 120 || got[1].Name != "Fishing" || got[1].Skill != 75 {
		t.Fatalf("professions = %+v, %v", got, err)
	}
}

func TestSummarizeEquipmentSkipsEmptySlots(t *testing.T) {
	raw := []byte(`{"equipped_items":[{"item":{"id":0},"slot":{"type":"SHIRT"}},
		{"item":{"id":7},"slot":{"type":"FEET"},"name":"Boots","quality":{"type":"COMMON"}}]}`)
	got, err := summarizeEquipment(raw)
	if err != nil || len(got) != 1 || got[0].Slot != "FEET" {
		t.Fatalf("equipment = %+v, %v", got, err)
	}
}
