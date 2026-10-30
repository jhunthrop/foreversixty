import SwiftUI

@main
struct ForeverSixtyApp: App {
    @State private var environment: AppEnvironment

    init() {
        let environment = AppEnvironment.fromBundle()
        CrashReporting.start(configuration: environment.configuration)
        _environment = State(initialValue: environment)
    }

    var body: some Scene {
        WindowGroup {
            RootView(environment: environment)
                .preferredColorScheme(.dark)
        }
    }
}
