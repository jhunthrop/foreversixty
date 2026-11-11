@testable import ForeverSixty
import Foundation
import Models
import Testing

/// One Activity row on Activity and Home's New for you (round 2): the feed's styled body, the other
/// player's crest, the trailing figure; and the Raid tab's ready icons and Loot history line.
@MainActor
@Suite("Activity rows")
struct ActivityRowTests {
    private func rows() throws -> [ActivityRow] { try RaidFixture.load(ActivityPage.self, "activity.list").rows }

    private func row(_ kind: String) throws -> ActivityRow {
        try #require(try rows().first { $0.kind == kind }, "no \(kind) row in activity.list")
    }

    @Test("loot: the winner's name in class colour, the item in its quality; the text is the body")
    func loot() throws {
        let loot = try row("loot.awarded")
        let parts = try #require(ActivityBody.parts(loot))
        #expect(parts.contains(.init(text: "Gorvash", tone: .className("warrior"))))
        #expect(parts.contains(.init(text: "Deathbringer", tone: .quality(4))))
        #expect(parts.map(\.text).joined() == loot.body)
        #expect(ActivityBody.figure(loot) == nil)
    }

    @Test("the guild ladder: strong figures, the passer's crest, the figure trailing")
    func guildsim() throws {
        let passed = try row("guildsim.passed")
        let parts = try #require(ActivityBody.parts(passed))
        #expect(parts.contains(.init(text: "12", tone: .strong)))
        #expect(ActivityCrest.className(passed, specs: .main) == "warrior")
        // Row 11: the figure once, worded as the others; the body no longer repeats it.
        #expect(ActivityBody.figure(passed) == "204 DPS")
        #expect(!passed.body.contains("203.6"))
    }

    @Test("a row without segments keeps its plain body; a sim keeps its figure")
    func plain() throws {
        let sim = try row("sim.finished")
        #expect(ActivityBody.parts(sim) == nil)
        #expect(ActivityBody.figure(sim) == "190 DPS")
    }

    @Test("Mark all read keeps a row's segments, crest and figure")
    func asRead() throws {
        let passed = try row("guildsim.passed")
        let read = passed.asRead
        #expect(read.read)
        #expect(read.segments == passed.segments)
        #expect(read.crest == passed.crest)
        #expect(read.figure == passed.figure)
    }

    @Test("Before you pull: the item's icon beside a check that names one; none for a consumable without an icon")
    func readyIcons() throws {
        let check = try RaidFixture.load(ReadyCheck.self, "character.ready")
        let ready = try #require(RaidFlowPresenter(now: RaidTime.now, clock: RaidTime.clock).ready(check))
        let icons = Dictionary(uniqueKeysWithValues: ready.rows.map { ($0.key, $0.icon) })
        #expect(icons["attune.onyxias-lair"] == "inv_jewelry_talisman_11")
        #expect(icons["enchant.feet"] == "inv_boots_07")
        #expect(icons["food"] == .some(nil))
        #expect(icons["flask"] == .some(nil))
    }

    @Test("Loot history: the last night's zone, kills and wipes, then the last raid's pieces")
    func lootLine() throws {
        let list = try RaidFixture.load(GuildRaids.self, "guild.raids")
        let night = try #require(list.rows.first { $0.raidId != nil })
        let presenter = RaidFlowPresenter(now: RaidTime.now, clock: RaidTime.clock)
        let line = try #require(presenter.lootLine(badge: nil, lastNight: night))
        #expect(line == Copy.raidPagesLastNight(zone: night.zone, kills: Copy.raidPagesKills(n: night.killCount), wipes: Copy.raidPagesWipes(n: night.wipeCount)))
        #expect(line == "Last: Onyxia's Lair, 1 kill, 2 wipes")
        #expect(presenter.lootLine(badge: nil, lastNight: nil) == nil)
    }
}
