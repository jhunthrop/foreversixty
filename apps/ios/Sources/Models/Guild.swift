import Foundation

/// A guild the account belongs to, as `GET /v1/me` serves `guilds`
/// (api/internal/auth/store.go `Guild`). Billing, claim and verify blocks are
/// left to the feature lanes that need them; unknown keys are ignored.
public struct Guild: Codable, Sendable, Equatable, Identifiable {
    public let id: Int64
    public let region: String
    public let ruleset: String
    public let name: String
    public let rank: String?
    public let verified: Bool
    public let verifiedCount: Int
    /// What this account may do in the guild ("approve", "settings", ...).
    /// The app draws a control exactly when its word is here, never from the rank.
    public let can: [String]
    public let rulesetLabel: String?
    public let slug: String?
    public let crestUrl: URL?
    public let faction: String?
    public let rosterCount: Int?

    private enum CodingKeys: String, CodingKey {
        case id, region, ruleset, name, rank, verified, verifiedCount, can
        case rulesetLabel, slug, crestUrl, faction, rosterCount
    }

    public init(from decoder: any Decoder) throws {
        let c = try decoder.container(keyedBy: CodingKeys.self)
        id = try c.decode(Int64.self, forKey: .id)
        region = try c.decode(String.self, forKey: .region)
        ruleset = try c.decode(String.self, forKey: .ruleset)
        name = try c.decode(String.self, forKey: .name)
        rank = try c.decodeIfPresent(String.self, forKey: .rank)
        verified = try c.decode(Bool.self, forKey: .verified)
        verifiedCount = try c.decode(Int.self, forKey: .verifiedCount)
        // Go serialises a nil slice as null.
        can = try c.decodeIfPresent([String].self, forKey: .can) ?? []
        rulesetLabel = try c.decodeIfPresent(String.self, forKey: .rulesetLabel)
        slug = try c.decodeIfPresent(String.self, forKey: .slug)
        crestUrl = try c.decodeIfPresent(URL.self, forKey: .crestUrl)
        faction = try c.decodeIfPresent(String.self, forKey: .faction)
        rosterCount = try c.decodeIfPresent(Int.self, forKey: .rosterCount)
    }
}
