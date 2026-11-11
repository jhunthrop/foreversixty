import Auth
import DesignSystem
import Networking
import SwiftUI

/// The Gear tab (flows spec 4.3, F5). Signed in: the character's list against what it wears,
/// Where to go tonight and every slot, over `BisModel`; Change opens the picker. Signed out: the
/// class rows, then the class's list (spec F3) with the picker for the spec and the level range.
/// Every item opens the item sheet.
struct BisScreen: View {
    @State private var model: BisModel
    @State private var sheet: ItemSheetRequest?
    /// Signed out: the class's list, pushed over the class rows.
    @State private var listShown = false
    /// Lane F12: the item page, pushed from the sheet's `More about this item`.
    @State private var itemPage: ItemPageRoute?
    @Environment(\.searchNavigator) private var search
    @Environment(\.scenePhase) private var scenePhase
    private let catalog: ItemCatalog
    /// Who in the guild needs an item; nil signed out.
    private let needs: (any ItemNeedsSource)?
    /// Lane F1's account, nil while its feature switch is off.
    private let session: SessionStore?
    /// Lists other tabs and links ask for (a guide's Gear row, `/app/bis`).
    private let router: BisRouter?

    init(
        model: BisModel, catalog: ItemCatalog, needs: (any ItemNeedsSource)? = nil, session: SessionStore? = nil,
        router: BisRouter? = nil
    ) {
        _model = State(initialValue: model)
        self.catalog = catalog
        self.needs = needs
        self.session = session
        self.router = router
    }

    /// The model restarts whenever the session's current character changes; nil while a signed-in
    /// session's characters are still on their way, so the signed-out first visit never fires early.
    /// The selected character's Battle.net portrait for the paperdoll's medallion (the seal without one).
    @State private var portrait: UIImage?
    @Environment(\.portraitSource) private var portraitSource

    private var avatarURL: URL? {
        session?.currentCharacter?.avatarUrl.flatMap(URL.init(string:))
    }

    private var owner: String?? {
        guard let session else { return .some(nil) }
        switch session.phase {
        case .importing: return nil
        case let .signedIn(me) where me == nil: return nil
        default: return .some(session.currentCharacter?.key)
        }
    }

    var body: some View {
        @Bindable var model = model
        NavigationStack {
            root
                .refreshable { await model.refresh() }
                .modifier(BisSearchMode(text: $model.searchText, isUniversal: search != nil))
                // Round 5 on iOS (m-ios-gear): the tab's large title on the band, under the profile button.
                .navigationTitle(AppTab.gear.title)
                .toolbarTitleDisplayMode(.large)
                .tabRootChrome()
                .navigationDestination(isPresented: $listShown) {
                    list.refreshable { await model.refresh() }.pushedBandBar(Copy.bisTitle)
                }
                .navigationDestination(item: $itemPage) { route in
                    if let search {
                        ItemPageScreen(route: route, dependencies: search.dependencies) { search.openList($0) }
                    }
                }
        }
        .sheet(isPresented: $model.isPickerPresented) {
            BisPickerSheet(context: model.context, catalog: model.catalog, build: model.build) { selection in
                Task { await model.select(selection) }
            }
        }
        .sheet(item: $sheet) { request in
            ItemSheetScreen(
                request: request, catalog: catalog, needs: needs,
                moreAbout: search.map { _ in { openItemPage(request.tooltip) } }
            )
        }
        .onChange(of: search?.bisRequest) { Task { await takeBisRequest() } }
        .onChange(of: router?.request) { _, request in Task { await apply(request) } }
        // Lane F15: a newer build waits behind the strip; nothing on screen moves.
        .onChange(of: scenePhase) { _, phase in
            if phase == .active { Task { await model.checkForNewData() } }
        }
        .task(id: avatarURL) {
            guard let avatarURL else { portrait = nil; return }
            portrait = await portraitSource.load(avatarURL)
        }
        .task(id: owner) {
            guard owner != nil else { return }
            await model.start()
            await apply(router?.request)
            await takeBisRequest()
            #if DEBUG
                // `-FSBisPicker open` opens the picker, so a capture script shows the real sheet.
                if UserDefaults.standard.string(forKey: "FSBisPicker") == "open" { model.isPickerPresented = true }
                await openFromLaunchArguments()
            #endif
        }
    }

    /// Signed in: Gear once the list is drawn, its notice when there is none; signed out: the classes.
    @ViewBuilder
    private var root: some View {
        if model.context.characterSelection == nil {
            GearClassesView(classes: model.context.classSlugs(model.catalog), name: model.context.className) { slug in
                listShown = true
                Task { await model.pickClass(slug) }
            }
        } else if case let .loaded(list, offlineSince) = model.phase, let inputs = model.gearInputs {
            GearListView(
                content: GearPresenter(context: model.context).content(list: list, inputs: inputs),
                build: model.build ?? "",
                portrait: portrait,
                actions: GearActions(openPicker: { model.isPickerPresented = true }, openItem: open)
            )
            .bandBar()
            .safeAreaInset(edge: .top, spacing: 0) { strips(offlineSince) }
        } else if let notice = GearPresenter(context: model.context).notice(phase: model.phase, title: model.specLabel(model.context.selection)) {
            // No list for the spec and band, or it did not load: the Gear tab's list line and one row.
            GearNoticeView(
                notice: notice,
                change: { model.isPickerPresented = true },
                retry: { Task { await model.refresh() } }
            )
            .bandBar()
        } else {
            list
        }
    }

