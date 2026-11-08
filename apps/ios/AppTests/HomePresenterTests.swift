@testable import ForeverSixty
import Foundation
import Models
import Networking
import Testing

/// F2's logic, one test list shared with Android (`HomePresenterTest.kt`): Home renders the
/// server's answer in the server's order and formats it; it never ranks or computes a gain.
@MainActor
@Suite("Home answer")
struct HomePresenterTests {

    @Test("a negative gain keeps its true minus")
    func negativeGain() {
        #expect(ItemNumbers.signedTenths(-0.8) == "\u{2212}0.8")
    }

    @Test("what a row replaces: the worn item, an empty slot, or an item our data does not know")
    func replaces() throws {
        let upgrades = try HomeFixtures.home().upgrades
        #expect(HomePresenter.replaces(upgrades.first { $0.slot == "ranged" }?.worn) == "Replaces Dark Iron Rifle, ilvl 55")
        #expect(HomePresenter.replaces(nil) == "Empty slot")
        #expect(HomePresenter.replaces(upgrades.first { $0.slot == "wrist" }?.worn) == "Not in our item data yet")
    }

    @Test("the offline stamp: today, this week, older")
    func stamp() {
        let time = HomeTime(now: HomeFixtures.now, calendar: HomeFixtures.utc)
        #expect(time.stamp(HomeFixtures.now.addingTimeInterval(-2 * 3600)) == "9:15 pm")
        #expect(time.stamp(HomeFixtures.now.addingTimeInterval(-2 * 86400)) == "Tue 11:15 pm")
        #expect(time.stamp(HomeFixtures.now.addingTimeInterval(-30 * 86400)) == "29 Sep, 11:15 pm")
    }
}

@MainActor
@Suite("Home store")
struct HomeStoreTests {
    private func store(_ result: Result<CharacterHome, APIError>, cache: MemoryHomeCache) -> HomeStore {
        HomeStore(fetch: { _ throws(APIError) in try result.get() }, cache: cache, now: { HomeFixtures.now })
    }

    @Test("a first load with nothing cached shows the answer and caches it")
    func firstLoad() async throws {
        let home = try HomeFixtures.home()
        let cache = MemoryHomeCache()
        let store = store(.success(home), cache: cache)
        await store.load(HomeFixtures.velaren)
        #expect(store.phase == .loaded(home, offlineSince: nil))
        #expect(cache.load(HomeFixtures.velaren.key)?.home == home)
    }

    @Test("no connection with a cached answer: the cached answer under the offline strip")
    func offline() async throws {
        let home = try HomeFixtures.home()
        let at = HomeFixtures.now.addingTimeInterval(-600)
        let cache = MemoryHomeCache([HomeFixtures.velaren.key: HomeCacheEntry(home: home, fetchedAt: at)])
        let store = store(.failure(.transport("offline")), cache: cache)
        await store.load(HomeFixtures.velaren)
        #expect(store.phase == .loaded(home, offlineSince: at))
    }

    @Test("no connection and nothing cached is the error state")
    func offlineEmpty() async {
        let store = store(.failure(.transport("offline")), cache: MemoryHomeCache())
        await store.load(HomeFixtures.velaren)
        #expect(store.phase == .failed(cached: nil))
    }

    @Test("a server error is the error state, never the offline strip")
    func serverError() async throws {
        let home = try HomeFixtures.home()
        let cache = MemoryHomeCache([HomeFixtures.velaren.key: HomeCacheEntry(home: home, fetchedAt: HomeFixtures.now)])
        let failure = APIError.server(status: 500, body: ErrorBody(code: "internal", message: "x"), requestID: nil)
        let store = store(.failure(failure), cache: cache)
        await store.load(HomeFixtures.velaren)
        #expect(store.phase == .failed(cached: home))
    }

    @Test("the cache entry survives a round trip through its file")
    func cacheRoundTrip() throws {
        let directory = FileManager.default.temporaryDirectory.appending(path: "home-cache-\(UUID().uuidString)")
        defer { try? FileManager.default.removeItem(at: directory) }
        let cache = FileHomeCache(directory: directory)
        let entry = HomeCacheEntry(home: try HomeFixtures.home(), fetchedAt: HomeFixtures.now)
        cache.save(entry, for: HomeFixtures.velaren.key)
        #expect(cache.load(HomeFixtures.velaren.key) == entry)
    }

