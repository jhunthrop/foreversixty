import Foundation
import Models
import Networking

/// The site's static files on this phone, one folder per game build (spec Wave 3 §14: lane F12
/// owns it; F9 and F15 reuse it). A file never changes under one build's path, so a file is read
/// cache first: memory, then disk, then the CDN, and a fetched body is kept in both. Swapping to a
/// new build prunes every other build; sign-out deletes everything; `sizeOnDisk` is what the
/// Offline data screen (F15.4) reports. Never sends the device token.
actor StaticDataStore {
    private let client: StaticDataClient
    private let directory: URL?
    private var memory: [String: Data] = [:]

    /// `directory` nil keeps the files in memory only (tests and previews).
    init(client: StaticDataClient, directory: URL? = StaticDataStore.defaultDirectory) {
        self.client = client
        self.directory = directory
    }

    /// `Caches/static-data`: the system may purge it, and every file can be fetched again.
    static var defaultDirectory: URL? {
        FileManager.default.urls(for: .cachesDirectory, in: .userDomainMask).first?.appending(path: "static-data")
    }

    /// The file, from memory or disk when this phone has it, else from the CDN (then kept).
    func file<T>(_ resource: StaticResource<T>) async throws(APIError) -> T {
        if let value = cached(resource) { return value }
        let (value, body) = try await client.fetch(resource)
        memory[resource.path] = body
        writeDisk(resource.path, body)
        return value
    }

    /// The saved copy alone, with no request; nil when it was never fetched or no longer decodes.
    func cached<T>(_ resource: StaticResource<T>) -> T? {
        guard let data = memory[resource.path] ?? readDisk(resource.path),
              let value = try? StaticDataClient.decode(resource, from: data)
        else { return nil }
        memory[resource.path] = data
        return value
    }

    /// Forgets every build but `build`, once that build is the one in use.
    func prune(keeping build: String) {
        memory = memory.filter { path, _ in Self.build(of: path) == build }
        guard let directory, let names = try? FileManager.default.contentsOfDirectory(atPath: directory.path) else { return }
        for name in names where name != build {
            try? FileManager.default.removeItem(at: directory.appending(path: name))
        }
    }

    /// Deletes every saved file (sign-out, and the Offline data screen's Delete).
    func removeAll() {
        memory = [:]
        guard let directory else { return }
        try? FileManager.default.removeItem(at: directory)
    }

    /// The bytes the saved files take on disk.
    func sizeOnDisk() -> Int {
        guard let directory,
              let files = FileManager.default.enumerator(at: directory, includingPropertiesForKeys: [.fileSizeKey])
        else { return 0 }
        return files.compactMap { ($0 as? URL).flatMap { try? $0.resourceValues(forKeys: [.fileSizeKey]).fileSize } }
            .reduce(0, +)
    }

    /// `/data/1.60.1.70658/items/shaman.json` → `1.60.1.70658`.
    private static func build(of path: String) -> String {
        String(relative(path).prefix { $0 != "/" })
    }

    private static func relative(_ path: String) -> String {
        path.split(separator: "/").dropFirst().filter { $0 != ".." }.joined(separator: "/")
    }

    private func diskURL(_ path: String) -> URL? {
        directory?.appending(path: Self.relative(path))
    }

    private func readDisk(_ path: String) -> Data? {
        diskURL(path).flatMap { try? Data(contentsOf: $0) }
    }

    /// Best effort: a full disk loses only the offline copy, never the answer on screen.
    private func writeDisk(_ path: String, _ data: Data) {
        guard let url = diskURL(path) else { return }
        try? FileManager.default.createDirectory(at: url.deletingLastPathComponent(), withIntermediateDirectories: true)
        try? data.write(to: url, options: .atomic)
    }
}
