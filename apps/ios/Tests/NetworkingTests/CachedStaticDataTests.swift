import Foundation
import Models
@testable import Networking
import Testing
import TestSupport

/// Answers from a closure, so a test can go offline between two requests.
private final class SwitchableTransport: HTTPTransport, @unchecked Sendable {
    private let lock = NSLock()
    private var answer: (URLRequest) throws(APIError) -> (Data, HTTPURLResponse)

    init(_ answer: @escaping (URLRequest) throws(APIError) -> (Data, HTTPURLResponse)) {
        self.answer = answer
    }

    func set(_ answer: @escaping (URLRequest) throws(APIError) -> (Data, HTTPURLResponse)) {
        lock.withLock { self.answer = answer }
    }

    func send(_ request: URLRequest) async throws(APIError) -> (Data, HTTPURLResponse) {
        let current = lock.withLock { answer }
        return try current(request)
    }
}

private func ok(_ body: Data, status: Int = 200) -> (URLRequest) throws(APIError) -> (Data, HTTPURLResponse) {
    { request throws(APIError) in
        guard let url = request.url,
              let response = HTTPURLResponse(url: url, statusCode: status, httpVersion: nil, headerFields: nil)
        else { throw .transport("no response") }
        return (body, response)
    }
}

private func offline(_: URLRequest) throws(APIError) -> (Data, HTTPURLResponse) {
    throw .transport("The Internet connection appears to be offline.")
}

@Suite("Cached static data")
struct CachedStaticDataTests {
    private let resource = StaticResource<BisFile>.getBis(build: "1.60.1.70658", spec: "warrior-fury")

    private func client(_ transport: any HTTPTransport, cache: any StaticFileCache, now: Date) -> CachedStaticDataClient {
        CachedStaticDataClient(client: StaticDataClient(transport: transport), cache: cache, clock: { now })
    }

    @Test("a fetched file is decoded, saved, and marked fresh")
    func freshFetch() async throws {
        let body = try FixtureLocation.data("cdn.bis.warrior-fury.json")
        let cache = InMemoryStaticFileCache()
        let now = Date(timeIntervalSince1970: 1_000)
        let answer = try await client(SwitchableTransport(ok(body)), cache: cache, now: now).load(resource)
        #expect(answer.value.spec == "warrior-fury")
        #expect(answer.isFromCache == false)
        #expect(answer.savedAt == now)
        #expect(cache.read(resource.path)?.data == body)
    }

    @Test("offline answers the saved copy with the time it was saved")
    func offlineServesTheCache() async throws {
        let body = try FixtureLocation.data("cdn.bis.warrior-fury.json")
        let cache = InMemoryStaticFileCache()
        let savedAt = Date(timeIntervalSince1970: 1_000)
        let transport = SwitchableTransport(ok(body))
        _ = try await client(transport, cache: cache, now: savedAt).load(resource)
        transport.set(offline)
        let answer = try await client(transport, cache: cache, now: Date(timeIntervalSince1970: 9_000)).load(resource)
        #expect(answer.isFromCache)
        #expect(answer.savedAt == savedAt)
        #expect(answer.value.bands.isEmpty == false)
    }

    @Test("offline with nothing saved is still the transport error")
    func offlineWithoutCache() async {
        let transport = SwitchableTransport(offline)
        await #expect(throws: APIError.transport("The Internet connection appears to be offline.")) {
            _ = try await client(transport, cache: InMemoryStaticFileCache(), now: .now).load(resource)
        }
    }

    @Test("a server error is an error, never the cached copy")
    func serverErrorIsNotOffline() async throws {
        let body = try FixtureLocation.data("cdn.bis.warrior-fury.json")
        let cache = InMemoryStaticFileCache()
        let transport = SwitchableTransport(ok(body))
        _ = try await client(transport, cache: cache, now: .now).load(resource)
        transport.set(ok(Data(), status: 404))
        await #expect(throws: APIError.unexpectedStatus(404)) {
            _ = try await client(transport, cache: cache, now: .now).load(resource)
        }
    }

    @Test("the saved copy reads without a request, for the first draw")
    func cachedOnly() async throws {
        let body = try FixtureLocation.data("cdn.bis.warrior-fury.json")
        let cache = InMemoryStaticFileCache()
        let reader = client(SwitchableTransport(ok(body)), cache: cache, now: Date(timeIntervalSince1970: 5))
        #expect(reader.cached(resource) == nil)
        _ = try await reader.load(resource)
        let saved = try #require(reader.cached(resource))
        #expect(saved.isFromCache)
        #expect(saved.savedAt == Date(timeIntervalSince1970: 5))
    }

    @Test("the disk cache keeps a file per path with its save time")
    func diskCache() throws {
        let directory = FileManager.default.temporaryDirectory.appending(path: UUID().uuidString)
        defer { try? FileManager.default.removeItem(at: directory) }
        let cache = DiskStaticFileCache(directory: directory)
        let savedAt = Date(timeIntervalSince1970: 1_700_000_000)
        cache.write(Data("{}".utf8), for: "/data/1.60.1.70658/bis/warrior-fury.json", at: savedAt)
        let entry = try #require(cache.read("/data/1.60.1.70658/bis/warrior-fury.json"))
        #expect(entry.data == Data("{}".utf8))
        #expect(abs(entry.savedAt.timeIntervalSince(savedAt)) < 1)
        #expect(cache.read("/data/1.60.1.70658/bis/warrior-arms.json") == nil)
    }
}
