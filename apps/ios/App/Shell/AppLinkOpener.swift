import Routing
import SwiftUI

/// Opens one of the app's own links on its tab (RootView.openLink), as OpenURLAction does for URLs.
struct OpenAppLinkAction {
    let handler: (DeepLink) -> Void

    func callAsFunction(_ link: DeepLink) { handler(link) }
}

extension EnvironmentValues {
    /// A sheet reads it to take the player from what it shows to the screen it names. Nil outside
    /// RootView (previews, tests).
    @Entry var openAppLink: OpenAppLinkAction?
}
