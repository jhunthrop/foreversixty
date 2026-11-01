@testable import ForeverSixty
import Foundation
import Models
import Networking

/// F12's test data, from the repo's active build 70707: every class's items
/// file folded into a search index as the site's generator does, the talents files, the spec
/// catalogue's guides, and the contract fixtures for the item-uses shard and the characters.
enum SearchTestData {
    static let build = "1.60.1.70707"
    static let classes = ["druid", "hunter", "mage", "paladin", "priest", "rogue", "shaman", "warlock", "warrior"]
    static let cape = 11623
    private static var buildDir: URL { BisTestData.repo.appending(path: "data/builds/\(build)") }

    /// Lightningbolltz, the boards' Troll Elemental Shaman at 60 (Horde).
    static let lightningbolltz = AccountCharacter(
        key: "us/normal/lightningbolltz", region: "us", ruleset: "normal", name: "Lightningbolltz",
        source: "bnet", className: "shaman", level: 60, faction: "horde", spec: "Elemental"
    )

    static let fury = AccountCharacter(
        key: "us/pvp/obnoxious-yell", region: "us", ruleset: "pvp", name: "Obnoxious Yell",
        source: "bnet", className: "warrior", level: 60, faction: "horde", spec: "Fury"
    )

    static var context: SearchContext {
        SearchContext(character: lightningbolltz, ownKeys: [lightningbolltz.key])
    }

    // MARK: The index

    static func itemsFile(_ classSlug: String) throws -> ItemsFile {
        var root = try JSONSerialization.jsonObject(with: Data(contentsOf: buildDir.appending(path: "items/\(classSlug).json"))) as? [String: Any] ?? [:]
        // The cape's 14 spell power is one on-equip spell that also gives 14 healing: the
        // `equip_stats` the next data run publishes (the committed 70707 files predate it).
        if var items = root["items"] as? [[String: Any]] {
            items = items.map { item in
                guard item["id"] as? Int == cape else { return item }
                return item.merging(["equip_stats": ["spell_power": 14, "healing": 14]]) { _, new in new }
            }
            root["items"] = items
        }
        return try WireFormat.decoder().decode(ItemsFile.self, from: JSONSerialization.data(withJSONObject: root))
    }

    static func talentsFile(_ classSlug: String) throws -> TalentsFile {
        try WireFormat.decoder().decode(TalentsFile.self, from: Data(contentsOf: buildDir.appending(path: "talents/\(classSlug).json")))
    }

    static func index() throws -> SearchIndex {
        var items: [Int: (item: Item, classes: [String])] = [:]
        var order: [Int] = []
        for classSlug in classes {
            for item in try itemsFile(classSlug).items {
                if items[item.id] == nil { order.append(item.id) }
                items[item.id, default: (item, [])].classes.append(classSlug)
            }
        }
        let searchItems = order.compactMap { id in
            items[id].map { SearchItem(
                id: id, name: $0.item.name, icon: $0.item.icon, quality: $0.item.quality, slot: $0.item.slot,
                itemLevel: $0.item.itemLevel, requiredLevel: $0.item.requiredLevel, classes: $0.classes
            ) }
        }
        let talents = try classes.flatMap { classSlug in
            try talentsFile(classSlug).trees.flatMap { tree in
                tree.talents.map { SearchTalent(
                    id: $0.id, name: $0.name, icon: $0.icon, className: classSlug, tree: tree.name,
                    specName: tree.name, maxRank: $0.maxRank
                ) }
            }
        }
        let guides = try BisTestData.catalog().specs.map { spec in
            SearchGuide(
                className: spec.classSlug, spec: SearchLinks.guideSpec(specKey: spec.spec, className: spec.classSlug),
                specName: spec.name, role: spec.role, icon: spec.icon
            )
        }
        return SearchIndex(build: build, itemUsesShards: 64, items: searchItems, talents: talents, guides: guides)
    }

    nonisolated(unsafe) private static var cachedCorpus: SearchCorpus?

