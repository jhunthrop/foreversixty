import DesignSystem
import SwiftUI

/// The app's tabs, one per v1 feature area in the lane plan. Lane D's spec
/// sets the final set, names and icons; feature lanes replace each placeholder.
enum AppTab: String, CaseIterable, Identifiable {
    case home, raids, bis, sims, account

    var id: String { rawValue }

    var title: String {
        switch self {
        case .home: Copy.shellTabCharacter
        case .raids: Copy.tabRaids
        case .bis: Copy.tabBis
        case .sims: Copy.shellTabSims
        case .account: Copy.shellTabAccount
        }
    }

    var systemImage: String {
        switch self {
        case .home: "person.crop.circle"
        case .raids: "calendar"
        case .bis: "list.star"
        case .sims: "chart.bar"
        case .account: "gearshape"
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

    @ViewBuilder
    private func screen(for tab: AppTab) -> some View {
        switch tab {
        // Lane F2: the home answer.
        case .home: HomeTab(dependencies: .live(environment), openBis: { selection = .bis })
        default: PlaceholderScreen(title: tab.title)
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
