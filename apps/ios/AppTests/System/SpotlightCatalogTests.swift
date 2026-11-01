@testable import ForeverSixty
import Foundation
import Models
import Networking
import Routing
import Testing

@Suite("Spotlight catalog")
struct SpotlightCatalogTests {
    private let answers = SystemStory.answers

    private func character(_ name: String, key: String, guild: Bool = true) throws -> AccountCharacter {
        let me = try RaidFixture.load(Me.self, "me")
        let base = try #require(me.characters.first)
        var json = try JSONSerialization.jsonObject(with: WireFormatTesting.encode(base)) as? [String: Any] ?? [:]
        json["key"] = key
        json["name"] = name
        json["class"] = "shaman"
        json["spec"] = "Elemental"
        json["level"] = 60
        if !guild { json["guild"] = NSNull() }
        return try WireFormatTesting.decode(AccountCharacter.self, json)
    }

    @Test("your characters with spec, class, level and guild")
    func own() throws {
        let row = SpotlightCatalog.row(try character("Lightningbolltz", key: SystemStory.key))
        #expect(row == SpotlightRow(id: SystemStory.key, title: "Lightningbolltz", subtitle: "Elemental Shaman · 60 · OLYMPUS XXVII"))
        let loner = SpotlightCatalog.row(try character("Lightningbolltz", key: SystemStory.key, guild: false))
        #expect(loner.subtitle == "Elemental Shaman · 60")
    }

    @Test("a watched character has no spec on the list")
    func watched() {
        let watched = WatchedCharacter(key: "us/pvp/morvane", region: "us", ruleset: "pvp", name: "Morvane", className: "warlock", level: 60, addedAt: SystemStory.now)
        #expect(SpotlightCatalog.row(watched).subtitle == "Warlock · 60")
    }

    @Test("caps: 9 of yours, 50 watched, 5 raids; your own never twice")
    func caps() throws {
        let own = try (0..<12).map { try character("C\($0)", key: "us/pvp/c\($0)") }
        let watched = (0..<60).map {
            WatchedCharacter(key: "us/pvp/w\($0)", region: "us", ruleset: "pvp", name: "W\($0)", className: nil, level: nil, addedAt: SystemStory.now)
        } + [WatchedCharacter(key: "us/pvp/c0", region: "us", ruleset: "pvp", name: "C0", className: nil, level: nil, addedAt: SystemStory.now)]
        let raids = try (0..<8).map { try SystemStory.raid(id: Int64(10 + $0), startsAt: "2026-12-\(10 + $0)T02:00:00Z") }
        let catalog = SpotlightCatalog.build(own: own, watched: watched, planned: raids, guildID: 2, guides: [], answers: answers, now: SystemStory.now)
        #expect(catalog.characters.count == 59)
        #expect(catalog.raids.count == 5)
        #expect(catalog.raids.first == SpotlightRow(id: "2/10", title: "Onyxia's Lair · Wed 8 pm", subtitle: "Raid · Wed 9 Dec"))
    }

    @Test("a saved guide")
    func guide() throws {
        let spec = try #require(SpecCatalog.bundled().spec("shaman-elemental"))
        let row = SpotlightCatalog.guideRow(.init(spec: spec, savedAt: RaidTime.date("2026-09-24T15:00:00Z")), clock: answers.clock)
        #expect(row == SpotlightRow(id: "shaman-elemental", title: "Elemental Shaman guide", subtitle: "Guide · Updated 24 Sep 2026"))
    }

    @Test("the entities open what they name")
    func links() {
        #expect(RaidEntity(id: "2/7", title: "", subtitle: "").link == .raid(guild: 2, raid: 7))
        #expect(GuideEntity(id: "shaman-elemental", title: "", subtitle: "").link == .guide(className: "shaman", spec: "elemental"))
    }
}

/// The wire format both ways, for patching an API model in a test.
enum WireFormatTesting {
    static func encode<T: Encodable>(_ value: T) throws -> Data { try WireFormat.encoder().encode(value) }

    static func decode<T: Decodable>(_ type: T.Type, _ json: [String: Any]) throws -> T {
        try WireFormat.decoder().decode(type, from: JSONSerialization.data(withJSONObject: json))
    }
}
