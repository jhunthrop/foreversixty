import Foundation
import Models
import Networking
import Observation
import Routing

/// What F11 needs from the app: the calls, the answer cache, the specs, the items catalog for the
/// Paperdoll and its tooltips, the clock, the icon build, and F10's report screen when on.
struct RankingsDependencies {
    let service: any RankingsService
    let cache: any RaidCache
    let specs: SpecCatalog
    let catalog: ItemCatalog
    let clock: ReportsClock
    let now: @Sendable () -> Date
    let iconBuild: @Sendable () async -> String?
    /// Lane F10's report, opened from a row's menu or a recent fight; nil while its switch is off.
    let reports: ReportsDependencies?

    @MainActor
    static func live(_ environment: AppEnvironment, catalog: ItemCatalog, reports: ReportsDependencies?) -> RankingsDependencies {
        let staticTransport: any HTTPTransport = switch environment.configuration.mode {
        case .live: URLSessionTransport()
        case .fixture: FixtureTransport.bundled() ?? FixtureTransport(directory: Bundle.main.bundleURL)
        }
        let statics = StaticDataClient(transport: staticTransport)
        return RankingsDependencies(
            service: LiveRankingsService(client: environment.client),
            cache: DiskRaidCache.rankings(),
            specs: .bundled(),
            catalog: catalog,
            clock: ReportsClock(),
            now: { Date() },
            iconBuild: { try? await statics.activeBuild().build },
            reports: reports
        )
    }

    /// The class's items file, for the sockets' icons, qualities and item levels.
    func items(build: String?) -> @Sendable (String) async -> [Int: Item] {
        let catalog = catalog
        return { className in
            guard let build else { return [:] }
            let request = ItemTooltipRequest(
                itemID: 0, fallback: HomeItemLook(id: 0, name: nil, quality: nil, icon: nil, itemLevel: nil),
                build: build, classSlug: className, spec: nil, specName: nil, band: nil, bandLabel: nil,
                preset: nil, faction: nil, sourceLine: nil, wornID: nil
            )
            return (try? await catalog.sources(for: request).items) ?? [:]
        }
    }
}

/// A character page to push.
struct PlayerRoute: Hashable, Identifiable {
    let ref: CharacterRef
    var id: String { ref.key }
}

/// Where a rankings or player link lands on Home's stack (spec §3.4).
@MainActor
@Observable
final class RankingsRouter {
    /// Rankings is pushed on Home.
    var showsBoard = false
    /// The encounter a `/app/rankings/{encounter}` link opens on.
    var encounter: String?
    /// A character page pushed on Home (a `/app/player/...` link).
    var player: PlayerRoute?

    static func handles(_ link: DeepLink) -> Bool {
        switch link {
        case .rankings, .player: true
        default: false
        }
    }

    func open(_ link: DeepLink) {
        switch link {
        case let .rankings(encounter):
            player = nil
            self.encounter = encounter
            showsBoard = true
        case let .player(region, ruleset, name):
            showsBoard = false
            player = PlayerRoute(ref: CharacterRef(region: region, ruleset: ruleset, slug: name, name: name.capitalized))
        default:
            break
        }
    }
}

extension DiskRaidCache {
    /// F11's answers, beside the raids cache and cleared with it at sign-out.
    static func rankings() -> DiskRaidCache {
        DiskRaidCache(directory: standard().directory.deletingLastPathComponent().appending(path: "rankings", directoryHint: .isDirectory))
    }
}
