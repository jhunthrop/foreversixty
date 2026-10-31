@testable import ForeverSixty
import Foundation
import Models
import Testing

@Suite("Reports logic")
struct ReportsLogicTests {
    private let now = ReportsTestData.now
    private let clock = ReportsTestData.clock

    // MARK: F10.1 list

    @Test func sections_splitAtSevenDays() throws {
        let rows = try ReportsTestData.list().rows
        let sections = ReportsList.sections(rows, now: now)
        #expect(sections.thisWeek.map(\.id) == ["m4xk9d2pqz7a", "ewlrx3enxbeb", "q7c2v9tk4mzh"])
        #expect(sections.earlier.map(\.id) == ["t3n8b5xw2kqe"])
    }

    @Test func rowWords_followTheState() throws {
        let rows = try ReportsTestData.list().rows
        #expect(ReportsCopy.rowMeta(rows[0], now: now, clock: clock) == "Updated 20 s ago")
        #expect(ReportsCopy.rowMeta(rows[1], now: now, clock: clock) == "14 Dec · 1 boss · 1 kill")
        #expect(ReportsCopy.rowMeta(rows[3], now: now, clock: clock) == "29 Nov · could not be read")
        #expect(rows.map { ReportRowState(status: $0.status) } == [.live, .complete, .processing, .failed])
        #expect(ReportsList.name(rows[1]) == "Onyxia, Monday")
        #expect(ReportsList.name(rows[0]) == "Blackrock Spire")
    }

    @Test func rowA11y_readsLiveSeconds() throws {
        let row = try ReportsTestData.list().rows[0]
        let words = ReportsCopy.rowA11y(row, now: now, clock: clock)
        #expect(words.hasPrefix("Blackrock Spire, Tuesday 15 December, 0 bosses, 0 kills, Private"))
        #expect(words.hasSuffix("live, updated 20 seconds ago"))
    }

    @Test func crest_usesTheLoggerThenTheCurrentCharacter() throws {
        let row = try ReportsTestData.list().rows[0]
        let account = try ReportsTestData.account(premium: false)
        #expect(ReportsList.crestClass(row, characters: account.characters, currentClass: "hunter") == "warrior")
        #expect(ReportsList.crestClass(row, characters: [], currentClass: "hunter") == "hunter")
        #expect(ReportsList.crestClass(row, characters: [], currentClass: nil) == nil)
    }

    @Test func visibility_glyphs() {
        #expect(ReportVisibility("public").systemImage == "globe")
        #expect(ReportVisibility("unlisted").systemImage == "link")
        #expect(ReportVisibility("guild").systemImage == "person.2")
        #expect(ReportVisibility("private").systemImage == "lock")
    }

    // MARK: F10.2 fights

    @Test func fightRows_bossPullsInOrderThenTheZonesTrash() throws {
        let rows = FightRows.rows(try ReportsTestData.report())
        #expect(rows.map(\.id) == ["fight-4", "fight-6", "fight-9", "trash-1"])
        #expect(FightRows.bossPulls(try ReportsTestData.report()) == 3)
    }

    @Test func fightRows_eachZonesTrashAfterItsBossPulls() throws {
        let rows = FightRows.rows(try ReportsTestData.twoZones())
        #expect(rows.map(\.id) == ["fight-4", "fight-6", "fight-9", "trash-1", "trash-2"])
        guard case let .trash(group) = rows[4] else { Issue.record("not trash"); return }
        #expect(ReportsCopy.meta(rows[4]) == "2 pulls · 2:35")
        #expect(group.zone == "Molten Core")
    }

    @Test func fightRows_trashWithNoGroupKeepsItsRows() throws {
        var report = try ReportsTestData.data("reports.get")
        report["trash_groups"] = [] as [Any]
        let rows = FightRows.rows(try ReportsTestData.decode(Report.self, report))
        #expect(rows.count == 9)
        #expect(rows.first == .loneTrash(try ReportsTestData.report().fights[0]))
    }

