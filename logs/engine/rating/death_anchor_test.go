package rating

import "testing"

// DeathAnchor is DeathsTab.svelte's row id: `death-${death.guid}-${death.at_ms}`.
func TestDeathAnchorMatchesDeathsTab(t *testing.T) {
	if got := DeathAnchor("Player-1136-09362954", 170000); got != "death-Player-1136-09362954-170000" {
		t.Fatalf("DeathAnchor = %q", got)
	}
}
