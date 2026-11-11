import Auth
import DesignSystem
import Models
import SwiftUI

/// Guides, Tier lists and What changed (spec F9), pushed from the profile menu under the tab it
/// was opened on (flows spec 2): the segments, preselected by the menu row, and the class pages,
/// guides and changelog entries it pushes in turn.
struct GuidesScreen: View {
    let dependencies: GuidesDependencies
    /// Lane F1's account, nil while its feature switch is off.
    let session: SessionStore?
    /// `guide_tiers`: the Tier lists segment.
    let tiersOn: Bool
    /// The segment the menu row opens.
    let initialSegment: GuidesSegment
    @Bindable var router: GuidesRouter
    /// A guide's Gear row: the Gear tab on that list.
    let openBis: (GuideGear) -> Void

    @State private var home: GuidesHomeStore
    @State private var tiers: TiersStore
    @State private var changes: ChangesStore
    @State private var segment: GuidesSegment
    @State private var pushed: GuidesRoute?

    init(
        dependencies: GuidesDependencies, session: SessionStore?, tiersOn: Bool, initialSegment: GuidesSegment = .guides,
        router: GuidesRouter, openBis: @escaping (GuideGear) -> Void
    ) {
        self.dependencies = dependencies
        self.session = session
        self.tiersOn = tiersOn
        self.initialSegment = initialSegment
        self.router = router
        self.openBis = openBis
        _home = State(initialValue: GuidesHomeStore(source: dependencies.source))
        _tiers = State(initialValue: TiersStore(source: dependencies.source, choice: .fallback))
        _changes = State(initialValue: ChangesStore(source: dependencies.source, now: dependencies.now))
        _segment = State(initialValue: GuidesSegment.resolved(initialSegment.rawValue, tiersOn: tiersOn))
    }

    private var character: AccountCharacter? { session?.currentCharacter }

    var body: some View {
        GuidesRootView(
            segment: segment, tiersOn: tiersOn, homePhase: home.phase, character: character,
            tiersPhase: tiers.phase, choice: tiers.choice, yourSpec: TiersLogic.specKey(character, specs: dependencies.specs),
            dates: changes.dates, changelog: changes.changelog, now: dependencies.now(), actions: rootActions
        )
        .refreshable { await load(segment) }
        .pushedBandBar(Copy.guidesTitle)
        .toolbar {
            if segment == .tiers {
                ToolbarItem(placement: .topBarTrailing) {
                    ShareButton(ShareText.tiers(tiers.choice.role), accessibility: Copy.shareAction)
                }
            }
        }
        .toolbarBackground(FSColor.bg, for: .navigationBar)
        .navigationDestination(item: $pushed) { route in
            GuidesRouteScreen(route: route, destination: destination)
        }
        .task { await home.load() }
        .task(id: segment) { await load(segment) }
        .task(id: character?.key) { tiers.choice = TierChoice.defaults(character, specs: dependencies.specs) }
        .onAppear { takePending() }
        .onChange(of: router.pending) { takePending() }
    }

    /// A guide link that arrived for this screen.
    private func takePending() {
        if let route = router.takePending() { pushed = route }
    }

    private var rootActions: GuidesRootActions {
        GuidesRootActions(
            selectSegment: { segment = $0 },
            openGuide: { pushed = .guide($0) },
            openClass: { pushed = .classPage($0) },
            openEntry: { pushed = .entry($0.id) },
            selectTiers: { tiers.choice = $0 },
            retry: { Task { await load(segment) } }
        )
    }

    private func load(_ segment: GuidesSegment) async {
        switch segment {
        case .guides: await home.load()
        case .tiers: await tiers.load()
        case .changes: await changes.load()
        }
    }

    /// A guide's name from the index (`Elemental Shaman`), for its head before the guide is drawn.
    private func guideTitle(_ spec: String) -> String {
        home.phase.value.flatMap { GuidesIndexLogic.spec(spec, in: $0.index)?.title } ?? Copy.guidesTitle
    }

    /// One pushed screen; `push` opens the next one over it.
    @ViewBuilder
    private func destination(_ route: GuidesRoute, push: @escaping (GuidesRoute) -> Void) -> some View {
        switch route {
        case let .classPage(slug):
            ScrollView {
                if case let .loaded(loaded, since) = home.phase {
                    GuideClassView(classSlug: slug, home: loaded, offlineSince: since, now: dependencies.now()) {
                        push(.guide($0.spec))
                    }
                } else {
                    GuidesHomeSkeleton().padding(GuidesLayout.gutter)
                }
            }
            .screenGround()
            .pushedBandBar(GameClass(rawValue: slug)?.name ?? slug)
        case let .guide(spec):
            GuideScreen(spec: spec, title: guideTitle(spec), dependencies: dependencies, openGear: openBis) { home.refreshSaved() }
        case let .entry(id):
            ScrollView {
                if let entry = changes.changelog.value?.entries.first(where: { $0.id == id }) {
                    ChangelogEntryView(entry: entry)
                }
            }
            .screenGround()
            .pushedBandBar(changes.changelog.value?.entries.first(where: { $0.id == id })?.title ?? Copy.guidesTitle)
        }
    }
}

