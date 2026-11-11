#if DEBUG
import Foundation
import Models

/// Debug-only: `-FSRefreshHold 2` holds the drawer's Battle.net refresh that many seconds, so the
/// review walk captures its `Refreshing…` frame (the fixture answers at once). Never in a release.
enum MenuDebugLaunch {
    static func holdRefresh() async {
        let arguments = ProcessInfo.processInfo.arguments
        guard let index = arguments.firstIndex(of: "-FSRefreshHold"), arguments.indices.contains(index + 1),
              let seconds = Double(arguments[index + 1])
        else { return }
        try? await Task.sleep(for: .seconds(seconds))
    }

    /// `-FSNotifications on` (or `off`) turns every group's push on (or off) in what the drawer reads,
    /// so the review walk captures `All on` and `All off` beside the fixture's some-off (round 7).
    static func notifications(_ settings: NotificationSettings?) -> NotificationSettings? {
        let arguments = ProcessInfo.processInfo.arguments
        guard let settings, let index = arguments.firstIndex(of: "-FSNotifications"), arguments.indices.contains(index + 1)
        else { return settings }
        let on = arguments[index + 1] == "on"
        let patch = NotificationsPatch(events: settings.events.map { NotificationChoice(kind: $0.kind, channel: "push", enabled: on) })
        return NotificationSettingsModel.applying(patch, to: settings)
    }
}
#endif
