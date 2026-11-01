import Foundation
import Models
import Networking
import Observation

/// What Your reports draws (spec F10.1): the rows so far, and the offline stamp.
struct ReportsContent: Equatable, Sendable {
    let rows: [ReportRow]
    /// `GET /v1/reports`'s total, so the list knows whether another page exists.
    let total: Int
    let page: Int
    /// Set while the cached answer shows because a refresh could not connect.
    let offlineSince: Date?
    /// Lane F15: set while the cached answer shows because Forever Sixty did not answer.
    var staleSince: Date?

    var hasMorePages: Bool { rows.count < total }
}

enum ReportsPhase: Equatable, Sendable {
    case signedOut
    case loading
    case loaded(ReportsContent)
    case failed
}

/// Loads Your reports (cache first) and its next pages.
@MainActor
@Observable
final class ReportsStore {
    private(set) var phase: ReportsPhase
    /// The last next-page read failed: the list ends with the error and `Try again`.
    private(set) var pageFailed = false

    private let service: any ReportsService
    private let cache: any RaidCache
    private let now: @Sendable () -> Date
    private let signedIn: Bool
    /// Lane F15: the connection, for the stale strip; the screen sets it, nil keeps v1's rule.
    var connectivity: (any Connectivity)?

    static let cacheKey = "reports"

    init(service: any ReportsService, cache: any RaidCache, now: @escaping @Sendable () -> Date, signedIn: Bool) {
        self.service = service
        self.cache = cache
        self.now = now
        self.signedIn = signedIn
        phase = signedIn ? .loading : .signedOut
    }

    private var content: ReportsContent? {
        if case let .loaded(content) = phase { return content }
        return nil
    }

    func load() async {
        guard signedIn else { return }
        if case .loading = phase, let cached = cache.read(ReportList.self, key: Self.cacheKey) {
            phase = .loaded(Self.content(cached.value, offlineSince: nil))
        }
        do {
            let list = try await service.list(page: 1)
            cache.write(list, key: Self.cacheKey, at: now())
            pageFailed = false
            phase = .loaded(Self.content(list, offlineSince: nil))
        } catch {
            if let cached = cache.read(ReportList.self, key: Self.cacheKey), let since = CacheFreshness.staleSince(
                error, connectivity: connectivity, savedAt: cached.savedAt, limit: .pages, now: now()
            ) {
                var shown = Self.content(cached.value, offlineSince: nil)
                shown.staleSince = since
                phase = .loaded(shown)
            } else if SignupRefusal.isOffline(error), let cached = cache.read(ReportList.self, key: Self.cacheKey) {
                phase = .loaded(Self.content(cached.value, offlineSince: cached.savedAt))
            } else if content == nil {
                // A refresh that fails keeps what is drawn; only a first load shows the error.
                phase = .failed
            }
        }
    }

    /// The next page, when the list has one.
    func loadMore() async {
        guard let content, content.offlineSince == nil, content.hasMorePages else { return }
        do {
            let next = try await service.list(page: content.page + 1)
            pageFailed = false
            phase = .loaded(ReportsContent(rows: content.rows + next.rows, total: next.total, page: next.page, offlineSince: nil))
        } catch {
            pageFailed = true
        }
    }

    private static func content(_ list: ReportList, offlineSince: Date?) -> ReportsContent {
        ReportsContent(rows: list.rows, total: list.total, page: list.page, offlineSince: offlineSince)
    }

    /// Home's Logs row sub from the newest report: the first page read (and cached), else the
    /// cached page, else nil while neither exists.
    static func homeLine(service: any ReportsService, cache: any RaidCache, now: @Sendable () -> Date) async -> String? {
        do {
            let list = try await service.list(page: 1)
            cache.write(list, key: cacheKey, at: now())
            return ReportsCopy.homeLine(list)
        } catch {
            return cache.read(ReportList.self, key: cacheKey).map { ReportsCopy.homeLine($0.value) }
        }
    }
}
