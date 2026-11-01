import DesignSystem
import Models
import Routing
import SwiftUI

/// The app's tabs, in the spec's order (Wave 3 §3.1): Home, Raids, BiS, Guides, You. Guides shows
/// only with its `guides` switch. Feature lanes plug their root screen into `RootView.screen(for:)`.
enum AppTab: String, CaseIterable, Identifiable {
    case home, raids, bis, guides, you

    var id: String { rawValue }

    var title: String {
        switch self {
        case .home: Copy.tabHome
        case .raids: Copy.tabRaids
        case .bis: Copy.tabBis
        case .guides: Copy.tabGuides
        case .you: Copy.tabYou
        }
    }

    var systemImage: String {
        switch self {
        case .home: "house"
        case .raids: "calendar"
        case .bis: "shield"
        case .guides: "book"
        case .you: "person"
        }
    }

    /// The tabs a build shows: Guides only while its switch is on.
    static func visible(guides: Bool) -> [AppTab] {
        allCases.filter { $0 != .guides || guides }
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
    /// Lane F12: the search cover and the BiS list a result asks for.
    @State private var search: SearchNavigator
    /// Lane F10: Home's Your reports stack, which report links push onto.
    @State private var reportsRouter = ReportsRouter()
    /// Lane F11: Home's Rankings and character pages, which rankings and player links push onto.
    @State private var rankingsRouter = RankingsRouter()
    /// Lane F9: the Guides tab's stack, which guide links push onto.
    @State private var guidesRouter = GuidesRouter()
    /// Lane F9: lists a guide's Gear row and `/app/bis` links ask the BiS tab for.
    @State private var bisRouter = BisRouter()
    /// Lane F9: keeps the guides of the player's classes saved.
    @State private var guidesPrefetcher: GuidesPrefetcher?
    @Environment(\.scenePhase) private var scenePhase
    /// Lane F15: where an intent or a Spotlight tap asks the app to go.
    @State private var intentRouter = IntentRouter.shared
    /// Lane F16: the app's raid calendar, kept up to date from the Raids tab. It never asks for
    /// access until the player flips its switch.
    @State private var calendarSync: CalendarSyncController

    init(environment: AppEnvironment) {
        self.environment = environment
        _search = State(initialValue: SearchNavigator(dependencies: .live(environment)))
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
            ForEach(AppTab.visible(guides: environment.isAvailable(.guides))) { tab in
                Tab(tab.title, systemImage: tab.systemImage, value: tab) {
                    screen(for: tab)
                }
            }
        }
        .tint(Theme.Palette.accent)
        .environment(environment.session)
        .environment(simsRouter)
        .environment(\.searchNavigator, searchIsOn ? search : nil)
        .fullScreenCover(isPresented: $search.isPresented) {
            SearchScreen(navigator: search).tint(Theme.Palette.accent)
        }
        .onChange(of: search.bisRequest) { _, request in
            if request != nil { selection = .bis }
        }
        .environment(reportsRouter)
        .environment(rankingsRouter)
        .task { await environment.appConfig.refresh() }
        .task { await environment.session.refresh() }
        .task { await environment.session.watchForDeadTokens() }
        .widgetRefreshTriggers(environment)
        .task(id: GuidesPrefetchTrigger(active: scenePhase == .active, classes: playerClasses)) { await prefetchGuides() }
        .onOpenURL { url in DeepLink(url: url).map(openLink) }
        // Lane F15: Siri, Shortcuts and Spotlight taps open what they name.
        .onChange(of: intentRouter.pending) { _, destination in
            guard let destination else { return }
            intentRouter.pending = nil
            openIntent(destination)
        }
        .task(id: SpotlightTrigger(
            active: scenePhase == .active, account: environment.session.phase.me,
            current: environment.session.currentCharacter?.key, enabled: environment.isAvailable(.system)
        )) { await refreshSpotlight() }
        .task {
            RaidNotificationCenter.shared.onOpen = openRaids
            #if DEBUG
                // Review captures only: the launch arguments in Raids/Debug/RaidDebugLaunch.swift.
                if let link = RaidDebugLaunch.link { openRaids(link) }
                await RaidDebugLaunch.postRaidPushIfAsked()
            #endif
        }
    }

