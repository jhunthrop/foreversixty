package main

import (
	"bytes"
	"context"
	"errors"
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
	if err := archiveRetiredPrices(context.Background(), gw, false, &out); err != nil {
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
	for _, status := range []string{"active", "trialing", "past_due", "incomplete", "unpaid", "paused"} {
		gw := premiumSubs(
			billing.StripeSubscriptionSummary{ID: "sub_old", Status: status, PriceLookupKey: "premium_yearly"},
			billing.StripeSubscriptionSummary{ID: "sub_gone", Status: "canceled", PriceLookupKey: "premium_monthly"},
		)
		var out bytes.Buffer
		if err := archiveRetiredPrices(context.Background(), gw, true, &out); err == nil {
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

func TestArchive_RefusesWhenListFails(t *testing.T) {
	gw := &billing.FakeGateway{Err: errors.New("stripe down")}
	if err := archiveRetiredPrices(context.Background(), gw, true, &bytes.Buffer{}); err == nil {
		t.Fatal("want the list error")
	}
	if len(gw.Archived) != 0 {
		t.Fatalf("archived %v after a failed list", gw.Archived)
	}
}

func TestArchive_RefusesMultiItemSubscription(t *testing.T) {
	for name, keys := range map[string][]string{
		"legacy as second item": {"premium_yearly_launch", "premium_monthly"},
		"two unrelated items":   {"premium_yearly_launch", "guild_monthly"},
		"legacy as first item":  {"premium_monthly", "premium_yearly_launch"},
	} {
		gw := premiumSubs(billing.StripeSubscriptionSummary{ID: "sub_multi", Status: "active", PriceLookupKey: keys[0], ItemLookupKeys: keys})
		var out bytes.Buffer
		if err := archiveRetiredPrices(context.Background(), gw, true, &out); err == nil {
			t.Fatalf("%s: want refusal", name)
		}
		if len(gw.Archived) != 0 || !strings.Contains(out.String(), "sub_multi") {
			t.Errorf("%s: archived %v, output %q", name, gw.Archived, out.String())
		}
	}
}

func TestArchive_ApplyArchivesOnlyTheTwoLegacyKeys(t *testing.T) {
	var subs []billing.StripeSubscriptionSummary
	for _, p := range catalog.Prices() {
		if p.Sold {
			subs = append(subs, billing.StripeSubscriptionSummary{ID: "sub_" + p.LookupKey, Status: "active", PriceLookupKey: p.LookupKey, ItemLookupKeys: []string{p.LookupKey}})
		}
	}
	gw := premiumSubs(subs...)
	var out bytes.Buffer
	if err := archiveRetiredPrices(context.Background(), gw, true, &out); err != nil {
		t.Fatal(err)
	}
	if want := []string{"premium_monthly", "premium_yearly"}; !slices.Equal(gw.Archived, want) {
		t.Fatalf("archived %v, want %v", gw.Archived, want)
	}
	for _, key := range gw.Archived {
		if !strings.Contains(out.String(), "archived "+key) {
			t.Errorf("output lacks %q: %q", "archived "+key, out.String())
		}
	}
	if strings.Contains(out.String(), "archiving") {
		t.Errorf("output %q should report archived, not archiving", out.String())
	}
}

func TestArchive_NoSoldKeyCanBeTargeted(t *testing.T) {
	if want := []string{"premium_monthly", "premium_yearly"}; !slices.Equal(retiredPriceKeys, want) {
		t.Fatalf("targets %v, want exactly %v", retiredPriceKeys, want)
	}
	for _, p := range catalog.Prices() {
		if p.Sold && slices.Contains(retiredPriceKeys, p.LookupKey) {
			t.Errorf("%s is sold but targeted for archiving", p.LookupKey)
		}
	}
}
