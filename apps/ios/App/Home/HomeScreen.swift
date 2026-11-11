import Auth
import DesignSystem
import DesignTokens
import Models
import Networking
import SwiftUI

/// The site pages Home opens in the browser.
enum HomeLinks {
    static let setup = URL(string: "https://foreversixty.gg/setup")!
    static let planner = URL(string: "https://foreversixty.gg/planner")!
}

/// What Home needs from the app: the API, the static data, the cache and the feature switch.
struct HomeDependencies {
    let client: APIClient
    let catalog: ItemCatalog
    let cache: any HomeCache
    let isEnabled: Bool
    /// Lane F5: Sims, pushed from the Sims row; nil while its feature switch is off.
    var sims: SimsDependencies?
    /// Lane F6: the Talents screen's files; nil while its feature switch is off.
    var talents: (any TalentsSource)?
    /// Lane F10: Your reports, pushed from the Logs row; nil while its feature switch is off.
    var reports: ReportsDependencies?
    /// Lane F11: Rankings and character pages, pushed from the Rankings row; nil while its switch is off.
    var rankings: RankingsDependencies?
    /// Lane F15: the CDN's active build now, for Talents' new-data strip.
    var latestBuild: @Sendable () async -> String? = { nil }
    /// The time-aware Home's calls (flows spec 4.1).
    var flow: HomeFlowDependencies?

    @MainActor
    static func live(_ environment: AppEnvironment) -> HomeDependencies {
        let catalog = ItemCatalog(store: environment.staticData)
        let reports: ReportsDependencies? = environment.isAvailable(.reports) ? .live(environment) : nil
        return HomeDependencies(
            client: environment.client,
            catalog: catalog,
            cache: FileHomeCache.standard() ?? MemoryHomeCache(),
            isEnabled: environment.isAvailable(.characterHome),
            sims: environment.isAvailable(.sims) ? .live(environment, catalog: catalog) : nil,
            talents: environment.isAvailable(.talents) ? LiveTalentsSource.make(environment) : nil,
            reports: reports,
            rankings: environment.isAvailable(.rankings) ? .live(environment, catalog: catalog, reports: reports) : nil,
            latestBuild: { [store = environment.staticData] in await store.latestBuild() },
            flow: .live(environment)
        )
    }
}

/// The Home tab: finds the character, then shows its answer. Signed in (lane F1), it follows the
/// session's current character and carries the selector; without a session it shows the main.
struct HomeTab: View {
    /// Switches to the BiS tab on the character's list (lane F3 owns that tab).
    let openBis: () -> Void
    /// Opens a destination elsewhere in the app (a tab, a link): Home's raid, report and activity rows.
    let open: (IntentDestination) -> Void
    /// Lane F1's account, nil while its feature switch is off.
    let session: SessionStore?

    @State private var dependencies: HomeDependencies

    @State private var fetchedMe: Me?
    @State private var meFailed = false

    /// The session's `/v1/me` when it has one, else Home's own read.
    private var me: Me? { session?.phase.me ?? fetchedMe }

    private var character: AccountCharacter? {
        guard let me else { return nil }
        guard let session else { return me.mainCharacter }
        return CharacterSelection.current(in: me, picked: session.pickedKey)
    }

    var body: some View {
        if !dependencies.isEnabled {
            PlaceholderScreen(title: Copy.tabHome)
        } else if let me, let character {
            HomeCharacterScreen(
                character: HomeCharacter(character),
                serverSims: me.entitlements.serverSims,
                reportsAccount: ReportsAccount.of(me, current: character),
                rankingsAccount: RankingsAccount.of(me, current: character),
                dependencies: dependencies,
                openBis: openBis,
                open: open
            )
            .id(character.key)
        } else {
            NavigationStack {
                VStack(alignment: .leading, spacing: 12) {
                    if me != nil {
                        Text(Copy.signinNoCharsTitle).font(ScaledType.sec).foregroundStyle(FSColor.strong)
                        Text(Copy.signinNoCharsBody).font(ScaledType.body).foregroundStyle(FSColor.muted)
                    } else if meFailed {
                        Text(Copy.homeAnswerFailed).font(ScaledType.lead).foregroundStyle(FSColor.wipe)
                        SecondaryButton(title: Copy.commonTryAgain, systemImage: "arrow.clockwise") { Task { await loadMe() } }
                    } else {
                        HomeSkeleton()
                    }
                }
                .frame(maxWidth: .infinity, maxHeight: .infinity, alignment: .topLeading)
                .background(FSColor.bg)
            }
            .task { await loadMe() }
        }
    }

    init(
        dependencies: HomeDependencies, session: SessionStore? = nil, openBis: @escaping () -> Void = {},
        open: @escaping (IntentDestination) -> Void = { _ in }
    ) {
        self.open = open
        _dependencies = State(initialValue: dependencies)
        self.session = session
        self.openBis = openBis
    }

