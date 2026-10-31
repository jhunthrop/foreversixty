@testable import ForeverSixty
import Foundation
import Models
import Networking
import Testing

/// apps/shared/widgets/figures.json: the golden vectors both apps hold their widget words to.
struct FigureVectors: Decodable {
    struct Case: Decodable {
        let name: String
        let snapshot: String
        let now: String
        let fetchedAt: String
        let zone: String
        let patch: [String: JSONValue]?
        let expect: [String: JSONValue]

        enum CodingKeys: String, CodingKey {
            case name, snapshot, now, zone, patch, expect
            case fetchedAt = "fetched_at"
        }
    }

    struct Upgrades: Decodable {
        struct Expect: Decodable {
            let header: String
            let rows: [[String]]
            let footer: String
            let rectTitle: String
        }

        struct NoGear: Decodable {
            let snapshot: String
            let sentence: String
        }

        let snapshot: String
        let expect: Expect
        let noGear: NoGear
    }

    let cases: [Case]
    let upgrades: Upgrades

    static func load() throws -> FigureVectors {
        let url = RaidFixture.directory
            .deletingLastPathComponent().deletingLastPathComponent().deletingLastPathComponent().deletingLastPathComponent()
            .appending(path: "apps/shared/widgets/figures.json")
        return try JSONDecoder().decode(FigureVectors.self, from: Data(contentsOf: url))
    }
}

enum WidgetFixtures {
    /// The fixture's data, with `patch` laid over data.next_raid.
    static func widget(_ name: String, patch: [String: JSONValue]? = nil) throws -> Widget {
        var root = try JSONSerialization.jsonObject(with: RaidFixture.data(name)) as? [String: Any] ?? [:]
        var data = root["data"] as? [String: Any] ?? [:]
        if let patch, var raid = data["next_raid"] as? [String: Any] {
            for (key, value) in patch { raid[key] = try JSONSerialization.jsonObject(with: JSONEncoder().encode(value), options: .fragmentsAllowed) }
            data["next_raid"] = raid
        }
        root["data"] = data
        let patched = try JSONSerialization.data(withJSONObject: root)
        return try WireFormat.decoder().decode(Envelope.self, from: patched).data
    }

    static func snapshot(_ name: String = "widget.get", fetchedAt: Date = WidgetTime.now, patch: [String: JSONValue]? = nil) throws -> WidgetSnapshot {
        WidgetSnapshot(widget: try widget(name, patch: patch), fetchedAt: fetchedAt)
    }

    private struct Envelope: Decodable { let data: Widget }
}

enum WidgetTime {
    static let chicago = TimeZone(identifier: "America/Chicago") ?? .gmt
    /// The board's moment: Tuesday 8 December, 5:39 pm in Chicago, 2 days 2 hours before the pull.
    static let now = RaidTime.date("2026-12-08T23:39:00Z")
}

@Suite("Widget figures: the golden vectors")
struct WidgetFiguresTests {
    private func string(_ value: JSONValue?) -> String?? {
        switch value {
        case .none: return .none
        case .some(.null): return .some(nil)
        case let .some(.string(text)): return .some(text)
        default: return .none
        }
    }

    private func bool(_ value: JSONValue?) -> Bool? {
        if case let .bool(flag) = value { return flag }
        return nil
    }

