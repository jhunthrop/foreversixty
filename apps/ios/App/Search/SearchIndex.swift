import Foundation
import Models
import Networking

/// One searchable name, prepared once: lowercased, with the start of every word after the first,
/// so a keystroke is a pure in-memory comparison (spec §11: keystroke to results ≤ 50 ms).
struct SearchEntry<Value: Sendable>: Sendable {
    let value: Value
    let name: String
    /// The lowercased name.
    let key: String
    /// Where each later word starts in `key` (`Elemental Fury` → the `f`).
    let wordStarts: [String.Index]

    init(_ value: Value, name: String) {
        self.value = value
        self.name = name
        let key = name.lowercased()
        self.key = key
        wordStarts = key.indices.filter { index in
            index != key.startIndex && !key[key.index(before: index)].isLetter && !key[key.index(before: index)].isNumber
                && (key[index].isLetter || key[index].isNumber)
        }
    }
}

/// A guide as the search lists it: `Elemental Shaman`, `Guide · DPS`.
struct SearchGuideHit: Equatable, Sendable {
    let guide: SearchGuide
    let title: String
}

/// The phone's search over one build's `search.json` (W3-C8), built once per build.
struct SearchCorpus: Sendable {
    let build: String
    let itemUsesShards: Int
    let items: [SearchEntry<SearchItem>]
    let talents: [SearchEntry<SearchTalent>]
    let guides: [SearchEntry<SearchGuideHit>]
    private let itemsByID: [Int: SearchItem]
    private let talentsByID: [Int: SearchTalent]

    init(_ index: SearchIndex) {
        build = index.build
        itemUsesShards = index.itemUsesShards
        items = index.items.map { SearchEntry($0, name: $0.name) }
        talents = index.talents.map { SearchEntry($0, name: $0.name) }
        guides = index.guides.map { guide in
            let title = SearchWords.specClass(spec: guide.specName, classSlug: guide.className)
            return SearchEntry(SearchGuideHit(guide: guide, title: title), name: title)
        }
        itemsByID = Dictionary(index.items.map { ($0.id, $0) }, uniquingKeysWith: { first, _ in first })
        talentsByID = Dictionary(index.talents.map { ($0.id, $0) }, uniquingKeysWith: { first, _ in first })
    }

    func item(_ id: Int) -> SearchItem? { itemsByID[id] }
    func talent(_ id: Int) -> SearchTalent? { talentsByID[id] }

    /// The item-uses shard that holds `id` (`id % item_uses_shards`).
    func shard(of id: Int) -> Int { itemUsesShards > 0 ? id % itemUsesShards : 0 }
}

/// Loads `search.json` for a build through the static data store (cache first, so it works
/// offline once fetched) and keeps the built corpus for the session (spec §11: warm ≤ 150 ms).
actor SearchIndexReader {
    private let store: StaticDataStore
    private var corpus: SearchCorpus?

    init(store: StaticDataStore) {
        self.store = store
    }

    func corpus(build: String) async throws(APIError) -> SearchCorpus {
        if let corpus, corpus.build == build { return corpus }
        let built = SearchCorpus(try await store.file(.getSearchIndex(build: build)))
        corpus = built
        // A newer build is in use: the store swaps to it (spec §14, the store's swap).
        await store.prune(keeping: build)
        return built
    }

    /// The item-uses shard that holds `id`, through the same store.
    func uses(of id: Int, in corpus: SearchCorpus) async throws(APIError) -> ItemUses? {
        let shard = try await store.file(.getItemUses(build: corpus.build, shard: corpus.shard(of: id)))
        return shard.items[String(id)]
    }

    /// A class's talents file, for the talent sheet.
    func talents(className: String, build: String) async throws(APIError) -> TalentsFile {
        try await store.file(.getTalents(build: build, className: className))
    }
}
