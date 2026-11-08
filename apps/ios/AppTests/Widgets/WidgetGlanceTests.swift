@testable import ForeverSixty
import Foundation
import Models
import Testing

/// The restyled widgets' and Live Activity's words (flows F1, F2, F5, E7), from the API's fixtures.
@Suite("Widget glance words")
struct WidgetGlanceTests {
    private let zone = WidgetTime.chicago
    private let now = WidgetTime.now

    private func without(raid: Bool = false, guild: Bool = false, zones: Bool = true) throws -> WidgetSnapshot {
        let base = try WidgetFixtures.widget("widget.get")
        return WidgetSnapshot(widget: Widget(
            generatedAt: base.generatedAt, firstRaidsAt: base.firstRaidsAt, firstRaidsZones: zones ? base.firstRaidsZones : nil,
            character: base.character, guild: guild ? nil : base.guild, nextRaid: raid || guild ? nil : base.nextRaid,
            lastRaid: base.lastRaid, upgrades: base.upgrades, live: base.live
        ), fetchedAt: now)
    }

    /// A push fixture's `aps.content-state`, decoded as ActivityKit decodes it.
    private func contentState(_ fixture: String) throws -> RaidActivityAttributes.ContentState {
        let root = try JSONSerialization.jsonObject(with: RaidFixture.data(fixture)) as? [String: Any]
        let aps = try #require(root?["aps"] as? [String: Any])
        let state = try JSONSerialization.data(withJSONObject: try #require(aps["content-state"]))
        return try JSONDecoder().decode(RaidActivityAttributes.ContentState.self, from: state)
    }

    @Test("a raid: the guild, the countdown, the time with its zone, the answer as the character (E7)")
    func raid() throws {
        let glance = RaidGlance(try WidgetFixtures.snapshot(), now: now, zone: zone)
        #expect(glance.kind == .raid)
        #expect(glance.label == "OLYMPUS XXVII")
        #expect(glance.big == "2d 2h")
        #expect(glance.name == "Onyxia's Lair")
        #expect(glance.lines == ["Thu 10 Dec, 8:00 pm CST", "In, as Lightningbolltz, Damage"])
        #expect(glance.selected == "in" && glance.canAnswer && !glance.asksSpec)
    }

    @Test("unanswered, healer: Not answered names the role; no spec asks for one")
    func unanswered() throws {
        let mine: [String: JSONValue] = ["my": .object(["state": .null, "spec": .string("Restoration"), "role": .string("healer")])]
        let glance = RaidGlance(try WidgetFixtures.snapshot(patch: mine), now: now, zone: zone)
        #expect(glance.lines.last == "Not answered, as Lightningbolltz, Healer")
        #expect(glance.selected == nil)
        let noSpec: [String: JSONValue] = ["my": .object(["state": .null, "spec": .null, "role": .string("dps")])]
        #expect(RaidGlance(try WidgetFixtures.snapshot(patch: noSpec), now: now, zone: zone).asksSpec)
    }

    @Test("under a day the figure is hours and minutes; a locked raid cannot be answered")
    func closeAndLocked() throws {
        let close = RaidGlance(try WidgetFixtures.snapshot(), now: RaidTime.date("2026-12-10T22:20:00Z"), zone: zone)
        #expect(close.big == "3h 40m")
        let locked = RaidGlance(try WidgetFixtures.snapshot(patch: ["status": .string("locked"), "can_sign_up": .bool(false)]), now: now, zone: zone)
        #expect(!locked.canAnswer && locked.lines.contains(Copy.widgetLocked))
    }

    @Test("no guild before the first raids: the countdown and the API's zones; after: Raids open now")
    func firstRaids() throws {
        let zones = "Barrow Deeps, Hyjal Summit, Onyxia's Lair"
        let before = RaidGlance(try without(guild: true), now: RaidTime.date("2026-11-06T17:39:00Z"), zone: zone)
        #expect(before.kind == .firstRaids && before.label == Copy.widgetFirstRaids)
        #expect(before.big == "32d 6h" && before.name == "Wed 9 Dec" && before.lines == [zones])
        let after = RaidGlance(try without(guild: true), now: now.addingTimeInterval(86_400 * 3), zone: zone)
        #expect(after.kind == .raidsOpen && after.big == Copy.widgetOpen && after.name == zones && after.lines == [Copy.widgetFindGuild])
        let guildNoRaid = RaidGlance(try without(raid: true), now: now.addingTimeInterval(86_400 * 3), zone: zone)
        #expect(guildNoRaid.lines == [Copy.widgetNoRaidYet])
        let noZones = RaidGlance(try without(guild: true, zones: false), now: now.addingTimeInterval(86_400 * 3), zone: zone)
        #expect(noZones.name == Copy.widgetFindGuild && noZones.lines.isEmpty)
    }

    @Test("signed out and no character each say what to do")
    func empty() {
        let out = RaidGlance(nil, now: now, zone: zone)
        #expect(out.kind == .empty && out.name == Copy.widgetSignedOut && out.lines == [Copy.widgetOpenApp])
    }

    @Test("the upgrade side is the next upgrade with its source, never the total")
    func upgrade() throws {
        let glance = UpgradeGlance(try WidgetFixtures.snapshot())
        #expect(glance.gain == "+5.3 DPS")
        #expect(glance.item == UpgradeGlance.Item(name: "Knowledge of the Timbermaw", quality: 4, icon: "inv_belt_09"))
        #expect(glance.source == "Blacksmithing")
        #expect(!glance.spoken.contains("18.3"))
        #expect(UpgradeGlance(nil).gain == nil && UpgradeGlance(nil).sentence == Copy.widgetSignedOut)
    }

    @Test("the Live Activity started from the widget: the API's icon and counts, the answer as the character")
    func live() throws {
        let raid = try #require(try WidgetFixtures.snapshot().widget.nextRaid)
        let attributes = RaidActivityAttributes(raid, characterName: "Thandrel")
        var state = RaidActivityAttributes.ContentState(raid, myState: "in")
        #expect(attributes.zoneIcon == "inv_misc_head_dragon_black")
        let figures = RaidLiveFigures(attributes, state, zone: zone)
        #expect(figures.starts == "Starts 8:00 pm CST · 24 in, 0 late · 2 tanks, 5 heals")
        #expect(figures.answer == "In, as Thandrel, Damage")
        #expect(figures.short == "Ony")
        #expect(!figures.done)
        state.inCount = nil
        state.tanks = nil
        #expect(RaidLiveFigures(attributes, state, zone: zone).starts == "Starts 8:00 pm CST · 24 of 40 signed up")
    }

    @Test("the server's start push decodes with its counts")
    func startPush() throws {
        let state = try contentState("surface.live-raid-start.apns")
        #expect(state.inCount == 36 && state.late == 2 && state.tanks == 2 && state.healers == 5)
        #expect(state.status == "open" && state.loot == nil)
    }

    @Test("the server's end push: done, the kills, every item in its colour, the report")
    func endPush() throws {
        let state = try contentState("surface.live-raid-end.apns")
        let raid = try #require(try WidgetFixtures.snapshot().widget.nextRaid)
        let figures = RaidLiveFigures(RaidActivityAttributes(raid, characterName: "Thandrel"), state, zone: zone)
        #expect(figures.done && figures.doneTitle == "Onyxia's Lair done")
        #expect(String(figures.doneLine.characters) == "1 kill · you got Sapphiron Drape · report ready")
        var two = state
        two.loot = (state.loot ?? []) + [RaidLiveLoot(name: "Deathbringer", quality: 4, icon: "inv_axe_09")]
        let twoLine = RaidLiveFigures(RaidActivityAttributes(raid, characterName: "Thandrel"), two, zone: zone).doneLine
        #expect(String(twoLine.characters) == "1 kill · you got Sapphiron Drape, Deathbringer · report ready")
    }
}