    @Test("fixture mode answers the home route and the tooltip's static files")
    func fixtureRoute() async throws {
        let transport = FixtureTransport(directory: HomeFixtures.directory)
        let client = APIClient(baseURL: try #require(URL(string: "https://api.foreversixty.gg")), transport: transport)
        let home = HomeStore.live(client: client, cache: MemoryHomeCache())
        await home.load(HomeFixtures.velaren)
        guard case let .loaded(answer, _) = home.phase else { Issue.record("not loaded"); return }
        #expect(answer.upgradeCount == 12)
        let catalog = ItemCatalog(client: StaticDataClient(transport: transport), directory: nil)
        let request = ItemTooltipRequest.forHome(
            HomeItemLook(id: 252_605, name: nil, quality: nil, icon: nil, itemLevel: nil), home: answer, character: HomeFixtures.velaren
        )
        let sources = try await catalog.sources(for: request)
        #expect(sources.items[252_605] != nil)
        #expect(sources.bis != nil && !sources.sets.isEmpty)
    }
}

@MainActor
@Suite("Item tooltip and gear")
struct ItemTooltipTests {
    private func request(_ id: Int, wornID: Int? = nil, sourceLine: String? = nil) throws -> ItemTooltipRequest {
        .forHome(HomeItemLook(id: id, name: "Fallback", quality: 4, icon: nil, itemLevel: 60),
                 home: try HomeFixtures.home(), character: HomeFixtures.velaren, sourceLine: sourceLine, wornID: wornID)
    }

    @Test("a pick: binds, armor type, stats, levels, its source, the next best and the footer")
    func pick() throws {
        let model = ItemTooltipModel.build(try request(252_605), sources: try HomeFixtures.tooltipSources())
        #expect(model.headerMeta == "Epic · Head")
        #expect(model.binds == "Binds when equipped")
        #expect(model.slotName == "Head")
        #expect(model.typeName == "Mail")
        #expect(model.armor?.hasSuffix(" Armor") == true)
        #expect(model.stats.allSatisfy { $0.hasPrefix("+") })
        #expect(model.itemLevel == "Item Level 61")
        #expect(model.source == "Crafted: Leatherworking")
        #expect(model.alternatives.count == 2)
        #expect(model.pickFooter == "Best in slot for Head · Beast Mastery 60 · raid")
    }

    @Test("a vendor item adds the standing it requires to the site's line")
    func reputation() throws {
        let sources = try HomeFixtures.tooltipSources()
        #expect(ItemTooltipModel.build(try request(276_202), sources: sources).source == "Vendor: Friz Frazzlespark · requires Friendly")
    }

    @Test("the worn runner-up leads with You wear this")
    func youWear() throws {
        let model = ItemTooltipModel.build(try request(9533, wornID: 11862), sources: try HomeFixtures.tooltipSources())
        #expect(model.alternatives.map(\.meta).contains("You wear this · \u{2212}2.3 DPS vs the pick"))
        #expect(model.alternatives.first?.source == "2% from Ancient Treasure")
    }

    @Test("a weapon reads Two-Hand and its type, as the client prints it")
    func weapon() throws {
        let model = ItemTooltipModel.build(try request(250_619), sources: try HomeFixtures.tooltipSources())
        #expect(model.slotName == "Two-Hand")
        #expect(model.typeName == "Polearm")
        #expect(model.weaponDamage?.hasSuffix(" Damage") == true)
        #expect(model.weaponDPS?.hasSuffix("damage per second)") == true)
    }

    @Test("an item the build does not list keeps the caller's look and says so")
    func unknown() throws {
        let model = ItemTooltipModel.build(try request(999_999, sourceLine: "Somewhere"), sources: try HomeFixtures.tooltipSources())
        #expect(!model.known)
        #expect(model.look.name == "Fallback")
        #expect(model.source == "Somewhere")
    }

    @Test("Your gear: 17 rows in frame order; the off hand names the two-hander")
    func gear() throws {
        let home = try HomeFixtures.home()
        let sources = try HomeFixtures.tooltipSources()
        let band = ItemTooltipModel.band(for: try request(0), in: sources.bis)
        let blocker = GearModel.offHandBlocker(home: home, gear: [:], items: sources.items, band: band)
        #expect(blocker == "Legionite Glaive")
        let rows = GearModel.rows(home: home, gear: [:], items: sources.items, offHandBlockedBy: blocker)
        #expect(rows.map(\.slot) == ItemVocabulary.frameSlots)
        #expect(rows.count == 17)
        #expect(rows.first { $0.slot == "off_hand" }?.wornName == "Two-hand in main hand")
        let trinket = try #require(rows.first { $0.slot == "trinket2" })
        #expect(trinket.worn == nil && trinket.wornName == "Empty" && trinket.pick?.id == 249_469)
        #expect(rows.first { $0.slot == "shoulder" }?.isBestInSlot == true)
        #expect(rows.first { $0.slot == "wrist" }?.wornName == "Wrist · Item 999999")
        #expect(rows.first { $0.slot == "hands" }?.wornName == "Mourning Handguards")
    }
}
