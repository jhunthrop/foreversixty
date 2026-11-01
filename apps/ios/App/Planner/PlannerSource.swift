import Foundation
import Models
import Networking
import UIKit

/// What the planner reads and writes: the static talent, race, combo and class files through
/// the offline cache, and the builds API. A protocol so every model is tested without a network.
protocol PlannerSource: Sendable {
    func activeBuild() async -> String
    func talents(build: String, classSlug: String) async throws(APIError) -> TalentsFile
    func races(build: String) async throws(APIError) -> [RaceInfo]
    func combos(build: String) async throws(APIError) -> [RaceClassCombo]
    func classes(build: String) async throws(APIError) -> [ClassInfo]
    func save(_ input: BuildInput) async throws(APIError) -> BuildSaved
    func build(id: String) async throws(APIError) -> Build
    func myBuilds() async throws(APIError) -> [Build]
    func remove(id: String) async throws(APIError)
}

/// The live source: the same CDN cache and active build the Talents screen reads, and the API.
struct LivePlannerSource: PlannerSource {
    let files: CachedStaticDataClient
    let talentsSource: LiveTalentsSource
    let client: APIClient

    func activeBuild() async -> String { await talentsSource.activeBuild() }

    func talents(build: String, classSlug: String) async throws(APIError) -> TalentsFile {
        try await talentsSource.talents(build: build, classSlug: classSlug).value
    }

    func races(build: String) async throws(APIError) -> [RaceInfo] { try await files.load(.getRaces(build: build)).value }
    func combos(build: String) async throws(APIError) -> [RaceClassCombo] { try await files.load(.getCombos(build: build)).value }
    func classes(build: String) async throws(APIError) -> [ClassInfo] { try await files.load(.getClasses(build: build)).value }
    func save(_ input: BuildInput) async throws(APIError) -> BuildSaved { try await client.send(.saveBuild(input)) }
    func build(id: String) async throws(APIError) -> Build { try await client.send(.getBuild(id: id)) }
    func myBuilds() async throws(APIError) -> [Build] { try await client.send(.listMyBuilds()).rows }
    func remove(id: String) async throws(APIError) { _ = try await client.send(.removeBuild(id: id)) }
}

/// The planner's haptics (spec F13.1 Interactions; the web's SEGMENT_TICK, LONG_PRESS, REJECT and
/// CONFIRM): selection on an added point, a medium impact on a removed one, error on a refusal,
/// success on a share.
@MainActor
protocol PlannerHaptics {
    func play(_ haptic: PlannerHaptic)
}

enum PlannerHaptic: Equatable, Sendable {
    case segmentTick, longPress, reject, confirm

    static func of(_ event: PlannerEvent) -> PlannerHaptic {
        switch event {
        case .added: .segmentTick
        case .removed: .longPress
        case .refused: .reject
        }
    }
}

@MainActor
struct SystemPlannerHaptics: PlannerHaptics {
    func play(_ haptic: PlannerHaptic) {
        switch haptic {
        case .segmentTick: UISelectionFeedbackGenerator().selectionChanged()
        case .longPress: UIImpactFeedbackGenerator(style: .medium).impactOccurred()
        case .reject: UINotificationFeedbackGenerator().notificationOccurred(.error)
        case .confirm: UINotificationFeedbackGenerator().notificationOccurred(.success)
        }
    }
}

/// The planner's dependencies, built once from the environment.
struct PlannerDependencies {
    let source: any PlannerSource
    let drafts: any PlannerDraftStore
    let haptics: any PlannerHaptics
    /// Saved shared builds, so a viewer opened once opens offline.
    let cache: any RaidCache

    @MainActor
    static func live(_ environment: AppEnvironment) -> PlannerDependencies {
        let talents = LiveTalentsSource.make(environment)
        let base = FileManager.default.urls(for: .applicationSupportDirectory, in: .userDomainMask).first
            ?? FileManager.default.temporaryDirectory
        return PlannerDependencies(
            source: LivePlannerSource(files: talents.files, talentsSource: talents, client: environment.client),
            drafts: UserDefaultsPlannerDraftStore(),
            haptics: SystemPlannerHaptics(),
            cache: DiskRaidCache(directory: base.appending(path: "builds", directoryHint: .isDirectory))
        )
    }
}

/// The race and class names the planner prints, from a build's `races.json` and `classes.json`.
struct PlannerNames: Equatable, Sendable {
    let races: [RaceInfo]
    let classes: [ClassInfo]

    static let empty = PlannerNames(races: [], classes: [])

    func race(id: Int) -> RaceInfo? { races.first { $0.id == id } }
    func race(slug: String?) -> RaceInfo? { races.first { $0.slug == slug } }
    func klass(id: Int) -> ClassInfo? { classes.first { $0.id == id } }

    /// The class's legal, non-placeholder races in `races.json` order (spec F13.1 Race chip).
    static func legalRaces(classID: Int, races: [RaceInfo], combos: [RaceClassCombo]) -> [RaceInfo] {
        let legal = Set(combos.filter { $0.classId == classID }.map(\.raceId))
        return races.filter { legal.contains($0.id) && !$0.placeholder }
    }
}

/// The planner's words built from data: `30 / 3 / 18`, the share text, the saved day.
enum PlannerCopy {
    static let totalsSeparator = " / "

    static func totals(_ totals: [Int]) -> String { totals.map(String.init).joined(separator: totalsSeparator) }

    /// `4 Nov` (spec F13.5, F13.6 `{d Mon}`).
    static func day(_ date: Date, timeZone: TimeZone = .current) -> String {
        let formatter = DateFormatter()
        formatter.locale = Locale(identifier: "en_GB")
        formatter.timeZone = timeZone
        formatter.dateFormat = "d MMM"
        return formatter.string(from: date)
    }

    /// The class name from a slug when no classes file names it: `shaman` to `Shaman`.
    static func className(_ slug: String, names: PlannerNames) -> String {
        names.classes.first { $0.slug == slug }?.name ?? slug.capitalized
    }
}
