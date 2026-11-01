import DesignSystem
import Models
import SwiftUI

/// The site page `How is this calculated?` opens.
enum RankingsLinks {
    static let ratings = URL(string: "https://foreversixty.gg/ratings")!
}

/// F11.1, pushed from Home's Rankings row or a rankings link: the board, its filters sheet, the
/// F3.2 picker signed out, and the character pages and reports pushed over it.
struct RankingsScreen: View {
    let dependencies: RankingsDependencies
    let account: RankingsAccount
    /// F10's reader, for a report of one of your characters (with Your parses); others open read-only.
    let reportsAccount: ReportsAccount

    @State private var store: RankingsStore
    @State private var watchList: WatchList
    @State private var showsFilters = false
    @State private var showsPicker = false
    @State private var player: PlayerRoute?
    @State private var report: OpenReport?
    @State private var build: String?

    init(dependencies: RankingsDependencies, account: RankingsAccount, reportsAccount: ReportsAccount = .signedOut, encounter: String? = nil) {
        self.dependencies = dependencies
        self.account = account
        self.reportsAccount = reportsAccount
        _store = State(initialValue: RankingsStore(
            service: dependencies.service, cache: dependencies.cache, specs: dependencies.specs, now: dependencies.now,
            account: account, encounter: encounter
        ).connected())
        _watchList = State(initialValue: WatchList(service: dependencies.service, signedIn: account.signedIn))
    }

    var body: some View {
        ScrollView {
            RankingsView(
                phase: store.phase, filters: store.filters, encounterName: store.encounterName, account: account,
                watching: watchList.keys, pageFailed: store.pageFailed, now: dependencies.now(), actions: actions
            )
        }
        .background(FSColor.bg)
        .refreshable { await store.load() }
        .reportsBar(title: Copy.rankTitle) { EmptyView() }
        .haptic(.success, trigger: watchList.confirmed)
        .sheet(isPresented: $showsFilters) {
            if let filters = store.filters, let setup = store.setup {
                FiltersSheet(
                    filters: filters, setup: setup,
                    apply: { next in Task { await store.apply(next) } },
                    reset: { Task { await store.reset() } }
                )
            }
        }
        .sheet(isPresented: $showsPicker) { picker }
        .navigationDestination(item: $player) { route in
            CharacterScreen(route: route, dependencies: dependencies, account: account, reportsAccount: reportsAccount, watchList: watchList)
        }
        .navigationDestination(item: $report) { open in
            if let reports = dependencies.reports {
                ReportScreen(route: open.route, dependencies: reports, account: open.yours ? reportsAccount : .signedOut)
            }
        }
        .task {
            build = await dependencies.iconBuild()
            await store.load()
            await watchList.load()
        }
    }

    private var actions: RankingsActions {
        RankingsActions(
            open: { row in
                if let ref = CharacterRef(key: row.player.key, name: row.player.name) { player = PlayerRoute(ref: ref) }
            },
            retry: { Task { await store.load() } },
            loadMore: { Task { await store.loadMore() } },
            openFilters: { showsFilters = store.setup != nil },
            chooseSpec: { showsPicker = true },
            toggleWatch: { row in
                if let ref = CharacterRef(key: row.player.key, name: row.player.name) { Task { await watchList.toggle(ref) } }
            },
            openReport: dependencies.reports == nil ? nil : { row in
                report = OpenReport(route: ReportRoute(id: row.reportId), yours: account.characterKeys.contains(row.player.key))
            }
        )
    }

    /// Signed out: the BiS picker's class and spec (spec F3.2) aim the board.
    @ViewBuilder
    private var picker: some View {
        let catalog = dependencies.specs
        let filters = store.filters
        let current = catalog.spec(classSlug: filters?.classSlug ?? "", named: filters?.spec ?? "")
        let selection = current.map {
            BisSelection(classSlug: $0.classSlug, spec: $0.spec, range: BisSelection.signedOutDefault(catalog).range, faction: BisSelection.factions[0])
        } ?? BisSelection.signedOutDefault(catalog)
        BisPickerSheet(context: BisContext(selection: selection), catalog: catalog, build: build, showsListChoices: false) { picked in
            Task { await store.choose(classSlug: picked.classSlug, spec: picked.spec) }
        }
    }
}

/// F11.4, pushed from a board row or a player link: the page, the Watch star and share, the
/// tooltip sheet for a socket, and a recent fight's report.
struct CharacterScreen: View {
    let route: PlayerRoute
    let dependencies: RankingsDependencies
    let account: RankingsAccount
    var reportsAccount: ReportsAccount = .signedOut
    let watchList: WatchList

    @State private var store: CharacterStore?
    @State private var build: String?
    @State private var tooltip: ItemTooltipRequest?
    @State private var report: OpenReport?
    @Environment(\.openURL) private var openURL
    @Environment(\.dismiss) private var dismiss
    @Environment(\.searchNavigator) private var searchNavigator