    private func loadMe() async {
        meFailed = false
        do {
            fetchedMe = try await dependencies.client.send(.getMe)
        } catch {
            meFailed = true
        }
    }
}

/// Home for one character: the answer and the item sheet.
struct HomeCharacterScreen: View {
    let character: HomeCharacter
    /// `entitlements.server_sims`, for Sims' Server runs section (lane F5).
    let serverSims: Bool
    /// Lane F10: who reads Your reports (crests, the coach card's `log_coach`).
    let reportsAccount: ReportsAccount
    /// Lane F11: who reads Rankings (the aim, your row, the `You` tags, Parse history).
    let rankingsAccount: RankingsAccount
    let dependencies: HomeDependencies
    let openBis: () -> Void
    let open: (IntentDestination) -> Void

    @State private var store: HomeStore
    /// The time-aware Home (flows spec 4.1); nil while its calls are not wired (tests).
    @State private var flow: HomeFlowModel?
    @State private var sheet: ItemSheetRequest?
    /// Lane F12: the item page, pushed from the sheet's `More about this item`.
    @State private var itemPage: ItemPageRoute?
    @Environment(\.searchNavigator) private var search
    @Environment(SimsRouter.self) private var simsRouter: SimsRouter?
    @Environment(ReportsRouter.self) private var reportsRouter: ReportsRouter?
    @Environment(RankingsRouter.self) private var rankingsRouter: RankingsRouter?
    /// Lane F11: the watch list of a character page a player link pushes on Home.
    @State private var watchList: WatchList?
    @Environment(\.scenePhase) private var scenePhase

    init(
        character: HomeCharacter,
        serverSims: Bool = false,
        reportsAccount: ReportsAccount = .signedOut,
        rankingsAccount: RankingsAccount = .signedOut,
        dependencies: HomeDependencies,
        openBis: @escaping () -> Void,
        open: @escaping (IntentDestination) -> Void = { _ in }
    ) {
        self.open = open
        self.character = character
        self.serverSims = serverSims
        self.reportsAccount = reportsAccount
        self.rankingsAccount = rankingsAccount
        self.dependencies = dependencies
        self.openBis = openBis
        let store = HomeStore.live(client: dependencies.client, cache: dependencies.cache)
        _store = State(initialValue: store)
        _flow = State(initialValue: dependencies.flow.map { HomeFlowModel(character: character, home: store, dependencies: $0) })
    }

    var body: some View {
        NavigationStack {
            HomeFlowView(content: flowContent, build: home?.dataBuild ?? "", actions: flowActions)
                .refreshable { await flow?.refresh() }
                .bandBar()
                .navigationTitle("")
                .navigationBarTitleDisplayMode(.inline)
                .navigationDestination(isPresented: simsShown) { sims }
            .navigationDestination(item: $itemPage) { route in
                if let search {
                    ItemPageScreen(route: route, dependencies: search.dependencies) { search.openList($0) }
                }
            }
            .navigationDestination(isPresented: reportsShown) { reports }
            .navigationDestination(isPresented: rankingsShown) { rankings }
            .navigationDestination(item: playerShown) { player($0) }
            .tabRootChrome()
            .searchButton()
        }
        .sheet(item: $sheet) { request in
            ItemSheetScreen(
                request: request, catalog: dependencies.catalog, needs: LiveItemNeedsSource(client: dependencies.client),
                moreAbout: search.map { _ in {
                    sheet = nil
                    let item = request.tooltip
                    itemPage = ItemPageRoute(itemID: item.itemID, fallback: item.fallback, classSlug: item.classSlug)
                } }
            )
        }
        .task { if let flow { await flow.load() } else { await store.load(character) } }
        // Offline the raid card's answers disable, as the Raid tab's (F1).
        .onChange(of: offlineStamp != nil, initial: true) { _, offline in flow?.offline = offline }
        .onChange(of: scenePhase) { _, phase in
            if phase == .active { Task { await store.refreshIfStale(character); await flow?.refresh() } }
        }
    }

    private var home: CharacterHome? {
        switch store.phase {
        case let .loaded(home, _): home
        case let .failed(cached): cached
        case .loading: nil
        }
    }

    /// What Home draws: the head always, the blocks once `/v1/home` answers.
    private var flowContent: HomeFlowContent {
        let presenter = HomeFlowPresenter(now: dependencies.flow?.now() ?? Date())
        let phase: HomeFlowPhase = switch flow?.phase {
        case .loaded: .loaded(flow.map { presenter.blocks($0.blocks, $0.inputs) } ?? [])
        case .failed: .failed
        case .loading, nil: .loading
        }
        return HomeFlowContent(
            head: presenter.head(character), phase: phase,
            staleStamp: store.staleSince.map { HomeTime(now: Date()).stamp($0) }, offlineStamp: offlineStamp
        )
    }

