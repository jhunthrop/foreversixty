import Foundation

/// `GET /v1/characters/{region}/{ruleset}/{name}`: the public character card with the Forever
/// Score and the gear panel (docs/contracts/app-v1.md "Characters"). Every score field is nil
/// for a character with no rated fight: render "no logs yet", never 0. The rules behind picks,
/// gaps and scores live in the API; the app never re-derives them. Feature lanes add the
/// fight lists (`best`, `history`) when a screen reads them.
public struct CharacterProfile: Decodable, Sendable, Equatable {
    public let character: Header
    public let score: Double?
    public let tier: String?
    public let asOf: String?
    public let components: ScoreComponents?
    public let gear: [GearRow]
    public let gearGapTotal: Double?
    public let gearSpec: String?
    public let supporter: Bool

    /// Who the card is for.
    public struct Header: Decodable, Sendable, Equatable {
        public let key: String
        public let region: String
        public let ruleset: String
        public let name: String
        public let `class`: String?
        public let level: Int?
        public let race: String?
        public let realm: String?
        public let guild: GuildLine?
        public let avatarUrl: URL?
        public let renderUrl: URL?
    }

    public struct GuildLine: Decodable, Sendable, Equatable {
        public let name: String
        public let rank: String
    }

    /// The score's five parts.
    public struct ScoreComponents: Decodable, Sendable, Equatable {
        public let gear: Double
        public let experience: Double
        public let parses: Double
        public let attendance: Double
        public let discipline: Double
    }

    /// One slot of the gear panel: what is worn, the band's pick and the points it adds.
    public struct GearRow: Decodable, Sendable, Equatable {
        public let slot: String
        public let worn: WornItem?
        public let pick: PickedItem?
        public let gap: Double?
    }

    /// `name` and `itemLevel` are nil when the data build does not know the item.
    public struct WornItem: Decodable, Sendable, Equatable {
        public let id: Int
        public let name: String?
        public let itemLevel: Int?
    }

    /// `chance` is the drop fraction, nil for a source with no rate.
    public struct PickedItem: Decodable, Sendable, Equatable {
        public let id: Int
        public let name: String
        public let source: String
        public let chance: Double?
    }
}
