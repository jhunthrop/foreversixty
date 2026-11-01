import Foundation
import Models
import Networking
import Observation

enum SimResultPhase: Equatable, Sendable {
    case loading
    case loaded(SimResultModel, offlineSince: Date?)
    case failed
}

/// Loads one result (F5.3), finds its list row for the headline, and marks it seen once per open.
@MainActor
@Observable
final class SimResultStore {
    private(set) var phase: SimResultPhase = .loading

    let id: String
    private let row: SimRow?
    private let service: any SimsService
    private let cache: any RaidCache
    private let specs: SpecCatalog
    private let now: @Sendable () -> Date
    private let markSeen: (String) async -> Void
    private var seenPosted = false
    /// Lane F15: set while the saved result shows because Forever Sixty did not answer.
    private(set) var staleSince: Date?
    /// Lane F15: the connection, for the stale strip; the screen sets it, nil keeps v1's rule.
    var connectivity: (any Connectivity)?

    init(
        id: String,
        row: SimRow?,
        service: any SimsService,
        cache: any RaidCache,
        specs: SpecCatalog,
        now: @escaping @Sendable () -> Date,
        markSeen: @escaping (String) async -> Void
    ) {
        self.id = id
        self.row = row
        self.service = service
        self.cache = cache
        self.specs = specs
        self.now = now
        self.markSeen = markSeen
    }

    private var cacheKey: String { "sim-\(id)" }

    func load() async {
        do {
            let result = try await service.result(id: id)
            cache.write(result, key: cacheKey, at: now())
            staleSince = nil
            phase = .loaded(await model(result), offlineSince: nil)
            await postSeenOnce()
        } catch {
            let cached = cache.read(SimResult.self, key: cacheKey)
            if let cached, let since = CacheFreshness.staleSince(
                error, connectivity: connectivity, savedAt: cached.savedAt, limit: .pages, now: now()
            ) {
                staleSince = since
                phase = .loaded(await model(cached.value), offlineSince: nil)
            } else if SignupRefusal.isOffline(error), let cached {
                phase = .loaded(await model(cached.value), offlineSince: cached.savedAt)
            } else {
                phase = .failed
            }
        }
    }

    /// The list row: the one passed in, else the first page's (a deep link or a push).
    private func model(_ result: SimResult) async -> SimResultModel {
        var found = row
        if found == nil {
            found = try? await service.list(page: 1, lane: nil).rows.first { $0.simId == id }
        }
        return SimResultModel(id: id, result: result, row: found, specs: specs)
    }

    private func postSeenOnce() async {
        guard !seenPosted else { return }
        seenPosted = true
        await markSeen(id)
    }
}
