@testable import ForeverSixty
import Foundation
import Models
import Networking

/// The flows' player (the contract fixtures' `me`): Gorvash, Fury Warrior of OLYMPUS XXVII
/// (guild 2), signed In to Onyxia's Lair (raid 31), the next raid at `RaidTime.now`.
enum FlowPlayer {
    static let guild: Int64 = 2
    static let raid: Int64 = 31
    static let key = "us/pvp/gorvash"
    /// The ask-spec fixture's unsigned character with three specs.
    static let askKey = "us/pvp/lightningbolltz"
}

/// Holds a write in flight until the test opens it.
final class FlowGate: @unchecked Sendable {
    private let mutex = NSLock()
    private var opened = false
    private var waiting: CheckedContinuation<Void, Never>?

    func wait() async {
        await withCheckedContinuation { continuation in
            let resumeNow = mutex.withLock { () -> Bool in
                if opened { return true }
                waiting = continuation
                return false
            }
            if resumeNow { continuation.resume() }
        }
    }

    func open() {
        let continuation = mutex.withLock { () -> CheckedContinuation<Void, Never>? in
            opened = true
            defer { waiting = nil }
            return waiting
        }
        continuation?.resume()
    }
}

/// The fixture-mode raid service the flows run on, as Android's FlowSupport: a sign-up echoes the
/// requested state (and Late's minutes) into the viewer's own row, so every reload draws the answer
/// the server would. A gate holds a write in flight; `failWith` fails the next writes; `offline`
/// fails every call as a lost connection.
final class EchoRaidService: RaidService, @unchecked Sendable {
    private let mutex = NSLock()
    private var raids: [Int64: RaidView]
    private let list: GuildRaids
    private(set) var requests: [(key: String, request: SignupRequest)] = []
    private(set) var inFlight = 0
    var gate: FlowGate?
    var failWith: APIError?
    var offline = false
    /// The raid page's raid, when a flow needs another fixture (the ask-spec raid).
    var raidOverride: RaidView?

    init() throws {
        list = try RaidFixture.load(GuildRaids.self, "guild.raids")
        raids = Dictionary((list.planned ?? []).map { ($0.id, $0) }) { first, _ in first }
    }

    var sent: [SignupRequest] { mutex.withLock { requests.map(\.request) } }

    private func online() throws(APIError) {
        if offline { throw .transport("offline") }
    }

    func me() async throws(APIError) -> Me {
        try online()
        return try RaidFixture.api(Me.self, "me")
    }

    func raids(guild: Int64) async throws(APIError) -> GuildRaids {
        try online()
        let planned = mutex.withLock { (list.planned ?? []).map { raids[$0.id] ?? $0 } }
        do {
            guard var json = try JSONSerialization.jsonObject(with: WireFormat.encoder().encode(list)) as? [String: Any] else {
                throw APIError.missingFixture("guild.raids does not encode as an object")
            }
            json["planned"] = try JSONSerialization.jsonObject(with: WireFormat.encoder().encode(planned))
            return try WireFormat.decoder().decode(GuildRaids.self, from: JSONSerialization.data(withJSONObject: json))
        } catch let error as APIError {
            throw error
        } catch {
            throw .missingFixture("raids: \(error)")
        }
    }

    func raid(guild: Int64, raid: Int64) async throws(APIError) -> RaidView {
        try online()
        if let raidOverride { return raidOverride }
        guard let view = mutex.withLock({ raids[raid] }) else { throw .server(status: 404, body: ErrorBody(code: "not_found", message: ""), requestID: nil) }
        return view
    }

    func signup(guild: Int64, raid: Int64, characterKey: String, _ request: SignupRequest) async throws(APIError) -> SignupAnswer {
        mutex.withLock {
            requests.append((characterKey, request))
            inFlight += 1
        }
        if let gate { await gate.wait() }
        mutex.withLock { inFlight -= 1 }
        try online()
        if let failWith { throw failWith }
        let base: RaidView
        if let raidOverride {
            base = raidOverride
        } else if let stored = mutex.withLock({ raids[raid] }) {
            base = stored
        } else {
            throw .missingFixture("no raid \(raid)")
        }
        let echoed = try Self.echo(base, key: characterKey, request)
        mutex.withLock { raids[raid] = echoed }
        if raidOverride != nil { raidOverride = echoed }
        return SignupAnswer(raid: echoed, moved: nil, confirmation: nil)
    }

