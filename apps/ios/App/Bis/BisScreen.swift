import Auth
import DesignSystem
import Networking
import SwiftUI

/// The BiS tab (spec F3): the list, its picker and the item tooltip, over `BisModel`.
/// Signed in (lane F1), the list follows the session's current character under the selector.
struct BisScreen: View {
    @State private var model: BisModel
    @State private var tooltip: ItemTooltipRequest?
    /// Lane F12: the item page, pushed from the tooltip's `More about this item`.
    @State private var itemPage: ItemPageRoute?
    @Environment(\.searchNavigator) private var search
    @Environment(\.scenePhase) private var scenePhase
    private let catalog: ItemCatalog
    /// Lane F1's account, nil while its feature switch is off.
    private let session: SessionStore?
    /// Lists other tabs and links ask for (a guide's Gear row, `/app/bis`).
    private let router: BisRouter?

    init(model: BisModel, catalog: ItemCatalog, session: SessionStore? = nil, router: BisRouter? = nil) {
        _model = State(initialValue: model)
        self.catalog = catalog
        self.session = session
        self.router = router
    }

    /// The model restarts whenever the session's current character changes; nil while a signed-in
    /// session's characters are still on their way, so the signed-out first visit never fires early.
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
            BisView(
                phase: model.phase,
                context: model.context,
                catalog: model.catalog,
                searchText: $model.searchText,
                onOpenPicker: { model.isPickerPresented = true },
                onTryAgain: { Task { await model.refresh() } },
                onOpenItem: { tooltip = $0 },
                build: model.build,
                newData: model.newData,
                onUpdate: { Task { await model.applyNewData() } }
            )
            .refreshable { await model.refresh() }
            // The system search, minimised to the bar's trailing button (spec F3.1), until lane
            // F12's universal search replaces it. The leading slot is the character selector
            // (lane F1), the same on Home.
            .modifier(BisSearchMode(text: $model.searchText, isUniversal: search != nil))
            .characterSelector(isOn: session != nil)
            .toolbarBackground(FSColor.bg, for: .navigationBar)
            .sheet(isPresented: $model.isPickerPresented) {
                BisPickerSheet(context: model.context, catalog: model.catalog, build: model.build) { selection in
                    Task { await model.select(selection) }
                }
            }
            .sheet(item: $tooltip) { request in
                ItemTooltipSheet(request: request, catalog: catalog, moreAbout: search.map { _ in { openItemPage(request) } })
            }
            .navigationDestination(item: $itemPage) { route in
                if let search {
                    ItemPageScreen(route: route, dependencies: search.dependencies) { search.openList($0) }
                }
            }
        }
        .onChange(of: search?.bisRequest) { Task { await takeBisRequest() } }
        .onChange(of: router?.request) { _, request in Task { await apply(request) } }
        // Lane F15: a newer build waits behind the strip; nothing on screen moves.
        .onChange(of: scenePhase) { _, phase in
            if phase == .active { Task { await model.checkForNewData() } }
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

    private func openItemPage(_ request: ItemTooltipRequest) {
        tooltip = nil
        itemPage = ItemPageRoute(itemID: request.itemID, fallback: request.fallback, classSlug: request.classSlug)
    }

    /// Lane F12: a list a search result asked for (nil: the character's own), once the model is up.
    private func takeBisRequest() async {
        guard model.build != nil, let request = search?.takeBisRequest() else { return }
        itemPage = nil
        if let selection = request.selection ?? model.context.characterSelection {
            await model.select(selection)
        }
    }

    /// Shows the asked list in the faction on screen; the picker closes.
    private func apply(_ request: BisRequest?) async {
        guard let request, let spec = model.catalog.spec(request.spec) else { return }
        model.isPickerPresented = false
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
                let classSlug = String(spec.prefix { $0 != "-" })
                let range = BisRange(band: BisRange.lastBand, preset: BisRange.raidPreset)
                await model.select(BisSelection(classSlug: classSlug, spec: spec, range: range, faction: "horde"))
            }
            guard let slot = defaults.string(forKey: "FSBisTooltip"), case let .loaded(list, _) = model.phase,
                  let build = model.build, let row = list.rows.first(where: { $0.slot == slot }), let pick = row.pick
            else { return }
            try? await Task.sleep(for: .seconds(1))
            tooltip = BisTooltip.request(row: row, pick: pick, list: list, context: model.context, build: build)
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
            bundledBuild: BundledBuild.read()
        )
        let model = BisModel(source: source, catalog: .bundled(), preferences: UserDefaultsBisPreferences())
        let store = environment.staticData
        model.latestBuild = { await store.latestBuild() }
        return BisScreen(model: model, catalog: ItemCatalog(store: environment.staticData), session: session, router: router)
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
