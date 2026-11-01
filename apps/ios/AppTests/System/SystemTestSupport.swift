@testable import ForeverSixty
import Foundation
import Models
import Networking
import Testing

/// The spec's one story (§5.7): Lightningbolltz, Elemental Shaman on OLYMPUS XXVII, In for
/// Onyxia's Lair on Thursday 10 December at 8 pm Chicago time, 25 of 40 signed up.
enum SystemStory {
    static let zone = TimeZone(identifier: "America/Chicago") ?? .gmt
    static let answers = IntentAnswers(clock: WidgetClock(zone: zone))
    /// Tuesday 8 December, 5:39 pm in Chicago.
    static let now = RaidTime.date("2026-12-08T23:39:00Z")
    /// Thursday 10 December, 8 pm in Chicago.
    static let thursday = "2026-12-11T02:00:00Z"
    static let key = "us/pvp/lightningbolltz"
    /// 25 of 40 with you In (lead ruling): 24 In (3 tanks, 6 healers, 15 DPS) and 1 Late.
    static let counts = ["in": 24, "late": 1, "tanks": 3, "healers": 6, "dps": 15]

    /// guild.raid, re-cast as the story's raid: Lightningbolltz's view, the given counts.
    static func raid(
        id: Int64 = 7, startsAt: String = thursday, counts: [String: Int] = SystemStory.counts,
        status: String = "planned", raidSpec: String? = "shaman-elemental", options: Int = 3
    ) throws -> RaidView {
        var root = try JSONSerialization.jsonObject(with: RaidFixture.data("guild.raid")) as? [String: Any] ?? [:]
        var data = root["data"] as? [String: Any] ?? [:]
        data["id"] = id
        data["starts_at"] = startsAt
        data["status"] = status
        var merged = data["counts"] as? [String: Any] ?? [:]
        for (key, value) in counts { merged[key] = value }
        data["counts"] = merged
        let specs: [[String: Any]] = [
            ["spec": "shaman-elemental", "name": "Elemental", "role": "dps", "icon": "spell_nature_lightning"],
            ["spec": "shaman-enhancement", "name": "Enhancement", "role": "dps", "icon": "spell_nature_lightningshield"],
            ["spec": "shaman-restoration", "name": "Restoration", "role": "healer", "icon": "spell_nature_magicimmunity"],
        ]
        var viewer = data["viewer"] as? [String: Any] ?? [:]
        viewer["characters"] = [[
            "character_key": key, "name": "Lightningbolltz", "class": "shaman", "spec": "shaman-elemental",
            "verified": true, "state": "in", "role": "dps", "raid_spec": raidSpec as Any? ?? NSNull(),
            "spec_name": "Elemental", "spec_icon": "spell_nature_lightning", "spec_options": Array(specs.prefix(options)),
        ]]
        data["viewer"] = viewer
        root["data"] = data
        return try WireFormat.decoder().decode(Envelope<RaidView>.self, from: JSONSerialization.data(withJSONObject: root)).data
    }

    /// widget.get with the story's count: 25 of 40, you In.
    static func widget() throws -> Widget { try WidgetFixtures.widget("widget.get", patch: ["signed_up": .number(25)]) }

    /// character.home.yours, its first upgrade's source as the list's own sentence.
    static func home() throws -> CharacterHome {
        var root = try JSONSerialization.jsonObject(with: HomeFixtures.data("character.home.yours")) as? [String: Any] ?? [:]
        var data = root["data"] as? [String: Any] ?? [:]
        var upgrades = data["upgrades"] as? [[String: Any]] ?? []
        if var first = upgrades.first, var pick = first["pick"] as? [String: Any] {
            pick["source_line"] = Copy.bisSourceCrafted(name: "Leatherworking")
            first["pick"] = pick
            upgrades[0] = first
        }
        data["upgrades"] = upgrades
        root["data"] = data
        return try WireFormat.decoder().decode(Envelope<CharacterHome>.self, from: JSONSerialization.data(withJSONObject: root)).data
    }

