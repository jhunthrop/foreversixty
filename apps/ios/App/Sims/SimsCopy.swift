import Auth
import Foundation
import Models

/// Every F5 sentence built from data (spec §5.5); the words themselves live in the copy catalog.
enum SimsCopy {
    static func estimate(_ estimate: RunEstimate) -> String {
        switch estimate {
        case .aboutAMinute: Copy.simsEstimateMinute
        case let .minutes(n): Copy.simsEstimateMinutes(n: n)
        case let .hours(n): Copy.simsEstimateHours(n: n)
        }
    }

    /// The muted words after `Running`: the estimate, `taking longer than usual`, or nothing.
    static func stateSuffix(_ row: SimRow, state: ServerRunState, now: Date) -> String {
        switch state {
        case .running: row.estimateSeconds.map { Copy.simsTakes(estimate: estimate(RunEstimate(seconds: $0))) } ?? ""
        case .runningOverdue: Copy.simsTakesLong
        default: ""
        }
    }

    static func failedLine(_ row: SimRow) -> String {
        guard let reason = row.failureReason, !reason.isEmpty else { return Copy.simsFailedPlain }
        return Copy.simsFailed(reason: reason)
    }

    static func finishedLine(_ row: SimRow, now: Date) -> String {
        Copy.simsFinished(age: age(row.finishedAt ?? row.createdAt, now: now))
    }

    /// Relative inside 14 days (`2 h ago`), else a short date (`2 Nov`), as the site's `ageWords`.
    static func age(_ date: Date?, now: Date) -> String {
        guard let date else { return RelativeAge.now.phrase }
        return now.timeIntervalSince(date) < relativeCeiling ? RelativeAge(from: date, to: now).phrase : shortDate(date)
    }

    private static let relativeCeiling: TimeInterval = 14 * 24 * 60 * 60

    /// `Saved from the browser · 26 Oct` or `Ran on the server · 26 Oct`.
    static func origin(_ row: SimRow) -> String {
        let date = row.createdAt.map(shortDate) ?? ""
        return row.lane == "server" ? Copy.simsRanServer(date: date) : Copy.simsSavedBrowser(date: date)
    }

    /// `14 Dec`, the spec's day-month order on every locale (as Android).
    static func shortDate(_ date: Date) -> String {
        let formatter = DateFormatter()
        formatter.locale = Locale(identifier: "en_US_POSIX")
        formatter.dateFormat = "d MMM"
        return formatter.string(from: date)
    }
}
