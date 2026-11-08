@testable import ForeverSixty
import DesignSystem
import Foundation
import Models
import Networking
import SwiftUI
import Testing

@MainActor
@Suite("Home flow parity")
struct HomeFlowParityTests {
    static let device = CGSize(width: 393, height: 852)
    typealias T = HomeFlowTestData

    private func check(_ content: HomeFlowContent, _ screen: String, testName: String = #function) throws {
        let view = HomeFlowView(content: content, build: T.build).environment(\.itemIconSource, RepoItemIconSource())
        try ParityCheck.assert(view, screen: screen, size: Self.device, testName: testName)
    }

    @Test("ios-home-raid-day")
    func raidDay() throws {
        try check(T.content([.raidCard(T.raid()), T.upgrades(bestRun: false), T.newForYou("day")]), "home-raid-day")
    }

    @Test("ios-home-raid-day-late")
    func raidDayLate() throws {
        try check(T.content([.raidCard(T.raid(late: true)), T.upgrades(bestRun: false), T.newForYou("day")]), "home-raid-day-late")
    }

    @Test("ios-home-raid-day-officer")
    func raidDayOfficer() throws {
        try check(
            T.content([.raidCard(T.raid(officer: true)), T.upgrades(bestRun: false), T.newForYou("day")]),
            "home-raid-day-officer"
        )
    }

    @Test("ios-home-raid-live")
    func raidLive() throws {
        try check(T.content([T.raidLive, T.newForYou("live")]), "home-raid-live")
    }

    @Test("ios-home-after-raid")
    func afterRaid() throws {
        try check(T.content([T.lastNight, T.upgrades(bestRun: true), T.newForYou("day")]), "home-after-raid")
    }

    @Test("ios-home-60")
    func sixty() throws {
        try check(T.content([T.upgrades(bestRun: true), .raidLine(T.raid()), T.newForYou("day")]), "home-60")
    }

    @Test("ios-home-60-noavatar")
    func sixtyNoAvatar() throws {
        try check(T.content([T.upgrades(bestRun: true), .raidLine(T.raid()), T.newForYou("day")]), "home-60-noavatar")
    }

    @Test("ios-home-60-answered")
    func sixtyAnswered() throws {
        try check(T.content([T.upgrades(bestRun: true), .raidLine(T.raid(answered: .signedIn)), T.newForYou("day")]), "home-60-answered")
    }

    @Test("ios-home-noguild")
    func noGuild() throws {
        try check(
            T.content([T.upgrades(bestRun: true), T.newForYou("day", unread: 2), .findGuild(Copy.homeFindGuildLine)]),
            "home-noguild"
        )
    }

    @Test("ios-home-34")
    func leveler() throws {
        let upgrades = HomeBlockContent.upgrades(HomeUpgradesContent(
            rows: T.levelerRows, bestRun: "Best run: Blackfathom Deeps, Grimclaw +0.9 DPS", note: nil
        ))
        try check(T.content([upgrades, T.newForYou("leveler", unread: 1)], head: T.felsnap), "home-34")
    }

    @Test("ios-home-loading")
    func loading() throws {
        try check(HomeFlowContent(head: T.thandrel, phase: .loading), "home-loading")
    }

    @Test("ios-home-error")
    func error() throws {
        try check(HomeFlowContent(head: T.thandrel, phase: .failed), "home-error")
    }

    @Test("ios-home-new")
    func noCharacter() throws {
        try ParityCheck.assert(HomeNoCharacterView {}, screen: "home-new", size: Self.device)
    }
}

/// The time-aware Home's logic: the server's blocks, the moment's rules, the words.
@MainActor
@Suite("Home flow")
struct HomeFlowLogicTests {
    private static let zone = TimeZone(identifier: "America/Chicago")!

    private var presenter: HomeFlowPresenter {
        var calendar = Calendar(identifier: .gregorian)
        calendar.timeZone = Self.zone
        return HomeFlowPresenter(now: HomeFlowTestData.now, clock: RaidClock(zone: Self.zone), calendar: calendar)
    }

