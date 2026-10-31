import SwiftUI

@main
struct ForeverSixtyApp: App {
    @State private var environment: AppEnvironment
    /// Lane F14: the silent widget.refresh push arrives through the app delegate.
    @UIApplicationDelegateAdaptor(AppDelegate.self) private var appDelegate

    init() {
        let environment = AppEnvironment.fromBundle()
        CrashReporting.start(configuration: environment.configuration)
        // Lane F4: the raid notifications' categories and their In, Late and Out answers.
        RaidNotificationCenter.shared.install(answerer: RaidPushAnswerer(service: LiveRaidService(client: environment.client)))
        _environment = State(initialValue: environment)
    }

    var body: some Scene {
        WindowGroup {
            RootView(environment: environment)
                .preferredColorScheme(.dark)
        }
        // Lane F14: about every 30 minutes the system lets the snapshot refresh in the background.
        .backgroundTask(.appRefresh(WidgetRefresher.taskID)) {
            WidgetRefresher.schedule()
            await refreshWidgets()
        }
    }

    @MainActor
    private func refreshWidgets() async {
        let session = environment.session
        await environment.widgets.refresh(
            signedIn: session.phase.isSignedIn, me: session.phase.me, current: session.currentCharacter?.key
        )
    }
}
