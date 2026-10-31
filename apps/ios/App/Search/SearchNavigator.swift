import DesignSystem
import Observation
import SwiftUI

/// A list the BiS tab is asked to open: a spec's list, or nil for the character's own.
struct BisRequest: Equatable, Sendable {
    let id = UUID()
    let selection: BisSelection?
}

/// F12's place in the app (spec §3.2): the cover over the tabs, and the BiS list a result asks
/// for. Nil in the environment while the `search` switch is off, so every surface behaves as
/// before: no button, F3 keeps its filter, F2.6 has no `More about this item`.
@MainActor
@Observable
final class SearchNavigator {
    let dependencies: SearchDependencies
    var isPresented = false
    /// RootView selects the BiS tab for it; the BiS screen opens the list and clears it.
    var bisRequest: BisRequest?

    init(dependencies: SearchDependencies) {
        self.dependencies = dependencies
    }

    func open() { isPresented = true }

    /// Leaves search (when it is up) for a list in the BiS tab.
    func openList(_ selection: BisSelection?) {
        isPresented = false
        bisRequest = BisRequest(selection: selection)
    }

    /// The request, once: the BiS screen takes it when it is ready to select.
    func takeBisRequest() -> BisRequest? {
        defer { bisRequest = nil }
        return bisRequest
    }
}

extension EnvironmentValues {
    @Entry var searchNavigator: SearchNavigator?
}

/// The magnifier at the trailing edge of a root bar (Home, Raids, BiS; spec §3.2).
private struct SearchButtonToolbar: ViewModifier {
    @Environment(\.searchNavigator) private var navigator

    func body(content: Content) -> some View {
        content.toolbar {
            if let navigator {
                ToolbarItem(placement: .topBarTrailing) {
                    Button { navigator.open() } label: {
                        Image(systemName: "magnifyingglass")
                    }
                    .accessibilityLabel(Copy.searchButton)
                }
            }
        }
    }
}

extension View {
    /// The search button, only while the `search` switch is on.
    func searchButton() -> some View { modifier(SearchButtonToolbar()) }
}
