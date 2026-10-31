import Foundation
import Models
import Networking

/// The build's static files the tooltip reads (`items/{class}.json`, `sets.json`,
/// `bis/{spec}.json`), through the shared static data store: fetched once per build, then read
/// from memory and disk (spec §5.0 Caching). Never sends the device token.
actor ItemCatalog {
    private let store: StaticDataStore

    init(store: StaticDataStore) {
        self.store = store
    }

    /// A catalog over its own store; `directory` nil keeps the files in memory (tests and previews).
    init(client: StaticDataClient, directory: URL? = StaticDataStore.defaultDirectory) {
        self.init(store: StaticDataStore(client: client, directory: directory))
    }

    /// The files a tooltip needs. Items are required; sets and the BiS file only add lines,
    /// so a missing one leaves its lines out rather than failing the sheet.
    func sources(for request: ItemTooltipRequest) async throws(APIError) -> ItemTooltipSources {
        let items = try await store.file(.getItems(build: request.build, className: request.classSlug))
        let sets = try? await store.file(.getSets(build: request.build))
        let bis: BisFile? = if let spec = request.spec {
            try? await store.file(.getBis(build: request.build, spec: spec))
        } else {
            nil
        }
        return ItemTooltipSources(
            items: Dictionary(items.items.map { ($0.id, $0) }, uniquingKeysWith: { first, _ in first }),
            sets: sets ?? [],
            bis: bis
        )
    }
}
