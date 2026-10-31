import Auth
import Foundation

/// The two time phrases Home prints: the sync age (`2 h ago`) and the offline stamp
/// (`9:41 pm` today, `Thu 9:41 pm` this week, else `10 Dec, 9:41 pm`; spec §5.0).
struct HomeTime {
    let now: Date
    var calendar: Calendar = .current

    func age(since date: Date) -> String {
        RelativeAge(from: date, to: now).phrase
    }

    func stamp(_ date: Date) -> String {
        OfflineStamp.format(date, now: now, calendar: calendar)
    }
}
