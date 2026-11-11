import Auth
import DesignSystem
import Models
import Networking
import Observation
import Routing
import SwiftUI

/// The raids feature's dependencies, built once from the app's environment.
@MainActor
struct RaidsDependencies {
    let service: any RaidService
    let cache: any RaidCache
    let haptics: any Haptics
    let permission: any NotificationPermission
    let askStore: any ReminderAskStore
    let now: @Sendable () -> Date
    let clock: RaidClock
    /// The active data build, for the game icons.
    let iconBuild: @Sendable () async -> String?
    /// Lane F10: the report F4.22's row pushes; nil while the reports switch is off.
    let reports: ReportsDependencies?
    /// Lane F15: the connection, for the stale strip; nil (tests) keeps v1's offline rule alone.
    var connectivity: (any Connectivity)?
    /// The flows' calls (flows spec 4.2): the badge, applicants, recruiting guilds, the Ready check.
    var flow: any RaidFlowService
    /// The officer's Discord mention.
    var ping: any RaidPingService
    /// What the site still does for the app (Apply, Claim, connecting the Discord bot).
    var site: any RaidSiteActions = LiveRaidSiteActions()

    static func live(_ environment: AppEnvironment) -> RaidsDependencies {
        // The active build names every icon URL. In fixture mode it comes from the bundled
        // cdn.active-build fixture, as Android's fixture engine answers it; live, from the site's
        // /data/active-build.json (published by web/scripts/sync-data.mjs).
        let staticTransport: any HTTPTransport = switch environment.configuration.mode {
        case .fixture: FixtureTransport.bundled() ?? URLSessionTransport()
        case .live: URLSessionTransport()
        }
        let statics = StaticDataClient(transport: staticTransport)
        var service: any RaidService = LiveRaidService(client: environment.client)
        var flow: any RaidFlowService = LiveRaidFlowService(client: environment.client)
        #if DEBUG
        if let fixture = RaidDebugLaunch.service(baseURL: environment.configuration.baseURL) { service = fixture }
        if let fixture = RaidDebugLaunch.flow(baseURL: environment.configuration.baseURL) { flow = fixture }
        #endif
        return RaidsDependencies(
            service: service,
            cache: DiskRaidCache.standard(),
            haptics: SystemHaptics(),
            permission: SystemNotificationPermission(),
            askStore: UserDefaultsReminderAskStore(),
            now: AppClock.now,
            clock: RaidClock(),
            iconBuild: { try? await statics.activeBuild().build },
            reports: environment.isAvailable(.reports) ? .live(environment) : nil,
            connectivity: PathConnectivity.shared,
            flow: flow,
            ping: LiveRaidPingService(client: environment.client)
        )
    }

    var signup: SignupController.Dependencies {
        SignupController.Dependencies(
            service: service, haptics: haptics, permission: permission, askStore: askStore, now: now, hold: .seconds(4)
        )
    }
}

/// A screen pushed on the Raids stack.
enum RaidsRoute: Hashable, Sendable {
    case raid(guild: Int64, raid: Int64, roster: Bool, open: RaidOpenAction?)
    /// The raid night page before and during the raid; Your raid night after it.
    case night(raid: Int64)
    /// F3: the raid page filtered to Not answered.
    case notAnswered(guild: Int64, raid: Int64)
    /// F3: one applicant.
    case applicant(guild: Int64, application: Int64)
    /// The guild's planned raids.
    case calendar
    /// Lane F10: a report from Your raid night.
    case report(id: String)
    /// Lane F16: the guild's roster and raid history.
    /// `only`: the characters the roster lists (the Raid tab's not-synced raiders); empty lists all.
    case roster(guild: Int64, only: [String] = [])
    case history(guild: Int64)
}

/// The Raids tab's navigation: deep links and notifications push here (spec 5.0, 6.4).
@MainActor
@Observable
final class RaidsRouter {
    var path: [RaidsRoute] = []

