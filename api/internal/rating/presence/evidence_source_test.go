// api/internal/rating/presence/evidence_source_test.go
package presence

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// Task L6.12: a flagged report counts nowhere. Every reader of reports a player or guild sees
// picks them through EvidenceReport (not private, not flagged) or NotFlagged (private allowed,
// flagged never). This is the check for it over every non-test Go file of the api module. It
// reads statements, not lines: each string expression (a literal, or a + chain of literals and
// calls, however many lines it spans) that holds a condition on a report's visibility must
// also name EvidenceReport or NotFlagged, and each Go comparison of a visibility field must be
// on the list below. The list is the conditions that are not readers of evidence, each with
// its reason: mostly access rules, which decide which reports a viewer may open (an owner
// manages their own report, flagged or not), and caching choices.
//
// What it cannot see: a reader that never names visibility at all. A query that joins reports
// (or reads fights, fight_metrics or rating_scores) with no visibility condition passes here
// whether or not it leaves a flagged report out, and so does a reader whose EvidenceReport or
// NotFlagged call is deleted outright. Those readers are each guarded by a test that seeds a
// flagged report through the real reader and the real database:
//
//   - the score's inputs (rating.Store ratedFights, itemLevel, experience, guildNights):
//     TestAFlaggedFightCountsForNothingInTheScore, TestAFlaggedReportsOnlyRatedFightGivesNoScore,
//     TestTheAttendanceComponentsNightsExcludeAFlaggedReport;
//   - presence.In: TestInFindsNobodyInAFlaggedReport;
//   - rating.Store.staleFights, ReadCharacterRating, flaggedRerates:
//     TestStaleFightsSkipsAFlaggedReportsUnratedFights,
//     TestReadCharacterRatingLeavesOutAFlaggedReport,
//     TestRerateFlaggedReachesTheNightsGuildmatesAndTheAccountsAlts,
//     TestRerateFlaggedGuildmatesAreTheMembersAndTheWindowsRaiders;
//   - the per-fight ratings endpoint (a Go check of reports.Flagged):
//     TestFightRatingsIs404FlaggedForAFlaggedReport;
//   - rankings.Store.CharacterFights and scoreHistory: TestAFlaggedFightCountsForNothingInTheScore,
//     TestCharacterScoreHistorySkipsAFlaggedReport;
//   - the public guild page (rankings.Store.Guild): TestGuildProgressionIgnoresAFlaggedReport,
//     TestGuildRosterBestIgnoresAFlaggedReport, TestGuildReportListLeavesOutAFlaggedReport;
//   - the guild centre: attendance, log verification, moderator evidence and the raid present list
//     (TestAFlaggedReportVerifiesNobodyAndCountsForNobody), bestParses
//     (TestRosterBestParseIgnoresAFlaggedReport), latestRatings
//     (TestRosterRatingIgnoresAFlaggedReport),
//     HomeRoster's class/spec/item level (TestRosterClassSpecAndItemLevelIgnoreAFlaggedReport),
//     bestByRole (TestProgressionBestByRoleIgnoresAFlaggedReport), fillTopParse
//     (TestRaidsTopParseIgnoresAFlaggedReport), progression's counts, which pick NotFlagged
//     for a verified member and EvidenceReport for anyone else (countedReports: names neither
//     rule beside a visibility, so a swapped branch passes here)
//     (TestProgressionCountsAPrivateKillAndNeverAFlaggedOne,
//     TestProgressionCountsFollowTheViewer), the Raids tab's rows and fights (visibilityClause,
//     a flagged report only to a verified member who owns it)
//     (TestRaidsShowAFlaggedReportOnlyToAMemberWhoOwnsIt), the home summary
//     (TestHomeSummaryCountsAPrivateKillAndNeverAFlaggedOne), the Loot tab's killed
//     (TestLootKilledIgnoresAFlaggedKillAndCountsAPrivateOne);
//   - the nightly record (dataaddon windowParses, nightsByGuild, bossesKilledByGuild):
//     TestWindowParsesSkipsAFlaggedReport, TestTheNightlyGuildRecordIgnoresAFlaggedNight;
//   - the inbox's last_raid: TestLastRaidIsNeverAFlaggedReport;
//   - the recent-reports feed (reports.Store.Recent): TestStoreRecentLeavesOutAFlaggedReport.
//
// The public leaderboards (rankings.Store.Rankings, the guild boards) and the sims' parse and
// input readers read fight_metrics with no report join: a report's ranked rows are deleted
// when it is flagged (RemoveReport, from parse/sample.go and reports/ingest.go), guarded by
// TestSampleFlagsAReportWhoseStoredEventsDoNotMatch and
// TestRemoveReportWithdrawsItsRowsAndRecordsWhy.

