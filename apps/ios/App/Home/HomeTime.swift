import Foundation

/// The two time phrases Home prints: the sync age (`2 h ago`) and the offline stamp
/// (`9:41 pm` today, `Thu 9:41 pm` this week, else `10 Dec, 9:41 pm`; spec §5.0).
struct HomeTime {
    let now: Date
    var calendar: Calendar = .current

    func age(since date: Date) -> String {
        let minutes = Int(now.timeIntervalSince(date) / 60)
        switch minutes {
        case ..<1: return Copy.homeAgeNow
        case ..<60: return Copy.homeAgeMinutes(n: minutes)
        case ..<(24 * 60): return Copy.homeAgeHours(n: minutes / 60)
        default: return Copy.homeAgeDays(n: minutes / (24 * 60))
        }
    }

    func stamp(_ date: Date) -> String {
        let time = format(date, "h:mm a").replacingOccurrences(of: "AM", with: "am").replacingOccurrences(of: "PM", with: "pm")
        if calendar.isDate(date, inSameDayAs: now) { return time }
        if let days = calendar.dateComponents([.day], from: calendar.startOfDay(for: date), to: calendar.startOfDay(for: now)).day,
           days < 7 {
            return "\(format(date, "EEE")) \(time)"
        }
        return "\(format(date, "d MMM")), \(time)"
    }

    private func format(_ date: Date, _ pattern: String) -> String {
        let formatter = DateFormatter()
        formatter.locale = Locale(identifier: "en_US_POSIX")
        formatter.calendar = calendar
        formatter.timeZone = calendar.timeZone
        formatter.dateFormat = pattern
        formatter.amSymbol = "AM"
        formatter.pmSymbol = "PM"
        return formatter.string(from: date)
    }
}
