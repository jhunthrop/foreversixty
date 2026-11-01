@testable import ForeverSixty
import Foundation
import Models
import Networking

/// The planner tests' data: the active build's real shaman file (icons and rank text included)
/// and the build's own races and combos, read in place.
enum PlannerTestData {
    static let build = BisTestData.build

    static func shaman() throws -> TalentsFile { try TalentTestData.shaman() }
    static func warrior() throws -> TalentsFile { try TalentTestData.warrior() }

    static func races() throws -> [RaceInfo] { try data("races") }
    static func combos() throws -> [RaceClassCombo] { try data("combos") }
    static func classes() throws -> [ClassInfo] { try BisTestData.classes() }

    private static func data<T: Decodable>(_ name: String) throws -> T {
        let url = BisTestData.repo.appending(path: "data/builds/\(build)/\(name).json")
        return try WireFormat.decoder().decode(T.self, from: Data(contentsOf: url))
    }

    static func talentID(_ file: TalentsFile, _ name: String) -> Int {
        TalentRules.trees(file).flatMap(\.talents).first { $0.name == name }?.id ?? -1
    }

    /// The Elemental guide build (`TalentTestData.shamanGuide`) as one legal order: 51 points.
    static func elementalOrder(_ index: PlannerIndex) -> [Int] {
        PlannerRules.orderFromRanks(index, treeRanks: TalentRules.ranks(ofBand: TalentTestData.shamanGuide)).order
    }

    static let savedAt = Date(timeIntervalSince1970: 1_793_750_400) // 4 Nov 2026 00:00 UTC

    static func build(id: String = "k7q2m9xa", order: [Int], title: String? = "Elemental raid", views: Int64 = 12) -> Build {
        Build(id: id, classId: 7, raceId: 2, treeVersion: Self.build, pointOrder: order, gear: [:], title: title, createdAt: savedAt, views: views)
    }
}

/// An in-memory source: the files above, and builds answered or failed on demand.
final class FakePlannerSource: PlannerSource, @unchecked Sendable {
    var talentFiles: [String: TalentsFile] = [:]
    var races: [RaceInfo] = []
    var combos: [RaceClassCombo] = []
    var classes: [ClassInfo] = []
    var builds: [String: Build] = [:]
    var mine: [Build] = []
    var saveResult: Result<BuildSaved, APIError> = .success(BuildSaved(id: "k7q2m9xa", url: "https://foreversixty.gg/b/k7q2m9xa"))
    var buildError: APIError?
    var removeError: APIError?
    private(set) var saved: [BuildInput] = []
    private(set) var removed: [String] = []

    func activeBuild() async -> String { PlannerTestData.build }

    func talents(build _: String, classSlug: String) async throws(APIError) -> TalentsFile {
        guard let file = talentFiles[classSlug] else { throw .unexpectedStatus(404) }
        return file
    }

    func races(build _: String) async throws(APIError) -> [RaceInfo] { races }
    func combos(build _: String) async throws(APIError) -> [RaceClassCombo] { combos }
    func classes(build _: String) async throws(APIError) -> [ClassInfo] { classes }

    func save(_ input: BuildInput) async throws(APIError) -> BuildSaved {
        saved.append(input)
        return try saveResult.get()
    }

    func build(id: String) async throws(APIError) -> Build {
        if let buildError { throw buildError }
        guard let build = builds[id] else { throw .unexpectedStatus(404) }
        return build
    }

    func myBuilds() async throws(APIError) -> [Build] { mine }

    func remove(id: String) async throws(APIError) {
        removed.append(id)
        if let removeError { throw removeError }
    }

    @MainActor
    static func shaman() throws -> FakePlannerSource {
        let source = FakePlannerSource()
        source.talentFiles = ["shaman": try PlannerTestData.shaman(), "warrior": try PlannerTestData.warrior()]
        source.races = try PlannerTestData.races()
        source.combos = try PlannerTestData.combos()
        source.classes = try PlannerTestData.classes()
        return source
    }
}

/// Records every haptic the planner plays, through the app's one facade.
typealias RecordingPlannerHaptics = RecordingHaptics

extension RecordingHaptics {
    var played: [HapticMoment] { moments }
}

final class MemoryPlannerDraftStore: PlannerDraftStore, @unchecked Sendable {
    private(set) var drafts: [String: PlannerDraft] = [:]
    func load(classSlug: String) -> PlannerDraft? { drafts[classSlug] }
    func save(_ draft: PlannerDraft, classSlug: String) { drafts[classSlug] = draft }
}

@MainActor
extension PlannerDependencies {
    static func test(
        source: any PlannerSource, drafts: any PlannerDraftStore = MemoryPlannerDraftStore(),
        haptics: any Haptics = RecordingPlannerHaptics(), cache: any RaidCache = MemoryRaidCache()
    ) -> PlannerDependencies {
        PlannerDependencies(source: source, drafts: drafts, haptics: haptics, cache: cache)
    }
}
