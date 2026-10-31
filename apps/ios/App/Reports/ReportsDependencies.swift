import Foundation
import Models
import Networking

/// What F10 needs from the app: the report calls, the answer cache, the specs, the clock, the icon build.
struct ReportsDependencies {
    let service: any ReportsService
    let cache: any RaidCache
    let specs: SpecCatalog
    let clock: ReportsClock
    let now: @Sendable () -> Date
    let iconBuild: @Sendable () async -> String?

    @MainActor
    static func live(_ environment: AppEnvironment) -> ReportsDependencies {
        let staticTransport: any HTTPTransport = switch environment.configuration.mode {
        case .live: URLSessionTransport()
        case .fixture: FixtureTransport.bundled() ?? FixtureTransport(directory: Bundle.main.bundleURL)
        }
        let statics = StaticDataClient(transport: staticTransport)
        return ReportsDependencies(
            service: LiveReportsService(client: environment.client),
            cache: DiskRaidCache.reports(),
            specs: .bundled(),
            clock: ReportsClock(),
            now: { Date() },
            iconBuild: { try? await statics.activeBuild().build }
        )
    }
}

/// Who is looking (spec F10.1, F10.4): signed in or not, the account's characters for the crests,
/// and whether the coach card may draw.
struct ReportsAccount: Equatable, Sendable {
    let signedIn: Bool
    let characters: [AccountCharacter]
    let currentClass: String?
    /// `entitlements.log_coach`.
    let logCoach: Bool

    static let signedOut = ReportsAccount(signedIn: false, characters: [], currentClass: nil, logCoach: false)

    static func of(_ me: Me?, current: AccountCharacter?) -> ReportsAccount {
        guard let me else { return .signedOut }
        return ReportsAccount(
            signedIn: true, characters: me.characters, currentClass: current?.className,
            logCoach: me.entitlements.logCoach
        )
    }
}

/// The site pages F10 opens in the browser.
enum ReportsLinks {
    static let companion = HomeLinks.setup

    static func web(_ id: String) -> URL { RaidLinks.report(id) }
}

extension DiskRaidCache {
    /// F10's answers, beside the raids cache and cleared with it at sign-out.
    static func reports() -> DiskRaidCache {
        DiskRaidCache(directory: standard().directory.deletingLastPathComponent().appending(path: "reports", directoryHint: .isDirectory))
    }
}
