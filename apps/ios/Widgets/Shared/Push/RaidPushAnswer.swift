import Foundation
import Models
import Networking
import Routing
import UserNotifications

/// The raid notification's category and actions (flows spec F1). The server names the category;
/// the app registers it at launch, and the content extension swaps Late for its minutes.
enum RaidPushCategory {
    /// In, Late, Bench and Out, all answered without opening the app.
    static let signup = "RAID_SIGNUP"
    /// `Choose a spec` (foreground, opens the spec sheet) and Out (background).
    static let askSpec = "RAID_ASK_SPEC"
    /// `Choose a character` (foreground, opens the raid page).
    static let choose = "RAID_CHOOSE"
    /// An officer's new applicant: Accept and Decline, answered without opening the app.
    static let applicant = "GUILD_APPLICANT"
    /// An officer's short raid (`raid.short`): Mention the short role in Discord.
    static let short = "RAID_SHORT"

    static let actionIn = "signup_in"
    static let actionLate = "signup_late"
    static let actionBench = "signup_bench"
    static let actionOut = "signup_out"
    static let actionAskSpec = "ask_spec"
    static let actionChoose = "choose"
    /// The API's action id on `guild.applicant` (POST .../applications/{app}/decline).
    static let actionDecline = "application_decline"
    /// The API's action id on `guild.applicant` (POST .../applications/{app}/accept).
    static let actionAccept = "application_accept"
    /// The API's action id on `raid.short` (POST .../raids/{raid}/ping with the short role).
    static let actionMention = "raid_mention_short"

    /// The Late chips (15, 30, 60 min) the content extension offers after Late.
    static let lateMinutes = [15, 30, 60]

    static func actionLate(minutes: Int) -> String { "\(actionLate)_\(minutes)" }

    /// The sign-up an action answers, or nil for a foreground one.
    static func signup(forAction id: String) -> PushSignup? {
        switch id {
        case actionIn: PushSignup(state: "in")
        case actionLate: PushSignup(state: "late")
        case actionBench: PushSignup(state: "bench")
        case actionOut: PushSignup(state: "out")
        default:
            lateMinutes.first { actionLate(minutes: $0) == id }.map { PushSignup(state: "late", lateMinutes: $0) }
        }
    }
}

/// One answer from a notification: the state, and Late's minutes when a chip chose them.
struct PushSignup: Equatable, Sendable {
    let state: String
    var lateMinutes: Int?

    /// Out never asks for a spec; every other answer places the character in a role.
    var carriesSpec: Bool { state != "out" }
}

/// A raid push's payload, read from the notification's userInfo (APNs: top-level `kind`, `group`
/// and `link`, the `data` strings).
struct RaidPushPayload: Equatable, Sendable {
    let guild: Int64
    let raid: Int64
    let characterKey: String?
    let askSpec: Bool
    let link: URL?

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
    }

    /// The screen a tap opens: the payload's app link, else the raid page.
    func deepLink(askingSpec: Bool = false) -> DeepLink {
        if !askingSpec, let link, let parsed = DeepLink(url: link) { return parsed }
        return .raid(guild: guild, raid: raid, spec: askingSpec ? "ask" : nil, state: nil)
    }
}

/// What answering did: the confirmation to post in place of the notification, or a screen to open.
enum RaidPushOutcome: Equatable, Sendable {
    struct Confirmation: Equatable, Sendable {
        let title: String
        let body: String
    }

    case answered(Confirmation)
    case open(DeepLink)
}

/// The two calls an answer needs. The app's RaidService and the extensions' intent calls both
/// have them.
protocol PushSignupCalls: Sendable {
    func raid(guild: Int64, raid: Int64) async throws(APIError) -> RaidView
    func signup(guild: Int64, raid: Int64, characterKey: String, _ request: SignupRequest) async throws(APIError) -> SignupAnswer
}

