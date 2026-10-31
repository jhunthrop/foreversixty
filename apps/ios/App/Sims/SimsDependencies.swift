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
    /// Lane F14: the server runs in flight, for their Live Activities; nil draws none.
    var liveSims: (@MainActor ([SimLiveRun]) async -> Void)?

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
            },
            liveSims: liveSims(environment)
        )
    }
}

extension SimsDependencies {
    /// Lane F14: the runs' Live Activities, behind the widgets switch and `server_sims`.
    @MainActor
    private static func liveSims(_ environment: AppEnvironment) -> (@MainActor ([SimLiveRun]) async -> Void)? {
        guard environment.isAvailable(.widgets) else { return nil }
        let controller = environment.liveActivities
        let session = environment.session
        return { runs in
            await controller.syncSims(runs, entitled: session.phase.me?.entitlements.serverSims == true)
        }
    }
}

extension DiskRaidCache {
    /// F5's answers, beside the raids cache and cleared with it at sign-out.
    static func sims() -> DiskRaidCache {
        DiskRaidCache(directory: standard().directory.deletingLastPathComponent().appending(path: "sims", directoryHint: .isDirectory))
    }
}