// sqlVisibility is a SQL condition on a report's visibility, in any case and across lines:
// `visibility <> 'private'`, `r.Visibility = $1`, `visibility::text = ...`, `visibility not in
// (...)`, `visibility in (...)`, and the reversed `'public' = r.visibility`. A column named
// default_visibility is not a report's visibility (the word boundary leaves it out).
var sqlVisibility = []*regexp.Regexp{
	regexp.MustCompile(`(?is)\bvisibility\b(\s*::\s*\w+)?\s*(<>|!=|=|\bnot\s+in\b|\bin\b)`),
	regexp.MustCompile(`(?is)(<>|!=|=)\s*(\w+\.)?visibility\b`),
}

// goVisibility is a Go operand that names a visibility (rep.Visibility, visibility).
var goVisibility = regexp.MustCompile(`(?i)\bvisibility\b`)

// predicate is a call of one of the two rules.
var predicate = regexp.MustCompile(`\b(EvidenceReport|NotFlagged)\(`)

// notEvidence is every visibility condition that is not a reader of evidence: the file
// (slash-separated, from the api module root), a substring of the expression, and why.
var notEvidence = []struct{ file, expr, why string }{
	{"internal/rating/presence/presence.go", `".visibility <> 'private' and "`,
		"EvidenceReport itself"},
	{"internal/reports/report.go", "visibility != Private",
		"Ranked: the ingest's choice whether a report is ranked and rated at all, not a reader"},
	{"internal/reports/store.go", "visibility in ($2, $3)",
		"the guild's report list for its members: which reports a viewer may open (access)"},
	{"internal/reports/store.go", "visibility = coalesce($3, visibility)",
		"the report edit's update, not a reader"},
	{"internal/reports/handler.go", "rep.Visibility == ",
		"the report's own page: who may open it (mayView), its cache scope and where its files are"},
	{"internal/reports/handler.go", "p.Visibility != nil",
		"the report edit: whether a patch changes the visibility at all"},
	{"internal/reports/card.go", "rep.Visibility == ",
		"the report card image's cache scope"},
	{"internal/reports/uploads.go", `in.Visibility == ""`,
		"an upload with no visibility given defaults to public"},
	{"internal/rating/handler.go", "rep.Visibility == reports.Public",
		"the fight ratings' cache scope (whether the caller may see the report is reports.MayView)"},
	{"internal/reports/appreads.go", "rep.Visibility == ",
		"the apps' fight meters: their cache scope (whether the caller may see the report is MayView)"},
	{"internal/reports/appreads.go", "rep.Visibility != ",
		"the apps' fight meters: whether a name links to its character page, which only a public or guild report's names do; the meters themselves count nothing"},
	{"internal/reports/mayview.go", "rep.Visibility == ",
		"MayView: the one rule for who may open a report, used by reports and rating (access)"},
	{"internal/guilds/home.go", "r.owner_id = $3 or r.visibility = 'public'",
		"HomeReports: which reports the viewer may open, its own included (access)"},
	{"internal/parse/parse.go", "rep.Visibility == reports.Private",
		"the Discord notice's private-logs opt-in: whether the owner asked to be told about their own private log, not a reader"},
}

// unit is one expression the check reads: its source text and where it starts.
type unit struct {
	text string
	line int
}

// stringUnits is every maximal string expression in f: a string literal, or a + chain that
// holds one (the chain's calls, such as presence.EvidenceReport("r"), are part of it).
func stringUnits(fset *token.FileSet, src []byte, f *ast.File) []unit {
	var out []unit
	ast.Inspect(f, func(n ast.Node) bool {
		switch e := n.(type) {
		case *ast.BinaryExpr:
			if e.Op == token.ADD && holdsString(e) {
				out = append(out, unitOf(fset, src, e))
				return false
			}
		case *ast.BasicLit:
			if e.Kind == token.STRING {
				out = append(out, unitOf(fset, src, e))
			}
		}
		return true
	})
	return out
}

