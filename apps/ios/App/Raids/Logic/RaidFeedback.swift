import Foundation
import UIKit
import UserNotifications

/// Whether notification permission is still undecided, and the system prompt.
protocol NotificationPermission: Sendable {
    func isUndetermined() async -> Bool
    func request() async
}

struct SystemNotificationPermission: NotificationPermission {
    func isUndetermined() async -> Bool {
        await UNUserNotificationCenter.current().notificationSettings().authorizationStatus == .notDetermined
    }

    func request() async {
        // Lane F7 registers the APNs token once permission is granted.
        _ = try? await UNUserNotificationCenter.current().requestAuthorization(options: [.alert, .sound, .badge])
    }
}

/// The reminder ask's rule (spec F4.4): after a successful In or Late, at most twice per device,
/// the second time at least 30 days after the first, never again after that.
struct ReminderAskPolicy: Sendable {
    static let maxAsks = 2
    static let gap: TimeInterval = 30 * 86400

    let shown: Int
    let lastShown: Date?

    func mayAsk(now: Date) -> Bool {
        guard shown < Self.maxAsks else { return false }
        guard let lastShown else { return true }
        return now.timeIntervalSince(lastShown) >= Self.gap
    }
}

/// Where the ask's count lives: UserDefaults per device.
protocol ReminderAskStore: Sendable {
    func policy() -> ReminderAskPolicy
    func recordShown(at date: Date)
}

struct UserDefaultsReminderAskStore: ReminderAskStore, @unchecked Sendable {
    let defaults: UserDefaults
    private static let countKey = "raids.reminderAsk.count"
    private static let lastKey = "raids.reminderAsk.last"

    init(defaults: UserDefaults = .standard) { self.defaults = defaults }

    func policy() -> ReminderAskPolicy {
        ReminderAskPolicy(shown: defaults.integer(forKey: Self.countKey), lastShown: defaults.object(forKey: Self.lastKey) as? Date)
    }

    func recordShown(at date: Date) {
        defaults.set(defaults.integer(forKey: Self.countKey) + 1, forKey: Self.countKey)
        defaults.set(date, forKey: Self.lastKey)
    }
}
