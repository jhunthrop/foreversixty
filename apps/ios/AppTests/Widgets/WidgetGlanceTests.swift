@testable import ForeverSixty
import Foundation
import Models
import Testing

/// The restyled widgets' and Live Activity's words (flows F1, F2, F5, E7).
@Suite("Widget glance words")
struct WidgetGlanceTests {
    private let zone = WidgetTime.chicago
    private let now = WidgetTime.now

    private func without(raid: Bool = false, guild: Bool = false, firstRaidsAt: String? = "2026-12-09") throws -> WidgetSnapshot {
        let base = try WidgetFixtures.widget("widget.get")
        return WidgetSnapshot(widget: Widget(
            generatedAt: base.generatedAt, firstRaidsAt: firstRaidsAt, character: base.character, guild: guild ? nil : base.guild,
            nextRaid: raid || guild ? nil : base.nextRaid, lastRaid: base.lastRaid, upgrades: base.upgrades, live: base.live
        ), fetchedAt: now)
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

    @Test("no guild before the first raids: the countdown to them; after: Raids open now")
    func firstRaids() throws {
        let before = RaidGlance(try without(guild: true), now: RaidTime.date("2026-11-06T17:39:00Z"), zone: zone)
        #expect(before.kind == .firstRaids && before.label == Copy.widgetFirstRaids)
        #expect(before.big == "32d 6h" && before.name == "Wed 9 Dec" && before.lines == [Copy.widgetFirstRaidsZones])
        let after = RaidGlance(try without(guild: true), now: now.addingTimeInterval(86_400 * 3), zone: zone)
        #expect(after.kind == .raidsOpen && after.big == Copy.widgetOpen && after.lines == [Copy.widgetFindGuild])
        let guildNoRaid = RaidGlance(try without(raid: true), now: now.addingTimeInterval(86_400 * 3), zone: zone)
        #expect(guildNoRaid.lines == [Copy.widgetNoRaidYet])
    }

    @Test("signed out and no character each say what to do")
    func empty() {
        let out = RaidGlance(nil, now: now, zone: zone)
        #expect(out.kind == .empty && out.name == Copy.widgetSignedOut && out.lines == [Copy.widgetOpenApp])
    }

    @Test("the upgrade side is the next upgrade, never the total")
    func upgrade() throws {
        let glance = UpgradeGlance(try WidgetFixtures.snapshot())
        #expect(glance.gain == "+5.3 DPS")
        #expect(glance.item == UpgradeGlance.Item(name: "Knowledge of the Timbermaw", quality: 4, icon: "inv_belt_09"))
        #expect(!glance.spoken.contains("18.3"))
        #expect(UpgradeGlance(nil).gain == nil && UpgradeGlance(nil).sentence == Copy.widgetSignedOut)
    }

    @Test("the Live Activity: the start with its zone, who is coming, the answer as the character")
    func live() throws {
        let raid = try #require(try WidgetFixtures.snapshot().widget.nextRaid)
        let attributes = RaidActivityAttributes(raid, characterName: "Thandrel")
        var state = RaidActivityAttributes.ContentState(raid, myState: "in")
        #expect(attributes.icon == "inv_misc_head_dragon_black")
        #expect(RaidLiveFigures(attributes, state, zone: zone).starts == "Starts 8:00 pm CST · 24 of 40 signed up")
        state.inCount = 18
        state.late = 3
        state.tanks = 2
        state.healers = 5
        let figures = RaidLiveFigures(attributes, state, zone: zone)
        #expect(figures.starts == "Starts 8:00 pm CST · 18 in, 3 late · 2 tanks, 5 heals")
        #expect(figures.answer == "In, as Thandrel, Damage")
        #expect(figures.short == "Ony")
        #expect(!figures.done)
    }

    @Test("the raid's end: done, the kills, the loot, the report")
    func liveDone() throws {
        let raid = try #require(try WidgetFixtures.snapshot().widget.nextRaid)
        var state = RaidActivityAttributes.ContentState(raid, myState: "in")
        state.status = RaidActivityAttributes.ContentState.done
        state.kills = 1
        state.loot = RaidLiveLoot(name: "Deathbringer", quality: 4, icon: "inv_axe_09")
        state.reportReady = true
        let figures = RaidLiveFigures(RaidActivityAttributes(raid, characterName: "Thandrel"), state, zone: zone)
        #expect(figures.done && figures.doneTitle == "Onyxia's Lair done")
        #expect(String(figures.doneLine.characters) == "1 kill · you got Deathbringer · report ready")
    }

    @Test("the content-state decodes the server's start push, and the new fields when sent")
    func decodes() throws {
        let start = Data(#"{"starts_at":1796954400,"invite_at":null,"status":"open","my_state":"in","role":"dps","signed_up":38,"size":40}"#.utf8)
        let state = try JSONDecoder().decode(RaidActivityAttributes.ContentState.self, from: start)
        #expect(state.inCount == nil && state.signedUp == 38)
        let full = Data(#"{"starts_at":1,"status":"done","role":"dps","signed_up":21,"in":18,"late":3,"tanks":2,"healers":5,"kills":1,"loot":{"name":"Deathbringer","quality":4,"icon":"inv_axe_09"},"report_ready":true}"#.utf8)
        let done = try JSONDecoder().decode(RaidActivityAttributes.ContentState.self, from: full)
        #expect(done.inCount == 18 && done.loot?.name == "Deathbringer" && done.reportReady == true)
    }
}
