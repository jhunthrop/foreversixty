import Foundation
import Models
import Networking
import Observation

/// The calls only the time-aware Home makes, behind a protocol for tests.
protocol HomeFlowService: Sendable {
    /// `GET /v1/home?character=`: the blocks in the moment's order.
    func sections(character: String) async throws(APIError) -> HomeBlocks
    /// `GET /v1/me/activity?group=`: New for you's rows and the unread count.
    func activity(group: String?) async throws(APIError) -> ActivityPage
    /// `GET /v1/guilds/{id}/raid-badge`: an officer's open items.
    func badge(guild: Int64) async throws(APIError) -> RaidBadge
}

struct LiveHomeFlowService: HomeFlowService {
    let client: APIClient

    func sections(character: String) async throws(APIError) -> HomeBlocks {
        try await client.send(.getHome(character: character))
    }

    func activity(group: String?) async throws(APIError) -> ActivityPage {
        try await client.send(.getActivity(group: group))
    }

    func badge(guild: Int64) async throws(APIError) -> RaidBadge {
        try await client.send(.getGuildRaidBadge(id: guild))
    }
}

/// The Ready check (flows spec F2): `GET /v1/characters/{region}/{ruleset}/{name}/ready?raid=`;
/// nil without an export or on any failure (the row is left out).
protocol ReadySource: Sendable {
    func ready(character: String, raid: Int64) async -> ReadyCheck?
}

/// The applicants an officer sees (F3): the pending rows of `GET /v1/guilds/{id}/applications`.
protocol ApplicantsSource: Sendable {
    func applicants(guild: Int64) async -> [ApplicantLine]
}

struct LiveReadySource: ReadySource {
    let client: APIClient

    func ready(character: String, raid: Int64) async -> ReadyCheck? {
        guard let key = try? CharacterKey(character) else { return nil }
        return try? await client.send(.getCharacterReady(region: key.region, ruleset: key.ruleset, name: key.name, raid: raid))
    }
}

struct LiveApplicantsSource: ApplicantsSource {
    static let pending = "pending"
    let client: APIClient

    func applicants(guild: Int64) async -> [ApplicantLine] {
        guard let page: GuildOfficerApplications = try? await client.send(.listGuildApplications(id: guild)) else { return [] }
        return ApplicantLine.pending(page)
    }
}

/// What the time-aware Home reads, built once from the app's environment.
struct HomeFlowDependencies {
    let service: any HomeFlowService
    let raids: any RaidService
    let ready: any ReadySource
    let applicants: any ApplicantsSource
    /// Best run from the cached list and loot files; nil without them.
    let bestRun: @Sendable (CharacterHome, HomeCharacter) async -> BestRun?
    /// The first raids' date, for Find a guild's line.
    let firstRaids: @Sendable () async -> Date?
    let now: @Sendable () -> Date
}

/// Loads Home's blocks for one character: `/v1/home` decides the moment and the order, and each
/// block's data comes from its own call; a block whose call fails is left out, never the Home.
@MainActor
@Observable
final class HomeFlowModel {
    enum Phase: Equatable {
        case loading
        case failed
        case loaded
    }

    /// A raid counts as live from its pull to four hours after (flows spec 4.1).
    nonisolated static let liveSpan: TimeInterval = 4 * 3600

    private(set) var phase: Phase = .loading
    private(set) var blocks: [HomeBlock] = []
    private var loaded: HomeFlowInputs
    /// The raid card's answer (F1), shared with the Raid tab's head.
    let answers: RaidAnswerModel

    private let dependencies: HomeFlowDependencies
    private let home: HomeStore

    /// `haptics`: the answers' (R14), as the raid page's; `savedHold`: how long `Saved.` stays (nil keeps it).
    init(
        character: HomeCharacter, home: HomeStore, dependencies: HomeFlowDependencies,
        haptics: any Haptics = SystemHaptics(), savedHold: Duration? = .seconds(4)
    ) {
        loaded = HomeFlowInputs(character: character)
        self.home = home
        self.dependencies = dependencies
        answers = RaidAnswerModel(service: dependencies.raids, haptics: haptics, hold: savedHold)
    }

    /// Everything Home draws from: the loaded data and the raid answer's state.
    var inputs: HomeFlowInputs {
        var inputs = loaded
        inputs.answering = answers.state
        return inputs
    }

