import Foundation

/// When a raid widget redraws without the network (spec §5.6 Refresh): every 30 minutes up to
/// the raid, hourly while it is more than two days away, plus the raid's start and end, the end
/// of a `Not saved` line and the moment the snapshot turns stale.
enum WidgetTimeline {
    static let step: TimeInterval = 30 * 60
    static let farStep: TimeInterval = 60 * 60
    static let farAfter: TimeInterval = 2 * 24 * 60 * 60
    /// A timeline never plans further than this; the app's refresh and the next reload follow.
    static let horizon: TimeInterval = 3 * 24 * 60 * 60

    static func dates(now: Date, snapshot: WidgetSnapshot?) -> [Date] {
        var dates: Set<Date> = [now]
        if let raid = snapshot?.widget.nextRaid, raid.status != "cancelled" {
            dates.formUnion(countdownSteps(now: now, to: raid.startsAt))
            dates.insert(raid.startsAt)
            dates.insert(raid.endsAt)
        }
        if let until = snapshot?.notSavedUntil, until > now { dates.insert(until) }
        if let fetched = snapshot?.fetchedAt {
            let stale = fetched.addingTimeInterval(TimeInterval(WidgetFigures.staleAfterHours) * 3600)
            if stale > now { dates.insert(stale) }
        }
        let end = now.addingTimeInterval(horizon)
        return dates.filter { $0 >= now && $0 <= end }.sorted()
    }

    /// Steps counted back from the raid, so each entry lands on a whole countdown value.
    private static func countdownSteps(now: Date, to start: Date) -> [Date] {
        var out: [Date] = []
        var at = start
        while at > now {
            out.append(at)
            at = at.addingTimeInterval(-(start.timeIntervalSince(at) >= farAfter ? farStep : step))
            if out.count > 200 { break }
        }
        return out
    }
}
