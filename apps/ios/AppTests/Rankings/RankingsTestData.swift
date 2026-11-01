@testable import ForeverSixty
import Foundation
import Models
import Networking

/// F11's test data: the contract fixtures (the Onyxia board, Voltaris and Lightningbolltz) and
/// variants derived from them. Every helper throws, so a missing file fails only its own test.
enum RankingsTestData {
    static let build = "1.60.1.70658"
    /// The day after the board's last kill (rankings.list `updated_at` is 17 Dec 19:56 UTC).
    static let now = RaidTime.date("2026-12-17T20:00:00Z")
    static let clock = ReportsClock(zone: TimeZone(identifier: "UTC") ?? .gmt)
    static let specs = SpecCatalog.bundled()

    struct Failure: Error, CustomStringConvertible { let description: String }

    static func page() throws -> RankingsPage { try RaidFixture.load(RankingsPage.self, "rankings.list") }
    static func offPage() throws -> RankingsPage { try RaidFixture.load(RankingsPage.self, "rankings.list.off-page") }
    static func emptyPage() throws -> RankingsPage { try RaidFixture.load(RankingsPage.self, "rankings.list.empty") }
    static func encounters() throws -> RankedEncounters { try RaidFixture.load(RankedEncounters.self, "rankings.encounters") }
    static func phases() throws -> Phases { try RaidFixture.load(Phases.self, "phases") }
    static func watching() throws -> WatchingList { try RaidFixture.load(WatchingList.self, "watching.list") }
    static func voltaris() throws -> CharacterProfile { try RaidFixture.load(CharacterProfile.self, "character.card.ranked") }
    static func lightningbolltz() throws -> CharacterProfile { try RaidFixture.load(CharacterProfile.self, "character.card.yours") }
    static func freshling() throws -> CharacterProfile { try RaidFixture.load(CharacterProfile.self, "character.card.empty") }

    /// The shaman items file, by id, for the gear sockets.
    static func items() throws -> [Int: Item] {
        let file = try WireFormat.decoder().decode(ItemsFile.self, from: RaidFixture.data("cdn.items.shaman"))
        return Dictionary(file.items.map { ($0.id, $0) }, uniquingKeysWith: { first, _ in first })
    }

    /// Lightningbolltz, the board's own character (14th of 212).
    static let lightning = RankingsCharacter(
        key: "us/normal/lightningbolltz", name: "Lightningbolltz", className: "shaman", spec: "Elemental", ruleset: "normal"
    )

    /// The signed-in account on the board: Lightningbolltz and its alt Sparkwick.
    static func account(history: Bool = false) -> RankingsAccount {
        RankingsAccount(
            signedIn: true, characterKeys: ["us/normal/lightningbolltz", "us/normal/sparkwick"], current: lightning, history: history
        )
    }

    /// The board's filters for Lightningbolltz on Onyxia, Phase 1.
    static let filters = RankingsFilters(
        encounter: "onyxia", metric: .dps, ruleset: "normal", region: "", phase: "raids-1", faction: "", todayOnly: false,
        classSlug: "shaman", spec: "Elemental"
    )

    /// A fixture JSON's `data`, to derive a variant from.
    static func data(_ name: String) throws -> [String: Any] {
        let root = try JSONSerialization.jsonObject(with: RaidFixture.data(name)) as? [String: Any]
        guard let data = root?["data"] as? [String: Any] else { throw Failure(description: "\(name) has no data") }
        return data
    }

    static func decode<Value: Decodable>(_ type: Value.Type, _ object: Any) throws -> Value {
        try WireFormat.decoder().decode(type, from: JSONSerialization.data(withJSONObject: object))
    }

    /// Voltaris with a field replaced (`gear: []` for no export, a two-hander, and so on).
    static func voltaris(_ change: (inout [String: Any]) -> Void) throws -> CharacterProfile {
        var data = try data("character.card.ranked")
        change(&data)
        return try decode(CharacterProfile.self, data)
    }

    /// The board page with `deduped` false: the unit counts kills.
    static func kills() throws -> RankingsPage {
        var data = try data("rankings.list")
        data["deduped"] = false
        return try decode(RankingsPage.self, data)
    }

    /// The board page with a second class on it (a mixed board draws crests).
    static func mixed() throws -> RankingsPage {
        var data = try data("rankings.list")
        var rows = data["rows"] as? [[String: Any]] ?? []
        if var player = rows[1]["player"] as? [String: Any] {
            player["class"] = "Mage"
            player["spec"] = "Fire"
            rows[1]["player"] = player
        }
        data["rows"] = rows
        return try decode(RankingsPage.self, data)
    }
}
