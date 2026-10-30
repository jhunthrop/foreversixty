import SwiftUI

@main
struct ForeverSixtyApp: App {
    @State private var environment: AppEnvironment

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
    }
}