    /// Opens a raid link on a fresh stack, so back goes to the tab's root.
    func open(_ link: DeepLink) {
        switch link {
        case let .raid(guild, raid, _, _, view) where view == RaidsRouter.unansweredView:
            path = [.notAnswered(guild: guild, raid: raid)]
        case let .raid(guild, raid, spec, state, _):
            let press = spec == "ask" ? SignupState.signedIn : state.flatMap(SignupState.init)
            path = [.raid(guild: guild, raid: raid, roster: false, open: press.map(RaidOpenAction.init))]
        case let .raidNight(guild, raid):
            path = [.raid(guild: guild, raid: raid, roster: false, open: nil), .night(raid: raid)]
        case let .guildApplicant(guild, application):
            path = [.applicant(guild: guild, application: application)]
        case .guild:
            path = []
        default:
            break
        }
    }

    /// Opens one of the tab's screens on a fresh stack (Home's raid actions, a push).
    func open(_ route: RaidsRoute) {
        path = [route]
    }

    /// The raid link's `view=` that opens Not answered (links.json).
    static let unansweredView = "unanswered"

    /// True for the links the Raids tab opens.
    static func handles(_ link: DeepLink) -> Bool {
        switch link {
        case .raid, .raidNight, .guild, .guildApplicant: true
        default: false
        }
    }
}

/// The Raid tab root (flows spec 4.2). Signed in, it follows the session's current character's
/// guild; otherwise `/v1/me`'s main character.
struct RaidsTab: View {
    let dependencies: RaidsDependencies
    @Bindable var router: RaidsRouter
    /// Lane F1's account, nil while its feature switch is off.
    let session: SessionStore?
    /// Lane F16's guild extras and calendar sync, nil while `guild_extras` is off.
    let extras: GuildExtrasDependencies?
    let calendar: CalendarSyncController?
    @State private var model: RaidsTabModel
    @State private var loader: RaidRootLoader
    @State private var answers: RaidAnswerModel
    @State private var rosterStore: RosterStore?
    @State private var iconBuild: String?
    @Environment(\.openURL) private var openURL

    init(
        dependencies: RaidsDependencies, router: RaidsRouter, session: SessionStore? = nil,
        extras: GuildExtrasDependencies? = nil, calendar: CalendarSyncController? = nil
    ) {
        self.dependencies = dependencies
        self.router = router
        self.session = session
        self.extras = extras
        self.calendar = calendar
        _model = State(initialValue: RaidsTabModel(
            service: dependencies.service, cache: dependencies.cache, now: dependencies.now, connectivity: dependencies.connectivity
        ))
        _loader = State(initialValue: RaidRootLoader(service: dependencies.flow, raids: dependencies.service, now: dependencies.now))
        _answers = State(initialValue: RaidAnswerModel(service: dependencies.service, haptics: dependencies.haptics))
    }

    var body: some View {
        NavigationStack(path: $router.path) {
            RaidRootView(content: rootContent, actions: rootActions)
                .refreshable { await reload() }
                .bandBar()
                // Round 5 on iOS (m-ios-raid): the tab's large title on the band, under the profile button.
                .navigationTitle(AppTab.guild.title)
                .toolbarTitleDisplayMode(.large)
                .tabRootChrome()
                .searchButton()
                .navigationDestination(for: RaidsRoute.self) { destination($0) }
        }
        .environment(\.raidIconBuild, iconBuild)
        .task { iconBuild = await dependencies.iconBuild() }
        .task(id: session?.currentCharacter?.key) {
            model.currentKey = session?.currentCharacter?.key
            await reload()
        }
    }

    private var content: RaidsTabContent? {
        if case let .loaded(content) = model.state { return content }
        return nil
    }

    private var composer: RaidRootComposer {
        let roster = rosterStore?.guild == content?.head.id ? rosterStore?.home?.roster : nil
        let presenter = RaidFlowPresenter(now: dependencies.now(), clock: dependencies.clock)
        return RaidRootComposer(
            presenter: presenter,
            members: roster?.count,
            atCap: roster?.filter { $0.level == RaidFlowPresenter.maxLevel }.count,
            synced: roster.map(presenter.syncedThisWeek),
            answering: answers.state,
            home: rosterStore?.guild == content?.head.id ? rosterStore?.home : nil,
            firstRaids: HomeFlowDependencies.firstRaids(BundledData.keyDates())
        )
    }

