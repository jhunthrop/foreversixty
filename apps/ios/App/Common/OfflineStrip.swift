import DesignSystem
import SwiftUI

/// When the shown answer was fetched, in the strip's words (spec §5.0): `9:41 pm` today,
/// `Thu 9:41 pm` this week, else `10 Dec, 9:41 pm`.
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

/// The offline strip (spec §5.0): 36 tall under the bar, a wifi-off glyph and the stamp.
/// Shared by every screen that draws a saved answer; announced once when it appears.
struct OfflineStrip: View {
    let savedAt: Date
    var now: Date = .now

    static let height: CGFloat = 36

    var body: some View {
        let text = Copy.offlineStrip(time: OfflineStamp.format(savedAt, now: now))
        HStack(spacing: Theme.Spacing.sm) {
            Image(systemName: "wifi.slash")
                .font(.system(size: 14, weight: .semibold))
                .foregroundStyle(Theme.Palette.nav)
                .accessibilityHidden(true)
            Text(text)
                .font(Theme.Typography.metaRole)
                .foregroundStyle(Theme.Palette.nav)
            Spacer(minLength: 0)
        }
        .padding(.horizontal, Theme.Spacing.md)
        .frame(minHeight: Self.height)
        .background(Theme.Palette.offlineGround)
        .overlay(alignment: .bottom) { Rectangle().fill(Theme.Palette.line).frame(height: 1) }
        .accessibilityElement(children: .combine)
        .onAppear { AccessibilityNotification.Announcement(text).post() }
    }
}
