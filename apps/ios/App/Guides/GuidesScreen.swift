import Auth
import DesignSystem
import Models
import SwiftUI

/// The Guides tab (spec F9): the segments over Guides, Tier lists and What changed, and the
/// stack of class pages, guides and changelog entries.
struct GuidesScreen: View {
    let dependencies: GuidesDependencies
    /// Lane F1's account, nil while its feature switch is off.
    let session: SessionStore?
    /// `guide_tiers`: the Tier lists segment.
    let tiersOn: Bool
    @Bindable var router: GuidesRouter
    /// A guide's Gear row: the BiS tab on that list.
    let openBis: (GuideGear) -> Void

    @State private var home: GuidesHomeStore
    @State private var tiers: TiersStore
    @State private var changes: ChangesStore
    @AppStorage(GuidesSegment.storageKey) private var storedSegment = GuidesSegment.guides.rawValue

    init(dependencies: GuidesDependencies, session: SessionStore?, tiersOn: Bool, router: GuidesRouter, openBis: @escaping (GuideGear) -> Void) {
        self.dependencies = dependencies
        self.session = session
        self.tiersOn = tiersOn
        self.router = router
        self.openBis = openBis
        _home = State(initialValue: GuidesHomeStore(source: dependencies.source))
        _tiers = State(initialValue: TiersStore(source: dependencies.source, choice: .fallback))
        _changes = State(initialValue: ChangesStore(source: dependencies.source, now: dependencies.now))
    }

    private var character: AccountCharacter? { session?.currentCharacter }
    private var segment: GuidesSegment { GuidesSegment.resolved(storedSegment, tiersOn: tiersOn) }

    var body: some View {
        NavigationStack(path: $router.path) {
            GuidesRootView(
                segment: segment, tiersOn: tiersOn, homePhase: home.phase, character: character,
                tiersPhase: tiers.phase, choice: tiers.choice, yourSpec: TiersLogic.specKey(character, specs: dependencies.specs),
                dates: changes.dates, changelog: changes.changelog, now: dependencies.now(), actions: rootActions
            )
            .refreshable { await load(segment) }
            .navigationTitle(Copy.guidesTitle)
            .navigationBarTitleDisplayMode(.inline)
            .toolbar(removing: .title)
            .characterSelector(isOn: session != nil)
            .toolbarBackground(FSColor.bg, for: .navigationBar)
            .navigationDestination(for: GuidesRoute.self, destination: destination)
        }
        .task { await home.load() }
        .task(id: segment) { await load(segment) }
        .task(id: character?.key) { tiers.choice = TierChoice.defaults(character, specs: dependencies.specs) }
    }

    private var rootActions: GuidesRootActions {
        GuidesRootActions(
            selectSegment: { storedSegment = $0.rawValue },
            openGuide: { router.path.append(.guide($0)) },
            openClass: { router.path.append(.classPage($0)) },
            openEntry: { router.path.append(.entry($0.id)) },
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

    @ViewBuilder
    private func destination(_ route: GuidesRoute) -> some View {
        switch route {
        case let .classPage(slug):
            ScrollView {
                if case let .loaded(loaded, since) = home.phase {
                    GuideClassView(classSlug: slug, home: loaded, offlineSince: since, now: dependencies.now()) {
                        router.path.append(.guide($0.spec))
                    }
                } else {
                    GuidesHomeSkeleton().padding(GuidesLayout.gutter)
                }
            }
            .background(FSColor.bg)
            .navigationBarTitleDisplayMode(.inline)
        case let .guide(spec):
            GuideScreen(spec: spec, dependencies: dependencies, openGear: openBis) { home.refreshSaved() }
        case let .entry(id):
            ScrollView {
                if let entry = changes.changelog.value?.entries.first(where: { $0.id == id }) {
                    ChangelogEntryView(entry: entry)
                }
            }
            .background(FSColor.bg)
            .navigationBarTitleDisplayMode(.inline)
        }
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
                if segment == .guides, case let .loaded(_, since) = homePhase {
                    GuidesOfflineStrip(since: since, now: now)
                }
                SimsTitle(text: Copy.guidesTitle)
                SegmentBar(items: GuidesSegment.available(tiersOn: tiersOn), selection: segment, title: \.title, select: actions.selectSegment)
                    .padding(.horizontal, GuidesLayout.gutter)
                    .padding(.bottom, 14)
                region.padding(.horizontal, GuidesLayout.gutter)
            }
            .frame(maxWidth: .infinity, alignment: .leading)
            .padding(.bottom, GuidesLayout.tabBarClearance)
        }
        .background(FSColor.bg)
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

    init(spec: String, dependencies: GuidesDependencies, openGear: @escaping (GuideGear) -> Void, onLoaded: @escaping () -> Void) {
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
            )
        )
        .background(FSColor.bg)
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
    /// A guide's bar (spec F9.2): back, the spec in the class colour, share.
    func guideBar(_ guide: GuideFile?) -> some View {
        navigationBarTitleDisplayMode(.inline)
            .toolbar {
                if let guide {
                    ToolbarItem(placement: .principal) {
                        Text(guide.specName).font(ScaledType.name).foregroundStyle(ClassColor.of(guide.classSlug))
                    }
                    ToolbarItem(placement: .topBarTrailing) {
                        ShareLink(item: GuideCopy.shareURL(guide)) { Image(systemName: "square.and.arrow.up") }
                            .accessibilityLabel(Copy.guidesShare)
                    }
                }
            }
    }
}
