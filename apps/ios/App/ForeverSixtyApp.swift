import SwiftUI

@main
struct ForeverSixtyApp: App {
    @State private var environment = AppEnvironment.fromBundle()

    var body: some Scene {
        WindowGroup {
            RootView(environment: environment)
                .preferredColorScheme(.dark)
        }
    }
}
