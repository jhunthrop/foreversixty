package main

import (
	"bytes"
	"context"
	"slices"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/billing"
	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
)

func premiumSubs(subs ...billing.StripeSubscriptionSummary) *billing.FakeGateway {
	return &billing.FakeGateway{ListSubscriptionsResult: map[string][]billing.StripeSubscriptionSummary{catalog.ProductIDPremium: subs}}
}

func TestArchive_DryRunArchivesNothing(t *testing.T) {
	gw := premiumSubs(billing.StripeSubscriptionSummary{ID: "sub_new", Status: "active", PriceLookupKey: "premium_monthly_launch"})
	var out bytes.Buffer
	if err := archiveUnsoldPrices(context.Background(), gw, false, &out); err != nil {
		t.Fatal(err)
	}
	if len(gw.Archived) != 0 {
		t.Fatalf("dry run archived %v", gw.Archived)
	}
	for _, key := range []string{"premium_monthly", "premium_yearly"} {
		if !strings.Contains(out.String(), "would archive "+key) {
			t.Errorf("output lacks %s: %q", key, out.String())
		}
	}
}

func TestArchive_RefusesWhenLegacySubscriptionLive(t *testing.T) {
	for _, status := range []string{"active", "trialing", "past_due"} {
		gw := premiumSubs(
			billing.StripeSubscriptionSummary{ID: "sub_old", Status: status, PriceLookupKey: "premium_yearly"},
			billing.StripeSubscriptionSummary{ID: "sub_gone", Status: "canceled", PriceLookupKey: "premium_monthly"},
		)
		var out bytes.Buffer
		if err := archiveUnsoldPrices(context.Background(), gw, true, &out); err == nil {
			t.Fatalf("%s: want refusal", status)
		}
		if len(gw.Archived) != 0 {
			t.Fatalf("%s: archived %v despite a live subscription", status, gw.Archived)
		}
		if !strings.Contains(out.String(), "sub_old") || strings.Contains(out.String(), "sub_gone") {
			t.Errorf("%s: output %q must name sub_old only", status, out.String())
		}
	}
}

func TestArchive_ApplyArchivesOnlyUnsold(t *testing.T) {
	gw := premiumSubs(billing.StripeSubscriptionSummary{ID: "sub_new", Status: "active", PriceLookupKey: "premium_yearly_launch"})
	if err := archiveUnsoldPrices(context.Background(), gw, true, &bytes.Buffer{}); err != nil {
		t.Fatal(err)
	}
	want := []string{"premium_monthly", "premium_yearly"}
	got := slices.Clone(gw.Archived)
	slices.Sort(got)
	if !slices.Equal(got, want) {
		t.Fatalf("archived %v, want %v", got, want)
	}
}
