@testable import ForeverSixty
import Foundation
import Models
import Networking
import Testing

/// Counts requests and can be switched offline, over the contract fixtures.
final class CountingTransport: HTTPTransport, @unchecked Sendable {
    private let lock = NSLock()
    private var paths: [String] = []
    private var offline = false
    private let fixtures = FixtureTransport(directory: HomeFixtures.directory)

    var requested: [String] { lock.withLock { paths } }
    func goOffline() { lock.withLock { offline = true } }

    func send(_ request: URLRequest) async throws(APIError) -> (Data, HTTPURLResponse) {
        let isOffline = lock.withLock {
            paths.append(request.url?.path ?? "")
            return offline
        }
        if isOffline { throw .transport("offline") }
        return try await fixtures.send(request)
    }
}

@Suite("Static data store")
struct StaticDataStoreTests {
    private let directory = FileManager.default.temporaryDirectory.appending(path: "static-data-\(UUID().uuidString)")
    private let build = SearchTestData.build

    private func store(_ transport: CountingTransport) -> StaticDataStore {
        StaticDataStore(client: StaticDataClient(transport: transport), directory: directory)
    }

    @Test("a file is fetched once, then read from memory; a new store reads it from disk offline")
    func cacheFirst() async throws {
        let transport = CountingTransport()
        let first = store(transport)
        let index = try await first.file(.getSearchIndex(build: build))
        _ = try await first.file(.getSearchIndex(build: build))
        #expect(transport.requested.count == 1)
        #expect(!index.items.isEmpty)

        transport.goOffline()
        let reopened = store(transport)
        let offline = try await reopened.file(.getSearchIndex(build: build))
        #expect(offline == index)
        #expect(transport.requested.count == 1)
        try? FileManager.default.removeItem(at: directory)
    }

    @Test("never fetched and offline: the transport error, not a crash")
    func offlineMiss() async {
        let transport = CountingTransport()
        transport.goOffline()
        await #expect(throws: APIError.self) { try await store(transport).file(.getSearchIndex(build: build)) }
    }

    @Test("swap prunes every build but the active one; size counts the bytes; delete empties it")
    func pruneSizeDelete() async throws {
        let transport = CountingTransport()
        let store = store(transport)
        _ = try await store.file(.getSearchIndex(build: build))
        let old = directory.appending(path: "1.60.1.70532")
        try FileManager.default.createDirectory(at: old, withIntermediateDirectories: true)
        try Data(repeating: 1, count: 100).write(to: old.appending(path: "search.json"))
        let before = await store.sizeOnDisk()
        #expect(before > 100)

        await store.prune(keeping: build)
        #expect(!FileManager.default.fileExists(atPath: old.path))
        #expect(FileManager.default.fileExists(atPath: directory.appending(path: build).path))
        #expect(await store.sizeOnDisk() == before - 100)

        await store.removeAll()
        #expect(await store.sizeOnDisk() == 0)
        transport.goOffline()
        #expect(await store.cached(.getSearchIndex(build: build)) == nil)
    }

    @Test("the item catalog reads through the store: one fetch per file")
    func catalogShares() async throws {
        let transport = CountingTransport()
        let catalog = ItemCatalog(store: StaticDataStore(client: StaticDataClient(transport: transport), directory: nil))
        let request = ItemTooltipRequest(
            itemID: 16866, fallback: HomeItemLook(id: 16866, name: nil, quality: nil, icon: nil, itemLevel: nil),
            build: build, classSlug: "warrior", spec: "warrior-fury", specName: nil, band: 60, bandLabel: nil,
            preset: "raid", faction: "horde", sourceLine: nil, wornID: nil
        )
        _ = try await catalog.sources(for: request)
        let count = transport.requested.count
        _ = try await catalog.sources(for: request)
        #expect(transport.requested.count == count)
    }
}
