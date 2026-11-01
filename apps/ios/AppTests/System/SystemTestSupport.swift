@testable import ForeverSixty
import Foundation
import Models
import Networking

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

    /// guild.raid, re-cast as the story's raid: Lightningbolltz's view, the given counts.
    static func raid(
        id: Int64 = 7, startsAt: String = thursday, counts: [String: Int] = ["in": 24, "late": 1, "tanks": 3, "healers": 6, "dps": 15],
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

    static func widget() throws -> Widget { try WidgetFixtures.widget("widget.get") }

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
            let counts = request.state == "in" ? ["in": 24, "late": 1] : ["in": 23, "late": 2]
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
