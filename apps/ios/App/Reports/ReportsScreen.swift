import DesignSystem
import Models
import SwiftUI

/// F10.1, pushed from Home's Logs row or a report link: loads the list and pushes F10.2.
struct ReportsScreen: View {
    let dependencies: ReportsDependencies
    let account: ReportsAccount
    /// Home's report navigation: the pushed report lives here, so a link can push it too.
    @Bindable var router: ReportsRouter
    var isSigningIn = false
    var signIn: () -> Void = {}

    @State private var store: ReportsStore
    @Environment(\.openURL) private var openURL

    init(dependencies: ReportsDependencies, account: ReportsAccount, router: ReportsRouter, isSigningIn: Bool = false, signIn: @escaping () -> Void = {}) {
        self.dependencies = dependencies
        self.account = account
        self.router = router
        self.isSigningIn = isSigningIn
        self.signIn = signIn
        _store = State(initialValue: ReportsStore(
            service: dependencies.service, cache: dependencies.cache, now: dependencies.now, signedIn: account.signedIn
        ))
    }

    var body: some View {
        ScrollView {
            ReportsListView(
                phase: store.phase, account: account, pageFailed: store.pageFailed, isSigningIn: isSigningIn,
                now: dependencies.now(), clock: dependencies.clock, actions: actions
            )
        }
        .background(FSColor.bg)
        .refreshable { await store.load() }
        .reportsBar(title: Copy.reportsBack) {
            Button { Task { await store.load() } } label: { Image(systemName: "arrow.clockwise") }
                .accessibilityLabel(Copy.reportsRefresh)
        }
        .navigationDestination(item: $router.report) { route in
            ReportScreen(route: route, dependencies: dependencies, account: account)
        }
        .task { await store.load() }
    }

    private var actions: ReportsActions {
        ReportsActions(
            open: { router.report = ReportRoute(id: $0.id) },
            retry: { Task { await store.load() } },
            loadMore: { Task { await store.loadMore() } },
            getCompanion: { openURL(ReportsLinks.companion) },
            signIn: signIn
        )
    }
}

/// F10.2, pushed from Your reports, Your raid night or `/app/report/{id}`. A live report re-reads
/// itself every 20 s while the screen shows.
struct ReportScreen: View {
    let route: ReportRoute
    let dependencies: ReportsDependencies

    @State private var store: ReportStore
    @State private var meters: MetersTarget?
    @State private var build: String?
    @Environment(\.openURL) private var openURL
    @Environment(\.scenePhase) private var scenePhase

    init(route: ReportRoute, dependencies: ReportsDependencies, account: ReportsAccount) {
        self.route = route
        self.dependencies = dependencies
        _store = State(initialValue: ReportStore(
            id: route.id, account: account, service: dependencies.service, cache: dependencies.cache, now: dependencies.now
        ))
    }

    var body: some View {
        ScrollView {
            ReportView(
                phase: store.phase, parses: store.parses, coach: store.coach, now: dependencies.now(), clock: dependencies.clock,
                actions: ReportActions(
                    openFight: { row in meters = MetersTarget.of(row, reportID: route.id) },
                    retry: { Task { await store.load() } },
                    retryParses: { Task { await store.retryParses() } }
                )
            )
        }
        .background(FSColor.bg)
        .refreshable { await store.load() }
        .reportsBar(title: store.report?.zone ?? Copy.reportsBack) {
            Button { openURL(ReportsLinks.web(route.id)) } label: { Image(systemName: "arrow.up.right.square") }
                .accessibilityLabel(Copy.reportOpenWeb)
        }
        .sheet(item: $meters) { target in
            MetersSheet(target: target, dependencies: dependencies, build: build)
                .presentationDetents([.large])
                .presentationDragIndicator(.visible)
        }
        .task {
            build = await dependencies.iconBuild()
            await store.load()
        }
        .task(id: scenePhase) { await poll() }
    }

    /// Every 20 s while foreground and the report is live; the task ends when the screen leaves.
    private func poll() async {
        guard scenePhase == .active else { return }
        while !Task.isCancelled {
            try? await Task.sleep(for: ReportsLive.pollInterval)
            guard !Task.isCancelled, ReportsLive.polls(store.report) else { continue }
            await store.loadReport()
        }
    }
}

/// F10.3 bound to its store.
struct MetersSheet: View {
    let dependencies: ReportsDependencies
    let build: String?
    @State private var store: MetersStore

    init(target: MetersTarget, dependencies: ReportsDependencies, build: String?) {
        self.dependencies = dependencies
        self.build = build
        _store = State(initialValue: MetersStore(target: target, service: dependencies.service))
    }

    var body: some View {
        MetersView(
            target: store.target, metric: store.metric, phase: store.phase, showAll: store.showAll,
            specs: dependencies.specs, build: build,
            select: { metric in Task { await store.select(metric) } },
            expand: { store.showAll = true },
            retry: { Task { await store.load() } }
        )
        .task { await store.load() }
    }
}

extension View {
    /// The F10 bar (board F10.1, F10.2): back and one trailing glass button; the condensed caps
    /// title is in the content, and `title` names the back button of the screen pushed over it.
    func reportsBar(title: String, @ViewBuilder trailing: () -> some View) -> some View {
        let button = trailing()
        return navigationTitle(title)
            .navigationBarTitleDisplayMode(.inline)
            .toolbar(removing: .title)
            .toolbar { ToolbarItem(placement: .topBarTrailing) { button } }
    }
}
