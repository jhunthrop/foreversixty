import Foundation
import Models
import Networking
import Routing

/// The raid notification's categories and actions (app-v1.md "Push payloads", spec 6.3). The
/// server names the category; the app registers it once at launch with these action ids.
enum RaidPushCategory {
    /// In, Late and Out, all answered in the background.
    static let signup = "RAID_SIGNUP"
    /// `Choose a spec` (foreground, opens the spec sheet) and Out (background).
    static let askSpec = "RAID_ASK_SPEC"
    /// `Choose a character` (foreground, opens the raid page).
    static let choose = "RAID_CHOOSE"

    static let actionIn = "signup_in"
    static let actionLate = "signup_late"
    static let actionOut = "signup_out"
    static let actionAskSpec = "ask_spec"
    static let actionChoose = "choose"

    /// The sign-up state a background action answers, or nil for a foreground one.
    static func state(forAction id: String) -> SignupState? {
        switch id {
        case actionIn: .signedIn
        case actionLate: .late
        case actionOut: .out
        default: nil
        }
    }
}

/// A raid push's payload, read from the notification's userInfo (APNs: top-level `link` and
/// `data` strings).
struct RaidPushPayload: Equatable, Sendable {
    struct Answer: Codable, Equatable, Sendable {
        let title: String
        let body: String
    }

    let guild: Int64
    let raid: Int64
    let characterKey: String?
    let askSpec: Bool
    let link: URL?
    let answers: [String: Answer]

    init?(userInfo: [AnyHashable: Any]) {
        let data = userInfo["data"] as? [String: Any] ?? [:]
        func string(_ key: String) -> String? { data[key] as? String ?? userInfo[key] as? String }
        guard let guild = string("guild_id").flatMap(Int64.init), let raid = string("raid_id").flatMap(Int64.init) else {
            return nil
        }
        self.guild = guild
        self.raid = raid
        characterKey = string("character_key").flatMap { $0.isEmpty ? nil : $0 }
        askSpec = string("ask_spec") == "true"
        link = (userInfo["link"] as? String).flatMap(URL.init(string:))
        answers = string("answers")
            .flatMap { try? JSONDecoder().decode([String: Answer].self, from: Data($0.utf8)) } ?? [:]
    }

    /// The screen a tap opens: the payload's app link, else the raid page.
    func deepLink(askingSpec: Bool = false) -> DeepLink {
        if !askingSpec, let link, let parsed = DeepLink(url: link) { return parsed }
        return .raid(guild: guild, raid: raid, spec: askingSpec ? "ask" : nil, state: nil)
    }
}

/// What an action did: the answer to post in place of the notification, or a screen to open.
enum RaidPushOutcome: Equatable, Sendable {
    case answered(RaidPushPayload.Answer)
    case open(DeepLink)
}

/// Answers a background In, Late or Out: the PUT with the chosen character, keeping its note
/// (the PUT writes note too), then the server's answer text for the result.
struct RaidPushAnswerer: Sendable {
    let service: any RaidService

    func answer(_ payload: RaidPushPayload, state: SignupState) async -> RaidPushOutcome {
        let view: RaidView
        do {
            view = try await service.raid(guild: payload.guild, raid: payload.raid)
        } catch {
            return .answered(payload.answer("failed"))
        }
        guard let character = Self.character(for: payload, in: view) else { return .open(payload.deepLink()) }
        if state.carriesSpec, character.raidSpec == nil, (character.specOptions?.count ?? 0) >= 2 {
            return .open(payload.deepLink(askingSpec: true))
        }
        let note = view.signups.first { $0.characterKey == character.characterKey }?.note ?? ""
        do {
            let answer = try await service.signup(
                guild: payload.guild, raid: payload.raid, characterKey: character.characterKey,
                SignupRequest(state: state.rawValue, role: nil, note: note, spec: nil)
            )
            var reply = payload.answer(state.rawValue)
            if let moved = answer.moved, let suffix = payload.answers["moved"] {
                reply = .init(title: reply.title, body: reply.body + suffix.body.replacingOccurrences(of: "{name}", with: moved))
            }
            return .answered(reply)
        } catch {
            switch error.code {
            case "raid_locked": return .answered(payload.answer("locked"))
            case "raid_cancelled": return .answered(payload.answer("cancelled"))
            default: return .answered(payload.answer("failed"))
            }
        }
    }

    /// The server's pick when it sent one, else the same rule from the view: the character
    /// already signed, else the only verified one. Two or more unsigned: nil (open the page).
    static func character(for payload: RaidPushPayload, in view: RaidView) -> RaidViewerCharacter? {
        let verified = view.viewer.characters.filter(\.verified)
        if let key = payload.characterKey { return verified.first { $0.characterKey == key } }
        if let signed = verified.first(where: { $0.state != nil }) { return signed }
        return verified.count == 1 ? verified.first : nil
    }
}

private extension RaidPushPayload {
    /// The server's answer for an outcome; without one, the app's own words for saved or failed.
    func answer(_ key: String) -> Answer {
        if let answer = answers[key] { return answer }
        return key == "failed" ? Answer(title: Copy.commonFailed, body: "") : Answer(title: Copy.commonSaved, body: "")
    }
}
