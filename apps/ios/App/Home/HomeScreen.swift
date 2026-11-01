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
            rankings: environment.isAvailable(.rankings) ? .live(environment, catalog: catalog, reports: reports) : nil
        )
    }
}

/// The Home tab: finds the character, then shows its answer. Signed in (lane F1), it follows the
/// session's current character and carries the selector; without a session it shows the main.
struct HomeTab: View {
    /// Switches to the BiS tab on the character's list (lane F3 owns that tab).
    let openBis: () -> Void
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
                finishedSims: me.finishedSims,
                serverSims: me.entitlements.serverSims,
                reportsAccount: ReportsAccount.of(me, current: character),
                rankingsAccount: RankingsAccount.of(me, current: character),
                dependencies: dependencies,
                showsSelector: session != nil,
                openBis: openBis
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

    init(dependencies: HomeDependencies, session: SessionStore? = nil, openBis: @escaping () -> Void = {}) {
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

/// Home for one character: the answer, Your gear and the tooltip sheet.
struct HomeCharacterScreen: View {
    let character: HomeCharacter
    let finishedSims: [FinishedSim]
    /// `entitlements.server_sims`, for Sims' Server runs section (lane F5).
    let serverSims: Bool
    /// Lane F10: who reads Your reports (crests, the coach card's `log_coach`).
    let reportsAccount: ReportsAccount
    /// Lane F11: who reads Rankings (the aim, your row, the `You` tags, Parse history).
    let rankingsAccount: RankingsAccount
    let dependencies: HomeDependencies
    /// Signed in, the toolbar carries lane F1's selector in place of the name title.
    let showsSelector: Bool
    let openBis: () -> Void

    @State private var store: HomeStore
    @State private var tooltip: ItemTooltipRequest?
    /// Lane F12: the item page, pushed from the tooltip's `More about this item`.
    @State private var itemPage: ItemPageRoute?
    @Environment(\.searchNavigator) private var search
    @State private var showGear = false
    @Environment(SimsRouter.self) private var simsRouter: SimsRouter?
    @Environment(ReportsRouter.self) private var reportsRouter: ReportsRouter?
    /// Lane F10: the Logs row's sub, from the newest report; nil until read.
    @State private var logsLine: String?
    @Environment(RankingsRouter.self) private var rankingsRouter: RankingsRouter?
    /// Lane F11: the Rankings row's sub; nil until read.
    @State private var rankingsLine: String?
    /// Lane F11: the watch list of a character page a player link pushes on Home.
    @State private var watchList: WatchList?
    @State private var showTalents = false
    /// Lane F6: one Talents model per character, shared by Home's row and the pushed screen.
    @State private var talentsModel: TalentsModel?
    @Environment(\.openURL) private var openURL
    @Environment(\.scenePhase) private var scenePhase

    init(
        character: HomeCharacter,
        finishedSims: [FinishedSim],
        serverSims: Bool = false,
        reportsAccount: ReportsAccount = .signedOut,
        rankingsAccount: RankingsAccount = .signedOut,
        dependencies: HomeDependencies,
        showsSelector: Bool = false,
        openBis: @escaping () -> Void
    ) {
        self.character = character
        self.finishedSims = finishedSims
        self.serverSims = serverSims
        self.reportsAccount = reportsAccount
        self.rankingsAccount = rankingsAccount
        self.dependencies = dependencies
        self.showsSelector = showsSelector
        self.openBis = openBis
        _store = State(initialValue: .live(client: dependencies.client, cache: dependencies.cache))
        _talentsModel = State(initialValue: dependencies.talents.map {
            TalentsModel(source: $0, character: TalentCharacter(character), catalog: .bundled())
        })
    }

    var body: some View {
        NavigationStack {
            ScrollView {
                HomeAnswerView(
                    phase: store.phase,
                    presenter: HomePresenter(
                        character: character, finishedSims: finishedSims, time: HomeTime(now: Date()),
                        talentStatus: talentsModel?.phase.presenter?.status,
                        logsLine: dependencies.reports == nil ? nil : logsLine ?? "",
                        rankingsLine: dependencies.rankings == nil ? nil : rankingsLine ?? ""
                    ),
                    build: home?.dataBuild ?? "",
                    offlineStamp: offlineStamp,
                    send: handle
                )
            }
            .background(FSColor.bg)
            .refreshable { await store.refresh(character) }
            .navigationTitle(showsSelector ? "" : character.name)
            .navigationBarTitleDisplayMode(.inline)
            .navigationDestination(isPresented: $showGear) { gear }
            .navigationDestination(isPresented: simsShown) { sims }
            .navigationDestination(isPresented: $showTalents) { talents }
            .navigationDestination(item: $itemPage) { route in
                if let search {
                    ItemPageScreen(route: route, dependencies: search.dependencies) { search.openList($0) }
                }
            }
            .navigationDestination(isPresented: reportsShown) { reports }
            .navigationDestination(isPresented: rankingsShown) { rankings }
            .navigationDestination(item: playerShown) { player($0) }
            .characterSelector(isOn: showsSelector)
            .searchButton()
        }
        .sheet(item: $tooltip) { request in
            ItemTooltipSheet(request: request, catalog: dependencies.catalog, moreAbout: search.map { _ in {
                tooltip = nil
                itemPage = ItemPageRoute(itemID: request.itemID, fallback: request.fallback, classSlug: request.classSlug)
            } })
        }
        .task { await store.load(character) }
        .task { await talentsModel?.load() }
        .task { await loadLogsLine() }
        .task { await loadRankingsLine() }
        .onChange(of: scenePhase) { _, phase in
            if phase == .active { Task { await store.refreshIfStale(character) } }
        }
    }

    private var home: CharacterHome? {
        switch store.phase {
        case let .loaded(home, _): home
        case let .failed(cached): cached
        case .loading: nil
        }
    }

    private var offlineStamp: String? {
        guard case let .loaded(_, since?) = store.phase else { return nil }
        return HomeTime(now: Date()).stamp(since)
    }

    @ViewBuilder
    private var gear: some View {
        if let home {
            GearScreen(home: home, character: character, catalog: dependencies.catalog) { look, wornID in
                tooltip = request(for: look, home: home, sourceLine: nil, wornID: wornID)
            }
        }
    }

    @ViewBuilder
    private var talents: some View {
        if let talentsModel {
            TalentsScreen(character: character, model: talentsModel)
        }
    }

    private func handle(_ intent: HomeIntent) {
        switch intent {
        case let .openItem(look, sourceLine, wornID):
            if let home { tooltip = request(for: look, home: home, sourceLine: sourceLine, wornID: wornID) }
        case .openGear:
            showGear = true
        case .action(.getAddon):
            openURL(HomeLinks.setup)
        case .action(.openPlanner):
            openURL(HomeLinks.planner)
        case .action(.openList):
            openBis()
        case .openNext(.sims):
            simsRouter?.showsList = dependencies.sims != nil
        case .openNext(.talents):
            showTalents = talentsModel != nil
        case .openNext(.logs):
            reportsRouter?.showsList = dependencies.reports != nil
        case .openNext(.rankings):
            rankingsRouter?.encounter = nil
            rankingsRouter?.showsBoard = dependencies.rankings != nil
        case .openNext:
            break
        case .retry:
            Task { await store.refresh(character) }
        }
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

    private func loadRankingsLine() async {
        guard let deps = dependencies.rankings, let current = rankingsAccount.current else { return }
        rankingsLine = await RankingsStore.homeLine(
            service: deps.service, cache: deps.cache, specs: deps.specs, now: deps.now(), current: current
        ) ?? ""
    }

    private func loadLogsLine() async {
        guard let deps = dependencies.reports, reportsAccount.signedIn else { return }
        logsLine = await ReportsStore.homeLine(service: deps.service, cache: deps.cache, now: deps.now)
    }

    private func request(for look: HomeItemLook, home: CharacterHome, sourceLine: String?, wornID: Int?) -> ItemTooltipRequest {
        .forHome(look, home: home, character: character, sourceLine: sourceLine, wornID: wornID)
    }
}

/// Your gear, with the build's items file loaded for the slots the answer does not name.
struct GearScreen: View {
    private static let probe = HomeItemLook(id: 0, name: nil, quality: nil, icon: nil, itemLevel: nil)

    let home: CharacterHome
    let character: HomeCharacter
    let catalog: ItemCatalog
    let open: (HomeItemLook, Int?) -> Void

    @State private var sources: ItemTooltipSources?

    var body: some View {
        let items = sources?.items ?? [:]
        let band = ItemTooltipModel.band(for: .forHome(Self.probe, home: home, character: character), in: sources?.bis)
        let blocker = GearModel.offHandBlocker(home: home, gear: character.gear, items: items, band: band)
        GearView(
            rows: GearModel.rows(home: home, gear: character.gear, items: items, offHandBlockedBy: blocker),
            intro: Copy.gearIntro(band: HomePresenter.bandHead(home, character: character) ?? ""),
            build: home.dataBuild,
            open: open
        )
        .task {
            sources = try? await catalog.sources(for: .forHome(Self.probe, home: home, character: character))
        }
    }
}