    func lock(guild: Int64, raid: Int64) async throws(APIError) -> RaidView { try await self.raid(guild: guild, raid: raid) }
    func unlock(guild: Int64, raid: Int64) async throws(APIError) -> RaidView { try await self.raid(guild: guild, raid: raid) }
    func night(raid: Int64) async throws(APIError) -> RaidNightSummary {
        try online()
        return try RaidFixture.api(RaidNightSummary.self, "me.raid-night")
    }

    /// The raid as the server answers a PUT: the character's state, its own row's state and minutes.
    static func echo(_ raid: RaidView, key: String, _ request: SignupRequest) throws(APIError) -> RaidView {
        do {
            let encoded = try WireFormat.encoder().encode(raid)
            guard var json = try JSONSerialization.jsonObject(with: encoded) as? [String: Any],
                  var viewer = json["viewer"] as? [String: Any],
                  let characters = viewer["characters"] as? [[String: Any]],
                  let signups = json["signups"] as? [[String: Any]]
            else { throw APIError.missingFixture("raid \(raid.id) does not encode as an object") }
            viewer["characters"] = characters.map { character in
                guard character["character_key"] as? String == key else { return character }
                var echoed = character
                echoed["state"] = request.state
                if let spec = request.spec { echoed["raid_spec"] = spec }
                return echoed
            }
            json["viewer"] = viewer
            json["signups"] = signups.map { signup in
                guard signup["character_key"] as? String == key, signup["own"] as? Bool == true else { return signup }
                var echoed = signup
                echoed["state"] = request.state
                echoed["late_minutes"] = request.lateMinutes ?? NSNull()
                return echoed
            }
            return try WireFormat.decoder().decode(RaidView.self, from: JSONSerialization.data(withJSONObject: json))
        } catch let error as APIError {
            throw error
        } catch {
            throw .missingFixture("echo: \(error)")
        }
    }
}

/// Waits (yielding the main actor) until the condition holds, or records nothing and returns false.
@MainActor
func flowSettle(_ condition: () -> Bool, tries: Int = 500) async -> Bool {
    for _ in 0 ..< tries {
        if condition() { return true }
        await Task.yield()
        try? await Task.sleep(for: .milliseconds(2))
    }
    return condition()
}

/// Whether any string the content holds (at any depth) is the line: a content-level check that a
/// screen draws a sentence, whichever field ends up carrying it.
func flowDraws(_ value: Any, _ line: String) -> Bool {
    if let text = value as? String { return text == line }
    return Mirror(reflecting: value).children.contains { flowDraws($0.value, line) }
}

// MARK: The Raid tab root

/// The Raid tab root as RaidsTab wires it (flows spec 4.2, F1): RaidsTabModel's state, the
/// RaidAnswerModel's answering state, through RaidRootComposer. `answer`, `late` and `retry` mirror
/// RaidsTab's private `answer(_:)`: the current raid, the tab character's viewer row and its note,
/// and the PUT's raid replacing the tab's list; `load` sets offline from the tab, as `reload` does.
@MainActor
final class RaidTabFlow {
    let service: EchoRaidService
    let tab: RaidsTabModel
    let answers: RaidAnswerModel
    let haptics: RecordingHaptics
    private let cache = MemoryRaidCache()

    init(service: EchoRaidService? = nil) throws {
        let service = try service ?? EchoRaidService()
        self.service = service
        let haptics = RecordingHaptics()
        self.haptics = haptics
        tab = RaidsTabModel(service: service, cache: cache, now: { RaidTime.now })
        answers = RaidAnswerModel(service: service, haptics: haptics, hold: nil)
    }

    func load() async {
        await tab.load()
        if let content { answers.offline = content.offlineSince != nil }
    }