    /// Lane F12: universal search, its button and the item page.
    private var searchIsOn: Bool { environment.isAvailable(.search) }

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
            BisScreen.live(environment, session: signInAvailable ? environment.session : nil, router: bisRouter)
        case .bis: PlaceholderScreen(title: tab.title)
        case .guides:
            GuidesScreen(
                dependencies: .live(environment, session: signInAvailable ? environment.session : nil),
                session: signInAvailable ? environment.session : nil,
                tiersOn: environment.isAvailable(.guideTiers),
                router: guidesRouter,
                openBis: { gear in
                    selection = .bis
                    bisRouter.openRaid(spec: gear.spec, band: gear.band)
                }
            )
        case .you where signInAvailable:
            YouScreen(
                session: environment.session,
                version: AccountCopy.version(environment.identity, data: environment.appConfig.current?.activeBuild),
                onPickClass: { selection = .bis },
                plansLink: environment.plansLink,
                liveSurfaces: environment.isAvailable(.widgets),
                onLiveSettingsChange: { Task { await environment.liveActivities.syncRaid(WidgetSnapshotStore.shared().load()) } },
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

    /// A link selects its tab: sim, report, rankings and player links Home (lanes F5, F10, F11),
    /// guide links Guides and list links BiS (lane F9), raid links Raids (lane F4).
    private func openLink(_ link: DeepLink) {
        if SimsRouter.handles(link) {
            guard environment.isAvailable(.sims) else { return }
            selection = .home
            simsRouter.open(link)
        } else if ReportsRouter.handles(link) {
            guard environment.isAvailable(.reports) else { return }
            selection = .home
            reportsRouter.open(link)
        } else if RankingsRouter.handles(link) {
            guard environment.isAvailable(.rankings) else { return }
            selection = .home
            rankingsRouter.open(link)
        } else if GuidesRouter.handles(link) {
            guard environment.isAvailable(.guides) else { return }
            selection = .guides
            guidesRouter.open(link)
        } else if BisRouter.handles(link) {
            guard environment.isAvailable(.bis) else { return }
            selection = .bis
            bisRouter.open(link)
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

    /// Lane F15: an intent's or a Spotlight tap's destination, opened as the in-app link would be.
    private func openIntent(_ destination: IntentDestination) {
        switch destination {
        case let .link(link): openLink(link)
        case let .tab(tab): selection = tab
        case let .character(key): openCharacter(key)
        }
    }

    /// A character opens F11.4 when `rankings` is on; else Home on it (your own characters only).
    private func openCharacter(_ key: String) {
        let parts = key.split(separator: "/").map(String.init)
        if environment.isAvailable(.rankings), parts.count == 3 {
            openLink(.player(region: parts[0], ruleset: parts[1], name: parts[2]))
            return
        }
        if environment.session.phase.me?.characters.contains(where: { $0.key == key }) == true {
            environment.session.pick(key)
        }
        selection = .home
    }

    private func refreshSpotlight() async {
        guard scenePhase == .active else { return }
        let session = environment.session
        await SpotlightRefresher.live(environment, session: signInAvailable ? session : nil).refresh(
            me: session.phase.me, currentKey: session.currentCharacter?.key, enabled: environment.isAvailable(.system)
        )
    }

    /// Lane F9: on launch and on foreground (at most hourly), the guides of every class the
    /// player has a character in are fetched, so they open offline.
    private func prefetchGuides() async {
        guard scenePhase == .active, environment.isAvailable(.guides) else { return }
        let prefetcher = guidesPrefetcher ?? GuidesPrefetcher(
            source: GuidesDependencies.live(environment, session: signInAvailable ? environment.session : nil).source,
            now: { Date() }
        )
        guidesPrefetcher = prefetcher
        await prefetcher.run(classes: playerClasses)
    }

    /// The classes the signed-in player has characters in.
    private var playerClasses: Set<String> {
        guard signInAvailable else { return [] }
        return Set(environment.session.phase.me?.characters.compactMap(\.className) ?? [])
    }
}

/// Lane F15: rebuilds Spotlight when the app comes forward, the account or character changes,
/// or the `system` switch flips.
private struct SpotlightTrigger: Equatable {
    let active: Bool
    let account: Me?
    let current: String?
    let enabled: Bool
}

/// Runs the guide prefetch again when the app comes forward or the player's classes change.
private struct GuidesPrefetchTrigger: Equatable {
    let active: Bool
    let classes: Set<String>
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
