import Auth
import DesignSystem
import Networking
import SwiftUI

/// The BiS tab (spec F3): the list, its picker and the item tooltip, over `BisModel`.
/// Signed in (lane F1), the list follows the session's current character under the selector.
struct BisScreen: View {
    @State private var model: BisModel
    @State private var tooltip: ItemTooltipRequest?
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
                build: model.build
            )
            .refreshable { await model.refresh() }
            // The system search, minimised to the bar's trailing button (spec F3.1). The leading
            // slot is the character selector (lane F1), the same on Home.
            .searchable(text: $model.searchText, prompt: Copy.bisSearch)
            .characterSelector(isOn: session != nil)
            .searchToolbarBehavior(.minimize)
            .toolbarBackground(FSColor.bg, for: .navigationBar)
            .sheet(isPresented: $model.isPickerPresented) {
                BisPickerSheet(context: model.context, catalog: model.catalog, build: model.build) { selection in
                    Task { await model.select(selection) }
                }
            }
            .sheet(item: $tooltip) { ItemTooltipSheet(request: $0, catalog: catalog) }
        }
        .onChange(of: router?.request) { _, request in Task { await apply(request) } }
        .task(id: owner) {
            guard owner != nil else { return }
            await model.start()
            await apply(router?.request)
            #if DEBUG
                // `-FSBisPicker open` opens the picker, so a capture script shows the real sheet.
                if UserDefaults.standard.string(forKey: "FSBisPicker") == "open" { model.isPickerPresented = true }
                await openFromLaunchArguments()
            #endif
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
        let transport: any HTTPTransport = switch environment.configuration.mode {
        case .live: URLSessionTransport()
        case .fixture: FixtureTransport.bundled() ?? FixtureTransport(directory: Bundle.main.bundleURL)
        }
        let staticClient = StaticDataClient(transport: transport)
        let source = LiveBisDataSource(
            files: CachedStaticDataClient(client: staticClient, cache: DiskStaticFileCache.applicationSupport()),
            appConfig: environment.appConfig,
            character: { session?.currentCharacter },
            bundledBuild: BundledBuild.read()
        )
        let model = BisModel(source: source, catalog: .bundled(), preferences: UserDefaultsBisPreferences())
        return BisScreen(model: model, catalog: ItemCatalog(client: staticClient), session: session, router: router)
    }
}
