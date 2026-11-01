@testable import ForeverSixty
import Foundation
import Models
import Testing

/// F11's rules (spec §5.3), ported from web/src/lib/rankings: the rank colour, the defaults, the
/// query, your row, paging, the words of the filter line, the answer and Home's line.
@Suite("Rankings logic")
struct RankingsLogicTests {
    // MARK: rank colour (row-view.ts)

    @Test func rankPercentile_firstIs100_lastJustAboveZero() {
        #expect(RankColor.percentile(rank: 1, total: 212) == 100)
        #expect(RankColor.percentile(rank: 212, total: 212) > 0)
        #expect(RankColor.percentile(rank: 1, total: 0) == 100)
    }

    @Test func rankColor_capsTheFirstRankAtPink99() {
        #expect(RankColor.tier(rank: 1, total: 212) == .pink)
        #expect(RankColor.tier(rank: 1, total: 1) == .pink)
    }

    @Test func rankColor_floors() {
        // 212 characters: rank 14 is 93.9 (purple), 204 is 4.2 (grey).
        #expect(RankColor.tier(rank: 14, total: 212) == .purple)
        #expect(RankColor.tier(rank: 204, total: 212) == .grey)
        #expect(RankColor.tier(rank: 3, total: 100) == .orange)
        #expect(RankColor.tier(rank: 26, total: 100) == .purple)
        #expect(RankColor.tier(rank: 51, total: 100) == .blue)
        #expect(RankColor.tier(rank: 76, total: 100) == .green)
        #expect(RankColor.tier(rank: 77, total: 100) == .grey)
    }

    // MARK: defaults (defaults.ts, phases.ts)

    @Test func phaseAt_isTheLastPhaseStartedByThen() throws {
        let phases = try RankingsTestData.phases().phases
        #expect(RankingsDefaulting.phaseAt(RankingsTestData.now, phases: phases) == "raids-1")
        #expect(RankingsDefaulting.phaseAt(RaidTime.date("2026-11-04T22:59:00Z"), phases: phases) == "beta")
        #expect(RankingsDefaulting.phaseAt(RaidTime.date("2026-11-04T23:00:00Z"), phases: phases) == "launch")
        #expect(RankingsDefaulting.phaseAt(RankingsTestData.now, phases: []) == "")
    }

    @Test func defaults_followTheCharacter_andTheRole() throws {
        let phases = try RankingsTestData.phases().phases
        let own = RankingsDefaulting.defaults(RankingsTestData.lightning, specs: RankingsTestData.specs, phases: phases, now: RankingsTestData.now)
        #expect(own == RankingsDefaults(classSlug: "shaman", spec: "Elemental", ruleset: "normal", phase: "raids-1", metric: .dps))

        let healer = RankingsCharacter(key: "us/pvp/a", name: "A", className: "Priest", spec: "holy", ruleset: "pvp")
        let heal = RankingsDefaulting.defaults(healer, specs: RankingsTestData.specs, phases: phases, now: RankingsTestData.now)
        #expect(heal.metric == .hps)
        #expect(heal.spec == "Holy")
        #expect(heal.classSlug == "priest")

        let tank = RankingsCharacter(key: "us/pvp/b", name: "B", className: "warrior", spec: "Protection", ruleset: "pvp")
        #expect(RankingsDefaulting.defaults(tank, specs: RankingsTestData.specs, phases: phases, now: RankingsTestData.now).metric == .damageTaken)
    }

    @Test func defaults_signedOut_aimNothing() throws {
        let phases = try RankingsTestData.phases().phases
        let open = RankingsDefaulting.defaults(nil, specs: RankingsTestData.specs, phases: phases, now: RankingsTestData.now)
        #expect(open == RankingsDefaults(classSlug: "", spec: "", ruleset: "", phase: "raids-1", metric: .dps))
    }

    @Test func defaultEncounter_isTheLastKillListed_elseTheFirstListed() throws {
        let encounters = try RankingsTestData.encounters().encounters
        let voltaris = try RankingsTestData.voltaris()
        #expect(RankingsDefaulting.defaultEncounter(history: voltaris.history, encounters: encounters) == "onyxia")
        // A wipe is skipped; the next kill (Thaurissan) decides.
        let wipeFirst = Array(voltaris.history.dropFirst())
        #expect(RankingsDefaulting.defaultEncounter(history: wipeFirst, encounters: encounters) == "emperor-dagran-thaurissan")
        #expect(RankingsDefaulting.defaultEncounter(history: [], encounters: encounters) == "onyxia")
        #expect(RankingsDefaulting.defaultEncounter(history: [], encounters: []) == "")
    }

    // MARK: the query

