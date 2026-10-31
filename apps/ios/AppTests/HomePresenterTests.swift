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
    @Test("rows read measured gains largest first, then not modelled, then empty; the card shows five")
    func presentationOrder() throws {
        let home = try HomeFixtures.home()
        let ordered = HomePresenter.presentationOrder(home.upgrades).map(\.slot)
        #expect(ordered == ["main_hand", "ranged", "legs", "feet", "chest", "head", "waist", "hands", "back", "finger2", "wrist", "trinket2"])
        let model = HomeFixtures.presenter().model(home)
        guard case let .upgrades(rows, footer) = model.body else { Issue.record("no rows"); return }
        #expect(rows.map(\.slot) == Array(ordered.prefix(5)))
        #expect(footer == "7 more upgrades · all 17 slots")
    }

    @Test("the hero: two figures, the note with the unmodelled count, the sync age")
    func hero() throws {
        let model = HomeFixtures.presenter().model(try HomeFixtures.home())
        #expect(model.heroMeta == "Orc Beast Mastery Hunter · Level 60 · Olympus XXVII")
        #expect(model.count == "12 upgrades")
        #expect(model.gain?.figure == "+102.5")
        #expect(model.gain?.unit == "DPS")
        #expect(model.note == "Best in slot for 60, together · 2 gains not modelled yet")
        #expect(model.synced == "Synced 2 h ago from the addon")
        #expect(model.bandHead == "60 · raid")
    }

    @Test(
        "Home's band is the site rule over the character's level, the list BiS opens on",
        arguments: [(49, 40, "40 to 49"), (50, 50, "50 to 59"), (59, 50, "50 to 59"), (60, 60, "60 · raid")]
    )
    func bandFollowsTheLevel(level: Int, band: Int, head: String) throws {
        let home = try HomeFixtures.home()
        let character = HomeFixtures.velaren(level: level)
        let range = try #require(HomePresenter.range(home, character: character))
        #expect(range == BisRange.forLevel(level))
        #expect(range.band == band)
        #expect(HomeFixtures.presenter(character: character).model(home).bandHead == head)
        let tooltip = ItemTooltipRequest.forHome(HomeItemLook(id: 1, name: nil, quality: nil, icon: nil, itemLevel: nil),
                                                 home: home, character: character)
        #expect(tooltip.band == range.band && tooltip.preset == range.preset && tooltip.bandLabel == head)
    }

    @Test("with no level known, the server's band goes through the same rule")
    func bandWithoutALevel() throws {
        let home = try HomeFixtures.home()
        #expect(HomePresenter.range(home, character: HomeFixtures.velaren(level: nil)) == BisRange.forLevel(60))
    }

    @Test("a measured gain is the server's figure, the source the server's formatted line")
    func gains() throws {
        let home = try HomeFixtures.home()
        let model = HomeFixtures.presenter().model(home)
        guard case let .upgrades(rows, _) = model.body else { Issue.record("no rows"); return }
        #expect(rows[0].gain == .measured(figure: "+21.2", unit: "DPS"))
        #expect(rows[0].meta == "Main hand · Crafted: Blacksmithing")
        #expect(HomePresenter.presentationOrder(home.upgrades).last?.gainWords == "Gain not modelled yet")
    }

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

    @Test("every slot matching says so in place of the figures, with the Already line")
    func allMatch() throws {
        let model = HomeFixtures.presenter().model(try HomeFixtures.home("character.home.empty"))
        #expect(model.body == .allMatch("Every slot matches the 60 list."))
        #expect(model.count == nil)
        #expect(model.already?.hasPrefix("Already best in slot: Head, Neck, Shoulder") == true)
    }

    @Test("no gear export: the sentence, Get the addon, the list button and the note")
    func noGear() throws {
        let model = HomeFixtures.presenter().model(try HomeFixtures.home("character.home.no-gear"))
        guard case let .message(text, buttons, note) = model.body else { Issue.record("no message"); return }
        #expect(text == "Add the addon to see which slots you can upgrade.")
        #expect(buttons.map(\.kind) == [.getAddon, .openList])
        #expect(buttons.map(\.title) == ["Get the addon", "Beast Mastery 60 list"])
        #expect(note?.hasPrefix("The addon sends your gear") == true)
        #expect(model.next.first?.line == "Not available yet: no talent export for this character. Open the addon once to send it.")
        #expect(model.next.last?.line == "No sim yet.")
    }

    @Test("under 20, no spec and no list each have their sentence", arguments: [
        ("under_20", "Best in slot starts at 20."),
        ("no_spec", "Set a spec in the planner to see this."),
        ("no_list", "No best in slot list published yet for this spec and level range."),
    ])
    func emptyStates(state: String, sentence: String) throws {
        let model = HomeFixtures.presenter().model(try HomeFixtures.home(state: state))
        guard case let .message(text, _, _) = model.body else { Issue.record("no message"); return }
        #expect(text == sentence)
    }

    @Test("talents: the differ count and its tag")
    func talents() throws {
        let row = HomeFixtures.presenter().model(try HomeFixtures.home()).next.first
        #expect(row?.line == "1 of your 51 points differ from the 60 build")
        #expect(row?.tag == .differs)
    }

    @Test("sims: the newest finished run for this spec")
    func sims() throws {
        let sim = FinishedSim(simId: "s1", spec: "hunter-beast-mastery", title: "t", dps: 812.4, finishedAt: HomeFixtures.now)
        let row = HomeFixtures.presenter(sims: [sim]).model(try HomeFixtures.home()).next.last
        #expect(row?.line == "Your Beast Mastery sim finished: 812 DPS")
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
        #expect(rows.first { $0.slot == "off_hand" }?.wornName == "none: Legionite Glaive is two-handed.")
        let trinket = try #require(rows.first { $0.slot == "trinket2" })
        #expect(trinket.worn == nil && trinket.wornName == "Empty" && trinket.pick?.id == 249_469)
        #expect(rows.first { $0.slot == "shoulder" }?.isBestInSlot == true)
        #expect(rows.first { $0.slot == "wrist" }?.wornName == "Not in our item data yet")
        #expect(rows.first { $0.slot == "hands" }?.wornName == "Mourning Handguards")
    }
}
