import Foundation
import Models
import Networking
import Observation

/// The editor's load states (spec F13.1): the talent file is the one thing it cannot draw without.
enum PlannerPhase: Equatable, Sendable {
    case loading
    case ready
    case failed
}

/// The share step's states (spec F13.3, F13.4).
enum PlannerShareState: Equatable, Sendable {
    case idle
    case sending
    /// 201 or 200: the system share sheet opens on this link.
    case shared(URL)
    /// 429, offline or anything else, said in the confirm sheet.
    case failed(String)
}

/// The F13.1 editor: loads the class's trees, races and draft, applies every change through
/// `PlannerState`, plays its haptic, saves the draft as you go, and shares.
@MainActor
@Observable
final class PlannerEditorModel {
    /// The share sheet's title field limit (spec F13.3).
    static let titleLimit = 60

    private(set) var phase: PlannerPhase = .loading
    private(set) var state: PlannerState?
    private(set) var races: [RaceInfo] = []
    private(set) var share: PlannerShareState = .idle

    let launch: PlannerLaunch
    private let dependencies: PlannerDependencies

    init(launch: PlannerLaunch, dependencies: PlannerDependencies) {
        self.launch = launch
        self.dependencies = dependencies
    }

    /// A model already drawn, for snapshots and tests.
    init(launch: PlannerLaunch, dependencies: PlannerDependencies, state: PlannerState, races: [RaceInfo]) {
        self.launch = launch
        self.dependencies = dependencies
        self.state = state
        self.races = races
        phase = .ready
    }

    private var haptics: any PlannerHaptics { dependencies.haptics }

    // MARK: Loading

    func load() async {
        phase = .loading
        let source = dependencies.source
        let build = await source.activeBuild()
        let file: TalentsFile
        do {
            file = try await source.talents(build: build, classSlug: launch.classSlug)
        } catch {
            phase = .failed
            return
        }
        let allRaces = (try? await source.races(build: build)) ?? []
        let combos = (try? await source.combos(build: build)) ?? []
        races = PlannerNames.legalRaces(classID: file.classId, races: allRaces, combos: combos)
        state = PlannerState.opening(
            launch, index: PlannerIndex(file), draft: dependencies.drafts.load(classSlug: launch.classSlug),
            legalRaces: races.map(\.slug)
        )
        phase = .ready
        saveDraft()
    }

    // MARK: Changes

    func add(_ id: Int) { apply { $0.add(id) } }
    func remove(_ id: Int) { apply { $0.remove(id) } }

    func undo() { change { $0.undo() } }
    func reset() { change { $0.reset() } }
    func setLevel(_ level: Int) { change { $0.setLevel(level) } }
    func setRace(_ slug: String) { change { $0.setRace(slug) } }
    func moveFocus(by step: Int) { change { $0.moveFocus(by: step) } }

    /// The strip's `+` and `−`, and the keyboard's `+` and `-`, on the strip's talent.
    func addFocused() { if let id = state?.focus?.id { add(id) } }
    func removeFocused() { if let id = state?.focus?.id { remove(id) } }

    private func apply(_ body: (inout PlannerState) -> PlannerEvent) {
        guard var next = state else { return }
        let event = body(&next)
        state = next
        haptics.play(PlannerHaptic.of(event))
        if event != .refused { saveDraft() }
    }

    private func change(_ body: (inout PlannerState) -> Void) {
        guard var next = state else { return }
        body(&next)
        state = next
        saveDraft()
    }

    private func saveDraft() {
        guard let state else { return }
        dependencies.drafts.save(state.draft, classSlug: launch.classSlug)
    }

    var raceName: String? { races.first { $0.slug == state?.raceSlug }?.name }

    // MARK: Sharing

    /// The POST body (spec F13.3): class, race, the active build's trees, the order, no gear.
    func shareInput(title: String) -> BuildInput? {
        guard let state, let race = races.first(where: { $0.slug == state.raceSlug }) else { return nil }
        let trimmed = title.trimmingCharacters(in: .whitespacesAndNewlines)
        return BuildInput(
            classId: state.index.file.classId, raceId: race.id, treeVersion: state.index.file.build,
            pointOrder: state.order, gear: [:], title: trimmed.isEmpty ? nil : String(trimmed.prefix(Self.titleLimit))
        )
    }

    /// Share anyway. Never queued offline: a failed write says so and the build stays here.
    func share(title: String) async {
        guard let input = shareInput(title: title) else {
            share = .failed(Copy.planShareFailed)
            return
        }
        share = .sending
        do {
            let saved = try await dependencies.source.save(input)
            guard let url = URL(string: saved.url) else {
                share = .failed(Copy.planShareFailed)
                return
            }
            share = .shared(url)
            haptics.play(.confirm)
        } catch {
            handleShareError(error)
        }
    }

    private func handleShareError(_ error: APIError) {
        if let field = Self.fieldSentence(error) {
            // 400: the field's sentence goes under the editor (spec F13.4).
            state?.show(refusal: field)
            share = .idle
            haptics.play(.reject)
            return
        }
        share = .failed(Self.shareMessage(error))
    }

    func dismissShare() { share = .idle }

    /// A 400's first field sentence, lowest point index first.
    static func fieldSentence(_ error: APIError) -> String? {
        guard case let .server(400, body, _) = error else { return nil }
        guard let fields = body.fields, !fields.isEmpty else { return body.message }
        return fields.sorted { pointIndex($0.key) < pointIndex($1.key) }.first?.value
    }

    private static func pointIndex(_ field: String) -> Int {
        Int(field.filter(\.isNumber)) ?? Int.max
    }

    static func shareMessage(_ error: APIError) -> String {
        switch error {
        case .transport: Copy.planOffline
        case .server(429, _, _), .unexpectedStatus(429): Copy.planRateLimited
        default: Copy.planShareFailed
        }
    }
}
