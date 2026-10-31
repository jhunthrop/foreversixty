import Foundation
import Models
import Networking
import Observation

/// The 20 pt status line under a character's segments.
enum StatusLine: Equatable, Sendable {
    /// `Saved.` in kill green, for 4 s.
    case saved(String)
    /// A refusal or failure in wipe red.
    case error(String)
    /// A neutral line: `{name} moved to Bench.`
    case info(String)

    var text: String {
        switch self {
        case let .saved(text), let .error(text), let .info(text): text
        }
    }
}

/// One character block's drawn state.
struct SignupBlock: Equatable, Sendable {
    /// The checked segment: optimistic while a save is out.
    var shown: SignupState?
    /// A press waiting for "Which spec tonight?" (the dashed outline).
    var pending: SignupState?
    var busy = false
    var status: StatusLine?
    var noteBusy = false
    var noteFailed = false
}

/// The spec sheet's question: a press waiting for a spec, or Change spec.
enum SpecAsk: Equatable, Sendable, Identifiable {
    case ask(characterKey: String, state: SignupState)
    case change(characterKey: String)

    var id: String {
        switch self {
        case let .ask(key, state): "ask:\(key):\(state.rawValue)"
        case let .change(key): "change:\(key)"
        }
    }

    var characterKey: String {
        switch self {
        case let .ask(key, _), let .change(key): key
        }
    }
}

/// The raid page's and the Raids tab card's sign-up, the officer tools and the reminder ask. One
/// raid view is the truth; a press draws at once, the PUT's raid replaces the view, a failure
/// restores the previous answer (spec F4.3, F4.11).
@MainActor
@Observable
final class SignupController {
    struct Dependencies {
        let service: any RaidService
        let haptics: any RaidHaptics
        let permission: any NotificationPermission
        let askStore: any ReminderAskStore
        let now: @Sendable () -> Date
        /// How long `Saved.`, `Copied` and the lock toast stay (4 s; tests pass zero to keep them).
        let hold: Duration?
    }

    let guild: Int64
    private(set) var raid: RaidView
    private(set) var blocks: [String: SignupBlock] = [:]
    private(set) var notes: [String: String] = [:]
    var specAsk: SpecAsk?
    /// The last write or refresh failed for lack of a connection: segments disable.
    var offline = false
    private(set) var rosterOverrides: [String: SignupState?] = [:]
    private(set) var rowErrors: [String: String] = [:]
    private(set) var officerBusy = false
    private(set) var officerError: String?
    private(set) var toast: String?
    private(set) var copied = false
    private(set) var showsReminderAsk = false
    /// Hears every new raid view, so the Raids tab's list follows the page.
    var onRaidChanged: ((RaidView) -> Void)?

    private let deps: Dependencies

    init(guild: Int64, raid: RaidView, dependencies: Dependencies) {
        self.guild = guild
        self.raid = raid
        self.deps = dependencies
        apply(raid)
        for signup in raid.signups where signup.own {
            notes[signup.characterKey] = signup.note
        }
    }

    var facts: RaidFacts { RaidFacts(raid: raid, now: deps.now()) }

    func block(_ key: String) -> SignupBlock { blocks[key] ?? SignupBlock() }

    func character(_ key: String) -> RaidViewerCharacter? {
        raid.viewer.characters.first { $0.characterKey == key }
    }

    /// The segments take a press: online, a verified character, a signable raid, and not locked
    /// unless the viewer may still write (an officer's own segments stay enabled).
    func segmentsEnabled(_ key: String) -> Bool {
        guard let character = character(key), character.verified, !offline else { return false }
        let facts = facts
        guard facts.isSignable else { return false }
        return facts.chip != .locked || facts.canMove
    }

