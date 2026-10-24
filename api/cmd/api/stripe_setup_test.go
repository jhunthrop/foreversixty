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
