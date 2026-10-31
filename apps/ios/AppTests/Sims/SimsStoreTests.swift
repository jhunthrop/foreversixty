@testable import ForeverSixty
import Foundation
import Models
import Networking
import Testing

/// A scripted sims service: each call answers from its closure; reruns and seen marks are recorded.
final class FakeSimsService: SimsService, @unchecked Sendable {
    var listAnswer: (Int, String?) throws(APIError) -> SimList = { _, lane throws(APIError) in
        let list = try RaidFixture.api(SimList.self, "sims.list")
        guard lane == SimLane.server else { return list }
        return SimList(rows: list.rows.filter { $0.lane == SimLane.server }, total: list.total, page: 1, perPage: list.perPage)
    }
    var progressAnswer: (String) throws(APIError) -> SimProgress = { _ throws(APIError) in try RaidFixture.api(SimProgress.self, "sims.progress") }
    var rerunAnswer: (String) throws(APIError) -> SimStarted = { _ throws(APIError) in try RaidFixture.api(SimStarted.self, "sims.rerun") }
    var resultAnswer: (String) throws(APIError) -> SimResult = { _ throws(APIError) in try RaidFixture.api(SimResult.self, "sims.result") }
    private(set) var lists: [(page: Int, lane: String?)] = []
    private(set) var reruns: [String] = []
    private(set) var seen: [String] = []
    private(set) var polled: [String] = []

    func list(page: Int, lane: String?) async throws(APIError) -> SimList {
        lists.append((page, lane))
        return try listAnswer(page, lane)
    }

    func result(id: String) async throws(APIError) -> SimResult { try resultAnswer(id) }

    func progress(id: String) async throws(APIError) -> SimProgress {
        polled.append(id)
        return try progressAnswer(id)
    }

    func rerun(id: String) async throws(APIError) -> SimStarted {
        reruns.append(id)
        return try rerunAnswer(id)
    }

    func markSeen(id: String) async throws(APIError) { seen.append(id) }
}

@MainActor
@Suite("Sims store")
struct SimsStoreTests {
    private func store(_ service: FakeSimsService, cache: MemoryRaidCache = MemoryRaidCache()) -> SimsStore {
        SimsStore(service: service, cache: cache, now: { SimsLogicTests.now })
    }

    private func content(_ store: SimsStore) throws -> SimsContent {
        guard case let .loaded(content) = store.phase else {
            Issue.record("not loaded: \(store.phase)")
            throw APIError.missingFixture("not loaded")
        }
        return content
    }

    @Test func load_readsServerRunsAndYourSims() async throws {
        let service = FakeSimsService()
        let store = store(service)
        await store.load()
        let content = try content(store)
        #expect(content.serverRows.count == 3)
        #expect(content.mine.count == 3)
        #expect(!content.hasMorePages)
        #expect(Set(service.lists.map { $0.lane ?? "all" }) == ["server", "all"])
    }

    @Test func offline_showsTheCacheUnderTheStrip() async throws {
        let cache = MemoryRaidCache()
        await store(FakeSimsService(), cache: cache).load()
        let offline = FakeSimsService()
        offline.listAnswer = { _, _ throws(APIError) in throw .transport("offline") }
        let store = store(offline, cache: cache)
        await store.load()
        #expect(try content(store).offlineSince == SimsLogicTests.now)
    }

    @Test func serverError_withNothingCached_fails() async {
        let service = FakeSimsService()
        service.listAnswer = { _, _ throws(APIError) in throw .unexpectedStatus(500) }
        let store = store(service)
        await store.load()
        #expect(store.phase == .failed)
    }

    @Test func emptyList_loadsEmpty() async throws {
        let service = FakeSimsService()
        service.listAnswer = { _, _ throws(APIError) in try RaidFixture.api(SimList.self, "sims.list.empty") }
        let store = store(service)
        await store.load()
        #expect(try content(store).mine.isEmpty)
    }

    @Test func poll_asksTheRunningRowAndReloadsWhenItEnds() async throws {
        let service = FakeSimsService()
        let store = store(service)
        await store.load()
        #expect(store.isPolling)
        await store.pollOnce()
        #expect(service.polled == ["q8wd2c9vh4tz"])
        #expect(store.progress["q8wd2c9vh4tz"]?.iterationsDone == 1200)
        #expect(service.lists.count == 2)
        service.progressAnswer = { _ throws(APIError) in SimProgress(state: "done", iterationsDone: 3000, dps: 990, stage: 0, combosDone: 0, combosTotal: 0) }
        await store.pollOnce()
        #expect(service.lists.count == 4)
    }

    @Test func runAgain_startsThenReloads() async {
        let service = FakeSimsService()
        let store = store(service)
        await store.load()
        await store.runAgain("b7nf5x1pa3le")
        #expect(service.reruns == ["b7nf5x1pa3le"])
        #expect(store.rerun["b7nf5x1pa3le"] == .started)
        await store.runAgain("b7nf5x1pa3le")
        #expect(service.reruns.count == 1)
    }

    @Test func runAgain_onALapsedPlan_saysPremium() async throws {
        let service = FakeSimsService()
        let refusal = try RaidFixture.error("sims.rerun.error-premium-required", status: 402)
        service.rerunAnswer = { _ throws(APIError) in throw refusal }
        let store = store(service)
        await store.load()
        await store.runAgain("b7nf5x1pa3le")
        #expect(store.rerun["b7nf5x1pa3le"] == .lapsed)
    }

    @Test func markSeen_clearsTheDot() async throws {
        let service = FakeSimsService()
        let store = store(service)
        await store.load()
        await store.markSeen("km3xibdrsiyr")
        #expect(service.seen == ["km3xibdrsiyr"])
        #expect(try content(store).serverRows.first { $0.simId == "km3xibdrsiyr" }?.seen == true)
    }

    @Test func loadMore_appendsTheNextPage() async throws {
        let service = FakeSimsService()
        service.listAnswer = { page, _ throws(APIError) in
            let list = try RaidFixture.api(SimList.self, "sims.list")
            return SimList(rows: page == 1 ? list.rows : [list.rows[0]], total: 4, page: page, perPage: 100)
        }
        let store = store(service)
        await store.load()
        #expect(try content(store).hasMorePages)
        await store.loadMore()
        #expect(try content(store).mine.count == 4)
        #expect(try !content(store).hasMorePages)
    }
}
