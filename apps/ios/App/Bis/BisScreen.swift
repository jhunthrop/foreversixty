import DesignSystem
import Networking
import SwiftUI

/// The BiS tab (spec F3): the list, its picker and the item tooltip, over `BisModel`.
struct BisScreen: View {
    @State private var model: BisModel
    @State private var tooltip: ItemTooltipRequest?
    private let iconBase: URL
    private let catalog: ItemCatalog

    init(model: BisModel, catalog: ItemCatalog, iconBase: URL = StaticDataClient.siteURL) {
        _model = State(initialValue: model)
        self.catalog = catalog
        self.iconBase = iconBase
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
            // slot is the character selector (lane F1), the same on Home and Raids.
            .searchable(text: $model.searchText, prompt: Copy.bisSearch)
            .searchToolbarBehavior(.minimize)
            .toolbarBackground(Theme.Palette.background, for: .navigationBar)
            .sheet(isPresented: $model.isPickerPresented) {
                BisPickerSheet(context: model.context, catalog: model.catalog) { selection in
                    Task { await model.select(selection) }
                }
            }
            .sheet(item: $tooltip) { ItemTooltipSheet(request: $0, catalog: catalog) }
        }
        .environment(\.iconSource, model.build.map { IconSource(baseURL: iconBase, build: $0) })
        .task {
            await model.start()
            #if DEBUG
                // `-FSBisPicker open` opens the picker, so a capture script shows the real sheet.
                if UserDefaults.standard.string(forKey: "FSBisPicker") == "open" { model.isPickerPresented = true }
            #endif
        }
    }
}

extension BisScreen {
    /// The live tab: the CDN through the offline cache, the build from the app config.
    @MainActor
    static func live(_ environment: AppEnvironment) -> BisScreen {
        let transport: any HTTPTransport = switch environment.configuration.mode {
        case .live: URLSessionTransport()
        case .fixture: FixtureTransport.bundled() ?? FixtureTransport(directory: Bundle.main.bundleURL)
        }
        let staticClient = StaticDataClient(transport: transport)
        let source = LiveBisDataSource(
            files: CachedStaticDataClient(client: staticClient, cache: DiskStaticFileCache.applicationSupport()),
            appConfig: environment.appConfig,
            client: environment.client,
            bundledBuild: BundledBuild.read()
        )
        let model = BisModel(source: source, catalog: .bundled(), preferences: UserDefaultsBisPreferences())
        return BisScreen(model: model, catalog: ItemCatalog(client: staticClient))
    }

    /// The `bis` feature switch; on in debug and fixture builds so the tab can be built and
    /// reviewed before the API turns it on.
    static func isEnabled(_ environment: AppEnvironment) -> Bool {
        #if DEBUG
            return true
        #else
            return environment.configuration.mode == .fixture || environment.appConfig.isEnabled(.bis)
        #endif
    }
}