    /// The failed write's sentence, while it shows.
    var answerError: String? { answers.error }

    /// Home's answers disable while its answers show from the cache for lack of a connection.
    var offline: Bool {
        get { answers.offline }
        set { answers.offline = newValue }
    }

    var character: HomeCharacter { inputs.character }

    func load() async {
        await home.load(inputs.character)
        await refresh()
    }

    func refresh() async {
        do {
            blocks = HomeBlock.of(try await dependencies.service.sections(character: inputs.character.key).sections)
        } catch {
            if blocks.isEmpty {
                phase = .failed
                return
            }
        }
        loaded.home = home.answer
        await loadBlockData()
        phase = .loaded
    }

    private func loadBlockData() async {
        let character = inputs.character
        async let activity = activity()
        async let raids = raids()
        async let firstRaids = blocks.contains(.guildCTA) ? dependencies.firstRaids() : nil
        async let bestRun = bestRun()
        loaded.activity = await activity
        loaded.firstRaids = await firstRaids
        loaded.bestRun = await bestRun
        let fetched = await raids
        loaded.raid = fetched.raid
        loaded.night = fetched.night
        if let guild = character.guildID, let raid = loaded.raid {
            async let badge = try? dependencies.service.badge(guild: guild)
            async let applicants = dependencies.applicants.applicants(guild: guild)
            async let ready = dependencies.ready.ready(character: character.key, raid: raid.id)
            loaded.badge = await badge
            loaded.applicants = await applicants
            loaded.ready = await ready
        }
    }

    private func activity() async -> ActivityPage? {
        let query = blocks.lazy.compactMap { block -> HomeActivityQuery? in
            if case let .activity(query) = block { return query }
            return nil
        }.first
        guard let query else { return nil }
        return try? await dependencies.service.activity(group: query.group)
    }

    private func bestRun() async -> BestRun? {
        guard HomeBlock.showsBestRun(blocks), let answer = home.answer else { return nil }
        return await dependencies.bestRun(answer, inputs.character)
    }

    /// The raid a raid block needs (the next one, or the one in progress) and last night's summary.
    private func raids() async -> (raid: RaidView?, night: RaidNightSummary?) {
        guard let guild = inputs.character.guildID else { return (nil, nil) }
        let wantsRaid = HomeBlock.needsRaid(blocks)
        let wantsNight = blocks.contains(.raidSummary)
        guard wantsRaid || wantsNight, let list = try? await dependencies.raids.raids(guild: guild) else { return (nil, nil) }
        let raid = wantsRaid ? Self.current(list.planned ?? [], now: dependencies.now()) : nil
        var night: RaidNightSummary?
        if wantsNight, let last = list.rows.compactMap(\.raidId).first {
            night = try? await dependencies.raids.night(raid: last)
        }
        return (raid, night)
    }

    /// The raid in progress, else the next one.
    nonisolated static func current(_ planned: [RaidView], now: Date) -> RaidView? {
        let live = planned.first { $0.startsAt <= now && now.timeIntervalSince($0.startsAt) < liveSpan }
        return live ?? planned.filter { $0.startsAt > now }.min { $0.startsAt < $1.startsAt }
    }

    // MARK: Answering (F1)

    /// In, Late, Bench or Out for the selected character; Late opens its minutes.
    func answer(_ state: SignupState) async {
        await send { [answers] guild, raid, key, note in await answers.answer(state, guild: guild, raid: raid, key: key, note: note) }
    }

    /// Late's minutes (15, 30, 60), kept open with the chosen one selected.
    func late(_ minutes: Int) async {
        await send { [answers] guild, raid, key, note in await answers.late(minutes, guild: guild, raid: raid, key: key, note: note) }
    }

    /// Try again after a failed write.
    func retry() async {
        await send { [answers] guild, raid, key, note in await answers.retry(guild: guild, raid: raid, key: key, note: note) }
    }

    private func send(_ write: (Int64, Int64, String, String) async -> RaidView?) async {
        guard let raid = loaded.raid, let guild = loaded.character.guildID else { return }
        let key = raid.viewer.characters.first { $0.characterKey == loaded.character.key }?.characterKey ?? loaded.character.key
        let note = raid.signups.first { $0.characterKey == key }?.note ?? ""
        if let view = await write(guild, raid.id, key, note) { loaded.raid = view }
    }
}
