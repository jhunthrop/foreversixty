import Foundation
import Models
import Networking
import Observation

/// One raid answer from the Raid tab's card or the raid night page (F1): In, Late, Bench or Out
/// for a character, Late's minutes after it; the answer shows at once and rolls back on failure.
@MainActor
@Observable
final class RaidAnswerModel {
    /// The answer just sent, shown before the raid reloads.
    private(set) var pending: SignupState?
    private(set) var lateOpen = false
    private(set) var error: String?

    private let service: any RaidService

    init(service: any RaidService) {
        self.service = service
    }

    /// Sends the answer; Late opens its minutes. Answers the raid as it now stands.
    func answer(_ state: SignupState, guild: Int64, raid: Int64, key: String, note: String) async -> RaidView? {
        lateOpen = state == .late
        return await send(state, minutes: nil, guild: guild, raid: raid, key: key, note: note)
    }

    /// Late's minutes (15, 30, 60); the chips close.
    func late(_ minutes: Int, guild: Int64, raid: Int64, key: String, note: String) async -> RaidView? {
        lateOpen = false
        return await send(.late, minutes: minutes, guild: guild, raid: raid, key: key, note: note)
    }

    private func send(_ state: SignupState, minutes: Int?, guild: Int64, raid: Int64, key: String, note: String) async -> RaidView? {
        let previous = pending
        pending = state
        error = nil
        let request = SignupRequest(state: state.rawValue, role: nil, note: note, spec: nil, lateMinutes: minutes)
        do {
            let answer = try await service.signup(guild: guild, raid: raid, characterKey: key, request)
            pending = nil
            NotificationCenter.default.post(name: .raidSignupSaved, object: nil)
            return answer.raid
        } catch {
            pending = previous
            self.error = SignupRefusal.sentence(for: error)
            return nil
        }
    }
}
