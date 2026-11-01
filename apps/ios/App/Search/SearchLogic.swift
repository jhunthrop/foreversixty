import Foundation
import Models

/// The scope chips (spec F12.1): `All` caps each group at five; a scope lifts the cap.
enum SearchScope: String, CaseIterable, Sendable {
    case all, items, characters, guides, talents

    /// The scope a group's `Show all {n}` opens.
    init(_ group: SearchGroupKind) {
        switch group {
        case .guides: self = .guides
        case .talents: self = .talents
        case .items: self = .items
        case .characters: self = .characters
        }
    }

    func shows(_ group: SearchGroupKind) -> Bool { self == .all || self == SearchScope(group) }
}

/// The result groups, in the order `All` draws them (spec F12.2).
enum SearchGroupKind: CaseIterable, Sendable {
    case guides, talents, items, characters
}

/// How a name matched, best first (spec F12.2: exact, prefix, word prefix, contains).
enum SearchMatch: Int, Comparable, Sendable {
    case exact, prefix, wordPrefix, contains

    static func < (lhs: SearchMatch, rhs: SearchMatch) -> Bool { lhs.rawValue < rhs.rawValue }
}

/// One group's rows and how many matched in all, for `Show all {n}`.
struct SearchGroup<Row: Sendable>: Sendable {
    let rows: [Row]
    let total: Int

    static var empty: SearchGroup { SearchGroup(rows: [], total: 0) }

    var hasMore: Bool { total > rows.count }
}

extension SearchGroup: Equatable where Row: Equatable {}

/// Who is searching: the current character's class (its items rank first), region and ruleset
/// (its characters rank first), and the account's own characters (the `You` tag).
struct SearchViewer: Equatable, Sendable {
    var classSlug: String?
    var region: String?
    var ruleset: String?
    var ownKeys: Set<String> = []

    static let anonymous = SearchViewer()
}

/// What the results draw for one query and scope.
struct SearchResults: Equatable, Sendable {
    let guides: SearchGroup<SearchGuideHit>
    let talents: SearchGroup<SearchTalent>
    let items: SearchGroup<SearchItem>
    let characters: SearchGroup<CharacterSearchRow>

    var isEmpty: Bool {
        guides.rows.isEmpty && talents.rows.isEmpty && items.rows.isEmpty && characters.rows.isEmpty
    }

    static let empty = SearchResults(guides: .empty, talents: .empty, items: .empty, characters: .empty)
}

/// The search rules (spec F12.2), pure: the corpus, the query and the viewer in, the rows out.
enum SearchLogic {
    /// `All` shows up to five per group.
    static let allCap = 5
    /// Characters are asked of the API from two letters, 250 ms after the last keystroke, up to 8.
    static let minCharacterQuery = 2
    static let characterDebounce: Duration = .milliseconds(250)
    static let characterLimit = 8
    /// Asked for more than shown, so the viewer's region and ruleset can rank first.
    static let characterAsk = 10

    static func normalized(_ query: String) -> String {
        query.trimmingCharacters(in: .whitespacesAndNewlines).lowercased()
    }

    static func match<Value>(_ entry: SearchEntry<Value>, query: String) -> SearchMatch? {
        match(key: entry.key, wordStarts: entry.wordStarts, query: query)
    }

    static func match(key: String, wordStarts: [String.Index], query: String) -> SearchMatch? {
        guard !query.isEmpty else { return nil }
        if key == query { return .exact }
        if key.hasPrefix(query) { return .prefix }
        if wordStarts.contains(where: { key[$0...].hasPrefix(query) }) { return .wordPrefix }
        return key.contains(query) ? .contains : nil
    }

    /// Whether this query asks the API for characters (two letters, a scope that shows them).
    static func asksCharacters(_ query: String, scope: SearchScope) -> Bool {
        scope.shows(.characters) && normalized(query).count >= minCharacterQuery
    }

    static func results(
        corpus: SearchCorpus?, query: String, scope: SearchScope, viewer: SearchViewer, characters: [CharacterSearchRow]
    ) -> SearchResults {
        let query = normalized(query)
        guard !query.isEmpty else { return .empty }
        return SearchResults(
            guides: group(scope, .guides) { guides(corpus?.guides ?? [], query: query, viewer: viewer) },
            talents: group(scope, .talents) { talents(corpus?.talents ?? [], query: query, viewer: viewer) },
            items: group(scope, .items) { items(corpus?.items ?? [], query: query, viewer: viewer) },
            characters: group(scope, .characters) { self.characters(characters, query: query, viewer: viewer) }
        )
    }

    private static func group<Row>(_ scope: SearchScope, _ kind: SearchGroupKind, rows: () -> [Row]) -> SearchGroup<Row> {
        guard scope.shows(kind) else { return .empty }
        let all = rows()
        return SearchGroup(rows: scope == .all ? Array(all.prefix(allCap)) : all, total: all.count)
    }