    @Test func fightMeta_killWipeAndTrash() throws {
        #expect(ReportsCopy.meta(try ReportsTestData.fightRow(9)) == "Kill · 4:28 · 40 players · 3 deaths")
        #expect(ReportsCopy.meta(try ReportsTestData.fightRow(4)) == "Wipe · 1:12 · boss at 61%")
        #expect(ReportsCopy.meta(try ReportsTestData.fightRow(6)) == "Wipe · 4:52 · boss at 22%")
        #expect(ReportsCopy.meta(try ReportsTestData.trashRow()) == "6 pulls · 7:40")
        #expect(ReportsCopy.title(try ReportsTestData.trashRow()) == "Trash · Onyxia's Lair")
    }

    @Test func wipeMeta_withoutBossHealth_isPlain() throws {
        var report = try ReportsTestData.data("reports.get")
        report["fights"] = (report["fights"] as? [[String: Any]] ?? []).map { fight in
            var fight = fight
            fight["boss_health_pct"] = -1
            return fight
        }
        let fight = try #require(try ReportsTestData.decode(Report.self, report).fights.first { $0.index == 4 })
        #expect(ReportsCopy.meta(.boss(fight)) == "Wipe · 1:12")
        #expect(ReportsCopy.fightA11y(.boss(fight)) == "Onyxia, wipe, 1 minute 12 seconds")
    }

    @Test func pullNumbers_numberABossPulledMoreThanOnce() throws {
        let report = try ReportsTestData.report()
        let pulls = FightRows.pullNumbers(report)
        #expect(pulls == [4: PullNumber(n: 1, of: 3), 6: PullNumber(n: 2, of: 3), 9: PullNumber(n: 3, of: 3)])
        #expect(ReportsCopy.title(try ReportsTestData.fightRow(9), pulls: pulls) == "Onyxia · pull 3 of 3")
        #expect(ReportsCopy.title(try ReportsTestData.trashRow(), pulls: pulls) == "Trash · Onyxia's Lair")
        let target = try #require(MetersTarget.of(try ReportsTestData.fightRow(4), reportID: report.id, pulls: pulls))
        #expect(target.name == "Onyxia · pull 1 of 3")
    }

    @Test func pullNumbers_aSinglePullKeepsThePlainName() throws {
        let pulls = FightRows.pullNumbers(try ReportsTestData.onlyWipes())
        let onlyKill = try ReportsTestData.decode(Report.self, {
            var data = try ReportsTestData.data("reports.get")
            data["fights"] = (data["fights"] as? [[String: Any]] ?? []).filter { ($0["index"] as? Int) != 4 && ($0["index"] as? Int) != 6 }
            return data
        }())
        #expect(pulls.count == 2)
        #expect(FightRows.pullNumbers(onlyKill).isEmpty)
        #expect(ReportsCopy.title(try ReportsTestData.fightRow(9), pulls: FightRows.pullNumbers(onlyKill)) == "Onyxia")
    }

    @Test func bossPulls_countWords() {
        #expect(ReportsCopy.bossPulls(0) == "No boss pulls")
        #expect(ReportsCopy.bossPulls(1) == "1 boss pull")
        #expect(ReportsCopy.bossPulls(3) == "3 boss pulls")
    }

    @Test func fightA11y_speaksTheTime() throws {
        #expect(ReportsCopy.fightA11y(try ReportsTestData.fightRow(4)) == "Onyxia, wipe, 1 minute 12 seconds, boss at 61 percent")
        #expect(ReportsCopy.fightA11y(try ReportsTestData.trashRow()) == "Trash in Onyxia's Lair, 6 pulls, 7 minutes 40 seconds")
    }

    @Test func reportMeta_weekdayAndLongDate() throws {
        #expect(ReportsCopy.reportMeta(try ReportsTestData.report(), clock: clock) == "Monday 14 December 2026 · logged by Obnoxious Yell")
    }

