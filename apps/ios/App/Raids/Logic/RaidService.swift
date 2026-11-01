import Foundation
import Models
import Networking

/// The raid calls the screens make. Live goes through `APIClient`; tests pass a fake.
protocol RaidService: Sendable {
    func me() async throws(APIError) -> Me
    func raids(guild: Int64) async throws(APIError) -> GuildRaids
    func raid(guild: Int64, raid: Int64) async throws(APIError) -> RaidView
    func signup(guild: Int64, raid: Int64, characterKey: String, _ request: SignupRequest) async throws(APIError) -> SignupAnswer
    func lock(guild: Int64, raid: Int64) async throws(APIError) -> RaidView
    func unlock(guild: Int64, raid: Int64) async throws(APIError) -> RaidView
    func night(raid: Int64) async throws(APIError) -> RaidNightSummary
}

struct LiveRaidService: RaidService {
    let client: APIClient

    func me() async throws(APIError) -> Me { try await client.send(.getMe) }

    func raids(guild: Int64) async throws(APIError) -> GuildRaids {
        try await client.send(.listGuildRaids(id: guild))
    }

    func raid(guild: Int64, raid: Int64) async throws(APIError) -> RaidView {
        try await client.send(.getGuildRaid(id: guild, raid: raid))
    }

    func signup(guild: Int64, raid: Int64, characterKey: String, _ request: SignupRequest) async throws(APIError) -> SignupAnswer {
        try await RaidSignupPUT.send(client, guild: guild, raid: raid, characterKey: characterKey, request)
    }

    func lock(guild: Int64, raid: Int64) async throws(APIError) -> RaidView {
        try await client.send(.lockGuildRaid(id: guild, raid: raid))
    }

    func unlock(guild: Int64, raid: Int64) async throws(APIError) -> RaidView {
        try await client.send(.unlockGuildRaid(id: guild, raid: raid))
    }

    func night(raid: Int64) async throws(APIError) -> RaidNightSummary {
        try await client.send(.getRaidNight(raid: raid))
    }
}

/// A character key, `region/ruleset/name-slug`, split for the sign-up route.
struct CharacterKey: Equatable, Sendable {
    let region: String
    let ruleset: String
    let name: String

    init(_ key: String) throws(APIError) {
        let parts = key.split(separator: "/", omittingEmptySubsequences: false).map(String.init)
        guard parts.count == 3, parts.allSatisfy({ !$0.isEmpty }) else {
            throw .decoding("not a character key: \(key)")
        }
        (region, ruleset, name) = (parts[0], parts[1], parts[2])
    }
}

/// The sentence a refused or failed write shows (spec 5.4 raid states and section 13 `raid.*`).
enum SignupRefusal {
    static func sentence(for error: APIError) -> String {
        switch error {
        case .transport: return Copy.raidOffline
        case let .server(_, body, _):
            switch body.code {
            case "raid_locked": return Copy.raidRaceLocked
            case "raid_cancelled": return Copy.raidRaceCancelled
            case "guild_plan_required": return Copy.raidRaceToolsEnded
            case "rate_limited": return Copy.raidRateLimited
            case "forbidden": return body.message
            default: return Copy.commonFailed
            }
        default: return Copy.commonFailed
        }
    }

    /// True when the failure was the connection, so the screen shows its offline state.
    static func isOffline(_ error: APIError) -> Bool {
        if case .transport = error { return true }
        return false
    }
}
