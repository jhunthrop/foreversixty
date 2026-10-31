import Foundation
import Models
import Networking
import Observation

/// What Sims draws (spec F5.1): the server rows, your sims, and the offline stamp.
struct SimsContent: Equatable, Sendable {
    let serverRows: [SimRow]
    let mine: [SimRow]
    /// `GET /v1/sims`'s total, so the list knows whether another page exists.
    let total: Int
    let page: Int
    /// Set while the cached answer shows because a refresh could not connect.
    let offlineSince: Date?
    /// A section that did not load says so on its own; the other keeps its rows (spec F5 States).
    var serverFailed = false
    var mineFailed = false

    var hasMorePages: Bool { mine.count < total }
}

enum SimsPhase: Equatable, Sendable {
    case loading
    case loaded(SimsContent)
    case failed
}

/// Loads both lists (cache first), polls in-flight runs, and runs a failed run again.
@MainActor
@Observable
final class SimsStore {
    private(set) var phase: SimsPhase = .loading
    /// The newest poll answer per in-flight run.
    private(set) var progress: [String: SimProgress] = [:]
    /// What each failed row's `Run again` shows.
    private(set) var rerun: [String: RerunStatus] = [:]
    /// The last next-page read failed: the list ends with the error and `Try again`.
    private(set) var pageFailed = false

    private let service: any SimsService
    private let cache: any RaidCache
    private let now: @Sendable () -> Date

    static let cacheKey = "sims"

    init(service: any SimsService, cache: any RaidCache, now: @escaping @Sendable () -> Date, initial: SimsPhase = .loading) {
        self.service = service
        self.cache = cache
        self.now = now
        self.phase = initial
    }

    private var content: SimsContent? {
        if case let .loaded(content) = phase { return content }
        return nil
    }

    func load() async {
        if case .loading = phase, let cached = cache.read(CachedSims.self, key: Self.cacheKey) {
            phase = .loaded(cached.value.content(offlineSince: nil))
        }
        let server = await answer { () async throws(APIError) in try await service.list(page: 1, lane: SimLane.server) }
        let mine = await answer { () async throws(APIError) in try await service.list(page: 1, lane: nil) }
        switch (server, mine) {
        case let (.success(server), .success(mine)):
            let answer = CachedSims(server: server.rows, mine: mine)
            cache.write(answer, key: Self.cacheKey, at: now())
            phase = .loaded(answer.content(offlineSince: nil))
        case let (.failure(error), _), let (_, .failure(error)):
            fail(error, server: try? server.get(), mine: try? mine.get())
        }
    }

    private func answer(_ call: () async throws(APIError) -> SimList) async -> Result<SimList, APIError> {
        do { return .success(try await call()) } catch { return .failure(error) }
    }

    private func fail(_ error: APIError, server: SimList?, mine: SimList?) {
        let cached = cache.read(CachedSims.self, key: Self.cacheKey)
        if SignupRefusal.isOffline(error), let cached {
            phase = .loaded(cached.value.content(offlineSince: cached.savedAt))
        } else if server != nil || mine != nil {
            phase = .loaded(SimsContent(
                serverRows: server?.rows ?? [], mine: mine?.rows ?? [], total: mine?.total ?? 0, page: mine?.page ?? 1,
                offlineSince: nil, serverFailed: server == nil, mineFailed: mine == nil
            ))
        } else if content == nil {
            // A refresh that fails keeps what is drawn; only a first load shows the error.
            phase = .failed
        }
    }

    /// The next page of Your sims, when the list has one.
    func loadMore() async {
        guard let content, content.offlineSince == nil, content.hasMorePages else { return }
        guard let next = try? await service.list(page: content.page + 1, lane: nil) else {
            pageFailed = true
            return
        }
        pageFailed = false
        phase = .loaded(SimsContent(
            serverRows: content.serverRows, mine: content.mine + next.rows,
            total: next.total, page: next.page, offlineSince: nil, serverFailed: content.serverFailed
        ))
    }

    /// One poll round: each in-flight run's progress; a run that ended reloads the lists.
    func pollOnce() async {
        guard let content else { return }
        var ended = false
        for id in ServerRuns.inFlightIDs(content.serverRows, now: now()) {
            guard let answer = try? await service.progress(id: id) else { continue }
            progress[id] = answer
            ended = ended || ServerRuns.hasEnded(answer)
        }
        if ended { await load() }
    }

    /// Whether anything is still running, so the screen keeps polling.
    var isPolling: Bool {
        guard let content else { return false }
        return !ServerRuns.inFlightIDs(content.serverRows, now: now()).isEmpty
    }

    /// `Run again` on a failed row: Starting…, then Started and the new run in Server runs.
    func runAgain(_ id: String) async {
        guard !(rerun[id]?.isLocked ?? false) else { return }
        rerun[id] = .pending
        do {
            _ = try await service.rerun(id: id)
            rerun[id] = .started
            await load()
        } catch {
            rerun[id] = RerunStatus.refused(error)
        }
    }

    /// Opening a finished result marks it seen, here and on the server.
    func markSeen(_ id: String) async {
        try? await service.markSeen(id: id)
        guard let content else { return }
        phase = .loaded(SimsContent(
            serverRows: content.serverRows.map { $0.simId == id ? $0.seen(true) : $0 },
            mine: content.mine.map { $0.simId == id ? $0.seen(true) : $0 },
            total: content.total, page: content.page, offlineSince: content.offlineSince,
            serverFailed: content.serverFailed, mineFailed: content.mineFailed
        ))
    }
}

/// The two first pages as cached together.
struct CachedSims: Codable, Sendable {
    let server: [SimRow]
    let mine: SimList

    func content(offlineSince: Date?) -> SimsContent {
        SimsContent(serverRows: server, mine: mine.rows, total: mine.total, page: mine.page, offlineSince: offlineSince)
    }
}

extension SimRow {
    func seen(_ seen: Bool) -> SimRow {
        SimRow(
            simId: simId, spec: spec, kind: kind, dps: dps, headline: headline, state: state,
            engineVersion: engineVersion, createdAt: createdAt, title: title, lane: lane,
            finishedAt: finishedAt, failureReason: failureReason, seen: seen, estimateSeconds: estimateSeconds
        )
    }
}
