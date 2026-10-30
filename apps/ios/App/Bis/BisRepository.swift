import Foundation
import Models
import Networking

/// Everything the BiS tab reads: the static files of one build (spec F3, data), each with its
/// offline copy, and the signed-in character. A protocol so the model is tested without a network.
protocol BisDataSource: Sendable {
    /// The build the lists are read from: `GET /v1/app/config`'s `active_build`, else the copy
    /// bundled with the app when the config has never been read.
    func activeBuild() async -> String
    func bis(build: String, spec: String) async throws(APIError) -> CachedFile<BisFile>
    func items(build: String, classSlug: String) async throws(APIError) -> CachedFile<ItemsFile>
    func classes(build: String) async throws(APIError) -> CachedFile<[ClassInfo]>
    /// The saved copies alone, for the first draw before any request.
    func cachedBis(build: String, spec: String) -> CachedFile<BisFile>?
    func cachedItems(build: String, classSlug: String) -> CachedFile<ItemsFile>?
    func cachedClasses(build: String) -> CachedFile<[ClassInfo]>?
    /// The current character, or nil signed out. F1's switcher replaces the main character.
    func currentCharacter() async -> AccountCharacter?
}

/// The live source: the CDN through the offline cache, the build from the app config.
struct LiveBisDataSource: BisDataSource {
    let files: CachedStaticDataClient
    let appConfig: AppConfigStore
    let client: APIClient
    /// Nil only for a broken build; the first offline launch then shows the not-saved sentence.
    let bundledBuild: String?

    func activeBuild() async -> String {
        if let build = appConfig.current?.activeBuild { return build }
        await appConfig.refresh()
        return appConfig.current?.activeBuild ?? bundledBuild ?? ""
    }

    func bis(build: String, spec: String) async throws(APIError) -> CachedFile<BisFile> {
        try await files.load(.getBis(build: build, spec: spec))
    }

    func items(build: String, classSlug: String) async throws(APIError) -> CachedFile<ItemsFile> {
        try await files.load(.getItems(build: build, className: classSlug))
    }

    func classes(build: String) async throws(APIError) -> CachedFile<[ClassInfo]> {
        try await files.load(.getClasses(build: build))
    }

    func cachedBis(build: String, spec: String) -> CachedFile<BisFile>? {
        files.cached(.getBis(build: build, spec: spec))
    }

    func cachedItems(build: String, classSlug: String) -> CachedFile<ItemsFile>? {
        files.cached(.getItems(build: build, className: classSlug))
    }

    func cachedClasses(build: String) -> CachedFile<[ClassInfo]>? {
        files.cached(.getClasses(build: build))
    }

    func currentCharacter() async -> AccountCharacter? {
        try? await client.send(.getMe).mainCharacter
    }
}

/// The active build bundled with the app (web/src/data/active-build.json, copied in by
/// project.yml): the fallback for a first launch with no connection.
enum BundledBuild {
    static let resourceName = "active-build"

    static func read(_ bundle: Bundle = .main) -> String? {
        guard let url = bundle.url(forResource: resourceName, withExtension: "json"),
              let data = try? Data(contentsOf: url)
        else { return nil }
        return try? WireFormat.decoder().decode(ActiveBuild.self, from: data).build
    }
}
