import Foundation
import Models
import Networking
import Observation

/// What the board is chosen from: the ranked encounters, the phases, and the viewer's defaults.
struct RankingsSetup: Equatable, Sendable {
    let encounters: [RankedEncounter]
    let phases: [PhaseBoundary]
    let defaults: RankingsDefaults
    /// The encounter the board opens on: a link's, else the default (defaults.ts).
    let encounter: String

    var filters: RankingsFilters { RankingsDefaulting.filters(defaults, encounter: encounter) }

    func encounterName(_ slug: String) -> String {
        encounters.first { $0.slug == slug || String($0.id) == slug }?.name ?? slug
    }

    static let encountersKey = "rankings-encounters"
    static let phasesKey = "rankings-phases"

    /// Reads the encounters, the phases and (signed in) the character's history in parallel,
    /// cached for offline. Nil when neither the API nor the cache has the encounters or phases.
    static func load(
        service: any RankingsService, cache: any RaidCache, specs: SpecCatalog, now: Date,
        current: RankingsCharacter?, encounter: String? = nil
    ) async -> RankingsSetup? {
        // The reads run together; the cache is touched after, one entry at a time.
        async let encounterRead = attempt { try await service.encounters() }
        async let phaseRead = attempt { try await service.phases() }
        async let historyRead = history(service: service, current: current)
        let (encounterResult, phaseResult, history) = await (encounterRead, phaseRead, historyRead)
        guard let encounters = settle(encounterResult, cache: cache, key: encountersKey, at: now)?.encounters,
              let phases = settle(phaseResult, cache: cache, key: phasesKey, at: now)?.phases
        else { return nil }
        let defaults = RankingsDefaulting.defaults(current, specs: specs, phases: phases, now: now)
        let opening = encounter.flatMap { link in encounters.first { $0.slug == link || String($0.id) == link }?.slug }
        return RankingsSetup(
            encounters: encounters, phases: phases, defaults: defaults,
            encounter: opening ?? RankingsDefaulting.defaultEncounter(history: history, encounters: encounters)
        )
    }

    private static func attempt<Value: Sendable>(_ call: @Sendable () async throws -> Value) async -> Value? {
        try? await call()
    }

    /// A fresh answer is cached; a failed read falls back to the cache.
    private static func settle<Value: Codable>(_ fresh: Value?, cache: any RaidCache, key: String, at now: Date) -> Value? {
        if let fresh {
            cache.write(fresh, key: key, at: now)
            return fresh
        }
        return cache.read(Value.self, key: key)?.value
    }

    /// The character's fights, for the default encounter; never blocks the board.
    private static func history(service: any RankingsService, current: RankingsCharacter?) async -> [CharacterFight] {
        guard let current, let ref = CharacterRef(key: current.key, name: current.name) else { return [] }
        return (try? await service.character(ref).history) ?? []
    }
}

/// The board as drawn: every row read so far and the first page's answer.
struct BoardContent: Equatable, Sendable {
    let first: RankingsPage
    let rows: [RankingRow]
    let page: Int
    /// The last page read added nothing new (fixture mode answers the same page).
    let exhausted: Bool
    /// Set while the cached board shows because a refresh could not connect.
    let offlineSince: Date?
    /// Lane F15: set while the cached answer shows because Forever Sixty did not answer.
    var staleSince: Date?

    var hasMorePages: Bool { !exhausted && rows.count < first.total }
}

enum BoardPhase: Equatable, Sendable {
    case loading
    case loaded(BoardContent)
    case failed
}

/// Loads the board (cache first) for the current filters, its next pages, and re-aims it.
@MainActor
@Observable
final class RankingsStore {
    private(set) var setup: RankingsSetup?
    private(set) var filters: RankingsFilters?
    private(set) var phase: BoardPhase = .loading
    /// The last next-page read failed: the board ends with the error and `Try again`.
    private(set) var pageFailed = false

    let account: RankingsAccount
    private let service: any RankingsService
    private let cache: any RaidCache
    private let specs: SpecCatalog
    private let now: @Sendable () -> Date
    private let linkedEncounter: String?
    /// Lane F15: the connection, for the stale strip; the screen sets it, nil keeps v1's rule.
    var connectivity: (any Connectivity)?

    init(
        service: any RankingsService, cache: any RaidCache, specs: SpecCatalog, now: @escaping @Sendable () -> Date,
        account: RankingsAccount, encounter: String? = nil
    ) {
        self.service = service
        self.cache = cache
        self.specs = specs
        self.now = now
        self.account = account
        linkedEncounter = encounter
    }

    var encounterName: String {
        guard let setup, let filters else { return "" }
        return setup.encounterName(filters.encounter)
    }

