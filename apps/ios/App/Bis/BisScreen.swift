import DesignSystem
import Networking
import SwiftUI

/// The BiS tab (spec F3): the list, its picker and the item tooltip, over `BisModel`.
struct BisScreen: View {
    @State private var model: BisModel
    @State private var tooltip: ItemTooltipContent?
    @State private var isSearching = false
    private let iconBase: URL

    init(model: BisModel, iconBase: URL = StaticDataClient.siteURL) {
        _model = State(initialValue: model)
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
                onOpenItem: { tooltip = $0 }
            )
            .refreshable { await model.refresh() }
            .searchable(text: $model.searchText, isPresented: $isSearching, prompt: Copy.bisSearch)
            .toolbar {
                // The leading slot is the character selector (lane F1), the same on Home and Raids.
                ToolbarItem(placement: .topBarTrailing) {
                    Button { isSearching = true } label: { Image(systemName: "magnifyingglass") }
                        .accessibilityLabel(Copy.bisSearch)
                }
            }
            .toolbarBackground(Theme.Palette.background, for: .navigationBar)
            .sheet(isPresented: $model.isPickerPresented) {
                BisPickerSheet(context: model.context, catalog: model.catalog) { selection in
                    Task { await model.select(selection) }
                }
            }
            .sheet(item: $tooltip) { ItemTooltipSheet(content: $0) }
        }
        .environment(\.iconSource, model.build.map { IconSource(baseURL: iconBase, build: $0) })
        .task { await model.start() }
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
        let source = LiveBisDataSource(
            files: CachedStaticDataClient(
                client: StaticDataClient(transport: transport),
                cache: DiskStaticFileCache.applicationSupport()
            ),
            appConfig: environment.appConfig,
            client: environment.client,
            bundledBuild: BundledBuild.read()
        )
        return BisScreen(model: BisModel(
            source: source,
            catalog: .bundled(),
            preferences: UserDefaultsBisPreferences()
        ))
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
