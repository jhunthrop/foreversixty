import Foundation

/// The app's clock. Debug builds take `-FSNow <ISO 8601 instant>` so the review captures run on
/// the same frozen instant as the snapshot goldens; every other build reads the real time.
enum AppClock {
    static let now: @Sendable () -> Date = {
        #if DEBUG
            if let text = UserDefaults.standard.string(forKey: "FSNow"),
               let pinned = try? Date.ISO8601FormatStyle().parse(text) {
                return { pinned }
            }
        #endif
        return { Date() }
    }()
}
