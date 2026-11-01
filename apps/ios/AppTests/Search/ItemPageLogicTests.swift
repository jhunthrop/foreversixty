@testable import ForeverSixty
import Foundation
import Models
import Networking
import Testing

@MainActor
@Suite("Item page")
struct ItemPageLogicTests {
    private let raid = BisSelection(classSlug: "shaman", spec: "shaman-elemental", range: BisRange(band: 60, preset: "raid"), faction: "horde")

    private func band(_ selection: BisSelection) throws -> BisBand {
        let bis = try #require(try SearchTestData.tooltipSources().bis)
        return try #require(bis.bands.first { $0.band == selection.range.band && $0.preset == selection.range.preset && $0.faction == selection.faction })
    }

    @Test("your pick: Spritecaster Cape is Elemental's Back pick at 60, Phase 1 raid, Horde")
    func pick() throws {
        let card = ItemPageLogic.yourPick(itemID: SearchTestData.cape, band: try band(raid), selection: raid, specName: "Elemental", className: "shaman")
        #expect(card?.title == "Your Back pick")
        #expect(card?.meta == "Elemental · 60 · Phase 1 raid · Horde")
    }

    @Test("runner-up: Satyrmane Sash is runner-up 2 for Waist, and Elemental's spec row agrees")
    func runnerUp() throws {
        let card = ItemPageLogic.yourPick(itemID: SearchTestData.runnerUp, band: try band(raid), selection: raid, specName: "Elemental", className: "shaman")
        #expect(card?.title == "Runner-up 2 for Waist")
        #expect(card?.meta == "Elemental · 60 · Phase 1 raid · Horde")
        let row = try SearchTestData.uses(SearchTestData.runnerUp).specs.first { $0.spec == "shaman-elemental" }
        #expect(row?.rank == 2 && row?.band == 60 && row?.preset == "raid")
        #expect(ItemSpecTag(rank: row?.rank ?? 0).text == "Runner-up 2")
    }

    @Test("bare list: the meta has no preset word")
    func bareMeta() throws {
        let bare = raid.withRange(BisRange(band: 50, preset: "bare"))
        let band = BisBand(
            spec: "shaman-elemental", band: 50, preset: "bare", faction: "horde", race: "troll", talents: "", talentPoints: 0,
            scoreUnit: "dps", referenceDpsPerPoint: 1, role: "dps", slots: [BisSlot(slot: "back", verified: true, itemId: SearchTestData.cape)]
        )
        let card = ItemPageLogic.yourPick(itemID: SearchTestData.cape, band: band, selection: bare, specName: "Elemental", className: "shaman")
        #expect(card?.meta == "Elemental · 50 to 59 · Horde")
    }

    @Test("not on the list, or no list: no card")
    func notOnList() throws {
        #expect(ItemPageLogic.yourPick(itemID: 16866, band: try band(raid), selection: raid, specName: "Elemental", className: "shaman") == nil)
        #expect(ItemPageLogic.yourPick(itemID: SearchTestData.cape, band: nil, selection: raid, specName: "Elemental", className: "shaman") == nil)
    }