    // MARK: Your parses

    @Test func parseRows_andWords() throws {
        let rows = ParseRows.rows(try ReportsTestData.me())
        #expect(rows.count == 1)
        #expect(ReportsCopy.parseMeta(rows[0]) == "Kill 4:28 · Obnoxious Yell, Fury")
        #expect(ReportsCopy.figure(rows[0].fight.value) == "712")
        #expect(rows[0].tier == .purple)
        #expect(ReportsCopy.parseA11y(rows[0]) == "Onyxia, kill, 4 minutes 28 seconds, 712 damage per second, 91st percentile, New best")
        #expect(ParseRows.rows(try ReportsTestData.meWithoutKill()).isEmpty)
    }

    @Test(arguments: [
        (0, ParseTier.grey), (24, .grey), (25, .green), (49, .green), (50, .blue), (74, .blue),
        (75, .purple), (94, .purple), (95, .orange), (98, .orange), (99, .pink), (100, .pink),
    ])
    func parseLadder_floors(percentile: Int, tier: ParseTier) {
        #expect(ParseTier(percentile) == tier)
    }

    // MARK: F10.3 meters

    @Test func meterBar_isTheValueOverTheTop() {
        #expect(MeterRows.fraction(712.4, top: 788.3) == 712.4 / 788.3)
        #expect(MeterRows.fraction(10, top: 0) == 0)
        #expect(MeterRows.fraction(900, top: 800) == 1)
    }

    @Test func meterCut_elevenThenShowAll() throws {
        let players = try ReportsTestData.meters().players
        #expect(MeterRows.visible(players, limit: MeterRows.defaultVisible, showAll: false).map(\.rank) == Array(1...11))
        #expect(MeterRows.hasMore(players, limit: MeterRows.defaultVisible, showAll: false))
        #expect(MeterRows.visible(players, limit: MeterRows.defaultVisible, showAll: true).count == 40)
        #expect(!MeterRows.hasMore(players, limit: MeterRows.defaultVisible, showAll: true))
    }

    @Test func meterCut_largestText_addsYourRowAfterTheFour() throws {
        let players = try ReportsTestData.meters().players
        let rows = MeterRows.visible(players, limit: MeterRows.largeVisible, showAll: false)
        #expect(rows.map(\.rank) == [1, 2, 3, 4, 6])
        #expect(rows.last?.player.mine == true)
        #expect(MeterRows.hasMore(players, limit: MeterRows.largeVisible, showAll: false))
    }

    @Test func meterA11y_marksYourRow() throws {
        let players = try ReportsTestData.meters().players
        let mine = try #require(players.firstIndex { $0.mine })
        #expect(ReportsCopy.meterA11y(rank: mine + 1, player: players[mine], metric: .damage) == "Rank 6, Obnoxious Yell, Fury Warrior, 712, your character")
    }

    @Test func meterValues_followTheSegment() throws {
        let player = try #require(try ReportsTestData.meters().players.first { $0.mine })
        #expect(MeterMetric.damage.value(player) == 712.4)
        #expect(MeterMetric.taken.value(player) == 58.3)
        #expect(MeterMetric.allCases.map(\.unit) == ["dps", "hps", "dtps"])
    }

    @Test func metersTarget_bossAndTrash() throws {
        let boss = try #require(MetersTarget.of(try ReportsTestData.fightRow(9), reportID: ReportsTestData.reportID))
        #expect(boss.n == 9 && boss.scope == nil && !boss.isTrash)
        #expect(ReportsCopy.metersHead(boss, meters: try ReportsTestData.meters()) == "Kill · 4:28 · 40 players")
        let trash = try #require(MetersTarget.of(try ReportsTestData.trashRow(), reportID: ReportsTestData.reportID))
        #expect(trash.n == 1 && trash.scope == "trash" && trash.isTrash)
        #expect(ReportsCopy.metersHead(trash, meters: nil) == "6 pulls · 7:40")
    }

