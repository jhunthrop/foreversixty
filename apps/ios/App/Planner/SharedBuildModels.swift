import Foundation
import Models
import Networking
import Observation

/// A saved build as the viewer and Your builds print it: names from its own tree version's
/// files, its totals summed from its own trees.
struct PlannerBuildSummary: Equatable, Sendable, Identifiable {
    let build: Build
    let classSlug: String
    let className: String
    let raceName: String
    let raceSlug: String?
    /// Points per tree; empty when the build's talent file did not load.
    let totals: [Int]

    var id: String { build.id }
    var totalsText: String { PlannerCopy.totals(totals) }

    /// `Orc Shaman 30 / 3 / 18`.
    var head: String { Copy.planSharedHead(race: raceName, className: className, totals: totalsText) }

    /// The row's title: the build's own title, else its head.
    var rowTitle: String {
        build.title.flatMap { $0.isEmpty ? nil : $0 }
            ?? Copy.planBuildsRowTitle(race: raceName, className: className, totals: totalsText)
    }

    func rowMeta(timeZone: TimeZone = .current) -> String {
        Copy.planBuildsMeta(race: raceName, className: className, totals: totalsText, date: PlannerCopy.day(build.createdAt, timeZone: timeZone))
    }

    func viewerMeta(timeZone: TimeZone = .current) -> String {
        let date = PlannerCopy.day(build.createdAt, timeZone: timeZone)
        let views = Int(build.views)
        guard let title = build.title, !title.isEmpty else { return Copy.planSharedMetaNoTitle(date: date, n: views) }
        return Copy.planSharedMeta(title: title, date: date, n: views)
    }

    /// The POST body that puts a removed build back on the list (spec F13.6 Undo).
    var input: BuildInput {
        BuildInput(
            classId: build.classId, raceId: build.raceId, treeVersion: build.treeVersion,
            pointOrder: build.pointOrder, gear: build.gear, title: build.title
        )
    }
}

/// Resolves a build's names and trees from its own `tree_version`'s files, falling back to the
/// active build's names when that version's files are gone.
struct PlannerSummarizer: Sendable {
    let source: any PlannerSource

    func summary(_ build: Build) async -> (PlannerBuildSummary, TalentsFile?) {
        let names = await names(build.treeVersion)
        let klass = names.klass(id: build.classId)
        let slug = klass?.slug ?? ""
        let file = slug.isEmpty ? nil : try? await source.talents(build: build.treeVersion, classSlug: slug)
        let totals = file.map { PlannerRules.pointsPerTree(PlannerIndex($0), order: build.pointOrder) } ?? []
        let summary = PlannerBuildSummary(
            build: build, classSlug: slug, className: klass?.name ?? PlannerCopy.className(slug, names: names),
            raceName: names.race(id: build.raceId)?.name ?? "", raceSlug: names.race(id: build.raceId)?.slug, totals: totals
        )
        return (summary, file)
    }

    private func names(_ version: String) async -> PlannerNames {
        if let classes = try? await source.classes(build: version), let races = try? await source.races(build: version) {
            return PlannerNames(races: races, classes: classes)
        }
        let active = await source.activeBuild()
        return PlannerNames(
            races: (try? await source.races(build: active)) ?? [], classes: (try? await source.classes(build: active)) ?? []
        )
    }
}

/// The viewer's states (spec F13.5).
enum SharedBuildPhase: Equatable, Sendable {
    case loading
    case loaded(SharedBuildContent)
    /// Unknown id: `This build does not exist.`
    case notFound
    /// Offline with no saved copy.
    case offline
    case failed
}

struct SharedBuildContent: Equatable, Sendable {
    let summary: PlannerBuildSummary
    /// The build's own trees (its `tree_version`, never the active build's); nil when that file
    /// did not load, and then only the head draws.
    let file: TalentsFile?
}

