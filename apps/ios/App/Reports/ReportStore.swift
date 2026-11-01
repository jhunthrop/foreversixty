import Foundation
import Models
import Networking
import Observation

/// A loaded report and whether it shows from the cache (spec F10.2, F10.5).
struct ReportContent: Equatable, Sendable {
    let report: Report
    let offlineSince: Date?
    /// Lane F15: set while the cached answer shows because Forever Sixty did not answer.
    var staleSince: Date?

    var fightRows: [FightRow] { FightRows.rows(report) }
    var pullNumbers: [Int: PullNumber] { FightRows.pullNumbers(report) }
}

enum ReportPhase: Equatable, Sendable {
    case loading
    case loaded(ReportContent)
    /// 404: private, removed, or not this account's to see.
    case notYours
    case failed
}

/// The Your parses card: hidden signed out, else its own load and failure (spec F10.2).
enum ParsesState: Equatable, Sendable {
    case hidden, loading, loaded(ReportMe), failed
}

/// Loads one report (cache first), the viewer's parses, and the Premium coach card.
@MainActor
@Observable
final class ReportStore {
    private(set) var phase: ReportPhase = .loading
    private(set) var parses: ParsesState
    private(set) var coach: CoachState = .hidden

    let id: String
    private let account: ReportsAccount
    private let service: any ReportsService
    private let cache: any RaidCache
    private let now: @Sendable () -> Date
    /// Lane F15: the connection, for the stale strip; the screen sets it, nil keeps v1's rule.
    var connectivity: (any Connectivity)?

    static let notFound = 404

    init(id: String, account: ReportsAccount, service: any ReportsService, cache: any RaidCache, now: @escaping @Sendable () -> Date) {
        self.id = id
        self.account = account
        self.service = service
        self.cache = cache
        self.now = now
        parses = account.signedIn ? .loading : .hidden
    }

    private var reportKey: String { "report-\(id)" }
    private var meKey: String { "report-me-\(id)" }

    var report: Report? {
        if case let .loaded(content) = phase { return content.report }
        return nil
    }

    /// The report, then (signed in) the parses and the coach. A refresh keeps what is drawn.
    func load() async {
        await loadReport()
        guard account.signedIn, report != nil else { return }
        await loadParses()
        await loadCoach()
    }

    func loadReport() async {
        if case .loading = phase, let cached = cache.read(Report.self, key: reportKey) {
            phase = .loaded(ReportContent(report: cached.value, offlineSince: nil))
        }
        do {
            let report = try await service.report(id: id)
            cache.write(report, key: reportKey, at: now())
            phase = .loaded(ReportContent(report: report, offlineSince: nil))
        } catch {
            if case let .server(status, _, _) = error, status == Self.notFound {
                phase = .notYours
            } else if let cached = cache.read(Report.self, key: reportKey), let since = CacheFreshness.staleSince(
                error, connectivity: connectivity, savedAt: cached.savedAt, limit: .pages, now: now()
            ) {
                phase = .loaded(ReportContent(report: cached.value, offlineSince: nil, staleSince: since))
            } else if SignupRefusal.isOffline(error), let cached = cache.read(Report.self, key: reportKey) {
                phase = .loaded(ReportContent(report: cached.value, offlineSince: cached.savedAt))
            } else if report == nil {
                phase = .failed
            }
        }
    }

    func loadParses() async {
        guard account.signedIn else { return }
        do {
            let me = try await service.me(id: id)
            cache.write(me, key: meKey, at: now())
            parses = .loaded(me)
        } catch {
            if SignupRefusal.isOffline(error), let cached = cache.read(ReportMe.self, key: meKey) {
                parses = .loaded(cached.value)
            } else if case .loaded = parses {
                return
            } else {
                parses = .failed
            }
        }
    }

    /// Only with `log_coach`: the night line for the first character with a kill, then that
    /// line's pull review. A free account makes no coach call at all.
    func loadCoach() async {
        guard case let .loaded(me) = parses, let player = Coach.player(logCoach: account.logCoach, me: me) else {
            coach = .hidden
            return
        }
        do {
            guard let night = try await service.coachNight(id: id, player: player) else {
                coach = .hidden
                return
            }
            coach = Coach.state(try await service.fightCoach(id: id, n: night.fightIndex, player: player))
        } catch {
            coach = .failed
        }
    }

    /// The parses card's Try again.
    func retryParses() async {
        parses = .loading
        await loadParses()
        await loadCoach()
    }
}