    var body: some View {
        ScrollView {
            CharacterView(
                ref: route.ref, phase: store?.phase ?? .loading, items: store?.items ?? [:], account: account,
                specs: dependencies.specs, build: build ?? "", watchFailed: watchList.failedKey == route.ref.key,
                now: dependencies.now(), clock: dependencies.clock, actions: actions
            )
        }
        .background(FSColor.bg)
        .refreshable { await store?.load() }
        .characterBar(
            title: route.ref.name, share: shareItem,
            star: account.signedIn && showsPage
                ? WatchStarState(name: displayName, watching: watchList.isWatching(route.ref.key), offline: isOffline)
                : nil
        ) { Task { await watchList.toggle(route.ref) } }
        .haptic(.success, trigger: watchList.confirmed)
        .sheet(item: $tooltip) { request in
            ItemTooltipSheet(request: request, catalog: dependencies.catalog)
        }
        .navigationDestination(item: $report) { open in
            if let reports = dependencies.reports {
                ReportScreen(route: open.route, dependencies: reports, account: open.yours ? reportsAccount : .signedOut)
            }
        }
        .task {
            build = await dependencies.iconBuild()
            let store = store ?? CharacterStore(
                ref: route.ref, service: dependencies.service, cache: dependencies.cache, now: dependencies.now,
                items: dependencies.items(build: build)
            ).connected()
            self.store = store
            await store.load()
            if watchList.keys.isEmpty { await watchList.load() }
        }
    }

    private var profile: CharacterProfile? {
        if case let .loaded(content) = store?.phase { return content.profile }
        return nil
    }

    private var showsPage: Bool { profile != nil }

    /// `Lightningbolltz, Elemental Shaman, level 60` (spec §5.7 Share).
    private var shareItem: ShareItem {
        let head = profile?.character
        return ShareText.character(
            key: route.ref.key, name: head?.name ?? route.ref.name,
            spec: profile.flatMap { CharacterPage.spec($0, specs: dependencies.specs) },
            className: head?.className, level: head?.level
        )
    }
    private var displayName: String { profile?.character.name ?? route.ref.name }

    private var isOffline: Bool {
        if case let .loaded(content) = store?.phase { return content.offlineSince != nil }
        return false
    }

    private var actions: CharacterActions {
        CharacterActions(
            retry: { Task { await store?.load() } },
            openHow: { openURL(RankingsLinks.ratings) },
            openSocket: { slot in
                guard let look = slot.look, let profile else { return }
                tooltip = ItemTooltipRequest(
                    itemID: look.id, fallback: look, build: build ?? "", classSlug: profile.character.className?.lowercased() ?? "",
                    spec: profile.gearSpec, specName: CharacterPage.spec(profile, specs: dependencies.specs), band: nil, bandLabel: nil,
                    preset: nil, faction: nil, sourceLine: nil, wornID: look.id
                )
            },
            openReport: dependencies.reports == nil ? nil : { line in
                report = OpenReport(route: ReportRoute(id: line.reportID), yours: account.characterKeys.contains(route.ref.key))
            },
            openHistory: { openURL(route.ref.historyURL) },
            searchAgain: {
                // Lane F15: Search with the name in the field, when Search is on; else back.
                if let searchNavigator { searchNavigator.open(query: route.ref.name) } else { dismiss() }
            }
        )
    }
}

/// A report opened from F11: with Your parses only for one of your characters (spec F11.4).
struct OpenReport: Hashable, Identifiable {
    let route: ReportRoute
    let yours: Bool
    var id: String { route.id }
}

/// The star's state in the bar; nil draws no star (signed out, or no page to watch).
struct WatchStarState: Equatable {
    let name: String
    let watching: Bool
    let offline: Bool
}

extension View {
    /// The character page's bar (board F11.4): back, then the Watch star and share, each a
    /// 44 pt target drawn without the shared glass ground so the two read as one group.
    func characterBar(title: String, share: ShareItem, star: WatchStarState?, toggle: @escaping () -> Void) -> some View {
        navigationTitle(title)
            .navigationBarTitleDisplayMode(.inline)
            .toolbar(removing: .title)
            .toolbar {
                if let star {
                    ToolbarItem(placement: .topBarTrailing) {
                        WatchStar(name: star.name, watching: star.watching, offline: star.offline, toggle: toggle)
                    }
                    .sharedBackgroundVisibility(.hidden)
                }
                ToolbarItem(placement: .topBarTrailing) {
                    ShareButton(item: share) {
                        Image(systemName: "square.and.arrow.up").foregroundStyle(FSColor.strong)
                            .frame(width: RankingsLayout.star, height: RankingsLayout.star)
                    }
                    .accessibilityLabel(Copy.charShare)
                }
                .sharedBackgroundVisibility(.hidden)
            }
    }
}

/// The Watch star (spec F11.4, §7): 44 pt; outline muted, filled gold; 50% and disabled offline.
struct WatchStar: View {
    let name: String
    let watching: Bool
    let offline: Bool
    let toggle: () -> Void

    var body: some View {
        Button(action: toggle) {
            Image(systemName: watching ? "star.fill" : "star")
                .foregroundStyle(watching ? FSColor.gold : FSColor.muted)
                .frame(width: RankingsLayout.star, height: RankingsLayout.star)
                .contentShape(Rectangle())
        }
        .disabled(offline)
        .opacity(offline ? RankingsLayout.disabledOpacity : 1)
        .accessibilityLabel(watching ? Copy.charUnwatchA11y(name: name) : Copy.charWatchA11y(name: name))
        .accessibilityAddTraits(watching ? .isSelected : [])
    }
}