    /// Your characters as the API names them (lead ruling): Lightningbolltz first, then Obnoxiousyell.
    static func ownCharacters() throws -> [AccountCharacter] {
        let me = try RaidFixture.load(Me.self, "me")
        let base = try JSONSerialization.jsonObject(with: WireFormat.encoder().encode(try #require(me.characters.first))) as? [String: Any] ?? [:]
        func make(_ key: String, _ name: String, _ className: String, _ spec: String) throws -> AccountCharacter {
            var json = base
            json["key"] = key
            json["name"] = name
            json["class"] = className
            json["spec"] = spec
            json["level"] = 60
            return try WireFormat.decoder().decode(AccountCharacter.self, from: JSONSerialization.data(withJSONObject: json))
        }
        return [
            try make(key, "Lightningbolltz", "shaman", "Elemental"),
            try make("us/pvp/obnoxiousyell", "Obnoxiousyell", "warrior", "Fury"),
        ]
    }

    /// Lightningbolltz with a talent export of `order` less its last `dropping` points.
    static func exporter(dropping: Int, file: TalentsFile, order: [Int]) throws -> AccountCharacter {
        let index = PlannerIndex(file)
        let ranks = PlannerRules.treeRanks(index, ranks: PlannerRules.ranks(Array(order.dropLast(dropping))))
        let me = try RaidFixture.load(Me.self, "me")
        var json = try JSONSerialization.jsonObject(with: WireFormat.encoder().encode(try #require(me.characters.first))) as? [String: Any] ?? [:]
        var build = json["build"] as? [String: Any] ?? [:]
        build["talents"] = ["trees": ranks.map { $0.map(String.init).joined() }, "points": ranks.map { $0.reduce(0, +) }]
        json["build"] = build
        json["key"] = key
        json["name"] = "Lightningbolltz"
        json["class"] = "shaman"
        json["spec"] = "Elemental"
        json["level"] = 60
        return try WireFormat.decoder().decode(AccountCharacter.self, from: JSONSerialization.data(withJSONObject: json))
    }

    /// The shared "saved on this phone" figures both platforms draw (lead ruling).
    static func offlineReport(savedAt: Date) -> OfflineReport {
        let build = "1.60.1.70707"
        return OfflineReport.build(files: [
            SavedStaticFile(relativePath: "\(build)/guides/shaman-elemental.json", bytes: 1_300_000, savedAt: savedAt),
            SavedStaticFile(relativePath: "\(build)/guides/warrior-fury.json", bytes: 1_100_000, savedAt: savedAt),
            SavedStaticFile(relativePath: "\(build)/bis/shaman-elemental.json", bytes: 3_100_000, savedAt: savedAt),
            SavedStaticFile(relativePath: "\(build)/talents/shaman.json", bytes: 600_000, savedAt: savedAt),
            SavedStaticFile(relativePath: "\(build)/search.json", bytes: 400_000, savedAt: savedAt),
        ], iconBytes: 41_000_000, iconsUpdated: savedAt)
    }

    private struct Envelope<Value: Decodable>: Decodable { let data: Value }
}

/// Scripted intent calls: each answers from its closure; every sign-up is recorded.
final class FakeIntentCalls: IntentCalls, @unchecked Sendable {
    var widgetAnswer: () throws(APIError) -> Widget = { () throws(APIError) in
        do { return try SystemStory.widget() } catch { throw .missingFixture("\(error)") }
    }
    var planned: [RaidView] = []
    var raidAnswer: () throws(APIError) -> RaidView = { () throws(APIError) in
        do { return try SystemStory.raid() } catch { throw .missingFixture("\(error)") }
    }
    var signupAnswer: (SignupRequest) throws(APIError) -> SignupAnswer = { request throws(APIError) in
        do {
            let counts = request.state == "in" ? SystemStory.counts : ["in": 23, "late": 2]
            return SignupAnswer(raid: try SystemStory.raid(counts: counts))
        } catch { throw .missingFixture("\(error)") }
    }
    var homeAnswer: () throws(APIError) -> CharacterHome = { () throws(APIError) in
        do { return try HomeFixtures.home("character.home.yours") } catch { throw .missingFixture("\(error)") }
    }
    private(set) var signups: [(raid: Int64, request: SignupRequest)] = []

    func widget(character: String?) async throws(APIError) -> Widget { try widgetAnswer() }

    func raids(guild: Int64) async throws(APIError) -> GuildRaids { GuildRaids(planned: planned) }

    func raid(guild: Int64, raid: Int64) async throws(APIError) -> RaidView { try raidAnswer() }

    func signup(guild: Int64, raid: Int64, characterKey: String, _ request: SignupRequest) async throws(APIError) -> SignupAnswer {
        signups.append((raid, request))
        return try signupAnswer(request)
    }

    func home(characterKey: String) async throws(APIError) -> CharacterHome { try homeAnswer() }
}

/// Answers Siri's questions from a script and records what was asked.
final class ScriptedAsker: IntentAsker, @unchecked Sendable {
    var choice: Int? = 0
    var confirms = true
    private(set) var questions: [String] = []
    private(set) var options: [[String]] = []

    func choose(_ question: String, options: [String]) async throws -> Int? {
        questions.append(question)
        self.options.append(options)
        return choice
    }

    func confirm(_ question: String) async throws -> Bool {
        questions.append(question)
        return confirms
    }
}
