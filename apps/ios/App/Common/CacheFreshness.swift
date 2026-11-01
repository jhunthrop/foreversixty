import Foundation
import Network
import Networking
import Synchronization

/// How long a cached answer stays fresh (spec §5.7 Offline mode): the Home answer and the raids
/// 5 minutes; rankings, character pages and reports 10 minutes; build data (guides, tiers, BiS,
/// talents, the search index) until the active build changes, so it never goes stale by age.
enum CacheLimit: Sendable {
    case homeAndRaids
    case pages
    case build

    var seconds: TimeInterval? {
        switch self {
        case .homeAndRaids: 5 * 60
        case .pages: 10 * 60
        case .build: nil
        }
    }
}

/// What a region shows after a failed load.
enum FailedLoad: Equatable, Sendable {
    /// No connection: the cache under `Offline. Showing what you had at {time}.`
    case offline(since: Date)
    /// Connected, but Forever Sixty did not answer and the cache is past its limit: the cache
    /// under `Forever Sixty did not answer. ...` with Retry.
    case stale(since: Date)
    /// No cache, a `4xx`, or a `5xx` on a cache still inside its limit: the region's own error
    /// state, as v1 draws it.
    case error
}

/// The strip decision (spec §5.7, the lead's ruling). Stale: a connection, the request failed or
/// timed out with no `4xx` (a transport failure, a timeout or a `5xx`), a cached answer older
/// than its limit. Offline: any other transport failure with a cached answer (v1). Otherwise the
/// region's error state. `connected` nil is unknown (tests and previews): v1's rule alone. Pure.
enum CacheFreshness {
    static func decide(_ error: APIError, connected: Bool?, cachedAt: Date?, limit: CacheLimit, now: Date) -> FailedLoad {
        guard let cachedAt, isUnanswered(error) else { return .error }
        if connected == true, let seconds = limit.seconds, now.timeIntervalSince(cachedAt) > seconds {
            return .stale(since: cachedAt)
        }
        if case .transport = error { return .offline(since: cachedAt) }
        return .error
    }

    /// The stale strip's time when `decide` says stale, else nil: the one check every store with
    /// a time-limited cache makes before its v1 offline path.
    static func staleSince(
        _ error: APIError, connectivity: (any Connectivity)?, savedAt: Date?, limit: CacheLimit, now: Date
    ) -> Date? {
        guard case let .stale(since) = decide(error, connected: connectivity?.isConnected, cachedAt: savedAt, limit: limit, now: now)
        else { return nil }
        return since
    }

    /// A failure the server did not answer for: transport (offline, timeout) or a `5xx`.
    static func isUnanswered(_ error: APIError) -> Bool {
        switch error {
        case .transport: true
        case let .server(status, _, _): status >= 500
        case let .unexpectedStatus(status): status >= 500
        default: false
        }
    }
}

/// Whether the phone has a usable connection, and whether it is unmetered (Wi-Fi only setting).
protocol Connectivity: Sendable {
    var isConnected: Bool { get }
    var isExpensive: Bool { get }
}

/// The system's path monitor, started once.
final class PathConnectivity: Connectivity {
    static let shared = PathConnectivity()

    private let monitor = NWPathMonitor()
    private let state = Mutex<(connected: Bool, expensive: Bool)>((true, false))

    private init() {
        monitor.pathUpdateHandler = { [weak self] path in
            self?.state.withLock { $0 = (path.status == .satisfied, path.isExpensive || path.isConstrained) }
        }
        monitor.start(queue: DispatchQueue(label: "gg.foreversixty.connectivity"))
    }

    var isConnected: Bool { state.withLock { $0.connected } }
    var isExpensive: Bool { state.withLock { $0.expensive } }
}

/// A fixed answer for tests and previews.
struct FixedConnectivity: Connectivity {
    var isConnected = true
    var isExpensive = false
}

/// A store with a time-limited cache that draws the stale strip (lane F15).
@MainActor
protocol StaleAware: AnyObject {
    var connectivity: (any Connectivity)? { get set }
}

extension StaleAware {
    /// The live store: the phone's connection decides offline or stale. Tests leave it nil.
    func connected(_ connectivity: (any Connectivity)? = PathConnectivity.shared) -> Self {
        self.connectivity = connectivity
        return self
    }
}

extension RankingsStore: StaleAware {}
extension CharacterStore: StaleAware {}
extension ReportsStore: StaleAware {}
extension ReportStore: StaleAware {}
extension SimsStore: StaleAware {}
extension SimResultStore: StaleAware {}
extension RosterStore: StaleAware {}
extension HistoryStore: StaleAware {}