    /// The status line when no press has spoken: the raid state's sentence.
    func restingStatus(_ key: String) -> StatusLine? {
        if let status = block(key).status { return status }
        if offline { return .info(Copy.raidOffline) }
        switch facts.chip {
        case .locked: return .info(facts.canMove ? Copy.officerMemberLocked : Copy.raidLocked)
        case .ended: return .info(Copy.raidEnded)
        default: return nil
        }
    }

    // MARK: Sign-up

    func press(_ key: String, _ state: SignupState) async {
        guard segmentsEnabled(key), !block(key).busy, let character = character(key) else { return }
        deps.haptics.selection()
        if state.carriesSpec, character.raidSpec == nil, (character.specOptions?.count ?? 0) >= 2 {
            blocks[key, default: SignupBlock()].pending = state
            specAsk = .ask(characterKey: key, state: state)
            return
        }
        await save(key, state: state, spec: nil)
    }

    func changeSpec(_ key: String) {
        guard segmentsEnabled(key) else { return }
        specAsk = .change(characterKey: key)
    }

    func pickSpec(_ option: RaidSpecOption) async {
        guard let ask = specAsk else { return }
        specAsk = nil
        switch ask {
        case let .ask(key, state):
            blocks[key, default: SignupBlock()].pending = nil
            await save(key, state: state, spec: option)
        case let .change(key):
            guard let state = block(key).shown else { return }
            await save(key, state: state, spec: option)
        }
    }

    /// Closing the sheet drops the press: nothing is saved.
    func closeSpecSheet() {
        if case let .ask(key, _) = specAsk { blocks[key, default: SignupBlock()].pending = nil }
        specAsk = nil
    }

    private func save(_ key: String, state: SignupState, spec: RaidSpecOption?) async {
        let previous = block(key).shown
        let name = character(key)?.name ?? ""
        blocks[key, default: SignupBlock()].shown = state
        blocks[key, default: SignupBlock()].busy = true
        blocks[key, default: SignupBlock()].status = nil
        let request = SignupRequest(state: state.rawValue, role: nil, note: notes[key] ?? "", spec: spec?.spec)
        do {
            let answer = try await deps.service.signup(guild: guild, raid: raid.id, characterKey: key, request)
            offline = false
            apply(answer.raid)
            let kept = spec.map { " " + Copy.specKept(spec: $0.name, name: name) } ?? ""
            setStatus(.saved(Copy.commonSaved + kept), for: key)
            if let moved = answer.moved, let other = raid.viewer.characters.first(where: { $0.name == moved && $0.characterKey != key }) {
                blocks[other.characterKey, default: SignupBlock()].status = .info(Copy.raidMoved(name: moved))
            }
            deps.haptics.success()
            if state.isComing { await considerReminderAsk() }
        } catch {
            blocks[key, default: SignupBlock()].shown = previous
            blocks[key, default: SignupBlock()].busy = false
            blocks[key, default: SignupBlock()].status = .error(SignupRefusal.sentence(for: error))
            offline = SignupRefusal.isOffline(error)
            deps.haptics.error()
        }
    }

    // MARK: Note

    func editNote(_ key: String, _ text: String) {
        notes[key] = String(text.prefix(Self.noteLimit))
    }

    static let noteLimit = 140
    static let noteCounterFrom = 120

    /// Saves the note with the current answer (Return or blur). The PUT writes note and state.
    func saveNote(_ key: String) async {
        guard let state = block(key).shown, segmentsEnabled(key) else { return }
        let stored = raid.signups.first { $0.characterKey == key }?.note ?? ""
        guard (notes[key] ?? "") != stored else { return }
        blocks[key, default: SignupBlock()].noteBusy = true
        blocks[key, default: SignupBlock()].noteFailed = false
        let request = SignupRequest(state: state.rawValue, role: nil, note: notes[key] ?? "", spec: nil)
        do {
            let answer = try await deps.service.signup(guild: guild, raid: raid.id, characterKey: key, request)
            apply(answer.raid)
        } catch {
            blocks[key, default: SignupBlock()].noteFailed = true
        }
        blocks[key, default: SignupBlock()].noteBusy = false
    }

