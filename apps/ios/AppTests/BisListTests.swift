@testable import ForeverSixty
import Foundation
import Models
import Networking
import Testing

/// The BiS list's rules (spec F3.1). The same list runs on Android (BisListTest.kt).
@Suite("BiS list")
struct BisListTests {
    private let utc = TimeZone.gmt

    private func raidList(worn: Set<Int> = []) throws -> BisList {
        let file = try BisTestData.bis()
        let band = try #require(file.band(for: BisTestData.raid.range, faction: "horde"))
        return BisListBuilder(items: try BisTestData.items().items, worn: worn, timeZone: utc)
            .list(file: file, band: band, specLabel: "Fury Warrior", rangeLabel: Copy.bisRangeP1)
    }

    @Test("the level range follows the site's band rule; 60 reads the raid list")
    func rangeForLevel() {
        #expect(BisRange.forLevel(58) == BisRange(band: 50, preset: "bare"))
        #expect(BisRange.forLevel(60) == BisRange(band: 60, preset: "raid"))
        #expect(BisRange.forLevel(12) == BisRange(band: 20, preset: "bare"))
        #expect(BisRange.forLevel(29) == BisRange(band: 20, preset: "bare"))
        #expect(BisRange.forLevel(49) == BisRange(band: 40, preset: "bare"))
        #expect(BisRange.forLevel(50) == BisRange(band: 50, preset: "bare"))
        #expect(BisRange.forLevel(59) == BisRange(band: 50, preset: "bare"))
        #expect(BisRange.forLevel(61) == BisRange(band: 60, preset: "raid"))
        #expect(BisRange.all.map { BisLabels.range($0) } == ["20 to 29", "30 to 39", "40 to 49", "50 to 59", "60", "60 · Phase 1 raid"])
    }

    @Test("a band is the file's own entry for range, faction and preset, else the bare list")
    func bandSelection() throws {
        let file = try BisTestData.bis()
        #expect(file.band(for: BisRange(band: 60, preset: "raid"), faction: "horde")?.preset == "raid")
        #expect(file.band(for: BisRange(band: 50, preset: "raid"), faction: "horde")?.preset == "bare")
        #expect(file.band(for: BisRange(band: 60, preset: "raid"), faction: "alliance") == nil)
    }

    @Test("17 rows in the file's order, never re-ranked")
    func rowsInFileOrder() throws {
        let list = try raidList()
        let file = try BisTestData.bis()
        let band = try #require(file.band(for: BisTestData.raid.range, faction: "horde"))
        #expect(list.rows.map(\.slot) == band.slots.map(\.slot))
        #expect(list.rows.count == 17)
        #expect(list.rows.compactMap(\.pick?.item.id) == band.slots.compactMap(\.itemId))
    }

    @Test("the set line carries the figure, the unit and the Updated date")
    func setDpsLine() throws {
        #expect(try raidList().setDpsLine == "894.1 DPS at best in slot for your level range · Updated 30 Oct 2026")
    }

    @Test("the next line: a signed gap with a true minus, same below 0.05, nothing else ranked")
    func nextLine() throws {
        let rows = try raidList().rows
        #expect(rows[0].pick?.next == "Next: Champion's Plate Helm, \u{2212}15.0 DPS")
        #expect(rows[2].pick?.next == "Next: Warlord's Plate Shoulders, same DPS")
        let none = BisListBuilder.nextLine(nil, unit: .dps)
        #expect(none == "Nothing else ranked")
        #expect(BisListBuilder.signedDelta(-0.049) == nil)
        #expect(BisListBuilder.signedDelta(0.05) == "+0.1")
    }

    @Test("healer and tank bands name their own measure")
    func units() {
        #expect(BisUnit(role: "healer").word == "HPS")
        #expect(BisUnit(role: "tank").word == "score")
        #expect(BisUnit(role: nil).word == "DPS")
    }

    @Test("the source line in the site's words")
    func sourceLines() throws {
        let rows = try raidList().rows
        #expect(rows[0].pick?.slotSource == "Head · Crafted: Blacksmithing")
        #expect(rows[1].pick?.slotSource == "Neck · Quest: Falrin's Vendetta")
        #expect(rows[2].pick?.slotSource == "Shoulder · Vendor: Creeg Bothunk")
        #expect(rows[5].pick?.slotSource == "Wrist · Warsong Outriders")
        #expect(rows[10].slotLabel == "Finger 1")
        #expect(rows[11].slotLabel == "Finger 2")
        #expect(rows[12].slotLabel == "Trinket 1")
        #expect(BisListBuilder.sourceLine("PvP rank 14 · Champion · Horde", kind: "vendor") == "PvP rank 14 · Champion · Horde")
    }

