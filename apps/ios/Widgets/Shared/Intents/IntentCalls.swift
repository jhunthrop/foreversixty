import Auth
import Foundation
import Models
import Networking

/// The one sign-up PUT every surface sends: the raid page and card, the widget, the Live
/// Activity, Siri, Shortcuts and the Control (spec §5.7 `The same rules as the app`).
enum RaidSignupPUT {
    static func send(
        _ client: APIClient, guild: Int64, raid: Int64, characterKey: String, _ request: SignupRequest
    ) async throws(APIError) -> SignupAnswer {
        let parts = characterKey.split(separator: "/", omittingEmptySubsequences: false).map(String.init)
        guard parts.count == 3, parts.allSatisfy({ !$0.isEmpty }) else {
            throw .decoding("not a character key: \(characterKey)")
        }
        return try await client.send(.putRaidSignup(
            id: guild, raid: raid, region: parts[0], ruleset: parts[1], name: parts[2], request
        ))
    }
}

/// What the intents read and write: the widget snapshot (character, guild, next raid), the
/// guild's planned raids, one raid's view, the sign-up and the Home answer. No new endpoint.
protocol IntentCalls: Sendable {
    func widget(character: String?) async throws(APIError) -> Widget
    func raids(guild: Int64) async throws(APIError) -> GuildRaids
    func raid(guild: Int64, raid: Int64) async throws(APIError) -> RaidView
    func signup(guild: Int64, raid: Int64, characterKey: String, _ request: SignupRequest) async throws(APIError) -> SignupAnswer
    func home(characterKey: String) async throws(APIError) -> CharacterHome
}

struct LiveIntentCalls: IntentCalls {
    let client: APIClient

    func widget(character: String?) async throws(APIError) -> Widget {
        try await client.send(.getWidget(character: character))
    }

    func raids(guild: Int64) async throws(APIError) -> GuildRaids {
        try await client.send(.listGuildRaids(id: guild))
    }

    func raid(guild: Int64, raid: Int64) async throws(APIError) -> RaidView {
        try await client.send(.getGuildRaid(id: guild, raid: raid))
    }

    func signup(guild: Int64, raid: Int64, characterKey: String, _ request: SignupRequest) async throws(APIError) -> SignupAnswer {
        try await RaidSignupPUT.send(client, guild: guild, raid: raid, characterKey: characterKey, request)
    }

    func home(characterKey: String) async throws(APIError) -> CharacterHome {
        let parts = characterKey.split(separator: "/").map(String.init)
        guard parts.count == 3 else { throw .decoding("not a character key: \(characterKey)") }
        return try await client.send(.getCharacterHome(region: parts[0], ruleset: parts[1], name: parts[2]))
    }
}

extension LiveIntentCalls {
    /// The process's own client, as the widget's intents build it (the App Group keychain token).
    static func fromBundle(_ bundle: Bundle = .main) -> LiveIntentCalls? {
        LiveWidgetCalls.fromBundle(bundle).map { LiveIntentCalls(client: $0.client) }
    }
}

/// The `system` switch as the app last read it, in the App Group so the widget extension's
/// Control reads it too (spec §5.7: off, an intent answers `intent.off`).
enum SystemSwitch {
    static let key = "system.enabled"

    static func isOn(_ defaults: UserDefaults = AppGroup.defaults) -> Bool {
        defaults.bool(forKey: key)
    }

    static func record(_ on: Bool, _ defaults: UserDefaults = AppGroup.defaults) {
        defaults.set(on, forKey: key)
    }
}
