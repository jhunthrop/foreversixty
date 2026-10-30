package discordbot

import (
	"strings"
	"testing"
)

func TestEmptyOrForgedNameAsksInsteadOfNamingNobody(t *testing.T) {
	w := newLookupWorld(t)
	linkTwoChars(w)
	for _, rep := range []Reply{
		w.modal(t, pressDiscord, charNameID, "   "),
		w.card(t, pressDiscord, charPickID, ""),
	} {
		assertCardLimits(t, rep.Message)
		if !rep.Update || contentOf(rep.Message) != "Which character? Type a name." || strings.Contains(contentOf(rep.Message), "named") {
			t.Fatalf("reply = %+v", rep)
		}
	}
	if rep := w.modal(t, "d-stranger", charNameID, ""); contentOf(rep.Message) != guideCopy.CharLeadUnlinked {
		t.Fatalf("unlinked = %+v", rep)
	}
}

func TestBareCharWithAMainThatNoLongerResolvesAsks(t *testing.T) {
	w := newLookupWorld(t)
	linkTwoChars(w)
	delete(w.cards.cards, "zulmara")
	rep := w.command(t, pressDiscord, "char", nil)
	assertCardLimits(t, rep.Message)
	if rep.Update || contentOf(rep.Message) != "Which character? Type a name." || hasCard(rep.Message) {
		t.Fatalf("reply = %+v", rep)
	}
}

func TestSomeoneElseLandsBeforeTheLinkWhateverTheOrder(t *testing.T) {
	own := []ownChar{{}, {}}
	link := Component{Kind: KindButton, Style: StyleLink, Label: "Character page", URL: "https://x.test"}
	share := Component{Kind: KindButton, Style: StyleSecondary, Label: "Share in channel", CustomID: "share:t"}
	for name, row := range map[string]ActionRow{"link last": {share, link}, "link first": {link, share}, "link only": {link}} {
		m := withOwnControls(Message{Components: []ActionRow{row}}, own, "x")
		labels := strings.Join(rowLabels(m.Components[0]), "|")
		li := strings.Index(labels, "Character page")
		se := strings.Index(labels, "Someone else")
		if se < 0 || se > li || len(m.Components) != 2 {
			t.Errorf("%s: row = %q rows = %d", name, labels, len(m.Components))
		}
	}
}