    @Test("sources: place and kind, a chance only when known, never 0%; quest level; world drop once")
    func sources() {
        let rows = ItemPageLogic.sources([
            ItemUseSource(kind: "dungeon", title: "High Interrogator Gerstahn", place: "Blackrock Depths"),
            ItemUseSource(kind: "raid", title: "Onyxia", place: "Onyxia's Lair", chance: 19.6),
            ItemUseSource(kind: "raid", title: "Ragnaros", place: "Molten Core", chance: 0.3),
            ItemUseSource(kind: "quest", title: "The Heart of the Mountain", place: "Horde", minLevel: 52),
            ItemUseSource(kind: "world_drop", title: "World drop"),
            ItemUseSource(kind: "vendor", title: "Meilosh"),
        ])
        #expect(rows.map(\.meta) == [
            "Blackrock Depths · Dungeon",
            "Onyxia's Lair · Raid · 20% from Onyxia",
            "Molten Core · Raid · <1% from Ragnaros",
            "Horde · Quest · Requires level 52",
            "",
            "Vendor",
        ])
        #expect(rows.map(\.chanceUnknown) == [true, false, false, false, false, false])
        #expect(rows[4].accessibilityLabel == "World drop")
        #expect(rows[0].accessibilityLabel == "High Interrogator Gerstahn, Blackrock Depths · Dungeon")
    }

    @Test("`Drop chance not known` prints once, only when a boss row has no chance")
    func chanceFootnote() {
        let unknown = ItemPageLogic.sources([
            ItemUseSource(kind: "dungeon", title: "A", place: "P"), ItemUseSource(kind: "dungeon", title: "B", place: "P"),
        ])
        #expect(ItemPageLogic.chanceFootnote(unknown) == "Drop chance not known")
        let known = ItemPageLogic.sources([ItemUseSource(kind: "raid", title: "Baron Geddon", place: "Molten Core", chance: 30)])
        #expect(ItemPageLogic.chanceFootnote(known) == nil)
        #expect(ItemPageLogic.chanceFootnote(ItemPageLogic.sources([ItemUseSource(kind: "vendor", title: "Meilosh")])) == nil)
    }

    @Test("not listed: Defiler's Mail Pauldrons is on no list at all, so no card and no rows")
    func notListed() throws {
        #expect(ItemPageLogic.yourPick(itemID: SearchTestData.notListed, band: try band(raid), selection: raid, specName: "Elemental", className: "shaman") == nil)
        let uses = try SearchTestData.uses(SearchTestData.notListed)
        #expect(uses.specs.isEmpty)
        #expect(ItemPageLogic.sources(uses.sources).map(\.kind) == ["vendor", "rep"])
    }

    @Test("best in slot for: the file's order, Pick or Runner-up n, raid or bare meta, six then more")
    func specs() throws {
        let catalog = try BisTestData.catalog()
        let rows = ItemPageLogic.specs(try SearchTestData.uses(SearchTestData.cape).specs, catalog: catalog)
        #expect(rows.first?.title == "Balance Druid")
        #expect(rows.first?.tag == .pick)
        #expect(rows.first?.meta == "60 · Phase 1 raid")
        #expect(rows.first?.icon != nil)
        let runner = ItemPageLogic.specs([ItemUseSpec(spec: "mage-fire", className: "mage", specName: "Fire", band: 50, preset: "bare", rank: 2)], catalog: catalog)
        #expect(runner.first?.tag == .runnerUp(2) && runner.first?.tag.text == "Runner-up 2")
        #expect(runner.first?.meta == "50")

        let many = Array(repeating: rows, count: 3).flatMap { $0 }.prefix(11)
        let shown = ItemPageLogic.visible(Array(many), expanded: false)
        #expect(shown.rows.count == 6 && shown.hidden == 5)
        #expect(ItemPageLogic.visible(Array(many), expanded: true).hidden == 0)
        #expect(ItemPageLogic.visible(Array(rows.prefix(6)), expanded: false).hidden == 0)
    }

    @Test("spell power, healing, mp5 and the like print only as the client's green Equip sentences")
    func equipLines() throws {
        let sources = try SearchTestData.tooltipSources()
        let model = { (id: Int) -> ItemTooltipModel in
            ItemTooltipModel.build(
                ItemTooltipRequest(itemID: id, fallback: HomeItemLook(id: id, name: nil, quality: nil, icon: nil, itemLevel: nil),
                                   build: "b", classSlug: "shaman", spec: nil, specName: nil, band: nil, bandLabel: nil, preset: nil,
                                   faction: nil, sourceLine: nil, wornID: nil),
                sources: sources
            )
        }
        let cape = model(SearchTestData.cape)
        #expect(Set(cape.stats) == ["+6 Intellect", "+6 Spirit", "+5 Stamina"])
        #expect(cape.equipLines == ["Equip: Increases damage and healing done by magical spells and effects by up to 14."])
        #expect(cape.binds == "Binds when picked up" && cape.typeName == "Cloth")

        let sash = model(SearchTestData.runnerUp)
        #expect(!sash.stats.contains { $0.contains(Copy.itemsStatSpellPower) })
        #expect(sash.equipLines == ["Equip: Increases damage and healing done by magical spells and effects by up to 14."])

        let pauldrons = model(SearchTestData.notListed)
        #expect(!pauldrons.stats.contains { $0.contains(Copy.itemsStatMp5) })
        #expect(pauldrons.equipLines == ["Equip: Restores 4 mana per 5 sec."])
        #expect(pauldrons.setName != nil)
    }

    @Test("a rating stat is an Equip sentence only when percent_stats states the same figure")
    func percentStats() {
        #expect(ItemVocabulary.percentLine("crit", value: 1) == "Equip: Improves your chance to get a critical strike by 1%.")
        #expect(ItemVocabulary.equipLine("crit", value: 1) == nil)
        #expect(ItemVocabulary.percentLine("defense", value: 7) == "Equip: Increased Defense +7.")
        #expect(ItemVocabulary.equipLine("defense", value: 7) == nil)
    }

    @Test("a spec row opens that list: its band, preset and faction, else the viewer's side")
    func selection() {
        let use = ItemUseSpec(spec: "mage-frost", className: "mage", specName: "Frost", band: 60, preset: "raid", faction: nil, rank: 0)
        let picked = ItemPageLogic.selection(use, viewerFaction: "alliance")
        #expect(picked == BisSelection(classSlug: "mage", spec: "mage-frost", range: BisRange(band: 60, preset: "raid"), faction: "alliance"))
        #expect(ItemPageLogic.selection(use, viewerFaction: nil).faction == "horde")
    }

    @Test("the tooltip reads the viewer's class file when they can use the item, else the first that can")
    func tooltipClass() throws {
        let cape = try #require(try SearchTestData.corpus().item(SearchTestData.cape))
        let own = BisSelection.forCharacter(SearchTestData.lightningbolltz, catalog: try BisTestData.catalog())
        let request = ItemPageLogic.tooltipRequest(item: cape, route: ItemPageRoute(itemID: cape.id), build: "b", character: SearchTestData.lightningbolltz, own: own)
        #expect(request.classSlug == "shaman")
        #expect(request.spec == "shaman-elemental" && request.preset == "raid" && request.faction == "horde")
        let helm = SearchItem(id: 16866, name: "Helm of Might", icon: "x", quality: 4, slot: "head", itemLevel: 66, requiredLevel: 60, classes: ["warrior"])
        #expect(ItemPageLogic.tooltipRequest(item: helm, route: ItemPageRoute(itemID: 16866), build: "b", character: SearchTestData.lightningbolltz, own: own).classSlug == "warrior")
    }

    @Test("the page loads the fixture item: head from the index, sources, specs; an unknown id says so")
    func store() async throws {
        let statics = StaticDataStore(client: StaticDataClient(transport: FixtureTransport(directory: HomeFixtures.directory)), directory: nil)
        let deps = SearchDependencies(
            index: SearchIndexReader(store: statics), catalog: ItemCatalog(store: statics), specs: try BisTestData.catalog(),
            characters: FakeCharacterSearch(.success([])), recent: MemoryRecentSearches(),
            build: { SearchTestData.build }, context: { SearchContext(character: SearchTestData.fury) }, sleep: { _ in }
        )
        let page = ItemPageStore(route: ItemPageRoute(itemID: SearchTestData.cape), dependencies: deps)
        await page.load()
        #expect(page.state.look.name == "Spritecaster Cape")
        #expect(page.state.meta == "Rare · Back · item level 52")
        guard case let .loaded(sources) = page.state.sources else { Issue.record("sources"); return }
        #expect(sources.count == 2)
        guard case let .loaded(specs) = page.state.specs else { Issue.record("specs"); return }
        #expect(!specs.isEmpty)
        #expect(page.state.pick == nil)

        let unknown = ItemPageStore(route: ItemPageRoute(itemID: 999_999_999), dependencies: deps)
        await unknown.load()
        #expect(!unknown.state.isKnown)
    }
}
