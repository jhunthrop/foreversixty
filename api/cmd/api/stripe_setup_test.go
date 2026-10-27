// api/cmd/api/stripe_setup_test.go
package main

import (
	"slices"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/billing"
	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
)

// Managed Payments refuses a Checkout whose product has no tax code, so
// every spec stripe-setup ensures carries one; a product created without
// it would sell nothing until a hand edit in the Dashboard.
func TestEveryPriceSpecCarriesTheSaaSTaxCode(t *testing.T) {
	for _, spec := range priceSpecs {
		if spec.TaxCode != billing.ProductTaxCodeSaaSPersonal {
			t.Errorf("%s: tax code %q, want %q", spec.LookupKey, spec.TaxCode, billing.ProductTaxCodeSaaSPersonal)
		}
	}
}

// priceSpecs is the catalog and nothing else: same keys, same order, the
// amount, the product and Stripe's interval word for each.
func TestPriceSpecs_FromCatalog(t *testing.T) {
	if len(priceSpecs) != len(catalog.Prices()) {
		t.Fatalf("priceSpecs has %d entries, catalog has %d", len(priceSpecs), len(catalog.Prices()))
	}
	wantNames := map[string]string{
		catalog.ProductIDPremium:   "Forever Sixty Premium",
		catalog.ProductIDGuild:     "Forever Sixty Guild",
		catalog.ProductIDRaidGuild: "Forever Sixty Raid Guild",
	}
	for i, p := range catalog.Prices() {
		s := priceSpecs[i]
		interval := "month"
		if p.Interval == catalog.Yearly {
			interval = "year"
		}
		if s.LookupKey != p.LookupKey || s.UnitAmountCents != p.AmountCents || s.ProductID != p.Offer.ProductID() ||
			s.Interval != interval || s.ProductName != wantNames[s.ProductID] {
			t.Errorf("priceSpecs[%d] = %+v, want catalog price %+v", i, s, p)
		}
	}
}

// Six portal configurations, one per family, each listing exactly that
// family's two keys under its one product; the founder families list only
// founder keys, so a monthly to yearly switch keeps the lock.
func TestPortalSpecs_SixFamiliesOneProductEach(t *testing.T) {
	specs := portalSpecs("https://foreversixty.gg")
	if len(specs) != 6 {
		t.Fatalf("%d portal specs, want 6", len(specs))
	}
	seen := map[string]bool{}
	for _, spec := range specs {
		if seen[spec.Key] {
			t.Errorf("family %s appears twice", spec.Key)
		}
		seen[spec.Key] = true
		if spec.PrivacyPolicyURL != "https://foreversixty.gg/privacy" || spec.TermsOfServiceURL != "https://foreversixty.gg/terms" || spec.DefaultReturnURL != "https://foreversixty.gg/account" {
			t.Errorf("%s: URLs = %q %q %q", spec.Key, spec.PrivacyPolicyURL, spec.TermsOfServiceURL, spec.DefaultReturnURL)
		}
		if len(spec.Products) != 1 {
			t.Fatalf("%s lists %d products, want 1", spec.Key, len(spec.Products))
		}
		p := spec.Products[0]
		if !slices.Equal(p.LookupKeys, catalog.FamilyLookupKeys(spec.Key)) {
			t.Errorf("%s keys = %v, want %v", spec.Key, p.LookupKeys, catalog.FamilyLookupKeys(spec.Key))
		}
		for _, key := range p.LookupKeys {
			price, _ := catalog.Lookup(key)
			if price.Offer.ProductID() != p.ProductID {
				t.Errorf("%s: key %s is not under product %s", spec.Key, key, p.ProductID)
			}
			family, _ := catalog.PortalFamily(key)
			if family != spec.Key {
				t.Errorf("%s lists %s, which belongs to %s", spec.Key, key, family)
			}
		}
	}
	for _, spec := range specs {
		founder := slices.Contains([]string{"premium_founder", "guild_founder", "raid_guild_founder"}, spec.Key)
		for _, key := range spec.Products[0].LookupKeys {
			if price, _ := catalog.Lookup(key); price.Founder != founder {
				t.Errorf("%s lists %s (founder=%v)", spec.Key, key, price.Founder)
			}
		}
	}
}
