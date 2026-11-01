import Foundation
import Models
import Networking
import Observation

/// The guild head (spec F4.1 item 1).
struct GuildHead: Equatable, Sendable {
    let id: Int64
    let name: String
    let crestURL: URL?
    let faction: String?
    /// `Horde · US Normal · 38 on the roster`, or nil when the account's guild row is missing.
    let line: String?
    /// The guild's page on the site (`/guild/{region}/{ruleset}/{slug}`), when the account row has it.
    let webURL: URL?

    init(character: CharacterGuild, guild: AccountGuild?) {
        id = character.id
        name = guild?.name ?? character.name
        crestURL = guild?.crestUrl.flatMap(URL.init(string:))
        faction = guild?.faction
        line = guild.map { guild in
            Copy.raidsGuildLine(
                faction: guild.faction.map(\.capitalized) ?? "",
                region: guild.region.uppercased(),
                ruleset: guild.rulesetLabel ?? guild.ruleset.capitalized,
                n: guild.rosterCount ?? guild.verifiedCount
            )
        }
        webURL = guild.flatMap { guild in
            guild.slug.map { StaticSite.url.appending(path: "guild/\(guild.region)/\(guild.ruleset)/\($0)") }
        }
    }
}

/// What the Raids tab draws (spec F4.1 and F4.13 to F4.17).
enum RaidsTabState: Equatable {
    case loading
    case signedOut
    case noGuild(name: String)
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
    /// Every character key on the account (lane F16's `You were there`).
    var accountKeys: Set<String> = []

    func replacing(_ raid: RaidView) -> RaidsTabContent {
        RaidsTabContent(
            head: head, character: character,
            planned: planned.map { $0.id == raid.id ? raid : $0 },
            nights: nights, waiting: waiting, offlineSince: offlineSince, accountKeys: accountKeys
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

    init(service: any RaidService, cache: any RaidCache, now: @escaping @Sendable () -> Date, initial: RaidsTabState = .loading) {
        self.service = service
        self.cache = cache
        self.now = now
        self.state = initial
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
        guard let character else { state = .noGuild(name: ""); return }
        guard let membership = character.guild else { state = .noGuild(name: character.name); return }
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
            if SignupRefusal.isOffline(error), let cached = cache.read(GuildRaids.self, key: key) {
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