    @Test("an unmodelled proc is said; the worn pick is tagged")
    func flags() throws {
        let rows = try raidList(worn: [12640]).rows
        #expect(rows[10].pick?.isUnmodelled == true)
        #expect(rows[0].pick?.isUnmodelled == false)
        #expect(rows[0].pick?.isWorn == true)
        #expect(rows[1].pick?.isWorn == false)
    }

    @Test("an off hand under a two-hander and a slot with no pick")
    func emptyRows() {
        let builder = BisListBuilder(items: [], worn: [], timeZone: utc)
        let offHand = builder.row(
            BisSlot(slot: "off_hand", verified: false, emptyReason: "two_hand_equipped"),
            unit: .dps,
            mainHand: "Argent Crusader"
        )
        #expect(offHand.content == .empty("Off hand · none: Argent Crusader is two-handed."))
        let trinket = builder.row(
            BisSlot(slot: "trinket1", verified: false, emptyReason: "no_sourced_item"),
            unit: .dps,
            mainHand: nil
        )
        #expect(trinket.content == .empty("Trinket 1 · no pick for this range."))
    }

    @Test("search keeps rows whose pick or runner-up matches")
    func search() throws {
        let list = try raidList()
        #expect(list.filtered(by: "lionheart").map(\.slot) == ["head"])
        #expect(list.filtered(by: "Champion's Plate Helm").map(\.slot) == ["head"])
        #expect(list.filtered(by: " ").count == 17)
    }

    @Test("every published BiS file of the active build decodes")
    func realFilesDecode() throws {
        let directory = BisTestData.repo.appending(path: "data/builds/\(BisTestData.build)/bis")
        let files = try FileManager.default.contentsOfDirectory(at: directory, includingPropertiesForKeys: nil)
            .filter { $0.pathExtension == "json" }
        #expect(files.count >= 28)
        for file in files {
            _ = try WireFormatProbe.decode(BisFile.self, file)
        }
    }

    @Test("the character's own list: its spec, range and faction")
    func characterSelection() throws {
        let me = try BisTestData.me()
        let catalog = try BisTestData.catalog()
        let main = try #require(me.mainCharacter)
        let own = try #require(BisSelection.forCharacter(main, catalog: catalog))
        #expect(own == BisSelection(classSlug: "warrior", spec: "warrior-fury", range: BisRange(band: 50, preset: "bare"), faction: "horde"))
        #expect(BisSelection.signedOutDefault(catalog).classSlug == "warrior")
        #expect(catalog.specs(of: "shaman").map(\.name) == ["Elemental", "Enhancement", "Restoration"])
    }

    @Test("Yours and the way back follow the character's list")
    func yoursAndBack() throws {
        let own = BisTestData.raid
        var context = BisContext(selection: own, characterSelection: own, characterName: "Obnoxious Yell")
        #expect(context.isYours)
        #expect(context.backTarget == nil)
        context.selection = own.withFaction("alliance")
        #expect(!context.isYours)
        #expect(context.backTarget?.name == "Obnoxious Yell")
        let signedOut = BisContext(selection: own)
        #expect(!signedOut.isYours)
        #expect(signedOut.backTarget == nil)
    }
}

enum WireFormatProbe {
    static func decode<T: Decodable>(_: T.Type, _ url: URL) throws -> T {
        try WireFormat.decoder().decode(T.self, from: Data(contentsOf: url))
    }
}

@Suite("BiS bundle")
struct BisBundleTests {
    @Test("the app carries the spec catalogue, the fallback build, the crests and the faction marks")
    func bundledResources() {
        #expect(SpecCatalog.bundled().specs.count >= 28)
        #expect(BundledBuild.read() != nil)
        #expect(BrandImages.crest("shaman") != nil)
        #expect(BrandImages.faction("horde") != nil)
    }
}

@Suite("BiS tooltip and dates")
struct BisTooltipTests {
    private func real(_ spec: String) throws -> BisFile {
        try WireFormatProbe.decode(BisFile.self, BisTestData.repo.appending(path: "data/builds/\(BisTestData.build)/bis/\(spec).json"))
    }

    private func realItems(_ classSlug: String) throws -> [Int: Item] {
        let file = try WireFormatProbe.decode(ItemsFile.self, BisTestData.repo.appending(path: "data/builds/\(BisTestData.build)/items/\(classSlug).json"))
        return Dictionary(file.items.map { ($0.id, $0) }, uniquingKeysWith: { first, _ in first })
    }

    private func furyItems() throws -> [Int: Item] {
        Dictionary(try BisTestData.items().items.map { ($0.id, $0) }, uniquingKeysWith: { first, _ in first })
    }

