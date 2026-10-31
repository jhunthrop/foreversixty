import Foundation
import Models
import Networking
import Observation

/// The last good home answer per character, on disk in Application Support and excluded from
/// backup (spec §5.0 Caching). Sign-out deletes it with `removeAll()`.
protocol HomeCache: Sendable {
    func load(_ key: String) -> HomeCacheEntry?
    func save(_ entry: HomeCacheEntry, for key: String)
    func removeAll()
}

struct HomeCacheEntry: Codable, Equatable, Sendable {
    let home: CharacterHome
    let fetchedAt: Date
}

struct FileHomeCache: HomeCache {
    let directory: URL

    static func standard() -> FileHomeCache? {
        FileManager.default.urls(for: .applicationSupportDirectory, in: .userDomainMask).first
            .map { FileHomeCache(directory: $0.appending(path: "home-answers")) }
    }

    func load(_ key: String) -> HomeCacheEntry? {
        guard let data = try? Data(contentsOf: url(key)) else { return nil }
        return try? WireFormat.decoder().decode(HomeCacheEntry.self, from: data)
    }

    func save(_ entry: HomeCacheEntry, for key: String) {
        guard let data = try? WireFormat.encoder().encode(entry) else { return }
        var folder = directory
        try? FileManager.default.createDirectory(at: folder, withIntermediateDirectories: true)
        var values = URLResourceValues()
        values.isExcludedFromBackup = true
        try? folder.setResourceValues(values)
        try? data.write(to: url(key), options: [.atomic, .completeFileProtectionUntilFirstUserAuthentication])
    }

    func removeAll() {
        try? FileManager.default.removeItem(at: directory)
    }

    private func url(_ key: String) -> URL {
        directory.appending(path: key.replacingOccurrences(of: "/", with: "_") + ".json")
    }
}

/// An in-memory cache for tests and previews.
final class MemoryHomeCache: HomeCache, @unchecked Sendable {
    private let lock = NSLock()
    private var entries: [String: HomeCacheEntry]

    init(_ entries: [String: HomeCacheEntry] = [:]) { self.entries = entries }

    func load(_ key: String) -> HomeCacheEntry? { lock.withLock { entries[key] } }
    func save(_ entry: HomeCacheEntry, for key: String) { lock.withLock { entries[key] = entry } }
    func removeAll() { lock.withLock { entries.removeAll() } }
}

/// What Home shows (spec F2.1 to F2.5).
enum HomePhase: Equatable {
    /// First load, nothing cached: skeletons.
    case loading
    /// An answer; `offlineSince` is the cached answer's time when the last fetch had no connection.
    case loaded(CharacterHome, offlineSince: Date?)
    /// The server failed. A cached answer still feeds the Next rows.
    case failed(cached: CharacterHome?)
}

/// Loads the home answer for one character: the cache draws at once, then the API refreshes it.
/// A refresh never replaces content with skeletons; no connection keeps the cached answer under
/// the offline strip; a server error is the error state.
@MainActor
@Observable
final class HomeStore {
    typealias Fetch = @Sendable (HomeCharacter) async throws(APIError) -> CharacterHome

    static let staleAfter: TimeInterval = 5 * 60

    private(set) var phase: HomePhase = .loading
    private(set) var fetchedAt: Date?
    private let fetch: Fetch
    private let cache: any HomeCache
    private let now: @Sendable () -> Date

    init(fetch: @escaping Fetch, cache: any HomeCache, now: @escaping @Sendable () -> Date = { Date() }) {
        self.fetch = fetch
        self.cache = cache
        self.now = now
    }

    /// The live store: `GET /v1/characters/{region}/{ruleset}/{name}/home`.
    static func live(client: APIClient, cache: any HomeCache) -> HomeStore {
        HomeStore(fetch: { character throws(APIError) in
            try await client.send(.getCharacterHome(region: character.region, ruleset: character.ruleset, name: character.slug))
        }, cache: cache)
    }

    /// Draws the cached answer, then fetches.
    func load(_ character: HomeCharacter) async {
        if let entry = cache.load(character.key) {
            phase = .loaded(entry.home, offlineSince: nil)
            fetchedAt = entry.fetchedAt
        } else {
            phase = .loading
        }
        await refresh(character)
    }

    /// Refreshes on foreground only when the answer is older than five minutes.
    func refreshIfStale(_ character: HomeCharacter) async {
        guard let fetchedAt, now().timeIntervalSince(fetchedAt) < Self.staleAfter else {
            await refresh(character)
            return
        }
    }

    func refresh(_ character: HomeCharacter) async {
        do {
            let home = try await fetch(character)
            let at = now()
            cache.save(HomeCacheEntry(home: home, fetchedAt: at), for: character.key)
            fetchedAt = at
            phase = .loaded(home, offlineSince: nil)
        } catch {
            fail(error, cached: cache.load(character.key))
        }
    }

    private func fail(_ error: APIError, cached: HomeCacheEntry?) {
        if case .transport = error, let cached {
            phase = .loaded(cached.home, offlineSince: cached.fetchedAt)
        } else {
            phase = .failed(cached: cached?.home)
        }
    }
}
