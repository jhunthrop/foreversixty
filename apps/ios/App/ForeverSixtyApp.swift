import DesignSystem
import SwiftUI

@main
struct ForeverSixtyApp: App {
    @State private var environment: AppEnvironment
    /// Lane F14: the silent widget.refresh push arrives through the app delegate.
    @UIApplicationDelegateAdaptor(AppDelegate.self) private var appDelegate

    init() {
        let environment = AppEnvironment.fromBundle()
        // Round 5's bars: the painted band behind every navigation bar, the owner's tab bar.
        BarAppearance.install()
        CrashReporting.start(configuration: environment.configuration)
        // Lane F4: the raid notifications' categories and their In, Late and Out answers.
        let flows = LiveRaidFlowService(client: environment.client)
        let ping = LiveRaidPingService(client: environment.client)
        RaidNotificationCenter.shared.install(
            answerer: RaidPushAnswerer(calls: LiveRaidService(client: environment.client)),
            officer: OfficerPushActions(
                accept: { guild, application in (try? await flows.accept(guild: guild, application: application)) != nil },
                decline: { guild, application in (try? await flows.decline(guild: guild, application: application)) != nil },
                mention: { guild, raid, role in try? await ping.ping(guild: guild, raid: raid, role: role) }
            )
        )
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
