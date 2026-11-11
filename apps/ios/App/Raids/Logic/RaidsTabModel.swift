import Foundation
import Models
import Networking
import Observation

/// The character's guild as the Raid tab names it (the guild head itself is retired: flows spec
/// appendix C; its name titles the guild's pages).
struct GuildHead: Equatable, Sendable {
    let id: Int64
    let name: String
    /// The guild's page on the site (`/guild/{region}/{ruleset}/{slug}`), when the account row has it.
    let webURL: URL?
    /// The roster's size, for the Roster row.
    let members: Int?

    init(character: CharacterGuild, guild: AccountGuild?) {
        id = character.id
        members = guild.map { $0.rosterCount ?? $0.verifiedCount }
        name = guild?.name ?? character.name
        webURL = guild.flatMap { guild in
            guild.slug.map { StaticSite.url.appending(path: "guild/\(guild.region)/\(guild.ruleset)/\($0)") }
        }
    }
}

/// What the Raids tab draws (spec F4.1 and F4.13 to F4.17).
enum RaidsTabState: Equatable {
    case loading
    case signedOut
    /// No guild (or no character): the guilds recruiting its spec.
    case noGuild(character: AccountCharacter?)
    case waiting(head: GuildHead, character: AccountCharacter)
    case loaded(RaidsTabContent)
    case failed
}

struct RaidsTabContent: Equatable {
    let head: GuildHead
    let character: AccountCharacter
    let planned: [RaidView]
    let nights: [RaidNightRow]
    /// The officer's waiting count; nil for anyone who cannot approve.
    let waiting: Int?
    /// Set while the cached answer shows because a refresh could not connect.
    let offlineSince: Date?
    /// Lane F15: set while the cached answer shows because Forever Sixty did not answer.
    var staleSince: Date?
    /// Every character key on the account (lane F16's `You were there`).
    var accountKeys: Set<String> = []

    func replacing(_ raid: RaidView) -> RaidsTabContent {
        RaidsTabContent(
            head: head, character: character,
            planned: planned.map { $0.id == raid.id ? raid : $0 },
            nights: nights, waiting: waiting, offlineSince: offlineSince, staleSince: staleSince, accountKeys: accountKeys
        )
    }
}

/// Loads `/v1/me` for the current character's guild, then the guild's raids; the cache draws first.
@MainActor
@Observable
final class RaidsTabModel {
    private(set) var state: RaidsTabState = .loading
    /// The character the tab is for; nil follows the account's main (lane F1's selector sets it).
    var currentKey: String?

    private let service: any RaidService
    private let cache: any RaidCache
    private let now: @Sendable () -> Date
    private let connectivity: (any Connectivity)?

    init(
        service: any RaidService, cache: any RaidCache, now: @escaping @Sendable () -> Date,
        initial: RaidsTabState = .loading, connectivity: (any Connectivity)? = nil
    ) {
        self.service = service
        self.cache = cache
        self.now = now
        self.state = initial
        self.connectivity = connectivity
    }

    func load() async {
        let me: Me
        do {
            me = try await service.me()
            cache.write(me, key: "me", at: now())
        } catch {
            if case .unauthorized = error { state = .signedOut; return }
            guard let cached = cache.read(Me.self, key: "me") else { state = .failed; return }
            me = cached.value
        }
        await load(for: me)
    }

    private func load(for me: Me) async {
        let character = me.characters.first { $0.key == currentKey } ?? me.mainCharacter
        guard let character else { state = .noGuild(character: nil); return }
        guard let membership = character.guild else { state = .noGuild(character: character); return }
        let head = GuildHead(character: membership, guild: me.guilds.first { $0.id == membership.id })
        guard membership.verified else { state = .waiting(head: head, character: character); return }
        let key = "guild-\(membership.id)-raids"
        if case .loading = state, let cached = cache.read(GuildRaids.self, key: key) {
            state = .loaded(content(head, character, cached.value, offlineSince: nil, me: me))
        }
        do {
            let raids = try await service.raids(guild: membership.id)
            cache.write(raids, key: key, at: now())
            state = .loaded(content(head, character, raids, offlineSince: nil, me: me))
        } catch {
            let cached = cache.read(GuildRaids.self, key: key)
            if let cached, let since = CacheFreshness.staleSince(
                error, connectivity: connectivity, savedAt: cached.savedAt, limit: .homeAndRaids, now: now()
            ) {
                var shown = content(head, character, cached.value, offlineSince: nil, me: me)
                shown.staleSince = since
                state = .loaded(shown)
            } else if SignupRefusal.isOffline(error), let cached {
                state = .loaded(content(head, character, cached.value, offlineSince: cached.savedAt, me: me))
            } else {
                state = .failed
            }
        }
    }

    private func content(_ head: GuildHead, _ character: AccountCharacter, _ raids: GuildRaids, offlineSince: Date?, me: Me) -> RaidsTabContent {
        RaidsTabContent(
            head: head, character: character,
            planned: raids.planned ?? [],
            nights: Array(raids.rows.filter { $0.raidId != nil }.prefix(Self.nightRows)),
            waiting: raids.waitingForApproval.flatMap { $0 > 0 ? $0 : nil },
            offlineSince: offlineSince,
            accountKeys: Set(me.characters.map(\.key))
        )
    }

    /// The raid nights row count (spec F4.1 item 4).
    static let nightRows = 3

    /// A sign-up on the card or the raid page replaced this raid.
    func replace(_ raid: RaidView) {
        guard case let .loaded(content) = state else { return }
        state = .loaded(content.replacing(raid))
    }
}