/// F13.5: one shared build, read only.
@MainActor
@Observable
final class SharedBuildModel {
    private(set) var phase: SharedBuildPhase = .loading
    let id: String
    private let source: any PlannerSource
    private let cache: any RaidCache

    init(id: String, source: any PlannerSource, cache: any RaidCache) {
        self.id = id
        self.source = source
        self.cache = cache
    }

    init(id: String, source: any PlannerSource, cache: any RaidCache, phase: SharedBuildPhase) {
        self.id = id
        self.source = source
        self.cache = cache
        self.phase = phase
    }

    private var cacheKey: String { "build-\(id)" }

    func load() async {
        phase = .loading
        let build: Build
        do {
            build = try await source.build(id: id)
            cache.write(build, key: cacheKey, at: Date())
        } catch {
            if error.isMissingFile {
                phase = .notFound
                return
            }
            guard let saved = cache.read(Build.self, key: cacheKey)?.value else {
                phase = Self.isOffline(error) ? .offline : .failed
                return
            }
            build = saved
        }
        let (summary, file) = await PlannerSummarizer(source: source).summary(build)
        phase = .loaded(SharedBuildContent(summary: summary, file: file))
    }

    private static func isOffline(_ error: APIError) -> Bool {
        if case .transport = error { return true }
        return false
    }
}

/// Your builds' states (spec F13.6).
enum YourBuildsPhase: Equatable, Sendable {
    case loading
    case loaded([PlannerBuildSummary])
    case failed
}

/// The `Removed.` snackbar: the row taken off and where it stood.
struct RemovedBuild: Equatable, Sendable {
    let summary: PlannerBuildSummary
    let position: Int
}

/// F13.6: the signed-in player's shared builds; swipe removes from the list, Undo puts it back.
@MainActor
@Observable
final class YourBuildsModel {
    /// How long the snackbar stays (spec F13.6).
    static let snackbarSeconds: Double = 5

    private(set) var phase: YourBuildsPhase = .loading
    private(set) var removed: RemovedBuild?
    /// A removal or an undo that failed, said under the list.
    private(set) var error: String?
    private let source: any PlannerSource

    init(source: any PlannerSource, phase: YourBuildsPhase = .loading, removed: RemovedBuild? = nil) {
        self.source = source
        self.phase = phase
        self.removed = removed
    }

    private var rows: [PlannerBuildSummary] {
        if case let .loaded(rows) = phase { return rows }
        return []
    }

    func load() async {
        if case .failed = phase { phase = .loading }
        do {
            let builds = try await source.myBuilds()
            let summarizer = PlannerSummarizer(source: source)
            var summaries: [PlannerBuildSummary] = []
            for build in builds { summaries.append(await summarizer.summary(build).0) }
            phase = .loaded(summaries)
        } catch {
            if rows.isEmpty { phase = .failed }
        }
    }

    /// Swipe: the row goes at once, then DELETE; a failed DELETE puts it back and says so.
    func remove(_ summary: PlannerBuildSummary) async {
        guard let position = rows.firstIndex(of: summary) else { return }
        var next = rows
        next.remove(at: position)
        phase = .loaded(next)
        removed = RemovedBuild(summary: summary, position: position)
        error = nil
        do {
            try await source.remove(id: summary.id)
        } catch {
            restore(summary, at: position)
            removed = nil
            self.error = Copy.planRemoveFailed
        }
    }

    /// Undo: saving the same build again puts it back on the list.
    func undo() async {
        guard let removed else { return }
        self.removed = nil
        do {
            _ = try await source.save(removed.summary.input)
            restore(removed.summary, at: removed.position)
        } catch {
            self.error = Copy.planShareFailed
        }
    }

    func dismissSnackbar() { removed = nil }

    private func restore(_ summary: PlannerBuildSummary, at position: Int) {
        var next = rows.filter { $0.id != summary.id }
        next.insert(summary, at: min(position, next.count))
        phase = .loaded(next)
    }
}
