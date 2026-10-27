// api/cmd/api/stripe_setup_test.go
package main

import (
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/billing"
)

// Managed Payments refuses a Checkout whose product has no tax code, so
// every spec stripe-setup ensures carries one; a product created without
// it would sell nothing until a hand edit in the Dashboard.
func TestEveryPriceSpecCarriesTheSaaSTaxCode(t *testing.T) {
	lookupKeys := map[string]bool{}
	for _, spec := range priceSpecs {
		if spec.TaxCode != billing.ProductTaxCodeSaaSPersonal {
			t.Errorf("%s: tax code %q, want %q", spec.LookupKey, spec.TaxCode, billing.ProductTaxCodeSaaSPersonal)
		}
		lookupKeys[spec.LookupKey] = true
	}
	for _, key := range []string{billing.LookupKeyPremiumMonthly, billing.LookupKeyPremiumYearly, billing.LookupKeyGuildMonthly, billing.LookupKeyGuildYearly} {
		if !lookupKeys[key] {
			t.Errorf("priceSpecs has no %s", key)
		}
	}
}

// The portal's switch set is each product's own prices and nothing
// across products: a Premium subscriber is never offered the guild plan
// there (that path is the site's own guild checkout, spec §2.3).
func TestPortalSpecOffersEachProductsOwnPricesOnTheSitesPages(t *testing.T) {
	spec := portalSpec("https://foreversixty.gg")
	if spec.PrivacyPolicyURL != "https://foreversixty.gg/privacy" || spec.TermsOfServiceURL != "https://foreversixty.gg/terms" || spec.DefaultReturnURL != "https://foreversixty.gg/account" {
		t.Errorf("portal URLs = %q %q %q, want the site's privacy, terms and account pages", spec.PrivacyPolicyURL, spec.TermsOfServiceURL, spec.DefaultReturnURL)
	}
	want := map[string][]string{
		billing.ProductIDPremium: {billing.LookupKeyPremiumMonthly, billing.LookupKeyPremiumYearly},
		billing.ProductIDGuild:   {billing.LookupKeyGuildMonthly, billing.LookupKeyGuildYearly},
	}
	if len(spec.Products) != len(want) {
		t.Fatalf("portal lists %d products, want %d", len(spec.Products), len(want))
	}
	for _, p := range spec.Products {
		keys, ok := want[p.ProductID]
		if !ok {
			t.Errorf("portal lists unknown product %s", p.ProductID)
			continue
		}
		if len(p.LookupKeys) != len(keys) {
			t.Errorf("%s: %v, want %v", p.ProductID, p.LookupKeys, keys)
			continue
		}
		for i := range keys {
			if p.LookupKeys[i] != keys[i] {
				t.Errorf("%s: %v, want %v", p.ProductID, p.LookupKeys, keys)
			}
		}
	}
}

// launchPrice is one row of the launch pricing table.
type launchPrice struct {
	product  string
	interval string
	cents    int64
}

// launchPriceTable is docs/superpowers/specs/2026-10-27-pricing-design.md §3, plus the
// original four lookup keys, which stay until the code that sells the launch prices ships.
var launchPriceTable = map[string]launchPrice{
	billing.LookupKeyPremiumMonthly:          {billing.ProductIDPremium, "month", 400},
	billing.LookupKeyPremiumYearly:           {billing.ProductIDPremium, "year", 4000},
	billing.LookupKeyPremiumMonthlyLaunch:    {billing.ProductIDPremium, "month", 600},
	billing.LookupKeyPremiumYearlyLaunch:     {billing.ProductIDPremium, "year", 5000},
	billing.LookupKeyPremiumFounderMonthly:   {billing.ProductIDPremium, "month", 500},
	billing.LookupKeyPremiumFounderYearly:    {billing.ProductIDPremium, "year", 4000},
	billing.LookupKeyGuildMonthly:            {billing.ProductIDGuild, "month", 1500},
	billing.LookupKeyGuildYearly:             {billing.ProductIDGuild, "year", 15000},
	billing.LookupKeyGuildFounderMonthly:     {billing.ProductIDGuild, "month", 1200},
	billing.LookupKeyGuildFounderYearly:      {billing.ProductIDGuild, "year", 12000},
	billing.LookupKeyRaidGuildMonthly:        {billing.ProductIDRaidGuild, "month", 2500},
	billing.LookupKeyRaidGuildYearly:         {billing.ProductIDRaidGuild, "year", 25000},
	billing.LookupKeyRaidGuildFounderMonthly: {billing.ProductIDRaidGuild, "month", 2000},
	billing.LookupKeyRaidGuildFounderYearly:  {billing.ProductIDRaidGuild, "year", 20000},
}

// The launch prices sit beside the original four under their own lookup keys, so ensuring
// them changes no live checkout until the code that sells them ships.
func TestPriceSpecsCarryTheLaunchPricingTable(t *testing.T) {
	got := map[string]billing.PriceSpec{}
	for _, spec := range priceSpecs {
		if _, dup := got[spec.LookupKey]; dup {
			t.Errorf("lookup key %s appears twice", spec.LookupKey)
		}
		got[spec.LookupKey] = spec
	}
	if len(got) != len(launchPriceTable) {
		t.Errorf("priceSpecs has %d lookup keys, want %d", len(got), len(launchPriceTable))
	}
	for key, want := range launchPriceTable {
		s, ok := got[key]
		if !ok {
			t.Errorf("priceSpecs has no %s", key)
			continue
		}
		if s.ProductID != want.product || s.Interval != want.interval || s.UnitAmountCents != want.cents {
			t.Errorf("%s = %s %s %d, want %s %s %d", key, s.ProductID, s.Interval, s.UnitAmountCents, want.product, want.interval, want.cents)
		}
	}
}