    static func corpus() throws -> SearchCorpus {
        if let cachedCorpus { return cachedCorpus }
        let built = SearchCorpus(try index())
        cachedCorpus = built
        return built
    }

    // MARK: Item uses and the tooltip

    /// The three items the captures use, with the site generator's item-uses rows for them over
    /// build 70707 (the file Android's tests read, one copy for both apps).
    static let runnerUp = 17755 // Satyrmane Sash: Elemental 60 raid, runner-up 2 for Waist
    static let notListed = 20203 // Defiler's Mail Pauldrons: on no list; a vendor, a reputation and a set

    static func uses(_ id: Int) throws -> ItemUses {
        let url = BisTestData.repo.appending(path: "apps/android/app/src/test/resources/search/item-uses-70707.json")
        let file = try WireFormat.decoder().decode([String: ItemUses].self, from: Data(contentsOf: url))
        return try requireValue(file[String(id)])
    }

    static func tooltipSources(classSlug: String = "shaman", spec: String = "shaman-elemental") throws -> ItemTooltipSources {
        let items = try itemsFile(classSlug).items
        let bis = try WireFormat.decoder().decode(BisFile.self, from: Data(contentsOf: buildDir.appending(path: "bis/\(spec).json")))
        let sets = try WireFormat.decoder().decode([ItemSet].self, from: Data(contentsOf: buildDir.appending(path: "sets.json")))
        return ItemTooltipSources(items: Dictionary(items.map { ($0.id, $0) }, uniquingKeysWith: { first, _ in first }), sets: sets, bis: bis)
    }

    // MARK: Characters

    static func row(_ name: String, className: String?, level: Int?, region: String = "us", ruleset: String = "normal") -> CharacterSearchRow {
        CharacterSearchRow(
            key: "\(region)/\(ruleset)/\(name.lowercased())", region: region, ruleset: ruleset, name: name,
            className: className, level: level, spec: nil
        )
    }

    static let lightningCharacters = [
        row("Lightningbolltz", className: "shaman", level: 60),
        row("Lightningrod", className: "mage", level: 52),
        row("Lightningpaw", className: "druid", level: 60, region: "eu", ruleset: "pvp"),
    ]

    static let eleCharacters = [
        row("Elementalx", className: "shaman", level: 60),
        row("Elegy", className: "priest", level: 58),
        row("Eleanora", className: "mage", level: 41, region: "eu", ruleset: "pvp"),
    ]

    static let recents = [
        RecentSearch(target: .item(id: cape), title: "Spritecaster Cape", icon: "inv_misc_cape_20", quality: 3, className: nil),
        RecentSearch(target: .guide(className: "shaman", spec: "elemental"), title: "Elemental Shaman", icon: "spell_nature_lightning", quality: nil, className: "shaman"),
    ]

    static func requireValue<T>(_ value: T?) throws -> T {
        guard let value else { throw CocoaError(.fileReadCorruptFile) }
        return value
    }
}

/// A character search that answers from a list and counts its calls.
final class FakeCharacterSearch: CharacterSearchService, @unchecked Sendable {
    private let lock = NSLock()
    private var calls: [String] = []
    let answer: Result<[CharacterSearchRow], APIError>

    init(_ answer: Result<[CharacterSearchRow], APIError>) {
        self.answer = answer
    }

    var asked: [String] { lock.withLock { calls } }

    func search(_ query: String, limit _: Int) async throws(APIError) -> [CharacterSearchRow] {
        lock.withLock { calls.append(query) }
        return try answer.get()
    }
}

/// Recent kept in memory.
final class MemoryRecentSearches: RecentSearchStorage, @unchecked Sendable {
    private let lock = NSLock()
    private var list: [RecentSearch]

    init(_ list: [RecentSearch] = []) { self.list = list }

    func load() -> [RecentSearch] { lock.withLock { list } }
    func save(_ list: [RecentSearch]) { lock.withLock { self.list = list } }
    func clear() { lock.withLock { list = [] } }
}
