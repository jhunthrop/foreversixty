import Foundation
import Models
import Networking
import Observation

/// What Raid history draws: the nights so far and the cursor for the next page.
struct HistoryContent: Equatable, Sendable {
    let rows: [RaidNightRow]
    let nextCursor: String?
    /// Set while the cached first page shows because a refresh could not connect.
    let offlineSince: Date?

    var hasMore: Bool { nextCursor != nil }
}

/// Loads the guild's logged raid nights 20 at a time (spec F16.3), the first page cache first.
@MainActor
@Observable
final class HistoryStore {
    private(set) var phase: GuildLoadPhase<HistoryContent> = .loading
    private(set) var loadingMore = false
    /// The last next-page read failed: `history.moreError` and Try again under the list.
    private(set) var moreFailed = false

    private let guild: Int64
    private let service: any GuildExtrasService
    private let cache: any RaidCache
    private let now: @Sendable () -> Date

    init(guild: Int64, service: any GuildExtrasService, cache: any RaidCache, now: @escaping @Sendable () -> Date) {
        self.guild = guild
        self.service = service
        self.cache = cache
        self.now = now
    }

    private var cacheKey: String { "guild-\(guild)-history" }

    var content: HistoryContent? {
        if case let .loaded(content, _) = phase { return content }
        return nil
    }

    func load() async {
        if case .loading = phase, let cached = cache.read(GuildRaids.self, key: cacheKey) {
            phase = .loaded(Self.content(cached.value, offlineSince: nil), offlineSince: nil)
        }
        do {
            let page = try await service.history(guild: guild, cursor: nil)
            cache.write(page, key: cacheKey, at: now())
            moreFailed = false
            phase = .loaded(Self.content(page, offlineSince: nil), offlineSince: nil)
        } catch {
            if SignupRefusal.isOffline(error), let cached = cache.read(GuildRaids.self, key: cacheKey) {
                let content = Self.content(cached.value, offlineSince: cached.savedAt)
                phase = .loaded(content, offlineSince: cached.savedAt)
            } else if content == nil {
                phase = .failed
            }
        }
    }

    /// The next page, when there is one; offline the cached page has no more.
    func loadMore() async {
        guard let content, content.offlineSince == nil, let cursor = content.nextCursor, !loadingMore else { return }
        loadingMore = true
        defer { loadingMore = false }
        do {
            let page = try await service.history(guild: guild, cursor: cursor)
            moreFailed = false
            let next = HistoryContent(rows: content.rows + page.rows, nextCursor: page.nextCursor, offlineSince: nil)
            phase = .loaded(next, offlineSince: nil)
        } catch {
            moreFailed = true
        }
    }

    private static func content(_ page: GuildRaids, offlineSince: Date?) -> HistoryContent {
        HistoryContent(rows: page.rows, nextCursor: offlineSince == nil ? page.nextCursor : nil, offlineSince: offlineSince)
    }
}
