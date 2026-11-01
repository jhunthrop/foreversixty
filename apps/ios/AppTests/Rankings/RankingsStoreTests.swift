@testable import ForeverSixty
import Foundation
import Models
import Networking
import Routing
import Testing

/// A scripted F11 service: answers from the fixtures unless told to fail, and records the reads.
actor FakeRankingsService: RankingsService {
    var pages: [RankingsPage]
    var failRankings: APIError?
    var failSetup: APIError?
    var characterResult: Result<CharacterProfile, APIError>
    var failWatch: APIError?
    private(set) var queries: [RankingsQuery] = []
    private(set) var watched: [String] = []
    private(set) var unwatched: [String] = []

    init(pages: [RankingsPage], character: Result<CharacterProfile, APIError>) {
        self.pages = pages
        characterResult = character
    }

    func set(failRankings: APIError?) { self.failRankings = failRankings }
    func set(failSetup: APIError?) { self.failSetup = failSetup }
    func set(character: Result<CharacterProfile, APIError>) { characterResult = character }
    func set(failWatch: APIError?) { self.failWatch = failWatch }

    func rankings(_ query: RankingsQuery) async throws(APIError) -> RankingsPage {
        queries.append(query)
        if let failRankings { throw failRankings }
        return pages[min(query.page - 1, pages.count - 1)]
    }

    func encounters() async throws(APIError) -> RankedEncounters {
        if let failSetup { throw failSetup }
        do { return try RankingsTestData.encounters() } catch { throw .decoding("\(error)") }
    }

    func phases() async throws(APIError) -> Phases {
        if let failSetup { throw failSetup }
        do { return try RankingsTestData.phases() } catch { throw .decoding("\(error)") }
    }

    func character(_: CharacterRef) async throws(APIError) -> CharacterProfile { try characterResult.get() }

    func watching() async throws(APIError) -> WatchingList {
        do { return try RankingsTestData.watching() } catch { throw .decoding("\(error)") }
    }

    func watch(_ ref: CharacterRef) async throws(APIError) {
        if let failWatch { throw failWatch }
        watched.append(ref.key)
    }

    func unwatch(_ ref: CharacterRef) async throws(APIError) {
        if let failWatch { throw failWatch }
        unwatched.append(ref.key)
    }
}

@MainActor
@Suite("Rankings stores")
struct RankingsStoreTests {
    private let offline = APIError.transport("offline")
    private let notFound = APIError.server(status: 404, body: ErrorBody(code: "not_found", message: "no such character"), requestID: nil)

    private func store(_ service: FakeRankingsService, cache: MemoryRaidCache = MemoryRaidCache(), account: RankingsAccount = RankingsTestData.account(), encounter: String? = nil) -> RankingsStore {
        RankingsStore(service: service, cache: cache, specs: RankingsTestData.specs, now: { RankingsTestData.now }, account: account, encounter: encounter)
    }

    private func service() throws -> FakeRankingsService {
        FakeRankingsService(pages: [try RankingsTestData.page()], character: .success(try RankingsTestData.lightningbolltz()))
    }

    @Test func load_aimsAtTheCharacter_andDefaults() async throws {
        let service = try service()
        let store = store(service)
        await store.load()
        #expect(store.filters == RankingsTestData.filters)
        #expect(store.encounterName == "Onyxia")
        guard case let .loaded(content) = store.phase else { Issue.record("not loaded"); return }
        #expect(content.rows.count == 100)
        let query = try #require(await service.queries.first)
        #expect(query.character == "us/normal/lightningbolltz")
        #expect(query.classSlug == "shaman")
        #expect(query.distinct == "character")
    }

    @Test func load_aLinkPicksTheEncounter() async throws {
        let store = store(try service(), encounter: "227")
        await store.load()
        #expect(store.filters?.encounter == "emperor-dagran-thaurissan")
    }

    @Test func signedOut_aimsNothing_thenThePickerAims() async throws {
        let service = try service()
        let store = store(service, account: .signedOut)
        await store.load()
        #expect(store.filters?.classSlug == "")
        #expect(await service.queries.first?.character == nil)
        await store.choose(classSlug: "priest", spec: "priest-holy")
        #expect(store.filters?.classSlug == "priest")
        #expect(store.filters?.spec == "Holy")
        #expect(store.filters?.metric == .hps)
    }

    @Test func paging_stopsWhenAPageAddsNothing() async throws {
        let service = try service()
        let store = store(service)
        await store.load()
        await store.loadMore()
        guard case let .loaded(content) = store.phase else { Issue.record("not loaded"); return }
        #expect(content.rows.count == 100)
        #expect(!content.hasMorePages)
        await store.loadMore()
        #expect(await service.queries.count == 2)
    }

    @Test func offline_showsTheCachedBoardStamped() async throws {
        let cache = MemoryRaidCache()
        let service = try service()
        await store(service, cache: cache).load()
        await service.set(failRankings: offline)
        await service.set(failSetup: offline)
        let again = store(service, cache: cache)
        await again.load()
        guard case let .loaded(content) = again.phase else { Issue.record("not loaded"); return }
        #expect(content.offlineSince == RankingsTestData.now)
    }

