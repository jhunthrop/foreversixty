package catalog

import (
	"testing"
	"time"
)

func TestPrices_MatchProductSpec(t *testing.T) {
	want := map[string]struct {
		offer    Offer
		interval Interval
		founder  bool
		cents    int64
	}{
		LookupKeyPremiumMonthlyLaunch:    {OfferPremium, Monthly, false, 600},
		LookupKeyPremiumYearlyLaunch:     {OfferPremium, Yearly, false, 5000},
		LookupKeyPremiumFounderMonthly:   {OfferPremium, Monthly, true, 500},
		LookupKeyPremiumFounderYearly:    {OfferPremium, Yearly, true, 4000},
		LookupKeyGuildMonthly:            {OfferGuild, Monthly, false, 1500},
		LookupKeyGuildYearly:             {OfferGuild, Yearly, false, 15000},
		LookupKeyGuildFounderMonthly:     {OfferGuild, Monthly, true, 1200},
		LookupKeyGuildFounderYearly:      {OfferGuild, Yearly, true, 12000},
		LookupKeyRaidGuildMonthly:        {OfferRaidGuild, Monthly, false, 2500},
		LookupKeyRaidGuildYearly:         {OfferRaidGuild, Yearly, false, 25000},
		LookupKeyRaidGuildFounderMonthly: {OfferRaidGuild, Monthly, true, 2000},
		LookupKeyRaidGuildFounderYearly:  {OfferRaidGuild, Yearly, true, 20000},
	}
	for key, w := range want {
		p, ok := Lookup(key)
		if !ok {
			t.Errorf("%s: not in the catalog", key)
			continue
		}
		if p.Offer != w.offer || p.Interval != w.interval || p.Founder != w.founder || p.AmountCents != w.cents || !p.Sold {
			t.Errorf("%s = %+v, want %+v sold", key, p, w)
		}
	}
	if len(Prices) != 14 {
		t.Errorf("catalog has %d prices, want 14", len(Prices))
	}
	for _, key := range []string{LookupKeyPremiumMonthly, LookupKeyPremiumYearly} {
		if p, ok := Lookup(key); !ok || p.Sold {
			t.Errorf("%s = %+v, want a legacy price that is not sold", key, p)
		}
	}
}

func TestCheckoutLookupKey_FounderBoundary(t *testing.T) {
	cases := []struct {
		offer    Offer
		interval Interval
		at       time.Time
		want     string
	}{
		{OfferPremium, Yearly, FounderWindowCloses.Add(-time.Second), LookupKeyPremiumFounderYearly},
		{OfferPremium, Yearly, FounderWindowCloses, LookupKeyPremiumYearlyLaunch},
		{OfferRaidGuild, Monthly, FounderWindowCloses.Add(-time.Second), LookupKeyRaidGuildFounderMonthly},
		{OfferRaidGuild, Monthly, FounderWindowCloses, LookupKeyRaidGuildMonthly},
		{OfferGuild, Yearly, FounderWindowCloses.Add(time.Hour), LookupKeyGuildYearly},
	}
	for _, c := range cases {
		if got := CheckoutLookupKey(c.offer, c.interval, c.at); got != c.want {
			t.Errorf("%s %s at %s = %s, want %s", c.offer, c.interval, c.at, got, c.want)
		}
	}
	if FounderWindowCloses != time.Date(2027, 1, 7, 8, 0, 0, 0, time.UTC) {
		t.Errorf("window closes at %s", FounderWindowCloses)
	}
}

func TestCheckoutLookupKey_NeverLegacy(t *testing.T) {
	times := []time.Time{{}, FounderWindowCloses.Add(-time.Second), FounderWindowCloses, FounderWindowCloses.AddDate(5, 0, 0)}
	for _, o := range []Offer{OfferPremium, OfferGuild, OfferRaidGuild} {
		for _, i := range []Interval{Monthly, Yearly} {
			for _, at := range times {
				key := CheckoutLookupKey(o, i, at)
				if key == LookupKeyPremiumMonthly || key == LookupKeyPremiumYearly {
					t.Errorf("%s %s at %s sells legacy %s", o, i, at, key)
				}
				if p, ok := Lookup(key); !ok || !p.Sold {
					t.Errorf("%s %s at %s = %s, which is not a sold price", o, i, at, key)
				}
			}
		}
	}
}

