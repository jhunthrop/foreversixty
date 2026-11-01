import Foundation
import Models
import Networking
import Observation

/// What a Guides region draws from one static file (v1 §5.0): skeletons, the file (with the
/// offline stamp when it is the saved copy), the not-saved sentence, or the error.
enum StaticPhase<Value: Equatable & Sendable>: Equatable, Sendable {
    case loading
    case loaded(Value, offlineSince: Date?)
    /// Offline, and this file was never saved.
    case notCached
    case failed

    var value: Value? {
        if case let .loaded(value, _) = self { return value }
        return nil
    }

    /// The saved copy, drawn at once before the request.
    static func saved(_ file: CachedFile<Value>?) -> StaticPhase? {
        file.map { .loaded($0.value, offlineSince: nil) }
    }

    /// The phase after a request: the fresh file, the saved copy when offline, the not-saved
    /// sentence, or the error. A refresh that fails keeps what is drawn (no skeleton, no error).
    func after(_ result: Result<CachedFile<Value>, APIError>) -> StaticPhase {
        switch result {
        case let .success(file):
            return .loaded(file.value, offlineSince: file.isFromCache ? file.savedAt : nil)
        case let .failure(error):
            if value != nil { return self }
            if case .transport = error { return .notCached }
            return .failed
        }
    }
}

extension Result where Failure == APIError {
    /// Runs a typed-throwing read into a Result.
    static func of(_ operation: @Sendable () async throws(APIError) -> Success) async -> Result {
        do {
            return try await .success(operation())
        } catch {
            return .failure(error)
        }
    }
}

/// The Guides tab's index and the saved guide files (spec F9.1, F9.6).
struct GuidesHome: Equatable, Sendable {
    let build: String
    let index: GuideIndex
    let saved: SavedGuides
}

@MainActor
@Observable
final class GuidesHomeStore {
    private(set) var phase: StaticPhase<GuidesHome> = .loading
    private let source: any GuidesSource

    init(source: any GuidesSource) {
        self.source = source
    }

    func load() async {
        let build = await source.activeBuild()
        if case .loading = phase, let saved = source.cachedIndex(build: build) {
            phase = .loaded(home(build, saved.value), offlineSince: nil)
        }
        let source = source
        let result = await Result.of { () async throws(APIError) in try await source.index(build: build) }
        phase = phase.after(result.map { CachedFile(value: home(build, $0.value), savedAt: $0.savedAt, isFromCache: $0.isFromCache) })
    }

    /// After a guide opens and saves: the Saved tags and the footer catch up.
    func refreshSaved() {
        guard case let .loaded(home, since) = phase else { return }
        phase = .loaded(self.home(home.build, home.index), offlineSince: since)
    }

    private func home(_ build: String, _ index: GuideIndex) -> GuidesHome {
        let specs = index.classes.flatMap { $0.specs.map(\.spec) }
        return GuidesHome(build: build, index: index, saved: source.savedBytes(build: build, specs: specs))
    }
}

/// One guide as read: the file and the build it came from.
struct GuideContent: Equatable, Sendable {
    let guide: GuideFile
    let build: String
}

/// F9.2: one guide, saved copy first; a newer build waits behind the new-data strip.
@MainActor
@Observable
final class GuideStore {
    private(set) var phase: StaticPhase<GuideContent> = .loading
    private(set) var newData: NewDataState = .none
    let spec: String
    private let source: any GuidesSource

    init(spec: String, source: any GuidesSource) {
        self.spec = spec
        self.source = source
    }

    func load() async {
        let build = await source.activeBuild()
        if case .loading = phase, let saved = source.cachedGuide(build: build, spec: spec) {
            phase = .loaded(GuideContent(guide: saved.value, build: build), offlineSince: nil)
        }
        phase = phase.after(await read(build))
    }

    /// On foreground: a newer active build shows the strip and changes nothing else.
    func checkForNewData() async {
        guard case let .loaded(content, _) = phase else { return }
        newData = NewDataCheck.state(shown: content.build, latest: await source.latestBuild())
    }

    /// `Update`: the guide of the new build replaces the open one, or the strip says it failed.
    func update() async {
        let build: String
        switch newData {
        case .none: return
        case let .ready(next), let .failed(next): build = next
        }
        switch await read(build) {
        case let .success(file) where !file.isFromCache:
            phase = .loaded(file.value, offlineSince: nil)
            newData = .none
        default:
            newData = .failed(build)
        }
    }

    private func read(_ build: String) async -> Result<CachedFile<GuideContent>, APIError> {
        let (spec, source) = (spec, source)
        let result = await Result.of { () async throws(APIError) in try await source.guide(build: build, spec: spec) }
        return result.map { CachedFile(value: GuideContent(guide: $0.value, build: build), savedAt: $0.savedAt, isFromCache: $0.isFromCache) }
    }
}

/// F9.4: the tier lists and the chosen role and faction.
@MainActor
@Observable
final class TiersStore {
    private(set) var phase: StaticPhase<TiersFile> = .loading
    var choice: TierChoice
    private let source: any GuidesSource

    init(source: any GuidesSource, choice: TierChoice) {
        self.source = source
        self.choice = choice
    }

    func load() async {
        let build = await source.activeBuild()
        if case .loading = phase, let saved = StaticPhase.saved(source.cachedTiers(build: build)) { phase = saved }
        let source = source
        phase = phase.after(await Result.of { () async throws(APIError) in try await source.tiers(build: build) })
    }
}

/// F9.5: the key dates (always answered: CDN, saved or bundled) and the changelog.
@MainActor
@Observable
final class ChangesStore {
    private(set) var dates: [KeyDateRow]?
    private(set) var changelog: StaticPhase<Changelog> = .loading
    private let source: any GuidesSource
    private let now: @Sendable () -> Date

    init(source: any GuidesSource, now: @escaping @Sendable () -> Date) {
        self.source = source
        self.now = now
    }

    func load() async {
        if case .loading = changelog, let saved = StaticPhase.saved(source.cachedChangelog()) { changelog = saved }
        let source = source
        async let dates = source.keyDates()
        let result = await Result.of { () async throws(APIError) in try await source.changelog() }
        self.dates = KeyDatesLogic.rows(await dates, now: now())
        changelog = changelog.after(result)
    }
}

/// Keeps the guides of the player's classes saved (spec F9 Sync): on launch and on foreground,
/// at most hourly.
@MainActor
final class GuidesPrefetcher {
    private let source: any GuidesSource
    private let now: @Sendable () -> Date
    private var last: Date?

    init(source: any GuidesSource, now: @escaping @Sendable () -> Date) {
        self.source = source
        self.now = now
    }

    func run(classes: Set<String>) async {
        guard !classes.isEmpty, GuidesSync.isDue(last: last, now: now()) else { return }
        last = now()
        let build = await source.activeBuild()
        guard let index = try? await source.index(build: build).value else { return }
        for spec in GuidesSync.specs(in: index, classes: classes) {
            _ = try? await source.guide(build: build, spec: spec)
        }
    }
}