    // MARK: Reminder ask

    private func considerReminderAsk() async {
        let now = deps.now()
        guard !showsReminderAsk, deps.askStore.policy().mayAsk(now: now), await deps.permission.isUndetermined() else { return }
        deps.askStore.recordShown(at: now)
        showsReminderAsk = true
    }

    func answerReminderAsk(turnOn: Bool) async {
        showsReminderAsk = false
        if turnOn { await deps.permission.request() }
    }

    // MARK: Officer tools

    /// An officer's move (spec F4.20): the officer PUT for the row's character, with its role.
    func move(_ row: RosterRow, to state: SignupState) async {
        guard facts.canMove, !offline else { return }
        rosterOverrides[row.characterKey] = .some(state)
        rowErrors[row.characterKey] = nil
        let request = SignupRequest(state: state.rawValue, role: row.role, note: row.note, spec: nil)
        do {
            let answer = try await deps.service.signup(guild: guild, raid: raid.id, characterKey: row.characterKey, request)
            rosterOverrides[row.characterKey] = nil
            apply(answer.raid)
            deps.haptics.success()
        } catch {
            rosterOverrides[row.characterKey] = nil
            rowErrors[row.characterKey] = SignupRefusal.sentence(for: error)
            deps.haptics.error()
        }
    }

    func toggleLock() async {
        guard facts.canLock, !officerBusy else { return }
        let locking = facts.chip != .locked
        officerBusy = true
        officerError = nil
        do {
            let view = locking
                ? try await deps.service.lock(guild: guild, raid: raid.id)
                : try await deps.service.unlock(guild: guild, raid: raid.id)
            apply(view)
            if locking {
                show(toast: Copy.officerLockedToast(n: view.counts.signedIn, b: view.counts.bench ?? 0))
            }
            deps.haptics.success()
        } catch {
            officerError = SignupRefusal.sentence(for: error)
            deps.haptics.error()
        }
        officerBusy = false
    }

    func copyList(clock: RaidClock) -> String {
        flash(copied: true)
        return RosterCopy.text(raid: raid, clock: clock, now: deps.now())
    }

    // MARK: Refresh

    /// A refetch replaces the view unless a save is out (it never moves a focused control).
    func refresh() async {
        guard !isWriting else { return }
        do {
            replace(with: try await deps.service.raid(guild: guild, raid: raid.id))
        } catch {
            offline = SignupRefusal.isOffline(error)
        }
    }

    /// A newer view from elsewhere (a refetch, the cache's successor); skipped while a write is out.
    func replace(with view: RaidView) {
        guard !isWriting else { return }
        offline = false
        apply(view)
    }

    private var isWriting: Bool { blocks.values.contains(where: \.busy) || !rosterOverrides.isEmpty }

    // MARK: State

    private func apply(_ view: RaidView) {
        raid = view
        for character in view.viewer.characters {
            var block = blocks[character.characterKey] ?? SignupBlock()
            block.shown = character.state.flatMap(SignupState.init)
            block.busy = false
            blocks[character.characterKey] = block
        }
        onRaidChanged?(view)
    }

    private func setStatus(_ status: StatusLine, for key: String) {
        blocks[key, default: SignupBlock()].status = status
        guard let hold = deps.hold else { return }
        Task { [weak self] in
            try? await Task.sleep(for: hold)
            guard let self, self.blocks[key]?.status == status else { return }
            self.blocks[key]?.status = nil
        }
    }

    private func show(toast text: String) {
        toast = text
        guard let hold = deps.hold else { return }
        Task { [weak self] in
            try? await Task.sleep(for: hold)
            if self?.toast == text { self?.toast = nil }
        }
    }

    private func flash(copied value: Bool) {
        copied = value
        guard let hold = deps.hold else { return }
        Task { [weak self] in
            try? await Task.sleep(for: hold)
            self?.copied = false
        }
    }
}
