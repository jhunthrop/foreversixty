import Models
import Networking

/// The Raid tab's number (flows spec 3): a member's unanswered next raid (1) or an officer's open
/// items, from `GET /v1/guilds/{id}/raid-badge`. Zero, and no guild, show none.
enum RaidBadgeCount {
    static func load(guild: Int64?, client: APIClient) async -> Int {
        guard let guild, let badge = try? await client.send(.getGuildRaidBadge(id: guild)) else { return 0 }
        return badge.count
    }
}