    @Test func error_withNothingCached() async throws {
        let service = try service()
        await service.set(failRankings: .unexpectedStatus(500))
        let store = store(service)
        await store.load()
        #expect(store.phase == .failed)
    }

    @Test func applyAndReset() async throws {
        let store = store(try service())
        await store.load()
        var next = try #require(store.filters)
        next.region = "eu"
        await store.apply(next)
        #expect(store.filters?.region == "eu")
        await store.reset()
        #expect(store.filters == RankingsTestData.filters)
    }

    @Test func homeLine_readsTheDefaults() async throws {
        let line = await RankingsStore.homeLine(
            service: try service(), cache: MemoryRaidCache(), specs: RankingsTestData.specs, now: RankingsTestData.now,
            current: RankingsTestData.lightning
        )
        #expect(line == "14th of 212 Elemental Shaman characters on Onyxia")
    }

    @Test func signOut_clearsTheRankingsCache() {
        let cache = DiskRaidCache.rankings()
        cache.write(["a"], key: "probe", at: RankingsTestData.now)
        #expect(cache.read([String].self, key: "probe") != nil)
        cache.removeAll()
        #expect(cache.read([String].self, key: "probe") == nil)
    }

    // MARK: character page

    private func character(_ service: FakeRankingsService, cache: MemoryRaidCache = MemoryRaidCache()) throws -> CharacterStore {
        let ref = try #require(CharacterRef(key: "us/normal/voltaris", name: "Voltaris"))
        let items = try RankingsTestData.items()
        return CharacterStore(ref: ref, service: service, cache: cache, now: { RankingsTestData.now }, items: { _ in items })
    }

    @Test func character_loadsAndReadsItems() async throws {
        let service = FakeRankingsService(pages: [], character: .success(try RankingsTestData.voltaris()))
        let store = try character(service)
        await store.load()
        guard case let .loaded(content) = store.phase else { Issue.record("not loaded"); return }
        #expect(content.profile.character.name == "Voltaris")
        #expect(store.items[252604]?.name == "Blue Dragonscale Helm")
    }

    @Test func character_404IsNotFound() async throws {
        let service = FakeRankingsService(pages: [], character: .failure(notFound))
        let store = try character(service)
        await store.load()
        #expect(store.phase == .notFound)
    }

    @Test func character_offlineShowsTheCache() async throws {
        let cache = MemoryRaidCache()
        let service = FakeRankingsService(pages: [], character: .success(try RankingsTestData.voltaris()))
        await (try character(service, cache: cache)).load()
        await service.set(character: .failure(offline))
        let again = try character(service, cache: cache)
        await again.load()
        guard case let .loaded(content) = again.phase else { Issue.record("not loaded"); return }
        #expect(content.offlineSince == RankingsTestData.now)
    }

    // MARK: watching

    @Test func watch_isOptimistic_andConfirms() async throws {
        let service = try service()
        let list = WatchList(service: service, signedIn: true)
        await list.load()
        #expect(list.isWatching("us/normal/voltaris"))
        let zulthrax = try #require(CharacterRef(key: "us/normal/zulthrax", name: "Zulthrax"))
        await list.toggle(zulthrax)
        #expect(list.isWatching(zulthrax.key))
        #expect(list.confirmed == 1)
        #expect(await service.watched == [zulthrax.key])
        let voltaris = try #require(CharacterRef(key: "us/normal/voltaris", name: "Voltaris"))
        await list.toggle(voltaris)
        #expect(!list.isWatching(voltaris.key))
        #expect(await service.unwatched == [voltaris.key])
    }

    @Test func watch_failureRevertsAndNames() async throws {
        let service = try service()
        await service.set(failWatch: .server(status: 409, body: ErrorBody(code: "limit", message: "limit"), requestID: nil))
        let list = WatchList(service: service, signedIn: true)
        let zulthrax = try #require(CharacterRef(key: "us/normal/zulthrax", name: "Zulthrax"))
        await list.toggle(zulthrax)
        #expect(!list.isWatching(zulthrax.key))
        #expect(list.failedKey == zulthrax.key)
        #expect(list.confirmed == 0)
    }

    @Test func router_opensRankingsAndPlayerLinks() throws {
        let router = RankingsRouter()
        let rankings = try #require(URL(string: "https://foreversixty.gg/app/rankings/onyxia").flatMap(DeepLink.init(url:)))
        #expect(RankingsRouter.handles(rankings))
        router.open(rankings)
        #expect(router.showsBoard)
        #expect(router.encounter == "onyxia")
        let player = try #require(URL(string: "https://foreversixty.gg/app/player/us/normal/voltaris").flatMap(DeepLink.init(url:)))
        router.open(player)
        #expect(router.player?.ref.key == "us/normal/voltaris")
        #expect(!router.showsBoard)
    }

    @Test func characterRef_links() throws {
        let ref = try #require(CharacterRef(key: "us/normal/voltaris", name: "Voltaris"))
        #expect(ref.webURL.absoluteString == "https://foreversixty.gg/character/us/normal/voltaris")
        #expect(ref.historyURL.absoluteString == "https://foreversixty.gg/character/us/normal/voltaris#history")
        #expect(CharacterRef(key: "nope", name: "x") == nil)
    }
}
