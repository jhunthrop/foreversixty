package discordbot

import (
	"context"
	"encoding/json"
	"strings"
	"testing"
)

func TestCardIDRoundTrips(t *testing.T) {
	id := cardCustomID(prefixBis, "chg", "warrior", "feral-bear", "p1", "main_hand")
	if id != "gbis:1:chg:warrior:feral-bear:p1:main_hand" || len(id) > maxCustomIDLen {
		t.Fatal(id)
	}
	p, ok := parseCardID(id)
	if !ok || p.Prefix != "gbis" || p.Version != "1" || p.Part != "chg" || strings.Join(p.State, "|") != "warrior|feral-bear|p1|main_hand" {
		t.Fatalf("%+v", p)
	}
	if got := cardCustomID(prefixBis, "c", "", "", "p1", ""); got != "gbis:1:c:::p1:" {
		t.Fatal(got)
	}
	if got := cardCustomID(prefixChar, "name"); got != "gchar:1:name:" {
		t.Fatal(got)
	}
	if customPrefix(id) != prefixBis {
		t.Fatal("router key")
	}
}

func TestParseCardIDNeedsThreeFields(t *testing.T) {
	for _, id := range []string{"", "gbis", "gbis:1"} {
		if _, ok := parseCardID(id); ok {
			t.Errorf("%q parsed", id)
		}
	}
	p, ok := parseCardID("gchar:1:name:")
	if !ok || len(p.State) != 1 || p.State[0] != "" {
		t.Fatalf("%+v %v", p, ok)
	}
}

func cardPressOf(id string) Interaction {
	return interactionOf(typeComponent, "d1", map[string]any{"custom_id": id})
}

func TestUnknownCardVersionIsStale(t *testing.T) {
	called := false
	parts := map[string]cardPart{"go": func(context.Context, Interaction, cardPress) (Reply, error) {
		called = true
		return Reply{}, nil
	}}
	rep, err := (&Service{}).cardHandler("bis", parts)(context.Background(), cardPressOf("gbis:9:go:::p1:"))
	if err != nil || called {
		t.Fatal(err, called)
	}
	if !rep.Update || rep.Message.Content != "This card is out of date. Run /bis again." ||
		len(rep.Message.Embeds) != 0 || len(rep.Message.Components) != 0 || !rep.Message.Ephemeral {
		t.Fatalf("%+v", rep)
	}
}

func TestUnknownCardPartFails(t *testing.T) {
	rep, err := (&Service{}).cardHandler("bis", nil)(context.Background(), cardPressOf("gbis:1:zzz:"))
	if err != nil || !rep.Update || rep.Message.Content != commonCopy.Failure {
		t.Fatalf("%+v %v", rep, err)
	}
}

func TestCardHandlerRunsThePart(t *testing.T) {
	parts := map[string]cardPart{"go": func(_ context.Context, _ Interaction, p cardPress) (Reply, error) {
		return Reply{Message: Message{Content: strings.Join(p.State, "|")}}, nil
	}}
	rep, err := (&Service{}).cardHandler("bis", parts)(context.Background(), cardPressOf("gbis:1:go:a:b"))
	if err != nil || rep.Message.Content != "a|b" {
		t.Fatalf("%+v %v", rep, err)
	}
}

func TestLockedSelectHasOneNoneOption(t *testing.T) {
	c := lockedSelect("gbis:1:s:::p1:", "Choose a class first")
	if c.Kind != KindStringSelect || !c.Disabled || len(c.Options) != 1 ||
		c.Options[0].Value != "none" || c.Options[0].Label != "Choose a class first" {
		t.Fatalf("%+v", c)
	}
}

func TestSelectDefaultMarshals(t *testing.T) {
	on, _ := json.Marshal(SelectOption{Label: "A", Value: "a", Default: true})
	off, _ := json.Marshal(SelectOption{Label: "A", Value: "a"})
	if string(on) != `{"label":"A","value":"a","default":true}` || string(off) != `{"label":"A","value":"a"}` {
		t.Fatal(string(on), string(off))
	}
}

func TestAssertCardLimitsAcceptsAGoodCard(t *testing.T) {
	assertCardLimits(t, Message{Ephemeral: true, Components: []ActionRow{
		{lockedSelect(cardCustomID(prefixBis, "s"), "Choose a class first")},
		{{Kind: KindButton, Style: StyleSuccess, Label: "Show BiS", CustomID: cardCustomID(prefixBis, "go")}},
	}})
}
