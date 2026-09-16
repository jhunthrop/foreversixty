// logs/cmd/forever-logs/mechanics_draft_test.go
package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
)

func TestMechanicsDraftProposesATableFromTheFixtureLog(t *testing.T) {
	var out, errOut bytes.Buffer
	err := run([]string{"mechanics-draft", "-encounter", "9001", "../../../web/src/fixtures/report/fixture.log"}, &out, &errOut)
	if err != nil {
		t.Fatal(err, errOut.String())
	}
	table, err := mechanics.Parse(out.Bytes())
	if err != nil {
		t.Fatalf("the draft must be a valid table: %v\n%s", err, out.String())
	}
	if table.EncounterID != 9001 {
		t.Fatalf("encounter = %d", table.EncounterID)
	}
	var lash *mechanics.Mechanic
	for i := range table.Mechanics {
		if table.Mechanics[i].SpellID == 334660 {
			lash = &table.Mechanics[i]
		}
	}
	if lash == nil || !strings.HasPrefix(lash.Note, "draft:") {
		t.Fatalf("Anima Lash must be drafted with evidence: %+v", table.Mechanics)
	}
	_ = json.Valid
}
