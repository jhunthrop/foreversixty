import Foundation
import Models
import Networking
import Observation

/// A screen's load: the answer, or what failed.
enum RaidFlowPhase<Value: Equatable>: Equatable {
    case loading
    case loaded(Value)
    case failed
}

/// Not answered (F3): the raid, the filter, and the raiders marked on the page. A mark waits the
/// hold on the row with Undo before it is sent, so a mis-tap is never a write.
@MainActor
@Observable
final class NotAnsweredModel {
    private(set) var phase: RaidFlowPhase<RaidView> = .loading
    var filter: NotAnsweredFilter = .notAnswered
    private(set) var marks: [String: SignupState] = [:]
    private(set) var error: String?

    let guild: Int64
    let raidID: Int64
    private let service: any RaidService
    private let hold: Duration
    private var holds: [String: Task<Void, Never>] = [:]

    init(guild: Int64, raid: Int64, service: any RaidService, hold: Duration) {
        self.guild = guild
        raidID = raid
        self.service = service
        self.hold = hold
    }

    var raid: RaidView? {
        if case let .loaded(raid) = phase { return raid }
        return nil
    }

    func load() async {
        do {
            phase = .loaded(try await service.raid(guild: guild, raid: raidID))
        } catch {
            if raid == nil { phase = .failed }
        }
    }

    /// Mark In or Bench: shown on the row at once, sent once the hold runs out.
    func mark(_ key: String, _ state: SignupState) {
        holds[key]?.cancel()
        marks[key] = state
        error = nil
        holds[key] = Task { [hold] in
            try? await Task.sleep(for: hold)
            guard !Task.isCancelled else { return }
            await self.send(key, state)
        }
    }

    func undo(_ key: String) {
        holds[key]?.cancel()
        holds[key] = nil
        marks[key] = nil
    }

    private func send(_ key: String, _ state: SignupState) async {
        holds[key] = nil
        let request = SignupRequest(state: state.rawValue, role: nil, note: nil, spec: nil, lateMinutes: nil)
        do {
            let answer = try await service.signup(guild: guild, raid: raidID, characterKey: key, request)
            marks[key] = nil
            phase = .loaded(answer.raid)
            NotificationCenter.default.post(name: .raidSignupSaved, object: nil)
        } catch {
            marks[key] = nil
            self.error = SignupRefusal.sentence(for: error)
        }
    }
}

/// The applicant (F3): the application from the officers' list, Accept or Decline.
@MainActor
@Observable
final class ApplicantModel {
    private(set) var phase: RaidFlowPhase<GuildOfficerApplication> = .loading
    private(set) var accepted = false
    private(set) var busy = false
    private(set) var error: String?

    let guild: Int64
    let applicationID: Int64
    private let service: any RaidFlowService

    init(guild: Int64, application: Int64, service: any RaidFlowService) {
        self.guild = guild
        applicationID = application
        self.service = service
    }

    func load() async {
        do {
            let page = try await service.applications(guild: guild)
            guard let row = page.rows.first(where: { $0.id == applicationID }) else { phase = .failed; return }
            accepted = row.status == "accepted"
            phase = .loaded(row)
        } catch {
            phase = .failed
        }
    }

    func accept() async {
        await decide { [service, guild, applicationID] () async throws(APIError) in
            try await service.accept(guild: guild, application: applicationID)
        }
        if error == nil { accepted = true }
    }

    /// Declines; true once the server took it (the screen goes back).
    func decline() async -> Bool {
        await decide { [service, guild, applicationID] () async throws(APIError) in
            try await service.decline(guild: guild, application: applicationID)
        }
        return error == nil
    }

    private func decide(_ call: @escaping () async throws(APIError) -> GuildApplicationDecision) async {
        busy = true
        error = nil
        do {
            phase = .loaded(try await call().application)
        } catch {
            self.error = SignupRefusal.sentence(for: error)
        }
        busy = false
    }
}

/// The raid page on raid night (F2): `GET /v1/me/raid-nights/{raid}` for who is coming and the
/// Ready check, the raid itself for the comp target and its time zone, and your answer.
@MainActor
@Observable
final class RaidNightPageModel {
    private(set) var phase: RaidFlowPhase<RaidNightSummary> = .loading
    private(set) var raid: RaidView?
    let answers: RaidAnswerModel

    let raidID: Int64
    private let service: any RaidService

    init(raid: Int64, service: any RaidService) {
        raidID = raid
        self.service = service
        answers = RaidAnswerModel(service: service)
    }

    /// The night names its guild; the raid follows from it (an after-raid night needs no raid).
    var guild: Int64? {
        if case let .loaded(night) = phase { return night.guildId }
        return nil
    }

    func load() async {
        guard let night = try? await service.night(raid: raidID) else {
            if case .loading = phase { phase = .failed }
            return
        }
        if night.phase != RaidNightPageScreen.afterPhase, let raid = try? await service.raid(guild: night.guildId, raid: raidID) {
            self.raid = raid
        }
        phase = .loaded(night)
    }

    func answer(_ state: SignupState) async {
        guard case let .loaded(night) = phase else { return }
        if await answers.answer(state, guild: night.guildId, raid: raidID, key: night.character.characterKey, note: night.you.note) != nil {
            await load()
        }
    }

    func late(_ minutes: Int) async {
        guard case let .loaded(night) = phase else { return }
        if await answers.late(minutes, guild: night.guildId, raid: raidID, key: night.character.characterKey, note: night.you.note) != nil {
            await load()
        }
    }
}