/// A pushed Guides screen that can push one more over itself (a class page opens a guide), so the
/// stack needs no typed path and works under any tab.
private struct GuidesRouteScreen<Content: View>: View {
    let route: GuidesRoute
    let destination: (GuidesRoute, @escaping (GuidesRoute) -> Void) -> Content
    @State private var next: GuidesRoute?

    var body: some View {
        destination(route) { next = $0 }
            .navigationDestination(item: $next) { GuidesRouteScreen(route: $0, destination: destination) }
    }
}

/// What the root hands back to its screen.
struct GuidesRootActions {
    var selectSegment: (GuidesSegment) -> Void = { _ in }
    var openGuide: (String) -> Void = { _ in }
    var openClass: (String) -> Void = { _ in }
    var openEntry: (ChangelogEntry) -> Void = { _ in }
    var selectTiers: (TierChoice) -> Void = { _ in }
    var retry: () -> Void = {}
}

/// The tab's root as drawn from its stores' phases: stateless, so snapshots drive every state.
/// The segments stay live while a region loads (F9.7).
struct GuidesRootView: View {
    let segment: GuidesSegment
    let tiersOn: Bool
    let homePhase: StaticPhase<GuidesHome>
    let character: AccountCharacter?
    var tiersPhase: StaticPhase<TiersFile> = .loading
    var choice: TierChoice = .fallback
    var yourSpec: String?
    var dates: [KeyDateRow]?
    var changelog: StaticPhase<Changelog> = .loading
    let now: Date
    var actions = GuidesRootActions()

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 0) {
                BandHead(title: Copy.guidesTitle)
                if segment == .guides, case let .loaded(_, since) = homePhase {
                    GuidesOfflineStrip(since: since, now: now)
                }
                SegmentBar(items: GuidesSegment.available(tiersOn: tiersOn), selection: segment, title: \.title, select: actions.selectSegment)
                    .padding(.horizontal, GuidesLayout.gutter)
                    .padding(.bottom, 14)
                region.padding(.horizontal, GuidesLayout.gutter)
            }
            .frame(maxWidth: .infinity, alignment: .leading)
            .padding(.bottom, GuidesLayout.tabBarClearance)
        }
        .screenGround()
    }

    @ViewBuilder
    private var region: some View {
        switch segment {
        case .guides:
            GuidesHomeView(
                phase: homePhase, character: character,
                openGuide: { actions.openGuide($0.spec) }, openClass: { actions.openClass($0.classSlug) }, retry: actions.retry
            )
        case .tiers:
            TiersView(
                phase: tiersPhase, choice: choice, yourSpec: yourSpec, select: actions.selectTiers,
                open: { actions.openGuide($0.spec) }, retry: actions.retry
            )
        case .changes:
            ChangesView(dates: dates, changelog: changelog, open: actions.openEntry, retry: actions.retry)
        }
    }
}

/// F9.2 bound to its store: the saved copy at once, the new-data check on foreground.
struct GuideScreen: View {
    /// Lane F13: a guide's build opens the editor as one legal order, `Cancel` in the bar.
    private var openPlanner: ((String, String) -> Void)? {
        guard let planner, planner.isOn else { return nil }
        return { classSlug, code in
            planner.openEditor(PlannerLaunch(classSlug: classSlug, source: .code(code), origin: .guide))
        }
    }

    let dependencies: GuidesDependencies
    let openGear: (GuideGear) -> Void
    let onLoaded: () -> Void

    @State private var store: GuideStore
    @Environment(\.scenePhase) private var scenePhase
    /// Lane F13: the talents section's planner row.
    @Environment(PlannerRouter.self) private var planner: PlannerRouter?

    /// The guide's name from the index, for the head before the guide itself is drawn.
    let title: String

    init(
        spec: String, title: String, dependencies: GuidesDependencies, openGear: @escaping (GuideGear) -> Void,
        onLoaded: @escaping () -> Void
    ) {
        self.title = title
        self.dependencies = dependencies
        self.openGear = openGear
        self.onLoaded = onLoaded
        _store = State(initialValue: GuideStore(spec: spec, source: dependencies.source))
    }

    var body: some View {
        GuideReaderView(
            phase: store.phase, newData: store.newData, now: dependencies.now(),
            actions: GuideActions(
                openGear: openGear,
                retry: { Task { await store.load() } },
                update: { Task { await store.update() } },
                openPlanner: openPlanner
            ),
            title: title
        )
        .screenGround()
        .guideBar(store.phase.value?.guide)
        .task {
            await store.load()
            onLoaded()
        }
        .task(id: scenePhase) {
            guard scenePhase == .active else { return }
            await store.checkForNewData()
        }
    }
}

extension View {
    /// A guide's bar (spec F9.2, N3): back and share on the clear bar; the band head names the guide.
    func guideBar(_ guide: GuideFile?) -> some View {
        pushedBandBar(guide.map(GuideCopy.title) ?? Copy.guidesTitle)
            .toolbar {
                if let guide {
                    ToolbarItem(placement: .topBarTrailing) {
                        ShareButton(ShareText.guide(guide), accessibility: Copy.guidesShare)
                    }
                }
            }
    }
}
