import Foundation
import Models
import Networking

/// The seven states of a Server runs row, as the site's `rowState` (web/src/lib/sim/server-runs.ts).
enum ServerRunState: Equatable, Sendable {
    case queued, running, runningOverdue, runningNoEstimate, done, doneUnseen, failed

    var isInFlight: Bool { self == .queued || self == .running || self == .runningOverdue || self == .runningNoEstimate }
    var isDone: Bool { self == .done || self == .doneUnseen }
}

/// A run's usual length in the one unit the state line names (`about 4 minutes`).
enum RunEstimate: Equatable, Sendable {
    case aboutAMinute
    case minutes(Int)
    case hours(Int)

    private static let secondsPerMinute = 60.0
    private static let minutesPerHour = 60.0
    /// From 90 minutes the estimate reads in hours, as on the site.
    private static let hoursFromMinutes = 90

    init(seconds: Int) {
        let minutes = Int((Double(seconds) / Self.secondsPerMinute).rounded())
        if minutes <= 1 {
            self = .aboutAMinute
        } else if minutes < Self.hoursFromMinutes {
            self = .minutes(minutes)
        } else {
            self = .hours(Int((Double(minutes) / Self.minutesPerHour).rounded()))
        }
    }
}

/// The rules under Sims (spec F5): pure data in, data out. Every sentence is in `SimsCopy`.
enum ServerRuns {
    /// Server runs shows the newest five; older ones are in Your sims (the site's `MAX_VISIBLE_RUNS`).
    static let maxVisible = 5
    /// A run is overdue once it has taken more than this many times its estimate.
    static let overdueFactor = 1.5
    /// How often an in-flight row asks `GET /v1/sims/{id}/progress` while the screen shows.
    static let pollInterval: Duration = .seconds(5)

    static func state(_ row: SimRow, now: Date) -> ServerRunState {
        switch row.state {
        case "queued": return .queued
        case "running":
            guard let estimate = row.estimateSeconds else { return .runningNoEstimate }
            return isOverdue(estimate: estimate, elapsed: elapsedSeconds(row, now: now)) ? .runningOverdue : .running
        case "error": return .failed
        default: return row.seen == false ? .doneUnseen : .done
        }
    }

    static func isOverdue(estimate: Int, elapsed: Int) -> Bool {
        Double(elapsed) > overdueFactor * Double(estimate)
    }

    /// Whole seconds since the run was created, never negative.
    static func elapsedSeconds(_ row: SimRow, now: Date) -> Int {
        guard let created = row.createdAt else { return 0 }
        return max(0, Int(now.timeIntervalSince(created)))
    }

    /// The newest five by creation time, and whether older rows were left out.
    static func visible(_ rows: [SimRow]) -> (rows: [SimRow], hasMore: Bool) {
        let newestFirst = rows.sorted { ($0.createdAt ?? .distantPast) > ($1.createdAt ?? .distantPast) }
        return (Array(newestFirst.prefix(maxVisible)), newestFirst.count > maxVisible)
    }

    /// The section shows for a Premium account, or one that still has server rows (a lapsed plan).
    static func showsSection(entitled: Bool, serverRows: [SimRow]) -> Bool {
        entitled || !serverRows.isEmpty
    }

    /// The rows the poll asks about: queued or running.
    static func inFlightIDs(_ rows: [SimRow], now: Date) -> [String] {
        rows.filter { state($0, now: now).isInFlight }.compactMap(\.simId)
    }

    /// The bar's fill, only when the run counts combinations (a bulk run); a plain run has no total.
    static func fraction(_ progress: SimProgress?) -> Double? {
        guard let progress, progress.combosTotal > 0 else { return nil }
        return min(1, Double(progress.combosDone) / Double(progress.combosTotal))
    }

    /// A finished poll answer: the list is read again for its headline and DPS.
    static func hasEnded(_ progress: SimProgress) -> Bool {
        progress.state == "done" || progress.state == "error"
    }
}

/// What a failed row's `Run again` shows (the site's `RerunStatus`).
enum RerunStatus: Equatable, Sendable {
    case idle, pending, started, failed, lapsed

    /// A refused rerun: 402 is a lapsed plan, anything else could not start.
    static func refused(_ error: APIError) -> RerunStatus {
        if case let .server(status, _, _) = error, status == 402 { return .lapsed }
        return .failed
    }

    /// The button cannot be pressed again while it starts or once it started.
    var isLocked: Bool { self == .pending || self == .started }
}