    private var rootContent: RaidRootContent {
        if session?.phase.isSignedIn == false { return .signedOut }
        return composer.content(model.state, extras: loader.extras)
    }

    private func reload() async {
        await model.load()
        if let content { answers.offline = content.offlineSince != nil }
        await loader.load(for: model.state)
        await refreshGuildExtras()
    }

    /// The roster's figures for the Roster row, and the calendar on every Raids refresh.
    private func refreshGuildExtras() async {
        guard let content else { return }
        if let calendar { await calendar.update(content.calendarInput) }
        guard let extras else { return }
        if rosterStore?.guild != content.head.id { rosterStore = extras.rosterStore(guild: content.head.id) }
        await rosterStore?.load()
    }

    /// A sign-up replaced a raid: the tab's list, then the calendar's event title.
    private func raidChanged(_ view: RaidView) {
        model.replace(view)
        guard let calendar, let content else { return }
        Task { await calendar.update(content.calendarInput) }
    }

    // MARK: Actions

    private var currentRaid: RaidView? {
        content.flatMap { HomeFlowModel.current($0.planned, now: dependencies.now()) }
    }

    private var rootActions: RaidRootActions {
        RaidRootActions(
            raid: NextRaidActions(
                answer: { state in answer { key, raid, note in await answers.answer(state, guild: raid.0, raid: raid.1, key: key, note: note) } },
                late: { minutes in answer { key, raid, note in await answers.late(minutes, guild: raid.0, raid: raid.1, key: key, note: note) } },
                retry: { answer { key, raid, note in await answers.retry(guild: raid.0, raid: raid.1, key: key, note: note) } },
                openRaid: { push { .raid(guild: $0, raid: $1.id, roster: false, open: nil) } },
                openReady: { push { _, raid in .night(raid: raid.id) } },
                previewMention: { push { .notAnswered(guild: $0, raid: $1.id) } },
                openApplicants: {
                    guard let guild = content?.head.id, let first = loader.extras.applicants.first else { return }
                    router.path.append(.applicant(guild: guild, application: first.id))
                }
            ),
            openPage: openPage,
            openRoster: { keys in
                guard let guild = content?.head.id, extras != nil else { return }
                router.path.append(.roster(guild: guild, only: keys))
            },
            // The cover list heads the raid page's sign-ups.
            openCover: { push { .raid(guild: $0, raid: $1.id, roster: true, open: nil) } },
            openLink: { openURL($0) },
            apply: { openURL(dependencies.site.apply(guild: $0.guildID)) },
            claim: { openURL(dependencies.site.claim) },
            findGuild: { openURL(dependencies.site.guilds) },
            signIn: { Task { await session?.signIn(deviceName: UIDevice.current.name) } },
            retry: { Task { await reload() } },
            guild: guildActions
        )
    }

    /// The Guild tab's own sections (spec 2026-11-10-guild-tab.md): the roster and loot pages in the
    /// app, the rest on the guild's page on the site.
    private var guildActions: GuildTabActions {
        let web = content?.head.webURL
        return GuildTabActions(
            openRoster: { openPage(.roster) },
            openLoot: { openPage(.loot) },
            openProgression: { if let web { openURL(GuildWeb.tab(web, "progression")) } },
            openReport: { id in
                if dependencies.reports != nil { router.path.append(.report(id: id)) } else { openURL(RaidLinks.report(id)) }
            },
            openPlanner: {
                guard let web, let raid = currentRaid else { return }
                openURL(GuildWeb.raid(web, raid: raid.id))
            },
            officer: { kind in
                guard let web else { return }
                switch kind {
                case .plan: openURL(GuildWeb.tab(web, "raids"))
                case .applications:
                    if let guild = content?.head.id, let first = loader.extras.applicants.first {
                        router.path.append(.applicant(guild: guild, application: first.id))
                    } else {
                        openURL(web.appending(path: "applications"))
                    }
                case .setup: openURL(web.appending(path: "setup"))
                }
            }
        )
    }

    private func push(_ route: (Int64, RaidView) -> RaidsRoute) {
        guard let guild = content?.head.id, let raid = currentRaid else { return }
        router.path.append(route(guild, raid))
    }