    private var content: BoardContent? {
        if case let .loaded(content) = phase { return content }
        return nil
    }

    /// The first load and pull to refresh.
    func load() async {
        if setup == nil {
            setup = await RankingsSetup.load(
                service: service, cache: cache, specs: specs, now: now(), current: account.current, encounter: linkedEncounter
            )
            filters = setup?.filters
        }
        guard let filters else {
            phase = .failed
            return
        }
        await loadBoard(filters)
    }

    /// A filter choice: applies at once (spec F11.2), back on page one.
    func apply(_ next: RankingsFilters) async {
        guard next != filters else { return }
        filters = next
        phase = .loading
        await loadBoard(next)
    }

    /// `Reset filters`: the viewer's defaults on the current encounter's default.
    func reset() async {
        guard let setup else { return }
        await apply(setup.filters)
    }

    /// Signed out: the F3.2 picker's class and spec aim the board.
    func choose(classSlug: String, spec: String) async {
        guard var next = filters else { return }
        next.classSlug = classSlug
        next.spec = specs.spec(spec)?.name ?? spec
        next.metric = RankMetric.forRole(specs.spec(spec)?.role)
        await apply(next)
    }

    func loadMore() async {
        guard let content, let filters, content.offlineSince == nil, content.hasMorePages else { return }
        do {
            let next = try await service.rankings(query(filters, page: content.page + 1))
            let merged = RankingsPaging.merge(content.rows, next.rows)
            pageFailed = false
            phase = .loaded(BoardContent(
                first: content.first, rows: merged, page: content.page + 1,
                exhausted: merged.count == content.rows.count, offlineSince: nil
            ))
        } catch {
            pageFailed = true
        }
    }

    private func loadBoard(_ filters: RankingsFilters) async {
        let first = query(filters, page: 1)
        let key = Self.cacheKey(first)
        if content == nil, let cached = cache.read(RankingsPage.self, key: key) {
            phase = .loaded(Self.content(cached.value, offlineSince: nil))
        }
        do {
            let page = try await service.rankings(first)
            guard filters == self.filters else { return }
            cache.write(page, key: key, at: now())
            pageFailed = false
            phase = .loaded(Self.content(page, offlineSince: nil))
        } catch {
            guard filters == self.filters else { return }
            if let cached = cache.read(RankingsPage.self, key: key), let since = CacheFreshness.staleSince(
                error, connectivity: connectivity, savedAt: cached.savedAt, limit: .pages, now: now()
            ) {
                var shown = Self.content(cached.value, offlineSince: nil)
                shown.staleSince = since
                phase = .loaded(shown)
            } else if SignupRefusal.isOffline(error), let cached = cache.read(RankingsPage.self, key: key) {
                phase = .loaded(Self.content(cached.value, offlineSince: cached.savedAt))
            } else if content == nil {
                // A refresh that fails keeps what is drawn; only a first load shows the error.
                phase = .failed
            }
        }
    }

    private func query(_ filters: RankingsFilters, page: Int) -> RankingsQuery {
        RankingsQuery.board(filters, character: account.current?.key, page: page)
    }

    private static func content(_ page: RankingsPage, offlineSince: Date?) -> BoardContent {
        BoardContent(first: page, rows: page.rows, page: page.page, exhausted: false, offlineSince: offlineSince)
    }

    /// One cache entry per board: every query field but the page.
    static func cacheKey(_ query: RankingsQuery) -> String {
        [
            "rankings", query.encounter, query.metric, query.spec, query.classSlug, query.phase, query.region,
            query.ruleset, query.faction, query.since, query.character,
        ]
        .map { $0 ?? "" }
        .joined(separator: "|")
    }

    /// Home's Rankings row sub: one read for the current character with the defaults; nil when
    /// neither the API nor the cache answers.
    static func homeLine(
        service: any RankingsService, cache: any RaidCache, specs: SpecCatalog, now: Date, current: RankingsCharacter
    ) async -> String? {
        guard let setup = await RankingsSetup.load(service: service, cache: cache, specs: specs, now: now, current: current),
              !setup.encounter.isEmpty
        else { return nil }
        let filters = setup.filters
        let query = RankingsQuery.board(filters, character: current.key, page: 1)
        let key = cacheKey(query)
        let page: RankingsPage
        do {
            page = try await service.rankings(query)
            cache.write(page, key: key, at: now)
        } catch {
            guard let cached = cache.read(RankingsPage.self, key: key) else { return nil }
            page = cached.value
        }
        return RankingsCopy.homeLine(page: page, filters: filters, encounterName: setup.encounterName(filters.encounter))
    }
}
