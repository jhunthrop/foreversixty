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
    /// The session's current character (lane F1's selector, spec §3.4), or nil signed out.
    func currentCharacter() async -> AccountCharacter?
    /// The key dates: the CDN's (or its saved copy), else the copy bundled with the app.
    func keyDates() async -> [KeyDate]
}

/// The live source: the CDN through the offline cache, the build from the app config.
struct LiveBisDataSource: BisDataSource {
    let files: CachedStaticDataClient
    let appConfig: AppConfigStore
    /// The session's current character; nil signed out or while sign-in is off.
    let character: @MainActor @Sendable () -> AccountCharacter?
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
        await character()
    }

    func keyDates() async -> [KeyDate] {
        if let dates = try? await files.load(.getKeyDates).value { return dates }
        return BundledData.keyDates()
    }
}

/// The active build bundled with the app (web/src/data/active-build.json, copied in by
/// project.yml): the fallback for a first launch with no connection.
enum BundledBuild {
    static let resourceName = "active-build"

    static func read(_ bundle: Bundle = .main) -> String? {
        BundledData.decode(resourceName, as: ActiveBuild.self, bundle: bundle)?.build
    }
}

/// The site's files bundled with the app for when the CDN cannot answer: the key dates
/// (web/src/data/dates.json; `/data/dates.json` is not published yet, app-v1.md's open item).
enum BundledData {
    static let datesResource = "dates"

    static func keyDates(_ bundle: Bundle = .main) -> [KeyDate] {
        decode(datesResource, as: [KeyDate].self, bundle: bundle) ?? []
    }

    static func decode<T: Decodable>(_ name: String, as _: T.Type, bundle: Bundle = .main) -> T? {
        guard let url = bundle.url(forResource: name, withExtension: "json"),
              let data = try? Data(contentsOf: url)
        else { return nil }
        return try? WireFormat.decoder().decode(T.self, from: data)
    }
}
