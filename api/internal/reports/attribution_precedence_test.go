package reports

import "testing"

func TestAttributeGuildTheLoggingCharacterBeatsTheOwnersOtherSingleGuild(t *testing.T) {
	h := newHarness(t)
	a, b := h.seedGuild("Alpha", "hardcore"), h.seedGuild("Beta", "hardcore")
	// The logging character is verified in A; the uploader's only other verified guild is B.
	h.seedGuildCharacter(a, h.owner, harnessLoggingKey, true)
	h.seedGuildCharacter(b, h.owner, "us/hardcore/altbael", true)
	id := h.createReport(Public)

	got, err := h.store.AttributeGuild(t.Context(), id)
	if err != nil {
		t.Fatal(err)
	}
	if !sameGuild(got, a) {
		t.Fatalf("attributed %v, want the logging character's guild %d, not %d", got, a, b)
	}
}

func TestAttributeGuildKeepsAGuildAReportAlreadyCarriesWhenAttributionWouldPickAnother(t *testing.T) {
	h := newHarness(t)
	a, b := h.seedGuild("Alpha", "hardcore"), h.seedGuild("Beta", "hardcore")
	id := h.createReport(Public)
	if _, err := h.store.Pool.Exec(t.Context(), `update reports set guild_id = $2 where id = $1`, id, a); err != nil {
		t.Fatal(err)
	}
	// Now the logging character verifies in B: attribution would pick B.
	h.seedGuildCharacter(b, h.owner, harnessLoggingKey, true)

	got, err := h.store.AttributeGuild(t.Context(), id)
	if err != nil {
		t.Fatal(err)
	}
	if !sameGuild(got, a) || !sameGuild(h.reportGuild(id), a) {
		t.Fatalf("attributed %v, stored %v, want guild %d kept", got, h.reportGuild(id), a)
	}
}