    private func request(_ file: BisFile, band: Int, preset: String, slotIndex: Int, wornID: Int? = nil) throws -> ItemTooltipRequest {
        let entry = try #require(file.band(for: BisRange(band: band, preset: preset), faction: "horde"))
        let slot = entry.slots[slotIndex]
        let id = try #require(slot.itemId)
        return ItemTooltipRequest(
            itemID: id,
            fallback: HomeItemLook(id: id, name: slot.itemName, quality: nil, icon: nil, itemLevel: nil),
            build: BisTestData.build, classSlug: "warrior", spec: file.spec, specName: "Fury Warrior",
            band: band, bandLabel: BisLabels.range(BisRange(band: band, preset: preset)), preset: preset, faction: "horde",
            sourceLine: slot.source, wornID: wornID
        )
    }

    @Test("the next best reads in the band's unit: DPS, HPS, score")
    func nextBestUnits() throws {
        let fury = try BisTestData.bis()
        let dps = ItemTooltipModel.build(try request(fury, band: 60, preset: "raid", slotIndex: 0), sources: ItemTooltipSources(items: try furyItems(), sets: [], bis: fury))
        #expect(dps.alternatives.first?.meta == "\u{2212}15.0 DPS vs the pick")
        for (spec, word) in [("priest-holy", "HPS"), ("warrior-protection", "score")] {
            let file = try real(spec)
            let index = try #require(file.band(for: BisRange(band: 60, preset: "raid"), faction: "horde")?.slots
                .firstIndex { ($0.alternatives ?? []).prefix(2).contains { abs($0.dpsDelta) >= 0.05 } })
            let model = ItemTooltipModel.build(try request(file, band: 60, preset: "raid", slotIndex: index), sources: ItemTooltipSources(items: try realItems(String(spec.prefix { $0 != "-" })), sets: [], bis: file))
            #expect(model.alternatives.contains { $0.meta.hasSuffix(" \(word) vs the pick") })
        }
    }

    @Test("the footer names the slot, spec and band, never a preset id")
    func footer() throws {
        let fury = try BisTestData.bis()
        let bare = ItemTooltipModel.build(try request(fury, band: 50, preset: "bare", slotIndex: 0), sources: ItemTooltipSources(items: try furyItems(), sets: [], bis: fury))
        #expect(bare.pickFooter == "Best in slot for Head · Fury Warrior 50 to 59")
        let raid = ItemTooltipModel.build(try request(fury, band: 60, preset: "raid", slotIndex: 10), sources: ItemTooltipSources(items: try furyItems(), sets: [], bis: fury))
        #expect(raid.pickFooter == "Best in slot for Finger 1 · Fury Warrior 60 · Phase 1 raid")
    }

    @Test("the worn runner-up leads with You wear this")
    func wornRunnerUp() throws {
        let fury = try BisTestData.bis()
        let alt = try #require(fury.band(for: BisRange(band: 60, preset: "raid"), faction: "horde")?.slots[0].alternatives?.first)
        let model = ItemTooltipModel.build(try request(fury, band: 60, preset: "raid", slotIndex: 0, wornID: alt.itemId), sources: ItemTooltipSources(items: try furyItems(), sets: [], bis: fury))
        #expect(model.alternatives.first?.meta.hasPrefix("You wear this · ") == true)
    }

    @Test("a tank list explains its score instead of a damage headline")
    func tankHeadline() throws {
        let file = try real("warrior-protection")
        let band = try #require(file.band(for: BisRange(band: 60, preset: "raid"), faction: "horde"))
        let list = BisListBuilder(items: [], worn: [], timeZone: .gmt).list(file: file, band: band, specLabel: "Protection Warrior", rangeLabel: "60")
        #expect(list.unit == .score)
        #expect(list.setDpsLine?.hasPrefix("Tank lists rank by score") == true)
        #expect(list.rows.compactMap(\.pick).contains { $0.next.hasSuffix(" score") })
    }

    @Test("the raid chip says when raids open, from the data, until they do")
    func raidsOpening() throws {
        let dates: [KeyDate] = try BisTestData.fixture("cdn.dates")
        let opening = try #require(RaidsOpening(dates))
        #expect(opening.label == "Dec 9")
        let raid = BisRange(band: 60, preset: "raid")
        #expect(BisLabels.range(raid, opening: opening, now: BisTestData.now) == "60 · Phase 1 raid · opens Dec 9")
        #expect(BisLabels.range(raid, opening: opening, now: opening.starts) == "60 · Phase 1 raid")
        #expect(BisLabels.range(BisRange(band: 60, preset: "bare"), opening: opening, now: BisTestData.now) == "60")
        #expect(RaidsOpening([]) == nil)
    }
}
