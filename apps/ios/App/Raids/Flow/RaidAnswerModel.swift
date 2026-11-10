import Foundation
import Models
import Networking
import Observation

/// Where one answer from the Raid tab's head, Home's card or the raid night page stands (F1): the
/// answer just sent (and Late's minutes) shown before the raid reloads, whether Late's minutes are
/// open, the status line, a write in flight, and whether the connection is gone.
struct RaidAnswering: Equatable, Sendable {
    var pending: SignupState?
    var pendingMinutes: Int?
    var lateOpen = false
    var status: StatusLine?
    var busy = false
    var offline = false

    /// The line under the answers: the last write's, else offline's sentence.
    var line: StatusLine? { status ?? (offline ? .info(Copy.raidOffline) : nil) }
}

/// One raid answer (F1): In, Late, Bench or Out for a character, Late's minutes after it. The answer
/// shows at once and rolls back on failure with the error and Try again; Saved. holds for `hold`;
/// offline nothing is sent. The haptics are the raid page's (R14): selection on the press, then
/// success or error.
@MainActor
@Observable
final class RaidAnswerModel {
    private(set) var state = RaidAnswering()

    /// The answer the last failed write tried, for Try again.
    private var failed: (state: SignupState, minutes: Int?)?
    private var clear: Task<Void, Never>?

    private let service: any RaidService
    private let haptics: any Haptics
    private let hold: Duration?

    init(service: any RaidService, haptics: any Haptics = SystemHaptics(), hold: Duration? = .seconds(4)) {
        self.service = service
        self.haptics = haptics
        self.hold = hold
    }

    var pending: SignupState? { state.pending }
    var lateOpen: Bool { state.lateOpen }
    var error: String? {
        if case let .error(text) = state.status { return text }
        return nil
    }

    /// The connection is gone (the screen's own load said so): the answers disable.
    var offline: Bool {
        get { state.offline }
        set { state.offline = newValue }
    }

    /// Sends the answer; Late opens its minutes. Answers the raid as it now stands.
    func answer(_ answer: SignupState, guild: Int64, raid: Int64, key: String, note: String) async -> RaidView? {
        guard pressable() else { return nil }
        state.lateOpen = answer == .late
        return await send(answer, minutes: nil, guild: guild, raid: raid, key: key, note: note)
    }

    /// Late's minutes (15, 30, 60): they stay open with the chosen one selected.
    func late(_ minutes: Int, guild: Int64, raid: Int64, key: String, note: String) async -> RaidView? {
        guard pressable() else { return nil }
        state.lateOpen = true
        return await send(.late, minutes: minutes, guild: guild, raid: raid, key: key, note: note)
    }

    /// Try again: the answer the last write failed to save.
    func retry(guild: Int64, raid: Int64, key: String, note: String) async -> RaidView? {
        guard let failed, pressable() else { return nil }
        state.lateOpen = failed.state == .late
        return await send(failed.state, minutes: failed.minutes, guild: guild, raid: raid, key: key, note: note)
    }

    /// The press's haptic; offline it is refused with the error haptic and the offline line.
    private func pressable() -> Bool {
        guard !state.busy else { return false }
        haptics.selection()
        guard !state.offline else {
            haptics.error()
            state.status = nil
            return false
        }
        return true
    }

    private func send(_ answer: SignupState, minutes: Int?, guild: Int64, raid: Int64, key: String, note: String) async -> RaidView? {
        let previous = (state.pending, state.pendingMinutes)
        clear?.cancel()
        state.pending = answer
        state.pendingMinutes = minutes
        state.status = nil
        state.busy = true
        let request = SignupRequest(state: answer.rawValue, role: nil, note: note, spec: nil, lateMinutes: minutes)
        do {
            let saved = try await service.signup(guild: guild, raid: raid, characterKey: key, request)
            state.pending = nil
            state.pendingMinutes = nil
            state.busy = false
            state.offline = false
            failed = nil
            show(.saved(Copy.commonSaved))
            haptics.success()
            NotificationCenter.default.post(name: .raidSignupSaved, object: nil)
            return saved.raid
        } catch {
            (state.pending, state.pendingMinutes) = previous
            state.busy = false
            if SignupRefusal.isOffline(error) { state.offline = true }
            failed = (answer, minutes)
            state.status = .error(SignupRefusal.sentence(for: error))
            haptics.error()
            return nil
        }
    }

    /// Saved. for the hold, then the line clears.
    private func show(_ status: StatusLine) {
        state.status = status
        guard let hold else { return }
        clear = Task { [weak self] in
            try? await Task.sleep(for: hold)
            guard !Task.isCancelled, let self, self.state.status == status else { return }
            self.state.status = nil
        }
    }

    /// Whether Try again has an answer to resend.
    var canRetry: Bool { failed != nil && !state.offline }
}
