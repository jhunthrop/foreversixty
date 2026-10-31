import Foundation
import Models
import Networking

/// The build's static files the tooltip reads (`items/{class}.json`, `sets.json`,
/// `bis/{spec}.json`), fetched once from the CDN and kept per build: in memory for the session,
/// on disk until a newer build is in use (spec §5.0 Caching). Never sends the device token.
actor ItemCatalog {
    private let client: StaticDataClient
    private let directory: URL?
    private var memory: [String: Data] = [:]

    /// `directory` nil keeps the files in memory only (tests and previews).
    init(client: StaticDataClient, directory: URL? = ItemCatalog.defaultDirectory) {
        self.client = client
        self.directory = directory
    }

    static var defaultDirectory: URL? {
        FileManager.default.urls(for: .cachesDirectory, in: .userDomainMask).first?.appending(path: "static-data")
    }

    /// The files a tooltip needs. Items are required; sets and the BiS file only add lines,
    /// so a missing one leaves its lines out rather than failing the sheet.
    func sources(for request: ItemTooltipRequest) async throws(APIError) -> ItemTooltipSources {
        let items = try await file(.getItems(build: request.build, className: request.classSlug), as: ItemsFile.self)
        let sets = try? await file(.getSets(build: request.build), as: [ItemSet].self)
        let bis: BisFile? = if let spec = request.spec {
            try? await file(.getBis(build: request.build, spec: spec), as: BisFile.self)
        } else {
            nil
        }
        return ItemTooltipSources(
            items: Dictionary(items.items.map { ($0.id, $0) }, uniquingKeysWith: { first, _ in first }),
            sets: sets ?? [],
            bis: bis
        )
    }

    /// Forgets every build but `keep` on disk, once a newer build is in use.
    func prune(keeping build: String) {
        guard let directory, let builds = try? FileManager.default.contentsOfDirectory(atPath: directory.path) else { return }
        for name in builds where name != build {
            try? FileManager.default.removeItem(at: directory.appending(path: name))
        }
    }

    private func file<T: Codable & Sendable>(_ resource: StaticResource<T>, as _: T.Type) async throws(APIError) -> T {
        if let data = memory[resource.path] ?? readDisk(resource.path), let value = try? decode(T.self, data) {
            memory[resource.path] = data
            return value
        }
        let value = try await client.get(resource)
        if let data = try? WireFormat.encoder().encode(value) {
            memory[resource.path] = data
            writeDisk(resource.path, data)
        }
        return value
    }

    private func decode<T: Decodable>(_: T.Type, _ data: Data) throws -> T {
        try WireFormat.decoder().decode(T.self, from: data)
    }

    private func diskURL(_ path: String) -> URL? {
        directory?.appending(path: String(path.dropFirst("/data/".count)))
    }

    private func readDisk(_ path: String) -> Data? {
        diskURL(path).flatMap { try? Data(contentsOf: $0) }
    }

    private func writeDisk(_ path: String, _ data: Data) {
        guard let url = diskURL(path) else { return }
        try? FileManager.default.createDirectory(at: url.deletingLastPathComponent(), withIntermediateDirectories: true)
        try? data.write(to: url, options: .atomic)
    }
}