    /// Answers the current raid for the tab's character (its viewer row's key and note).
    private func answer(_ send: @escaping (String, (Int64, Int64), String) async -> RaidView?) {
        guard let content, let raid = currentRaid else { return }
        let key = raid.viewer.characters.first { $0.characterKey == content.character.key }?.characterKey ?? content.character.key
        let note = raid.signups.first { $0.characterKey == key }?.note ?? ""
        Task {
            if let view = await send(key, (content.head.id, raid.id), note) { raidChanged(view) }
        }
    }

    private func openPage(_ page: GuildPage) {
        guard let content else { return }
        let guild = content.head.id
        switch page {
        case .roster where extras != nil: router.path.append(.roster(guild: guild))
        case .loot where extras != nil: router.path.append(.history(guild: guild))
        case .calendar: router.path.append(.calendar)
        case .roster, .loot, .sims:
            if let url = content.head.webURL { openURL(page == .sims ? url.appending(path: "sims") : url) }
        }
    }

    // MARK: Destinations

    @ViewBuilder
    private func destination(_ route: RaidsRoute) -> some View {
        switch route {
        case let .raid(guild, raid, roster, open):
            RaidPageScreen(
                model: RaidPageModel(
                    guild: guild, raid: raid, deps: dependencies.signup, cache: dependencies.cache, openAction: open,
                    onChange: { view in raidChanged(view) }, connectivity: dependencies.connectivity
                ),
                clock: dependencies.clock, now: dependencies.now, currentKey: content?.character.key,
                guildName: content?.head.name ?? "", guildURL: content?.head.webURL,
                openNight: { router.path.append(.night(raid: $0)) },
                back: { router.path.removeAll() },
                scrollToRoster: roster
            )
        case let .night(raid):
            RaidNightPageScreen(
                model: RaidNightPageModel(raid: raid, service: dependencies.service),
                clock: dependencies.clock, now: dependencies.now,
                openRoster: { if let guild = content?.head.id { router.path.append(.raid(guild: guild, raid: raid, roster: true, open: nil)) } },
                after: { AnyView(afterNight(raid)) }
            )
        case let .notAnswered(guild, raid):
            NotAnsweredScreen(
                model: NotAnsweredModel(
                    guild: guild, raid: raid, service: dependencies.service, ping: dependencies.ping, hold: .seconds(4)
                ),
                guildName: content?.head.name ?? "", site: dependencies.site, clock: dependencies.clock
            )
        case let .applicant(guild, application):
            ApplicantScreen(
                model: ApplicantModel(guild: guild, application: application, service: dependencies.flow),
                guildName: content?.head.name ?? "", now: dependencies.now,
                back: { if !router.path.isEmpty { router.path.removeLast() } }
            )
        case .calendar:
            RaidCalendarScreen(
                planned: content?.planned ?? [], clock: dependencies.clock, now: dependencies.now, currentKey: content?.character.key
            ) { raid in
                if let guild = content?.head.id { router.path.append(.raid(guild: guild, raid: raid.id, roster: false, open: nil)) }
            }
        case let .report(id):
            if let reports = dependencies.reports {
                ReportScreen(
                    route: ReportRoute(id: id), dependencies: reports,
                    account: ReportsAccount.of(session?.phase.me, current: session?.currentCharacter)
                )
            }
        case let .roster(guild, only):
            if let extras {
                RosterRoute(extras: extras, guild: guild, shared: rosterStore, only: only.isEmpty ? nil : Set(only))
            }
        case let .history(guild):
            if let extras {
                HistoryRoute(
                    extras: extras, guild: guild,
                    yourKeys: content?.accountKeys ?? [],
                    openNight: { router.path.append(.night(raid: $0)) }
                )
            }
        }
    }

    /// Your raid night once it is over (spec F4.22).
    private func afterNight(_ raid: Int64) -> some View {
        RaidNightView(
            model: RaidNightModel(
                raid: raid, service: dependencies.service, cache: dependencies.cache, now: dependencies.now, connectivity: dependencies.connectivity
            ),
            clock: dependencies.clock, now: dependencies.now,
            openReport: dependencies.reports == nil ? nil : { router.path.append(.report(id: $0)) }
        )
    }
}
