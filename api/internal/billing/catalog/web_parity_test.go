package catalog

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"testing"
	"time"
)

// webPlansPath is the site's one price file; it writes the same amounts and
// founder end as this catalog, and this test keeps the two from drifting.
const webPlansPath = "../../../../web/src/lib/billing/plans.ts"

var (
	webPlanBlock  = regexp.MustCompile(`(?s)key: '(\w+)',(.*?)\n  \}`)
	webPlanAmount = regexp.MustCompile(`\b(monthly|yearly|founderMonthly|founderYearly): (\d+),`)
	webFounderEnd = regexp.MustCompile(`FOUNDER_WINDOW_CLOSES = '([^']+)'`)
	webMemberCap  = regexp.MustCompile(`GUILD_TIER_MEMBER_LIMIT = (\d+);`)
)

// webField is plans.ts's field name for a catalog price.
func webField(p Price) string {
	switch {
	case p.Founder && p.Interval == Monthly:
		return "founderMonthly"
	case p.Founder:
		return "founderYearly"
	}
	return string(p.Interval)
}

func readWebPlans(t *testing.T) string {
	t.Helper()
	raw, err := os.ReadFile(webPlansPath)
	if err != nil {
		t.Fatal(err)
	}
	return string(raw)
}

// webAmounts is every plan amount plans.ts writes, as cents keyed "plan/field".
func webAmounts(t *testing.T, src string) map[string]int64 {
	t.Helper()
	out := map[string]int64{}
	for _, block := range webPlanBlock.FindAllStringSubmatch(src, -1) {
		for _, m := range webPlanAmount.FindAllStringSubmatch(block[2], -1) {
			dollars, err := strconv.ParseInt(m[2], 10, 64)
			if err != nil {
				t.Fatal(err)
			}
			out[block[1]+"/"+m[1]] = dollars * 100
		}
	}
	return out
}

func TestWebPlansMatchTheCatalog(t *testing.T) {
	src := readWebPlans(t)
	web := webAmounts(t, src)
	sold := 0
	for _, p := range Prices() {
		if !p.Sold {
			continue
		}
		sold++
		key := fmt.Sprintf("%s/%s", p.Offer, webField(p))
		if got, ok := web[key]; !ok || got != p.AmountCents {
			t.Errorf("%s: plans.ts %d cents (found %v), catalog %d", p.LookupKey, got, ok, p.AmountCents)
		}
	}
	if sold != 12 || len(web) != sold {
		t.Errorf("catalog sells %d prices, plans.ts writes %d amounts; want 12 each", sold, len(web))
	}
	end := webFounderEnd.FindStringSubmatch(src)
	if end == nil {
		t.Fatal("plans.ts has no FOUNDER_WINDOW_CLOSES")
	}
	if at, err := time.Parse(time.RFC3339, end[1]); err != nil || !at.Equal(FounderWindowCloses()) {
		t.Errorf("plans.ts founder end %s, catalog %s", end[1], FounderWindowCloses().Format(time.RFC3339))
	}
	if limit := webMemberCap.FindStringSubmatch(src); limit == nil || limit[1] != strconv.Itoa(GuildTierMemberLimit) {
		t.Errorf("plans.ts member limit %v, catalog %d", limit, GuildTierMemberLimit)
	}
}
