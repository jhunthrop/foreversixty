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
    /// `GET /v1/guilds/{id}/setup`: an officer's setup checklist (the Guild tab's Setup row).
    func setup(guild: Int64) async throws(APIError) -> GuildSetup
    /// `GET /v1/guilds/{id}/applications/mine`: the player's own applications to a guild.
    func myApplications(guild: Int64) async throws(APIError) -> GuildApplications
    /// `GET /v1/guilds/{id}/loot/nights?limit=1`: the newest night's drops and who won them.
    func lootNights(guild: Int64) async throws(APIError) -> LootNights
}

/// The officer's Discord mention (F3): `POST /v1/guilds/{id}/raids/{raid}/ping` with an optional
/// role (`tank`, `healer`, `dps`; none mentions everyone unanswered).
protocol RaidPingService: Sendable {
    func ping(guild: Int64, raid: Int64, role: String?) async throws(APIError) -> RaidPingResult
}

/// What the app cannot do itself yet, so the site does it: applying to a guild and claiming one
/// (the report names the API they need), and the raid page when no Discord bot is connected.
protocol RaidSiteActions: Sendable {
    /// The raid's page on the site (connect the Discord bot there).
    func raid(guild: Int64, raid: Int64) -> URL
    /// Applying to a recruiting guild.
    func apply(guild: Int64) -> URL
    /// Finding a guild: the site's guild finder.
    var guilds: URL { get }
    /// Claiming a guild that is not here.
    var claim: URL { get }
}

struct LiveRaidPingService: RaidPingService {
    let client: APIClient

    func ping(guild: Int64, raid: Int64, role: String?) async throws(APIError) -> RaidPingResult {
        try await client.send(.pingGuildRaid(id: guild, raid: raid, RaidPingRequest(role: role)))
    }
}

extension ApplicantCharacter {
    /// `Alliance Dwarf`, when the API names them.
    var origin: String? {
        let words = [faction?.capitalized, race.map(HomePresenter.words)].compactMap(\.self).filter { !$0.isEmpty }
        return words.isEmpty ? nil : words.joined(separator: " ")
    }
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

    func setup(guild: Int64) async throws(APIError) -> GuildSetup {
        try await client.send(.getGuildSetup(id: guild))
    }

    func myApplications(guild: Int64) async throws(APIError) -> GuildApplications {
        try await client.send(.myGuildApplications(id: guild))
    }

    func lootNights(guild: Int64) async throws(APIError) -> LootNights {
        try await client.send(.getGuildLootNights(id: guild, limit: 1))
    }
}

struct LiveRaidSiteActions: RaidSiteActions {
    func raid(guild: Int64, raid: Int64) -> URL { RaidLinks.web(guild: guild, raid: raid) }

    func apply(guild: Int64) -> URL { StaticSite.url.appending(path: "app/guild/\(guild)") }

    /// The site's guild finder is `/guild` (web nav's Find a guild); `/guilds` is no page.
    var guilds: URL { StaticSite.url.appending(path: "guild") }

    var claim: URL { URL(string: "guild#claim", relativeTo: StaticSite.url)?.absoluteURL ?? guilds }
}

extension AccountCharacter {
    /// The specs.json key (`warrior-fury`), from the class slug and the spec's name.
    var specKey: String? {
        guard let className, let spec, !spec.isEmpty else { return nil }
        return "\(className)-\(spec.lowercased().replacingOccurrences(of: " ", with: "-"))"
    }
}
