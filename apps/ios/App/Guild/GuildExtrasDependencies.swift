import Foundation
import Models
import Networking

/// The guild extras' dependencies (spec F16), built from the app's environment.
struct GuildExtrasDependencies {
    let service: any GuildExtrasService
    let cache: any RaidCache
    let linkStore: any CalendarLinkStore
    let now: @Sendable () -> Date
    let clock: RaidClock
    /// Lane F15: the connection, for the stale strip; nil (tests) keeps v1's offline rule alone.
    var connectivity: (any Connectivity)?

    static func live(_ environment: AppEnvironment) -> GuildExtrasDependencies {
        GuildExtrasDependencies(
            service: service(environment),
            cache: DiskRaidCache.standard(),
            linkStore: KeychainCalendarLinkStore(),
            now: AppClock.now,
            clock: RaidClock(),
            connectivity: PathConnectivity.shared
        )
    }

    private static func service(_ environment: AppEnvironment) -> any GuildExtrasService {
        if environment.configuration.mode == .fixture,
           let folder = Bundle.main.url(forResource: FixtureTransport.bundledFolder, withExtension: nil) {
            return FixtureGuildExtrasService(directory: folder, baseURL: environment.configuration.baseURL)
        }
        return LiveGuildExtrasService(client: environment.client)
    }

    @MainActor
    func rosterStore(guild: Int64) -> RosterStore {
        RosterStore(guild: guild, service: service, cache: cache, now: now).connected(connectivity)
    }

    @MainActor
    func historyStore(guild: Int64) -> HistoryStore {
        HistoryStore(guild: guild, service: service, cache: cache, now: now).connected(connectivity)
    }

    @MainActor
    func linkModel(account: String) -> CalendarLinkModel {
        CalendarLinkModel(service: service, store: linkStore, account: account)
    }
}

/// The planned raids for the calendar when the switch turns on before the Raids tab loaded:
/// `/v1/me`, the current character's verified guild, its raids.
struct CalendarInputLoader {
    let service: any RaidService
    let currentKey: @MainActor () -> String?

    @MainActor
    func load() async -> CalendarInput? {
        guard let me = try? await service.me() else { return nil }
        let character = me.characters.first { $0.key == currentKey() } ?? me.mainCharacter
        guard let membership = character?.guild, membership.verified,
              let raids = try? await service.raids(guild: membership.id)
        else { return nil }
        let name = me.guilds.first { $0.id == membership.id }?.name ?? membership.name
        return CalendarInput(guildID: membership.id, guildName: name, raids: raids.planned ?? [])
    }
}

extension RaidsTabContent {
    /// The calendar's input from the tab's loaded raids.
    var calendarInput: CalendarInput {
        CalendarInput(guildID: head.id, guildName: head.name, raids: planned)
    }
}
