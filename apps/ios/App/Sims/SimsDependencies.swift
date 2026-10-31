import Foundation
import Networking
import SwiftUI

/// What F5 needs from the app: the sim calls, the answer cache, the specs, the item catalog, the
/// icon build, and whether a See plans link may show in this storefront.
struct SimsDependencies {
    let service: any SimsService
    let cache: any RaidCache
    let specs: SpecCatalog
    let catalog: ItemCatalog
    let now: @Sendable () -> Date
    let iconBuild: @Sendable () async -> String?
    /// F8's rule: See plans only where the storefront may link to web checkout (fail closed).
    let mayLinkToPlans: @MainActor () async -> Bool

    @MainActor
    static func live(_ environment: AppEnvironment, catalog: ItemCatalog) -> SimsDependencies {
        let staticTransport: any HTTPTransport = switch environment.configuration.mode {
        case .live: URLSessionTransport()
        case .fixture: FixtureTransport.bundled() ?? FixtureTransport(directory: Bundle.main.bundleURL)
        }
        let statics = StaticDataClient(transport: staticTransport)
        return SimsDependencies(
            service: LiveSimsService(client: environment.client),
            cache: DiskRaidCache.sims(),
            specs: .bundled(),
            catalog: catalog,
            now: { Date() },
            iconBuild: { try? await statics.activeBuild().build },
            mayLinkToPlans: {
                await environment.plansLink.resolve()
                return environment.plansLink.mayLink
            }
        )
    }
}

extension DiskRaidCache {
    /// F5's answers, beside the raids cache and cleared with it at sign-out.
    static func sims() -> DiskRaidCache {
        DiskRaidCache(directory: standard().directory.deletingLastPathComponent().appending(path: "sims", directoryHint: .isDirectory))
    }
}
