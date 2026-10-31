import DesignSystem
import Routing
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

extension AppTab {
    /// The tab the app opens on: Home, or in debug builds the `-FSTab <tab>` launch argument, so
    /// a capture script can open a tab on the shared simulator without taps.
    static var launchTab: AppTab {
        #if DEBUG
            if let name = UserDefaults.standard.string(forKey: "FSTab"), let tab = AppTab(rawValue: name) { return tab }
        #endif
        return .home
    }
}

struct RootView: View {
    let environment: AppEnvironment
    @State private var selection: AppTab = .launchTab
    /// Lane F4: the Raids tab's stack, which raid links and notifications push onto.
    @State private var raidsRouter = RaidsRouter()
    /// Lane F5: Home's Sims stack, which sim links push onto.
    @State private var simsRouter = SimsRouter()
    /// Lane F16: the app's raid calendar, kept up to date from the Raids tab. It never asks for
    /// access until the player flips its switch.
    @State private var calendarSync: CalendarSyncController

    init(environment: AppEnvironment) {
        self.environment = environment
        let service = RaidsDependencies.live(environment).service
        let session = environment.session
        let loader = CalendarInputLoader(service: service, currentKey: { session.currentCharacter?.key })
        _calendarSync = State(initialValue: CalendarSyncController(
            adapter: EventKitCalendarAdapter(), store: UserDefaultsCalendarSyncStore(), now: AppClock.now, clock: RaidClock(),
            loadInput: { await loader.load() }
        ))
    }

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
        .environment(simsRouter)
        .task { await environment.appConfig.refresh() }
        .task { await environment.session.refresh() }
        .task { await environment.session.watchForDeadTokens() }
        .onOpenURL { url in DeepLink(url: url).map(openLink) }
        .task {
            RaidNotificationCenter.shared.onOpen = openRaids
            #if DEBUG
                // Review captures only: the launch arguments in Raids/Debug/RaidDebugLaunch.swift.
                if let link = RaidDebugLaunch.link { openRaids(link) }
                await RaidDebugLaunch.postRaidPushIfAsked()
            #endif
        }
    }

    /// Lane F1: the account states and the character selector.
    private var signInAvailable: Bool { environment.isAvailable(.signIn) }

    /// The tab slots. Each lane replaces its placeholder line: F2 Home, F4 Raids, F3 BiS, F1 You.
    @ViewBuilder
    private func screen(for tab: AppTab) -> some View {
        switch tab {
        case .home where signInAvailable:
            HomeAccountGate(session: environment.session, onPickClass: { selection = .bis }) {
                HomeTab(dependencies: .live(environment), session: environment.session, openBis: { selection = .bis })
            }
        case .home: HomeTab(dependencies: .live(environment), openBis: { selection = .bis })
        case .raids where environment.isAvailable(.raids):
            RaidsTab(
                dependencies: .live(environment), router: raidsRouter, session: signInAvailable ? environment.session : nil,
                extras: guildExtras, calendar: guildExtras == nil ? nil : calendarSync
            )
        case .raids: PlaceholderScreen(title: tab.title)
        case .bis where environment.isAvailable(.bis):
            BisScreen.live(environment, session: signInAvailable ? environment.session : nil)
        case .bis: PlaceholderScreen(title: tab.title)
        case .you where signInAvailable:
            YouScreen(
                session: environment.session,
                version: AccountCopy.version(environment.identity, data: environment.appConfig.current?.activeBuild),
                onPickClass: { selection = .bis },
                plansLink: environment.plansLink,
                calendar: guildExtras.map { ($0, calendarSync) }
            )
        case .you: PlaceholderScreen(title: tab.title)
        }
    }
}

extension RootView {
    /// Lane F16: the guild extras' dependencies, nil while `guild_extras` is off.
    private var guildExtras: GuildExtrasDependencies? {
        environment.isAvailable(.guildExtras) ? .live(environment) : nil
    }

    /// A link selects its tab: sim links Home (lane F5), raid links Raids (lane F4).
    private func openLink(_ link: DeepLink) {
        if SimsRouter.handles(link) {
            guard environment.isAvailable(.sims) else { return }
            selection = .home
            simsRouter.open(link)
        } else {
            openRaids(link)
        }
    }

    /// Lane F4: a raid, raid-night or guild link selects the Raids tab and pushes its screen.
    private func openRaids(_ link: DeepLink) {
        guard environment.isAvailable(.raids), RaidsRouter.handles(link) else { return }
        selection = .raids
        raidsRouter.open(link)
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
