import Foundation

/// `GET /v1/me`'s data (api/internal/auth/me_view.go `Me`). Blocks the
/// scaffold does not use yet (finished sims, notifications, seats, Discord)
/// are ignored on decode and added by the feature lanes that need them.
public struct Me: Codable, Sendable, Equatable {
    public let user: User
    public let characters: [CharacterCard]
    public let guilds: [Guild]
    public let entitlements: Entitlements
    public let bnetImportedAt: Date?
    public let mainCharacterKey: String?

    /// The chosen main, else the first linked character.
    public var mainCharacter: CharacterCard? {
        characters.first { $0.key == mainCharacterKey } ?? characters.first
    }

    private enum CodingKeys: String, CodingKey {
        case user, characters, guilds, entitlements, bnetImportedAt, mainCharacterKey
    }

    public init(from decoder: any Decoder) throws {
        let c = try decoder.container(keyedBy: CodingKeys.self)
        user = try c.decode(User.self, forKey: .user)
        // Go serialises a nil slice as null.
        characters = try c.decodeIfPresent([CharacterCard].self, forKey: .characters) ?? []
        guilds = try c.decodeIfPresent([Guild].self, forKey: .guilds) ?? []
        entitlements = try c.decode(Entitlements.self, forKey: .entitlements)
        bnetImportedAt = try c.decodeIfPresent(Date.self, forKey: .bnetImportedAt)
        mainCharacterKey = try c.decodeIfPresent(String.self, forKey: .mainCharacterKey)
    }
}
