import Foundation

/// When the shown answer was fetched, in the strip's words (spec §5.0): `9:41 pm` today,
/// `Thu 9:41 pm` this week, else `10 Dec, 9:41 pm`. Shared by every screen that draws a saved answer.
enum OfflineStamp {
    static let daysInWeek = 7

    static func format(_ date: Date, now: Date = .now, calendar: Calendar = .current) -> String {
        let formatter = DateFormatter()
        formatter.locale = Locale(identifier: "en_US_POSIX")
        formatter.calendar = calendar
        formatter.timeZone = calendar.timeZone
        formatter.amSymbol = "am"
        formatter.pmSymbol = "pm"
        let startOfToday = calendar.startOfDay(for: now)
        let days = calendar.dateComponents([.day], from: calendar.startOfDay(for: date), to: startOfToday).day ?? 0
        formatter.dateFormat = switch days {
        case 0: "h:mm a"
        case 1..<daysInWeek: "EEE h:mm a"
        default: "d MMM, h:mm a"
        }
        return formatter.string(from: date)
    }
}
