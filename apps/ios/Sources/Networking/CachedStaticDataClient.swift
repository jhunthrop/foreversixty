import Foundation
import Synchronization

/// One saved static file: its body and when it was fetched.
public struct StaticFileEntry: Sendable, Equatable {
    public let data: Data
    public let savedAt: Date

    public init(data: Data, savedAt: Date) {
        self.data = data
        self.savedAt = savedAt
    }
}

/// Where static files are kept for offline use, keyed by their `/data/...` path. A build's files
/// never change under one path, so a newer build simply uses new keys.
public protocol StaticFileCache: Sendable {
    func read(_ path: String) -> StaticFileEntry?
    func write(_ data: Data, for path: String, at date: Date)
}

/// The cache on disk: one file per path under `directory`, its modification date the save time.
/// The app passes a folder in Application Support, excluded from backup (spec §5.0 Caching).
public struct DiskStaticFileCache: StaticFileCache {
    public let directory: URL

    public init(directory: URL) {
        self.directory = directory
    }

    /// `Application Support/static-data`, excluded from iCloud and device backups.
    public static func applicationSupport() -> DiskStaticFileCache {
        let base = (try? FileManager.default.url(
            for: .applicationSupportDirectory, in: .userDomainMask, appropriateFor: nil, create: true
        )) ?? FileManager.default.temporaryDirectory
        var directory = base.appending(path: "static-data", directoryHint: .isDirectory)
        try? FileManager.default.createDirectory(at: directory, withIntermediateDirectories: true)
        var values = URLResourceValues()
        values.isExcludedFromBackup = true
        try? directory.setResourceValues(values)
        return DiskStaticFileCache(directory: directory)
    }

    public func read(_ path: String) -> StaticFileEntry? {
        let url = file(for: path)
        guard let data = try? Data(contentsOf: url),
              let date = try? url.resourceValues(forKeys: [.contentModificationDateKey]).contentModificationDate
        else { return nil }
        return StaticFileEntry(data: data, savedAt: date)
    }

    /// Best effort: a full disk loses only the offline copy, never the answer on screen.
    public func write(_ data: Data, for path: String, at date: Date) {
        let url = file(for: path)
        try? FileManager.default.createDirectory(
            at: url.deletingLastPathComponent(), withIntermediateDirectories: true
        )
        guard (try? data.write(to: url, options: .atomic)) != nil else { return }
        try? FileManager.default.setAttributes([.modificationDate: date], ofItemAtPath: url.path)
    }

    /// `/data/1.60.1.70658/bis/warrior-fury.json` is `1.60.1.70658/bis/warrior-fury.json` here: one
    /// folder per build, so swapping builds deletes whole folders (the app's StaticDataStore).
    private func file(for path: String) -> URL {
        let parts = path.split(separator: "/").filter { $0 != ".." && !$0.isEmpty }
        let relative = (parts.first == "data" ? parts.dropFirst() : parts[...]).joined(separator: "/")
        return directory.appending(path: relative)
    }
}

/// For tests and previews.
public final class InMemoryStaticFileCache: StaticFileCache {
    private let entries = Mutex<[String: StaticFileEntry]>([:])

    public init() {}

    public func read(_ path: String) -> StaticFileEntry? { entries.withLock { $0[path] } }

    public func write(_ data: Data, for path: String, at date: Date) {
        entries.withLock { $0[path] = StaticFileEntry(data: data, savedAt: date) }
    }
}

/// A static file as the app has it: fetched now, or the saved copy when offline.
public struct CachedFile<Value: Sendable>: Sendable {
    public let value: Value
    /// When the body was fetched.
    public let savedAt: Date
    /// True when the network was out and this is the saved copy.
    public let isFromCache: Bool

    public init(value: Value, savedAt: Date, isFromCache: Bool) {
        self.value = value
        self.savedAt = savedAt
        self.isFromCache = isFromCache
    }
}

extension CachedFile: Equatable where Value: Equatable {}

/// Static files with an offline copy (spec §5.0): every good fetch is saved; a fetch that fails
/// for lack of a connection answers the saved copy instead. Any other failure (a 404, a body
/// that does not decode) stays an error, because that is the error state, not offline.
public struct CachedStaticDataClient: Sendable {
    private let client: StaticDataClient
    private let cache: any StaticFileCache
    private let clock: @Sendable () -> Date

    public init(
        client: StaticDataClient,
        cache: any StaticFileCache,
        clock: @escaping @Sendable () -> Date = { Date() }
    ) {
        self.client = client
        self.cache = cache
        self.clock = clock
    }

    public func load<Response>(_ resource: StaticResource<Response>) async throws(APIError) -> CachedFile<Response> {
        do {
            let (value, body) = try await client.fetch(resource)
            let now = clock()
            cache.write(body, for: resource.path, at: now)
            return CachedFile(value: value, savedAt: now, isFromCache: false)
        } catch {
            guard case .transport = error, let saved = cached(resource) else { throw error }
            return saved
        }
    }

    /// The saved copy alone, with no request: what a screen draws at once before it refreshes.
    public func cached<Response>(_ resource: StaticResource<Response>) -> CachedFile<Response>? {
        guard let entry = cache.read(resource.path),
              let value = try? StaticDataClient.decode(resource, from: entry.data)
        else { return nil }
        return CachedFile(value: value, savedAt: entry.savedAt, isFromCache: true)
    }
}
