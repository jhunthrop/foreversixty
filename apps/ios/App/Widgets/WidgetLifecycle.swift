import SwiftUI
import UIKit

/// The silent `widget.refresh` push (fixture surface.widget-refresh.apns). No server sends it
/// until lane A2's senders and Apple's push key exist; the handler is ready for that day.
final class AppDelegate: NSObject, UIApplicationDelegate {
    @MainActor static var onWidgetRefresh: (@MainActor () async -> Void)?

    func application(
        _ application: UIApplication,
        didReceiveRemoteNotification userInfo: [AnyHashable: Any]
    ) async -> UIBackgroundFetchResult {
        guard WidgetRefresher.isRefreshPush(userInfo) else { return .noData }
        await Self.onWidgetRefresh?()
        return .newData
    }
}

/// When the snapshot refreshes from inside the app: every foreground, a sign-up saved in the
/// app, a session change, a Live surfaces switch, and the silent push.
struct WidgetRefreshTriggers: ViewModifier {
    let environment: AppEnvironment
    @Environment(\.scenePhase) private var scenePhase

    func body(content: Content) -> some View {
        content
            .task(id: scenePhase) {
                if scenePhase == .active { await refresh() } else if scenePhase == .background { WidgetRefresher.schedule() }
            }
            .task(id: environment.session.phase.isSignedIn) { await refresh() }
            .task(id: environment.session.currentCharacter?.key) { await refresh() }
            .onReceive(NotificationCenter.default.publisher(for: .raidSignupSaved)) { _ in
                Task { await refresh() }
            }
            .task {
                AppDelegate.onWidgetRefresh = { await refresh() }
                LiveSurfaceHooks.onLiveAnswer = { @Sendable in
                    await MainActor.run { NotificationCenter.default.post(name: .raidSignupSaved, object: nil) }
                }
            }
    }

    @MainActor
    private func refresh() async {
        let session = environment.session
        await environment.widgets.refresh(
            signedIn: session.phase.isSignedIn, me: session.phase.me, current: session.currentCharacter?.key
        )
    }
}

extension View {
    func widgetRefreshTriggers(_ environment: AppEnvironment) -> some View {
        modifier(WidgetRefreshTriggers(environment: environment))
    }
}
