package discordbot

import (
	"context"
	"testing"
)

// card presses a component of a card (type 3) as discordUser, with the values a select picked,
// and returns the routed handler's reply.
func (w *lookupWorld) card(t *testing.T, discordUser, customID string, values ...string) Reply {
	t.Helper()
	data := map[string]any{"custom_id": customID}
	if len(values) > 0 {
		data["values"] = values
	}
	return w.route(t, interactionOf(typeComponent, discordUser, data))
}

// modal submits a card's modal (type 5) with the text typed into its one input.
func (w *lookupWorld) modal(t *testing.T, discordUser, customID, typed string) Reply {
	t.Helper()
	data := map[string]any{"custom_id": customID, "components": []map[string]any{
		{"type": 1, "components": []map[string]any{{"type": 4, "custom_id": "name", "value": typed}}},
	}}
	return w.route(t, interactionOf(typeModalSubmit, discordUser, data))
}

func (w *lookupWorld) route(t *testing.T, in Interaction) Reply {
	t.Helper()
	h, ok := w.svc.components[customPrefix(in.CustomID())]
	if !ok {
		t.Fatalf("no route for custom id %q", in.CustomID())
	}
	rep, err := h(context.Background(), in)
	if err != nil {
		t.Fatal(err)
	}
	return rep
}