/// Answers from the lock screen: the PUT with the chosen character, keeping its note (the PUT
/// writes note too), then the server's confirmation (`SignupAnswer.confirmation`).
struct RaidPushAnswerer: Sendable {
    let calls: any PushSignupCalls

    func answer(_ payload: RaidPushPayload, _ signup: PushSignup) async -> RaidPushOutcome {
        let view: RaidView
        do {
            view = try await calls.raid(guild: payload.guild, raid: payload.raid)
        } catch {
            return .answered(Self.failed)
        }
        guard let character = Self.character(for: payload, in: view) else { return .open(payload.deepLink()) }
        if signup.carriesSpec, character.raidSpec == nil, (character.specOptions?.count ?? 0) >= 2 {
            return .open(payload.deepLink(askingSpec: true))
        }
        let note = view.signups.first { $0.characterKey == character.characterKey }?.note ?? ""
        let request = SignupRequest(state: signup.state, role: nil, note: note, spec: nil, lateMinutes: signup.lateMinutes)
        do {
            let answer = try await calls.signup(guild: payload.guild, raid: payload.raid, characterKey: character.characterKey, request)
            return .answered(answer.confirmation.map { .init(title: $0.title, body: $0.body) } ?? .init(title: Copy.commonSaved, body: ""))
        } catch {
            switch error.code {
            case "raid_locked": return .answered(.init(title: Copy.raidLocked, body: ""))
            case "raid_cancelled": return .answered(.init(title: Copy.raidCancelled, body: ""))
            default: return .answered(Self.failed)
            }
        }
    }

    private static var failed: RaidPushOutcome.Confirmation { .init(title: Copy.commonFailed, body: "") }

    /// The server's pick when it sent one, else the same rule from the view: the character
    /// already signed, else the only verified one. Two or more unsigned: nil (open the page).
    static func character(for payload: RaidPushPayload, in view: RaidView) -> RaidViewerCharacter? {
        let verified = view.viewer.characters.filter(\.verified)
        if let key = payload.characterKey { return verified.first { $0.characterKey == key } }
        if let signed = verified.first(where: { $0.state != nil }) { return signed }
        return verified.count == 1 ? verified.first : nil
    }
}

extension LiveIntentCalls: PushSignupCalls {}

extension RaidPushOutcome.Confirmation {
    /// Posts the confirmation in place of the question: the same identifier and thread, no
    /// actions (flows spec F1: a quiet confirmation naming the character). `tapThrough` keeps where a
    /// tap on it lands.
    func post(replacing identifier: String, thread: String, tapThrough: [String: PushTapValue] = [:]) async {
        let content = UNMutableNotificationContent()
        content.title = title
        content.body = body
        content.threadIdentifier = thread
        content.userInfo = tapThrough.mapValues(\.value)
        let request = UNNotificationRequest(identifier: identifier, content: content, trigger: nil)
        try? await UNUserNotificationCenter.current().add(request)
    }
}

/// One tap-through value of a push's userInfo: a string (`kind`, `link`) or the `data` strings.
enum PushTapValue: Equatable, Sendable {
    case text(String)
    case strings([String: String])

    static let keys = ["kind", "group", "link", "data"]

    static func tapThrough(_ userInfo: [AnyHashable: Any]) -> [String: PushTapValue] {
        keys.reduce(into: [:]) { kept, key in
            if let text = userInfo[key] as? String {
                kept[key] = .text(text)
            } else if let strings = userInfo[key] as? [String: Any] {
                kept[key] = .strings(strings.compactMapValues { $0 as? String })
            }
        }
    }

    var value: Any {
        switch self {
        case let .text(text): text
        case let .strings(strings): strings
        }
    }
}

/// UIKit's completion for a notification response, called back on the main thread only.
struct MainThreadCompletion: @unchecked Sendable {
    private let completion: () -> Void

    init(_ completion: @escaping () -> Void) {
        self.completion = completion
    }

    @MainActor
    func run() {
        completion()
    }
}
