// Package catalog is the single source of truth for what Forever Sixty
// sells: the offers, the fourteen Stripe lookup keys and their amounts,
// the founder window, the portal families and the tier rules. It is a
// leaf package and imports nothing internal.
// See docs/superpowers/specs/2026-10-27-pricing-design.md §3, §4 and §8.
package catalog

import (
	"slices"
	"time"
)

// CheckoutPath is the site's checkout page, the one place a plan is bought. The web builds its
// own links to it; a test in this package keeps the two the same.
const CheckoutPath = "/premium/checkout"

// Stripe Product ids, fixed so stripe-setup is idempotent without Search.
const (
	ProductIDPremium   = "prod_fs_premium"
	ProductIDGuild     = "prod_fs_guild"
	ProductIDRaidGuild = "prod_fs_raid_guild"
)

// The original four lookup keys. The Premium pair is no longer sold; the
// Guild pair is the Guild tier's list price.
const (
	LookupKeyPremiumMonthly = "premium_monthly"
	LookupKeyPremiumYearly  = "premium_yearly"
	LookupKeyGuildMonthly   = "guild_monthly"
	LookupKeyGuildYearly    = "guild_yearly"
)

// The launch lookup keys, beside the originals because an existing key's
// price is never changed in place.
const (
	LookupKeyPremiumMonthlyLaunch    = "premium_monthly_launch"
	LookupKeyPremiumYearlyLaunch     = "premium_yearly_launch"
	LookupKeyPremiumFounderMonthly   = "premium_founder_monthly"
	LookupKeyPremiumFounderYearly    = "premium_founder_yearly"
	LookupKeyGuildFounderMonthly     = "guild_founder_monthly"
	LookupKeyGuildFounderYearly      = "guild_founder_yearly"
	LookupKeyRaidGuildMonthly        = "raid_guild_monthly"
	LookupKeyRaidGuildYearly         = "raid_guild_yearly"
	LookupKeyRaidGuildFounderMonthly = "raid_guild_founder_monthly"
	LookupKeyRaidGuildFounderYearly  = "raid_guild_founder_yearly"
)

// Offer is something a customer can buy.
type Offer string

// The three offers.
const (
	OfferPremium   Offer = "premium"
	OfferGuild     Offer = "guild"
	OfferRaidGuild Offer = "raid_guild"
)

// Entitlement plan names, matching entitlements.PlanPremium and PlanGuild
// (this leaf package cannot import them).
const (
	planPremium = "premium"
	planGuild   = "guild"
)

// Tier names stored on a guild plan.
const (
	TierGuild     = "guild"
	TierRaidGuild = "raid_guild"
)

// ParseOffer parses the wire name of an offer.
func ParseOffer(s string) (Offer, bool) {
	switch o := Offer(s); o {
	case OfferPremium, OfferGuild, OfferRaidGuild:
		return o, true
	}
	return "", false
}

// EntitlementPlan is the entitlements plan the offer grants.
func (o Offer) EntitlementPlan() string {
	if o == OfferPremium {
		return planPremium
	}
	return planGuild
}

// Tier is the guild tier the offer buys, or "" for Premium.
func (o Offer) Tier() string {
	switch o {
	case OfferGuild:
		return TierGuild
	case OfferRaidGuild:
		return TierRaidGuild
	}
	return ""
}

// ProductID is the Stripe product the offer's prices live under.
func (o Offer) ProductID() string {
	switch o {
	case OfferGuild:
		return ProductIDGuild
	case OfferRaidGuild:
		return ProductIDRaidGuild
	}
	return ProductIDPremium
}

// Interval is a billing interval.
type Interval string

// The two billing intervals.
const (
	Monthly Interval = "monthly"
	Yearly  Interval = "yearly"
)

// Price is one Stripe price: its lookup key and what it sells.
type Price struct {
	LookupKey   string
	Offer       Offer
	Interval    Interval
	Founder     bool
	AmountCents int64
	// Sold is false for the legacy Premium prices, which no checkout offers.
	Sold bool
}

// prices lists all fourteen lookup keys, each family monthly then yearly.
var prices = []Price{
	{LookupKeyPremiumMonthly, OfferPremium, Monthly, false, 400, false},
	{LookupKeyPremiumYearly, OfferPremium, Yearly, false, 4000, false},
	{LookupKeyPremiumMonthlyLaunch, OfferPremium, Monthly, false, 600, true},
	{LookupKeyPremiumYearlyLaunch, OfferPremium, Yearly, false, 5000, true},
	{LookupKeyPremiumFounderMonthly, OfferPremium, Monthly, true, 500, true},
	{LookupKeyPremiumFounderYearly, OfferPremium, Yearly, true, 4000, true},
	{LookupKeyGuildMonthly, OfferGuild, Monthly, false, 1500, true},
	{LookupKeyGuildYearly, OfferGuild, Yearly, false, 15000, true},
	{LookupKeyGuildFounderMonthly, OfferGuild, Monthly, true, 1200, true},
	{LookupKeyGuildFounderYearly, OfferGuild, Yearly, true, 12000, true},
	{LookupKeyRaidGuildMonthly, OfferRaidGuild, Monthly, false, 2500, true},
	{LookupKeyRaidGuildYearly, OfferRaidGuild, Yearly, false, 25000, true},
	{LookupKeyRaidGuildFounderMonthly, OfferRaidGuild, Monthly, true, 2000, true},
	{LookupKeyRaidGuildFounderYearly, OfferRaidGuild, Yearly, true, 20000, true},
}