    @Test func query_fromFilters_leavesAllOut_andDedupes() {
        let query = RankingsQuery.board(RankingsTestData.filters, character: "us/normal/lightningbolltz", page: 2)
        #expect(query == RankingsQuery(
            encounter: "onyxia", metric: "dps", spec: "Elemental", classSlug: "shaman", phase: "raids-1", region: nil,
            ruleset: "normal", faction: nil, since: nil, distinct: "character", character: "us/normal/lightningbolltz", page: 2
        ))
        var today = RankingsTestData.filters
        today.todayOnly = true
        today.classSlug = ""
        today.spec = ""
        today.region = "eu"
        let open = RankingsQuery.board(today, character: nil, page: 1)
        #expect(open.since == "today")
        #expect(open.classSlug == nil)
        #expect(open.spec == nil)
        #expect(open.region == "eu")
        #expect(open.character == nil)
    }

    // MARK: your row (own-rank.ts, row-view.ts)

    @Test func ownRow_inPlace_pinned_none() throws {
        let page = try RankingsTestData.page()
        #expect(OwnRow.placement(rows: page.rows, characterRank: page.characterRank) == .inPlace(14))
        let off = try RankingsTestData.offPage()
        #expect(OwnRow.placement(rows: off.rows, characterRank: off.characterRank) == .pinned(204))
        #expect(OwnRow.placement(rows: page.rows, characterRank: nil) == .none)
    }

    @Test func rowsCarryCrests_onlyOnAMixedBoard() throws {
        #expect(!RankingsBoard.showsCrests(try RankingsTestData.page().rows))
        #expect(RankingsBoard.showsCrests(try RankingsTestData.mixed().rows))
    }

    @Test func youTag_onAnyOfTheAccountsCharacters() throws {
        let rows = try RankingsTestData.offPage().rows + [try #require(RankingsTestData.offPage().characterRank?.row)]
        let keys = Set(RankingsTestData.account().characterKeys)
        #expect(rows.filter { RankingsBoard.isYours($0, keys: keys) }.map(\.player.name) == ["Lightningbolltz", "Sparkwick"])
    }

    // MARK: paging

    @Test func paging_dedupesByRank() throws {
        let page = try RankingsTestData.page()
        #expect(RankingsPaging.merge(page.rows, page.rows) == page.rows)
        let first = Array(page.rows.prefix(10))
        let next = Array(page.rows.dropFirst(5).prefix(10))
        #expect(RankingsPaging.merge(first, next).map(\.rank) == Array(1...15))
    }

    // MARK: words

    @Test func filterLine_namesEveryFilter_andAllWords() {
        #expect(RankingsCopy.filterLine(RankingsTestData.filters, encounterName: "Onyxia")
            == "Onyxia · Normal · All regions · Phase 1 · Damage")
        var other = RankingsTestData.filters
        other.ruleset = ""
        other.region = "eu"
        other.phase = "launch"
        other.metric = .execution
        other.faction = "horde"
        other.todayOnly = true
        #expect(RankingsCopy.filterLine(other, encounterName: "Onyxia")
            == "Onyxia · Every ruleset · EU · Launch · Execution · Horde · Today only")
    }

    @Test func rowWords() throws {
        let row = try RankingsTestData.page().rows[0]
        #expect(RankingsCopy.value(row, metric: .dps) == "462")
        #expect(RankingsCopy.value(row, metric: .execution) == "97%")
        #expect(RankingsCopy.rowMeta(row, now: RankingsTestData.now) == "Hordebreakers · 40 players · 3:13 · 2\u{00A0}d\u{00A0}ago")
        #expect(RankingsCopy.unit(.damageTaken) == "DTPS")
        #expect(RankingsCopy.a11y(row, total: 212, metric: .dps, yours: false)
            == "Rank 1 of 212, Voltaris, 462 damage per second, Hordebreakers")
        #expect(RankingsCopy.a11y(row, total: 212, metric: .dps, yours: true).hasSuffix(", your character"))
    }

    @Test func answer_ranked_named_unranked() throws {
        let page = try RankingsTestData.page()
        let mine = RankingsCopy.answer(page: page, filters: RankingsTestData.filters, encounterName: "Onyxia", current: RankingsTestData.lightning)
        #expect(mine == .ranked(
            ordinal: "14th", tier: .purple, unit: "of 212 Elemental Shaman characters",
            meta: "Your best kill on Onyxia, Normal, Phase 1.", name: nil
        ))
        let off = try RankingsTestData.offPage()
        let alt = RankingsCopy.answer(page: off, filters: RankingsTestData.filters, encounterName: "Onyxia", current: RankingsTestData.lightning)
        #expect(alt == .ranked(
            ordinal: "204th", tier: .grey, unit: "of 212 Elemental Shaman characters",
            meta: "Sparkwick · best kill on Onyxia, Normal, Phase 1.", name: "Sparkwick"
        ))
        let kills = RankingsCopy.answer(page: try RankingsTestData.kills(), filters: RankingsTestData.filters, encounterName: "Onyxia", current: RankingsTestData.lightning)
        guard case let .ranked(_, _, unit, _, _) = kills else { Issue.record("not ranked"); return }
        #expect(unit == "of 212 Elemental Shaman kills")
        let none = RankingsCopy.answer(page: try RankingsTestData.emptyPage(), filters: RankingsTestData.filters, encounterName: "Onyxia", current: RankingsTestData.lightning)
        #expect(none == .unranked(title: "Not ranked yet on Onyxia.", meta: "A logged kill puts you on the board."))
        #expect(RankingsCopy.answer(page: page, filters: RankingsTestData.filters, encounterName: "Onyxia", current: nil) == nil)
    }

    @Test func homeLine() throws {
        #expect(RankingsCopy.homeLine(page: try RankingsTestData.page(), filters: RankingsTestData.filters, encounterName: "Onyxia")
            == "14th of 212 Elemental Shaman characters on Onyxia")
        #expect(RankingsCopy.homeLine(page: try RankingsTestData.emptyPage(), filters: RankingsTestData.filters, encounterName: "Onyxia")
            == "Not ranked yet.")
    }

    @Test func encounterMeta() throws {
        let onyxia = try #require(RankingsTestData.encounters().encounters.first)
        #expect(RankingsCopy.encounterMeta(onyxia) == "Onyxia's Lair · 40 players · 38 ranked kills")
    }

    @Test func phaseChips_newestFirst_withTheirWords() throws {
        let phases = try RankingsTestData.phases().phases
        #expect(RankingsCopy.phaseChoices(phases).map(\.label) == ["Phase 1", "Launch", "Beta", "Pre-beta"])
    }
}
