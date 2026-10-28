package coach

// The history list's web fixtures (spec §3.5, §10.3): two pages of one Fury warrior's stored
// reviews, newest first, each item shaped as the history route scans it (historySQL: the
// stored review's fix-first title, the pack's spec, boss and kill). The titles come from the
// real rule-only renderer and the worked model review, so nothing here types a title.

import (
	"testing"
	"time"
)

// webHistoryNights is how many Molten Core nights the list holds after the Onyxia night.
const webHistoryNights = 3

// webHistoryReports names each night's report, newest first: the e2e report, then older ones.
var webHistoryReports = []string{webReportID, "fixture2mc02", "fixture2mc03", "fixture2mc04"}

// historyItemOf is one stored review as storedReviews returns it.
func historyItemOf(reportID string, fight int, p Pack, doc ReviewDoc, source string, at time.Time) historyItem {
	it := historyItem{
		ReportID: reportID, FightIndex: fight, PlayerKey: webFuryKey,
		Spec: p.Context.SpecLabel, Boss: p.Context.Boss, Kill: p.Context.Kill,
		Source: source, GeneratedAt: at,
	}
	if doc.FixFirst != nil && doc.FixFirst.Title != "" {
		it.FixFirst = &fixTitle{Title: doc.FixFirst.Title}
	}
	it.cursor = historyKey{At: at, ReportID: reportID, FightIndex: fight}
	return it
}

// webHistoryItems is every stored review, newest first: the Onyxia night (the worked kill
// read by the model, then the same pull as a wipe read from the log only), then Molten Core nights a
// week apart, eight bosses each in reverse kill order.
func webHistoryItems(t *testing.T, workedPack Pack, furyDoc ReviewDoc, persona Persona) []historyItem {
	t.Helper()
	wipeIn := furyWorked(t)
	wipeIn.Fight.Kill = false
	wipe := Review(wipeIn)
	out := []historyItem{
		historyItemOf(webHistoryReports[0], webFight+1, wipe, RenderRules(wipe, persona.Drills), SourceRules, webReviewedAt.Add(time.Minute)),
		historyItemOf(webHistoryReports[0], webFight, workedPack, furyDoc, SourceModel, webReviewedAt),
	}
	night := moltenCoreNightPacks(t)
	for n := 1; n <= webHistoryNights; n++ {
		base := webReviewedAt.AddDate(0, 0, -7*n)
		for i := len(night) - 1; i >= 0; i-- {
			f := night[i]
			at := base.Add(time.Duration(i) * time.Minute)
			out = append(out, historyItemOf(webHistoryReports[n], f.FightIndex, f.Pack, RenderRules(f.Pack, persona.Drills), SourceRules, at))
		}
	}
	return out
}

// webHistoryPages is the route's two pages over items: a full page with its cursor, then the rest.
func webHistoryPages(items []historyItem) (first, older historyPage) {
	first = historyPage{Reviews: items[:historyPageSize], NextCursor: encodeCursor(items[historyPageSize-1].cursor)}
	older = historyPage{Reviews: items[historyPageSize:]}
	return first, older
}

func TestWebHistory_PagesAreTheRoutesShape(t *testing.T) {
	fury := webPersona(t, "warrior-fury")
	worked := Review(furyWorked(t))
	items := webHistoryItems(t, worked, furyModelDoc(t, worked, fury), fury)
	if len(items) <= historyPageSize {
		t.Fatalf("%d items; the list needs more than a page to show Show older", len(items))
	}
	for i := 1; i < len(items); i++ {
		if !items[i].GeneratedAt.Before(items[i-1].GeneratedAt) {
			t.Fatalf("item %d is not older than item %d", i, i-1)
		}
	}
	first, older := webHistoryPages(items)
	after, err := decodeCursor(first.NextCursor)
	if err != nil || *after != first.Reviews[historyPageSize-1].cursor {
		t.Fatalf("cursor %q decodes to %+v, %v", first.NextCursor, after, err)
	}
	if older.NextCursor != "" || len(older.Reviews) == 0 {
		t.Fatalf("the older page is the last: %+v", older)
	}
	var wipes, none, rules int
	for _, it := range items {
		if !it.Kill {
			wipes++
		}
		if it.FixFirst == nil {
			none++
		}
		if it.Source == SourceRules {
			rules++
		}
	}
	if wipes == 0 || none == 0 || rules == 0 || rules == len(items) {
		t.Fatalf("the list shows a wipe (%d), a review with nothing to fix (%d) and both sources (%d rules of %d)",
			wipes, none, rules, len(items))
	}
}