    private var list: some View {
        @Bindable var model = model
        return BisView(
            phase: model.phase,
            context: model.context,
            catalog: model.catalog,
            searchText: $model.searchText,
            onOpenPicker: { model.isPickerPresented = true },
            onTryAgain: { Task { await model.refresh() } },
            onOpenItem: { sheet = ItemSheetRequest(tooltip: $0, gain: nil, guildID: guildID) },
            build: model.build,
            newData: model.newData,
            onUpdate: { Task { await model.applyNewData() } }
        )
    }

    private func strips(_ offlineSince: Date?) -> some View {
        VStack(spacing: 0) {
            if let offlineSince { OfflineStrip(text: Copy.offlineStrip(time: OfflineStamp.format(offlineSince, now: .now))) }
            NewDataStrip(model.newData) { Task { await model.applyNewData() } }
        }
    }

    private var guildID: Int64? { model.context.isYours ? model.character?.guild?.id : nil }

    /// A Gear row's item: the slot's pick or one of its runners-up, in the item sheet.
    private func open(_ ref: GearItemRef) {
        guard case let .loaded(list, _) = model.phase, let build = model.build,
              let row = list.rows.first(where: { $0.slot == ref.slot }), let pick = row.pick
        else { return }
        let runnerUp = ref.itemID == pick.item.id ? nil : pick.runnersUp.first { $0.item.id == ref.itemID }
        let tooltip = BisTooltip.request(row: row, pick: pick, list: list, context: model.context, build: build, runnerUp: runnerUp)
        sheet = ItemSheetRequest(tooltip: tooltip, gain: ref.gain, guildID: guildID)
    }

    private func openItemPage(_ request: ItemTooltipRequest) {
        sheet = nil
        itemPage = ItemPageRoute(itemID: request.itemID, fallback: request.fallback, classSlug: request.classSlug)
    }

    /// Lane F12: a list a search result asked for (nil: the character's own), once the model is up.
    private func takeBisRequest() async {
        guard model.build != nil, let request = search?.takeBisRequest() else { return }
        itemPage = nil
        if let selection = request.selection ?? model.context.characterSelection {
            if model.context.characterSelection == nil { listShown = true }
            await model.select(selection)
        }
    }

    /// Shows the asked list in the faction on screen; the picker closes.
    private func apply(_ request: BisRequest?) async {
        guard let request, let spec = model.catalog.spec(request.spec) else { return }
        model.isPickerPresented = false
        if model.context.characterSelection == nil { listShown = true }
        let faction = model.context.selection.faction
        await model.select(BisSelection(classSlug: spec.classSlug, spec: spec.spec, range: request.range, faction: faction))
    }

    #if DEBUG
        /// `-FSBisSpec priest-holy` opens that spec's 60 raid Horde list, and `-FSBisTooltip
        /// shoulder` that row's tooltip, so a capture script shows a healer or tank tooltip
        /// without taps (simctl cannot tap).
        private func openFromLaunchArguments() async {
            let defaults = UserDefaults.standard
            if let spec = defaults.string(forKey: "FSBisSpec") {
                model.isPickerPresented = false
                if model.context.characterSelection == nil { listShown = true }
                let classSlug = String(spec.prefix { $0 != "-" })
                let range = BisRange(band: BisRange.lastBand, preset: BisRange.raidPreset)
                await model.select(BisSelection(classSlug: classSlug, spec: spec, range: range, faction: "horde"))
            }
            guard let slot = defaults.string(forKey: "FSBisTooltip"), case let .loaded(list, _) = model.phase,
                  let build = model.build, let row = list.rows.first(where: { $0.slot == slot }), let pick = row.pick
            else { return }
            try? await Task.sleep(for: .seconds(1))
            let tooltip = BisTooltip.request(row: row, pick: pick, list: list, context: model.context, build: build)
            sheet = ItemSheetRequest(tooltip: tooltip, gain: nil, guildID: guildID)
        }
    #endif
}

extension BisScreen {
    /// The live tab: the CDN through the offline cache, the build from the app config.
    @MainActor
    static func live(_ environment: AppEnvironment, session: SessionStore?, router: BisRouter? = nil) -> BisScreen {
        let source = LiveBisDataSource(
            files: environment.staticData.files,
            appConfig: environment.appConfig,
            character: { session?.currentCharacter },
            bundledBuild: BundledBuild.read(),
            home: GearHome.live(client: environment.client, cache: FileHomeCache.standard() ?? MemoryHomeCache())
        )
        let model = BisModel(source: source, catalog: .bundled(), preferences: UserDefaultsBisPreferences())
        let store = environment.staticData
        model.latestBuild = { await store.latestBuild() }
        return BisScreen(
            model: model, catalog: ItemCatalog(store: environment.staticData),
            needs: session == nil ? nil : LiveItemNeedsSource(client: environment.client), session: session, router: router
        )
    }
}


/// F3's own filter while lane F12's search is off; the universal search button once it is on
/// (spec Wave 3 §3.2: F3 keeps no second search).
private struct BisSearchMode: ViewModifier {
    @Binding var text: String
    let isUniversal: Bool

    func body(content: Content) -> some View {
        if isUniversal {
            content.searchButton()
        } else {
            content
                .searchable(text: $text, prompt: Copy.bisSearch)
                .searchToolbarBehavior(.minimize)
        }
    }
}
