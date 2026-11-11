import DesignSystem
import Models
import Networking
import Routing
import SwiftUI

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

/// The app: the four tabs (Home, Raid, Gear, Activity) under the profile menu (flows spec 3), the
/// links, pushes and intents that land on their answers, and the account's background work.
struct RootView: View {
    let environment: AppEnvironment
    @State var shell = AppShell(selection: .launchTab)
    /// Lane F4: the Raid tab's stack, which raid links and notifications push onto.
    @State var raidsRouter = RaidsRouter()
    /// Lane F5: Home's Sims stack, which sim links push onto.
    @State var simsRouter = SimsRouter()
    /// Lane F12: the search cover and the list a result asks the Gear tab for.
    @State var search: SearchNavigator
    /// Lane F10: Home's Your reports stack, which report links push onto.
    @State var reportsRouter = ReportsRouter()
    /// Lane F11: Home's Rankings and character pages, which rankings and player links push onto.
    @State var rankingsRouter = RankingsRouter()
    /// Lane F9: guide links, pushed under Home.
    @State var guidesRouter = GuidesRouter()
    /// Lane F9: lists a guide's Gear row and `/app/bis` links ask the Gear tab for.
    @State var bisRouter = BisRouter()
    /// Lane F9: keeps the guides of the player's classes saved.
    @State private var guidesPrefetcher: GuidesPrefetcher?
    @Environment(\.scenePhase) private var scenePhase
    /// Lane F15: where an intent or a Spotlight tap asks the app to go.
    @State private var intentRouter = IntentRouter.shared
    /// Lane F16: the app's raid calendar, kept up to date from the Raid tab. It never asks for
    /// access until the player flips its switch.
    @State var calendarSync: CalendarSyncController
    /// Lane F13: the planner's editor and link pages, behind `planner`.
    @State var plannerRouter: PlannerRouter
    @Environment(\.openURL) var openURL
    /// Flows spec F8: Pick your main, then the guild member's notification ask.
    @State private var firstRun: FirstRunModel
    /// Answers saved anywhere in the app: the Raid badge reads again after each.
    @State private var raidSignupSaves = 0

    init(environment: AppEnvironment) {
        self.environment = environment
        let session = environment.session
        _firstRun = State(initialValue: FirstRunModel(dependencies: .init(
            setMain: { key in try await session.setMain(key) },
            permission: SystemNotificationPermission(),
            asked: DefaultsFirstRunAskStore()
        )))
        _search = State(initialValue: SearchNavigator(dependencies: .live(environment)))
        let service = RaidsDependencies.live(environment).service
        let loader = CalendarInputLoader(service: service, currentKey: { session.currentCharacter?.key })
        _calendarSync = State(initialValue: CalendarSyncController(
            adapter: EventKitCalendarAdapter(), store: UserDefaultsCalendarSyncStore(), now: AppClock.now, clock: RaidClock(),
            loadInput: { await loader.load() }
        ))
        _plannerRouter = State(initialValue: PlannerRouter(dependencies: .live(environment)) {
            environment.isAvailable(.planner)
        })
    }

    var body: some View {
        if environment.updateGate.isRequired {
            UpdateRequiredView(storeURL: environment.updateGate.storeURL)
        } else {
            tabs
        }
    }