    @Test("every surface's words match figures.json")
    func vectors() throws {
        let vectors = try FigureVectors.load()
        #expect(vectors.cases.count >= 10)
        for item in vectors.cases {
            let snapshot = try WidgetFixtures.snapshot(item.snapshot, fetchedAt: RaidTime.date(item.fetchedAt), patch: item.patch)
            let figures = WidgetFigures(snapshot, now: RaidTime.date(item.now), zone: TimeZone(identifier: item.zone) ?? .gmt)
            let got: [String: String?] = [
                "tag": figures.tag, "tagKind": figures.tagKind?.rawValue, "countdown": figures.countdown, "short": figures.short,
                "time": figures.time, "timeShort": figures.timeShort, "inline": figures.inline, "you": figures.you,
                "signedUp": figures.signedUp, "stale": figures.stale, "liveTitle": figures.liveTitle, "liveUntil": figures.liveUntil,
            ]
            for (key, value) in got {
                guard let expected = string(item.expect[key]) else { continue }
                #expect(value == expected, "\(item.name): \(key)")
            }
            if let state = string(item.expect["state"]) ?? nil { #expect("\(figures.state)" == state, "\(item.name): state") }
            if let buttons = bool(item.expect["buttons"]) { #expect(figures.buttons == buttons, "\(item.name): buttons") }
            if let live = bool(item.expect["liveActive"]) { #expect(figures.liveActive == live, "\(item.name): liveActive") }
        }
    }

    @Test("the upgrade widgets' words match figures.json")
    func upgrades() throws {
        let vectors = try FigureVectors.load().upgrades
        let figures = UpgradeFigures(try WidgetFixtures.snapshot(vectors.snapshot))
        #expect(figures.header == vectors.expect.header)
        #expect(figures.rows.map { [$0.name, $0.gain] } == vectors.expect.rows)
        #expect(figures.footer == vectors.expect.footer)
        #expect(figures.rectTitle == vectors.expect.rectTitle)
        let noGear = UpgradeFigures(try WidgetFixtures.snapshot(vectors.noGear.snapshot))
        #expect(noGear.rows.isEmpty && noGear.sentence == vectors.noGear.sentence && !noGear.opensApp)
    }

    @Test("signed out, and no Home answer yet, each have words")
    func emptyStates() throws {
        let signedOut = WidgetFigures(nil, now: WidgetTime.now, zone: WidgetTime.chicago)
        #expect(signedOut.state == .signedOut && signedOut.inline == Copy.widgetSignedOut)
        #expect(UpgradeFigures(nil).sentence == Copy.widgetSignedOut)
        let base = try WidgetFixtures.widget("widget.get")
        let noUpgrades = Widget(
            generatedAt: base.generatedAt, character: base.character, guild: base.guild, nextRaid: base.nextRaid,
            lastRaid: base.lastRaid, upgrades: nil, live: base.live
        )
        let figures = UpgradeFigures(WidgetSnapshot(widget: noUpgrades, fetchedAt: WidgetTime.now))
        #expect(figures.sentence == Copy.widgetOpenApp && figures.opensApp)
        let noRaid = Widget(
            generatedAt: base.generatedAt, character: base.character, guild: base.guild, nextRaid: nil,
            lastRaid: base.lastRaid, upgrades: base.upgrades, live: base.live
        )
        #expect(WidgetFigures(WidgetSnapshot(widget: noRaid, fetchedAt: WidgetTime.now), now: WidgetTime.now, zone: WidgetTime.chicago).state == .noRaid)
    }

    @Test("every surface agrees: lock screen, home screen and Live Activity read one snapshot")
    func surfacesAgree() throws {
        let snapshot = try WidgetFixtures.snapshot(patch: ["signed_up": .number(38)])
        let now = RaidTime.date("2026-12-11T01:21:48Z")
        let widget = WidgetFigures(snapshot, now: now, zone: WidgetTime.chicago)
        let raid = try #require(snapshot.widget.nextRaid)
        let attributes = RaidActivityAttributes(raid, characterName: "Lightningbolltz")
        let state = RaidActivityAttributes.ContentState(raid, myState: widget.checked)
        let live = WidgetFigures(.live(attributes, state), now: now, zone: WidgetTime.chicago)
        #expect(live.signedUp == widget.signedUp)
        #expect(live.countdown == widget.countdown)
        #expect(live.short == widget.short)
        #expect(live.liveTitle == widget.liveTitle)
        #expect(live.liveUntil == widget.liveUntil)
        #expect(live.inline == widget.inline)
    }

    @Test("a pending press draws at once, and Not saved shows for ten seconds")
    func pendingAndNotSaved() throws {
        var snapshot = try WidgetFixtures.snapshot()
        snapshot.pending = PendingAnswer(raidID: try #require(snapshot.widget.nextRaid).raidId, state: "late")
        #expect(WidgetFigures(snapshot, now: WidgetTime.now, zone: WidgetTime.chicago).checked == "late")
        snapshot.pending = nil
        snapshot.notSavedUntil = WidgetTime.now.addingTimeInterval(10)
        snapshot.notSavedState = "out"
        #expect(WidgetFigures(snapshot, now: WidgetTime.now.addingTimeInterval(9), zone: WidgetTime.chicago).notSaved == "out")
        #expect(WidgetFigures(snapshot, now: WidgetTime.now.addingTimeInterval(11), zone: WidgetTime.chicago).notSaved == nil)
    }
}
