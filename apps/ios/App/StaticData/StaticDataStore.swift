import Foundation
import Models
import Networking

/// The site's static files on this phone: the app's ONE static data store (lane F15 merged F12's
/// store with the F9, BiS and Talents `CachedStaticDataClient`). One folder per game build in
/// one directory (`Application Support/static-data`, excluded from backup), one build swap
/// (`prune`), one delete (`removeAll`) and one size report per category (`report`, the Offline
/// data screen). A file never changes under one build's path, so `file` reads cache first:
/// memory, then disk, then the CDN. `files` is the same disk through `CachedStaticDataClient`,
/// for the screens that draw the saved copy first and refresh it. Never sends the device token.
actor StaticDataStore {
    private let client: StaticDataClient
    private let directory: URL?
    private let disk: (any StaticFileCache)?
    private var memory: [String: Data] = [:]
    /// The disk the store keeps (memory only for tests), shared with `files`.
    nonisolated let fileCache: any StaticFileCache
    /// The CDN through the same disk cache: BiS, Guides and Talents read through it.
    nonisolated let files: CachedStaticDataClient

    /// `directory` nil keeps the files in memory only (tests and previews).
    init(client: StaticDataClient, directory: URL? = StaticDataStore.defaultDirectory) {
        self.client = client
        self.directory = directory
        let disk: (any StaticFileCache)? = directory.map(DiskStaticFileCache.init(directory:))
        self.disk = disk
        let cache = disk ?? InMemoryStaticFileCache()
        fileCache = cache
        files = CachedStaticDataClient(client: client, cache: cache)
    }

    /// The active build as the CDN says it now; nil when it cannot be read now (offline). The
    /// new-data strip and Update now compare against it.
    nonisolated func latestBuild() async -> String? {
        guard let answer = try? await files.load(.getActiveBuild), !answer.isFromCache else { return nil }
        return answer.value.build
    }

    /// `Application Support/static-data`, excluded from backup: the one directory.
    static var defaultDirectory: URL? {
        DiskStaticFileCache.applicationSupport().directory
    }

    /// The file, from memory or disk when this phone has it, else from the CDN (then kept).
    func file<T>(_ resource: StaticResource<T>) async throws(APIError) -> T {
        if let value = cached(resource) { return value }
        let (value, body) = try await client.fetch(resource)
        memory[resource.path] = body
        disk?.write(body, for: resource.path, at: Date())
        return value
    }

    /// The saved copy alone, with no request; nil when it was never fetched or no longer decodes.
    func cached<T>(_ resource: StaticResource<T>) -> T? {
        guard let data = memory[resource.path] ?? disk?.read(resource.path)?.data,
              let value = try? StaticDataClient.decode(resource, from: data)
        else { return nil }
        memory[resource.path] = data
        return value
    }

    /// Forgets every build but `build`, once that build is the one in use.
    func prune(keeping build: String) {
        memory = memory.filter { path, _ in StaticFileLayout.build(of: path) == build }
        guard let directory, let names = try? FileManager.default.contentsOfDirectory(atPath: directory.path) else { return }
        for name in names where StaticFileLayout.isBuildFolder(name) && name != build {
            try? FileManager.default.removeItem(at: directory.appending(path: name))
        }
    }

    /// Deletes every saved file (sign-out, and the Offline data screen's Remove).
    func removeAll() {
        memory = [:]
        guard let directory else { return }
        try? FileManager.default.removeItem(at: directory)
    }

    /// The bytes the saved files take on disk.
    func sizeOnDisk() -> Int {
        savedFiles().reduce(0) { $0 + $1.bytes }
    }

    /// Every saved file of `build`: its path under the build, size and save time.
    func savedFiles(build: String? = nil) -> [SavedStaticFile] {
        guard let directory,
              let walker = FileManager.default.enumerator(
                  at: directory, includingPropertiesForKeys: [.fileSizeKey, .contentModificationDateKey, .isRegularFileKey]
              )
        else { return [] }
        let base = directory.standardizedFileURL.path + "/"
        return walker.compactMap { item -> SavedStaticFile? in
            guard let url = item as? URL,
                  let values = try? url.resourceValues(forKeys: [.fileSizeKey, .contentModificationDateKey, .isRegularFileKey]),
                  values.isRegularFile == true
            else { return nil }
            let relative = url.standardizedFileURL.path.replacingOccurrences(of: base, with: "")
            return SavedStaticFile(relativePath: relative, bytes: values.fileSize ?? 0, savedAt: values.contentModificationDate ?? .distantPast)
        }
        .filter { build == nil || StaticFileLayout.build(ofRelative: $0.relativePath) == build }
    }
}

/// One saved static file, for the Offline data report.
struct SavedStaticFile: Equatable, Sendable {
    /// Under the store's directory: `1.60.1.70658/guides/shaman-elemental.json`.
    let relativePath: String
    let bytes: Int
    let savedAt: Date
}

/// Where a static path lives on disk: `/data/1.60.1.70658/items/shaman.json` is
/// `1.60.1.70658/items/shaman.json` under the store's directory.
enum StaticFileLayout {
    static func build(of path: String) -> String {
        build(ofRelative: path.split(separator: "/").dropFirst().joined(separator: "/"))
    }

    static func build(ofRelative relative: String) -> String {
        String(relative.prefix { $0 != "/" })
    }

    /// A build folder is named like `1.60.1.70658`; root files (`active-build.json`) are not.
    static func isBuildFolder(_ name: String) -> Bool {
        name.first?.isNumber == true && !name.hasSuffix(".json")
    }
}
