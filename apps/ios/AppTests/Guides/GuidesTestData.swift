@testable import ForeverSixty
import Foundation
import Models
import Networking

/// F9's test data: the W3-C1 contract fixtures (bare static files, no envelope), read in place,
/// and the characters the boards draw. Every helper throws, so a missing file fails only its test.
enum GuidesTestData {
    /// The build the guide fixtures were written on: its icons are in data/builds.
    static let build = "1.60.1.70707"
    /// 2026-11-01 18:00 UTC: Name reservation is running, the beta is over, first raids are ahead.
    static let now = Date(timeIntervalSince1970: 1_793_556_000)

    static func file<Value: Decodable>(_ type: Value.Type, _ name: String) throws -> Value {
        try WireFormat.decoder().decode(type, from: HomeFixtures.data(name))
    }

    static func index() throws -> GuideIndex { try file(GuideIndex.self, "cdn.guides.index") }
    static func elemental() throws -> GuideFile { try file(GuideFile.self, "cdn.guides.shaman-elemental") }
    static func fire() throws -> GuideFile { try file(GuideFile.self, "cdn.guides.mage-fire") }
    static func protection() throws -> GuideFile { try file(GuideFile.self, "cdn.guides.warrior-protection") }
    static func tiers() throws -> TiersFile { try file(TiersFile.self, "cdn.tiers") }
    static func changelog() throws -> Changelog { try file(Changelog.self, "cdn.changelog") }
    static func dates() throws -> [KeyDate] { try file([KeyDate].self, "cdn.dates") }

    /// The boards' sample player: Lightningbolltz, an Orc Elemental Shaman at 60.
    static let lightningbolltz = character(name: "Lightningbolltz", className: "shaman", spec: "Elemental", race: "orc")
    /// The tier board's sample player: Emberquill, a Troll Fire Mage at 60.
    static let emberquill = character(name: "Emberquill", className: "mage", spec: "Fire", race: "troll")
    /// A spec that is first: Obnoxious Yell, the `me` fixture's Fury Warrior.
    static let obnoxiousYell = character(name: "Obnoxious Yell", className: "warrior", spec: "Fury", race: "troll")
    /// A Protection Warrior, for the tank list.
    static let tank = character(name: "Shieldwall", className: "warrior", spec: "Protection", race: "tauren")
    /// A Holy Priest, for the healer list.
    static let healer = character(name: "Mendwell", className: "priest", spec: "Holy", race: "troll")

    static func character(name: String, className: String, spec: String, race: String, faction: String = "horde") -> AccountCharacter {
        let slug = name.lowercased().replacingOccurrences(of: " ", with: "-")
        let json: [String: Any] = [
            "key": "us/normal/\(slug)", "region": "us", "ruleset": "normal", "name": name, "source": "addon",
            "class": className, "level": 60, "faction": faction, "race": race, "spec": spec,
        ]
        // The JSON is built here from literals; a decoding failure is a broken test helper.
        return try! WireFormat.decoder().decode(AccountCharacter.self, from: JSONSerialization.data(withJSONObject: json))
    }

    /// The saved guide files the boards draw: Elemental and Restoration Shaman, Protection Warrior.
    static func saved(_ specs: [String: Int] = ["shaman-elemental": 8563, "shaman-restoration": 9120, "warrior-protection": 13873]) -> SavedGuides {
        SavedGuides(bytes: specs)
    }
}

/// A `GuidesSource` over the fixtures, with switches for offline and for a newer build.
final class FixtureGuidesSource: GuidesSource, @unchecked Sendable {
    var build = GuidesTestData.build
    var latest: String?
    var offline = false
    var failing = false
    var cachedSpecs: Set<String> = []

    private func answer<Value>(_ make: () throws -> Value) throws(APIError) -> CachedFile<Value> {
        if offline { throw .transport("offline") }
        if failing { throw .unexpectedStatus(500) }
        do {
            return CachedFile(value: try make(), savedAt: GuidesTestData.now, isFromCache: false)
        } catch {
            throw .missingFixture("\(error)")
        }
    }

    func activeBuild() async -> String { build }
    func latestBuild() async -> String? { offline ? nil : (latest ?? build) }

    func index(build: String) async throws(APIError) -> CachedFile<GuideIndex> {
        try answer { try GuidesTestData.index() }
    }

    func cachedIndex(build: String) -> CachedFile<GuideIndex>? {
        (try? GuidesTestData.index()).map { CachedFile(value: $0, savedAt: GuidesTestData.now, isFromCache: true) }
    }

    func guide(build: String, spec: String) async throws(APIError) -> CachedFile<GuideFile> {
        if offline, let saved = cachedGuide(build: build, spec: spec) { return saved }
        return try answer { try GuidesTestData.file(GuideFile.self, "cdn.guides.\(spec)") }
    }

    func cachedGuide(build: String, spec: String) -> CachedFile<GuideFile>? {
        guard cachedSpecs.contains(spec) else { return nil }
        return (try? GuidesTestData.file(GuideFile.self, "cdn.guides.\(spec)"))
            .map { CachedFile(value: $0, savedAt: GuidesTestData.now, isFromCache: true) }
    }

    func savedBytes(build: String, specs: [String]) -> SavedGuides {
        SavedGuides(bytes: Dictionary(uniqueKeysWithValues: specs.filter(cachedSpecs.contains).map { ($0, 1000) }))
    }

    func tiers(build: String) async throws(APIError) -> CachedFile<TiersFile> {
        try answer { try GuidesTestData.tiers() }
    }

    func cachedTiers(build: String) -> CachedFile<TiersFile>? { nil }

    func changelog() async throws(APIError) -> CachedFile<Changelog> {
        try answer { try GuidesTestData.changelog() }
    }

    func cachedChangelog() -> CachedFile<Changelog>? { nil }

    func keyDates() async -> [KeyDate] { (try? GuidesTestData.dates()) ?? [] }
}
