import DesignSystem
import SwiftUI

/// The app's four tabs, in the spec's order (§3.1): Home, Raids, BiS, You. Feature lanes plug
/// their root screen into `RootView.screen(for:)`, one case each; the enum itself stays fixed.
enum AppTab: String, CaseIterable, Identifiable {
    case home, raids, bis, you

    var id: String { rawValue }

    var title: String {
        switch self {
        case .home: Copy.tabHome
        case .raids: Copy.tabRaids
        case .bis: Copy.tabBis
        case .you: Copy.tabYou
        }
    }

    var systemImage: String {
        switch self {
        case .home: "house"
        case .raids: "calendar"
        case .bis: "shield"
        case .you: "person"
        }
    }
}

struct RootView: View {
    let environment: AppEnvironment
    @State private var selection: AppTab = .home

    var body: some View {
        if environment.updateGate.isRequired {
            UpdateRequiredView(storeURL: environment.updateGate.storeURL)
        } else {
            tabs
        }
    }

    private var tabs: some View {
        TabView(selection: $selection) {
            ForEach(AppTab.allCases) { tab in
                Tab(tab.title, systemImage: tab.systemImage, value: tab) {
                    screen(for: tab)
                }
            }
        }
        .tint(Theme.Palette.accent)
        .task { await environment.appConfig.refresh() }
    }

    /// The tab slots. Each lane replaces its placeholder line: F2 Home, F4 Raids, F3 BiS, F1 You.
    @ViewBuilder
    private func screen(for tab: AppTab) -> some View {
        switch tab {
        case .home: HomeTab(dependencies: .live(environment), openBis: { selection = .bis })
        case .raids: PlaceholderScreen(title: tab.title)
        case .bis: PlaceholderScreen(title: tab.title)
        case .you: PlaceholderScreen(title: tab.title)
        }
    }
}

struct PlaceholderScreen: View {
    let title: String

    var body: some View {
        NavigationStack {
            Text(Copy.shellComingSoon)
                .foregroundStyle(Theme.Palette.muted)
                .frame(maxWidth: .infinity, maxHeight: .infinity)
                .background(Theme.Palette.background)
                .navigationTitle(title)
        }
    }
}