    private var tabs: some View {
        TabView(selection: $shell.selection) {
            ForEach(AppTab.allCases) { tab in
                Tab(tab.title, image: tab.icon, value: tab) {
                    screen(for: tab).environment(\.tabChrome, chrome(for: tab)).tint(Theme.Palette.accent)
                }
                .badge(shell.badges[tab] ?? 0)
            }
        }
        // The selected tab's gold (spec "Tab-bar icons"); the screens keep the accent.
        .tint(Theme.Palette.tabSelected)
        .modifier(ProfileDrawer(shell: shell) { menu })
        .environment(environment.session)
        .environment(simsRouter)
        .environment(\.searchNavigator, searchIsOn ? search : nil)
        .environment(\.openAppLink, OpenAppLinkAction { openLink($0) })
        .fullScreenCover(isPresented: $search.isPresented) {
            SearchScreen(navigator: search).tint(Theme.Palette.accent)
        }
        .onChange(of: search.bisRequest) { _, request in
            if request != nil { shell.selection = .gear }
        }
        .environment(reportsRouter)
        .environment(rankingsRouter)
        .environment(plannerRouter)
        .modifier(PlannerHost(router: plannerRouter, newBuild: plannerNewBuild))
        .task { await environment.appConfig.refresh() }
        .task { await environment.session.refresh() }
        .task { await environment.session.watchForDeadTokens() }
        // Flows spec 5.1: the device token follows the account; the ask itself comes in context (E6).
        .task(id: PushTrigger(signedIn: environment.session.phase.isSignedIn, active: scenePhase == .active)) {
            PushRegistrar.shared.calls = PushTokenCalls(client: environment.client, appVersion: environment.identity.version)
            await PushRegistrar.shared.sync(signedIn: environment.session.phase.isSignedIn)
        }
        .widgetRefreshTriggers(environment)
        // The Raid tab's badge, for the selected character's guild, on every foreground.
        .task(id: RaidBadgeTrigger(guild: raidBadgeGuild, active: scenePhase == .active, signupSaves: raidSignupSaves)) {
            shell.badges[.raid] = await RaidBadgeCount.load(guild: raidBadgeGuild, client: raidBadgeClient)
        }
        // The Activity tab's badge at shell start and on every foreground, as the Raid badge (round 5, i1).
        .task(id: ActivityBadgeTrigger(signedIn: environment.session.phase.isSignedIn, active: scenePhase == .active)) {
            guard scenePhase == .active else { return }
            shell.badges[.activity] = await ActivityBadgeCount.load(
                signedIn: environment.session.phase.isSignedIn, service: activityService
            )
        }
        .onReceive(NotificationCenter.default.publisher(for: .raidSignupSaved)) { _ in raidSignupSaves += 1 }
        .modifier(FirstRunHost(session: environment.session, model: firstRun, isOn: signInAvailable))
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
            RaidNotificationCenter.shared.onOpen = openIntent
            #if DEBUG
                // Review captures only: the launch arguments in Raids/Debug/RaidDebugLaunch.swift.
                // Lane F13's captures open the planner the same way (UITests/PlannerCaptureTests),
                // and lane F15's open guides and character pages (UITests/SystemCaptureTests).
                if let link = RaidDebugLaunch.link { openLink(link) }
                // `-FSMenu <row>` pushes a menu row on the launch tab (Settings, Notifications...),
                // as a tap on the profile menu would, so a capture opens it without taps.
                if let row = UserDefaults.standard.string(forKey: "FSMenu").flatMap(MenuDestination.init(rawValue:)) {
                    shell.open(row)
                }
                await RaidDebugLaunch.postPushIfAsked()
            #endif
        }
    }

    /// Lane F12: universal search, its button and the item page.
    var searchIsOn: Bool { environment.isAvailable(.search) }

    /// Lane F1: the account states and the selected character.
    var signInAvailable: Bool { environment.isAvailable(.signIn) }

    /// What each tab root needs from the shell: the profile button's character and the menu's pushes.
    private func chrome(for tab: AppTab) -> TabChrome {
        TabChrome(
            tab: tab, shell: shell, character: signInAvailable ? environment.session.currentCharacter : nil,
            destination: { AnyView(menuScreen($0)) }
        )
    }

    /// The tab roots. Each screen marks its root with `.tabRootChrome()` inside its stack.
    @ViewBuilder
    private func screen(for tab: AppTab) -> some View {
        switch tab {
        case .home where signInAvailable && environment.session.homeGateHolds:
            // Lane F15: rankings and player links open over the gate (public data).
            PublicRankingsHost(
                dependencies: HomeDependencies.live(environment).rankings,
                account: RankingsAccount.of(environment.session.phase.me, current: nil), router: rankingsRouter
            ) {
                HomeAccountGate(session: environment.session, onPickClass: { shell.selection = .gear }) { EmptyView() }
            }
        case .home where signInAvailable:
            HomeAccountGate(session: environment.session, onPickClass: { shell.selection = .gear }) {
                HomeTab(dependencies: .live(environment), session: environment.session, openBis: { shell.selection = .gear }, open: openIntent)
            }
        case .home: HomeTab(dependencies: .live(environment), openBis: { shell.selection = .gear }, open: openIntent)
        case .raid where environment.isAvailable(.raids):
            RaidsTab(
                dependencies: .live(environment), router: raidsRouter, session: signInAvailable ? environment.session : nil,
                extras: guildExtras, calendar: guildExtras == nil ? nil : calendarSync
            )
        case .raid: PlaceholderScreen(title: tab.title)
        case .gear where environment.isAvailable(.bis):
            BisScreen.live(environment, session: signInAvailable ? environment.session : nil, router: bisRouter)
        case .gear: PlaceholderScreen(title: tab.title)
        case .activity:
            ActivityScreen(
                service: activityService,
                session: signInAvailable ? environment.session : nil,
                build: environment.appConfig.current?.activeBuild ?? BundledBuild.read() ?? "",
                onUnread: { shell.badges[.activity] = $0 },
                open: openIntent
            )
        }
    }

    /// Lane F16: the guild extras' dependencies, nil while `guild_extras` is off.
    var guildExtras: GuildExtrasDependencies? {
        environment.isAvailable(.guildExtras) ? .live(environment) : nil
    }

    /// Lane F15: the Offline data screen's model factory, nil while `system` is off. The closure is
    /// main-actor isolated by name: Swift 6.2 (CI's Xcode 26.6) does not carry the parameter's
    /// `@MainActor` type into a closure literal inside a ternary and rejects it as non-Sendable.
    var offlineData: (@MainActor () -> OfflineDataModel)? {
        guard environment.isAvailable(.system) else { return nil }
        let environment = environment
        return { @MainActor in OfflineDataModel.live(environment) }
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
        // Lane F15: `Keep my classes ready` drives it; `Download on Wi-Fi only` holds new guides.
        let settings = OfflineSettings()
        guard scenePhase == .active, environment.isAvailable(.guides), settings.keepClassesReady,
              settings.mayDownload(on: PathConnectivity.shared) else { return }
        let prefetcher = guidesPrefetcher ?? GuidesPrefetcher(
            source: GuidesDependencies.live(environment, session: signInAvailable ? environment.session : nil).source,
            now: { Date() }
        )
        guidesPrefetcher = prefetcher
        await prefetcher.run(classes: playerClasses)
    }

    /// Lane F13: Your builds' `Plan a new build` row, for the current character.
    func plannerNewBuild() -> PlannerNewBuild? {
        let character = environment.session.currentCharacter
        return PlannerNewBuild(className: character?.className, race: character?.race, level: character?.level)
    }

    /// Activity's calls, for the tab and its badge.
    private var activityService: LiveActivityService { LiveActivityService(client: environment.client) }

    /// The Raid badge's calls: the same answers the Raid tab reads (round 5, i2: in the review
    /// captures the officer fixture's badge, as Android's tab and badge share one load).
    private var raidBadgeClient: APIClient {
        #if DEBUG
            if let fixture = RaidDebugLaunch.client(baseURL: environment.configuration.baseURL) { return fixture }
        #endif
        return environment.client
    }

    /// The selected character's guild, whose raids the Raid badge counts.
    private var raidBadgeGuild: Int64? {
        guard signInAvailable, environment.isAvailable(.raids) else { return nil }
        return environment.session.currentCharacter?.guild?.id
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

/// Reads the Activity badge again on sign-in and on a foreground.
private struct ActivityBadgeTrigger: Equatable {
    let signedIn: Bool
    let active: Bool
}

/// Reads the Raid badge again for a new guild, a foreground, or a saved answer.
private struct RaidBadgeTrigger: Equatable {
    let guild: Int64?
    let active: Bool
    let signupSaves: Int
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
                .tabRootChrome()
        }
    }
}

/// When the push token is checked: a sign-in or sign-out, and every foreground (a permission
/// turned on in Settings registers on return).
private struct PushTrigger: Equatable {
    let signedIn: Bool
    let active: Bool
}
