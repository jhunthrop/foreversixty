import Foundation
import Models
import Networking
import Observation

/// A guild-extras screen's load state: the cache draws first, offline sits under the strip.
enum GuildLoadPhase<Content: Equatable & Sendable>: Equatable, Sendable {
    case loading
    case loaded(Content, offlineSince: Date?)
    case failed
}

/// Loads `GET /v1/guilds/{id}/home` for the roster (spec F16.1), cache first. The Raids tab's
/// Roster row reads its count from here too, so the pushed roster opens drawn.
@MainActor
@Observable
final class RosterStore {
    private(set) var phase: GuildLoadPhase<GuildHome> = .loading

    let guild: Int64
    private let service: any GuildExtrasService
    private let cache: any RaidCache
    private let now: @Sendable () -> Date

    /// Lane F15: the stale strip's time, and the connection that decides it (nil keeps v1's rule).
    private(set) var staleSince: Date?
    var connectivity: (any Connectivity)?

    init(guild: Int64, service: any GuildExtrasService, cache: any RaidCache, now: @escaping @Sendable () -> Date) {
        self.guild = guild
        self.service = service
        self.cache = cache
        self.now = now
    }

    private var cacheKey: String { "guild-\(guild)-home" }

    var home: GuildHome? {
        if case let .loaded(home, _) = phase { return home }
        return nil
    }

    var list: RosterList? { home.map(RosterList.init(home:)) }
    var onRoster: Int? { home?.counts.onRoster }
    /// The viewer's character in this guild, one of "your characters" for history.
    var viewerKey: String? { home?.viewer.characterKey }

    var offlineSince: Date? {
        if case let .loaded(_, since) = phase { return since }
        return nil
    }

    func load() async {
        if case .loading = phase, let cached = cache.read(GuildHome.self, key: cacheKey) {
            phase = .loaded(cached.value, offlineSince: nil)
        }
        do {
            let home = try await service.home(guild: guild)
            cache.write(home, key: cacheKey, at: now())
            staleSince = nil
            phase = .loaded(home, offlineSince: nil)
        } catch {
            let cached = cache.read(GuildHome.self, key: cacheKey)
            if let cached, let since = CacheFreshness.staleSince(
                error, connectivity: connectivity, savedAt: cached.savedAt, limit: .pages, now: now()
            ) {
                staleSince = since
                phase = .loaded(cached.value, offlineSince: nil)
            } else if SignupRefusal.isOffline(error), let cached {
                phase = .loaded(cached.value, offlineSince: cached.savedAt)
            } else if home == nil {
                phase = .failed
            }
        }
    }
}
