import AppIntents
import CoreSpotlight
import Foundation
import Routing

/// Lane F15: the widgets' one CharacterEntity is the Spotlight character too (spec §5.7), so
/// Spotlight, Siri and Shortcuts see the same things.
extension CharacterEntity: IndexedEntity {}

/// A planned raid in Spotlight: `Onyxia's Lair · Thu 8 pm`. The id is `{guild}/{raid}`.
struct RaidEntity: IndexedEntity {
    static let typeDisplayRepresentation = TypeDisplayRepresentation(name: "intent.raid")
    static let defaultQuery = RaidQuery()

    let id: String
    let title: String
    let subtitle: String

    var displayRepresentation: DisplayRepresentation {
        DisplayRepresentation(title: "\(title)", subtitle: "\(subtitle)", image: .init(systemName: "calendar"))
    }

    /// `/app/raid/{guild}/{raid}`.
    var link: DeepLink? {
        let parts = id.split(separator: "/").compactMap { Int64($0) }
        return parts.count == 2 ? .raid(guild: parts[0], raid: parts[1]) : nil
    }
}

/// A guide saved on the phone: `Elemental Shaman guide`. The id is the file key, `shaman-elemental`.
struct GuideEntity: IndexedEntity {
    static let typeDisplayRepresentation = TypeDisplayRepresentation(name: "guides.title")
    static let defaultQuery = GuideQuery()

    let id: String
    let title: String
    let subtitle: String

    var displayRepresentation: DisplayRepresentation {
        DisplayRepresentation(title: "\(title)", subtitle: "\(subtitle)", image: .init(systemName: "book"))
    }

    /// `/app/guides/{class}/{spec}`.
    var link: DeepLink? {
        guard let dash = id.firstIndex(of: "-") else { return nil }
        return .guide(className: String(id[..<dash]), spec: String(id[id.index(after: dash)...]))
    }
}

/// An entity built from one Spotlight row.
protocol SpotlightRowEntity: IndexedEntity where ID == String {
    init(id: String, title: String, subtitle: String)
}

extension RaidEntity: SpotlightRowEntity {}
extension GuideEntity: SpotlightRowEntity {}

/// Resolves a tapped raid or guide from the catalog the app last indexed.
protocol SpotlightRowQuery: EntityQuery where Entity: SpotlightRowEntity {
    static var part: KeyPath<SpotlightCatalog, [SpotlightRow]> { get }
}

extension SpotlightRowQuery {
    private var rows: [SpotlightRow] { SpotlightStore().load()[keyPath: Self.part] }

    func entities(for identifiers: [String]) async throws -> [Entity] {
        let rows = rows
        return identifiers.compactMap { id in
            rows.first { $0.id == id }.map { Entity(id: $0.id, title: $0.title, subtitle: $0.subtitle) }
        }
    }

    func suggestedEntities() async throws -> [Entity] {
        rows.map { Entity(id: $0.id, title: $0.title, subtitle: $0.subtitle) }
    }
}

struct RaidQuery: SpotlightRowQuery {
    typealias Entity = RaidEntity
    static var part: KeyPath<SpotlightCatalog, [SpotlightRow]> { \.raids }
}

struct GuideQuery: SpotlightRowQuery {
    typealias Entity = GuideEntity
    static var part: KeyPath<SpotlightCatalog, [SpotlightRow]> { \.guides }
}

/// The last indexed catalog, so a Spotlight tap resolves to its entity.
struct SpotlightStore {
    static let key = "spotlight.catalog"
    var defaults: UserDefaults = .standard

    func load() -> SpotlightCatalog {
        guard let data = defaults.data(forKey: Self.key) else { return .empty }
        return (try? JSONDecoder().decode(SpotlightCatalog.self, from: data)) ?? .empty
    }

    func save(_ catalog: SpotlightCatalog) {
        defaults.set(try? JSONEncoder().encode(catalog), forKey: Self.key)
    }

    func clear() { defaults.removeObject(forKey: Self.key) }
}

// MARK: Opening a tapped entity

struct OpenCharacterIntent: OpenIntent {
    static let title: LocalizedStringResource = "intent.character"
    static let isDiscoverable = false

    @Parameter(title: "intent.character") var target: CharacterEntity

    @MainActor
    func perform() async throws -> some IntentResult {
        IntentRouter.shared.open(target.characterKey.map { .character(key: $0) } ?? .tab(.home))
        return .result()
    }
}

struct OpenRaidIntent: OpenIntent {
    static let title: LocalizedStringResource = "intent.raid"
    static let isDiscoverable = false

    @Parameter(title: "intent.raid") var target: RaidEntity

    @MainActor
    func perform() async throws -> some IntentResult {
        IntentRouter.shared.open(target.link.map(IntentDestination.link) ?? .tab(.raids))
        return .result()
    }
}

struct OpenGuideIntent: OpenIntent {
    static let title: LocalizedStringResource = "guides.title"
    static let isDiscoverable = false

    @Parameter(title: "guides.title") var target: GuideEntity

    @MainActor
    func perform() async throws -> some IntentResult {
        IntentRouter.shared.open(target.link.map(IntentDestination.link) ?? .tab(.guides))
        return .result()
    }
}

/// Writes the catalog to Spotlight: each type's old entries go, the new ones are indexed. Sign
/// out deletes the whole index (spec §5.7).
struct SpotlightIndexer: Sendable {
    func index(_ catalog: SpotlightCatalog) async {
        SpotlightStore().save(catalog)
        WidgetSnapshotStore.shared().saveSpotlightCharacters(catalog.characters.map { WidgetCharacterChoice(key: $0.id, name: $0.title) })
        let index = CSSearchableIndex.default()
        try? await index.deleteAppEntities(ofType: CharacterEntity.self)
        try? await index.deleteAppEntities(ofType: RaidEntity.self)
        try? await index.deleteAppEntities(ofType: GuideEntity.self)
        try? await index.indexAppEntities(catalog.characters.map { CharacterEntity(id: $0.id, name: $0.title, subtitle: $0.subtitle) })
        try? await index.indexAppEntities(catalog.raids.map { RaidEntity(id: $0.id, title: $0.title, subtitle: $0.subtitle) })
        try? await index.indexAppEntities(catalog.guides.map { GuideEntity(id: $0.id, title: $0.title, subtitle: $0.subtitle) })
    }

    func clear() async {
        SpotlightStore().clear()
        try? await CSSearchableIndex.default().deleteAllSearchableItems()
    }
}
