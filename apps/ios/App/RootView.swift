import DesignSystem
import SwiftUI

/// The app's tabs, one per v1 feature area in the lane plan. Lane D's spec
/// sets the final set, names and icons; feature lanes replace each placeholder.
enum AppTab: String, CaseIterable, Identifiable {
    case home, raids, bis, sims, you

    var id: String { rawValue }

    var title: String {
        switch self {
        case .home: Copy.shellTabCharacter
        case .raids: Copy.tabRaids
        case .bis: Copy.tabBis
        case .sims: Copy.shellTabSims
        case .you: Copy.shellTabAccount
        }
    }

    var systemImage: String {
        switch self {
        case .home: "person.crop.circle"
        case .raids: "calendar"
        case .bis: "list.star"
        case .sims: "chart.bar"
        case .you: "gearshape"
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
        .environment(environment.session)
        .task { await environment.appConfig.refresh() }
        .task { await environment.session.refresh() }
        .task { await environment.session.watchForDeadTokens() }
    }

    /// Lane F1: the account states, on by default in debug and fixture builds.
    private var signInAvailable: Bool {
        environment.appConfig.isAvailable(.signIn, mode: environment.configuration.mode, isDebugBuild: Self.isDebugBuild)
    }

    #if DEBUG
    private static let isDebugBuild = true
    #else
    private static let isDebugBuild = false
    #endif

    @ViewBuilder
    private func screen(for tab: AppTab) -> some View {
        switch tab {
        case .home where signInAvailable:
            HomeAccountGate(session: environment.session, onPickClass: { selection = .bis }) {
                HomeScreen(client: environment.client)
            }
        case .home: HomeScreen(client: environment.client)
        case .you where signInAvailable:
            YouScreen(
                session: environment.session,
                version: AccountCopy.version(environment.identity, data: environment.appConfig.current?.activeBuild),
                onPickClass: { selection = .bis }
            )
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
