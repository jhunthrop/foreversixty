import Foundation
import Models
import Networking

// MERGE NOTE (lanes F2 and F3): F2's ItemCatalog.swift replaces this file at integration. Same
// name, same `init(client:)` and `sources(for:)`, so BiS's call site compiles against either.

/// The build's static files a tooltip reads (`items/{class}.json`, `bis/{spec}.json`), through
/// the offline cache, so a tooltip opens offline from what the list already saved.
actor ItemCatalog {
    private let files: CachedStaticDataClient

    init(client: StaticDataClient, cache: any StaticFileCache = DiskStaticFileCache.applicationSupport()) {
        files = CachedStaticDataClient(client: client, cache: cache)
    }

    /// Items are required; the BiS file only adds lines, so a missing one leaves them out.
    func sources(for request: ItemTooltipRequest) async throws(APIError) -> ItemTooltipSources {
        let items = try await files.load(.getItems(build: request.build, className: request.classSlug)).value
        let bis: BisFile? = if let spec = request.spec {
            try? await files.load(.getBis(build: request.build, spec: spec)).value
        } else {
            nil
        }
        return ItemTooltipSources(
            items: Dictionary(items.items.map { ($0.id, $0) }, uniquingKeysWith: { first, _ in first }),
            sets: [],
            bis: bis
        )
    }
}
