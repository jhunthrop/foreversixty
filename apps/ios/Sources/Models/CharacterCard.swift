import Foundation

/// The character card: one character linked to the account, as `GET /v1/me`
/// serves each entry of `characters` (api/internal/auth/store.go `Character`).
/// Every field the API omits when unknown is optional here.
public struct CharacterCard: Codable, Sendable, Equatable, Identifiable {
    public let key: String
    public let region: String
    public let ruleset: String
    public let name: String
    public let `class`: String?
    public let realm: String?
    public let level: Int?
    public let faction: String?
    public let race: String?
    public let gender: String?
    public let itemLevel: Int?
    public let avatarUrl: URL?
    public let renderUrl: URL?
    /// "export" or "bnet": which path last wrote this character.
    public let source: String
    public let build: CharacterBuild?
    public let guild: CharacterGuild?
    /// The primary talent tree, absent until known.
    public let spec: String?

    public var id: String { key }
}

/// The character's latest simmable build (Go `CharacterBuild`).
public struct CharacterBuild: Codable, Sendable, Equatable {
    public let source: String
    public let capturedAt: Date
    /// Slot name to item id.
    public let gear: [String: Int]?
    public let talents: BuildTalents?
    public let level: Int?
    public let dataBuild: String?
    public let medianSyncGapSec: Int?
    public let syncError: String?
}

/// Per-tree rank strings and point totals, in tree order (Go `BuildTalents`).
public struct BuildTalents: Codable, Sendable, Equatable {
    public let trees: [String]
    public let points: [Int]
}

/// The character's guild membership (Go `CharacterGuild`).
public struct CharacterGuild: Codable, Sendable, Equatable {
    public let id: Int64
    public let name: String
    public let rank: String
    public let rankIndex: Int?
    public let verified: Bool
    public let attendance: Attendance?
}

/// Raids attended over raids counted (Go `Attendance`).
public struct Attendance: Codable, Sendable, Equatable {
    public let present: Int
    public let nights: Int
}