// Prices returns a copy of the fourteen prices, each family monthly then
// yearly.
func Prices() []Price { return slices.Clone(prices) }

// founderWindowCloses is the one end instant of the founder window,
// 6 January 2027 23:59:59 America/Los_Angeles, exclusive.
var founderWindowCloses = time.Date(2027, 1, 7, 8, 0, 0, 0, time.UTC)

// FounderWindowCloses is the instant the founder window ends (exclusive).
func FounderWindowCloses() time.Time { return founderWindowCloses }

// FounderOpen reports whether a checkout created at now gets founder prices.
func FounderOpen(now time.Time) bool { return now.Before(founderWindowCloses) }

// Lookup finds a price by lookup key.
func Lookup(lookupKey string) (Price, bool) {
	for _, p := range prices {
		if p.LookupKey == lookupKey {
			return p, true
		}
	}
	return Price{}, false
}

func find(o Offer, i Interval, founder bool) (Price, bool) {
	for _, p := range prices {
		if p.Sold && p.Offer == o && p.Interval == i && p.Founder == founder {
			return p, true
		}
	}
	return Price{}, false
}

// CheckoutLookupKey is the lookup key a checkout created at now sells: the
// founder price while the window is open, the list price after. It answers
// "" for an unknown offer or interval, and never a legacy key.
func CheckoutLookupKey(o Offer, i Interval, now time.Time) string {
	p, _ := find(o, i, FounderOpen(now))
	return p.LookupKey
}

// UpgradeLookupKey maps a Guild key to the Raid guild key with the same
// founder flag and interval, so the founder price survives the upgrade.
func UpgradeLookupKey(fromLookupKey string) (string, bool) {
	from, ok := Lookup(fromLookupKey)
	if !ok || from.Offer != OfferGuild {
		return "", false
	}
	to, ok := find(OfferRaidGuild, from.Interval, from.Founder)
	return to.LookupKey, ok
}

var portalFamilies = []string{
	"premium", "premium_founder", "guild", "guild_founder", "raid_guild", "raid_guild_founder",
}

// PortalFamilies returns the customer portal configurations, one per family.
func PortalFamilies() []string { return slices.Clone(portalFamilies) }

// PortalFamily is the portal configuration a subscription on lookupKey
// belongs to. The legacy Premium keys use the plain premium family.
func PortalFamily(lookupKey string) (string, bool) {
	p, ok := Lookup(lookupKey)
	if !ok {
		return "", false
	}
	if p.Founder {
		return string(p.Offer) + "_founder", true
	}
	return string(p.Offer), true
}

// FamilyLookupKeys is a family's two sold keys, monthly then yearly.
func FamilyLookupKeys(family string) []string {
	var keys []string
	for _, p := range prices {
		if f, _ := PortalFamily(p.LookupKey); f == family && p.Sold {
			keys = append(keys, p.LookupKey)
		}
	}
	return keys
}

// Tier rules.
const (
	// GuildTierMemberLimit is the verified-member count the Guild tier covers;
	// a soft cap, nothing switches off above it.
	GuildTierMemberLimit = 25
	// GuildTrialDays is the card-up-front trial on a guild's first plan.
	GuildTrialDays = 7
	SeatsGuild     = 1
	SeatsRaidGuild = 3

	SeatMoveLock      = 30 * 24 * time.Hour
	SeatReturnedShown = 14 * 24 * time.Hour
	ComebackValid     = 30 * 24 * time.Hour
)

// SeatsForTier is how many Premium seats a guild tier includes.
func SeatsForTier(tier string) int {
	switch tier {
	case TierGuild:
		return SeatsGuild
	case TierRaidGuild:
		return SeatsRaidGuild
	}
	return 0
}

// PriceFacts is the price a row was bought at, from its lookup key. A row
// with no catalog price (a grant) has null interval and amount.
type PriceFacts struct {
	Founder     bool    `json:"founder"`
	Interval    *string `json:"interval"`
	AmountCents *int64  `json:"amount_cents"`
}

// FactsOf reads founder and interval from the catalog. The amount is the
// row's own when it has one, else the catalog's; a key outside the catalog
// with no amount has none.
func FactsOf(lookupKey string, unitAmountCents int64) PriceFacts {
	var facts PriceFacts
	price, known := Lookup(lookupKey)
	if known {
		interval := string(price.Interval)
		facts.Founder, facts.Interval = price.Founder, &interval
	}
	switch {
	case unitAmountCents > 0:
		facts.AmountCents = &unitAmountCents
	case known:
		amount := price.AmountCents
		facts.AmountCents = &amount
	}
	return facts
}