    private func sections(_ types: [(String, String?)]) -> [HomeSection] {
        types.map { HomeSection(type: $0.0, limit: $0.0 == "upgrades" || $0.0 == "activity" ? 3 : nil, style: $0.1) }
    }

    @Test("the blocks follow the server's order and skip types the app does not know")
    func blocks() {
        let blocks = HomeBlock.of(sections([("head", nil), ("next_raid", "card"), ("upgrades", nil), ("feed", nil), ("activity", nil)]))
        #expect(blocks == [
            .head, .nextRaid(card: true), .upgrades(limit: 3),
            .activity(HomeActivityQuery(limit: 3, group: nil, exclude: [])),
        ])
        #expect(!HomeBlock.showsBestRun(blocks))
        #expect(HomeBlock.showsBestRun(HomeBlock.of(sections([("upgrades", nil), ("next_raid", "line")]))))
    }

    @Test("the fixture's Home decodes into its blocks")
    func fixture() throws {
        let envelope = try WireFormat.decoder().decode(Envelope<HomeBlocks>.self, from: try HomeFixtures.data("home"))
        let blocks = HomeBlock.of(try #require(envelope.data).sections)
        #expect(blocks.first == .head)
        #expect(blocks.contains(.nextRaid(card: false)))
    }

    @Test("New for you reads past the kinds Last night shows, to fill its rows")
    func excludes() {
        let rows = ActivityTestData.rows
        let query = HomeActivityQuery(limit: 3, group: nil, exclude: ["loot.awarded", "report.ready"])
        #expect(query.rows(rows).map(\.id) == [8, 5, 4])
    }

    @Test("the head: level and spec, then the guild at 60 or the levels to the next band")
    func head() {
        let sixty = HomeCharacter(
            key: "us/normal/thandrel", region: "us", ruleset: "normal", name: "Thandrel", race: "human", className: "warrior",
            specName: "Fury", level: 60, guild: "Olympus XXVII", faction: "alliance"
        )
        #expect(presenter.head(sixty).facts == ["Level 60 Fury Warrior", "Olympus XXVII"])
        let leveler = HomeCharacter(
            key: "us/normal/felsnap", region: "us", ruleset: "normal", name: "Felsnap", race: "dwarf", className: "hunter",
            specName: "Beast Mastery", level: 34, guild: nil, faction: "alliance"
        )
        #expect(presenter.head(leveler).facts == ["Level 34 Beast Mastery Hunter", "6 levels to band 40"])
    }

    @Test("the countdown: hours and minutes to the pull, none once it has started")
    func countdown() {
        let start = HomeFlowTestData.now.addingTimeInterval(2 * 3600 + 48 * 60 + 30)
        #expect(presenter.clockParts(to: start) == [FSClockPart("02", "hours"), FSClockPart("48", "min")])
        #expect(presenter.clockParts(to: HomeFlowTestData.now.addingTimeInterval(-60)).isEmpty)
    }

    @Test("the raid in progress wins over the next one; past raids are never next")
    func currentRaid() throws {
        let raid = try RaidFixture.raid()
        let now = raid.startsAt.addingTimeInterval(3600)
        #expect(HomeFlowModel.current([raid], now: now)?.id == raid.id)
        #expect(HomeFlowModel.current([raid], now: raid.startsAt.addingTimeInterval(5 * 3600)) == nil)
        #expect(HomeFlowModel.current([raid], now: raid.startsAt.addingTimeInterval(-60))?.id == raid.id)
    }

    @Test("Best run: the dungeon pick or runner-up with the largest gain, its instance and chance")
    func bestRun() throws {
        // Head: a crafted pick with a dungeon runner-up 0.4 behind; finger: a dungeon pick; the
        // runner-up in Maraudon is above a level 34's reach.
        let slots = try WireFormat.decoder().decode([BisSlot].self, from: Data("""
        [{"slot":"head","verified":true,"item_id":1,"item_name":"Lionheart Helm","source_kind":"crafted","alternatives":[
           {"item_id":2,"item_name":"Blackcrow","source":"Blackrock Spire","source_kind":"dungeon","dps_delta":-0.4}]},
         {"slot":"finger","verified":true,"item_id":3,"item_name":"Ring of Precision","source_kind":"dungeon","alternatives":[
           {"item_id":4,"item_name":"Far Ring","source":"Maraudon","source_kind":"dungeon","dps_delta":0.5}]}]
        """.utf8))
        let loot = try WireFormat.decoder().decode(LootFile.self, from: Data("""
        {"sources":[
          {"id":"dungeon:brs","kind":"dungeon","name":"Blackrock Spire","level_min":30,"bosses":[
            {"id":"b1","name":"Boss","items":[2],"item_chances":{"2":8.2}}]},
          {"id":"dungeon:bfd","kind":"dungeon","name":"Blackfathom Deeps","level_min":20,"items":[3]},
          {"id":"dungeon:mara","kind":"dungeon","name":"Maraudon","level_min":40,"items":[4]}]}
        """.utf8))
        let run = try #require(BestRun.best(slots: slots, gains: ["head": 2.3, "finger": 0.6], loot: loot, level: 34))
        #expect(run == BestRun(place: "Blackrock Spire", item: "Blackcrow", gain: 2.3 - 0.4, chance: 8.2))
        let finger = try #require(BestRun.best(slots: slots, gains: ["finger": 0.6], loot: loot, level: 34))
        #expect(finger.place == "Blackfathom Deeps")
        #expect(BestRun.best(slots: slots, gains: [:], loot: loot, level: 34) == nil)
    }

    @Test("Best run's line names the chance when the loot file has one")
    func bestRunLine() {
        let run = BestRun(place: "Blackrock Spire", item: "Blackcrow", gain: 1.94, chance: 8.2)
        #expect(presenter.bestRunLine(run, unit: "DPS") == "Best run: Blackrock Spire, Blackcrow +1.9 DPS, 8%")
        let none = BestRun(place: "Blackfathom Deeps", item: "Ring of Precision", gain: 0.6, chance: nil)
        #expect(presenter.bestRunLine(none, unit: "DPS") == "Best run: Blackfathom Deeps, Ring of Precision +0.6 DPS")
    }

    @Test("before the pull: the items not ok, in their own words, with the export's age")
    func ready() {
        func item(_ key: String, _ state: String, _ line: String) -> ReadyItem {
            ReadyItem(key: key, state: state, title: key, line: line, whereToGet: nil)
        }
        let check = ReadyCheck(exportedAt: HomeFlowTestData.now.addingTimeInterval(-9 * 3600), items: [
            item("attune.onyxias-lair", "ok", "Attuned"), item("flask", "miss", "No flask in your bags"), item("enchant.feet", "warn", "boots not enchanted"),
        ])
        let line = presenter.ready(check)
        #expect(line?.title == "2 things before you pull")
        #expect(line?.sub == "No flask in your bags; boots not enchanted · from your last log-out, 9 h ago")
        #expect(presenter.ready(ReadyCheck(exportedAt: nil, items: [item("flask", "ok", "Flask")])) == nil)
    }

    @Test("Last night's loot line carries the gain over the worn piece when the night has one")
    func lootLine() throws {
        let night = try RaidFixture.load(RaidNightSummary.self, "me.raid-night")
        let loot = try #require(night.loot.first)
        #expect(presenter.lootLine(loot).hasPrefix("You got it from \(loot.boss) · \(loot.slot.lowercased())"))
        let content = presenter.lastNight(night)
        #expect(content.reportIcon == (night.zoneIcon ?? RaidIcons.generic))
    }

    @Test("rows read measured gains largest first, then not modelled, then empty")
    func presentationOrder() throws {
        let rows = HomePresenter.presentationOrder(try HomeFixtures.home().upgrades)
        let measured = rows.prefix { $0.modelled && $0.gain != nil }.compactMap(\.gain)
        #expect(measured == measured.sorted(by: >))
    }
}
