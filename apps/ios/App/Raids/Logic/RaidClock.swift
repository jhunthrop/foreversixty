import Foundation

/// The raid's times as the viewer reads them (spec 5.4 F4.1 and F4.2; the site's raid-copy.ts):
/// clocks, the date block, the long date, the planned row's `{when}`, the countdown and the
/// guild-time line. Everything takes the zone and the instant, so a test pins both.
struct RaidClock: Sendable {
    let zone: TimeZone
    let locale: Locale

    init(zone: TimeZone = .current, locale: Locale = Locale(identifier: "en_GB")) {
        self.zone = zone
        self.locale = locale
    }

    private static let minute: TimeInterval = 60
    private static let hour: TimeInterval = 3600
    private static let day: TimeInterval = 86400
    /// A raid within this many calendar days reads by weekday (`Thursday 8:00 pm`).
    private static let weekdayWindowDays = 6

    private func format(_ template: String, _ date: Date, zone: TimeZone? = nil) -> String {
        let formatter = DateFormatter()
        // The zone's abbreviation (CST) comes from the US data; en_GB spells it GMT-6.
        formatter.locale = template.contains("z") ? Locale(identifier: "en_US") : locale
        formatter.timeZone = zone ?? self.zone
        formatter.dateFormat = template
        formatter.amSymbol = "am"
        formatter.pmSymbol = "pm"
        return formatter.string(from: date)
    }

    /// `7:45 pm`.
    func clock(_ date: Date) -> String { format("h:mm a", date) }

    /// `8:00 pm CST` in the guild's zone.
    func guildClock(_ date: Date, guildZone: TimeZone) -> String {
        format("h:mm a zzz", date, zone: guildZone)
    }

    /// The date block: `THU`, `10`, `DEC`.
    func dateBlock(_ date: Date) -> (weekday: String, day: String, month: String) {
        (format("EEE", date).uppercased(), format("d", date), format("MMM", date).uppercased())
    }

    /// `Thursday 10 December`.
    func longDate(_ date: Date) -> String { format("EEEE d MMMM", date) }

    /// `Thu 10 Dec` (the spec ask's date and the copied list).
    func shortDate(_ date: Date) -> String { format("EEE d MMM", date) }

    /// The night screen's meta date: weekday and `d Month` separately.
    func weekday(_ date: Date) -> String { format("EEEE", date) }
    func dayMonth(_ date: Date) -> String { format("d MMMM", date) }

    /// The planned row's `{when}`: `Tonight 8:00 pm`, `Tomorrow 8:00 pm`, `Saturday 7:00 pm`,
    /// `Thu 17 Dec, 8:00 pm`.
    func when(_ date: Date, now: Date) -> String {
        let time = clock(date)
        switch calendarDays(from: now, to: date) {
        case 0: return Copy.raidsWhenTonight(clock: time)
        case 1: return Copy.raidsWhenTomorrow(clock: time)
        case 2...Self.weekdayWindowDays: return Copy.raidsWhenWeekday(weekday: weekday(date), clock: time)
        default: return Copy.raidsWhenDate(date: shortDate(date), clock: time)
        }
    }

    /// True when the viewer's zone and the guild's differ at that instant.
    func offsetDiffers(at date: Date, guildZone: TimeZone) -> Bool {
        zone.secondsFromGMT(for: date) != guildZone.secondsFromGMT(for: date)
    }

    private func calendarDays(from: Date, to: Date) -> Int {
        var calendar = Calendar(identifier: .gregorian)
        calendar.timeZone = zone
        let start = calendar.startOfDay(for: from)
        let end = calendar.startOfDay(for: to)
        return calendar.dateComponents([.day], from: start, to: end).day ?? 0
    }

    /// The site's countdown (`raid-copy.ts countdown`): `in 2 days 2 h`, `starting now`,
    /// `started 20 min ago`, `ended 1 h ago`.
    func countdown(to start: Date, end: Date, now: Date) -> String {
        let untilStart = start.timeIntervalSince(now)
        if abs(untilStart) < Self.minute { return Copy.raidCountdownNow }
        if untilStart > 0 { return Copy.raidCountdownIn(span: Self.span(untilStart)) }
        if now < end { return Copy.raidCountdownStarted(span: Self.span(-untilStart)) }
        return Copy.raidCountdownEnded(span: Self.span(now.timeIntervalSince(end)))
    }

    /// `2 days 2 h`, `5 h 20 min`, `12 min`.
    static func span(_ seconds: TimeInterval) -> String {
        let (days, hours, minutes) = parts(seconds)
        if days > 0 { return [Copy.raidsSpanDays(n: days), hours > 0 ? Copy.raidsSpanHours(n: hours) : nil].compactMap(\.self).joined(separator: " ") }
        if hours > 0 { return [Copy.raidsSpanHours(n: hours), minutes > 0 ? Copy.raidsSpanMinutes(n: minutes) : nil].compactMap(\.self).joined(separator: " ") }
        return Copy.raidsSpanMinutes(n: minutes)
    }

    /// The same span in whole words, for VoiceOver (7.1: `2 days 3 hours`, never `h`).
    static func spanWords(_ seconds: TimeInterval) -> String {
        let (days, hours, minutes) = parts(seconds)
        if days > 0 { return [Copy.raidsSpanWordsDays(n: days), hours > 0 ? Copy.raidsSpanWordsHours(n: hours) : nil].compactMap(\.self).joined(separator: " ") }
        if hours > 0 { return [Copy.raidsSpanWordsHours(n: hours), minutes > 0 ? Copy.raidsSpanWordsMinutes(n: minutes) : nil].compactMap(\.self).joined(separator: " ") }
        return Copy.raidsSpanWordsMinutes(n: minutes)
    }

    private static func parts(_ seconds: TimeInterval) -> (Int, Int, Int) {
        let total = max(0, seconds)
        return (Int(total / day), Int(total.truncatingRemainder(dividingBy: day) / hour), Int(total.truncatingRemainder(dividingBy: hour) / minute))
    }

    /// The offline strip's `{time}`: `9:41 pm` today, `Thu 9:41 pm` this week, else `10 Dec, 9:41 pm`.
    func stripTime(_ date: Date, now: Date) -> String {
        var calendar = Calendar(identifier: .gregorian)
        calendar.timeZone = zone
        return OfflineStamp.format(date, now: now, calendar: calendar)
    }
}
