@testable import ForeverSixty
import Foundation
import Models
import Testing

/// The Guild tab's sections (spec 2026-11-10-guild-tab.md section 2) from the roster fixture
/// (guild.home.roster: 38 members, 30 at 60, 5 at 55 to 59).
@Suite("Guild tab presenter")
struct GuildTabPresenterTests {
    /// The day after the fixture's newest seen_at (2026-12-16 21:40).
    static let now = Date(timeIntervalSince1970: 1_797_508_800)
    let presenter = GuildTabPresenter(now: Self.now, clock: RaidClock(zone: TimeZone(identifier: "UTC")!))

    private func inputs(level: Int, officer: GuildSetup? = nil) throws -> GuildTabInputs {
        var inputs = GuildTabInputs(home: try GuildFixture.home())
        inputs.level = level
        inputs.characterKey = GuildFixture.viewerKey
        inputs.setup = officer
        return inputs
    }

    @Test("below 60: the level spread and Close to 60")
    func leveler() throws {
        let content = try #require(presenter.content(try inputs(level: 36), guildName: "OLYMPUS XXVII"))
        #expect(content.spread == GuildSpread(atCap: 30, close: 5, third: 36, thirdLabel: "your level", own: true))
        #expect(content.leveler)
        #expect(content.close.allSatisfy { Int($0.trail).map { (55..<60).contains($0) } ?? false })
        // Close to 60 never repeats a name Recently on shows.
        #expect(Set(content.close.map(\.name)).isDisjoint(with: content.recent.map(\.name)))
        #expect(content.loot == nil)
        #expect(content.rosterMeta == "30 of 38 at 60")
    }

    @Test("at 60: the raiders figure in place of your level, no Close to 60")
    func atCap() throws {
        let content = try #require(presenter.content(try inputs(level: 60), guildName: "OLYMPUS XXVII"))
        #expect(content.spread?.own == false)
        #expect(!content.leveler)
        #expect(content.close.isEmpty)
        #expect(content.recent.count == GuildTabPresenter.recentShown)
    }

    @Test("Recently on is newest first and leaves the viewer out")
    func recent() throws {
        let content = try #require(presenter.content(try inputs(level: 60), guildName: ""))
        let home = try GuildFixture.home()
        let viewer = home.members.first { $0.characterKey == GuildFixture.viewerKey }?.name
        #expect(!content.recent.contains { $0.name == viewer })
        #expect(content.recent.allSatisfy { $0.trail.hasSuffix("ago") || $0.trail == "just now" })
    }

    @Test("the officer rows follow viewer.officer_tools")
    func officerRows() throws {
        let member = try #require(presenter.content(try inputs(level: 60), guildName: ""))
        #expect(member.officer.isEmpty)
    }

    @Test("the pages left: Guild sims at 60, Loot history under 60, never Roster")
    func pagesLeft() {
        let atCap = RaidRootComposer.pages(RaidFlowTestData.pages, tab: GuildTabTestData.tab(leveler: false)).rows.map(\.page)
        let leveler = RaidRootComposer.pages(RaidFlowTestData.pages, tab: GuildTabTestData.tab(leveler: true)).rows.map(\.page)
        #expect(atCap == [.calendar, .sims])
        #expect(leveler == [.calendar, .loot])
    }

    @Test("seen_at reads the member's local clock in the viewer's zone")
    func seenAt() {
        let utc = TimeZone(identifier: "UTC")!
        #expect(GuildTabPresenter.seenAt("2026-12-16 21:40", zone: utc) == Date(timeIntervalSince1970: 1_797_457_200))
        #expect(GuildTabPresenter.seenAt("yesterday", zone: utc) == nil)
        #expect(GuildTabPresenter.seenAt(nil, zone: utc) == nil)
    }

    @Test("a report's length in hours and minutes")
    func duration() {
        #expect(GuildTabPresenter.duration(5_520_000) == "1 h 32 min")
        #expect(GuildTabPresenter.duration(2_460_000) == "41 min")
    }

    @Test("every setup step has its word")
    func steps() {
        for key in ["officer_ranks", "discord_bot", "recruiting", "raid_schedule", "invite_link", "addon_sync", "report_visibility", "linked_role"] {
            #expect(GuildTabPresenter.stepWord(key) != nil)
        }
        #expect(GuildTabPresenter.stepWord("other") == nil)
    }
}
