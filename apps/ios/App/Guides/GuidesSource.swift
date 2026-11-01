import Auth
import Foundation
import Models
import Networking

/// Everything the Guides tab reads (spec F9 Data): the W3-C1 static files through the offline
/// cache, the active build and the key dates. A protocol so the stores are tested without a network.
protocol GuidesSource: Sendable {
    /// The build the guides are read from: `/data/active-build.json` (or its saved copy), else the
    /// app config's, else the copy bundled with the app.
    func activeBuild() async -> String
    /// A fresh read of the active build for the new-data check; nil when it cannot be read now.
    func latestBuild() async -> String?
    func index(build: String) async throws(APIError) -> CachedFile<GuideIndex>
    func cachedIndex(build: String) -> CachedFile<GuideIndex>?
    func guide(build: String, spec: String) async throws(APIError) -> CachedFile<GuideFile>
    func cachedGuide(build: String, spec: String) -> CachedFile<GuideFile>?
    /// Which of `specs` have their guide file saved, and its size on disk.
    func savedBytes(build: String, specs: [String]) -> SavedGuides
    func tiers(build: String) async throws(APIError) -> CachedFile<TiersFile>
    func cachedTiers(build: String) -> CachedFile<TiersFile>?
    func changelog() async throws(APIError) -> CachedFile<Changelog>
    func cachedChangelog() -> CachedFile<Changelog>?
    func keyDates() async -> [KeyDate]
}

/// The live source: the CDN through `CachedStaticDataClient`, as BiS and Talents read it.
struct LiveGuidesSource: GuidesSource {
    let files: CachedStaticDataClient
    let cache: any StaticFileCache
    let bisSource: LiveBisDataSource

    func activeBuild() async -> String {
        if let build = try? await files.load(.getActiveBuild).value.build { return build }
        return await bisSource.activeBuild()
    }

    func latestBuild() async -> String? {
        guard let answer = try? await files.load(.getActiveBuild), !answer.isFromCache else { return nil }
        return answer.value.build
    }

    func index(build: String) async throws(APIError) -> CachedFile<GuideIndex> {
        try await files.load(.getGuideIndex(build: build))
    }

    func cachedIndex(build: String) -> CachedFile<GuideIndex>? { files.cached(.getGuideIndex(build: build)) }

    func guide(build: String, spec: String) async throws(APIError) -> CachedFile<GuideFile> {
        try await files.load(.getGuide(build: build, spec: spec))
    }

    func cachedGuide(build: String, spec: String) -> CachedFile<GuideFile>? {
        files.cached(.getGuide(build: build, spec: spec))
    }

    func savedBytes(build: String, specs: [String]) -> SavedGuides {
        let pairs = specs.compactMap { spec in
            cache.read(StaticResource<GuideFile>.getGuide(build: build, spec: spec).path).map { (spec, $0.data.count) }
        }
        return SavedGuides(bytes: Dictionary(pairs) { first, _ in first })
    }

    func tiers(build: String) async throws(APIError) -> CachedFile<TiersFile> { try await files.load(.getTiers(build: build)) }

    func cachedTiers(build: String) -> CachedFile<TiersFile>? { files.cached(.getTiers(build: build)) }

    func changelog() async throws(APIError) -> CachedFile<Changelog> { try await files.load(.getChangelog) }

    func cachedChangelog() -> CachedFile<Changelog>? { files.cached(.getChangelog) }

    func keyDates() async -> [KeyDate] { await bisSource.keyDates() }
}

/// What the Guides tab needs from the app.
struct GuidesDependencies {
    let source: any GuidesSource
    let specs: SpecCatalog
    let now: @Sendable () -> Date

    @MainActor
    static func live(_ environment: AppEnvironment, session: SessionStore?) -> GuidesDependencies {
        let transport: any HTTPTransport = switch environment.configuration.mode {
        case .live: URLSessionTransport()
        case .fixture: FixtureTransport.bundled() ?? FixtureTransport(directory: Bundle.main.bundleURL)
        }
        let cache = DiskStaticFileCache.applicationSupport()
        let files = CachedStaticDataClient(client: StaticDataClient(transport: transport), cache: cache)
        let bis = LiveBisDataSource(
            files: files, appConfig: environment.appConfig,
            character: { session?.currentCharacter }, bundledBuild: BundledBuild.read()
        )
        return GuidesDependencies(
            source: LiveGuidesSource(files: files, cache: cache, bisSource: bis),
            specs: .bundled(),
            now: { Date() }
        )
    }
}
