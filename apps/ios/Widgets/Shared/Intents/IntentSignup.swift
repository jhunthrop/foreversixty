import Foundation
import Models
import Networking

/// Which raid an intent means (spec §5.7 `Which raid`): the next raid starting within 24 hours;
/// else the next planned one; two or more within 24 hours ask `Which raid?`.
enum RaidPick: Equatable, Sendable {
    case none
    case one(RaidView)
    case ask([RaidView])

    static let window: TimeInterval = 24 * 3600

    static func pick(_ planned: [RaidView], now: Date) -> RaidPick {
        let open = planned
            .filter { $0.status != "cancelled" && end(of: $0) > now }
            .sorted { $0.startsAt < $1.startsAt }
        let soon = open.filter { $0.startsAt.timeIntervalSince(now) <= window }
        if soon.count >= 2 { return .ask(soon) }
        return open.first.map(RaidPick.one) ?? .none
    }

    private static func end(of raid: RaidView) -> Date {
        raid.startsAt.addingTimeInterval(TimeInterval(raid.durationMinutes * 60))
    }
}

/// The questions a sign-up may ask, answered by Siri's dialogs in the intent and by a script in
/// tests. A nil answer is Cancel.
protocol IntentAsker: Sendable {
    func choose(_ question: String, options: [String]) async throws -> Int?
    func confirm(_ question: String) async throws -> Bool
}

/// The sign-up's outcome: the sentence to say, and whether it saved.
struct IntentReply: Equatable, Sendable {
    let sentence: String
    let saved: Bool

    static func refused(_ sentence: String) -> IntentReply { IntentReply(sentence: sentence, saved: false) }
    /// Cancel pressed on a question: nothing is said and nothing saved.
    static let cancelled = IntentReply(sentence: "", saved: false)
}

/// Sign up for a raid from Siri, Shortcuts, the Action button or the Control (spec §5.7). Asks
/// which raid and which spec when the app would, asks once to confirm only when `askFirst` (run
/// by voice), then sends the app's own PUT and answers with the server's counts.
struct IntentSignup: Sendable {
    let calls: any IntentCalls
    let answers: IntentAnswers
    let now: @Sendable () -> Date
    let enabled: @Sendable () -> Bool

    func run(state: String, character: String?, askFirst: Bool, asker: any IntentAsker) async throws -> IntentReply {
        guard enabled() else { return .refused(Copy.intentOff) }
        do {
            return try await signUp(state: state, character: character, askFirst: askFirst, asker: asker)
        } catch let error as APIError {
            if case .unauthorized = error { return .refused(Copy.intentSignedOut) }
            return .refused(IntentAnswers.refusal(error))
        }
    }

    private func signUp(state: String, character: String?, askFirst: Bool, asker: any IntentAsker) async throws -> IntentReply {
        let widget = try await calls.widget(character: character)
        guard let you = widget.character else { return .refused(Copy.intentNoCharacter) }
        guard let guild = widget.guild else { return .refused(Copy.intentNoGuild(name: you.name)) }
        let pick = RaidPick.pick(try await calls.raids(guild: guild.id).planned ?? [], now: now())
        guard pick != .none else { return .refused(Copy.intentNoRaid) }
        guard let raid = try await chooseRaid(pick, asker: asker) else { return .cancelled }
        let viewer = raid.viewer.characters.first { $0.characterKey == you.key }
        guard let spec = try await chooseSpec(viewer, state: state, guild: guild.name, asker: asker) else { return .cancelled }
        let question = answers.confirm(name: you.name, state: state, zone: raid.zoneName, startsAt: raid.startsAt)
        if askFirst, try await !asker.confirm(question) { return .cancelled }
        let request = SignupRequest(state: state, role: nil, note: noteOf(raid, key: you.key), spec: spec.value)
        let answer = try await calls.signup(guild: guild.id, raid: raid.id, characterKey: you.key, request)
        let role = answer.raid.viewer.characters.first { $0.characterKey == you.key }?.role ?? viewer?.role ?? "dps"
        return IntentReply(sentence: answers.signedUp(state: state, raid: answer.raid, characterName: you.name, role: role), saved: true)
    }

    private func chooseRaid(_ pick: RaidPick, asker: any IntentAsker) async throws -> RaidView? {
        switch pick {
        case .none: return nil
        case let .one(raid): return raid
        case let .ask(raids):
            let titles = raids.map { answers.raidTitle(zone: $0.zoneName, startsAt: $0.startsAt) }
            return try await asker.choose(Copy.intentWhich, options: titles).map { raids[$0] }
        }
    }

    /// The spec the PUT carries: none needed (`.some(nil)`), a chosen one, or nil on Cancel. Asked
    /// as the raid page asks: In or Late, no raid spec yet and two or more specs to choose from.
    private func chooseSpec(_ viewer: RaidViewerCharacter?, state: String, guild: String, asker: any IntentAsker) async throws -> SpecChoice? {
        guard let viewer, state != "out", viewer.raidSpec == nil, let options = viewer.specOptions, options.count >= 2 else {
            return SpecChoice(value: nil)
        }
        let question = Copy.intentSpec(name: viewer.name, guild: guild)
        return try await asker.choose(question, options: options.map(\.name)).map { SpecChoice(value: options[$0].spec) }
    }

    private func noteOf(_ raid: RaidView, key: String) -> String {
        raid.signups.first { $0.characterKey == key }?.note ?? ""
    }

    private struct SpecChoice { let value: String? }
}