    private var flowActions: HomeFlowActions {
        let raidTab = { open(.tab(.raid)) }
        // Each raid action opens its screen under Raid; without the raid loaded, the tab.
        let raidScreen = { (route: (Int64, Int64) -> RaidsRoute?) in
            guard let guild = flow?.character.guildID, let raid = flow?.inputs.raid?.id, let screen = route(guild, raid) else {
                raidTab()
                return
            }
            open(.raid(screen))
        }
        let raidNight = { raidScreen { _, raid in .night(raid: raid) } }
        return HomeFlowActions(
            raid: NextRaidActions(
                answer: { state in Task { await flow?.answer(state) } },
                late: { minutes in Task { await flow?.late(minutes) } },
                retry: { Task { await flow?.retry() } },
                openRaid: { raidScreen { .raid(guild: $0, raid: $1, roster: false, open: nil) } },
                openReady: raidNight,
                previewMention: { raidScreen { .notAnswered(guild: $0, raid: $1) } },
                openApplicants: { raidScreen { guild, _ in flow?.inputs.applicants.first.map { .applicant(guild: guild, application: $0.id) } } }
            ),
            openRaidNight: raidNight,
            openItem: { row in
                if let home { sheet = request(for: row.look, home: home, sourceLine: row.sourceLine, wornID: row.wornID, gain: row.gain) }
            },
            openGear: openBis,
            openLoot: { loot in if let home { sheet = request(for: loot.look, home: home, sourceLine: nil, wornID: nil, gain: nil) } },
            openReport: { id in open(.link(.reportPage(id: id))) },
            simAgain: { simsRouter?.showsList = dependencies.sims != nil },
            openActivity: { row in
                PushLanding.destination(kind: row.kind, link: URL(string: row.link), data: [:]).map(open)
            },
            allActivity: { open(.tab(.activity)) },
            findGuild: raidTab,
            retry: { Task { await flow?.refresh() } }
        )
    }

    private var offlineStamp: String? {
        guard case let .loaded(_, since?) = store.phase else { return nil }
        return HomeTime(now: Date()).stamp(since)
    }


    /// Lane F5: Sims on Home's stack, driven by the router so a sim link can push it too.
    private var simsShown: Binding<Bool> {
        Binding(get: { simsRouter?.showsList ?? false }, set: { simsRouter?.showsList = $0 })
    }

    @ViewBuilder
    private var sims: some View {
        if let deps = dependencies.sims, let simsRouter {
            SimsScreen(dependencies: deps, entitled: serverSims, router: simsRouter)
        }
    }

    /// Lane F10: Your reports on Home's stack, driven by the router so a report link can push it too.
    private var reportsShown: Binding<Bool> {
        Binding(get: { reportsRouter?.showsList ?? false }, set: { reportsRouter?.showsList = $0 })
    }

    @ViewBuilder
    private var reports: some View {
        if let deps = dependencies.reports, let reportsRouter {
            ReportsScreen(dependencies: deps, account: reportsAccount, router: reportsRouter)
        }
    }

    /// Lane F11: Rankings on Home's stack, driven by the router so a rankings link can push it too.
    private var rankingsShown: Binding<Bool> {
        Binding(get: { rankingsRouter?.showsBoard ?? false }, set: { rankingsRouter?.showsBoard = $0 })
    }

    /// Lane F11: a character page a player link pushes on Home.
    private var playerShown: Binding<PlayerRoute?> {
        Binding(get: { rankingsRouter?.player }, set: { rankingsRouter?.player = $0 })
    }

    @ViewBuilder
    private var rankings: some View {
        if let deps = dependencies.rankings {
            RankingsScreen(dependencies: deps, account: rankingsAccount, reportsAccount: reportsAccount, encounter: rankingsRouter?.encounter)
                .id(rankingsRouter?.encounter)
        }
    }

    @ViewBuilder
    private func player(_ route: PlayerRoute) -> some View {
        if let deps = dependencies.rankings {
            let list = watchList ?? WatchList(service: deps.service, signedIn: rankingsAccount.signedIn)
            CharacterScreen(route: route, dependencies: deps, account: rankingsAccount, reportsAccount: reportsAccount, watchList: list)
                .onAppear { if watchList == nil { watchList = list } }
        }
    }


    private func request(
        for look: HomeItemLook, home: CharacterHome, sourceLine: String?, wornID: Int?, gain: HomeGain?
    ) -> ItemSheetRequest {
        ItemSheetRequest(
            tooltip: .forHome(look, home: home, character: character, sourceLine: sourceLine, wornID: wornID),
            gain: gain, guildID: character.guildID
        )
    }
}
