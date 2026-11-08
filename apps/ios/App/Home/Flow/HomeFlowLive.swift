import Foundation
import Models
import Networking

extension HomeFlowDependencies {
    /// The key date of the first raids (`role: raids`).
    static let firstRaidsRole = "raids"

    @MainActor
    static func live(_ environment: AppEnvironment) -> HomeFlowDependencies {
        let store = environment.staticData
        return HomeFlowDependencies(
            service: LiveHomeFlowService(client: environment.client),
            raids: RaidsDependencies.live(environment).service,
            ready: LiveReadySource(client: environment.client),
            applicants: LiveApplicantsSource(client: environment.client),
            bestRun: { home, character in await bestRun(store: store, home: home, character: character) },
            firstRaids: { firstRaids(BundledData.keyDates()) },
            now: AppClock.now
        )
    }

    /// Best run from the cached list and loot files of the answer's build (flows spec F5).
    static func bestRun(store: StaticDataStore, home: CharacterHome, character: HomeCharacter) async -> BestRun? {
        guard let spec = home.spec, let range = HomePresenter.range(home, character: character),
              let bis = try? await store.file(.getBis(build: home.dataBuild, spec: spec)),
              let loot = try? await store.file(.getLoot(build: home.dataBuild))
        else { return nil }
        let bands = bis.bands.filter { $0.band == range.band && $0.preset == (home.preset ?? range.preset) }
        let band = bands.first { $0.faction == character.faction } ?? bands.first
        guard let band else { return nil }
        return BestRun.best(slots: band.slots, gains: BestRun.gains(home.upgrades), loot: loot, level: character.level)
    }

    static func firstRaids(_ dates: [KeyDate]) -> Date? {
        guard let iso = dates.first(where: { $0.role == firstRaidsRole })?.iso else { return nil }
        return try? Date(iso + "T00:00:00Z", strategy: .iso8601)
    }
}