    // MARK: F10.4 coach

    @Test func coach_asksOnlyWithLogCoachAndAKill() throws {
        let me = try ReportsTestData.me()
        #expect(Coach.player(logCoach: true, me: me) == "us/pvp/obnoxious-yell")
        #expect(Coach.player(logCoach: false, me: me) == nil)
        #expect(Coach.player(logCoach: true, me: try ReportsTestData.meWithoutKill()) == nil)
    }

    @Test func coach_sectionsPresentInOrder() throws {
        let review = try ReportsTestData.review()
        let sections = Coach.sections(review.review)
        #expect(sections.map(\.kind) == [.fixFirst, .practice, .then, .doingWell, .couldntCheck])
        #expect(sections[0].lines == [.init(title: "Bloodthirst: 9 casts lost", body: review.review.fixFirst?.why ?? "")])
        #expect(sections[2].lines.count == 2)
        #expect(Coach.state(review) == .ready(review))
    }

    @Test func coach_sectionsSkipWhatIsMissing() throws {
        var data = try ReportsTestData.data("coach.fight")
        var review = data["review"] as? [String: Any] ?? [:]
        review["fixFirst"] = NSNull()
        review["thenFix"] = [] as [Any]
        data["review"] = review
        data["status"] = "writing"
        let decoded = try ReportsTestData.decode(CoachReview.self, data)
        #expect(Coach.sections(decoded.review).map(\.kind) == [.doingWell, .couldntCheck])
        #expect(Coach.state(decoded) == .writing)
    }

    @Test func coach_footer() throws {
        let footer = ReportsCopy.coachFooter(try ReportsTestData.review(), clock: clock)
        #expect(footer.hasPrefix("Reviewed 14 Dec · Based on: Measured against"))
        #expect(footer.hasSuffix("damage dealers; Measured against your own rage bar."))
        #expect(!footer.contains(".;"))
    }

    @Test func coach_footer_dropsAnExactDuplicateLabel() throws {
        var data = try ReportsTestData.data("coach.fight")
        let labels = data["basis_labels"] as? [String] ?? []
        data["basis_labels"] = labels + [labels[2]]
        let review = try ReportsTestData.decode(CoachReview.self, data)
        let footer = ReportsCopy.coachFooter(review, clock: clock)
        #expect(footer.components(separatedBy: "Measured against your own rage bar").count == 2)
    }

    @Test func metersHead_countsThePullsRoster() throws {
        let boss = try #require(MetersTarget.of(try ReportsTestData.fightRow(9), reportID: ReportsTestData.reportID))
        #expect(ReportsCopy.metersHead(boss, meters: try ReportsTestData.healing()) == "Kill · 4:28 · 40 players")
    }

    // MARK: Live and Home

    @Test func live_pollsOnlyALiveReport() throws {
        #expect(ReportsLive.polls(try ReportsTestData.liveReport()))
        #expect(!ReportsLive.polls(try ReportsTestData.report()))
        #expect(!ReportsLive.polls(nil))
    }

    @Test func homeLine_newestReport() throws {
        // The newest finished raid with a boss pull: never the live dungeon run above it.
        #expect(ReportsCopy.homeLine(try ReportsTestData.list()) == "Your last raid: Onyxia's Lair, 1 kill")
        let list = try ReportsTestData.list()
        let noRaid = ReportList(rows: list.rows.filter { !ReportsList.isLastRaid($0) }, total: 3, page: 1, perPage: 100)
        #expect(ReportsCopy.homeLine(noRaid) == "No reports yet.")
        #expect(ReportsCopy.homeLine(try RaidFixture.load(ReportList.self, "reports.list.empty")) == "No reports yet.")
    }

    @Test func durations() {
        #expect(ReportsCopy.duration(268_000) == "4:28")
        #expect(ReportsCopy.spokenDuration(192_000) == "3 minutes 12 seconds")
        #expect(ReportsCopy.ordinal(91) == "91st")
    }
}