func TestUpgradeLookupKey(t *testing.T) {
	cases := map[string]string{
		LookupKeyGuildFounderYearly:  LookupKeyRaidGuildFounderYearly,
		LookupKeyGuildMonthly:        LookupKeyRaidGuildMonthly,
		LookupKeyGuildFounderMonthly: LookupKeyRaidGuildFounderMonthly,
		LookupKeyGuildYearly:         LookupKeyRaidGuildYearly,
	}
	for from, want := range cases {
		if got, ok := UpgradeLookupKey(from); !ok || got != want {
			t.Errorf("UpgradeLookupKey(%s) = %s, %v, want %s", from, got, ok, want)
		}
	}
	for _, from := range []string{LookupKeyPremiumMonthly, LookupKeyPremiumFounderYearly, LookupKeyRaidGuildMonthly, "nope"} {
		if got, ok := UpgradeLookupKey(from); ok {
			t.Errorf("UpgradeLookupKey(%s) = %s, want false", from, got)
		}
	}
}

func TestPortalFamily(t *testing.T) {
	cases := map[string]string{
		LookupKeyPremiumMonthly:          "premium",
		LookupKeyPremiumYearly:           "premium",
		LookupKeyPremiumMonthlyLaunch:    "premium",
		LookupKeyPremiumYearlyLaunch:     "premium",
		LookupKeyPremiumFounderMonthly:   "premium_founder",
		LookupKeyPremiumFounderYearly:    "premium_founder",
		LookupKeyGuildMonthly:            "guild",
		LookupKeyGuildYearly:             "guild",
		LookupKeyGuildFounderMonthly:     "guild_founder",
		LookupKeyGuildFounderYearly:      "guild_founder",
		LookupKeyRaidGuildMonthly:        "raid_guild",
		LookupKeyRaidGuildYearly:         "raid_guild",
		LookupKeyRaidGuildFounderMonthly: "raid_guild_founder",
		LookupKeyRaidGuildFounderYearly:  "raid_guild_founder",
	}
	for key, want := range cases {
		if got, ok := PortalFamily(key); !ok || got != want {
			t.Errorf("PortalFamily(%s) = %s, %v, want %s", key, got, ok, want)
		}
	}
	if _, ok := PortalFamily("nope"); ok {
		t.Error("an unknown key has no family")
	}
	for _, family := range PortalFamilies {
		keys := FamilyLookupKeys(family)
		if len(keys) != 2 {
			t.Fatalf("%s has keys %v, want two", family, keys)
		}
		m, _ := Lookup(keys[0])
		y, _ := Lookup(keys[1])
		if m.Interval != Monthly || y.Interval != Yearly {
			t.Errorf("%s keys %v are not monthly then yearly", family, keys)
		}
	}
}

func TestOfferAndTierRules(t *testing.T) {
	if o, ok := ParseOffer("raid_guild"); !ok || o != OfferRaidGuild {
		t.Errorf("ParseOffer raid_guild = %v, %v", o, ok)
	}
	if _, ok := ParseOffer("gold"); ok {
		t.Error("unknown offer parsed")
	}
	for o, w := range map[Offer][3]string{
		OfferPremium:   {"premium", "", ProductIDPremium},
		OfferGuild:     {"guild", "guild", ProductIDGuild},
		OfferRaidGuild: {"guild", "raid_guild", ProductIDRaidGuild},
	} {
		if o.EntitlementPlan() != w[0] || o.Tier() != w[1] || o.ProductID() != w[2] {
			t.Errorf("%s = %s %s %s, want %v", o, o.EntitlementPlan(), o.Tier(), o.ProductID(), w)
		}
	}
	if SeatsForTier("guild") != SeatsGuild || SeatsForTier("raid_guild") != SeatsRaidGuild || SeatsForTier("") != 0 {
		t.Error("SeatsForTier")
	}
}