// comparisonUnits is every Go == or != comparison in f with an operand that names a
// visibility.
func comparisonUnits(fset *token.FileSet, src []byte, f *ast.File) []unit {
	var out []unit
	ast.Inspect(f, func(n ast.Node) bool {
		e, ok := n.(*ast.BinaryExpr)
		if !ok || (e.Op != token.EQL && e.Op != token.NEQ) {
			return true
		}
		for _, side := range []ast.Expr{e.X, e.Y} {
			if _, lit := side.(*ast.BasicLit); !lit && goVisibility.MatchString(unitOf(fset, src, side).text) {
				out = append(out, unitOf(fset, src, e))
				break
			}
		}
		return true
	})
	return out
}

func holdsString(n ast.Node) bool {
	found := false
	ast.Inspect(n, func(m ast.Node) bool {
		if l, ok := m.(*ast.BasicLit); ok && l.Kind == token.STRING {
			found = true
		}
		return !found
	})
	return found
}

func unitOf(fset *token.FileSet, src []byte, n ast.Node) unit {
	start, end := fset.Position(n.Pos()), fset.Position(n.End())
	return unit{text: string(src[start.Offset:end.Offset]), line: start.Line}
}

func TestEveryReportVisibilityConditionIsTheEvidenceRule(t *testing.T) {
	root, err := filepath.Abs(filepath.Join("..", "..", ".."))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(root, "go.mod")); err != nil {
		t.Fatalf("%s is not the api module root: %v", root, err)
	}
	used := make([]bool, len(notEvidence))
	allowed := func(rel string, u unit) bool {
		ok := false
		for i, a := range notEvidence {
			if a.file == rel && strings.Contains(u.text, a.expr) {
				used[i], ok = true, true
			}
		}
		return ok
	}
	var scanned int
	err = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if name := d.Name(); path != root && (name == "node_modules" || name == "testdata" || strings.HasPrefix(name, ".")) {
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)
		src, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, path, src, 0)
		if err != nil {
			return err
		}
		scanned++
		for _, u := range stringUnits(fset, src, f) {
			condition := false
			for _, re := range sqlVisibility {
				condition = condition || re.MatchString(u.text)
			}
			if condition && !predicate.MatchString(u.text) && !allowed(rel, u) {
				t.Errorf("%s:%d reads reports by visibility without presence.EvidenceReport or NotFlagged:\n\t%s",
					rel, u.line, strings.TrimSpace(u.text))
			}
		}
		for _, u := range comparisonUnits(fset, src, f) {
			if !allowed(rel, u) {
				t.Errorf("%s:%d compares a visibility in Go, which is not on the access list:\n\t%s", rel, u.line, u.text)
			}
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if scanned < 100 {
		t.Fatalf("scanned %d Go files under %s, want the whole api module", scanned, root)
	}
	for i, a := range notEvidence {
		if !used[i] {
			t.Errorf("the allowlist entry %s %q matches nothing: remove it", a.file, a.expr)
		}
	}
}

// The check itself catches the forms a line grep missed: a condition split over lines, in
// another case, with a cast, with not in, and a Go field comparison.
func TestTheVisibilityCheckCatchesEveryForm(t *testing.T) {
	src := "package x\n\nfunc f(rep struct{ Visibility string }) {\n" +
		"\t_ = `select 1 from reports r\n\t\twhere r.visibility\n\t\t  <> 'private'`\n" +
		"\t_ = `where R.VISIBILITY = 'public'`\n" +
		"\t_ = `where r.visibility::text = $1`\n" +
		"\t_ = `where r.visibility not in ('private')`\n" +
		"\t_ = `where 'public' = r.visibility`\n" +
		"\t_ = `where r.visibility <> 'private' and ` + EvidenceReport(\"r\")\n" +
		"\t_ = `default_visibility = $2`\n" +
		"\t_ = rep.Visibility == \"public\"\n}\n" +
		"func EvidenceReport(string) string { return \"\" }\n"
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "x.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	var caught int
	for _, u := range stringUnits(fset, []byte(src), f) {
		condition := false
		for _, re := range sqlVisibility {
			condition = condition || re.MatchString(u.text)
		}
		if condition && !predicate.MatchString(u.text) {
			caught++
		}
	}
	if caught != 5 {
		t.Errorf("caught %d SQL conditions, want the 5 without the rule (multi-line, upper case, cast, not in, reversed)", caught)
	}
	if n := len(comparisonUnits(fset, []byte(src), f)); n != 1 {
		t.Errorf("caught %d Go comparisons, want rep.Visibility == \"public\"", n)
	}
}
