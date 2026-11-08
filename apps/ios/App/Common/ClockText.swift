import Foundation

/// A time of day as the app writes it everywhere: `6:02 pm`, lower-case, a plain space (the
/// system formats put a narrow no-break space before an upper-case PM).
enum ClockText {
    static func time(_ date: Date, calendar: Calendar = .current, locale: Locale = .current) -> String {
        let formatter = DateFormatter()
        formatter.calendar = calendar
        formatter.timeZone = calendar.timeZone
        formatter.locale = locale
        formatter.setLocalizedDateFormatFromTemplate("h:mm a")
        formatter.amSymbol = "am"
        formatter.pmSymbol = "pm"
        return formatter.string(from: date).replacingOccurrences(of: "\u{202F}", with: " ")
    }

    /// `3:00 pm` for `15:00`; the text itself when it is not a time.
    static func time(hhmm: String, locale: Locale = .current) -> String {
        let parts = hhmm.split(separator: ":").compactMap { Int($0) }
        var calendar = Calendar(identifier: .gregorian)
        calendar.timeZone = .gmt
        guard parts.count == 2, let date = calendar.date(from: DateComponents(hour: parts[0], minute: parts[1])) else {
            return hhmm
        }
        return time(date, calendar: calendar, locale: locale)
    }
}
