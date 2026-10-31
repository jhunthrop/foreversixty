import DesignSystem
import Models
import SwiftUI

/// A result to push: its id, and the list row when the list supplied it.
struct SimResultRoute: Hashable, Identifiable {
    let id: String
    let row: SimRow?

    static func == (lhs: SimResultRoute, rhs: SimResultRoute) -> Bool { lhs.id == rhs.id }
    func hash(into hasher: inout Hasher) { hasher.combine(id) }
}

/// F5.1, pushed from Home's Sims row: loads, polls in-flight runs every 5 s while it shows, and
/// pushes F5.3 for a finished row.
struct SimsScreen: View {
    let dependencies: SimsDependencies
    /// `entitlements.server_sims` from the account's `/v1/me`.
    let entitled: Bool
    /// Home's sim navigation: the pushed result lives here, so a link can push it too.
    @Bindable var router: SimsRouter

    @State private var store: SimsStore
    @State private var build: String?
    @State private var mayLinkToPlans = false
    @Environment(\.openURL) private var openURL
    @Environment(\.scenePhase) private var scenePhase

    init(dependencies: SimsDependencies, entitled: Bool, router: SimsRouter) {
        self.dependencies = dependencies
        self.entitled = entitled
        self.router = router
        _store = State(initialValue: SimsStore(service: dependencies.service, cache: dependencies.cache, now: dependencies.now))
    }

    var body: some View {
        ScrollView {
            SimsListView(
                phase: store.phase, entitled: entitled, progress: store.progress, rerun: store.rerun,
                pageFailed: store.pageFailed, mayLinkToPlans: mayLinkToPlans, specs: dependencies.specs,
                build: build, now: dependencies.now(), actions: actions
            )
        }
        .background(FSColor.bg)
        .refreshable { await store.load() }
        .navigationTitle(Copy.simsTitle)
        .navigationBarTitleDisplayMode(.inline)
        .navigationDestination(item: $router.result) { route in
            SimResultScreen(route: route, dependencies: dependencies) { id in await store.markSeen(id) }
        }
        .task {
            build = await dependencies.iconBuild()
            mayLinkToPlans = await dependencies.mayLinkToPlans()
            await store.load()
        }
        .task(id: scenePhase) { await poll() }
    }

    private var actions: SimsActions {
        SimsActions(
            open: { row in row.simId.map { router.result = SimResultRoute(id: $0, row: row) } },
            runAgain: { id in Task { await store.runAgain(id) } },
            retry: { Task { await store.load() } },
            loadMore: { Task { await store.loadMore() } },
            openPlans: { openURL(SimsDependencies.plansURL) }
        )
    }

    /// Every 5 s while foreground and something is in flight; the task ends when the screen leaves.
    private func poll() async {
        guard scenePhase == .active else { return }
        while !Task.isCancelled {
            try? await Task.sleep(for: ServerRuns.pollInterval)
            guard !Task.isCancelled, store.isPolling else { continue }
            await store.pollOnce()
        }
    }
}

/// F5.3, pushed from Sims, a `sim.finished` push or `/app/sim/{id}`.
struct SimResultScreen: View {
    let route: SimResultRoute
    let dependencies: SimsDependencies

    @State private var store: SimResultStore
    @State private var looks: [Int: SimGearLook] = [:]
    @State private var build: String?
    @State private var tooltip: ItemTooltipRequest?
    @Environment(\.openURL) private var openURL

    init(route: SimResultRoute, dependencies: SimsDependencies, markSeen: @escaping (String) async -> Void) {
        self.route = route
        self.dependencies = dependencies
        _store = State(initialValue: SimResultStore(
            id: route.id, row: route.row, service: dependencies.service, cache: dependencies.cache,
            specs: dependencies.specs, now: dependencies.now, markSeen: markSeen
        ))
    }

    var body: some View {
        ScrollView {
            SimResultView(
                phase: store.phase, looks: looks, build: build, now: dependencies.now(),
                openItem: { cell in tooltip = request(for: cell) },
                openBreakdown: { openURL($0) },
                retry: { Task { await store.load() } }
            )
        }
        .background(FSColor.bg)
        .navigationTitle(title)
        .navigationBarTitleDisplayMode(.inline)
        .sheet(item: $tooltip) { request in ItemTooltipSheet(request: request, catalog: dependencies.catalog) }
        .task {
            build = await dependencies.iconBuild()
            await store.load()
            await loadLooks()
        }
    }

    private var title: String {
        if case let .loaded(model, _) = store.phase { return model.title }
        return Copy.simsTitle
    }

    private var model: SimResultModel? {
        if case let .loaded(model, _) = store.phase { return model }
        return nil
    }

    private func request(for cell: SimGearCell) -> ItemTooltipRequest? {
        guard let model, let build, let look = looks[cell.itemID] else { return nil }
        return SimGearLooks.request(cell: cell, look: look, model: model, build: build)
    }

    private func loadLooks() async {
        guard let model, let build else { return }
        looks = await SimGearLooks.load(model: model, build: build, catalog: dependencies.catalog)
    }
}

/// The gear grid's names and icons, from the build's items file through the shared catalog.
enum SimGearLooks {
    static func load(model: SimResultModel, build: String, catalog: ItemCatalog) async -> [Int: SimGearLook] {
        guard let probe = model.gear.first,
              let sources = try? await catalog.sources(for: request(cell: probe, look: nil, model: model, build: build))
        else { return [:] }
        return Dictionary(uniqueKeysWithValues: model.gear.compactMap { cell in
            sources.items[cell.itemID].map { (cell.itemID, SimGearLook(name: $0.name, icon: $0.icon, quality: $0.quality)) }
        })
    }

    static func request(cell: SimGearCell, look: SimGearLook?, model: SimResultModel, build: String) -> ItemTooltipRequest {
        ItemTooltipRequest(
            itemID: cell.itemID,
            fallback: HomeItemLook(id: cell.itemID, name: look?.name, quality: look?.quality, icon: look?.icon, itemLevel: nil),
            build: build, classSlug: model.classSlug, spec: nil, specName: model.specName,
            band: nil, bandLabel: nil, preset: nil, faction: nil, sourceLine: nil, wornID: nil
        )
    }
}