    /// Items: match, then the viewer's class-usable items, then quality, then item level.
    static func items(_ entries: [SearchEntry<SearchItem>], query: String, viewer: SearchViewer) -> [SearchItem] {
        ranked(entries, query: query) { lhs, rhs in
            let usable = { (item: SearchItem) in viewer.classSlug.map(item.classes.contains) ?? false }
            if usable(lhs) != usable(rhs) { return usable(lhs) ? .orderedAscending : .orderedDescending }
            if lhs.quality != rhs.quality { return lhs.quality > rhs.quality ? .orderedAscending : .orderedDescending }
            if lhs.itemLevel != rhs.itemLevel { return lhs.itemLevel > rhs.itemLevel ? .orderedAscending : .orderedDescending }
            return .orderedSame
        }
    }

    /// Talents: match, then the viewer's class.
    static func talents(_ entries: [SearchEntry<SearchTalent>], query: String, viewer: SearchViewer) -> [SearchTalent] {
        ranked(entries, query: query) { lhs, rhs in classFirst(lhs.className, rhs.className, viewer: viewer) }
    }

    /// Guides: match, then the viewer's class.
    static func guides(_ entries: [SearchEntry<SearchGuideHit>], query: String, viewer: SearchViewer) -> [SearchGuideHit] {
        ranked(entries, query: query) { lhs, rhs in classFirst(lhs.guide.className, rhs.guide.className, viewer: viewer) }
    }

    /// Characters as the API answered, the viewer's region and ruleset first, then by match; at most 8.
    static func characters(_ rows: [CharacterSearchRow], query: String, viewer: SearchViewer) -> [CharacterSearchRow] {
        let entries = rows.map { SearchEntry($0, name: $0.name) }
        let home = { (row: CharacterSearchRow) in row.region == viewer.region && row.ruleset == viewer.ruleset }
        let sorted = ranked(entries, query: query, keepsUnmatched: true, groupFirst: true) { lhs, rhs in
            if home(lhs) != home(rhs) { return home(lhs) ? .orderedAscending : .orderedDescending }
            return .orderedSame
        }
        return Array(sorted.prefix(characterLimit))
    }

    private static func classFirst(_ lhs: String, _ rhs: String, viewer: SearchViewer) -> ComparisonResult {
        guard let mine = viewer.classSlug, (lhs == mine) != (rhs == mine) else { return .orderedSame }
        return lhs == mine ? .orderedAscending : .orderedDescending
    }

    /// Matching entries by match quality, then `tiebreak`, then name. With `groupFirst` the
    /// tiebreak outranks the match (characters: home region first). `keepsUnmatched` keeps rows
    /// the server matched by its own rule (a slug), after every local match.
    private static func ranked<Value>(
        _ entries: [SearchEntry<Value>], query: String, keepsUnmatched: Bool = false, groupFirst: Bool = false,
        tiebreak: (Value, Value) -> ComparisonResult
    ) -> [Value] {
        let scored: [(entry: SearchEntry<Value>, match: Int)] = entries.compactMap { entry in
            if let match = match(entry, query: query) { return (entry, match.rawValue) }
            return keepsUnmatched ? (entry, SearchMatch.contains.rawValue + 1) : nil
        }
        return scored.sorted { lhs, rhs in
            let order = tiebreak(lhs.entry.value, rhs.entry.value)
            if groupFirst, order != .orderedSame { return order == .orderedAscending }
            if lhs.match != rhs.match { return lhs.match < rhs.match }
            if order != .orderedSame { return order == .orderedAscending }
            return lhs.entry.key < rhs.entry.key
        }.map(\.entry.value)
    }
}

/// What a result opens, and how Recent remembers it.
enum SearchTarget: Codable, Hashable, Sendable {
    case item(id: Int)
    /// `key` is the character's `{region}/{ruleset}/{slug}`, the web page's path.
    case character(key: String)
    case guide(className: String, spec: String)
    case talent(id: Int, className: String)
}

/// One opened result, as Recent draws it.
struct RecentSearch: Codable, Equatable, Sendable, Identifiable {
    let target: SearchTarget
    let title: String
    let icon: String?
    let quality: Int?
    let className: String?

    var id: SearchTarget { target }
}

/// Recent (spec F12.1): the last 20 things opened on this phone, newest first, each once.
enum RecentSearches {
    static let cap = 20

    static func adding(_ entry: RecentSearch, to list: [RecentSearch]) -> [RecentSearch] {
        Array(([entry] + list.filter { $0.target != entry.target }).prefix(cap))
    }
}

/// Where Recent is kept: on this phone only, cleared at sign-out.
protocol RecentSearchStorage: Sendable {
    func load() -> [RecentSearch]
    func save(_ list: [RecentSearch])
    func clear()
}

struct UserDefaultsRecentSearches: RecentSearchStorage {
    private static let key = "gg.foreversixty.search.recent"
    // UserDefaults is documented thread-safe.
    private nonisolated(unsafe) let defaults: UserDefaults

    init(defaults: UserDefaults = .standard) {
        self.defaults = defaults
    }

    func load() -> [RecentSearch] {
        defaults.data(forKey: Self.key).flatMap { try? JSONDecoder().decode([RecentSearch].self, from: $0) } ?? []
    }

    func save(_ list: [RecentSearch]) {
        defaults.set(try? JSONEncoder().encode(list), forKey: Self.key)
    }

    func clear() {
        defaults.removeObject(forKey: Self.key)
    }
}