    var root: RaidRootContent {
        RaidRootComposer(
            presenter: RaidFlowPresenter(now: RaidTime.now, clock: RaidTime.clock),
            answering: answers.state
        ).content(tab.state, extras: RaidRootExtras())
    }

    var guild: RaidGuildContent? {
        if case let .guild(guild) = root { return guild }
        return nil
    }

    var head: NextRaidContent? { guild?.next }

    private var content: RaidsTabContent? {
        if case let .loaded(content) = tab.state { return content }
        return nil
    }

    func answer(_ state: SignupState) async {
        await send { [answers] key, guild, raid, note in await answers.answer(state, guild: guild, raid: raid, key: key, note: note) }
    }

    func late(_ minutes: Int) async {
        await send { [answers] key, guild, raid, note in await answers.late(minutes, guild: guild, raid: raid, key: key, note: note) }
    }

    func retry() async {
        await send { [answers] key, guild, raid, note in await answers.retry(guild: guild, raid: raid, key: key, note: note) }
    }

    /// Starts an answer without waiting for its write (the in-flight step).
    func startAnswer(_ state: SignupState) -> Task<Void, Never> {
        Task { await answer(state) }
    }

    private func send(_ write: (String, Int64, Int64, String) async -> RaidView?) async {
        guard let content, let raid = HomeFlowModel.current(content.planned, now: RaidTime.now) else { return }
        let key = raid.viewer.characters.first { $0.characterKey == content.character.key }?.characterKey ?? content.character.key
        let note = raid.signups.first { $0.characterKey == key }?.note ?? ""
        if let view = await write(key, content.head.id, raid.id, note) { tab.replace(view) }
    }
}

// MARK: Home's raid card

struct FlowHomeService: HomeFlowService {
    func sections(character: String) async throws(APIError) -> HomeBlocks {
        HomeBlocks(sections: [HomeSection(type: "next_raid", limit: nil, style: "card")])
    }

    func activity(group: String?) async throws(APIError) -> ActivityPage { throw .transport("not in this flow") }
    func badge(guild: Int64) async throws(APIError) -> RaidBadge { throw .transport("not in this flow") }
}

struct NoReady: ReadySource {
    func ready(character: String, raid: Int64) async -> ReadyCheck? { nil }
}

struct NoApplicants: ApplicantsSource {
    func applicants(guild: Int64) async -> [ApplicantLine] { [] }
}

/// Home's raid card (flows spec 4.1, F1) through HomeFlowModel and HomeFlowPresenter, on the same
/// echoing service.
@MainActor
final class HomeCardFlow {
    let service: EchoRaidService
    let model: HomeFlowModel
    let haptics: RecordingHaptics

    init() throws {
        let service = try EchoRaidService()
        self.service = service
        let haptics = RecordingHaptics()
        self.haptics = haptics
        let me = try RaidFixture.load(Me.self, "me")
        let character = try requireCharacter(me)
        let dependencies = HomeFlowDependencies(
            service: FlowHomeService(), raids: service, ready: NoReady(), applicants: NoApplicants(),
            bestRun: { _, _ in nil }, firstRaids: { nil }, now: { RaidTime.now }
        )
        let home = HomeStore(fetch: { _ throws(APIError) in throw .transport("not in this flow") }, cache: MemoryHomeCache(), now: { RaidTime.now })
        model = HomeFlowModel(character: HomeCharacter(character), home: home, dependencies: dependencies, haptics: haptics, savedHold: nil)
    }

    func load() async { await model.load() }

    var card: NextRaidContent? {
        HomeFlowPresenter(now: RaidTime.now, clock: RaidTime.clock).nextRaid(model.inputs)
    }

    var blocks: [HomeBlockContent] {
        HomeFlowPresenter(now: RaidTime.now, clock: RaidTime.clock).blocks(model.blocks, model.inputs)
    }
}

private func requireCharacter(_ me: Me) throws -> AccountCharacter {
    guard let character = me.characters.first(where: { $0.key == FlowPlayer.key }) else {
        throw RaidFixture.Missing(name: "me", reason: "no \(FlowPlayer.key)")
    }
    return character
}
