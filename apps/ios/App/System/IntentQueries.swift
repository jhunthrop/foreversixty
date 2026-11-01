import Foundation
import Models
import Networking

/// The Next upgrade answer: the snippet's words, or a sentence when there is nothing to show.
enum UpgradeReply: Equatable, Sendable {
    case upgrade(UpgradeAnswer, build: String)
    case sentence(String)

    var sentence: String {
        switch self {
        case let .upgrade(answer, _): answer.sentence
        case let .sentence(text): text
        }
    }
}

/// The read-only intents' answers (spec §5.7): Next raid, Who is coming and Next upgrade, from
/// the widget snapshot, the next raid's view and the Home answer. Every refusal is a sentence.
struct IntentQueries: Sendable {
    let calls: any IntentCalls
    let answers: IntentAnswers
    let enabled: @Sendable () -> Bool

    /// `Onyxia's Lair, Thursday at 8 pm. You are In. 25 of 40 signed up.`
    func nextRaid(character: String?) async -> String {
        await answer { () throws(APIError) -> String in
            let widget = try await calls.widget(character: character)
            if let refusal = Self.refusal(widget) { return refusal }
            guard let raid = widget.nextRaid else { return Copy.intentNoRaid }
            return answers.nextRaid(raid)
        }
    }

    /// `24 of 40 signed up: 3 tanks, 6 healers, 14 DPS, 1 late.`
    func whoIsComing(character: String?) async -> String {
        await answer { () throws(APIError) -> String in
            let widget = try await calls.widget(character: character)
            if let refusal = Self.refusal(widget) { return refusal }
            guard let next = widget.nextRaid else { return Copy.intentNoRaid }
            return answers.whoIsComing(try await calls.raid(guild: next.guildId, raid: next.raidId))
        }
    }

    func nextUpgrade(character: String?) async -> UpgradeReply {
        guard enabled() else { return .sentence(Copy.intentOff) }
        do throws(APIError) {
            let widget = try await calls.widget(character: character)
            guard let you = widget.character else { return .sentence(Copy.intentNoCharacter) }
            let home = try await calls.home(characterKey: you.key)
            guard let upgrade = UpgradeAnswer(home: home, characterName: you.name) else {
                return .sentence(Copy.intentNoUpgrade(name: you.name))
            }
            return .upgrade(upgrade, build: home.dataBuild)
        } catch {
            return .sentence(Self.failure(error))
        }
    }

    private func answer(_ body: () async throws(APIError) -> String) async -> String {
        guard enabled() else { return Copy.intentOff }
        do throws(APIError) {
            return try await body()
        } catch {
            return Self.failure(error)
        }
    }

    private static func refusal(_ widget: Widget) -> String? {
        guard let you = widget.character else { return Copy.intentNoCharacter }
        return widget.guild == nil ? Copy.intentNoGuild(name: you.name) : nil
    }

    private static func failure(_ error: APIError) -> String {
        switch error {
        case .unauthorized: Copy.intentSignedOut
        case .transport: Copy.intentOffline
        default: Copy.commonFailed
        }
    }
}
