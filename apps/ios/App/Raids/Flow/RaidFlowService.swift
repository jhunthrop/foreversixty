import Foundation
import Models
import Networking

/// The calls the Raid tab's flows make beyond the raid list (flows spec 4.2, F2, F3), behind a
/// protocol for tests.
protocol RaidFlowService: Sendable {
    /// `GET /v1/guilds/{id}/raid-badge`: an officer's open items (the Needs you count).
    func badge(guild: Int64) async throws(APIError) -> RaidBadge
    /// `GET /v1/guilds/{id}/applications`.
    func applications(guild: Int64) async throws(APIError) -> GuildOfficerApplications
    func accept(guild: Int64, application: Int64) async throws(APIError) -> GuildApplicationDecision
    func decline(guild: Int64, application: Int64) async throws(APIError) -> GuildApplicationDecision
    /// `GET /v1/guilds/recruiting`: the guilds recruiting a spec on a faction.
    func recruiting(spec: String, faction: String, region: String, ruleset: String) async throws(APIError) -> RecruitingGuilds
    /// The Ready check for one raid; nil without an export or on any failure.
    func ready(character: String, raid: Int64) async -> ReadyCheck?
}

/// What the app cannot do itself yet, so the site does it: the role-scoped Discord mention, a
/// direct message, applying to a guild and claiming one (the report names the API they need).
protocol RaidSiteActions: Sendable {
    /// Where Mention goes: the raid's page on the site, which posts the ping.
    func mention(guild: Int64, raid: Int64, role: String?) -> URL
    /// A Discord direct message to a raider.
    func message(characterKey: String) -> URL
    /// Applying to a recruiting guild.
    func apply(guild: Int64) -> URL
    /// Claiming a guild that is not here, and finding one.
    var guilds: URL { get }
}

struct LiveRaidFlowService: RaidFlowService {
    let client: APIClient

    func badge(guild: Int64) async throws(APIError) -> RaidBadge {
        try await client.send(.getGuildRaidBadge(id: guild))
    }

    func applications(guild: Int64) async throws(APIError) -> GuildOfficerApplications {
        try await client.send(.listGuildApplications(id: guild))
    }

    func accept(guild: Int64, application: Int64) async throws(APIError) -> GuildApplicationDecision {
        try await client.send(.acceptGuildApplication(id: guild, app: application))
    }

    func decline(guild: Int64, application: Int64) async throws(APIError) -> GuildApplicationDecision {
        try await client.send(.declineGuildApplication(id: guild, app: application))
    }

    func recruiting(spec: String, faction: String, region: String, ruleset: String) async throws(APIError) -> RecruitingGuilds {
        try await client.send(.listRecruitingGuilds(spec: spec, faction: faction, region: region, ruleset: ruleset))
    }

    func ready(character: String, raid: Int64) async -> ReadyCheck? {
        await LiveReadySource(client: client).ready(character: character, raid: raid)
    }
}

struct LiveRaidSiteActions: RaidSiteActions {
    static let discord = URL(string: "https://discord.com/channels/@me")!

    func mention(guild: Int64, raid: Int64, role: String?) -> URL { RaidLinks.web(guild: guild, raid: raid) }

    func message(characterKey: String) -> URL { Self.discord }

    func apply(guild: Int64) -> URL { StaticSite.url.appending(path: "app/guild/\(guild)") }

    var guilds: URL { StaticSite.url.appending(path: "guilds") }
}

extension AccountCharacter {
    /// The specs.json key (`warrior-fury`), from the class slug and the spec's name.
    var specKey: String? {
        guard let className, let spec, !spec.isEmpty else { return nil }
        return "\(className)-\(spec.lowercased().replacingOccurrences(of: " ", with: "-"))"
    }
}
