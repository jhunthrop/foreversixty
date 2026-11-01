@testable import ForeverSixty
import Foundation
import Models
import Networking
import Testing

@Suite("The strip decision: offline, stale or the region's error")
struct CacheFreshnessTests {
    private let now = SystemStory.now
    private var old: Date { now.addingTimeInterval(-11 * 60) }
    private var young: Date { now.addingTimeInterval(-60) }
    private let timeout = APIError.transport("The request timed out.")
    private let serverDown = APIError.unexpectedStatus(503)

    @Test("no connection with a cache is offline, whatever its age")
    func offline() {
        #expect(CacheFreshness.decide(timeout, connected: false, cachedAt: young, limit: .homeAndRaids, now: now) == .offline(since: young))
        #expect(CacheFreshness.decide(timeout, connected: false, cachedAt: old, limit: .pages, now: now) == .offline(since: old))
    }

    @Test("connected, no answer, a cache past its limit is stale (transport, timeout or 5xx)")
    func stale() {
        #expect(CacheFreshness.decide(timeout, connected: true, cachedAt: old, limit: .homeAndRaids, now: now) == .stale(since: old))
        #expect(CacheFreshness.decide(serverDown, connected: true, cachedAt: old, limit: .pages, now: now) == .stale(since: old))
        let body = ErrorBody(code: "internal", message: "boom")
        #expect(CacheFreshness.decide(.server(status: 500, body: body, requestID: nil), connected: true, cachedAt: old, limit: .pages, now: now) == .stale(since: old))
    }

    @Test("the limits: 5 minutes for Home and raids, 10 for pages, never for build data")
    func limits() {
        let six = now.addingTimeInterval(-6 * 60)
        #expect(CacheFreshness.decide(timeout, connected: true, cachedAt: six, limit: .homeAndRaids, now: now) == .stale(since: six))
        #expect(CacheFreshness.decide(timeout, connected: true, cachedAt: six, limit: .pages, now: now) == .offline(since: six))
        let month = now.addingTimeInterval(-30 * 86400)
        #expect(CacheFreshness.decide(timeout, connected: true, cachedAt: month, limit: .build, now: now) == .offline(since: month))
    }

    @Test("a 4xx is always the region's error; no cache is the region's error")
    func errors() {
        let body = ErrorBody(code: "not_found", message: "gone")
        #expect(CacheFreshness.decide(.server(status: 404, body: body, requestID: nil), connected: true, cachedAt: old, limit: .pages, now: now) == .error)
        #expect(CacheFreshness.decide(.unexpectedStatus(400), connected: true, cachedAt: old, limit: .pages, now: now) == .error)
        #expect(CacheFreshness.decide(timeout, connected: true, cachedAt: nil, limit: .pages, now: now) == .error)
        #expect(CacheFreshness.decide(serverDown, connected: true, cachedAt: young, limit: .pages, now: now) == .error)
    }

    @Test("unknown connectivity keeps v1's rule: transport with a cache is offline")
    func unknown() {
        #expect(CacheFreshness.decide(timeout, connected: nil, cachedAt: old, limit: .homeAndRaids, now: now) == .offline(since: old))
        #expect(CacheFreshness.decide(serverDown, connected: nil, cachedAt: old, limit: .homeAndRaids, now: now) == .error)
    }
}

@MainActor
@Suite("Home's stale strip")
struct HomeStaleTests {
    @Test("connected, the API down, an answer older than five minutes: the cache under the stale strip")
    func stale() async throws {
        let home = try HomeFixtures.home()
        let saved = HomeFixtures.now.addingTimeInterval(-6 * 60)
        let cache = MemoryHomeCache([HomeFixtures.velaren.key: HomeCacheEntry(home: home, fetchedAt: saved)])
        let store = HomeStore(
            fetch: { _ throws(APIError) in throw .unexpectedStatus(502) }, cache: cache,
            now: { HomeFixtures.now }, connectivity: FixedConnectivity(isConnected: true)
        )
        await store.load(HomeFixtures.velaren)
        #expect(store.phase == .loaded(home, offlineSince: nil))
        #expect(store.staleSince == saved)
    }
}
